package gRPC

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/service/jwt"
)

const (
	AuthMetaKey = "authorization"
	UserContextKey = "user_id"
)

type AuthInterceptor struct {
	jwtManager *jwt.TokenManager
	accessRoles map[string][]string
}

func NewAuthInterceptor(jwt *jwt.TokenManager, accessRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager: jwt,
		accessRoles: accessRoles,
	}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		claims, err := i.authenticate(ctx); if err != nil {
			return nil, err
		}

		err = i.authorize(claims, info.FullMethod); if err != nil {
			return nil, err
		}
		ctx = metadata.AppendToOutgoingContext(ctx, UserContextKey, claims.UniqueID)
		return handler(ctx, req)
	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
func (i *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		claims, err := i.authenticate(stream.Context()); if err != nil {
			return err
		}

		err = i.authorize(claims, info.FullMethod); if err != nil {
			return err
		}
		stream.SetHeader(metadata.New(map[string]string{UserContextKey: claims.UniqueID}))
		return handler(srv, stream)
	}
}

func (i *AuthInterceptor) authenticate(ctx context.Context) (*core.AuthClaims, error) {
	meta, ok := metadata.FromIncomingContext(ctx); if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := meta[AuthMetaKey]; if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := i.jwtManager.Verify(accessToken); if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %w", err)
	}

	return claims, nil
}

func (i *AuthInterceptor) authorize(claims *core.AuthClaims, method string) error {
	accessibleRoles, ok := i.accessRoles[method]; if !ok {
		return nil
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}