// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FantasticActivityData.proto

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

type FantasticActivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DMNLGAFIJBL []uint32                `protobuf:"varint,6,rep,packed,name=DMNLGAFIJBL,proto3" json:"DMNLGAFIJBL,omitempty"`
	AIHLHFCCGAM []uint32                `protobuf:"varint,3,rep,packed,name=AIHLHFCCGAM,proto3" json:"AIHLHFCCGAM,omitempty"`
	NOGNIMDOGJD map[uint32]*HLDGLCIGODH `protobuf:"bytes,12,rep,name=NOGNIMDOGJD,proto3" json:"NOGNIMDOGJD,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AFIMOPINBMJ map[uint32]uint32       `protobuf:"bytes,1,rep,name=AFIMOPINBMJ,proto3" json:"AFIMOPINBMJ,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DBNMCDGLFKJ uint32                  `protobuf:"varint,14,opt,name=DBNMCDGLFKJ,proto3" json:"DBNMCDGLFKJ,omitempty"`
	DMEJBNMACHI []uint32                `protobuf:"varint,2,rep,packed,name=DMEJBNMACHI,proto3" json:"DMEJBNMACHI,omitempty"`
	GKDKMCGLIEB []uint32                `protobuf:"varint,8,rep,packed,name=GKDKMCGLIEB,proto3" json:"GKDKMCGLIEB,omitempty"`
}

func (x *FantasticActivityData) Reset() {
	*x = FantasticActivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FantasticActivityData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FantasticActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FantasticActivityData) ProtoMessage() {}

func (x *FantasticActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_FantasticActivityData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FantasticActivityData.ProtoReflect.Descriptor instead.
func (*FantasticActivityData) Descriptor() ([]byte, []int) {
	return file_FantasticActivityData_proto_rawDescGZIP(), []int{0}
}

func (x *FantasticActivityData) GetDMNLGAFIJBL() []uint32 {
	if x != nil {
		return x.DMNLGAFIJBL
	}
	return nil
}

func (x *FantasticActivityData) GetAIHLHFCCGAM() []uint32 {
	if x != nil {
		return x.AIHLHFCCGAM
	}
	return nil
}

func (x *FantasticActivityData) GetNOGNIMDOGJD() map[uint32]*HLDGLCIGODH {
	if x != nil {
		return x.NOGNIMDOGJD
	}
	return nil
}

func (x *FantasticActivityData) GetAFIMOPINBMJ() map[uint32]uint32 {
	if x != nil {
		return x.AFIMOPINBMJ
	}
	return nil
}

func (x *FantasticActivityData) GetDBNMCDGLFKJ() uint32 {
	if x != nil {
		return x.DBNMCDGLFKJ
	}
	return 0
}

func (x *FantasticActivityData) GetDMEJBNMACHI() []uint32 {
	if x != nil {
		return x.DMEJBNMACHI
	}
	return nil
}

func (x *FantasticActivityData) GetGKDKMCGLIEB() []uint32 {
	if x != nil {
		return x.GKDKMCGLIEB
	}
	return nil
}

var File_FantasticActivityData_proto protoreflect.FileDescriptor

var file_FantasticActivityData_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48,
	0x4c, 0x44, 0x47, 0x4c, 0x43, 0x49, 0x47, 0x4f, 0x44, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe5, 0x03, 0x0a, 0x15, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4d,
	0x4e, 0x4c, 0x47, 0x41, 0x46, 0x49, 0x4a, 0x42, 0x4c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0b, 0x44, 0x4d, 0x4e, 0x4c, 0x47, 0x41, 0x46, 0x49, 0x4a, 0x42, 0x4c, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x49, 0x48, 0x4c, 0x48, 0x46, 0x43, 0x43, 0x47, 0x41, 0x4d, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0b, 0x41, 0x49, 0x48, 0x4c, 0x48, 0x46, 0x43, 0x43, 0x47, 0x41, 0x4d, 0x12, 0x49,
	0x0a, 0x0b, 0x4e, 0x4f, 0x47, 0x4e, 0x49, 0x4d, 0x44, 0x4f, 0x47, 0x4a, 0x44, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x4f, 0x47, 0x4e,
	0x49, 0x4d, 0x44, 0x4f, 0x47, 0x4a, 0x44, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4e, 0x4f,
	0x47, 0x4e, 0x49, 0x4d, 0x44, 0x4f, 0x47, 0x4a, 0x44, 0x12, 0x49, 0x0a, 0x0b, 0x41, 0x46, 0x49,
	0x4d, 0x4f, 0x50, 0x49, 0x4e, 0x42, 0x4d, 0x4a, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x46, 0x49, 0x4d, 0x4f, 0x50, 0x49, 0x4e, 0x42,
	0x4d, 0x4a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x41, 0x46, 0x49, 0x4d, 0x4f, 0x50, 0x49,
	0x4e, 0x42, 0x4d, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x42, 0x4e, 0x4d, 0x43, 0x44, 0x47, 0x4c,
	0x46, 0x4b, 0x4a, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x42, 0x4e, 0x4d, 0x43,
	0x44, 0x47, 0x4c, 0x46, 0x4b, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4d, 0x45, 0x4a, 0x42, 0x4e,
	0x4d, 0x41, 0x43, 0x48, 0x49, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4d, 0x45,
	0x4a, 0x42, 0x4e, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4b, 0x44, 0x4b,
	0x4d, 0x43, 0x47, 0x4c, 0x49, 0x45, 0x42, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x47,
	0x4b, 0x44, 0x4b, 0x4d, 0x43, 0x47, 0x4c, 0x49, 0x45, 0x42, 0x1a, 0x4c, 0x0a, 0x10, 0x4e, 0x4f,
	0x47, 0x4e, 0x49, 0x4d, 0x44, 0x4f, 0x47, 0x4a, 0x44, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x48, 0x4c, 0x44, 0x47, 0x4c, 0x43, 0x49, 0x47, 0x4f, 0x44, 0x48, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x46, 0x49, 0x4d,
	0x4f, 0x50, 0x49, 0x4e, 0x42, 0x4d, 0x4a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FantasticActivityData_proto_rawDescOnce sync.Once
	file_FantasticActivityData_proto_rawDescData = file_FantasticActivityData_proto_rawDesc
)

