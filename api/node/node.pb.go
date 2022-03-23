// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/node/node.proto

package node

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_Unknown         Status = 0
	Status_WaitChainOnline Status = 1
	Status_Initialized     Status = 2
	Status_Starting        Status = 3
	Status_Running         Status = 4
	Status_Warning         Status = 5
	Status_Error           Status = 6
	Status_Updating        Status = 7
	Status_Stopping        Status = 8
	Status_Stopped         Status = 9
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "WaitChainOnline",
		2: "Initialized",
		3: "Starting",
		4: "Running",
		5: "Warning",
		6: "Error",
		7: "Updating",
		8: "Stopping",
		9: "Stopped",
	}
	Status_value = map[string]int32{
		"Unknown":         0,
		"WaitChainOnline": 1,
		"Initialized":     2,
		"Starting":        3,
		"Running":         4,
		"Warning":         5,
		"Error":           6,
		"Updating":        7,
		"Stopping":        8,
		"Stopped":         9,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_node_node_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_node_node_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{0}
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace        string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Cluster          string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Account          string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	ExternalIp       string `protobuf:"bytes,5,opt,name=externalIp,proto3" json:"externalIp,omitempty"`
	Port             int32  `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	Chain            string `protobuf:"bytes,7,opt,name=chain,proto3" json:"chain,omitempty"`
	StorageSize      int64  `protobuf:"varint,13,opt,name=storageSize,proto3" json:"storageSize,omitempty"`
	StorageClassName string `protobuf:"bytes,14,opt,name=storageClassName,proto3" json:"storageClassName,omitempty"`
	LogLevel         string `protobuf:"bytes,15,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
	Status           Status `protobuf:"varint,16,opt,name=status,proto3,enum=node.Status" json:"status,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[0]
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
	return file_api_node_node_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Node) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Node) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Node) GetExternalIp() string {
	if x != nil {
		return x.ExternalIp
	}
	return ""
}

func (x *Node) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Node) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *Node) GetStorageSize() int64 {
	if x != nil {
		return x.StorageSize
	}
	return 0
}

func (x *Node) GetStorageClassName() string {
	if x != nil {
		return x.StorageClassName
	}
	return ""
}

func (x *Node) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *Node) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Unknown
}

type NodeSimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Status    Status `protobuf:"varint,3,opt,name=status,proto3,enum=node.Status" json:"status,omitempty"`
}

func (x *NodeSimpleResponse) Reset() {
	*x = NodeSimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSimpleResponse) ProtoMessage() {}

func (x *NodeSimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSimpleResponse.ProtoReflect.Descriptor instead.
func (*NodeSimpleResponse) Descriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{1}
}

func (x *NodeSimpleResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeSimpleResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NodeSimpleResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Unknown
}

type ListNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Chain     string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (x *ListNodeRequest) Reset() {
	*x = ListNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeRequest) ProtoMessage() {}

func (x *ListNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeRequest.ProtoReflect.Descriptor instead.
func (*ListNodeRequest) Descriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{2}
}

func (x *ListNodeRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListNodeRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

type NodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodeList) Reset() {
	*x = NodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeList) ProtoMessage() {}

func (x *NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeList.ProtoReflect.Descriptor instead.
func (*NodeList) Descriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{3}
}

func (x *NodeList) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type NodeStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *NodeStartRequest) Reset() {
	*x = NodeStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStartRequest) ProtoMessage() {}

func (x *NodeStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStartRequest.ProtoReflect.Descriptor instead.
func (*NodeStartRequest) Descriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{4}
}

func (x *NodeStartRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeStartRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NodeStopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *NodeStopRequest) Reset() {
	*x = NodeStopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStopRequest) ProtoMessage() {}

func (x *NodeStopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStopRequest.ProtoReflect.Descriptor instead.
func (*NodeStopRequest) Descriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{5}
}

func (x *NodeStopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeStopRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ReloadConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ReloadConfigRequest) Reset() {
	*x = ReloadConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadConfigRequest) ProtoMessage() {}

func (x *ReloadConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadConfigRequest.ProtoReflect.Descriptor instead.
func (*ReloadConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{6}
}

func (x *ReloadConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReloadConfigRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NodeDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *NodeDeleteRequest) Reset() {
	*x = NodeDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_node_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeDeleteRequest) ProtoMessage() {}

func (x *NodeDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_node_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeDeleteRequest.ProtoReflect.Descriptor instead.
func (*NodeDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_node_node_proto_rawDescGZIP(), []int{7}
}

func (x *NodeDeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeDeleteRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_api_node_node_proto protoreflect.FileDescriptor

var file_api_node_node_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x6c, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x45, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x2c, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x43, 0x0a, 0x0f, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22,
	0x47, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2a,
	0x97, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x61, 0x69, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x06,
	0x12, 0x0c, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x07, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x10, 0x09, 0x32, 0xda, 0x02, 0x0a, 0x0b, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x49, 0x6e, 0x69,
	0x74, 0x12, 0x0a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x18, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x16, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x15, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x61, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_node_node_proto_rawDescOnce sync.Once
	file_api_node_node_proto_rawDescData = file_api_node_node_proto_rawDesc
)

func file_api_node_node_proto_rawDescGZIP() []byte {
	file_api_node_node_proto_rawDescOnce.Do(func() {
		file_api_node_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_node_node_proto_rawDescData)
	})
	return file_api_node_node_proto_rawDescData
}

var file_api_node_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_node_node_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_node_node_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: node.Status
	(*Node)(nil),                // 1: node.Node
	(*NodeSimpleResponse)(nil),  // 2: node.NodeSimpleResponse
	(*ListNodeRequest)(nil),     // 3: node.ListNodeRequest
	(*NodeList)(nil),            // 4: node.NodeList
	(*NodeStartRequest)(nil),    // 5: node.NodeStartRequest
	(*NodeStopRequest)(nil),     // 6: node.NodeStopRequest
	(*ReloadConfigRequest)(nil), // 7: node.ReloadConfigRequest
	(*NodeDeleteRequest)(nil),   // 8: node.NodeDeleteRequest
	(*emptypb.Empty)(nil),       // 9: google.protobuf.Empty
}
var file_api_node_node_proto_depIdxs = []int32{
	0, // 0: node.Node.status:type_name -> node.Status
	0, // 1: node.NodeSimpleResponse.status:type_name -> node.Status
	1, // 2: node.NodeList.nodes:type_name -> node.Node
	1, // 3: node.NodeService.Init:input_type -> node.Node
	3, // 4: node.NodeService.List:input_type -> node.ListNodeRequest
	5, // 5: node.NodeService.Start:input_type -> node.NodeStartRequest
	6, // 6: node.NodeService.Stop:input_type -> node.NodeStopRequest
	7, // 7: node.NodeService.ReloadConfig:input_type -> node.ReloadConfigRequest
	8, // 8: node.NodeService.Delete:input_type -> node.NodeDeleteRequest
	2, // 9: node.NodeService.Init:output_type -> node.NodeSimpleResponse
	4, // 10: node.NodeService.List:output_type -> node.NodeList
	2, // 11: node.NodeService.Start:output_type -> node.NodeSimpleResponse
	9, // 12: node.NodeService.Stop:output_type -> google.protobuf.Empty
	9, // 13: node.NodeService.ReloadConfig:output_type -> google.protobuf.Empty
	9, // 14: node.NodeService.Delete:output_type -> google.protobuf.Empty
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_node_node_proto_init() }
func file_api_node_node_proto_init() {
	if File_api_node_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_node_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_node_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSimpleResponse); i {
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
		file_api_node_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeRequest); i {
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
		file_api_node_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeList); i {
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
		file_api_node_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStartRequest); i {
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
		file_api_node_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStopRequest); i {
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
		file_api_node_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadConfigRequest); i {
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
		file_api_node_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeDeleteRequest); i {
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
			RawDescriptor: file_api_node_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_node_node_proto_goTypes,
		DependencyIndexes: file_api_node_node_proto_depIdxs,
		EnumInfos:         file_api_node_node_proto_enumTypes,
		MessageInfos:      file_api_node_node_proto_msgTypes,
	}.Build()
	File_api_node_node_proto = out.File
	file_api_node_node_proto_rawDesc = nil
	file_api_node_node_proto_goTypes = nil
	file_api_node_node_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	Init(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NodeSimpleResponse, error)
	List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*NodeList, error)
	Start(ctx context.Context, in *NodeStartRequest, opts ...grpc.CallOption) (*NodeSimpleResponse, error)
	Stop(ctx context.Context, in *NodeStopRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReloadConfig(ctx context.Context, in *ReloadConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *NodeDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Init(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NodeSimpleResponse, error) {
	out := new(NodeSimpleResponse)
	err := c.cc.Invoke(ctx, "/node.NodeService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*NodeList, error) {
	out := new(NodeList)
	err := c.cc.Invoke(ctx, "/node.NodeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Start(ctx context.Context, in *NodeStartRequest, opts ...grpc.CallOption) (*NodeSimpleResponse, error) {
	out := new(NodeSimpleResponse)
	err := c.cc.Invoke(ctx, "/node.NodeService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Stop(ctx context.Context, in *NodeStopRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/node.NodeService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ReloadConfig(ctx context.Context, in *ReloadConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/node.NodeService/ReloadConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Delete(ctx context.Context, in *NodeDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/node.NodeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	Init(context.Context, *Node) (*NodeSimpleResponse, error)
	List(context.Context, *ListNodeRequest) (*NodeList, error)
	Start(context.Context, *NodeStartRequest) (*NodeSimpleResponse, error)
	Stop(context.Context, *NodeStopRequest) (*emptypb.Empty, error)
	ReloadConfig(context.Context, *ReloadConfigRequest) (*emptypb.Empty, error)
	Delete(context.Context, *NodeDeleteRequest) (*emptypb.Empty, error)
}

// UnimplementedNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (*UnimplementedNodeServiceServer) Init(context.Context, *Node) (*NodeSimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedNodeServiceServer) List(context.Context, *ListNodeRequest) (*NodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedNodeServiceServer) Start(context.Context, *NodeStartRequest) (*NodeSimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedNodeServiceServer) Stop(context.Context, *NodeStopRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedNodeServiceServer) ReloadConfig(context.Context, *ReloadConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadConfig not implemented")
}
func (*UnimplementedNodeServiceServer) Delete(context.Context, *NodeDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Init(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).List(ctx, req.(*ListNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Start(ctx, req.(*NodeStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Stop(ctx, req.(*NodeStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ReloadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ReloadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeService/ReloadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ReloadConfig(ctx, req.(*ReloadConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Delete(ctx, req.(*NodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _NodeService_Init_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NodeService_List_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _NodeService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _NodeService_Stop_Handler,
		},
		{
			MethodName: "ReloadConfig",
			Handler:    _NodeService_ReloadConfig_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NodeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/node/node.proto",
}
