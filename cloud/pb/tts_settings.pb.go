// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pb/tts_settings.proto

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

type TTSSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceIdBasic string `protobuf:"bytes,1,opt,name=VoiceIdBasic,proto3" json:"VoiceIdBasic,omitempty"`
	VoiceIdPlus  string `protobuf:"bytes,2,opt,name=VoiceIdPlus,proto3" json:"VoiceIdPlus,omitempty"`
	Tier         Tier   `protobuf:"varint,3,opt,name=Tier,proto3,enum=stredono.Tier" json:"Tier,omitempty"`
	// string = EventType name (more stable than int32 I guess)
	EventSettings map[string]*EventSettings `protobuf:"bytes,4,rep,name=EventSettings,proto3" json:"EventSettings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TTSSettings) Reset() {
	*x = TTSSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tts_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TTSSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TTSSettings) ProtoMessage() {}

func (x *TTSSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tts_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TTSSettings.ProtoReflect.Descriptor instead.
func (*TTSSettings) Descriptor() ([]byte, []int) {
	return file_pb_tts_settings_proto_rawDescGZIP(), []int{0}
}

func (x *TTSSettings) GetVoiceIdBasic() string {
	if x != nil {
		return x.VoiceIdBasic
	}
	return ""
}

func (x *TTSSettings) GetVoiceIdPlus() string {
	if x != nil {
		return x.VoiceIdPlus
	}
	return ""
}

func (x *TTSSettings) GetTier() Tier {
	if x != nil {
		return x.Tier
	}
	return Tier_BASIC
}

func (x *TTSSettings) GetEventSettings() map[string]*EventSettings {
	if x != nil {
		return x.EventSettings
	}
	return nil
}

type EventSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinimumValue  int32  `protobuf:"varint,1,opt,name=MinimumValue,proto3" json:"MinimumValue,omitempty"`
	EnableTTS     bool   `protobuf:"varint,2,opt,name=EnableTTS,proto3" json:"EnableTTS,omitempty"`
	MinimumForTTS *int32 `protobuf:"varint,3,opt,name=MinimumForTTS,proto3,oneof" json:"MinimumForTTS,omitempty"`
}

func (x *EventSettings) Reset() {
	*x = EventSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tts_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSettings) ProtoMessage() {}

func (x *EventSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tts_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSettings.ProtoReflect.Descriptor instead.
func (*EventSettings) Descriptor() ([]byte, []int) {
	return file_pb_tts_settings_proto_rawDescGZIP(), []int{1}
}

func (x *EventSettings) GetMinimumValue() int32 {
	if x != nil {
		return x.MinimumValue
	}
	return 0
}

func (x *EventSettings) GetEnableTTS() bool {
	if x != nil {
		return x.EnableTTS
	}
	return false
}

func (x *EventSettings) GetMinimumForTTS() int32 {
	if x != nil && x.MinimumForTTS != nil {
		return *x.MinimumForTTS
	}
	return 0
}

var File_pb_tts_settings_proto protoreflect.FileDescriptor

var file_pb_tts_settings_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x74, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e,
	0x6f, 0x1a, 0x0e, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x0b, 0x54, 0x54, 0x53, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x50, 0x6c, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x50, 0x6c, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f,
	0x2e, 0x54, 0x69, 0x65, 0x72, 0x52, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x54, 0x54,
	0x53, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x59, 0x0a, 0x12, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x69,
	0x6d, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x54, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x54, 0x53, 0x12, 0x29, 0x0a, 0x0d, 0x4d, 0x69,
	0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x46, 0x6f, 0x72, 0x54, 0x54, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x46, 0x6f, 0x72, 0x54,
	0x54, 0x53, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75,
	0x6d, 0x46, 0x6f, 0x72, 0x54, 0x54, 0x53, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6b, 0x75, 0x6c, 0x69, 0x6b, 0x30, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_tts_settings_proto_rawDescOnce sync.Once
	file_pb_tts_settings_proto_rawDescData = file_pb_tts_settings_proto_rawDesc
)

func file_pb_tts_settings_proto_rawDescGZIP() []byte {
	file_pb_tts_settings_proto_rawDescOnce.Do(func() {
		file_pb_tts_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_tts_settings_proto_rawDescData)
	})
	return file_pb_tts_settings_proto_rawDescData
}

var file_pb_tts_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_tts_settings_proto_goTypes = []interface{}{
	(*TTSSettings)(nil),   // 0: stredono.TTSSettings
	(*EventSettings)(nil), // 1: stredono.EventSettings
	nil,                   // 2: stredono.TTSSettings.EventSettingsEntry
	(Tier)(0),             // 3: stredono.Tier
}
var file_pb_tts_settings_proto_depIdxs = []int32{
	3, // 0: stredono.TTSSettings.Tier:type_name -> stredono.Tier
	2, // 1: stredono.TTSSettings.EventSettings:type_name -> stredono.TTSSettings.EventSettingsEntry
	1, // 2: stredono.TTSSettings.EventSettingsEntry.value:type_name -> stredono.EventSettings
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_tts_settings_proto_init() }
func file_pb_tts_settings_proto_init() {
	if File_pb_tts_settings_proto != nil {
		return
	}
	file_pb_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_tts_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TTSSettings); i {
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
		file_pb_tts_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSettings); i {
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
	file_pb_tts_settings_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_tts_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_tts_settings_proto_goTypes,
		DependencyIndexes: file_pb_tts_settings_proto_depIdxs,
		MessageInfos:      file_pb_tts_settings_proto_msgTypes,
	}.Build()
	File_pb_tts_settings_proto = out.File
	file_pb_tts_settings_proto_rawDesc = nil
	file_pb_tts_settings_proto_goTypes = nil
	file_pb_tts_settings_proto_depIdxs = nil
}