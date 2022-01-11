// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/chain/chain.proto

package chain

import (
	context "context"
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

type ConsensusType int32

const (
	ConsensusType_BFT  ConsensusType = 0
	ConsensusType_Raft ConsensusType = 1
)

// Enum value maps for ConsensusType.
var (
	ConsensusType_name = map[int32]string{
		0: "BFT",
		1: "Raft",
	}
	ConsensusType_value = map[string]int32{
		"BFT":  0,
		"Raft": 1,
	}
)

func (x ConsensusType) Enum() *ConsensusType {
	p := new(ConsensusType)
	*p = x
	return p
}

func (x ConsensusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConsensusType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_chain_chain_proto_enumTypes[0].Descriptor()
}

func (ConsensusType) Type() protoreflect.EnumType {
	return &file_api_chain_chain_proto_enumTypes[0]
}

func (x ConsensusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConsensusType.Descriptor instead.
func (ConsensusType) EnumDescriptor() ([]byte, []int) {
	return file_api_chain_chain_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_Online      Status = 0
	Status_Publicizing Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Online",
		1: "Publicizing",
	}
	Status_value = map[string]int32{
		"Online":      0,
		"Publicizing": 1,
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
	return file_api_chain_chain_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_chain_chain_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_chain_chain_proto_rawDescGZIP(), []int{1}
}

type Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace       string        `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id              string        `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp       int64         `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevHash        string        `protobuf:"bytes,5,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	BlockInterval   int32         `protobuf:"varint,6,opt,name=blockInterval,proto3" json:"blockInterval,omitempty"`
	BlockLimit      int32         `protobuf:"varint,7,opt,name=blockLimit,proto3" json:"blockLimit,omitempty"`
	EnableTls       bool          `protobuf:"varint,8,opt,name=enableTls,proto3" json:"enableTls,omitempty"`
	ConsensusType   ConsensusType `protobuf:"varint,9,opt,name=consensusType,proto3,enum=chain.ConsensusType" json:"consensusType,omitempty"`
	NetworkImage    string        `protobuf:"bytes,10,opt,name=networkImage,proto3" json:"networkImage,omitempty"`
	ConsensusImage  string        `protobuf:"bytes,11,opt,name=consensusImage,proto3" json:"consensusImage,omitempty"`
	ExecutorImage   string        `protobuf:"bytes,12,opt,name=executorImage,proto3" json:"executorImage,omitempty"`
	StorageImage    string        `protobuf:"bytes,13,opt,name=storageImage,proto3" json:"storageImage,omitempty"`
	ControllerImage string        `protobuf:"bytes,14,opt,name=controllerImage,proto3" json:"controllerImage,omitempty"`
	KmsImage        string        `protobuf:"bytes,15,opt,name=kmsImage,proto3" json:"kmsImage,omitempty"`
}

func (x *Chain) Reset() {
	*x = Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chain_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain) ProtoMessage() {}

func (x *Chain) ProtoReflect() protoreflect.Message {
	mi := &file_api_chain_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain.ProtoReflect.Descriptor instead.
func (*Chain) Descriptor() ([]byte, []int) {
	return file_api_chain_chain_proto_rawDescGZIP(), []int{0}
}

func (x *Chain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chain) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Chain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chain) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Chain) GetPrevHash() string {
	if x != nil {
		return x.PrevHash
	}
	return ""
}

func (x *Chain) GetBlockInterval() int32 {
	if x != nil {
		return x.BlockInterval
	}
	return 0
}

func (x *Chain) GetBlockLimit() int32 {
	if x != nil {
		return x.BlockLimit
	}
	return 0
}

func (x *Chain) GetEnableTls() bool {
	if x != nil {
		return x.EnableTls
	}
	return false
}

func (x *Chain) GetConsensusType() ConsensusType {
	if x != nil {
		return x.ConsensusType
	}
	return ConsensusType_BFT
}

func (x *Chain) GetNetworkImage() string {
	if x != nil {
		return x.NetworkImage
	}
	return ""
}

func (x *Chain) GetConsensusImage() string {
	if x != nil {
		return x.ConsensusImage
	}
	return ""
}

