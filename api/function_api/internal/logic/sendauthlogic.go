package logic

import (
	"TongChi_shop/rpc/shop"
	"context"
	"errors"

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

func (l *SendAuthLogic) SendAuth(req *types.SendAuthReq) (*types.SendAuthResp, error) {
	resp, err := l.svcCtx.FunctionRpc.SendEmailAuth(l.ctx, &shop.SendEmailAuthReq{
		Email: req.Email,
	})
	if err != nil || resp.OK == 0 {
		return nil, errors.New("发送失败")
	}
	return &types.SendAuthResp{
		OK: true,
	}, nil
}
