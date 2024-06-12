package goodsservelogic

import (
	"TongChi_shop/model/specification_model"
	"TongChi_shop/model/user_model"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
	"TongChi_shop/tool/raw_field"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GoodsList 获取商品列表
func (l *GoodsListLogic) GoodsList(in *shop.GoodsListReq) (*shop.GoodsListResp, error) {
	// todo: add your logic here and delete this line
	userModel := user_model.NewUserInfosModel(l.svcCtx.DB)
	role := userModel.UserRole(in.UserId)
	if role == 0 {
		return nil, errors.New("用户不存在")
	}
	sql, _ := raw_field.UpdateFieldMap(in, user_model.UserGoodsSqlFileMap)
	for i := 0; i < len(sql); i++ {
		sql[i] = sql[i] + " = ? "
	}
	//获取用户的商户以及商户对应的商品信息
	count, goods := userModel.FindUsersGoods(in.UserId, in.Page, in.Limit, "")

	goodsIdList := make([]interface{}, 0, len(goods))
	//将对应的规格信息放到对应的商品中同时加载goodsId列表
	goodsListMap := make(map[int32]*[]specification_model.SpecificationInfos)
	for i := 0; i < len(goods); i++ {
		goodsListMap[goods[i].GoodsId] = &goods[i].Specifications
		goodsIdList = append(goodsIdList, goods[i].GoodsId)
	}

	//获取各个商品对应的规格信息
	specificationModel := specification_model.NewSpecificationModel(l.svcCtx.DB)
	specifications := specificationModel.FindSpecification(goodsIdList...)

	for _, infos := range specifications {
		if specification, ok := goodsListMap[infos.GoodsId]; ok {
			*specification = append(*specification, infos)
		}
	}

	resp := &shop.GoodsListResp{
		List:   make([]*shop.GoodsList, 0, len(goods)),
		Count:  count,
		Limit:  in.Limit,
		Number: (count + in.Limit - 1) / in.Limit,
		Page:   in.Page,
	}

	for _, good := range goods {
		resp.List = append(resp.List, good.ToGoodsListResp())
	}
	return resp, nil
}
