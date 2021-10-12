// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: search.proto

package proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	proto3 "go.kicksware.com/api/services/products/api/gRPC/proto"
	proto2 "go.kicksware.com/api/services/references/api/gRPC/proto"
	proto1 "go.kicksware.com/api/shared/api/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

	Tag           string                `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	RequestParams *proto1.RequestParams `protobuf:"bytes,2,opt,name=requestParams,proto3" json:"requestParams,omitempty"`
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

func (x *SearchTag) GetRequestParams() *proto1.RequestParams {
	if x != nil {
		return x.RequestParams
	}
	return nil
}

type SearchFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field         string                `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value         string                `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	RequestParams *proto1.RequestParams `protobuf:"bytes,6,opt,name=requestParams,proto3" json:"requestParams,omitempty"`
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

func (x *SearchFilter) GetRequestParams() *proto1.RequestParams {
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
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x76,
	0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0x88, 0x03, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x67, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x53, 0x4b, 0x55, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0xbf, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x61, 0x67, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x04, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x49, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x69, 0x6d, 0x6f, 0x74, 0x68, 0x2d, 0x79, 0x2f, 0x6b, 0x69, 0x63, 0x6b, 0x73,
	0x77, 0x61, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x52, 0x50, 0x43,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*SearchTag)(nil),                // 0: proto.SearchTag
	(*SearchFilter)(nil),             // 1: proto.SearchFilter
	(*proto1.RequestParams)(nil),     // 2: proto.RequestParams
	(*proto2.ReferenceFilter)(nil),   // 3: proto.ReferenceFilter
	(*proto3.ProductFilter)(nil),     // 4: proto.ProductFilter
	(*proto2.ReferenceResponse)(nil), // 5: proto.ReferenceResponse
	(*proto3.ProductResponse)(nil),   // 6: proto.ProductResponse
}
var file_search_proto_depIdxs = []int32{
	2,  // 0: proto.SearchTag.requestParams:type_name -> proto.RequestParams
	2,  // 1: proto.SearchFilter.requestParams:type_name -> proto.RequestParams
	0,  // 2: proto.SearchReferencesService.Search:input_type -> proto.SearchTag
	1,  // 3: proto.SearchReferencesService.SearchBy:input_type -> proto.SearchFilter
	1,  // 4: proto.SearchReferencesService.SearchSKU:input_type -> proto.SearchFilter
	1,  // 5: proto.SearchReferencesService.SearchBrand:input_type -> proto.SearchFilter
	1,  // 6: proto.SearchReferencesService.SearchModel:input_type -> proto.SearchFilter
	3,  // 7: proto.SearchReferencesService.Sync:input_type -> proto.ReferenceFilter
	0,  // 8: proto.SearchProductService.Search:input_type -> proto.SearchTag
	1,  // 9: proto.SearchProductService.SearchBy:input_type -> proto.SearchFilter
	4,  // 10: proto.SearchProductService.Sync:input_type -> proto.ProductFilter
	5,  // 11: proto.SearchReferencesService.Search:output_type -> proto.ReferenceResponse
	5,  // 12: proto.SearchReferencesService.SearchBy:output_type -> proto.ReferenceResponse
	5,  // 13: proto.SearchReferencesService.SearchSKU:output_type -> proto.ReferenceResponse
	5,  // 14: proto.SearchReferencesService.SearchBrand:output_type -> proto.ReferenceResponse
	5,  // 15: proto.SearchReferencesService.SearchModel:output_type -> proto.ReferenceResponse
	5,  // 16: proto.SearchReferencesService.Sync:output_type -> proto.ReferenceResponse
	6,  // 17: proto.SearchProductService.Search:output_type -> proto.ProductResponse
	6,  // 18: proto.SearchProductService.SearchBy:output_type -> proto.ProductResponse
	6,  // 19: proto.SearchProductService.Sync:output_type -> proto.ProductResponse
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
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
	Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error)
	SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error)
	SearchSKU(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error)
	SearchBrand(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error)
	SearchModel(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error)
	Sync(ctx context.Context, in *proto2.ReferenceFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error)
}

type searchReferencesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchReferencesServiceClient(cc grpc.ClientConnInterface) SearchReferencesServiceClient {
	return &searchReferencesServiceClient{cc}
}

func (c *searchReferencesServiceClient) Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error) {
	out := new(proto2.ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchReferencesServiceClient) SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error) {
	out := new(proto2.ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/SearchBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchReferencesServiceClient) SearchSKU(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error) {
	out := new(proto2.ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/SearchSKU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchReferencesServiceClient) SearchBrand(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error) {
	out := new(proto2.ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/SearchBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchReferencesServiceClient) SearchModel(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error) {
	out := new(proto2.ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/SearchModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchReferencesServiceClient) Sync(ctx context.Context, in *proto2.ReferenceFilter, opts ...grpc.CallOption) (*proto2.ReferenceResponse, error) {
	out := new(proto2.ReferenceResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchReferencesService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchReferencesServiceServer is the server API for SearchReferencesService service.
type SearchReferencesServiceServer interface {
	Search(context.Context, *SearchTag) (*proto2.ReferenceResponse, error)
	SearchBy(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error)
	SearchSKU(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error)
	SearchBrand(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error)
	SearchModel(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error)
	Sync(context.Context, *proto2.ReferenceFilter) (*proto2.ReferenceResponse, error)
}

// UnimplementedSearchReferencesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchReferencesServiceServer struct {
}

func (*UnimplementedSearchReferencesServiceServer) Search(context.Context, *SearchTag) (*proto2.ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedSearchReferencesServiceServer) SearchBy(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBy not implemented")
}
func (*UnimplementedSearchReferencesServiceServer) SearchSKU(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSKU not implemented")
}
func (*UnimplementedSearchReferencesServiceServer) SearchBrand(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBrand not implemented")
}
func (*UnimplementedSearchReferencesServiceServer) SearchModel(context.Context, *SearchFilter) (*proto2.ReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchModel not implemented")
}
func (*UnimplementedSearchReferencesServiceServer) Sync(context.Context, *proto2.ReferenceFilter) (*proto2.ReferenceResponse, error) {
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

func _SearchReferencesService_SearchSKU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchReferencesServiceServer).SearchSKU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchReferencesService/SearchSKU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchReferencesServiceServer).SearchSKU(ctx, req.(*SearchFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchReferencesService_SearchBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchReferencesServiceServer).SearchBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchReferencesService/SearchBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchReferencesServiceServer).SearchBrand(ctx, req.(*SearchFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchReferencesService_SearchModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchReferencesServiceServer).SearchModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchReferencesService/SearchModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchReferencesServiceServer).SearchModel(ctx, req.(*SearchFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchReferencesService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.ReferenceFilter)
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
		return srv.(SearchReferencesServiceServer).Sync(ctx, req.(*proto2.ReferenceFilter))
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
			MethodName: "SearchSKU",
			Handler:    _SearchReferencesService_SearchSKU_Handler,
		},
		{
			MethodName: "SearchBrand",
			Handler:    _SearchReferencesService_SearchBrand_Handler,
		},
		{
			MethodName: "SearchModel",
			Handler:    _SearchReferencesService_SearchModel_Handler,
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
	Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*proto3.ProductResponse, error)
	SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto3.ProductResponse, error)
	Sync(ctx context.Context, in *proto3.ProductFilter, opts ...grpc.CallOption) (*proto3.ProductResponse, error)
}

type searchProductServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchProductServiceClient(cc grpc.ClientConnInterface) SearchProductServiceClient {
	return &searchProductServiceClient{cc}
}

func (c *searchProductServiceClient) Search(ctx context.Context, in *SearchTag, opts ...grpc.CallOption) (*proto3.ProductResponse, error) {
	out := new(proto3.ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchProductService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchProductServiceClient) SearchBy(ctx context.Context, in *SearchFilter, opts ...grpc.CallOption) (*proto3.ProductResponse, error) {
	out := new(proto3.ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchProductService/SearchBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchProductServiceClient) Sync(ctx context.Context, in *proto3.ProductFilter, opts ...grpc.CallOption) (*proto3.ProductResponse, error) {
	out := new(proto3.ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.SearchProductService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchProductServiceServer is the server API for SearchProductService service.
type SearchProductServiceServer interface {
	Search(context.Context, *SearchTag) (*proto3.ProductResponse, error)
	SearchBy(context.Context, *SearchFilter) (*proto3.ProductResponse, error)
	Sync(context.Context, *proto3.ProductFilter) (*proto3.ProductResponse, error)
}

// UnimplementedSearchProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchProductServiceServer struct {
}

func (*UnimplementedSearchProductServiceServer) Search(context.Context, *SearchTag) (*proto3.ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedSearchProductServiceServer) SearchBy(context.Context, *SearchFilter) (*proto3.ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBy not implemented")
}
func (*UnimplementedSearchProductServiceServer) Sync(context.Context, *proto3.ProductFilter) (*proto3.ProductResponse, error) {
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
	in := new(proto3.ProductFilter)
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
		return srv.(SearchProductServiceServer).Sync(ctx, req.(*proto3.ProductFilter))
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