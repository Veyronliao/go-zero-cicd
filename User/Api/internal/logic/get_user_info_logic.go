package logic

import (
	"context"
	"fmt"

	common "Bolog/Common"
	"Bolog/User/Api/internal/svc"
	"Bolog/User/Api/internal/types"
	"Bolog/User/Rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp common.HttpResponse, err error) {
	// todo: add your logic here and delete this line
	//id := l.ctx.Value("UserId").(string)
	id := int32(req.Id)
	fmt.Println("id:?", req.Id)
	//userid, _ := strconv.ParseInt(id, 10, 64)
	userinfo, err := l.svcCtx.UserRpc.GetUserInfoById(l.ctx, &userclient.GetUserInfoReq{
		Id: id,
	})
	fmt.Println(userinfo)
	if err != nil {
		return common.FailureResponse("", err.Error(), 0), err
	}
	return common.SuccessResponse(userinfo.Data, "success"), nil
}
