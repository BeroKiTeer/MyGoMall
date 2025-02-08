package main

import (
	"context"
	"user/biz/service"
	user "user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	resp, err = service.NewLogoutService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	resp, err = service.NewUpdateUserService(ctx).Run(req)

	return resp, err
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	resp, err = service.NewGetUserInfoService(ctx).Run(req)

	return resp, err
}
