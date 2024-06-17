// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"TongChi_shop/api/task_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthentication},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/TaskUpdate",
					Handler: TaskUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/taskList",
					Handler: TaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/taskType",
					Handler: TaskTypeHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/task"),
	)
}