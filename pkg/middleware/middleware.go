package middleware

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func NewWhiteListMatcher(whiteList map[string]struct{}) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

const (
	UserId = "USER-ID"
)

func FillInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			FillUserIdFromCtx(ctx)

			return handler(ctx, req)
		}
	}
}

func FillUserIdFromCtx(ctx context.Context) string {
	token, ok := jwt.FromContext(ctx)
	if ok {
		if claims, ok := token.(jwtv5.MapClaims); ok {
			if userId, ok := claims[UserId]; ok {
				ctx = context.WithValue(ctx, UserId, userId)
				return userId.(string)
			}

		} else {
			fmt.Println("匹配登录信息失败")
		}
	}

	return ""
}

func GetUserId(ctx context.Context) string {
	return ctx.Value(UserId).(string)
}
