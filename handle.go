package main

import (
	"context"

	"kitex-multi-protocol/kitex_gen/hello/example/helloservice"
)

type HelloServiceImpl struct{}

func (s *HelloServiceImpl) SayHello(ctx context.Context, req *helloservice.HelloReq) (*helloservice.HelloResp, error) {
	return &helloservice.HelloResp{
		Message: "Hello, " + req.Name,
	}, nil
}
