package service

import (
	"context"
	"kratos-use/api/mini/user/v1"
	api_user "kratos-use/api/mini/user/v1"
	"kratos-use/internal/biz"
	"kratos-use/internal/entity"
	"kratos-use/pkg/middleware"
)

type UserService struct {
	v1.UnimplementedUserServer

	userUseCase *biz.UserUsecase
}

func NewUserService(userUseCase *biz.UserUsecase) *UserService {
	return &UserService{
		userUseCase: userUseCase,
	}
}

func (s *UserService) GetLoginUserInfo(ctx context.Context, req *api_user.GetLoginUserInfoReq) (*v1.GetLoginUserInfoResp, error) {
	userId := middleware.GetUserId(ctx)
	userInfo, err := s.userUseCase.GetUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &v1.GetLoginUserInfoResp{
		Id:       userInfo.ID,
		UserName: userInfo.Name,
		Mobile:   userInfo.Mobile,
	}, nil
}

func (s *UserService) Register(ctx context.Context, req *api_user.RegisterReq) (*api_user.RegisterResp, error) {
	userId, err := s.userUseCase.Register(ctx, &entity.RegisterInput{
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &api_user.RegisterResp{UserId: userId}, nil
}
func (s *UserService) Login(ctx context.Context, req *api_user.LoginReq) (*api_user.LoginResp, error) {
	token, err := s.userUseCase.Login(ctx, &entity.LoginInput{
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &api_user.LoginResp{Token: token}, nil
}
