// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: acl-config.proto

package proto

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Config stores ACL configuration parameters.
type ACLConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ccName - should be empty, field will be filled while Init transaction, filler will be ignored
	CcName string `protobuf:"bytes,1,opt,name=ccName,proto3" json:"ccName,omitempty"`
	// adminSKIEncoded - ACL admin public key SKI in hex format
	AdminSKIEncoded string `protobuf:"bytes,2,opt,name=adminSKIEncoded,proto3" json:"adminSKIEncoded,omitempty"`
	// validators - set of validators public keys
	Validators []*ACLValidator `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (x *ACLConfig) Reset() {
	*x = ACLConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLConfig) ProtoMessage() {}

func (x *ACLConfig) ProtoReflect() protoreflect.Message {
	mi := &file_acl_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLConfig.ProtoReflect.Descriptor instead.
func (*ACLConfig) Descriptor() ([]byte, []int) {
	return file_acl_config_proto_rawDescGZIP(), []int{0}
}

func (x *ACLConfig) GetCcName() string {
	if x != nil {
		return x.CcName
	}
	return ""
}

func (x *ACLConfig) GetAdminSKIEncoded() string {
	if x != nil {
		return x.AdminSKIEncoded
	}
	return ""
}

func (x *ACLConfig) GetValidators() []*ACLValidator {
	if x != nil {
		return x.Validators
	}
	return nil
}

type ACLValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	KeyType   string `protobuf:"bytes,2,opt,name=keyType,proto3" json:"keyType,omitempty"`
}

func (x *ACLValidator) Reset() {
	*x = ACLValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLValidator) ProtoMessage() {}

func (x *ACLValidator) ProtoReflect() protoreflect.Message {
	mi := &file_acl_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLValidator.ProtoReflect.Descriptor instead.
func (*ACLValidator) Descriptor() ([]byte, []int) {
	return file_acl_config_proto_rawDescGZIP(), []int{1}
}

func (x *ACLValidator) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *ACLValidator) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

var File_acl_config_proto protoreflect.FileDescriptor

var file_acl_config_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x63, 0x6c, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x09, 0x41, 0x43, 0x4c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x53, 0x4b, 0x49, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x12, 0xfa, 0x42, 0x0f, 0x72, 0x0d, 0x32, 0x0b, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x61,
	0x2d, 0x66, 0x5d, 0x2b, 0x24, 0x52, 0x0f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x4b, 0x49, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x43, 0x4c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x46, 0x0a, 0x0c, 0x41,
	0x43, 0x4c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x6f, 0x69, 0x64, 0x65, 0x61, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x61, 0x63,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acl_config_proto_rawDescOnce sync.Once
	file_acl_config_proto_rawDescData = file_acl_config_proto_rawDesc
)

func file_acl_config_proto_rawDescGZIP() []byte {
	file_acl_config_proto_rawDescOnce.Do(func() {
		file_acl_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_acl_config_proto_rawDescData)
	})
	return file_acl_config_proto_rawDescData
}

var file_acl_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_acl_config_proto_goTypes = []interface{}{
	(*ACLConfig)(nil),    // 0: proto.ACLConfig
	(*ACLValidator)(nil), // 1: proto.ACLValidator
}
var file_acl_config_proto_depIdxs = []int32{
	1, // 0: proto.ACLConfig.validators:type_name -> proto.ACLValidator
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_acl_config_proto_init() }
func file_acl_config_proto_init() {
	if File_acl_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acl_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLConfig); i {
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
		file_acl_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLValidator); i {
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
			RawDescriptor: file_acl_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acl_config_proto_goTypes,
		DependencyIndexes: file_acl_config_proto_depIdxs,
		MessageInfos:      file_acl_config_proto_msgTypes,
	}.Build()
	File_acl_config_proto = out.File
	file_acl_config_proto_rawDesc = nil
	file_acl_config_proto_goTypes = nil
	file_acl_config_proto_depIdxs = nil
}
