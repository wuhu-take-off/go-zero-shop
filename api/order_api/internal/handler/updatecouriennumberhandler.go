package handler

import (
	"net/http"

	"TongChi_shop/api/order_api/internal/logic"
	"TongChi_shop/api/order_api/internal/svc"
	"TongChi_shop/api/order_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateCourienNumberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCourierNumberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateCourienNumberLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCourienNumber(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
