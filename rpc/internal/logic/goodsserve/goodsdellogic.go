package goodsservelogic

import (
	"TongChi_shop/model/goods_model"
	"context"
	"errors"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsDelLogic {
	return &GoodsDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商品
func (l *GoodsDelLogic) GoodsDel(in *shop.GoodsDelReq) (*shop.GoodsDelResp, error) {
	// todo: add your logic here and delete this line
	goodsModel := goods_model.NewGoodsInfosModel(l.svcCtx.DB)
	if err := goodsModel.DelGoodsInfos(in.GoodsId); err != nil {
		return nil, errors.New("删除失败")
	}
	return &shop.GoodsDelResp{
		OK: true,
	}, nil
}
