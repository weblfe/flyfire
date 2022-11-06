package grpc

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	grpc2 "google.golang.org/grpc"
)

func NewClientConnection(r registry.Discovery, logger log.Logger, endpoint string) grpc2.ClientConnInterface {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// TODO: should we set less
		grpc.WithTimeout(30*time.Second),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			AppendHeader(),
			mmd.Client(),
			recovery.Recovery(),
			logging.Client(logger),
		),
		//	grpc.WithUnaryInterceptor(),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func AppendHeader() middleware.Middleware {
	var NeedAppendHeader = []string{"Accept-Language"}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				params := make([]string, 0)
				for _, key := range NeedAppendHeader {
					value := tr.RequestHeader().Get(key)
					if len(value) > 0 {
						params = append(params, key, value)
					}
				}
				if len(params) > 0 {
					ctx = metadata.AppendToClientContext(ctx, params...)
				}
			}
			return handler(ctx, req)
		}
	}
}