func (x *Chain) GetExecutorImage() string {
	if x != nil {
		return x.ExecutorImage
	}
	return ""
}

func (x *Chain) GetStorageImage() string {
	if x != nil {
		return x.StorageImage
	}
	return ""
}

func (x *Chain) GetControllerImage() string {
	if x != nil {
		return x.ControllerImage
	}
	return ""
}

func (x *Chain) GetKmsImage() string {
	if x != nil {
		return x.KmsImage
	}
	return ""
}

type ChainSimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Status    Status `protobuf:"varint,3,opt,name=status,proto3,enum=chain.Status" json:"status,omitempty"`
}

func (x *ChainSimpleResponse) Reset() {
	*x = ChainSimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chain_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainSimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainSimpleResponse) ProtoMessage() {}

func (x *ChainSimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chain_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainSimpleResponse.ProtoReflect.Descriptor instead.
func (*ChainSimpleResponse) Descriptor() ([]byte, []int) {
	return file_api_chain_chain_proto_rawDescGZIP(), []int{1}
}

func (x *ChainSimpleResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChainSimpleResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ChainSimpleResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Online
}

type ListChainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ListChainRequest) Reset() {
	*x = ListChainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chain_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChainRequest) ProtoMessage() {}

func (x *ListChainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chain_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChainRequest.ProtoReflect.Descriptor instead.
func (*ListChainRequest) Descriptor() ([]byte, []int) {
	return file_api_chain_chain_proto_rawDescGZIP(), []int{2}
}

func (x *ListChainRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ChainList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chains []*ChainSimpleResponse `protobuf:"bytes,1,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (x *ChainList) Reset() {
	*x = ChainList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chain_chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainList) ProtoMessage() {}

func (x *ChainList) ProtoReflect() protoreflect.Message {
	mi := &file_api_chain_chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainList.ProtoReflect.Descriptor instead.
func (*ChainList) Descriptor() ([]byte, []int) {
	return file_api_chain_chain_proto_rawDescGZIP(), []int{3}
}

func (x *ChainList) GetChains() []*ChainSimpleResponse {
	if x != nil {
		return x.Chains
	}
	return nil
}

type ChainOnlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ChainOnlineRequest) Reset() {
	*x = ChainOnlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chain_chain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainOnlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainOnlineRequest) ProtoMessage() {}

func (x *ChainOnlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chain_chain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainOnlineRequest.ProtoReflect.Descriptor instead.
func (*ChainOnlineRequest) Descriptor() ([]byte, []int) {
	return file_api_chain_chain_proto_rawDescGZIP(), []int{4}
}

func (x *ChainOnlineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChainOnlineRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_api_chain_chain_proto protoreflect.FileDescriptor

var file_api_chain_chain_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0xff,
	0x03, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6c, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x6d, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x6d, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x6e, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x30, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x3f, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2a, 0x22, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x42, 0x46, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x61, 0x66, 0x74, 0x10, 0x01, 0x2a,
	0x25, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69,
	0x7a, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x32, 0xb4, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12,
	0x0c, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x1a, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x61,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2d,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chain_chain_proto_rawDescOnce sync.Once
	file_api_chain_chain_proto_rawDescData = file_api_chain_chain_proto_rawDesc
)

func file_api_chain_chain_proto_rawDescGZIP() []byte {
	file_api_chain_chain_proto_rawDescOnce.Do(func() {
		file_api_chain_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chain_chain_proto_rawDescData)
	})
	return file_api_chain_chain_proto_rawDescData
}

var file_api_chain_chain_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_chain_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_chain_chain_proto_goTypes = []interface{}{
	(ConsensusType)(0),          // 0: chain.ConsensusType
	(Status)(0),                 // 1: chain.Status
	(*Chain)(nil),               // 2: chain.Chain
	(*ChainSimpleResponse)(nil), // 3: chain.ChainSimpleResponse
	(*ListChainRequest)(nil),    // 4: chain.ListChainRequest
	(*ChainList)(nil),           // 5: chain.ChainList
	(*ChainOnlineRequest)(nil),  // 6: chain.ChainOnlineRequest
}
var file_api_chain_chain_proto_depIdxs = []int32{
	0, // 0: chain.Chain.consensusType:type_name -> chain.ConsensusType
	1, // 1: chain.ChainSimpleResponse.status:type_name -> chain.Status
	3, // 2: chain.ChainList.chains:type_name -> chain.ChainSimpleResponse
	2, // 3: chain.ChainService.Init:input_type -> chain.Chain
	4, // 4: chain.ChainService.List:input_type -> chain.ListChainRequest
	6, // 5: chain.ChainService.Online:input_type -> chain.ChainOnlineRequest
	3, // 6: chain.ChainService.Init:output_type -> chain.ChainSimpleResponse
	5, // 7: chain.ChainService.List:output_type -> chain.ChainList
	3, // 8: chain.ChainService.Online:output_type -> chain.ChainSimpleResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_chain_chain_proto_init() }
func file_api_chain_chain_proto_init() {
	if File_api_chain_chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chain_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain); i {
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
		file_api_chain_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainSimpleResponse); i {
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
		file_api_chain_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChainRequest); i {
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
		file_api_chain_chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainList); i {
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
		file_api_chain_chain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainOnlineRequest); i {
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
			RawDescriptor: file_api_chain_chain_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chain_chain_proto_goTypes,
		DependencyIndexes: file_api_chain_chain_proto_depIdxs,
		EnumInfos:         file_api_chain_chain_proto_enumTypes,
		MessageInfos:      file_api_chain_chain_proto_msgTypes,
	}.Build()
	File_api_chain_chain_proto = out.File
	file_api_chain_chain_proto_rawDesc = nil
	file_api_chain_chain_proto_goTypes = nil
	file_api_chain_chain_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChainServiceClient is the client API for ChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainServiceClient interface {
	Init(ctx context.Context, in *Chain, opts ...grpc.CallOption) (*ChainSimpleResponse, error)
	List(ctx context.Context, in *ListChainRequest, opts ...grpc.CallOption) (*ChainList, error)
	Online(ctx context.Context, in *ChainOnlineRequest, opts ...grpc.CallOption) (*ChainSimpleResponse, error)
}

type chainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChainServiceClient(cc grpc.ClientConnInterface) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) Init(ctx context.Context, in *Chain, opts ...grpc.CallOption) (*ChainSimpleResponse, error) {
	out := new(ChainSimpleResponse)
	err := c.cc.Invoke(ctx, "/chain.ChainService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) List(ctx context.Context, in *ListChainRequest, opts ...grpc.CallOption) (*ChainList, error) {
	out := new(ChainList)
	err := c.cc.Invoke(ctx, "/chain.ChainService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) Online(ctx context.Context, in *ChainOnlineRequest, opts ...grpc.CallOption) (*ChainSimpleResponse, error) {
	out := new(ChainSimpleResponse)
	err := c.cc.Invoke(ctx, "/chain.ChainService/Online", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServiceServer is the server API for ChainService service.
type ChainServiceServer interface {
	Init(context.Context, *Chain) (*ChainSimpleResponse, error)
	List(context.Context, *ListChainRequest) (*ChainList, error)
	Online(context.Context, *ChainOnlineRequest) (*ChainSimpleResponse, error)
}

// UnimplementedChainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChainServiceServer struct {
}

func (*UnimplementedChainServiceServer) Init(context.Context, *Chain) (*ChainSimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedChainServiceServer) List(context.Context, *ListChainRequest) (*ChainList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedChainServiceServer) Online(context.Context, *ChainOnlineRequest) (*ChainSimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Online not implemented")
}

func RegisterChainServiceServer(s *grpc.Server, srv ChainServiceServer) {
	s.RegisterService(&_ChainService_serviceDesc, srv)
}

func _ChainService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.ChainService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).Init(ctx, req.(*Chain))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.ChainService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).List(ctx, req.(*ListChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_Online_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainOnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).Online(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.ChainService/Online",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).Online(ctx, req.(*ChainOnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chain.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ChainService_Init_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ChainService_List_Handler,
		},
		{
			MethodName: "Online",
			Handler:    _ChainService_Online_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/chain/chain.proto",
}