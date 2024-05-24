// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerSettingInfo.proto

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

type PlayerSettingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B1                bool              `protobuf:"varint,2,opt,name=b1,proto3" json:"b1,omitempty"`
	B2                bool              `protobuf:"varint,8,opt,name=b2,proto3" json:"b2,omitempty"`
	B3                bool              `protobuf:"varint,4,opt,name=b3,proto3" json:"b3,omitempty"`
	B4                bool              `protobuf:"varint,3,opt,name=b4,proto3" json:"b4,omitempty"`
	B5                bool              `protobuf:"varint,13,opt,name=b5,proto3" json:"b5,omitempty"`
	B6                bool              `protobuf:"varint,5,opt,name=b6,proto3" json:"b6,omitempty"`
	DisplayRecordType DisplayRecordType `protobuf:"varint,12,opt,name=display_record_type,json=displayRecordType,proto3,enum=DisplayRecordType" json:"display_record_type,omitempty"`
}

func (x *PlayerSettingInfo) Reset() {
	*x = PlayerSettingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerSettingInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerSettingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerSettingInfo) ProtoMessage() {}

func (x *PlayerSettingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerSettingInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerSettingInfo.ProtoReflect.Descriptor instead.
func (*PlayerSettingInfo) Descriptor() ([]byte, []int) {
	return file_PlayerSettingInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerSettingInfo) GetB1() bool {
	if x != nil {
		return x.B1
	}
	return false
}

func (x *PlayerSettingInfo) GetB2() bool {
	if x != nil {
		return x.B2
	}
	return false
}

func (x *PlayerSettingInfo) GetB3() bool {
	if x != nil {
		return x.B3
	}
	return false
}

func (x *PlayerSettingInfo) GetB4() bool {
	if x != nil {
		return x.B4
	}
	return false
}

func (x *PlayerSettingInfo) GetB5() bool {
	if x != nil {
		return x.B5
	}
	return false
}

func (x *PlayerSettingInfo) GetB6() bool {
	if x != nil {
		return x.B6
	}
	return false
}

func (x *PlayerSettingInfo) GetDisplayRecordType() DisplayRecordType {
	if x != nil {
		return x.DisplayRecordType
	}
	return DisplayRecordType_BATTLE_RECORD_NONE
}

var File_PlayerSettingInfo_proto protoreflect.FileDescriptor

var file_PlayerSettingInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x31, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x62, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x32, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x62, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x33, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x62, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x34, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x62, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x35, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x62, 0x35, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x36, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x62, 0x36, 0x12, 0x42, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerSettingInfo_proto_rawDescOnce sync.Once
	file_PlayerSettingInfo_proto_rawDescData = file_PlayerSettingInfo_proto_rawDesc
)

func file_PlayerSettingInfo_proto_rawDescGZIP() []byte {
	file_PlayerSettingInfo_proto_rawDescOnce.Do(func() {
		file_PlayerSettingInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerSettingInfo_proto_rawDescData)
	})
	return file_PlayerSettingInfo_proto_rawDescData
}

var file_PlayerSettingInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerSettingInfo_proto_goTypes = []interface{}{
	(*PlayerSettingInfo)(nil), // 0: PlayerSettingInfo
	(DisplayRecordType)(0),    // 1: DisplayRecordType
}
var file_PlayerSettingInfo_proto_depIdxs = []int32{
	1, // 0: PlayerSettingInfo.display_record_type:type_name -> DisplayRecordType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerSettingInfo_proto_init() }
func file_PlayerSettingInfo_proto_init() {
	if File_PlayerSettingInfo_proto != nil {
		return
	}
	file_DisplayRecordType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerSettingInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerSettingInfo); i {
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
			RawDescriptor: file_PlayerSettingInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerSettingInfo_proto_goTypes,
		DependencyIndexes: file_PlayerSettingInfo_proto_depIdxs,
		MessageInfos:      file_PlayerSettingInfo_proto_msgTypes,
	}.Build()
	File_PlayerSettingInfo_proto = out.File
	file_PlayerSettingInfo_proto_rawDesc = nil
	file_PlayerSettingInfo_proto_goTypes = nil
	file_PlayerSettingInfo_proto_depIdxs = nil
}