package handler

import (
	"net/http"

	"TongChi_shop/api/goods_api/internal/logic"
	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GoodsUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsUpDateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGoodsUpdateLogic(r.Context(), svcCtx)
		resp, err := l.GoodsUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
