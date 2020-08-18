package gRPC

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/core/model"
	gRPCSrv "github.com/timoth-y/kicksware-platform/middleware-service/service-common/service/gRPC"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/pipe"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"
)

type referencePipe struct {
	client proto.ReferenceServiceClient
	auth   *gRPCSrv.AuthClientInterceptor
}

func NewSneakerReferencePipe(config env.ServiceConfig) pipe.SneakerReferencePipe {
	auth := gRPCSrv.NewAuthClientInterceptor(config.Common.InnerServiceFormat)
	return &referencePipe{
		client: newRemoteClient(config, auth),
		auth: auth,
	}
}

func newRemoteClient(config env.ServiceConfig, auth *gRPCSrv.AuthClientInterceptor) proto.ReferenceServiceClient {
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024 * 1024 * 50),
		),
		grpc.WithUnaryInterceptor(auth.Unary()),
	}
	if config.Security.TLSCertificate.EnableTLS {
		tls, err := gRPCSrv.LoadClientTLSCredentials(config.Security.TLSCertificate); if err != nil {
			glog.Fatalln("cannot load TLS credentials: ", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(tls))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(config.Common.InnerServiceFormat, opts...); if err != nil {
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
		RequestParams: proto.RequestParams{}.FromNative(params),
	}
	resp, err := p.client.GetReferences(ctx, filter); if err != nil {
		return nil, err
	}
	refs = proto.ReferencesToNative(resp.References)
	return
}