package logic

import (
	"context"

	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadimgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoadimgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadimgLogic {
	return &LoadimgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoadimgLogic) Loadimg(req *types.LoadImgReq) (resp *types.LoadImgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
