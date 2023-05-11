package resp

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

func MustError(msg string, err error) BaseError {
	return BaseError{
		Msg:  msg,
		Code: BizError,
		Err:  err,
	}
}

func MustBizError(code int, msg string, err error) BaseError {
	return BaseError{
		Msg:  msg,
		Code: code,
		Err:  err,
	}
}
