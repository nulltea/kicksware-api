package gRPC

import (
	"context"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/meta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthClientInterceptor struct {
	endpoint string
	authClient  proto.AuthServiceClient
	accessToken *meta.AuthToken
}

func NewAuthClientInterceptor(authEndpoint string) *AuthClientInterceptor {
	return &AuthClientInterceptor{
		endpoint: authEndpoint,
		authClient: newAuthClient(authEndpoint),
	}
}

func newAuthClient(serviceEndpoint string) proto.AuthServiceClient {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(serviceEndpoint, opts...); if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}

	return proto.NewAuthServiceClient(conn)
}

// Unary returns a client interceptor to authenticate unary RPC
func (i *AuthClientInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(i.attachToken(ctx), method, req, reply, cc, opts...)
	}
}

// Stream returns a client interceptor to authenticate stream RPC
func (i *AuthClientInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		return streamer(i.attachToken(ctx), desc, cc, method, opts...)
	}
}

func (i *AuthClientInterceptor) requestAccessToken(ctx context.Context) (*meta.AuthToken, error) {
	if i.accessToken != nil {
		if i.accessToken.Expires.Nanosecond() <= time.Now().Nanosecond() {
			newToken, err := i.authClient.Refresh(ctx, proto.AuthToken{}.FromNative(i.accessToken)); if err == nil {
				i.accessToken = newToken.ToNative()
				return i.accessToken, nil
			}
		} else {
			return i.accessToken, nil
		}
	}

	resp, err := i.authClient.Guest(ctx, &empty.Empty{}); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	return resp.ToNative(), nil
}

func (i *AuthClientInterceptor) attachToken(ctx context.Context) context.Context {
	token, err := i.requestAccessToken(ctx); if err != nil {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, AuthMetaKey, token.Token)
}