// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: interservice/compliance/ingest/ingest/compliance.proto

package ingest

import (
	context "context"
	compliance "github.com/chef/automate/api/interservice/compliance/ingest/events/compliance"
	event "github.com/chef/automate/api/interservice/event"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ProjectUpdateStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectUpdateStatusReq) Reset() {
	*x = ProjectUpdateStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectUpdateStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectUpdateStatusReq) ProtoMessage() {}

func (x *ProjectUpdateStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectUpdateStatusReq.ProtoReflect.Descriptor instead.
func (*ProjectUpdateStatusReq) Descriptor() ([]byte, []int) {
	return file_interservice_compliance_ingest_ingest_compliance_proto_rawDescGZIP(), []int{0}
}

type ProjectUpdateStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State                 string                 `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty" toml:"state,omitempty" mapstructure:"state,omitempty"`
	EstimatedTimeComplete *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=estimated_time_complete,json=estimatedTimeComplete,proto3" json:"estimated_time_complete,omitempty" toml:"estimated_time_complete,omitempty" mapstructure:"estimated_time_complete,omitempty"`
	PercentageComplete    float32                `protobuf:"fixed32,3,opt,name=percentage_complete,json=percentageComplete,proto3" json:"percentage_complete,omitempty" toml:"percentage_complete,omitempty" mapstructure:"percentage_complete,omitempty"`
}

func (x *ProjectUpdateStatusResp) Reset() {
	*x = ProjectUpdateStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectUpdateStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectUpdateStatusResp) ProtoMessage() {}

func (x *ProjectUpdateStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectUpdateStatusResp.ProtoReflect.Descriptor instead.
func (*ProjectUpdateStatusResp) Descriptor() ([]byte, []int) {
	return file_interservice_compliance_ingest_ingest_compliance_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectUpdateStatusResp) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ProjectUpdateStatusResp) GetEstimatedTimeComplete() *timestamppb.Timestamp {
	if x != nil {
		return x.EstimatedTimeComplete
	}
	return nil
}

func (x *ProjectUpdateStatusResp) GetPercentageComplete() float32 {
	if x != nil {
		return x.PercentageComplete
	}
	return 0
}

var File_interservice_compliance_ingest_ingest_compliance_proto protoreflect.FileDescriptor

var file_interservice_compliance_ingest_ingest_compliance_proto_rawDesc = []byte{
	0x0a, 0x36, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x22, 0xb4, 0x01, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x15, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x12, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x32, 0xcd, 0x03, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x40, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x73,
	0x67, 0x1a, 0x2d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0xa4, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x45, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x46, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_compliance_ingest_ingest_compliance_proto_rawDescOnce sync.Once
	file_interservice_compliance_ingest_ingest_compliance_proto_rawDescData = file_interservice_compliance_ingest_ingest_compliance_proto_rawDesc
)

func file_interservice_compliance_ingest_ingest_compliance_proto_rawDescGZIP() []byte {
	file_interservice_compliance_ingest_ingest_compliance_proto_rawDescOnce.Do(func() {
		file_interservice_compliance_ingest_ingest_compliance_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_compliance_ingest_ingest_compliance_proto_rawDescData)
	})
	return file_interservice_compliance_ingest_ingest_compliance_proto_rawDescData
}

