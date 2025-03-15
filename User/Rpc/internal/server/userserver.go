// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"Bolog/User/Rpc/internal/logic"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) CreateUser(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserRespon, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServer) DeleteUser(ctx context.Context, in *user.DeleteUserReq) (*user.DeleteUserResponse, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *UserServer) UpdateUser(ctx context.Context, in *user.UpdateUserReq) (*user.UpdateUserRespon, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserServer) UpdateUserPassWord(ctx context.Context, in *user.UpdateUserPassWordReq) (*user.UpdateUserPassWordResponse, error) {
	l := logic.NewUpdateUserPassWordLogic(ctx, s.svcCtx)
	return l.UpdateUserPassWord(in)
}

func (s *UserServer) GetUserInfoById(ctx context.Context, in *user.GetUserInfoReq) (*user.GetUserInfoResponse, error) {
	l := logic.NewGetUserInfoByIdLogic(ctx, s.svcCtx)
	return l.GetUserInfoById(in)
}

func (s *UserServer) GetUserListPage(ctx context.Context, in *user.GetUserListPageReq) (*user.GetUserListPageResponse, error) {
	l := logic.NewGetUserListPageLogic(ctx, s.svcCtx)
	return l.GetUserListPage(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
