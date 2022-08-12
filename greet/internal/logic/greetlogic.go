package logic

import (
	"context"
	"fmt"
	"go-zero-template/greet/internal/models"

	"go-zero-template/greet/internal/svc"
	"go-zero-template/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	var users []models.BusinessUser
	if err := l.svcCtx.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	fmt.Println(users)
	return &types.Response{Message: users[0].BifUserBid}, nil
}
