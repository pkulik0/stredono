// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pb/bot_message.proto

package pb

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

type BotMessage_AnnouncementData_AnnouncementColor int32

const (
	BotMessage_AnnouncementData_PRIMARY BotMessage_AnnouncementData_AnnouncementColor = 0
	BotMessage_AnnouncementData_BLUE    BotMessage_AnnouncementData_AnnouncementColor = 1
	BotMessage_AnnouncementData_GREEN   BotMessage_AnnouncementData_AnnouncementColor = 2
	BotMessage_AnnouncementData_ORANGE  BotMessage_AnnouncementData_AnnouncementColor = 3
	BotMessage_AnnouncementData_PURPLE  BotMessage_AnnouncementData_AnnouncementColor = 4
)

// Enum value maps for BotMessage_AnnouncementData_AnnouncementColor.
var (
	BotMessage_AnnouncementData_AnnouncementColor_name = map[int32]string{
		0: "PRIMARY",
		1: "BLUE",
		2: "GREEN",
		3: "ORANGE",
		4: "PURPLE",
	}
	BotMessage_AnnouncementData_AnnouncementColor_value = map[string]int32{
		"PRIMARY": 0,
		"BLUE":    1,
		"GREEN":   2,
		"ORANGE":  3,
		"PURPLE":  4,
	}
)

func (x BotMessage_AnnouncementData_AnnouncementColor) Enum() *BotMessage_AnnouncementData_AnnouncementColor {
	p := new(BotMessage_AnnouncementData_AnnouncementColor)
	*p = x
	return p
}

func (x BotMessage_AnnouncementData_AnnouncementColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BotMessage_AnnouncementData_AnnouncementColor) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_bot_message_proto_enumTypes[0].Descriptor()
}

func (BotMessage_AnnouncementData_AnnouncementColor) Type() protoreflect.EnumType {
	return &file_pb_bot_message_proto_enumTypes[0]
}

