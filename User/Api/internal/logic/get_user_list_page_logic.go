package logic

import (
	"context"

	"Bolog/User/Api/internal/svc"
	"Bolog/User/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListPageLogic {
	return &GetUserListPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListPageLogic) GetUserListPage(req *types.GetUserListPageRequest) (resp *types.GetUserListPageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
