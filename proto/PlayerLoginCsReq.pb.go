// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerLoginCsReq.proto

package proto

import (
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

type PlayerLoginCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceInfo            string       `protobuf:"bytes,2,opt,name=DeviceInfo,proto3" json:"DeviceInfo,omitempty"`
	Language              LanguageType `protobuf:"varint,12,opt,name=Language,proto3,enum=LanguageType" json:"Language,omitempty"`
	ClientVersion         string       `protobuf:"bytes,15,opt,name=ClientVersion,proto3" json:"ClientVersion,omitempty"`
	DeviceUuid            string       `protobuf:"bytes,5,opt,name=DeviceUuid,proto3" json:"DeviceUuid,omitempty"`
	Token                 string       `protobuf:"bytes,9,opt,name=Token,proto3" json:"Token,omitempty"`
	ClientDataVersion     uint32       `protobuf:"varint,10,opt,name=ClientDataVersion,proto3" json:"ClientDataVersion,omitempty"`
	PlatformType          PlatformType `protobuf:"varint,7,opt,name=PlatformType,proto3,enum=PlatformType" json:"PlatformType,omitempty"`
	Signature             string       `protobuf:"bytes,13,opt,name=Signature,proto3" json:"Signature,omitempty"`
	SystemVersion         string       `protobuf:"bytes,6,opt,name=SystemVersion,proto3" json:"SystemVersion,omitempty"`
	AccountType           uint32       `protobuf:"varint,1812,opt,name=AccountType,proto3" json:"AccountType,omitempty"`
	Null1                 bool         `protobuf:"varint,61,opt,name=Null1,proto3" json:"Null1,omitempty"`
	ChecksumClientVersion string       `protobuf:"bytes,11,opt,name=ChecksumClientVersion,proto3" json:"ChecksumClientVersion,omitempty"`
	Null2                 bool         `protobuf:"varint,1588,opt,name=Null2,proto3" json:"Null2,omitempty"`
	CurTimezone           string       `protobuf:"bytes,1759,opt,name=CurTimezone,proto3" json:"CurTimezone,omitempty"`
}

func (x *PlayerLoginCsReq) Reset() {
	*x = PlayerLoginCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLoginCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginCsReq) ProtoMessage() {}

func (x *PlayerLoginCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLoginCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginCsReq.ProtoReflect.Descriptor instead.
func (*PlayerLoginCsReq) Descriptor() ([]byte, []int) {
	return file_PlayerLoginCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLoginCsReq) GetDeviceInfo() string {
	if x != nil {
		return x.DeviceInfo
	}
	return ""
}

func (x *PlayerLoginCsReq) GetLanguage() LanguageType {
	if x != nil {
		return x.Language
	}
	return LanguageType_LANGUAGE_NONE
}

func (x *PlayerLoginCsReq) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *PlayerLoginCsReq) GetDeviceUuid() string {
	if x != nil {
		return x.DeviceUuid
	}
	return ""
}

func (x *PlayerLoginCsReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PlayerLoginCsReq) GetClientDataVersion() uint32 {
	if x != nil {
		return x.ClientDataVersion
	}
	return 0
}

func (x *PlayerLoginCsReq) GetPlatformType() PlatformType {
	if x != nil {
		return x.PlatformType
	}
	return PlatformType_EDITOR
}

func (x *PlayerLoginCsReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PlayerLoginCsReq) GetSystemVersion() string {
	if x != nil {
		return x.SystemVersion
	}
	return ""
}

func (x *PlayerLoginCsReq) GetAccountType() uint32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *PlayerLoginCsReq) GetNull1() bool {
	if x != nil {
		return x.Null1
	}
	return false
}

func (x *PlayerLoginCsReq) GetChecksumClientVersion() string {
	if x != nil {
		return x.ChecksumClientVersion
	}
	return ""
}

func (x *PlayerLoginCsReq) GetNull2() bool {
	if x != nil {
		return x.Null2
	}
	return false
}

func (x *PlayerLoginCsReq) GetCurTimezone() string {
	if x != nil {
		return x.CurTimezone
	}
	return ""
}

var File_PlayerLoginCsReq_proto protoreflect.FileDescriptor

var file_PlayerLoginCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x04, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x55, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x11,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x0c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x94, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x75, 0x6c, 0x6c, 0x31, 0x18, 0x3d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x4e, 0x75, 0x6c, 0x6c, 0x31, 0x12, 0x34, 0x0a, 0x15, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x15, 0x0a, 0x05, 0x4e, 0x75, 0x6c, 0x6c, 0x32, 0x18, 0xb4, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x4e, 0x75, 0x6c, 0x6c, 0x32, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0xdf, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43,
	0x75, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerLoginCsReq_proto_rawDescOnce sync.Once
	file_PlayerLoginCsReq_proto_rawDescData = file_PlayerLoginCsReq_proto_rawDesc
)

func file_PlayerLoginCsReq_proto_rawDescGZIP() []byte {
	file_PlayerLoginCsReq_proto_rawDescOnce.Do(func() {
		file_PlayerLoginCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLoginCsReq_proto_rawDescData)
	})
	return file_PlayerLoginCsReq_proto_rawDescData
}

var file_PlayerLoginCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerLoginCsReq_proto_goTypes = []interface{}{
	(*PlayerLoginCsReq)(nil), // 0: PlayerLoginCsReq
	(LanguageType)(0),        // 1: LanguageType
	(PlatformType)(0),        // 2: PlatformType
}
var file_PlayerLoginCsReq_proto_depIdxs = []int32{
	1, // 0: PlayerLoginCsReq.Language:type_name -> LanguageType
	2, // 1: PlayerLoginCsReq.PlatformType:type_name -> PlatformType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerLoginCsReq_proto_init() }
func file_PlayerLoginCsReq_proto_init() {
	if File_PlayerLoginCsReq_proto != nil {
		return
	}
	file_LanguageType_proto_init()
	file_PlatformType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerLoginCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginCsReq); i {
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
			RawDescriptor: file_PlayerLoginCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLoginCsReq_proto_goTypes,
		DependencyIndexes: file_PlayerLoginCsReq_proto_depIdxs,
		MessageInfos:      file_PlayerLoginCsReq_proto_msgTypes,
	}.Build()
	File_PlayerLoginCsReq_proto = out.File
	file_PlayerLoginCsReq_proto_rawDesc = nil
	file_PlayerLoginCsReq_proto_goTypes = nil
	file_PlayerLoginCsReq_proto_depIdxs = nil
}
