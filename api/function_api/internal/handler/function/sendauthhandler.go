package function

import (
	"net/http"

	"TongChi_shop/api/function_api/internal/logic/function"
	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/function_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := function.NewSendAuthLogic(r.Context(), svcCtx)
		resp, err := l.SendAuth(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
