package gRPC

import (
	"context"

	"github.com/golang/glog"
	"go.kicksware.com/api/shared/core"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"

	"go.kicksware.com/api/services/references/core/model"
	gRPCSrv "go.kicksware.com/api/shared/api/gRPC"

	"go.kicksware.com/api/services/references/api/gRPC/proto"
	common "go.kicksware.com/api/shared/api/proto"
	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/search/core/pipe"
	"go.kicksware.com/api/services/search/env"
)

type referencePipe struct {
	client proto.ReferenceServiceClient
	auth   *gRPCSrv.AuthClientInterceptor
}

func NewSneakerReferencePipe(config env.ServiceConfig, service core.AuthService) pipe.SneakerReferencePipe {
	auth := gRPCSrv.NewAuthClientInterceptor(config.Common.RpcEndpointFormat, config.Auth.TLSCertificate, service)
	return &referencePipe{
		client: newRemoteClient(config, auth),
		auth:   auth,
	}
}

func newRemoteClient(config env.ServiceConfig, auth *gRPCSrv.AuthClientInterceptor) proto.ReferenceServiceClient {
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024 * 1024 * 50),
		),
		grpc.WithUnaryInterceptor(auth.Unary()),
	}
	if config.Auth.TLSCertificate.EnableTLS {
		tls, err := gRPCSrv.LoadClientTLSCredentials(config.Auth.TLSCertificate); if err != nil {
			glog.Fatalln("cannot load TLS credentials: ", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(tls))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(config.Common.RpcEndpointFormat, opts...); if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}

	return proto.NewReferenceServiceClient(conn)
}


func (p *referencePipe) FetchOne(code string) (ref *model.SneakerReference, err error) {
	ctx := context.Background()
	filter := &proto.ReferenceFilter{
		ReferenceID: []string{code},
	}
	resp, err := p.client.GetReferences(ctx, filter); if err != nil {
		return nil, err
	}
	ref = resp.References[0].ToNative()
	return
}

func (p *referencePipe) Fetch(codes []string, params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	ctx := context.Background()
	filter := &proto.ReferenceFilter{
		ReferenceID: codes,
	}
	resp, err := p.client.GetReferences(ctx, filter); if err != nil {
		return nil, err
	}
	refs = proto.ReferencesToNative(resp.References)
	return
}

func (p *referencePipe) FetchAll(params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	ctx := context.Background()
	resp, err := p.client.GetReferences(ctx, &proto.ReferenceFilter{}); if err != nil {
		return nil, err
	}
	refs = proto.ReferencesToNative(resp.References)
	return
}

func (p *referencePipe) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) (refs[]*model.SneakerReference, err error) {
	ctx := context.Background()
	str, err := structpb.NewStruct(query); if err != nil {
		return nil, err
	}
	filter := &proto.ReferenceFilter{
		RequestQuery: str,
		RequestParams: common.RequestParams{}.FromNative(params),
	}
	resp, err := p.client.GetReferences(ctx, filter); if err != nil {
		return nil, err
	}
	refs = proto.ReferencesToNative(resp.References)
	return
}
