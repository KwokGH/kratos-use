package biz

import (
	"context"
	"errors"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUserUsecase,
	NewDiaryUsecase,
)

// Transaction 新增事务接口方法
type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}

var (
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
)
