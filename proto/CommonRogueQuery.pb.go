// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CommonRogueQuery.proto

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

type CommonRogueQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueUpdate *RogueUpdate `protobuf:"bytes,6,opt,name=rogue_update,json=rogueUpdate,proto3" json:"rogue_update,omitempty"`
	RogueQuery  *RogueQuery  `protobuf:"bytes,2,opt,name=rogue_query,json=rogueQuery,proto3" json:"rogue_query,omitempty"`
}

func (x *CommonRogueQuery) Reset() {
	*x = CommonRogueQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonRogueQuery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRogueQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRogueQuery) ProtoMessage() {}

func (x *CommonRogueQuery) ProtoReflect() protoreflect.Message {
	mi := &file_CommonRogueQuery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRogueQuery.ProtoReflect.Descriptor instead.
func (*CommonRogueQuery) Descriptor() ([]byte, []int) {
	return file_CommonRogueQuery_proto_rawDescGZIP(), []int{0}
}

func (x *CommonRogueQuery) GetRogueUpdate() *RogueUpdate {
	if x != nil {
		return x.RogueUpdate
	}
	return nil
}

func (x *CommonRogueQuery) GetRogueQuery() *RogueQuery {
	if x != nil {
		return x.RogueQuery
	}
	return nil
}

var File_CommonRogueQuery_proto protoreflect.FileDescriptor

var file_CommonRogueQuery_proto_rawDesc = []byte{
	0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x2f, 0x0a, 0x0c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonRogueQuery_proto_rawDescOnce sync.Once
	file_CommonRogueQuery_proto_rawDescData = file_CommonRogueQuery_proto_rawDesc
)

func file_CommonRogueQuery_proto_rawDescGZIP() []byte {
	file_CommonRogueQuery_proto_rawDescOnce.Do(func() {
		file_CommonRogueQuery_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonRogueQuery_proto_rawDescData)
	})
	return file_CommonRogueQuery_proto_rawDescData
}

var file_CommonRogueQuery_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonRogueQuery_proto_goTypes = []interface{}{
	(*CommonRogueQuery)(nil), // 0: CommonRogueQuery
	(*RogueUpdate)(nil),      // 1: RogueUpdate
	(*RogueQuery)(nil),       // 2: RogueQuery
}
var file_CommonRogueQuery_proto_depIdxs = []int32{
	1, // 0: CommonRogueQuery.rogue_update:type_name -> RogueUpdate
	2, // 1: CommonRogueQuery.rogue_query:type_name -> RogueQuery
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_CommonRogueQuery_proto_init() }
func file_CommonRogueQuery_proto_init() {
	if File_CommonRogueQuery_proto != nil {
		return
	}
	file_RogueUpdate_proto_init()
	file_RogueQuery_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonRogueQuery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRogueQuery); i {
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
			RawDescriptor: file_CommonRogueQuery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonRogueQuery_proto_goTypes,
		DependencyIndexes: file_CommonRogueQuery_proto_depIdxs,
		MessageInfos:      file_CommonRogueQuery_proto_msgTypes,
	}.Build()
	File_CommonRogueQuery_proto = out.File
	file_CommonRogueQuery_proto_rawDesc = nil
	file_CommonRogueQuery_proto_goTypes = nil
	file_CommonRogueQuery_proto_depIdxs = nil
}
