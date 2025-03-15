package handler

import (
	"net/http"

	"Bolog/User/Api/internal/logic"
	"Bolog/User/Api/internal/svc"
	"Bolog/User/Api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserListPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserListPageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetUserListPageLogic(r.Context(), svcCtx)
		resp, err := l.GetUserListPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
