package logic

import (
	"context"

	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserRespon, error) {
	usermodel := new(models.UserBasic)
	usermodel.ID = uint(in.Id)
	usermodel.UserName = in.UserName
	usermodel.Email = in.Email
	usermodel.Telephone = in.Telephone
	usermodel.Age = int(in.Age)
	usermodel.Gender = int(in.Gender)
	result := l.svcCtx.DB.Save(usermodel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.UpdateUserRespon{
		Result: 1,
	}, nil
}
