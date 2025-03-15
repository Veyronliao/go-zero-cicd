package logic

import (
	"context"

	common "Bolog/Common"
	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserRespon, error) {
	usermodel := models.UserBasic{
		UserName:  in.UserName,
		PassWord:  common.MD5(in.PassWord),
		Gender:    int(in.Gender),
		Email:     in.Email,
		Telephone: in.Telephone,
		Age:       int(in.Age),
	}
	result := l.svcCtx.DB.Create(&usermodel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.CreateUserRespon{
		Result: 1,
	}, nil
}
