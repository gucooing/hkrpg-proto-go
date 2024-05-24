// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SyncRogueDialogueEventDataScNotify.proto

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

type SyncRogueDialogueEventDataScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueDialogueEvent []*RogueDialogueEvent `protobuf:"bytes,13,rep,name=rogue_dialogue_event,json=rogueDialogueEvent,proto3" json:"rogue_dialogue_event,omitempty"`
}

func (x *SyncRogueDialogueEventDataScNotify) Reset() {
	*x = SyncRogueDialogueEventDataScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncRogueDialogueEventDataScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRogueDialogueEventDataScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRogueDialogueEventDataScNotify) ProtoMessage() {}

func (x *SyncRogueDialogueEventDataScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncRogueDialogueEventDataScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRogueDialogueEventDataScNotify.ProtoReflect.Descriptor instead.
func (*SyncRogueDialogueEventDataScNotify) Descriptor() ([]byte, []int) {
	return file_SyncRogueDialogueEventDataScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRogueDialogueEventDataScNotify) GetRogueDialogueEvent() []*RogueDialogueEvent {
	if x != nil {
		return x.RogueDialogueEvent
	}
	return nil
}

var File_SyncRogueDialogueEventDataScNotify_proto protoreflect.FileDescriptor

var file_SyncRogueDialogueEventDataScNotify_proto_rawDesc = []byte{
	0x0a, 0x28, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x22, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x45, 0x0a, 0x14, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x12, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncRogueDialogueEventDataScNotify_proto_rawDescOnce sync.Once
	file_SyncRogueDialogueEventDataScNotify_proto_rawDescData = file_SyncRogueDialogueEventDataScNotify_proto_rawDesc
)

func file_SyncRogueDialogueEventDataScNotify_proto_rawDescGZIP() []byte {
	file_SyncRogueDialogueEventDataScNotify_proto_rawDescOnce.Do(func() {
		file_SyncRogueDialogueEventDataScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncRogueDialogueEventDataScNotify_proto_rawDescData)
	})
	return file_SyncRogueDialogueEventDataScNotify_proto_rawDescData
}

var file_SyncRogueDialogueEventDataScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncRogueDialogueEventDataScNotify_proto_goTypes = []interface{}{
	(*SyncRogueDialogueEventDataScNotify)(nil), // 0: SyncRogueDialogueEventDataScNotify
	(*RogueDialogueEvent)(nil),                 // 1: RogueDialogueEvent
}
var file_SyncRogueDialogueEventDataScNotify_proto_depIdxs = []int32{
	1, // 0: SyncRogueDialogueEventDataScNotify.rogue_dialogue_event:type_name -> RogueDialogueEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SyncRogueDialogueEventDataScNotify_proto_init() }
func file_SyncRogueDialogueEventDataScNotify_proto_init() {
	if File_SyncRogueDialogueEventDataScNotify_proto != nil {
		return
	}
	file_RogueDialogueEvent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncRogueDialogueEventDataScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRogueDialogueEventDataScNotify); i {
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
			RawDescriptor: file_SyncRogueDialogueEventDataScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncRogueDialogueEventDataScNotify_proto_goTypes,
		DependencyIndexes: file_SyncRogueDialogueEventDataScNotify_proto_depIdxs,
		MessageInfos:      file_SyncRogueDialogueEventDataScNotify_proto_msgTypes,
	}.Build()
	File_SyncRogueDialogueEventDataScNotify_proto = out.File
	file_SyncRogueDialogueEventDataScNotify_proto_rawDesc = nil
	file_SyncRogueDialogueEventDataScNotify_proto_goTypes = nil
	file_SyncRogueDialogueEventDataScNotify_proto_depIdxs = nil
}
