package logic

import (
	"context"

	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendAuthLogic {
	return &SendAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendAuthLogic) SendAuth(req *types.SendAuthReq) (resp *types.SendAuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
