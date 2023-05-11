package rest

const (
	BizError = iota + 10001
	ParseError
	MysqlError
	RedisError
	TokenError
)

type BaseError struct {
	Msg  string `json:"msg"`  //错误信息
	Code int    `json:"code"` //错误码
	Err  error
}

func (b BaseError) Error() string {
	return b.Err.Error()
}

func MustDefaultError(err error) BaseError {
	return BaseError{
		Msg:  err.Error(),
		Code: BizError,
		Err:  err,
	}
}

func MustError(code int, msg string, err error) BaseError {
	if msg == "" {
		msg = err.Error()
	}
	return BaseError{
		Msg:  msg,
		Code: code,
		Err:  err,
	}
}

func MustBizError(msg string, err error) BaseError {
	return MustError(BizError, msg, err)
}

func MustMysqlError(msg string, err error) BaseError {
	return MustError(MysqlError, msg, err)
}

func MustRedisError(msg string, err error) BaseError {
	return MustError(RedisError, msg, err)
}

func MustParseError(msg string, err error) BaseError {
	return MustError(ParseError, msg, err)
}

func MustTokenError(msg string, err error) BaseError {
	return MustError(TokenError, msg, err)
}
