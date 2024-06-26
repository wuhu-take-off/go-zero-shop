package logic

import (
	"TongChi_shop/rpc/shop"
	json_number_transition "TongChi_shop/tool/json.number_transition"
	"context"
	"errors"

	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsUpdateLogic {
	return &GoodsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsUpdateLogic) GoodsUpdate(req *types.GoodsUpDateReq) (resp *types.GoodsUpDateResp, err error) {
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}

	var SpecificationList []*shop.SpecificationResp
	for _, list := range req.List {
		SpecificationList = append(SpecificationList, &shop.SpecificationResp{
			SpecificationId:  list.SpecificationId,
			Size:             list.Size,
			DisplayInventory: list.DisplayInventory,
			RealInventory:    list.RealInventory,
			Score:            list.Score,
		})
	}
	res, err := l.svcCtx.GoodsRpc.GoodsUpdate(l.ctx, &shop.GoodsUpdateReq{
		UserId:            int32(userId),
		GoodsId:           req.GoodsId,
		GoodsName:         req.GoodsName,
		GoodsTypeId:       req.GoodsTypeId,
		Status:            req.Status,
		Describe:          req.Description,
		SpecificationList: SpecificationList,
	})
	if err != nil || !res.OK {
		return nil, err
	}
	return
}
