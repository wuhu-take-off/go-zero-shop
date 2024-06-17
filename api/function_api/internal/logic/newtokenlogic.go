package logic

import (
	"context"

	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewTokenLogic {
	return &NewTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewTokenLogic) NewToken(req *types.NewTokenReq) (resp *types.NewTokenResp, err error) {
	// todo: add your logic here and delete this line

	return
}
