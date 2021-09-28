package logic

import (
	"context"
	"errors"
	"github.com/acger/user-svc/common"
	"github.com/acger/user-svc/model"
	"gorm.io/gorm"

	"github.com/acger/user-svc/internal/svc"
	"github.com/acger/user-svc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoRsp, error) {
	var u model.User
	var r *gorm.DB

	if len(in.Account) > 0 {
		r = l.svcCtx.DB.Where("account = ? ", in.Account).First(&u)
	} else {
		r = l.svcCtx.DB.First(&u, in.Id)
	}

	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return &user.UserInfoRsp{
			Code:    10003,
			Message: common.ErrorCode[10003],
		}, nil
	}

	if r.Error != nil {
		return &user.UserInfoRsp{
			Code:    10001,
			Message: common.ErrorCode[10001],
		}, nil
	}

	return &user.UserInfoRsp{
		Code: 0,
		User: &user.UserDetail{
			Id:       uint64(u.ID),
			Account:  u.Account,
			Name:     u.Name,
			Avatar:   u.Avatar,
			Status:   u.Status,
			Password: u.Password,
		},
	}, nil
}
