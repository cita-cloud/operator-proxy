// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/allinone/allinone.proto

package allinone

import (
	context "context"
	chain "github.com/cita-cloud/operator-proxy/api/chain"
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

type AllInOneCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace        string              `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id               string              `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp        int64               `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevHash         string              `protobuf:"bytes,5,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	BlockInterval    int32               `protobuf:"varint,6,opt,name=blockInterval,proto3" json:"blockInterval,omitempty"`
	BlockLimit       int32               `protobuf:"varint,7,opt,name=blockLimit,proto3" json:"blockLimit,omitempty"`
	EnableTls        bool                `protobuf:"varint,8,opt,name=enableTls,proto3" json:"enableTls,omitempty"`
	ConsensusType    chain.ConsensusType `protobuf:"varint,9,opt,name=consensusType,proto3,enum=chain.ConsensusType" json:"consensusType,omitempty"`
	NetworkImage     string              `protobuf:"bytes,10,opt,name=networkImage,proto3" json:"networkImage,omitempty"`
	ConsensusImage   string              `protobuf:"bytes,11,opt,name=consensusImage,proto3" json:"consensusImage,omitempty"`
	ExecutorImage    string              `protobuf:"bytes,12,opt,name=executorImage,proto3" json:"executorImage,omitempty"`
	StorageImage     string              `protobuf:"bytes,13,opt,name=storageImage,proto3" json:"storageImage,omitempty"`
	ControllerImage  string              `protobuf:"bytes,14,opt,name=controllerImage,proto3" json:"controllerImage,omitempty"`
	KmsImage         string              `protobuf:"bytes,15,opt,name=kmsImage,proto3" json:"kmsImage,omitempty"`
	NodeCount        int32               `protobuf:"varint,16,opt,name=nodeCount,proto3" json:"nodeCount,omitempty"`
	StorageSize      int64               `protobuf:"varint,17,opt,name=storageSize,proto3" json:"storageSize,omitempty"`
	StorageClassName string              `protobuf:"bytes,18,opt,name=storageClassName,proto3" json:"storageClassName,omitempty"`
	LogLevel         string              `protobuf:"bytes,19,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
	Version          string              `protobuf:"bytes,20,opt,name=version,proto3" json:"version,omitempty"`
	AdminAddress     string              `protobuf:"bytes,21,opt,name=adminAddress,proto3" json:"adminAddress,omitempty"`
}

func (x *AllInOneCreateRequest) Reset() {
	*x = AllInOneCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_allinone_allinone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllInOneCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllInOneCreateRequest) ProtoMessage() {}

func (x *AllInOneCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_allinone_allinone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllInOneCreateRequest.ProtoReflect.Descriptor instead.
func (*AllInOneCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_allinone_allinone_proto_rawDescGZIP(), []int{0}
}

func (x *AllInOneCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AllInOneCreateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AllInOneCreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AllInOneCreateRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AllInOneCreateRequest) GetPrevHash() string {
	if x != nil {
		return x.PrevHash
	}
	return ""
}

func (x *AllInOneCreateRequest) GetBlockInterval() int32 {
	if x != nil {
		return x.BlockInterval
	}
	return 0
}

func (x *AllInOneCreateRequest) GetBlockLimit() int32 {
	if x != nil {
		return x.BlockLimit
	}
	return 0
}

func (x *AllInOneCreateRequest) GetEnableTls() bool {
	if x != nil {
		return x.EnableTls
	}
	return false
}

func (x *AllInOneCreateRequest) GetConsensusType() chain.ConsensusType {
	if x != nil {
		return x.ConsensusType
	}
	return chain.ConsensusType_UnknownConsensusType
}

func (x *AllInOneCreateRequest) GetNetworkImage() string {
	if x != nil {
		return x.NetworkImage
	}
	return ""
}

func (x *AllInOneCreateRequest) GetConsensusImage() string {
	if x != nil {
		return x.ConsensusImage
	}
	return ""
}

func (x *AllInOneCreateRequest) GetExecutorImage() string {
	if x != nil {
		return x.ExecutorImage
	}
	return ""
}

func (x *AllInOneCreateRequest) GetStorageImage() string {
	if x != nil {
		return x.StorageImage
	}
	return ""
}

func (x *AllInOneCreateRequest) GetControllerImage() string {
	if x != nil {
		return x.ControllerImage
	}
	return ""
}

func (x *AllInOneCreateRequest) GetKmsImage() string {
	if x != nil {
		return x.KmsImage
	}
	return ""
}

func (x *AllInOneCreateRequest) GetNodeCount() int32 {
	if x != nil {
		return x.NodeCount
	}
	return 0
}

func (x *AllInOneCreateRequest) GetStorageSize() int64 {
	if x != nil {
		return x.StorageSize
	}
	return 0
}

func (x *AllInOneCreateRequest) GetStorageClassName() string {
	if x != nil {
		return x.StorageClassName
	}
	return ""
}

func (x *AllInOneCreateRequest) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *AllInOneCreateRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AllInOneCreateRequest) GetAdminAddress() string {
	if x != nil {
		return x.AdminAddress
	}
	return ""
}

type AllInOneCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *AllInOneCreateResponse) Reset() {
	*x = AllInOneCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_allinone_allinone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllInOneCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllInOneCreateResponse) ProtoMessage() {}

func (x *AllInOneCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_allinone_allinone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllInOneCreateResponse.ProtoReflect.Descriptor instead.
func (*AllInOneCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_allinone_allinone_proto_rawDescGZIP(), []int{1}
}

func (x *AllInOneCreateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AllInOneCreateResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_api_allinone_allinone_proto protoreflect.FileDescriptor

var file_api_allinone_allinone_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x6c, 0x6c, 0x69, 0x6e, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61,
	0x6c, 0x6c, 0x69, 0x6e, 0x6f, 0x6e, 0x65, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x05, 0x0a, 0x15, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x4f, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
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
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4a, 0x0a, 0x16, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x4f,
	0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x32, 0x5e, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x4f, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x2e, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x6f, 0x6e, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x49, 0x6e,
	0x4f, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x6f, 0x6e, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x49,
	0x6e, 0x4f, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x69, 0x74, 0x61, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x6c, 0x69, 0x6e, 0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_allinone_allinone_proto_rawDescOnce sync.Once
	file_api_allinone_allinone_proto_rawDescData = file_api_allinone_allinone_proto_rawDesc
)

func file_api_allinone_allinone_proto_rawDescGZIP() []byte {
	file_api_allinone_allinone_proto_rawDescOnce.Do(func() {
		file_api_allinone_allinone_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_allinone_allinone_proto_rawDescData)
	})
	return file_api_allinone_allinone_proto_rawDescData
}

var file_api_allinone_allinone_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_allinone_allinone_proto_goTypes = []interface{}{
	(*AllInOneCreateRequest)(nil),  // 0: allinone.AllInOneCreateRequest
	(*AllInOneCreateResponse)(nil), // 1: allinone.AllInOneCreateResponse
	(chain.ConsensusType)(0),       // 2: chain.ConsensusType
}
var file_api_allinone_allinone_proto_depIdxs = []int32{
	2, // 0: allinone.AllInOneCreateRequest.consensusType:type_name -> chain.ConsensusType
	0, // 1: allinone.AllInOneService.Create:input_type -> allinone.AllInOneCreateRequest
	1, // 2: allinone.AllInOneService.Create:output_type -> allinone.AllInOneCreateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_allinone_allinone_proto_init() }
func file_api_allinone_allinone_proto_init() {
	if File_api_allinone_allinone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_allinone_allinone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllInOneCreateRequest); i {
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
		file_api_allinone_allinone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllInOneCreateResponse); i {
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
			RawDescriptor: file_api_allinone_allinone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_allinone_allinone_proto_goTypes,
		DependencyIndexes: file_api_allinone_allinone_proto_depIdxs,
		MessageInfos:      file_api_allinone_allinone_proto_msgTypes,
	}.Build()
	File_api_allinone_allinone_proto = out.File
	file_api_allinone_allinone_proto_rawDesc = nil
	file_api_allinone_allinone_proto_goTypes = nil
	file_api_allinone_allinone_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AllInOneServiceClient is the client API for AllInOneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AllInOneServiceClient interface {
	Create(ctx context.Context, in *AllInOneCreateRequest, opts ...grpc.CallOption) (*AllInOneCreateResponse, error)
}

type allInOneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAllInOneServiceClient(cc grpc.ClientConnInterface) AllInOneServiceClient {
	return &allInOneServiceClient{cc}
}

func (c *allInOneServiceClient) Create(ctx context.Context, in *AllInOneCreateRequest, opts ...grpc.CallOption) (*AllInOneCreateResponse, error) {
	out := new(AllInOneCreateResponse)
	err := c.cc.Invoke(ctx, "/allinone.AllInOneService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllInOneServiceServer is the server API for AllInOneService service.
type AllInOneServiceServer interface {
	Create(context.Context, *AllInOneCreateRequest) (*AllInOneCreateResponse, error)
}

// UnimplementedAllInOneServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAllInOneServiceServer struct {
}

func (*UnimplementedAllInOneServiceServer) Create(context.Context, *AllInOneCreateRequest) (*AllInOneCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterAllInOneServiceServer(s *grpc.Server, srv AllInOneServiceServer) {
	s.RegisterService(&_AllInOneService_serviceDesc, srv)
}

func _AllInOneService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllInOneCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllInOneServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allinone.AllInOneService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllInOneServiceServer).Create(ctx, req.(*AllInOneCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AllInOneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allinone.AllInOneService",
	HandlerType: (*AllInOneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AllInOneService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/allinone/allinone.proto",
}
