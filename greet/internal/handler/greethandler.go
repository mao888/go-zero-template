package handler

import (
	"go-zero-template/greet/common/response"
	"go-zero-template/greet/internal/logic"
	"go-zero-template/greet/internal/svc"
	"go-zero-template/greet/internal/types"
	"go-zero-template/greet/internal/validator"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := validator.ValidateGreet(req); err != nil {
			response.Response(w, nil, err)
			return
		}
		l := logic.NewGreetLogic(r.Context(), ctx)
		resp, err := l.Greet(&req)
		response.Response(w, resp, err)

	}
}
