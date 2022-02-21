package logic

import (
	"context"
	"github.com/acger/user-svc/model"

	"github.com/acger/user-svc/internal/svc"
	"github.com/acger/user-svc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *user.UserListReq) (*user.UserListRsp, error) {

	var u []model.User
	m := l.svcCtx.DB

	r := m.Find(&u, in.Id)

	if r.Error != nil {
		return &user.UserListRsp{
			Code: 0,
		}, nil
	}

	var list []*user.UserInfo

	for _, item := range u {
		list = append(list, &user.UserInfo{
			Id:      uint64(item.ID),
			Account: item.Account,
			Name:    item.Name,
			Avatar:  item.Avatar,
			Status:  item.Status,
		})
	}

	return &user.UserListRsp{
		Code: 0,
		User: list,
	}, nil
}
