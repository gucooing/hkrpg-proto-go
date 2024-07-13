// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TreasureDungeonLevel.proto

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

type TreasureDungeonLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IKHIMOPBOIK bool                         `protobuf:"varint,808,opt,name=IKHIMOPBOIK,proto3" json:"IKHIMOPBOIK,omitempty"`
	FDOBNHFFPBE []*TreasureDungeonRecordData `protobuf:"bytes,14,rep,name=FDOBNHFFPBE,proto3" json:"FDOBNHFFPBE,omitempty"`
	CPCGPGBDKDG uint32                       `protobuf:"varint,11,opt,name=CPCGPGBDKDG,proto3" json:"CPCGPGBDKDG,omitempty"`
	MapId       uint32                       `protobuf:"varint,3,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	PMFCBFMFNLI uint32                       `protobuf:"varint,15,opt,name=PMFCBFMFNLI,proto3" json:"PMFCBFMFNLI,omitempty"`
	CAIKGAFLPPG bool                         `protobuf:"varint,1957,opt,name=CAIKGAFLPPG,proto3" json:"CAIKGAFLPPG,omitempty"`
	AvatarList  []*GJHNLALENEP               `protobuf:"bytes,1040,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	FGMBCCJEGJI bool                         `protobuf:"varint,2038,opt,name=FGMBCCJEGJI,proto3" json:"FGMBCCJEGJI,omitempty"`
	GIADELONBAA []*FODJFBNFPJC               `protobuf:"bytes,8,rep,name=GIADELONBAA,proto3" json:"GIADELONBAA,omitempty"`
	EDODDFIDMII uint32                       `protobuf:"varint,4,opt,name=EDODDFIDMII,proto3" json:"EDODDFIDMII,omitempty"`
	JNBJLDBJJOM []*GJHNLALENEP               `protobuf:"bytes,1137,rep,name=JNBJLDBJJOM,proto3" json:"JNBJLDBJJOM,omitempty"`
	ItemList    []*DOJPPNDKNAC               `protobuf:"bytes,2041,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	IPOECDDLFOH []*OMPCPEGHAID               `protobuf:"bytes,386,rep,name=IPOECDDLFOH,proto3" json:"IPOECDDLFOH,omitempty"`
	BuffList    []*OAPECJLDCGL               `protobuf:"bytes,929,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	KFKMBIDIIOD uint32                       `protobuf:"varint,1,opt,name=KFKMBIDIIOD,proto3" json:"KFKMBIDIIOD,omitempty"`
	MPPNAMEEEFM uint32                       `protobuf:"varint,12,opt,name=MPPNAMEEEFM,proto3" json:"MPPNAMEEEFM,omitempty"`
	AKLPPEMEEAJ uint32                       `protobuf:"varint,9,opt,name=AKLPPEMEEAJ,proto3" json:"AKLPPEMEEAJ,omitempty"`
	GADHCMJJEAB uint32                       `protobuf:"varint,1699,opt,name=GADHCMJJEAB,proto3" json:"GADHCMJJEAB,omitempty"`
}

func (x *TreasureDungeonLevel) Reset() {
	*x = TreasureDungeonLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureDungeonLevel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureDungeonLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureDungeonLevel) ProtoMessage() {}

func (x *TreasureDungeonLevel) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureDungeonLevel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureDungeonLevel.ProtoReflect.Descriptor instead.
func (*TreasureDungeonLevel) Descriptor() ([]byte, []int) {
	return file_TreasureDungeonLevel_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureDungeonLevel) GetIKHIMOPBOIK() bool {
	if x != nil {
		return x.IKHIMOPBOIK
	}
	return false
}

func (x *TreasureDungeonLevel) GetFDOBNHFFPBE() []*TreasureDungeonRecordData {
	if x != nil {
		return x.FDOBNHFFPBE
	}
	return nil
}

func (x *TreasureDungeonLevel) GetCPCGPGBDKDG() uint32 {
	if x != nil {
		return x.CPCGPGBDKDG
	}
	return 0
}

func (x *TreasureDungeonLevel) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *TreasureDungeonLevel) GetPMFCBFMFNLI() uint32 {
	if x != nil {
		return x.PMFCBFMFNLI
	}
	return 0
}

func (x *TreasureDungeonLevel) GetCAIKGAFLPPG() bool {
	if x != nil {
		return x.CAIKGAFLPPG
	}
	return false
}