func file_FantasticActivityData_proto_rawDescGZIP() []byte {
	file_FantasticActivityData_proto_rawDescOnce.Do(func() {
		file_FantasticActivityData_proto_rawDescData = protoimpl.X.CompressGZIP(file_FantasticActivityData_proto_rawDescData)
	})
	return file_FantasticActivityData_proto_rawDescData
}

var file_FantasticActivityData_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_FantasticActivityData_proto_goTypes = []interface{}{
	(*FantasticActivityData)(nil), // 0: FantasticActivityData
	nil,                           // 1: FantasticActivityData.NOGNIMDOGJDEntry
	nil,                           // 2: FantasticActivityData.AFIMOPINBMJEntry
	(*HLDGLCIGODH)(nil),           // 3: HLDGLCIGODH
}
var file_FantasticActivityData_proto_depIdxs = []int32{
	1, // 0: FantasticActivityData.NOGNIMDOGJD:type_name -> FantasticActivityData.NOGNIMDOGJDEntry
	2, // 1: FantasticActivityData.AFIMOPINBMJ:type_name -> FantasticActivityData.AFIMOPINBMJEntry
	3, // 2: FantasticActivityData.NOGNIMDOGJDEntry.value:type_name -> HLDGLCIGODH
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_FantasticActivityData_proto_init() }
func file_FantasticActivityData_proto_init() {
	if File_FantasticActivityData_proto != nil {
		return
	}
	file_HLDGLCIGODH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FantasticActivityData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FantasticActivityData); i {
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
			RawDescriptor: file_FantasticActivityData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FantasticActivityData_proto_goTypes,
		DependencyIndexes: file_FantasticActivityData_proto_depIdxs,
		MessageInfos:      file_FantasticActivityData_proto_msgTypes,
	}.Build()
	File_FantasticActivityData_proto = out.File
	file_FantasticActivityData_proto_rawDesc = nil
	file_FantasticActivityData_proto_goTypes = nil
	file_FantasticActivityData_proto_depIdxs = nil
}
