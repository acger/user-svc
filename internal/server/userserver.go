// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/acger/user-svc/internal/logic"
	"github.com/acger/user-svc/internal/svc"
	"github.com/acger/user-svc/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) UserAdd(ctx context.Context, in *user.UserAddReq) (*user.UserAddRsp, error) {
	l := logic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

func (s *UserServer) UserUpdate(ctx context.Context, in *user.UserUpdateReq) (*user.Response, error) {
	l := logic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoReq) (*user.UserInfoRsp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) UserList(ctx context.Context, in *user.UserListReq) (*user.UserListRsp, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}
