// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"Bolog/User/Api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/delete",
				Handler: DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/userpage",
				Handler: GetUserListPageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: UserLoginHandler(serverCtx),
			},
		},
	)
}
