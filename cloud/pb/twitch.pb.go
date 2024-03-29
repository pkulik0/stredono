// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pb/twitch.proto

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

type TokenEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid            string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	EncryptedToken []byte `protobuf:"bytes,2,opt,name=EncryptedToken,proto3" json:"EncryptedToken,omitempty"`
}

func (x *TokenEntry) Reset() {
	*x = TokenEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_twitch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenEntry) ProtoMessage() {}

func (x *TokenEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pb_twitch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenEntry.ProtoReflect.Descriptor instead.
func (*TokenEntry) Descriptor() ([]byte, []int) {
	return file_pb_twitch_proto_rawDescGZIP(), []int{0}
}

func (x *TokenEntry) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TokenEntry) GetEncryptedToken() []byte {
	if x != nil {
		return x.EncryptedToken
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_twitch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_pb_twitch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_pb_twitch_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type TwitchReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Cost      int64  `protobuf:"varint,3,opt,name=Cost,proto3" json:"Cost,omitempty"`
	IsEnabled bool   `protobuf:"varint,4,opt,name=IsEnabled,proto3" json:"IsEnabled,omitempty"`
}

func (x *TwitchReward) Reset() {
	*x = TwitchReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_twitch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwitchReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwitchReward) ProtoMessage() {}

func (x *TwitchReward) ProtoReflect() protoreflect.Message {
	mi := &file_pb_twitch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwitchReward.ProtoReflect.Descriptor instead.
func (*TwitchReward) Descriptor() ([]byte, []int) {
	return file_pb_twitch_proto_rawDescGZIP(), []int{2}
}

func (x *TwitchReward) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TwitchReward) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TwitchReward) GetCost() int64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *TwitchReward) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

type TwitchUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	DisplayName       string `protobuf:"bytes,2,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	Login             string `protobuf:"bytes,3,opt,name=Login,proto3" json:"Login,omitempty"`
	AvatarUrl         string `protobuf:"bytes,4,opt,name=AvatarUrl,proto3" json:"AvatarUrl,omitempty"`
	Description       string `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	CreationTimestamp int64  `protobuf:"varint,6,opt,name=CreationTimestamp,proto3" json:"CreationTimestamp,omitempty"`
}

func (x *TwitchUser) Reset() {
	*x = TwitchUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_twitch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwitchUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwitchUser) ProtoMessage() {}

func (x *TwitchUser) ProtoReflect() protoreflect.Message {
	mi := &file_pb_twitch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwitchUser.ProtoReflect.Descriptor instead.
func (*TwitchUser) Descriptor() ([]byte, []int) {
	return file_pb_twitch_proto_rawDescGZIP(), []int{3}
}

func (x *TwitchUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TwitchUser) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *TwitchUser) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *TwitchUser) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *TwitchUser) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TwitchUser) GetCreationTimestamp() int64 {
	if x != nil {
		return x.CreationTimestamp
	}
	return 0
}

var File_pb_twitch_proto protoreflect.FileDescriptor

var file_pb_twitch_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x1a, 0x0e, 0x70, 0x62, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0a, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x4d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x64, 0x0a, 0x0c, 0x54, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x54, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6b, 0x75, 0x6c,
	0x69, 0x6b, 0x30, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_twitch_proto_rawDescOnce sync.Once
	file_pb_twitch_proto_rawDescData = file_pb_twitch_proto_rawDesc
)

func file_pb_twitch_proto_rawDescGZIP() []byte {
	file_pb_twitch_proto_rawDescOnce.Do(func() {
		file_pb_twitch_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_twitch_proto_rawDescData)
	})
	return file_pb_twitch_proto_rawDescData
}

var file_pb_twitch_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_twitch_proto_goTypes = []interface{}{
	(*TokenEntry)(nil),   // 0: stredono.TokenEntry
	(*Token)(nil),        // 1: stredono.Token
	(*TwitchReward)(nil), // 2: stredono.TwitchReward
	(*TwitchUser)(nil),   // 3: stredono.TwitchUser
}
var file_pb_twitch_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_twitch_proto_init() }
func file_pb_twitch_proto_init() {
	if File_pb_twitch_proto != nil {
		return
	}
	file_pb_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_twitch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenEntry); i {
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
		file_pb_twitch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_pb_twitch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwitchReward); i {
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
		file_pb_twitch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwitchUser); i {
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
			RawDescriptor: file_pb_twitch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_twitch_proto_goTypes,
		DependencyIndexes: file_pb_twitch_proto_depIdxs,
		MessageInfos:      file_pb_twitch_proto_msgTypes,
	}.Build()
	File_pb_twitch_proto = out.File
	file_pb_twitch_proto_rawDesc = nil
	file_pb_twitch_proto_goTypes = nil
	file_pb_twitch_proto_depIdxs = nil
}
