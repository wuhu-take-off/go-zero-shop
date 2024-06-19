package logic

import (
	"TongChi_shop/rpc/shop"
	json_number_transition "TongChi_shop/tool/json.number_transition"
	"TongChi_shop/tool/raw_field"
	"context"
	"errors"

	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsAddLogic {
	return &GoodsAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsAddLogic) GoodsAdd(req *types.GoodsAddReq) (resp *types.GoodsAddResp, err error) {
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}
	if !raw_field.CheckMaxMin(req) {
		return nil, errors.New("格式错误")
	}

	var list []*shop.SpecificationReq
	for _, s := range req.Specification {
		if s.Inventory < 0 || s.Score <= 0 {
			continue
		}

		list = append(list, &shop.SpecificationReq{
			Size:             s.Size,
			DisplayInventory: s.Inventory,
			RealInventory:    s.Inventory,
			Score:            s.Score,
		})
	}
	add, err := l.svcCtx.GoodsRpc.GoodsAdd(l.ctx, &shop.GoodsAddReq{
		UserId:            int32(userId),
		GoodsName:         req.GoodsName,
		GoodsTypeId:       req.GoodsType,
		Status:            req.Status,
		Description:       req.Description,
		SpecificationList: list,
	})
	if err != nil || !add.OK {
		return nil, err
	}
	return
}
