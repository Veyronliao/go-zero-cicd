package logic

import (
	"context"

	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIdLogic) GetUserInfoById(in *user.GetUserInfoReq) (*user.GetUserInfoResponse, error) {
	userbasic := &models.UserBasic{}
	result := l.svcCtx.DB.First(userbasic, uint(in.Id))
	if result.Error != nil {
		return nil, result.Error
	}
	data := user.PageListRespData{
		Id:       int32(userbasic.ID),
		UserName: userbasic.UserName,
		//PassWord:   userbasic.PassWord,
		Gender:    int32(userbasic.Gender),
		Age:       int32(userbasic.Age),
		Email:     userbasic.Email,
		Telephone: userbasic.Telephone,
		//CreateTime: userbasic.CreatedAt.Format("2006-01-02 15:04:05"),
		CreateTime: userbasic.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	response := user.GetUserInfoResponse{
		Data: &data,
	}
	return &response, nil
}
