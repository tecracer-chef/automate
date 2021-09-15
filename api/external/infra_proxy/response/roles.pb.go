// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: external/infra_proxy/response/roles.proto

package response

import (
	proto "github.com/golang/protobuf/proto"
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

type Roles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of the roles item.
	Roles []*RoleListItem `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	// Starting page for the results.
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// Total number of records.
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Roles) Reset() {
	*x = Roles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roles) ProtoMessage() {}

func (x *Roles) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roles.ProtoReflect.Descriptor instead.
func (*Roles) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_roles_proto_rawDescGZIP(), []int{0}
}

func (x *Roles) GetRoles() []*RoleListItem {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Roles) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Roles) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type RoleListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Desscription of the role.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Environment for the role.
	Environments []string `protobuf:"bytes,3,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *RoleListItem) Reset() {
	*x = RoleListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleListItem) ProtoMessage() {}

func (x *RoleListItem) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleListItem.ProtoReflect.Descriptor instead.
func (*RoleListItem) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_roles_proto_rawDescGZIP(), []int{1}
}

func (x *RoleListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleListItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoleListItem) GetEnvironments() []string {
	if x != nil {
		return x.Environments
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the chef object.
	ChefType string `protobuf:"bytes,2,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty"`
	// Descrption of the role.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Role default attributes JSON.
	DefaultAttributes string `protobuf:"bytes,4,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Role override attributes JSON.
	OverrideAttributes string `protobuf:"bytes,5,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	// Json class name.
	JsonClass string `protobuf:"bytes,6,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty"`
	// Run list for the role.
	RunList []string `protobuf:"bytes,7,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_roles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_roles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_roles_proto_rawDescGZIP(), []int{2}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetChefType() string {
	if x != nil {
		return x.ChefType
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetDefaultAttributes() string {
	if x != nil {
		return x.DefaultAttributes
	}
	return ""
}

func (x *Role) GetOverrideAttributes() string {
	if x != nil {
		return x.OverrideAttributes
	}
	return ""
}

func (x *Role) GetJsonClass() string {
	if x != nil {
		return x.JsonClass
	}
	return ""
}

func (x *Role) GetRunList() []string {
	if x != nil {
		return x.RunList
	}
	return nil
}

type ExpandedRunList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the run list collection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// List of the run list.
	RunList []*RunList `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
}

func (x *ExpandedRunList) Reset() {
	*x = ExpandedRunList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_roles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedRunList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedRunList) ProtoMessage() {}

func (x *ExpandedRunList) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_roles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedRunList.ProtoReflect.Descriptor instead.
func (*ExpandedRunList) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_roles_proto_rawDescGZIP(), []int{3}
}

func (x *ExpandedRunList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExpandedRunList) GetRunList() []*RunList {
	if x != nil {
		return x.RunList
	}
	return nil
}

type RoleEnvironments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Role environment list.
	Environments []string `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *RoleEnvironments) Reset() {
	*x = RoleEnvironments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_roles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleEnvironments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleEnvironments) ProtoMessage() {}

func (x *RoleEnvironments) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_roles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleEnvironments.ProtoReflect.Descriptor instead.
func (*RoleEnvironments) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_roles_proto_rawDescGZIP(), []int{4}
}

func (x *RoleEnvironments) GetEnvironments() []string {
	if x != nil {
		return x.Environments
	}
	return nil
}

var File_external_infra_proxy_response_roles_proto protoreflect.FileDescriptor

var file_external_infra_proxy_response_roles_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7d, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x68,
	0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xf3, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x66, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6d,
	0x0a, 0x0f, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x4a, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x36, 0x0a,
	0x10, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_infra_proxy_response_roles_proto_rawDescOnce sync.Once
	file_external_infra_proxy_response_roles_proto_rawDescData = file_external_infra_proxy_response_roles_proto_rawDesc
)

func file_external_infra_proxy_response_roles_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_response_roles_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_response_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_response_roles_proto_rawDescData)
	})
	return file_external_infra_proxy_response_roles_proto_rawDescData
}

var file_external_infra_proxy_response_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_external_infra_proxy_response_roles_proto_goTypes = []interface{}{
	(*Roles)(nil),            // 0: chef.automate.api.infra_proxy.response.Roles
	(*RoleListItem)(nil),     // 1: chef.automate.api.infra_proxy.response.RoleListItem
	(*Role)(nil),             // 2: chef.automate.api.infra_proxy.response.Role
	(*ExpandedRunList)(nil),  // 3: chef.automate.api.infra_proxy.response.ExpandedRunList
	(*RoleEnvironments)(nil), // 4: chef.automate.api.infra_proxy.response.RoleEnvironments
	(*RunList)(nil),          // 5: chef.automate.api.infra_proxy.response.RunList
}
var file_external_infra_proxy_response_roles_proto_depIdxs = []int32{
	1, // 0: chef.automate.api.infra_proxy.response.Roles.roles:type_name -> chef.automate.api.infra_proxy.response.RoleListItem
	5, // 1: chef.automate.api.infra_proxy.response.ExpandedRunList.run_list:type_name -> chef.automate.api.infra_proxy.response.RunList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_response_roles_proto_init() }
func file_external_infra_proxy_response_roles_proto_init() {
	if File_external_infra_proxy_response_roles_proto != nil {
		return
	}
	file_external_infra_proxy_response_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_response_roles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roles); i {
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
		file_external_infra_proxy_response_roles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleListItem); i {
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
		file_external_infra_proxy_response_roles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_external_infra_proxy_response_roles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedRunList); i {
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
		file_external_infra_proxy_response_roles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleEnvironments); i {
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
			RawDescriptor: file_external_infra_proxy_response_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_response_roles_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_response_roles_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_response_roles_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_response_roles_proto = out.File
	file_external_infra_proxy_response_roles_proto_rawDesc = nil
	file_external_infra_proxy_response_roles_proto_goTypes = nil
	file_external_infra_proxy_response_roles_proto_depIdxs = nil
}
