package client

import (
	"context"
	"proteitestcase/cmd/server/service"
	"proteitestcase/logger"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Interceptor(
	ctx context.Context,
	method string,
	req,
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	lg := logger.GRPCLogger()

	startTime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := lg.Info()
	if err != nil {
		logger = lg.Error().Err(err)
	}

	logger.Str("protocol", "grpc").
		Int("status_code", int(statusCode)).
		Str("status_text", statusCode.String()).
		Dur("duration", duration).
		Msg("received a gRPC response")
	return err
}

func SetTokenToContext(token string) *service.AuthToken {
	requestToken := new(service.AuthToken)
	requestToken.Token = token
	return requestToken
}