var file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interservice_compliance_ingest_ingest_compliance_proto_goTypes = []interface{}{
	(*ProjectUpdateStatusReq)(nil),  // 0: chef.automate.domain.compliance.ingest.ingest.ProjectUpdateStatusReq
	(*ProjectUpdateStatusResp)(nil), // 1: chef.automate.domain.compliance.ingest.ingest.ProjectUpdateStatusResp
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
	(*compliance.Report)(nil),       // 3: chef.automate.domain.compliance.ingest.events.compliance.Report
	(*event.EventMsg)(nil),          // 4: chef.automate.domain.event.api.EventMsg
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
	(*event.EventResponse)(nil),     // 6: chef.automate.domain.event.api.EventResponse
}
var file_interservice_compliance_ingest_ingest_compliance_proto_depIdxs = []int32{
	2, // 0: chef.automate.domain.compliance.ingest.ingest.ProjectUpdateStatusResp.estimated_time_complete:type_name -> google.protobuf.Timestamp
	3, // 1: chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService.ProcessComplianceReport:input_type -> chef.automate.domain.compliance.ingest.events.compliance.Report
	4, // 2: chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService.HandleEvent:input_type -> chef.automate.domain.event.api.EventMsg
	0, // 3: chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService.ProjectUpdateStatus:input_type -> chef.automate.domain.compliance.ingest.ingest.ProjectUpdateStatusReq
	5, // 4: chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService.ProcessComplianceReport:output_type -> google.protobuf.Empty
	6, // 5: chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService.HandleEvent:output_type -> chef.automate.domain.event.api.EventResponse
	1, // 6: chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService.ProjectUpdateStatus:output_type -> chef.automate.domain.compliance.ingest.ingest.ProjectUpdateStatusResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interservice_compliance_ingest_ingest_compliance_proto_init() }
func file_interservice_compliance_ingest_ingest_compliance_proto_init() {
	if File_interservice_compliance_ingest_ingest_compliance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectUpdateStatusReq); i {
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
		file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectUpdateStatusResp); i {
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
			RawDescriptor: file_interservice_compliance_ingest_ingest_compliance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_compliance_ingest_ingest_compliance_proto_goTypes,
		DependencyIndexes: file_interservice_compliance_ingest_ingest_compliance_proto_depIdxs,
		MessageInfos:      file_interservice_compliance_ingest_ingest_compliance_proto_msgTypes,
	}.Build()
	File_interservice_compliance_ingest_ingest_compliance_proto = out.File
	file_interservice_compliance_ingest_ingest_compliance_proto_rawDesc = nil
	file_interservice_compliance_ingest_ingest_compliance_proto_goTypes = nil
	file_interservice_compliance_ingest_ingest_compliance_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComplianceIngesterServiceClient is the client API for ComplianceIngesterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComplianceIngesterServiceClient interface {
	ProcessComplianceReport(ctx context.Context, in *compliance.Report, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HandleEvent(ctx context.Context, in *event.EventMsg, opts ...grpc.CallOption) (*event.EventResponse, error)
	ProjectUpdateStatus(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateStatusResp, error)
}

type complianceIngesterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceIngesterServiceClient(cc grpc.ClientConnInterface) ComplianceIngesterServiceClient {
	return &complianceIngesterServiceClient{cc}
}

func (c *complianceIngesterServiceClient) ProcessComplianceReport(ctx context.Context, in *compliance.Report, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService/ProcessComplianceReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceIngesterServiceClient) HandleEvent(ctx context.Context, in *event.EventMsg, opts ...grpc.CallOption) (*event.EventResponse, error) {
	out := new(event.EventResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService/HandleEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceIngesterServiceClient) ProjectUpdateStatus(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateStatusResp, error) {
	out := new(ProjectUpdateStatusResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService/ProjectUpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceIngesterServiceServer is the server API for ComplianceIngesterService service.
type ComplianceIngesterServiceServer interface {
	ProcessComplianceReport(context.Context, *compliance.Report) (*emptypb.Empty, error)
	HandleEvent(context.Context, *event.EventMsg) (*event.EventResponse, error)
	ProjectUpdateStatus(context.Context, *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error)
}

// UnimplementedComplianceIngesterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedComplianceIngesterServiceServer struct {
}

func (*UnimplementedComplianceIngesterServiceServer) ProcessComplianceReport(context.Context, *compliance.Report) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessComplianceReport not implemented")
}
func (*UnimplementedComplianceIngesterServiceServer) HandleEvent(context.Context, *event.EventMsg) (*event.EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleEvent not implemented")
}
func (*UnimplementedComplianceIngesterServiceServer) ProjectUpdateStatus(context.Context, *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectUpdateStatus not implemented")
}

func RegisterComplianceIngesterServiceServer(s *grpc.Server, srv ComplianceIngesterServiceServer) {
	s.RegisterService(&_ComplianceIngesterService_serviceDesc, srv)
}

func _ComplianceIngesterService_ProcessComplianceReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(compliance.Report)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIngesterServiceServer).ProcessComplianceReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService/ProcessComplianceReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIngesterServiceServer).ProcessComplianceReport(ctx, req.(*compliance.Report))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceIngesterService_HandleEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(event.EventMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIngesterServiceServer).HandleEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService/HandleEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIngesterServiceServer).HandleEvent(ctx, req.(*event.EventMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceIngesterService_ProjectUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectUpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIngesterServiceServer).ProjectUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService/ProjectUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIngesterServiceServer).ProjectUpdateStatus(ctx, req.(*ProjectUpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComplianceIngesterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.compliance.ingest.ingest.ComplianceIngesterService",
	HandlerType: (*ComplianceIngesterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessComplianceReport",
			Handler:    _ComplianceIngesterService_ProcessComplianceReport_Handler,
		},
		{
			MethodName: "HandleEvent",
			Handler:    _ComplianceIngesterService_HandleEvent_Handler,
		},
		{
			MethodName: "ProjectUpdateStatus",
			Handler:    _ComplianceIngesterService_ProjectUpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/compliance/ingest/ingest/compliance.proto",
}