func (x BotMessage_AnnouncementData_AnnouncementColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BotMessage_AnnouncementData_AnnouncementColor.Descriptor instead.
func (BotMessage_AnnouncementData_AnnouncementColor) EnumDescriptor() ([]byte, []int) {
	return file_pb_bot_message_proto_rawDescGZIP(), []int{0, 1, 0}
}

type BotMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatID  string `protobuf:"bytes,1,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	// Types that are assignable to Data:
	//
	//	*BotMessage_Normal
	//	*BotMessage_Announcement
	Data isBotMessage_Data `protobuf_oneof:"Data"`
}

func (x *BotMessage) Reset() {
	*x = BotMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bot_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotMessage) ProtoMessage() {}

func (x *BotMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bot_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotMessage.ProtoReflect.Descriptor instead.
func (*BotMessage) Descriptor() ([]byte, []int) {
	return file_pb_bot_message_proto_rawDescGZIP(), []int{0}
}

func (x *BotMessage) GetChatID() string {
	if x != nil {
		return x.ChatID
	}
	return ""
}

func (x *BotMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (m *BotMessage) GetData() isBotMessage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *BotMessage) GetNormal() *BotMessage_NormalData {
	if x, ok := x.GetData().(*BotMessage_Normal); ok {
		return x.Normal
	}
	return nil
}

func (x *BotMessage) GetAnnouncement() *BotMessage_AnnouncementData {
	if x, ok := x.GetData().(*BotMessage_Announcement); ok {
		return x.Announcement
	}
	return nil
}

type isBotMessage_Data interface {
	isBotMessage_Data()
}

type BotMessage_Normal struct {
	Normal *BotMessage_NormalData `protobuf:"bytes,3,opt,name=Normal,proto3,oneof"`
}

type BotMessage_Announcement struct {
	Announcement *BotMessage_AnnouncementData `protobuf:"bytes,4,opt,name=Announcement,proto3,oneof"`
}

func (*BotMessage_Normal) isBotMessage_Data() {}

func (*BotMessage_Announcement) isBotMessage_Data() {}

type BotMessage_NormalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplyMessageID string `protobuf:"bytes,1,opt,name=ReplyMessageID,proto3" json:"ReplyMessageID,omitempty"`
}

func (x *BotMessage_NormalData) Reset() {
	*x = BotMessage_NormalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bot_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotMessage_NormalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotMessage_NormalData) ProtoMessage() {}

func (x *BotMessage_NormalData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bot_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotMessage_NormalData.ProtoReflect.Descriptor instead.
func (*BotMessage_NormalData) Descriptor() ([]byte, []int) {
	return file_pb_bot_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BotMessage_NormalData) GetReplyMessageID() string {
	if x != nil {
		return x.ReplyMessageID
	}
	return ""
}

type BotMessage_AnnouncementData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color BotMessage_AnnouncementData_AnnouncementColor `protobuf:"varint,1,opt,name=Color,proto3,enum=stredono.BotMessage_AnnouncementData_AnnouncementColor" json:"Color,omitempty"`
}

func (x *BotMessage_AnnouncementData) Reset() {
	*x = BotMessage_AnnouncementData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bot_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotMessage_AnnouncementData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotMessage_AnnouncementData) ProtoMessage() {}

func (x *BotMessage_AnnouncementData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bot_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotMessage_AnnouncementData.ProtoReflect.Descriptor instead.
func (*BotMessage_AnnouncementData) Descriptor() ([]byte, []int) {
	return file_pb_bot_message_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BotMessage_AnnouncementData) GetColor() BotMessage_AnnouncementData_AnnouncementColor {
	if x != nil {
		return x.Color
	}
	return BotMessage_AnnouncementData_PRIMARY
}

var File_pb_bot_message_proto protoreflect.FileDescriptor

var file_pb_bot_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x62, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f,
	0x22, 0xb7, 0x03, 0x0a, 0x0a, 0x42, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x39, 0x0a, 0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x42, 0x6f, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x0c,
	0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x42, 0x6f,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0c, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x34, 0x0a, 0x0a, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x1a,
	0xb0, 0x01, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x42,
	0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x22, 0x4d, 0x0a, 0x11, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x4d,
	0x41, 0x52, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x52, 0x50, 0x4c, 0x45,
	0x10, 0x04, 0x42, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6b, 0x75, 0x6c, 0x69, 0x6b, 0x30,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_bot_message_proto_rawDescOnce sync.Once
	file_pb_bot_message_proto_rawDescData = file_pb_bot_message_proto_rawDesc
)

func file_pb_bot_message_proto_rawDescGZIP() []byte {
	file_pb_bot_message_proto_rawDescOnce.Do(func() {
		file_pb_bot_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_bot_message_proto_rawDescData)
	})
	return file_pb_bot_message_proto_rawDescData
}

var file_pb_bot_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_bot_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_bot_message_proto_goTypes = []interface{}{
	(BotMessage_AnnouncementData_AnnouncementColor)(0), // 0: stredono.BotMessage.AnnouncementData.AnnouncementColor
	(*BotMessage)(nil),                  // 1: stredono.BotMessage
	(*BotMessage_NormalData)(nil),       // 2: stredono.BotMessage.NormalData
	(*BotMessage_AnnouncementData)(nil), // 3: stredono.BotMessage.AnnouncementData
}
var file_pb_bot_message_proto_depIdxs = []int32{
	2, // 0: stredono.BotMessage.Normal:type_name -> stredono.BotMessage.NormalData
	3, // 1: stredono.BotMessage.Announcement:type_name -> stredono.BotMessage.AnnouncementData
	0, // 2: stredono.BotMessage.AnnouncementData.Color:type_name -> stredono.BotMessage.AnnouncementData.AnnouncementColor
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_bot_message_proto_init() }
func file_pb_bot_message_proto_init() {
	if File_pb_bot_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_bot_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotMessage); i {
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
		file_pb_bot_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotMessage_NormalData); i {
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
		file_pb_bot_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotMessage_AnnouncementData); i {
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
	file_pb_bot_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BotMessage_Normal)(nil),
		(*BotMessage_Announcement)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_bot_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_bot_message_proto_goTypes,
		DependencyIndexes: file_pb_bot_message_proto_depIdxs,
		EnumInfos:         file_pb_bot_message_proto_enumTypes,
		MessageInfos:      file_pb_bot_message_proto_msgTypes,
	}.Build()
	File_pb_bot_message_proto = out.File
	file_pb_bot_message_proto_rawDesc = nil
	file_pb_bot_message_proto_goTypes = nil
	file_pb_bot_message_proto_depIdxs = nil
}
