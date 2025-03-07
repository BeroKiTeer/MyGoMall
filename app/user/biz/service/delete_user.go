package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"user/biz/dal/mysql"
	"user/biz/model"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	// Finish your business logic.
	//查询用户id是否有效
	if req.UserId < 0 {
		return nil, errors.New("无效的用户ID！")
	}

	exist, err := model.UserExistsByID(mysql.DB, req.UserId)
	if exist == false {
		klog.Error("用户不存在！")
		return nil, errors.New("用户不存在！")
	} else if err != nil {
		klog.Error("查询用户是否存在时出错！", err)
		return nil, err
	}

	//若存在，获取该用户对象
	DeleteUser, err := model.GetUserById(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	//删除用户
	err = model.DeleteUser(mysql.DB, DeleteUser)

	if err != nil {
		klog.Error("删除用户失败！", err)
		return nil, err
	}
	Success := err == nil

	return &user.DeleteUserResp{Success: Success}, err
}
