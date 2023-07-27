package casbin

import (
	"fmt"
	"github.com/CloverOS/go-zero-pkg/pkg/redis"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	redisv9 "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
)

const Prefix = "role"

type Casbin struct {
	syncedEnforcer *casbin.SyncedEnforcer
}

func NewGormCasbin(gormDb *gorm.DB, watch *redis.Config) *Casbin {
	rbac := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := model.NewModelFromString(rbac)
	if err != nil {
		panic("casbin model router failed: " + err.Error())
	}
	if gormDb == nil {
		panic("gorm db is nil")
	}
	a, _ := gormadapter.NewAdapterByDBUseTableName(gormDb, "role", "api_resource")
	syncedEnforcer, _ := casbin.NewSyncedEnforcer(m, a)
	w, _ := rediswatcher.NewWatcher(watch.Addr, rediswatcher.WatcherOptions{
		Options: redisv9.Options{
			Network:  "tcp",
			Password: watch.Password,
		},
		Channel: "/casbin",
		// Only exists in test, generally be true
		IgnoreSelf: false,
	})
	err = syncedEnforcer.SetWatcher(w)
	if err != nil {
		panic("casbin set watcher failed: " + err.Error())
	}
	err = w.SetUpdateCallback(func(s string) {
		fmt.Println("------>update casbin :" + s)
	})
	if err != nil {
		panic("casbin update callback failed: " + err.Error())
	}
	return &Casbin{syncedEnforcer: syncedEnforcer}
}

func (c *Casbin) ClearPermission(roleId int) (bool, error) {
	_, err := c.syncedEnforcer.DeletePermissionsForUser(Prefix + strconv.Itoa(roleId))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Casbin) Verify(roleId int, path string, method string) (success bool) {
	success, _ = c.syncedEnforcer.Enforce(Prefix+strconv.Itoa(roleId), path, method)
	return success
}

func (c *Casbin) AddPermissions(roleId int, apiResource []ApiResource) (bool, error) {
	_, err := c.ClearPermission(roleId)
	if err != nil {
		return false, err
	}
	var permissions [][]string
	for _, v := range apiResource {
		permissions = append(permissions, []string{v.GetPath(), v.GetMethod()})
	}
	success, err := c.syncedEnforcer.AddPermissionsForUser(Prefix+strconv.Itoa(roleId), permissions...)
	if err != nil {
		return false, err
	}
	return success, err
}
