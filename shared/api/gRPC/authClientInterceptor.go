package gRPC

import (
	"context"
	"time"

	"github.com/golang/glog"
	"go.kicksware.com/api/services/users/api/gRPC/proto"
	"go.kicksware.com/api/services/users/core/meta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"go.kicksware.com/api/shared/core"
	tlsMeta "go.kicksware.com/api/shared/core/meta"
)

type AuthClientInterceptor struct {
	endpoint string
	authClient  proto.AuthServiceClient
	accessToken *meta.AuthToken
	authService core.AuthService
}

func NewAuthClientInterceptor(authEndpoint string, tls *tlsMeta.TLSCertificate, service core.AuthService) *AuthClientInterceptor {
	return &AuthClientInterceptor{
		endpoint: authEndpoint,
		authClient: newAuthClient(authEndpoint, tls),
		authService: service,
	}
}

func newAuthClient(serviceEndpoint string, tlsCert *tlsMeta.TLSCertificate) proto.AuthServiceClient {
	var opts []grpc.DialOption
	if tlsCert != nil && tlsCert.EnableTLS{
		tls, err := LoadClientTLSCredentials(tlsCert); if err != nil {
			glog.Fatalln("cannot load TLS credentials: ", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(tls))
	} else {
		opts = append(opts, grpc.WithInsecure())
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
		if !i.accessToken.Expires.IsZero() && i.accessToken.Expires.Nanosecond() <= time.Now().Nanosecond() {
			newToken, err := i.authClient.Refresh(ctx, proto.AuthToken{}.FromNative(i.accessToken)); if err == nil {
				i.accessToken = newToken.ToNative()
				return i.accessToken, nil
			}
		} else {
			return i.accessToken, nil
		}
	}

	if token, ok := tryRetrieveToken(ctx); ok {
		return token, nil
	}

	resp, err := i.authClient.Guest(ctx, i.accessKey()); if err == nil {
		return resp.ToNative(), nil
	}

	glog.Errorln(err)
	return nil, err
}

func (i *AuthClientInterceptor) attachToken(ctx context.Context) context.Context {
	token, err := i.requestAccessToken(ctx); if err != nil {
		return ctx
	}
	i.accessToken = token
	return metadata.AppendToOutgoingContext(ctx, AuthMetaKey, token.Token)
}

func tryRetrieveToken(ctx context.Context) (*meta.AuthToken, bool) {
	md, ok := metadata.FromIncomingContext(ctx); if !ok {
		return nil, false
	}

	values, ok := md[AuthMetaKey]; if !ok || len(values) == 0 {
		return nil, false
	}

	return &meta.AuthToken{
		Token:   values[0],
		Success: true,
		Expires: time.Time{},
	}, true
}

func (i *AuthClientInterceptor) accessKey() *proto.AccessKey {
	return &proto.AccessKey{
		Key: i.authService.AccessKey(),
	}
}
