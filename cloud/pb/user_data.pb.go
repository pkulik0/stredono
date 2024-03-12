// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pb/user_data.proto

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

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings   *UserSettings     `protobuf:"bytes,1,opt,name=Settings,proto3" json:"Settings,omitempty"`
	Media      *MediaRequest     `protobuf:"bytes,2,opt,name=Media,proto3" json:"Media,omitempty"`
	Commands   map[string]string `protobuf:"bytes,3,rep,name=Commands,proto3" json:"Commands,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OverlayKey string            `protobuf:"bytes,4,opt,name=OverlayKey,proto3" json:"OverlayKey,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_pb_user_data_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetSettings() *UserSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *UserData) GetMedia() *MediaRequest {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *UserData) GetCommands() map[string]string {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *UserData) GetOverlayKey() string {
	if x != nil {
		return x.OverlayKey
	}
	return ""
}

type UserSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tips   *TipSettings    `protobuf:"bytes,1,opt,name=Tips,proto3" json:"Tips,omitempty"`
	Events *EventsSettings `protobuf:"bytes,2,opt,name=Events,proto3" json:"Events,omitempty"`
	Alerts []*Alert        `protobuf:"bytes,3,rep,name=Alerts,proto3" json:"Alerts,omitempty"`
}

func (x *UserSettings) Reset() {
	*x = UserSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSettings) ProtoMessage() {}

func (x *UserSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSettings.ProtoReflect.Descriptor instead.
func (*UserSettings) Descriptor() ([]byte, []int) {
	return file_pb_user_data_proto_rawDescGZIP(), []int{1}
}

func (x *UserSettings) GetTips() *TipSettings {
	if x != nil {
		return x.Tips
	}
	return nil
}

func (x *UserSettings) GetEvents() *EventsSettings {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *UserSettings) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_pb_user_data_proto protoreflect.FileDescriptor

var file_pb_user_data_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x1a, 0x16,
	0x70, 0x62, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x62, 0x2f, 0x74, 0x74, 0x73, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70,
	0x62, 0x2f, 0x74, 0x69, 0x70, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x70, 0x62, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87,
	0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x2c, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x3c, 0x0a,
	0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4f,
	0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x54, 0x69, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f,
	0x6e, 0x6f, 0x2e, 0x54, 0x69, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x04,
	0x54, 0x69, 0x70, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x06,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e,
	0x6f, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6b,
	0x75, 0x6c, 0x69, 0x6b, 0x30, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_data_proto_rawDescOnce sync.Once
	file_pb_user_data_proto_rawDescData = file_pb_user_data_proto_rawDesc
)

func file_pb_user_data_proto_rawDescGZIP() []byte {
	file_pb_user_data_proto_rawDescOnce.Do(func() {
		file_pb_user_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_data_proto_rawDescData)
	})
	return file_pb_user_data_proto_rawDescData
}

var file_pb_user_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_user_data_proto_goTypes = []interface{}{
	(*UserData)(nil),       // 0: stredono.UserData
	(*UserSettings)(nil),   // 1: stredono.UserSettings
	nil,                    // 2: stredono.UserData.CommandsEntry
	(*MediaRequest)(nil),   // 3: stredono.MediaRequest
	(*TipSettings)(nil),    // 4: stredono.TipSettings
	(*EventsSettings)(nil), // 5: stredono.EventsSettings
	(*Alert)(nil),          // 6: stredono.Alert
}
var file_pb_user_data_proto_depIdxs = []int32{
	1, // 0: stredono.UserData.Settings:type_name -> stredono.UserSettings
	3, // 1: stredono.UserData.Media:type_name -> stredono.MediaRequest
	2, // 2: stredono.UserData.Commands:type_name -> stredono.UserData.CommandsEntry
	4, // 3: stredono.UserSettings.Tips:type_name -> stredono.TipSettings
	5, // 4: stredono.UserSettings.Events:type_name -> stredono.EventsSettings
	6, // 5: stredono.UserSettings.Alerts:type_name -> stredono.Alert
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_user_data_proto_init() }
func file_pb_user_data_proto_init() {
	if File_pb_user_data_proto != nil {
		return
	}
	file_pb_media_request_proto_init()
	file_pb_tts_settings_proto_init()
	file_pb_tip_settings_proto_init()
	file_pb_events_settings_proto_init()
	file_pb_alert_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_user_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_pb_user_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSettings); i {
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
			RawDescriptor: file_pb_user_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_user_data_proto_goTypes,
		DependencyIndexes: file_pb_user_data_proto_depIdxs,
		MessageInfos:      file_pb_user_data_proto_msgTypes,
	}.Build()
	File_pb_user_data_proto = out.File
	file_pb_user_data_proto_rawDesc = nil
	file_pb_user_data_proto_goTypes = nil
	file_pb_user_data_proto_depIdxs = nil
}
