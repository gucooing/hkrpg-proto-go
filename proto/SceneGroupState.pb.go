// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneGroupState.proto

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

type SceneGroupState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDefault bool   `protobuf:"varint,10,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	State     uint32 `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
	GroupId   uint32 `protobuf:"varint,11,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *SceneGroupState) Reset() {
	*x = SceneGroupState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGroupState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGroupState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGroupState) ProtoMessage() {}

func (x *SceneGroupState) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGroupState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGroupState.ProtoReflect.Descriptor instead.
func (*SceneGroupState) Descriptor() ([]byte, []int) {
	return file_SceneGroupState_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGroupState) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *SceneGroupState) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SceneGroupState) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_SceneGroupState_proto protoreflect.FileDescriptor

var file_SceneGroupState_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGroupState_proto_rawDescOnce sync.Once
	file_SceneGroupState_proto_rawDescData = file_SceneGroupState_proto_rawDesc
)

func file_SceneGroupState_proto_rawDescGZIP() []byte {
	file_SceneGroupState_proto_rawDescOnce.Do(func() {
		file_SceneGroupState_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGroupState_proto_rawDescData)
	})
	return file_SceneGroupState_proto_rawDescData
}

var file_SceneGroupState_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGroupState_proto_goTypes = []interface{}{
	(*SceneGroupState)(nil), // 0: SceneGroupState
}
var file_SceneGroupState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneGroupState_proto_init() }
func file_SceneGroupState_proto_init() {
	if File_SceneGroupState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneGroupState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGroupState); i {
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
			RawDescriptor: file_SceneGroupState_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGroupState_proto_goTypes,
		DependencyIndexes: file_SceneGroupState_proto_depIdxs,
		MessageInfos:      file_SceneGroupState_proto_msgTypes,
	}.Build()
	File_SceneGroupState_proto = out.File
	file_SceneGroupState_proto_rawDesc = nil
	file_SceneGroupState_proto_goTypes = nil
	file_SceneGroupState_proto_depIdxs = nil
}
