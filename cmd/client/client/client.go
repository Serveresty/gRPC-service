package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func Interceptor(
	ctx context.Context,
	method string,
	req,
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	log.Println("--> unary client interceptor: ", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Println(reply)
	return err
}
