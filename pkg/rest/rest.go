package rest

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		base, ok := err.(BaseError)
		if ok {
			body.Code = base.Code
			body.Msg = base.Msg
		} else {
			body.Code = BizError
			body.Msg = err.Error()
		}
	} else {
		body.Code = 200
		body.Msg = "Success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
