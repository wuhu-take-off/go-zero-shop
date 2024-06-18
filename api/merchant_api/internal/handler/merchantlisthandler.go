package handler

import (
	"net/http"

	"TongChi_shop/api/merchant_api/internal/logic"
	"TongChi_shop/api/merchant_api/internal/svc"
	"TongChi_shop/api/merchant_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MerchantListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMerchantListLogic(r.Context(), svcCtx)
		resp, err := l.MerchantList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
