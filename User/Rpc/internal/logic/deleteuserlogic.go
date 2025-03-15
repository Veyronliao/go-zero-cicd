package logic

import (
	"context"

	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserReq) (*user.DeleteUserResponse, error) {
	deleteuser := new(models.UserBasic)
	deleteuser.ID = uint(in.Id)
	result := l.svcCtx.DB.Delete(deleteuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.DeleteUserResponse{
		Result: 1,
	}, nil
}
