// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetPunkLordBattleRecordScRsp.proto

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

type GetPunkLordBattleRecordScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CNJBLBIAPLC []*PunkLordBattleReplay `protobuf:"bytes,8,rep,name=CNJBLBIAPLC,proto3" json:"CNJBLBIAPLC,omitempty"`
	OIMOMMOJJBJ []*PunkLordBattleRecord `protobuf:"bytes,14,rep,name=OIMOMMOJJBJ,proto3" json:"OIMOMMOJJBJ,omitempty"`
	Retcode     uint32                  `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	LGOHHIGCHCG *IBJHGNNGJCD            `protobuf:"bytes,10,opt,name=LGOHHIGCHCG,proto3" json:"LGOHHIGCHCG,omitempty"`
}

func (x *GetPunkLordBattleRecordScRsp) Reset() {
	*x = GetPunkLordBattleRecordScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPunkLordBattleRecordScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPunkLordBattleRecordScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPunkLordBattleRecordScRsp) ProtoMessage() {}

func (x *GetPunkLordBattleRecordScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPunkLordBattleRecordScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPunkLordBattleRecordScRsp.ProtoReflect.Descriptor instead.
func (*GetPunkLordBattleRecordScRsp) Descriptor() ([]byte, []int) {
	return file_GetPunkLordBattleRecordScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPunkLordBattleRecordScRsp) GetCNJBLBIAPLC() []*PunkLordBattleReplay {
	if x != nil {
		return x.CNJBLBIAPLC
	}
	return nil
}

func (x *GetPunkLordBattleRecordScRsp) GetOIMOMMOJJBJ() []*PunkLordBattleRecord {
	if x != nil {
		return x.OIMOMMOJJBJ
	}
	return nil
}

func (x *GetPunkLordBattleRecordScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPunkLordBattleRecordScRsp) GetLGOHHIGCHCG() *IBJHGNNGJCD {
	if x != nil {
		return x.LGOHHIGCHCG
	}
	return nil
}

var File_GetPunkLordBattleRecordScRsp_proto protoreflect.FileDescriptor

var file_GetPunkLordBattleRecordScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x49, 0x42, 0x4a, 0x48, 0x47, 0x4e, 0x4e, 0x47, 0x4a, 0x43, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xda, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x37, 0x0a, 0x0b, 0x43, 0x4e, 0x4a, 0x42, 0x4c, 0x42, 0x49, 0x41, 0x50, 0x4c, 0x43, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x0b, 0x43, 0x4e,
	0x4a, 0x42, 0x4c, 0x42, 0x49, 0x41, 0x50, 0x4c, 0x43, 0x12, 0x37, 0x0a, 0x0b, 0x4f, 0x49, 0x4d,
	0x4f, 0x4d, 0x4d, 0x4f, 0x4a, 0x4a, 0x42, 0x4a, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x4f, 0x49, 0x4d, 0x4f, 0x4d, 0x4d, 0x4f, 0x4a, 0x4a,
	0x42, 0x4a, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b,
	0x4c, 0x47, 0x4f, 0x48, 0x48, 0x49, 0x47, 0x43, 0x48, 0x43, 0x47, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x42, 0x4a, 0x48, 0x47, 0x4e, 0x4e, 0x47, 0x4a, 0x43, 0x44, 0x52,
	0x0b, 0x4c, 0x47, 0x4f, 0x48, 0x48, 0x49, 0x47, 0x43, 0x48, 0x43, 0x47, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPunkLordBattleRecordScRsp_proto_rawDescOnce sync.Once
	file_GetPunkLordBattleRecordScRsp_proto_rawDescData = file_GetPunkLordBattleRecordScRsp_proto_rawDesc
)

func file_GetPunkLordBattleRecordScRsp_proto_rawDescGZIP() []byte {
	file_GetPunkLordBattleRecordScRsp_proto_rawDescOnce.Do(func() {
		file_GetPunkLordBattleRecordScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPunkLordBattleRecordScRsp_proto_rawDescData)
	})
	return file_GetPunkLordBattleRecordScRsp_proto_rawDescData
}

var file_GetPunkLordBattleRecordScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPunkLordBattleRecordScRsp_proto_goTypes = []interface{}{
	(*GetPunkLordBattleRecordScRsp)(nil), // 0: GetPunkLordBattleRecordScRsp
	(*PunkLordBattleReplay)(nil),         // 1: PunkLordBattleReplay
	(*PunkLordBattleRecord)(nil),         // 2: PunkLordBattleRecord
	(*IBJHGNNGJCD)(nil),                  // 3: IBJHGNNGJCD
}
var file_GetPunkLordBattleRecordScRsp_proto_depIdxs = []int32{
	1, // 0: GetPunkLordBattleRecordScRsp.CNJBLBIAPLC:type_name -> PunkLordBattleReplay
	2, // 1: GetPunkLordBattleRecordScRsp.OIMOMMOJJBJ:type_name -> PunkLordBattleRecord
	3, // 2: GetPunkLordBattleRecordScRsp.LGOHHIGCHCG:type_name -> IBJHGNNGJCD
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetPunkLordBattleRecordScRsp_proto_init() }
func file_GetPunkLordBattleRecordScRsp_proto_init() {
	if File_GetPunkLordBattleRecordScRsp_proto != nil {
		return
	}
	file_PunkLordBattleReplay_proto_init()
	file_IBJHGNNGJCD_proto_init()
	file_PunkLordBattleRecord_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetPunkLordBattleRecordScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPunkLordBattleRecordScRsp); i {
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
			RawDescriptor: file_GetPunkLordBattleRecordScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPunkLordBattleRecordScRsp_proto_goTypes,
		DependencyIndexes: file_GetPunkLordBattleRecordScRsp_proto_depIdxs,
		MessageInfos:      file_GetPunkLordBattleRecordScRsp_proto_msgTypes,
	}.Build()
	File_GetPunkLordBattleRecordScRsp_proto = out.File
	file_GetPunkLordBattleRecordScRsp_proto_rawDesc = nil
	file_GetPunkLordBattleRecordScRsp_proto_goTypes = nil
	file_GetPunkLordBattleRecordScRsp_proto_depIdxs = nil
}
