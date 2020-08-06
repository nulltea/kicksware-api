// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: search.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SearchTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag           string         `protobuf:"bytes,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	RequestParams *RequestParams `protobuf:"bytes,2,opt,name=RequestParams,proto3" json:"RequestParams,omitempty"`
}

func (x *SearchTag) Reset() {
	*x = SearchTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTag) ProtoMessage() {}

func (x *SearchTag) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTag.ProtoReflect.Descriptor instead.
func (*SearchTag) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchTag) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *SearchTag) GetRequestParams() *RequestParams {
	if x != nil {
		return x.RequestParams
	}
	return nil
}

type SearchFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field         string         `protobuf:"bytes,1,opt,name=Field,proto3" json:"Field,omitempty"`
	Value         string         `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	SKU           string         `protobuf:"bytes,3,opt,name=SKU,proto3" json:"SKU,omitempty"`
	Model         string         `protobuf:"bytes,4,opt,name=Model,proto3" json:"Model,omitempty"`
	Brand         string         `protobuf:"bytes,5,opt,name=Brand,proto3" json:"Brand,omitempty"`
	RequestParams *RequestParams `protobuf:"bytes,6,opt,name=RequestParams,proto3" json:"RequestParams,omitempty"`
}

func (x *SearchFilter) Reset() {
	*x = SearchFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFilter) ProtoMessage() {}

func (x *SearchFilter) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFilter.ProtoReflect.Descriptor instead.
func (*SearchFilter) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchFilter) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SearchFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SearchFilter) GetSKU() string {
	if x != nil {
		return x.SKU
	}
	return ""
}

func (x *SearchFilter) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *SearchFilter) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *SearchFilter) GetRequestParams() *RequestParams {
	if x != nil {
		return x.RequestParams
	}
	return nil
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x54,
	0x61, 0x67, 0x12, 0x3a, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xb4,
	0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x4b, 0x55, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x4b, 0x55, 0x12, 0x14, 0x0a,
	0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x3a, 0x0a, 0x0d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0xca, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x67, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x79, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0xbf, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x67, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x12, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x04,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_search_proto_goTypes = []interface{}{
	(*SearchTag)(nil),         // 0: proto.SearchTag
	(*SearchFilter)(nil),      // 1: proto.SearchFilter
	(*RequestParams)(nil),     // 2: proto.RequestParams
	(*ReferenceFilter)(nil),   // 3: proto.ReferenceFilter
	(*ProductFilter)(nil),     // 4: proto.ProductFilter
	(*ReferenceResponse)(nil), // 5: proto.ReferenceResponse
	(*ProductResponse)(nil),   // 6: proto.ProductResponse
}
var file_search_proto_depIdxs = []int32{
	2, // 0: proto.SearchTag.RequestParams:type_name -> proto.RequestParams
	2, // 1: proto.SearchFilter.RequestParams:type_name -> proto.RequestParams
	0, // 2: proto.SearchReferencesService.Search:input_type -> proto.SearchTag
	1, // 3: proto.SearchReferencesService.SearchBy:input_type -> proto.SearchFilter
	3, // 4: proto.SearchReferencesService.Sync:input_type -> proto.ReferenceFilter
	0, // 5: proto.SearchProductService.Search:input_type -> proto.SearchTag
	1, // 6: proto.SearchProductService.SearchBy:input_type -> proto.SearchFilter
	4, // 7: proto.SearchProductService.Sync:input_type -> proto.ProductFilter
	5, // 8: proto.SearchReferencesService.Search:output_type -> proto.ReferenceResponse
	5, // 9: proto.SearchReferencesService.SearchBy:output_type -> proto.ReferenceResponse
	5, // 10: proto.SearchReferencesService.Sync:output_type -> proto.ReferenceResponse
	6, // 11: proto.SearchProductService.Search:output_type -> proto.ProductResponse
	6, // 12: proto.SearchProductService.SearchBy:output_type -> proto.ProductResponse
	6, // 13: proto.SearchProductService.Sync:output_type -> proto.ProductResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	file_common_proto_init()
	file_reference_proto_init()
	file_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SearchReferencesServiceClient is the client API for SearchReferencesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchReferencesServiceClient interface {
	Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*ReferenceResponse, error)
	SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*ReferenceResponse, error)
	Sync(ctx context.Context, in *ReferenceFilter, opts ...grpc.CallOption) (*ReferenceResponse, error)
}

type searchReferencesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchReferencesServiceClient(cc grpc.ClientConnInterface) SearchReferencesServiceClient {
	return &searchReferencesServiceClient{cc}
}

func (c *searchReferencesServiceClient) Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*ReferenceResponse, error) {
	out := new(ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchReferencesServiceClient) SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*ReferenceResponse, error) {
	out := new(ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/SearchBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchReferencesServiceClient) Sync(ctx context.Context, in *ReferenceFilter, opts ...grpc.CallOption) (*ReferenceResponse, error) {
	out := new(ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchReferencesServiceServer is the server API for SearchReferencesService service.
type SearchReferencesServiceServer interface {
	Search(context.Context, *SearchTag) (*ReferenceResponse, error)
	SearchBy(context.Context, *SearchFilter) (*ReferenceResponse, error)
	Sync(context.Context, *ReferenceFilter) (*ReferenceResponse, error)
}

// UnimplementedSearchReferencesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchReferencesServiceServer struct {
}

func (*UnimplementedSearchReferencesServiceServer) Search(context.Context, *SearchTag) (*ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedSearchReferencesServiceServer) SearchBy(context.Context, *SearchFilter) (*ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBy not implemented")
}
func (*UnimplementedSearchReferencesServiceServer) Sync(context.Context, *ReferenceFilter) (*ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}

func RegisterSearchReferencesServiceServer(s *grpc.Server, srv SearchReferencesServiceServer) {
	s.RegisterService(&_SearchReferencesService_serviceDesc, srv)
}

func _SearchReferencesService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchReferencesServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchReferencesService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchReferencesServiceServer).Search(ctx, req.(*SearchTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchReferencesService_SearchBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchReferencesServiceServer).SearchBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchReferencesService/SearchBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchReferencesServiceServer).SearchBy(ctx, req.(*SearchFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchReferencesService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferenceFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchReferencesServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchReferencesService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchReferencesServiceServer).Sync(ctx, req.(*ReferenceFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchReferencesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchReferencesService",
	HandlerType: (*SearchReferencesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchReferencesService_Search_Handler,
		},
		{
			MethodName: "SearchBy",
			Handler:    _SearchReferencesService_SearchBy_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _SearchReferencesService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}

// SearchProductServiceClient is the client API for SearchProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchProductServiceClient interface {
	Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*ProductResponse, error)
	SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*ProductResponse, error)
	Sync(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error)
}

type searchProductServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchProductServiceClient(cc grpc.ClientConnInterface) SearchProductServiceClient {
	return &searchProductServiceClient{cc}
}

func (c *searchProductServiceClient) Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchProductService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchProductServiceClient) SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchProductService/SearchBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchProductServiceClient) Sync(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchProductService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchProductServiceServer is the server API for SearchProductService service.
type SearchProductServiceServer interface {
	Search(context.Context, *SearchTag) (*ProductResponse, error)
	SearchBy(context.Context, *SearchFilter) (*ProductResponse, error)
	Sync(context.Context, *ProductFilter) (*ProductResponse, error)
}

// UnimplementedSearchProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchProductServiceServer struct {
}

func (*UnimplementedSearchProductServiceServer) Search(context.Context, *SearchTag) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedSearchProductServiceServer) SearchBy(context.Context, *SearchFilter) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBy not implemented")
}
func (*UnimplementedSearchProductServiceServer) Sync(context.Context, *ProductFilter) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}

func RegisterSearchProductServiceServer(s *grpc.Server, srv SearchProductServiceServer) {
	s.RegisterService(&_SearchProductService_serviceDesc, srv)
}

func _SearchProductService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchProductServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchProductService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchProductServiceServer).Search(ctx, req.(*SearchTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchProductService_SearchBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchProductServiceServer).SearchBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchProductService/SearchBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchProductServiceServer).SearchBy(ctx, req.(*SearchFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchProductService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchProductServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchProductService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchProductServiceServer).Sync(ctx, req.(*ProductFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchProductService",
	HandlerType: (*SearchProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchProductService_Search_Handler,
		},
		{
			MethodName: "SearchBy",
			Handler:    _SearchProductService_SearchBy_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _SearchProductService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}