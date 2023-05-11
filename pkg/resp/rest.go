package resp

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err BaseError) {
	var body Body
	if err != NilBaseError {
		body.Code = err.Code
		body.Msg = err.Msg
	} else {
		body.Code = 200
		body.Msg = "Success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
