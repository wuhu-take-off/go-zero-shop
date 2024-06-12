package functionservelogic

import (
	"context"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadImgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadImgLogic {
	return &LoadImgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoadImgLogic) LoadImg(in *shop.LoadImgReq) (*shop.LoadImgResp, error) {
	// todo: add your logic here and delete this line

	return &shop.LoadImgResp{}, nil
}
