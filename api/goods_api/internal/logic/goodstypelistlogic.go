package logic

import (
	"context"

	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsTypeListLogic {
	return &GoodsTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsTypeListLogic) GoodsTypeList(req *types.GoodsTypeListReq) (resp *types.GoodsTypeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
