package logic

import (
	"context"
	"go-zero-template/greet/common/errorx"
	"go-zero-template/greet/internal/svc"
	"go-zero-template/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetNameLogic {
	return &GreetNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetNameLogic) GreetName(req *types.Request) (resp *types.Response, err error) {
	logx.Info("req==", req)
	return nil, errorx.NewUserNotFoundError("用户不存在")
}
