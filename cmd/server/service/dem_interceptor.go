package service

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> unary server interceptor: ", info.FullMethod)
	if info.FullMethod != "/api.AuthService/Login" {
		log.Println(req)
	}
	return handler(ctx, req)
}
