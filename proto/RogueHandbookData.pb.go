// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueHandbookData.proto

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

type RogueHandbookData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueEvent    []*RogueHandbookEvent   `protobuf:"bytes,12,rep,name=rogue_event,json=rogueEvent,proto3" json:"rogue_event,omitempty"`
	BuffList      []*RogueHandbookBuff    `protobuf:"bytes,15,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	RogueAeonList []*RogueHandbookAeon    `protobuf:"bytes,9,rep,name=rogue_aeon_list,json=rogueAeonList,proto3" json:"rogue_aeon_list,omitempty"`
	MiracleList   []*RogueHandbookMiracle `protobuf:"bytes,3,rep,name=miracle_list,json=miracleList,proto3" json:"miracle_list,omitempty"`
}

func (x *RogueHandbookData) Reset() {
	*x = RogueHandbookData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueHandbookData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueHandbookData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueHandbookData) ProtoMessage() {}

func (x *RogueHandbookData) ProtoReflect() protoreflect.Message {
	mi := &file_RogueHandbookData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueHandbookData.ProtoReflect.Descriptor instead.
func (*RogueHandbookData) Descriptor() ([]byte, []int) {
	return file_RogueHandbookData_proto_rawDescGZIP(), []int{0}
}

func (x *RogueHandbookData) GetRogueEvent() []*RogueHandbookEvent {
	if x != nil {
		return x.RogueEvent
	}
	return nil
}

func (x *RogueHandbookData) GetBuffList() []*RogueHandbookBuff {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *RogueHandbookData) GetRogueAeonList() []*RogueHandbookAeon {
	if x != nil {
		return x.RogueAeonList
	}
	return nil
}

func (x *RogueHandbookData) GetMiracleList() []*RogueHandbookMiracle {
	if x != nil {
		return x.MiracleList
	}
	return nil
}

var File_RogueHandbookData_proto protoreflect.FileDescriptor

var file_RogueHandbookData_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x6f, 0x6b, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x75,
	0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x41, 0x65, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x0b, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a,
	0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b,
	0x42, 0x75, 0x66, 0x66, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x41, 0x65, 0x6f, 0x6e, 0x52, 0x0d, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x6d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b,
	0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueHandbookData_proto_rawDescOnce sync.Once
	file_RogueHandbookData_proto_rawDescData = file_RogueHandbookData_proto_rawDesc
)

func file_RogueHandbookData_proto_rawDescGZIP() []byte {
	file_RogueHandbookData_proto_rawDescOnce.Do(func() {
		file_RogueHandbookData_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueHandbookData_proto_rawDescData)
	})
	return file_RogueHandbookData_proto_rawDescData
}

var file_RogueHandbookData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueHandbookData_proto_goTypes = []interface{}{
	(*RogueHandbookData)(nil),    // 0: RogueHandbookData
	(*RogueHandbookEvent)(nil),   // 1: RogueHandbookEvent
	(*RogueHandbookBuff)(nil),    // 2: RogueHandbookBuff
	(*RogueHandbookAeon)(nil),    // 3: RogueHandbookAeon
	(*RogueHandbookMiracle)(nil), // 4: RogueHandbookMiracle
}
var file_RogueHandbookData_proto_depIdxs = []int32{
	1, // 0: RogueHandbookData.rogue_event:type_name -> RogueHandbookEvent
	2, // 1: RogueHandbookData.buff_list:type_name -> RogueHandbookBuff
	3, // 2: RogueHandbookData.rogue_aeon_list:type_name -> RogueHandbookAeon
	4, // 3: RogueHandbookData.miracle_list:type_name -> RogueHandbookMiracle
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_RogueHandbookData_proto_init() }
func file_RogueHandbookData_proto_init() {
	if File_RogueHandbookData_proto != nil {
		return
	}
	file_RogueHandbookEvent_proto_init()
	file_RogueHandbookMiracle_proto_init()
	file_RogueHandbookBuff_proto_init()
	file_RogueHandbookAeon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueHandbookData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueHandbookData); i {
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
			RawDescriptor: file_RogueHandbookData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueHandbookData_proto_goTypes,
		DependencyIndexes: file_RogueHandbookData_proto_depIdxs,
		MessageInfos:      file_RogueHandbookData_proto_msgTypes,
	}.Build()
	File_RogueHandbookData_proto = out.File
	file_RogueHandbookData_proto_rawDesc = nil
	file_RogueHandbookData_proto_goTypes = nil
	file_RogueHandbookData_proto_depIdxs = nil
}
