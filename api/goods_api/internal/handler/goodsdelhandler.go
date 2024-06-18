package handler

import (
	"net/http"

	"TongChi_shop/api/goods_api/internal/logic"
	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GoodsDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGoodsDelLogic(r.Context(), svcCtx)
		resp, err := l.GoodsDel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
