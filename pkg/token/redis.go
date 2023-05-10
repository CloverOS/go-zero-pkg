package token

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/matoous/go-nanoid/v2"
	"time"
)

type Token[T any] struct {
	db *redis.Client
}

const (
	Prefix = "Auth:"
)

func NewTokenService[T any](db *redis.Client) Token[T] {
	return Token[T]{db: db}
}

func (t *Token[T]) CreateToken(data interface{}, expiration time.Duration, roleId ...string) (string, error) {
	key, err := gonanoid.New()
	marshal, err := json.Marshal(data)
	if err != nil {
		return "", errors.New("json marshal failed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	set, err := t.db.SetNX(ctx, Prefix+key, marshal, expiration).Result()
	if err != nil || !set {
		return "", errors.New("create failed！")
	}
	if len(roleId) > 1 && roleId[0] != "" && roleId[1] != "" {
		set, err = t.db.SetNX(ctx, Prefix+roleId[0]+":"+roleId[1]+":"+key, key, expiration).Result()
	}
	if err != nil || !set {
		return "", errors.New("create failed！")
	}
	return key, nil
}

func (t *Token[T]) UpdateTokenWithoutRefreshExpireTime(key string, data interface{}) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return errors.New("json marshal failed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = t.db.Set(ctx, Prefix+key, marshal, redis.KeepTTL).Result()
	return err
}

func (t *Token[T]) DeleteToken(key ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := t.db.Del(ctx, key...).Result()
	return err
}

func (t *Token[T]) ClearAllLogin(prefix string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.db.Keys(ctx, Prefix+prefix+":*").Result()
	if err != nil {
		return err
	}
	for _, v := range result {
		key, err := t.db.Get(ctx, v).Result()
		if err != nil {
			return err
		}
		err = t.DeleteToken(Prefix+key, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Token[T]) GetTokenByRoleId(roleId string) ([]string, error) {
	var tokens []string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := t.db.Keys(ctx, Prefix+roleId+":*").Result()
	if err != nil {
		return tokens, err
	}
	for i := range result {
		strings, err := t.db.Get(ctx, result[i]).Result()
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, strings)
	}
	return tokens, nil
}

func (t *Token[T]) GetDataListFromRoleId(roleId string) ([]T, error) {
	var data []T
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.db.Keys(ctx, Prefix+roleId+":*").Result()
	if err != nil {
		return data, err
	}
	for _, v := range result {
		key, err := t.db.Get(ctx, v).Result()
		if err != nil {
			return data, err
		}
		fromToken, err := t.GetDataFromToken(key)
		if err == nil {
			data = append(data, fromToken)
		}
	}
	return data, err
}

func (t *Token[T]) GetDataFromToken(token string) (T, error) {
	var data T
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := t.db.Get(ctx, Prefix+token).Result()
	if err != nil {
		return data, err
	}
	e := json.Unmarshal([]byte(result), &data)
	if e != nil {
		return data, e
	}
	return data, err
}
