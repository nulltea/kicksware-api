package gRPC

import (
	"context"
	"log"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"

	"search-service/api/gRPC/proto"
	"search-service/core/meta"
	"search-service/core/model"
	"search-service/core/pipe"
	"search-service/core/service"
	"search-service/env"
)

type referencePipe struct {
	client               proto.ReferencePipeServiceClient
	auth                 service.AuthService
}

func NewSneakerReferencePipe(auth service.AuthService, config env.CommonConfig) pipe.SneakerReferencePipe {
	return &referencePipe{
		newRemoteClient(config.InnerServiceFormat),
		auth,
	}
}

func newRemoteClient(serviceEndpoint string) proto.ReferencePipeServiceClient {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(serviceEndpoint, opts...); if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	return proto.NewReferencePipeServiceClient(conn)
}

func (p *referencePipe) authenticate() (string, error) {
	token, err := p.auth.Authenticate(); if err != nil {
		log.Fatalln(errors.Wrap(err, "search-service::startup.InnerServiceAuth: authenticate failed"))
		return "", err
	}
	return token, nil
}


func (p *referencePipe) FetchOne(code string) (ref *model.SneakerReference, err error) {
	ctx := context.Background()
	filter := &proto.ReferenceFilter{
		ReferenceID: []string{code},
	}
	ctl, err := p.client.GetReferences(ctx, filter); if err != nil {
		return nil, err
	}
	resp, err := ctl.Recv(); if err != nil {
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
	ctl, err := p.client.GetReferences(ctx, filter); if err != nil {
		return nil, err
	}
	resp, err := ctl.Recv(); if err != nil {
		return nil, err
	}
	refs = proto.ReferencesToNative(resp.References)
	return
}

func (p *referencePipe) FetchAll(params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	ctx := context.Background()
	ctl, err := p.client.GetReferences(ctx, nil); if err != nil {
		return nil, err
	}
	resp, err := ctl.Recv(); if err != nil {
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
	ctl, err := p.client.GetReferences(ctx, filter); if err != nil {
		return nil, err
	}
	resp, err := ctl.Recv(); if err != nil {
		return nil, err
	}
	refs = proto.ReferencesToNative(resp.References)
	return
}