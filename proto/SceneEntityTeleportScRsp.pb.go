// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneEntityTeleportScRsp.proto

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

type SceneEntityTeleportScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode          uint32        `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ClientPosVersion uint32        `protobuf:"varint,1,opt,name=client_pos_version,json=clientPosVersion,proto3" json:"client_pos_version,omitempty"`
	EntityMotion     *EntityMotion `protobuf:"bytes,7,opt,name=entity_motion,json=entityMotion,proto3" json:"entity_motion,omitempty"`
}

func (x *SceneEntityTeleportScRsp) Reset() {
	*x = SceneEntityTeleportScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityTeleportScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityTeleportScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityTeleportScRsp) ProtoMessage() {}

func (x *SceneEntityTeleportScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityTeleportScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityTeleportScRsp.ProtoReflect.Descriptor instead.
func (*SceneEntityTeleportScRsp) Descriptor() ([]byte, []int) {
	return file_SceneEntityTeleportScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityTeleportScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SceneEntityTeleportScRsp) GetClientPosVersion() uint32 {
	if x != nil {
		return x.ClientPosVersion
	}
	return 0
}

func (x *SceneEntityTeleportScRsp) GetEntityMotion() *EntityMotion {
	if x != nil {
		return x.EntityMotion
	}
	return nil
}

var File_SceneEntityTeleportScRsp_proto protoreflect.FileDescriptor

var file_SceneEntityTeleportScRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x18, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0d, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_SceneEntityTeleportScRsp_proto_rawDescOnce sync.Once
	file_SceneEntityTeleportScRsp_proto_rawDescData = file_SceneEntityTeleportScRsp_proto_rawDesc
)

func file_SceneEntityTeleportScRsp_proto_rawDescGZIP() []byte {
	file_SceneEntityTeleportScRsp_proto_rawDescOnce.Do(func() {
		file_SceneEntityTeleportScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityTeleportScRsp_proto_rawDescData)
	})
	return file_SceneEntityTeleportScRsp_proto_rawDescData
}

var file_SceneEntityTeleportScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityTeleportScRsp_proto_goTypes = []interface{}{
	(*SceneEntityTeleportScRsp)(nil), // 0: SceneEntityTeleportScRsp
	(*EntityMotion)(nil),             // 1: EntityMotion
}
var file_SceneEntityTeleportScRsp_proto_depIdxs = []int32{
	1, // 0: SceneEntityTeleportScRsp.entity_motion:type_name -> EntityMotion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneEntityTeleportScRsp_proto_init() }
func file_SceneEntityTeleportScRsp_proto_init() {
	if File_SceneEntityTeleportScRsp_proto != nil {
		return
	}
	file_EntityMotion_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityTeleportScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityTeleportScRsp); i {
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
			RawDescriptor: file_SceneEntityTeleportScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityTeleportScRsp_proto_goTypes,
		DependencyIndexes: file_SceneEntityTeleportScRsp_proto_depIdxs,
		MessageInfos:      file_SceneEntityTeleportScRsp_proto_msgTypes,
	}.Build()
	File_SceneEntityTeleportScRsp_proto = out.File
	file_SceneEntityTeleportScRsp_proto_rawDesc = nil
	file_SceneEntityTeleportScRsp_proto_goTypes = nil
	file_SceneEntityTeleportScRsp_proto_depIdxs = nil
}
