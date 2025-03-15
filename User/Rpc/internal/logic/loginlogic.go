package logic

import (
	"context"

	common "Bolog/Common"
	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResponse, error) {
	userbasic := &models.UserBasic{}
	result := l.svcCtx.DB.Where("user_name=? and pass_word=?", in.Username, common.MD5(in.Password)).Find(userbasic)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.LoginResponse{
		Result: 1,
		UserId: int32(userbasic.ID),
	}, nil
}
