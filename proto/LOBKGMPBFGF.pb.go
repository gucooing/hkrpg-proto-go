// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LOBKGMPBFGF.proto

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

type LOBKGMPBFGF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId     uint32         `protobuf:"varint,6,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	JPLKINPPOLD []*ILJGLGIKBOD `protobuf:"bytes,10,rep,name=JPLKINPPOLD,proto3" json:"JPLKINPPOLD,omitempty"`
}

func (x *LOBKGMPBFGF) Reset() {
	*x = LOBKGMPBFGF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LOBKGMPBFGF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LOBKGMPBFGF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LOBKGMPBFGF) ProtoMessage() {}

func (x *LOBKGMPBFGF) ProtoReflect() protoreflect.Message {
	mi := &file_LOBKGMPBFGF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LOBKGMPBFGF.ProtoReflect.Descriptor instead.
func (*LOBKGMPBFGF) Descriptor() ([]byte, []int) {
	return file_LOBKGMPBFGF_proto_rawDescGZIP(), []int{0}
}

func (x *LOBKGMPBFGF) GetGoodsId() uint32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *LOBKGMPBFGF) GetJPLKINPPOLD() []*ILJGLGIKBOD {
	if x != nil {
		return x.JPLKINPPOLD
	}
	return nil
}

var File_LOBKGMPBFGF_proto protoreflect.FileDescriptor

var file_LOBKGMPBFGF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x4f, 0x42, 0x4b, 0x47, 0x4d, 0x50, 0x42, 0x46, 0x47, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4c, 0x4a, 0x47, 0x4c, 0x47, 0x49, 0x4b, 0x42, 0x4f, 0x44,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0b, 0x4c, 0x4f, 0x42, 0x4b, 0x47, 0x4d,
	0x50, 0x42, 0x46, 0x47, 0x46, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x50, 0x4c, 0x4b, 0x49, 0x4e, 0x50, 0x50, 0x4f, 0x4c, 0x44, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x4c, 0x4a, 0x47, 0x4c, 0x47, 0x49, 0x4b,
	0x42, 0x4f, 0x44, 0x52, 0x0b, 0x4a, 0x50, 0x4c, 0x4b, 0x49, 0x4e, 0x50, 0x50, 0x4f, 0x4c, 0x44,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_LOBKGMPBFGF_proto_rawDescOnce sync.Once
	file_LOBKGMPBFGF_proto_rawDescData = file_LOBKGMPBFGF_proto_rawDesc
)

func file_LOBKGMPBFGF_proto_rawDescGZIP() []byte {
	file_LOBKGMPBFGF_proto_rawDescOnce.Do(func() {
		file_LOBKGMPBFGF_proto_rawDescData = protoimpl.X.CompressGZIP(file_LOBKGMPBFGF_proto_rawDescData)
	})
	return file_LOBKGMPBFGF_proto_rawDescData
}

var file_LOBKGMPBFGF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LOBKGMPBFGF_proto_goTypes = []interface{}{
	(*LOBKGMPBFGF)(nil), // 0: LOBKGMPBFGF
	(*ILJGLGIKBOD)(nil), // 1: ILJGLGIKBOD
}
var file_LOBKGMPBFGF_proto_depIdxs = []int32{
	1, // 0: LOBKGMPBFGF.JPLKINPPOLD:type_name -> ILJGLGIKBOD
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LOBKGMPBFGF_proto_init() }
func file_LOBKGMPBFGF_proto_init() {
	if File_LOBKGMPBFGF_proto != nil {
		return
	}
	file_ILJGLGIKBOD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LOBKGMPBFGF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LOBKGMPBFGF); i {
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
			RawDescriptor: file_LOBKGMPBFGF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LOBKGMPBFGF_proto_goTypes,
		DependencyIndexes: file_LOBKGMPBFGF_proto_depIdxs,
		MessageInfos:      file_LOBKGMPBFGF_proto_msgTypes,
	}.Build()
	File_LOBKGMPBFGF_proto = out.File
	file_LOBKGMPBFGF_proto_rawDesc = nil
	file_LOBKGMPBFGF_proto_goTypes = nil
	file_LOBKGMPBFGF_proto_depIdxs = nil
}
