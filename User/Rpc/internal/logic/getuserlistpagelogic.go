package logic

import (
	"context"

	common "Bolog/Common"
	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/svc"
	"Bolog/User/Rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListPageLogic {
	return &GetUserListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListPageLogic) GetUserListPage(in *user.GetUserListPageReq) (*user.GetUserListPageResponse, error) {
	pageSize := common.IF(in.Pagesize == 0, 20, in.Pagesize).(int)
	offset := common.IF(in.Pageindex == 0, 0, (in.Pageindex-1)*in.Pagesize).(int)
	var count int64
	tx := l.svcCtx.DB.Model(new(models.UserBasic)).Select("Id,Name,PassWord,Gender,Email,Telephone,CreatedAt,UpdatedAt,DeletedAt")
	if in.Selectoption.UserName != "" {
		tx.Where("UserName like ?", "%"+in.Selectoption.UserName+"%")
	}
	if in.Selectoption.Email != "" {
		tx.Where("Email like ?", "%"+in.Selectoption.Email+"%")
	}
	if in.Selectoption.Telephone != "" {
		tx.Where("Telephone like ?", "%"+in.Selectoption.Telephone+"%")
	}
	if in.Selectoption.Gender >= 0 {
		tx.Where("Gender = ?", in.Selectoption.Gender)
	}
	data := make([]*models.UserBasic, 0)
	err := tx.Count(&count).Limit(pageSize).Offset(offset).Find(data).Error
	if err != nil {
		logx.Error("[db error:]", err)
		return nil, err
	}
	resp := new(user.GetUserListPageResponse)
	list := make([]*user.PageListRespData, 0)
	for _, v := range list {
		item := new(user.PageListRespData)
		item.Id = v.Id
		item.Age = v.Age
		item.CreateTime = v.CreateTime
		item.Email = v.Email
		item.Telephone = v.Telephone
		item.UserName = v.UserName
		item.Gender = v.Gender
		list = append(list, item)
	}
	resp.Data = list
	resp.Count = int32(count)
	return resp, nil
}
