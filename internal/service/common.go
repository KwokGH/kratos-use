package service

import (
	"context"
	"time"

	pb "kratos-use/api/common"

	"github.com/go-kratos/kratos/v2"
)

type CommonService struct {
	//pb.UnimplementedCommonServer
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (s *CommonService) Ping(ctx context.Context, req *pb.Empty) (*pb.PingResp, error) {
	resp := &pb.PingResp{
		Version: "未知",
		Name:    "未知",
		Time:    time.Now().String(),
	}
	appInfo, ok := kratos.FromContext(ctx)
	if ok {
		resp.Version = appInfo.Version()
		resp.Name = appInfo.Name()
	}
	return resp, nil // 返回pong
}
