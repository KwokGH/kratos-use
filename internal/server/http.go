package server

import (
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	api_common "kratos-use/api/common"
	api_diary "kratos-use/api/mini/diary/v1"
	api_user "kratos-use/api/mini/user/v1"
	"kratos-use/internal/conf"
	"kratos-use/internal/service"
	"kratos-use/pkg/middleware"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	bootstrap *conf.Bootstrap,
	logger log.Logger,
	commonService *service.CommonService,
	userService *service.UserService,
	diaryService *service.DiaryService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.Tracer(),
			logging.Server(logger),
			validate.Validator(),
			selector.Server(
				jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
					return []byte(bootstrap.App.Auth.AccessSecret), nil
				})).Match(middleware.NewWhiteListMatcher(whiteList)).Build(),
			middleware.FillInfo(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}

	c := bootstrap.Server

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Http.Timeout)*time.Second))
	}
	srv := http.NewServer(opts...)

	api_common.RegisterCommonHTTPServer(srv, commonService)
	api_user.RegisterUserHTTPServer(srv, userService)
	api_diary.RegisterDiaryHTTPServer(srv, diaryService)

	return srv
}

var whiteList map[string]struct{}

func init() {
	whiteList = make(map[string]struct{})
	whiteList[api_common.OperationCommonPing] = struct{}{}
	whiteList[api_user.OperationUserLogin] = struct{}{}
	whiteList[api_user.OperationUserRegister] = struct{}{}
}
