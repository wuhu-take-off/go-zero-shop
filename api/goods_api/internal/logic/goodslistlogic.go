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

type GoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsListLogic) GoodsList(req *types.GoodsListReq) (resp *types.GoodsListResp, err error) {
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}
	list, err := l.svcCtx.GoodsRpc.GoodsList(l.ctx, &shop.GoodsListReq{
		UserId:      int32(userId),
		Status:      req.Status,
		GoodsTypeId: req.GoodsTypeId,
		MerchantId:  req.MerchantId,
		GoodsName:   req.GoodsName,
		GoodsId:     req.GoodsID,
		Page:        req.Page,
		Limit:       req.Limit,
	})
	if err != nil {
		return nil, err
	}
	var goodsList []*types.GoodsList
	for _, g := range list.List {
		var specificationList []*types.SpecificationList
		for _, specificationResp := range g.SpecificationList {
			specificationList = append(specificationList, &types.SpecificationList{
				SpecificationId:  specificationResp.SpecificationId,
				Size:             specificationResp.Size,
				DisplayInventory: specificationResp.DisplayInventory,
				RealInventory:    specificationResp.RealInventory,
				Score:            specificationResp.Score,
			})
		}

		goodsList = append(goodsList, &types.GoodsList{
			MerchantName: g.MerchantName,
			GoodsId:      g.GoodsId,
			GoodsName:    g.GoodsName,
			GoodsTypeId:  g.GoodsTypeId,
			Status:       g.Status,
			List:         specificationList,
		})
	}
	return &types.GoodsListResp{
		Count:  list.Count,
		Limit:  list.Limit,
		Number: list.Number,
		Page:   list.Page,
		List:   goodsList,
	}, nil
}
