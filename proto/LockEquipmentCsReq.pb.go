// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LockEquipmentCsReq.proto

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

type LockEquipmentCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsProtected     bool     `protobuf:"varint,9,opt,name=is_protected,json=isProtected,proto3" json:"is_protected,omitempty"`
	EquipmentIdList []uint32 `protobuf:"varint,13,rep,packed,name=equipment_id_list,json=equipmentIdList,proto3" json:"equipment_id_list,omitempty"`
}

func (x *LockEquipmentCsReq) Reset() {
	*x = LockEquipmentCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockEquipmentCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockEquipmentCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockEquipmentCsReq) ProtoMessage() {}

func (x *LockEquipmentCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_LockEquipmentCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockEquipmentCsReq.ProtoReflect.Descriptor instead.
func (*LockEquipmentCsReq) Descriptor() ([]byte, []int) {
	return file_LockEquipmentCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *LockEquipmentCsReq) GetIsProtected() bool {
	if x != nil {
		return x.IsProtected
	}
	return false
}

func (x *LockEquipmentCsReq) GetEquipmentIdList() []uint32 {
	if x != nil {
		return x.EquipmentIdList
	}
	return nil
}

var File_LockEquipmentCsReq_proto protoreflect.FileDescriptor

var file_LockEquipmentCsReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4c, 0x6f, 0x63, 0x6b, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x12, 0x4c, 0x6f,
	0x63, 0x6b, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_LockEquipmentCsReq_proto_rawDescOnce sync.Once
	file_LockEquipmentCsReq_proto_rawDescData = file_LockEquipmentCsReq_proto_rawDesc
)

func file_LockEquipmentCsReq_proto_rawDescGZIP() []byte {
	file_LockEquipmentCsReq_proto_rawDescOnce.Do(func() {
		file_LockEquipmentCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_LockEquipmentCsReq_proto_rawDescData)
	})
	return file_LockEquipmentCsReq_proto_rawDescData
}

var file_LockEquipmentCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LockEquipmentCsReq_proto_goTypes = []interface{}{
	(*LockEquipmentCsReq)(nil), // 0: LockEquipmentCsReq
}
var file_LockEquipmentCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LockEquipmentCsReq_proto_init() }
func file_LockEquipmentCsReq_proto_init() {
	if File_LockEquipmentCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LockEquipmentCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockEquipmentCsReq); i {
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
			RawDescriptor: file_LockEquipmentCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LockEquipmentCsReq_proto_goTypes,
		DependencyIndexes: file_LockEquipmentCsReq_proto_depIdxs,
		MessageInfos:      file_LockEquipmentCsReq_proto_msgTypes,
	}.Build()
	File_LockEquipmentCsReq_proto = out.File
	file_LockEquipmentCsReq_proto_rawDesc = nil
	file_LockEquipmentCsReq_proto_goTypes = nil
	file_LockEquipmentCsReq_proto_depIdxs = nil
}