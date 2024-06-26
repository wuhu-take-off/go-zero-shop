package logic

import (
	"TongChi_shop/rpc/shop"
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
	img, err := l.svcCtx.FunctionRpc.LoadImg(l.ctx, &shop.LoadImgReq{
		GoodsId: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.LoadImgResp)
	for i := 0; i < len(img.Imgs); i++ {
		t := types.Img{
			Data: img.Imgs[i].Img,
		}
		resp.Imgs = append(resp.Imgs, &t)
	}

	return
}
