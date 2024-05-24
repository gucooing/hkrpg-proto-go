// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AvatarProperty.proto

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

type AvatarProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxHp   float64 `protobuf:"fixed64,1,opt,name=max_hp,json=maxHp,proto3" json:"max_hp,omitempty"`
	Attack  float64 `protobuf:"fixed64,2,opt,name=attack,proto3" json:"attack,omitempty"`
	Defence float64 `protobuf:"fixed64,3,opt,name=defence,proto3" json:"defence,omitempty"`
	Speed   float64 `protobuf:"fixed64,4,opt,name=speed,proto3" json:"speed,omitempty"`
	LeftHp  float64 `protobuf:"fixed64,5,opt,name=left_hp,json=leftHp,proto3" json:"left_hp,omitempty"`
	LeftSp  float64 `protobuf:"fixed64,6,opt,name=left_sp,json=leftSp,proto3" json:"left_sp,omitempty"`
	MaxSp   float64 `protobuf:"fixed64,7,opt,name=max_sp,json=maxSp,proto3" json:"max_sp,omitempty"`
}

func (x *AvatarProperty) Reset() {
	*x = AvatarProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarProperty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarProperty) ProtoMessage() {}

func (x *AvatarProperty) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarProperty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarProperty.ProtoReflect.Descriptor instead.
func (*AvatarProperty) Descriptor() ([]byte, []int) {
	return file_AvatarProperty_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarProperty) GetMaxHp() float64 {
	if x != nil {
		return x.MaxHp
	}
	return 0
}

func (x *AvatarProperty) GetAttack() float64 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *AvatarProperty) GetDefence() float64 {
	if x != nil {
		return x.Defence
	}
	return 0
}

func (x *AvatarProperty) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *AvatarProperty) GetLeftHp() float64 {
	if x != nil {
		return x.LeftHp
	}
	return 0
}

func (x *AvatarProperty) GetLeftSp() float64 {
	if x != nil {
		return x.LeftSp
	}
	return 0
}

func (x *AvatarProperty) GetMaxSp() float64 {
	if x != nil {
		return x.MaxSp
	}
	return 0
}

var File_AvatarProperty_proto protoreflect.FileDescriptor

var file_AvatarProperty_proto_rawDesc = []byte{
	0x0a, 0x14, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78,
	0x5f, 0x68, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x48, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x64, 0x65, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x66, 0x74,
	0x5f, 0x68, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x65, 0x66, 0x74, 0x48,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x73, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x6c, 0x65, 0x66, 0x74, 0x53, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61,
	0x78, 0x5f, 0x73, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x53,
	0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarProperty_proto_rawDescOnce sync.Once
	file_AvatarProperty_proto_rawDescData = file_AvatarProperty_proto_rawDesc
)

func file_AvatarProperty_proto_rawDescGZIP() []byte {
	file_AvatarProperty_proto_rawDescOnce.Do(func() {
		file_AvatarProperty_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarProperty_proto_rawDescData)
	})
	return file_AvatarProperty_proto_rawDescData
}

var file_AvatarProperty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarProperty_proto_goTypes = []interface{}{
	(*AvatarProperty)(nil), // 0: AvatarProperty
}
var file_AvatarProperty_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarProperty_proto_init() }
func file_AvatarProperty_proto_init() {
	if File_AvatarProperty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarProperty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarProperty); i {
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
			RawDescriptor: file_AvatarProperty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarProperty_proto_goTypes,
		DependencyIndexes: file_AvatarProperty_proto_depIdxs,
		MessageInfos:      file_AvatarProperty_proto_msgTypes,
	}.Build()
	File_AvatarProperty_proto = out.File
	file_AvatarProperty_proto_rawDesc = nil
	file_AvatarProperty_proto_goTypes = nil
	file_AvatarProperty_proto_depIdxs = nil
}