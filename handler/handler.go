package handler

import (
	"context"
	"github.com/liuyp5181/base/config"
	"github.com/liuyp5181/base/log"
	pb "github.com/liuyp5181/template/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceImpl struct {
	pb.UnimplementedGreeterServer
}

func (s *ServiceImpl) HelloWorld(ctx context.Context, req *pb.HelloWorldReq) (resp *pb.HelloWorldRes, err error) {
	log.Info(req, config.ServiceName)
	if len(req.Question) == 0 {
		return nil, status.Errorf(codes.Internal, "req is nil")
	}
	resp = &pb.HelloWorldRes{Type: pb.Type_ONE, Answer: "hi"}
	return resp, nil
}
