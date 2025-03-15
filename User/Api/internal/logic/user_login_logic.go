package logic

import (
	common "Bolog/Common"
	"Bolog/User/Api/internal/svc"
	"Bolog/User/Api/internal/types"
	"Bolog/User/Rpc/userclient"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp common.HttpResponse, err error) {
	loginModel := userclient.LoginReq{
		Username: req.UserName,
		Password: common.MD5(req.PassWord),
	}
	data, err := l.svcCtx.UserRpc.Login(l.ctx, &loginModel)
	if err != nil {
		return common.FailureResponse(nil, err.Error(), 0), nil
	}
	expire := l.svcCtx.Config.Auth.AccessExpire
	token, _ := common.GenerateToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), expire, int64(data.UserId))
	return common.SuccessResponse(token, ""), nil
}