func (x *TreasureDungeonLevel) GetAvatarList() []*GJHNLALENEP {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetFGMBCCJEGJI() bool {
	if x != nil {
		return x.FGMBCCJEGJI
	}
	return false
}

func (x *TreasureDungeonLevel) GetGIADELONBAA() []*FODJFBNFPJC {
	if x != nil {
		return x.GIADELONBAA
	}
	return nil
}

func (x *TreasureDungeonLevel) GetEDODDFIDMII() uint32 {
	if x != nil {
		return x.EDODDFIDMII
	}
	return 0
}

func (x *TreasureDungeonLevel) GetJNBJLDBJJOM() []*GJHNLALENEP {
	if x != nil {
		return x.JNBJLDBJJOM
	}
	return nil
}

func (x *TreasureDungeonLevel) GetItemList() []*DOJPPNDKNAC {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetIPOECDDLFOH() []*OMPCPEGHAID {
	if x != nil {
		return x.IPOECDDLFOH
	}
	return nil
}

func (x *TreasureDungeonLevel) GetBuffList() []*OAPECJLDCGL {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetKFKMBIDIIOD() uint32 {
	if x != nil {
		return x.KFKMBIDIIOD
	}
	return 0
}

func (x *TreasureDungeonLevel) GetMPPNAMEEEFM() uint32 {
	if x != nil {
		return x.MPPNAMEEEFM
	}
	return 0
}

func (x *TreasureDungeonLevel) GetAKLPPEMEEAJ() uint32 {
	if x != nil {
		return x.AKLPPEMEEAJ
	}
	return 0
}

func (x *TreasureDungeonLevel) GetGADHCMJJEAB() uint32 {
	if x != nil {
		return x.GADHCMJJEAB
	}
	return 0
}

var File_TreasureDungeonLevel_proto protoreflect.FileDescriptor

var file_TreasureDungeonLevel_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4d,
	0x50, 0x43, 0x50, 0x45, 0x47, 0x48, 0x41, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x44, 0x4f, 0x4a, 0x50, 0x50, 0x4e, 0x44, 0x4b, 0x4e, 0x41, 0x43, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x41, 0x50, 0x45, 0x43, 0x4a, 0x4c, 0x44, 0x43, 0x47, 0x4c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x4f, 0x44, 0x4a, 0x46, 0x42, 0x4e, 0x46, 0x50,
	0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4a, 0x48, 0x4e, 0x4c,
	0x41, 0x4c, 0x45, 0x4e, 0x45, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x05, 0x0a,
	0x14, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0b, 0x49, 0x4b, 0x48, 0x49, 0x4d, 0x4f, 0x50,
	0x42, 0x4f, 0x49, 0x4b, 0x18, 0xa8, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x4b, 0x48,
	0x49, 0x4d, 0x4f, 0x50, 0x42, 0x4f, 0x49, 0x4b, 0x12, 0x3c, 0x0a, 0x0b, 0x46, 0x44, 0x4f, 0x42,
	0x4e, 0x48, 0x46, 0x46, 0x50, 0x42, 0x45, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x46, 0x44, 0x4f, 0x42, 0x4e,
	0x48, 0x46, 0x46, 0x50, 0x42, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x50, 0x43, 0x47, 0x50, 0x47,
	0x42, 0x44, 0x4b, 0x44, 0x47, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x50, 0x43,
	0x47, 0x50, 0x47, 0x42, 0x44, 0x4b, 0x44, 0x47, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x4d, 0x46, 0x43, 0x42, 0x46, 0x4d, 0x46, 0x4e, 0x4c, 0x49, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x4d, 0x46, 0x43, 0x42, 0x46, 0x4d, 0x46, 0x4e, 0x4c,
	0x49, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x41, 0x49, 0x4b, 0x47, 0x41, 0x46, 0x4c, 0x50, 0x50, 0x47,
	0x18, 0xa5, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x41, 0x49, 0x4b, 0x47, 0x41, 0x46,
	0x4c, 0x50, 0x50, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x90, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4a, 0x48,
	0x4e, 0x4c, 0x41, 0x4c, 0x45, 0x4e, 0x45, 0x50, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x47, 0x4d, 0x42, 0x43, 0x43, 0x4a, 0x45,
	0x47, 0x4a, 0x49, 0x18, 0xf6, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x47, 0x4d, 0x42,
	0x43, 0x43, 0x4a, 0x45, 0x47, 0x4a, 0x49, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x49, 0x41, 0x44, 0x45,
	0x4c, 0x4f, 0x4e, 0x42, 0x41, 0x41, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46,
	0x4f, 0x44, 0x4a, 0x46, 0x42, 0x4e, 0x46, 0x50, 0x4a, 0x43, 0x52, 0x0b, 0x47, 0x49, 0x41, 0x44,
	0x45, 0x4c, 0x4f, 0x4e, 0x42, 0x41, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x44, 0x4f, 0x44, 0x44,
	0x46, 0x49, 0x44, 0x4d, 0x49, 0x49, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x44,
	0x4f, 0x44, 0x44, 0x46, 0x49, 0x44, 0x4d, 0x49, 0x49, 0x12, 0x2f, 0x0a, 0x0b, 0x4a, 0x4e, 0x42,
	0x4a, 0x4c, 0x44, 0x42, 0x4a, 0x4a, 0x4f, 0x4d, 0x18, 0xf1, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x47, 0x4a, 0x48, 0x4e, 0x4c, 0x41, 0x4c, 0x45, 0x4e, 0x45, 0x50, 0x52, 0x0b, 0x4a,
	0x4e, 0x42, 0x4a, 0x4c, 0x44, 0x42, 0x4a, 0x4a, 0x4f, 0x4d, 0x12, 0x2a, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xf9, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x44, 0x4f, 0x4a, 0x50, 0x50, 0x4e, 0x44, 0x4b, 0x4e, 0x41, 0x43, 0x52, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x49, 0x50, 0x4f, 0x45, 0x43, 0x44,
	0x44, 0x4c, 0x46, 0x4f, 0x48, 0x18, 0x82, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f,
	0x4d, 0x50, 0x43, 0x50, 0x45, 0x47, 0x48, 0x41, 0x49, 0x44, 0x52, 0x0b, 0x49, 0x50, 0x4f, 0x45,
	0x43, 0x44, 0x44, 0x4c, 0x46, 0x4f, 0x48, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0xa1, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x41,
	0x50, 0x45, 0x43, 0x4a, 0x4c, 0x44, 0x43, 0x47, 0x4c, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x46, 0x4b, 0x4d, 0x42, 0x49, 0x44, 0x49, 0x49,
	0x4f, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x46, 0x4b, 0x4d, 0x42, 0x49,
	0x44, 0x49, 0x49, 0x4f, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x50, 0x50, 0x4e, 0x41, 0x4d, 0x45,
	0x45, 0x45, 0x46, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x50, 0x50, 0x4e,
	0x41, 0x4d, 0x45, 0x45, 0x45, 0x46, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x4b, 0x4c, 0x50, 0x50,
	0x45, 0x4d, 0x45, 0x45, 0x41, 0x4a, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4b,
	0x4c, 0x50, 0x50, 0x45, 0x4d, 0x45, 0x45, 0x41, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x47, 0x41, 0x44,
	0x48, 0x43, 0x4d, 0x4a, 0x4a, 0x45, 0x41, 0x42, 0x18, 0xa3, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x47, 0x41, 0x44, 0x48, 0x43, 0x4d, 0x4a, 0x4a, 0x45, 0x41, 0x42, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TreasureDungeonLevel_proto_rawDescOnce sync.Once
	file_TreasureDungeonLevel_proto_rawDescData = file_TreasureDungeonLevel_proto_rawDesc
)

func file_TreasureDungeonLevel_proto_rawDescGZIP() []byte {
	file_TreasureDungeonLevel_proto_rawDescOnce.Do(func() {
		file_TreasureDungeonLevel_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureDungeonLevel_proto_rawDescData)
	})
	return file_TreasureDungeonLevel_proto_rawDescData
}

var file_TreasureDungeonLevel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureDungeonLevel_proto_goTypes = []interface{}{
	(*TreasureDungeonLevel)(nil),      // 0: TreasureDungeonLevel
	(*TreasureDungeonRecordData)(nil), // 1: TreasureDungeonRecordData
	(*GJHNLALENEP)(nil),               // 2: GJHNLALENEP
	(*FODJFBNFPJC)(nil),               // 3: FODJFBNFPJC
	(*DOJPPNDKNAC)(nil),               // 4: DOJPPNDKNAC
	(*OMPCPEGHAID)(nil),               // 5: OMPCPEGHAID
	(*OAPECJLDCGL)(nil),               // 6: OAPECJLDCGL
}
var file_TreasureDungeonLevel_proto_depIdxs = []int32{
	1, // 0: TreasureDungeonLevel.FDOBNHFFPBE:type_name -> TreasureDungeonRecordData
	2, // 1: TreasureDungeonLevel.avatar_list:type_name -> GJHNLALENEP
	3, // 2: TreasureDungeonLevel.GIADELONBAA:type_name -> FODJFBNFPJC
	2, // 3: TreasureDungeonLevel.JNBJLDBJJOM:type_name -> GJHNLALENEP
	4, // 4: TreasureDungeonLevel.item_list:type_name -> DOJPPNDKNAC
	5, // 5: TreasureDungeonLevel.IPOECDDLFOH:type_name -> OMPCPEGHAID
	6, // 6: TreasureDungeonLevel.buff_list:type_name -> OAPECJLDCGL
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_TreasureDungeonLevel_proto_init() }
func file_TreasureDungeonLevel_proto_init() {
	if File_TreasureDungeonLevel_proto != nil {
		return
	}
	file_OMPCPEGHAID_proto_init()
	file_DOJPPNDKNAC_proto_init()
	file_OAPECJLDCGL_proto_init()
	file_FODJFBNFPJC_proto_init()
	file_TreasureDungeonRecordData_proto_init()
	file_GJHNLALENEP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TreasureDungeonLevel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureDungeonLevel); i {
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
			RawDescriptor: file_TreasureDungeonLevel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureDungeonLevel_proto_goTypes,
		DependencyIndexes: file_TreasureDungeonLevel_proto_depIdxs,
		MessageInfos:      file_TreasureDungeonLevel_proto_msgTypes,
	}.Build()
	File_TreasureDungeonLevel_proto = out.File
	file_TreasureDungeonLevel_proto_rawDesc = nil
	file_TreasureDungeonLevel_proto_goTypes = nil
	file_TreasureDungeonLevel_proto_depIdxs = nil
}
