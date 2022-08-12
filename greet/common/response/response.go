package response

import (
	"go-zero-template/greet/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var code int
	switch e := err.(type) {
	case *errorx.CodeError:
		code = e.ErrCode()
	default:
		code = http.StatusInternalServerError
	}
	var body Body
	if err != nil {
		body.Code = code
		body.Msg = err.Error()
	} else {
		body.Msg = "ok"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
