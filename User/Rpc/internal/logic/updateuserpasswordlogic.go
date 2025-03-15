package logic

import (
	"context"
	"fmt"

	common "Bolog/Common"
	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPassWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPassWordLogic {
	return &UpdateUserPassWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserPassWordLogic) UpdateUserPassWord(in *user.UpdateUserPassWordReq) (*user.UpdateUserPassWordResponse, error) {
	updateuser := new(models.UserBasic)
	updateuser.ID = uint(in.Id)
	updateuser.PassWord = common.MD5(in.PassWord)
	fmt.Println(updateuser.PassWord)
	result := l.svcCtx.DB.Model(updateuser).Update("pass_word", updateuser.PassWord)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.UpdateUserPassWordResponse{
		Result: 1,
	}, nil
}
