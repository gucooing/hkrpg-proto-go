// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HBPLKCPLECE.proto

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

type HBPLKCPLECE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         HDEDBPNEFID `protobuf:"varint,1,opt,name=type,proto3,enum=HDEDBPNEFID" json:"type,omitempty"`
	MIDKHPPKGCJ  uint32      `protobuf:"varint,2,opt,name=MIDKHPPKGCJ,proto3" json:"MIDKHPPKGCJ,omitempty"`
	DisplayValue uint32      `protobuf:"varint,3,opt,name=display_value,json=displayValue,proto3" json:"display_value,omitempty"`
}

func (x *HBPLKCPLECE) Reset() {
	*x = HBPLKCPLECE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HBPLKCPLECE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HBPLKCPLECE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HBPLKCPLECE) ProtoMessage() {}

func (x *HBPLKCPLECE) ProtoReflect() protoreflect.Message {
	mi := &file_HBPLKCPLECE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HBPLKCPLECE.ProtoReflect.Descriptor instead.
func (*HBPLKCPLECE) Descriptor() ([]byte, []int) {
	return file_HBPLKCPLECE_proto_rawDescGZIP(), []int{0}
}

func (x *HBPLKCPLECE) GetType() HDEDBPNEFID {
	if x != nil {
		return x.Type
	}
	return HDEDBPNEFID_BATTLE_STATICTIC_EVENT_NONE
}

func (x *HBPLKCPLECE) GetMIDKHPPKGCJ() uint32 {
	if x != nil {
		return x.MIDKHPPKGCJ
	}
	return 0
}

func (x *HBPLKCPLECE) GetDisplayValue() uint32 {
	if x != nil {
		return x.DisplayValue
	}
	return 0
}

var File_HBPLKCPLECE_proto protoreflect.FileDescriptor

var file_HBPLKCPLECE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x42, 0x50, 0x4c, 0x4b, 0x43, 0x50, 0x4c, 0x45, 0x43, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x44, 0x45, 0x44, 0x42, 0x50, 0x4e, 0x45, 0x46, 0x49, 0x44,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x0b, 0x48, 0x42, 0x50, 0x4c, 0x4b, 0x43,
	0x50, 0x4c, 0x45, 0x43, 0x45, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x48, 0x44, 0x45, 0x44, 0x42, 0x50, 0x4e, 0x45, 0x46, 0x49,
	0x44, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x49, 0x44, 0x4b, 0x48,
	0x50, 0x50, 0x4b, 0x47, 0x43, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x49,
	0x44, 0x4b, 0x48, 0x50, 0x50, 0x4b, 0x47, 0x43, 0x4a, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HBPLKCPLECE_proto_rawDescOnce sync.Once
	file_HBPLKCPLECE_proto_rawDescData = file_HBPLKCPLECE_proto_rawDesc
)

func file_HBPLKCPLECE_proto_rawDescGZIP() []byte {
	file_HBPLKCPLECE_proto_rawDescOnce.Do(func() {
		file_HBPLKCPLECE_proto_rawDescData = protoimpl.X.CompressGZIP(file_HBPLKCPLECE_proto_rawDescData)
	})
	return file_HBPLKCPLECE_proto_rawDescData
}

var file_HBPLKCPLECE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HBPLKCPLECE_proto_goTypes = []interface{}{
	(*HBPLKCPLECE)(nil), // 0: HBPLKCPLECE
	(HDEDBPNEFID)(0),    // 1: HDEDBPNEFID
}
var file_HBPLKCPLECE_proto_depIdxs = []int32{
	1, // 0: HBPLKCPLECE.type:type_name -> HDEDBPNEFID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HBPLKCPLECE_proto_init() }
func file_HBPLKCPLECE_proto_init() {
	if File_HBPLKCPLECE_proto != nil {
		return
	}
	file_HDEDBPNEFID_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HBPLKCPLECE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HBPLKCPLECE); i {
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
			RawDescriptor: file_HBPLKCPLECE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HBPLKCPLECE_proto_goTypes,
		DependencyIndexes: file_HBPLKCPLECE_proto_depIdxs,
		MessageInfos:      file_HBPLKCPLECE_proto_msgTypes,
	}.Build()
	File_HBPLKCPLECE_proto = out.File
	file_HBPLKCPLECE_proto_rawDesc = nil
	file_HBPLKCPLECE_proto_goTypes = nil
	file_HBPLKCPLECE_proto_depIdxs = nil
}
