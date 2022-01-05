// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: external/infra_proxy/response/nodes.proto

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

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node list.
	Nodes []*NodeAttribute `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// Starting page for the results.
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// Total number of records.
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *Nodes) GetNodes() []*NodeAttribute {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Nodes) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Nodes) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type NodeAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Node last checkin.
	CheckIn string `protobuf:"bytes,3,opt,name=check_in,json=checkIn,proto3" json:"check_in,omitempty"`
	// Node uptime.
	Uptime string `protobuf:"bytes,4,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// Node platform.
	Platform string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty"`
	// Node environment name.
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	// Node policy group.
	PolicyGroup string `protobuf:"bytes,7,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	// Node FQDN.
	Fqdn string `protobuf:"bytes,8,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// Node IP address.
	IpAddress string `protobuf:"bytes,9,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *NodeAttribute) Reset() {
	*x = NodeAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAttribute) ProtoMessage() {}

func (x *NodeAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAttribute.ProtoReflect.Descriptor instead.
func (*NodeAttribute) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{1}
}

func (x *NodeAttribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeAttribute) GetCheckIn() string {
	if x != nil {
		return x.CheckIn
	}
	return ""
}

func (x *NodeAttribute) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

func (x *NodeAttribute) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *NodeAttribute) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *NodeAttribute) GetPolicyGroup() string {
	if x != nil {
		return x.PolicyGroup
	}
	return ""
}

func (x *NodeAttribute) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *NodeAttribute) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type DeleteNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteNode) Reset() {
	*x = DeleteNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNode) ProtoMessage() {}

func (x *DeleteNode) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNode.ProtoReflect.Descriptor instead.
func (*DeleteNode) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node ID.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Node environment.
	Environment string `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	// Node policy name.
	PolicyName string `protobuf:"bytes,4,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// Node policy group.
	PolicyGroup string `protobuf:"bytes,5,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	// Node run-list.
	RunList []string `protobuf:"bytes,6,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	// Node tags.
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// Node automatic attributes JSON.
	AutomaticAttributes string `protobuf:"bytes,8,opt,name=automatic_attributes,json=automaticAttributes,proto3" json:"automatic_attributes,omitempty"`
	// Node default attributes JSON.
	DefaultAttributes string `protobuf:"bytes,9,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Node normal attributes JSON.
	NormalAttributes string `protobuf:"bytes,10,opt,name=normal_attributes,json=normalAttributes,proto3" json:"normal_attributes,omitempty"`
	// Node override attributes JSON.
	OverrideAttributes string `protobuf:"bytes,11,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{3}
}

func (x *Node) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *Node) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *Node) GetPolicyGroup() string {
	if x != nil {
		return x.PolicyGroup
	}
	return ""
}

func (x *Node) GetRunList() []string {
	if x != nil {
		return x.RunList
	}
	return nil
}

func (x *Node) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Node) GetAutomaticAttributes() string {
	if x != nil {
		return x.AutomaticAttributes
	}
	return ""
}

func (x *Node) GetDefaultAttributes() string {
	if x != nil {
		return x.DefaultAttributes
	}
	return ""
}

func (x *Node) GetNormalAttributes() string {
	if x != nil {
		return x.NormalAttributes
	}
	return ""
}

func (x *Node) GetOverrideAttributes() string {
	if x != nil {
		return x.OverrideAttributes
	}
	return ""
}

type UpdateNodeTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node tags.
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateNodeTags) Reset() {
	*x = UpdateNodeTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeTags) ProtoMessage() {}

func (x *UpdateNodeTags) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeTags.ProtoReflect.Descriptor instead.
func (*UpdateNodeTags) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNodeTags) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateNodeEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Node environment name.
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *UpdateNodeEnvironment) Reset() {
	*x = UpdateNodeEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeEnvironment) ProtoMessage() {}

func (x *UpdateNodeEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeEnvironment.ProtoReflect.Descriptor instead.
func (*UpdateNodeEnvironment) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNodeEnvironment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNodeEnvironment) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type UpdateNodeAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Node attributes JSON.
	Attributes string `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *UpdateNodeAttributes) Reset() {
	*x = UpdateNodeAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeAttributes) ProtoMessage() {}

