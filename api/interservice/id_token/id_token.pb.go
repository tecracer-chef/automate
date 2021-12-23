// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: interservice/id_token/id_token.proto

package id_token

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

type ValidateIdTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id_token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" toml:"token,omitempty" mapstructure:"token,omitempty"`
}

func (x *ValidateIdTokenRequest) Reset() {
	*x = ValidateIdTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_id_token_id_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateIdTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateIdTokenRequest) ProtoMessage() {}

func (x *ValidateIdTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_id_token_id_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateIdTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateIdTokenRequest) Descriptor() ([]byte, []int) {
	return file_interservice_id_token_id_token_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateIdTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateIdTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsInvalid bool `protobuf:"varint,1,opt,name=is_invalid,json=isInvalid,proto3" json:"is_invalid,omitempty" toml:"is_invalid,omitempty" mapstructure:"is_invalid,omitempty"`
}

func (x *ValidateIdTokenResponse) Reset() {
	*x = ValidateIdTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_id_token_id_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateIdTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateIdTokenResponse) ProtoMessage() {}

func (x *ValidateIdTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_id_token_id_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateIdTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateIdTokenResponse) Descriptor() ([]byte, []int) {
	return file_interservice_id_token_id_token_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateIdTokenResponse) GetIsInvalid() bool {
	if x != nil {
		return x.IsInvalid
	}
	return false
}

var File_interservice_id_token_id_token_proto protoreflect.FileDescriptor

var file_interservice_id_token_id_token_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x17, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32,
	0x9d, 0x01, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_id_token_id_token_proto_rawDescOnce sync.Once
	file_interservice_id_token_id_token_proto_rawDescData = file_interservice_id_token_id_token_proto_rawDesc
)

func file_interservice_id_token_id_token_proto_rawDescGZIP() []byte {
	file_interservice_id_token_id_token_proto_rawDescOnce.Do(func() {
		file_interservice_id_token_id_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_id_token_id_token_proto_rawDescData)
	})
	return file_interservice_id_token_id_token_proto_rawDescData
}

var file_interservice_id_token_id_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interservice_id_token_id_token_proto_goTypes = []interface{}{
	(*ValidateIdTokenRequest)(nil),  // 0: chef.automate.domain.id_token.ValidateIdTokenRequest
	(*ValidateIdTokenResponse)(nil), // 1: chef.automate.domain.id_token.ValidateIdTokenResponse
}
var file_interservice_id_token_id_token_proto_depIdxs = []int32{
	0, // 0: chef.automate.domain.id_token.ValidateIdTokenService.ValidateIdToken:input_type -> chef.automate.domain.id_token.ValidateIdTokenRequest
	1, // 1: chef.automate.domain.id_token.ValidateIdTokenService.ValidateIdToken:output_type -> chef.automate.domain.id_token.ValidateIdTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_id_token_id_token_proto_init() }
func file_interservice_id_token_id_token_proto_init() {
	if File_interservice_id_token_id_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_id_token_id_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateIdTokenRequest); i {
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
		file_interservice_id_token_id_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateIdTokenResponse); i {
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
			RawDescriptor: file_interservice_id_token_id_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_id_token_id_token_proto_goTypes,
		DependencyIndexes: file_interservice_id_token_id_token_proto_depIdxs,
		MessageInfos:      file_interservice_id_token_id_token_proto_msgTypes,
	}.Build()
	File_interservice_id_token_id_token_proto = out.File
	file_interservice_id_token_id_token_proto_rawDesc = nil
	file_interservice_id_token_id_token_proto_goTypes = nil
	file_interservice_id_token_id_token_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ValidateIdTokenServiceClient is the client API for ValidateIdTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidateIdTokenServiceClient interface {
	ValidateIdToken(ctx context.Context, in *ValidateIdTokenRequest, opts ...grpc.CallOption) (*ValidateIdTokenResponse, error)
}

type validateIdTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidateIdTokenServiceClient(cc grpc.ClientConnInterface) ValidateIdTokenServiceClient {
	return &validateIdTokenServiceClient{cc}
}

func (c *validateIdTokenServiceClient) ValidateIdToken(ctx context.Context, in *ValidateIdTokenRequest, opts ...grpc.CallOption) (*ValidateIdTokenResponse, error) {
	out := new(ValidateIdTokenResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.id_token.ValidateIdTokenService/ValidateIdToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidateIdTokenServiceServer is the server API for ValidateIdTokenService service.
type ValidateIdTokenServiceServer interface {
	ValidateIdToken(context.Context, *ValidateIdTokenRequest) (*ValidateIdTokenResponse, error)
}

// UnimplementedValidateIdTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedValidateIdTokenServiceServer struct {
}

func (*UnimplementedValidateIdTokenServiceServer) ValidateIdToken(context.Context, *ValidateIdTokenRequest) (*ValidateIdTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateIdToken not implemented")
}

func RegisterValidateIdTokenServiceServer(s *grpc.Server, srv ValidateIdTokenServiceServer) {
	s.RegisterService(&_ValidateIdTokenService_serviceDesc, srv)
}

func _ValidateIdTokenService_ValidateIdToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateIdTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidateIdTokenServiceServer).ValidateIdToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.id_token.ValidateIdTokenService/ValidateIdToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidateIdTokenServiceServer).ValidateIdToken(ctx, req.(*ValidateIdTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValidateIdTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.id_token.ValidateIdTokenService",
	HandlerType: (*ValidateIdTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateIdToken",
			Handler:    _ValidateIdTokenService_ValidateIdToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/id_token/id_token.proto",
}
