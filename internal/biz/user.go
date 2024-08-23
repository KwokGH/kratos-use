package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	api_common "kratos-use/api/common"
	"kratos-use/internal/conf"
	"kratos-use/pkg/jwtx"
	"time"

	"kratos-use/internal/entity"
	"kratos-use/pkg/unique"
)

// UserRepo is a Greater repo.
type UserRepo interface {
	CreateForRegister(ctx context.Context, input *entity.RegisterInput) (string, error)
	GetUserById(ctx context.Context, id string) (*entity.UserDTO, error)
	GetUserByAccount(ctx context.Context, id string) (*entity.UserDTO, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo      UserRepo
	log       *log.Helper
	tx        Transaction
	bootstrap *conf.Bootstrap
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, tx Transaction, bootstrap *conf.Bootstrap, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, tx: tx, bootstrap: bootstrap, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, input *entity.RegisterInput) (string, error) {
	userId, err := uc.repo.CreateForRegister(ctx, input)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (uc *UserUsecase) Login(ctx context.Context, input *entity.LoginInput) (string, error) {
	userInfo, err := uc.repo.GetUserByAccount(ctx, input.Account)
	if err != nil {
		return "", err
	}
	// 验证用户密码
	md5Password := unique.GetMd5(input.Password + userInfo.PasswordSalt)
	if md5Password != userInfo.Password {
		uc.log.WithContext(ctx).Warnw("msg", "密码不正确", "userId", userInfo.ID)
		return "", api_common.ErrorBadRequest("用户名或密码错误")
	}

	// 构造token
	auth := uc.bootstrap.App.Auth
	token, err := jwtx.CreateJwtToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, map[string]interface{}{
		"userId": userInfo.ID,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, id string) (*entity.UserDTO, error) {
	return uc.repo.GetUserById(ctx, id)
}