func (x *UpdateNodeAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeAttributes.ProtoReflect.Descriptor instead.
func (*UpdateNodeAttributes) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNodeAttributes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNodeAttributes) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

type NodeExpandedRunList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the run list collection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// List of the run list.
	RunList []*RunList `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
}

func (x *NodeExpandedRunList) Reset() {
	*x = NodeExpandedRunList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeExpandedRunList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeExpandedRunList) ProtoMessage() {}

func (x *NodeExpandedRunList) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_nodes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeExpandedRunList.ProtoReflect.Descriptor instead.
func (*NodeExpandedRunList) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_nodes_proto_rawDescGZIP(), []int{7}
}

func (x *NodeExpandedRunList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeExpandedRunList) GetRunList() []*RunList {
	if x != nil {
		return x.RunList
	}
	return nil
}

var File_external_infra_proxy_response_nodes_proto protoreflect.FileDescriptor

var file_external_infra_proxy_response_nodes_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7e, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0xfa, 0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x20, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x88,
	0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75,
	0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x4d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4a,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x13, 0x4e, 0x6f,
	0x64, 0x65, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x4a, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_external_infra_proxy_response_nodes_proto_rawDescOnce sync.Once
	file_external_infra_proxy_response_nodes_proto_rawDescData = file_external_infra_proxy_response_nodes_proto_rawDesc
)

func file_external_infra_proxy_response_nodes_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_response_nodes_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_response_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_response_nodes_proto_rawDescData)
	})
	return file_external_infra_proxy_response_nodes_proto_rawDescData
}

var file_external_infra_proxy_response_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_external_infra_proxy_response_nodes_proto_goTypes = []interface{}{
	(*Nodes)(nil),                 // 0: chef.automate.api.infra_proxy.response.Nodes
	(*NodeAttribute)(nil),         // 1: chef.automate.api.infra_proxy.response.NodeAttribute
	(*DeleteNode)(nil),            // 2: chef.automate.api.infra_proxy.response.DeleteNode
	(*Node)(nil),                  // 3: chef.automate.api.infra_proxy.response.Node
	(*UpdateNodeTags)(nil),        // 4: chef.automate.api.infra_proxy.response.UpdateNodeTags
	(*UpdateNodeEnvironment)(nil), // 5: chef.automate.api.infra_proxy.response.UpdateNodeEnvironment
	(*UpdateNodeAttributes)(nil),  // 6: chef.automate.api.infra_proxy.response.UpdateNodeAttributes
	(*NodeExpandedRunList)(nil),   // 7: chef.automate.api.infra_proxy.response.NodeExpandedRunList
	(*RunList)(nil),               // 8: chef.automate.api.infra_proxy.response.RunList
}
var file_external_infra_proxy_response_nodes_proto_depIdxs = []int32{
	1, // 0: chef.automate.api.infra_proxy.response.Nodes.nodes:type_name -> chef.automate.api.infra_proxy.response.NodeAttribute
	8, // 1: chef.automate.api.infra_proxy.response.NodeExpandedRunList.run_list:type_name -> chef.automate.api.infra_proxy.response.RunList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_response_nodes_proto_init() }
func file_external_infra_proxy_response_nodes_proto_init() {
	if File_external_infra_proxy_response_nodes_proto != nil {
		return
	}
	file_external_infra_proxy_response_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_response_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
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
		file_external_infra_proxy_response_nodes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAttribute); i {
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
		file_external_infra_proxy_response_nodes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNode); i {
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
		file_external_infra_proxy_response_nodes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_external_infra_proxy_response_nodes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeTags); i {
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
		file_external_infra_proxy_response_nodes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeEnvironment); i {
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
		file_external_infra_proxy_response_nodes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeAttributes); i {
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
		file_external_infra_proxy_response_nodes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeExpandedRunList); i {
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
			RawDescriptor: file_external_infra_proxy_response_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_response_nodes_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_response_nodes_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_response_nodes_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_response_nodes_proto = out.File
	file_external_infra_proxy_response_nodes_proto_rawDesc = nil
	file_external_infra_proxy_response_nodes_proto_goTypes = nil
	file_external_infra_proxy_response_nodes_proto_depIdxs = nil
}
