// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/account/account.proto

package account

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

type Role int32

const (
	Role_Admin     Role = 0
	Role_Consensus Role = 1
	Role_Ordinary  Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "Admin",
		1: "Consensus",
		2: "Ordinary",
	}
	Role_value = map[string]int32{
		"Admin":     0,
		"Consensus": 1,
		"Ordinary":  2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_api_account_account_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_api_account_account_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_api_account_account_proto_rawDescGZIP(), []int{0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Chain       string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	KmsPassword string `protobuf:"bytes,4,opt,name=kmsPassword,proto3" json:"kmsPassword,omitempty"`
	Role        Role   `protobuf:"varint,5,opt,name=role,proto3,enum=account.Role" json:"role,omitempty"`
	Domain      string `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty"`
	Address     string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_api_account_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Account) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *Account) GetKmsPassword() string {
	if x != nil {
		return x.KmsPassword
	}
	return ""
}

func (x *Account) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_Admin
}

func (x *Account) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AccountSampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *AccountSampleResponse) Reset() {
	*x = AccountSampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountSampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountSampleResponse) ProtoMessage() {}

func (x *AccountSampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountSampleResponse.ProtoReflect.Descriptor instead.
func (*AccountSampleResponse) Descriptor() ([]byte, []int) {
	return file_api_account_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountSampleResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountSampleResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ListAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Chain     string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (x *ListAccountRequest) Reset() {
	*x = ListAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountRequest) ProtoMessage() {}

func (x *ListAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountRequest.ProtoReflect.Descriptor instead.
func (*ListAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_account_proto_rawDescGZIP(), []int{2}
}

func (x *ListAccountRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListAccountRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

type AccountList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *AccountList) Reset() {
	*x = AccountList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountList) ProtoMessage() {}

func (x *AccountList) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountList.ProtoReflect.Descriptor instead.
func (*AccountList) Descriptor() ([]byte, []int) {
	return file_api_account_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountList) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_api_account_account_proto protoreflect.FileDescriptor

var file_api_account_account_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc8, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x6d, 0x73, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b,
	0x6d, 0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x49, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x22, 0x3b, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2a, 0x2e, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x10,
	0x02, 0x32, 0x87, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x61, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_account_account_proto_rawDescOnce sync.Once
	file_api_account_account_proto_rawDescData = file_api_account_account_proto_rawDesc
)

func file_api_account_account_proto_rawDescGZIP() []byte {
	file_api_account_account_proto_rawDescOnce.Do(func() {
		file_api_account_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_account_account_proto_rawDescData)
	})
	return file_api_account_account_proto_rawDescData
}

var file_api_account_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_account_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_account_account_proto_goTypes = []interface{}{
	(Role)(0),                     // 0: account.Role
	(*Account)(nil),               // 1: account.Account
	(*AccountSampleResponse)(nil), // 2: account.AccountSampleResponse
	(*ListAccountRequest)(nil),    // 3: account.ListAccountRequest
	(*AccountList)(nil),           // 4: account.AccountList
}
var file_api_account_account_proto_depIdxs = []int32{
	0, // 0: account.Account.role:type_name -> account.Role
	1, // 1: account.AccountList.accounts:type_name -> account.Account
	1, // 2: account.AccountService.CreateAccount:input_type -> account.Account
	3, // 3: account.AccountService.ListAccount:input_type -> account.ListAccountRequest
	1, // 4: account.AccountService.CreateAccount:output_type -> account.Account
	4, // 5: account.AccountService.ListAccount:output_type -> account.AccountList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_account_account_proto_init() }
func file_api_account_account_proto_init() {
	if File_api_account_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_account_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_api_account_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountSampleResponse); i {
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
		file_api_account_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountRequest); i {
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
		file_api_account_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountList); i {
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
			RawDescriptor: file_api_account_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_account_account_proto_goTypes,
		DependencyIndexes: file_api_account_account_proto_depIdxs,
		EnumInfos:         file_api_account_account_proto_enumTypes,
		MessageInfos:      file_api_account_account_proto_msgTypes,
	}.Build()
	File_api_account_account_proto = out.File
	file_api_account_account_proto_rawDesc = nil
	file_api_account_account_proto_goTypes = nil
	file_api_account_account_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
	ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*AccountList, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/account.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*AccountList, error) {
	out := new(AccountList)
	err := c.cc.Invoke(ctx, "/account.AccountService/ListAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	CreateAccount(context.Context, *Account) (*Account, error)
	ListAccount(context.Context, *ListAccountRequest) (*AccountList, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) CreateAccount(context.Context, *Account) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (*UnimplementedAccountServiceServer) ListAccount(context.Context, *ListAccountRequest) (*AccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/ListAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ListAccount(ctx, req.(*ListAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _AccountService_ListAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/account/account.proto",
}
