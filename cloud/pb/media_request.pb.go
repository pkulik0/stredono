// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pb/media_request.proto

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

type MediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEnabled         bool                      `protobuf:"varint,1,opt,name=IsEnabled,proto3" json:"IsEnabled,omitempty"`
	IsPlaying         bool                      `protobuf:"varint,2,opt,name=IsPlaying,proto3" json:"IsPlaying,omitempty"`
	Settings          *MediaRequestSettings     `protobuf:"bytes,3,opt,name=Settings,proto3" json:"Settings,omitempty"`
	Queue             []*MediaRequest_QueueItem `protobuf:"bytes,4,rep,name=Queue,proto3" json:"Queue,omitempty"`
	CurrentQueueIndex int32                     `protobuf:"varint,5,opt,name=CurrentQueueIndex,proto3" json:"CurrentQueueIndex,omitempty"`
	RequireApproval   bool                      `protobuf:"varint,6,opt,name=RequireApproval,proto3" json:"RequireApproval,omitempty"`
}

func (x *MediaRequest) Reset() {
	*x = MediaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_media_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaRequest) ProtoMessage() {}

func (x *MediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_media_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaRequest.ProtoReflect.Descriptor instead.
func (*MediaRequest) Descriptor() ([]byte, []int) {
	return file_pb_media_request_proto_rawDescGZIP(), []int{0}
}

func (x *MediaRequest) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *MediaRequest) GetIsPlaying() bool {
	if x != nil {
		return x.IsPlaying
	}
	return false
}

func (x *MediaRequest) GetSettings() *MediaRequestSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *MediaRequest) GetQueue() []*MediaRequest_QueueItem {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *MediaRequest) GetCurrentQueueIndex() int32 {
	if x != nil {
		return x.CurrentQueueIndex
	}
	return 0
}

func (x *MediaRequest) GetRequireApproval() bool {
	if x != nil {
		return x.RequireApproval
	}
	return false
}

type MediaRequestSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinRole         Role  `protobuf:"varint,1,opt,name=MinRole,proto3,enum=stredono.Role" json:"MinRole,omitempty"`
	MinViews        int32 `protobuf:"varint,2,opt,name=MinViews,proto3" json:"MinViews,omitempty"`
	MinLikes        int32 `protobuf:"varint,3,opt,name=MinLikes,proto3" json:"MinLikes,omitempty"`
	RequireApproval bool  `protobuf:"varint,4,opt,name=RequireApproval,proto3" json:"RequireApproval,omitempty"`
}

func (x *MediaRequestSettings) Reset() {
	*x = MediaRequestSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_media_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaRequestSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaRequestSettings) ProtoMessage() {}

func (x *MediaRequestSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pb_media_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaRequestSettings.ProtoReflect.Descriptor instead.
func (*MediaRequestSettings) Descriptor() ([]byte, []int) {
	return file_pb_media_request_proto_rawDescGZIP(), []int{1}
}

func (x *MediaRequestSettings) GetMinRole() Role {
	if x != nil {
		return x.MinRole
	}
	return Role_NORMAL
}

func (x *MediaRequestSettings) GetMinViews() int32 {
	if x != nil {
		return x.MinViews
	}
	return 0
}

func (x *MediaRequestSettings) GetMinLikes() int32 {
	if x != nil {
		return x.MinLikes
	}
	return 0
}

func (x *MediaRequestSettings) GetRequireApproval() bool {
	if x != nil {
		return x.RequireApproval
	}
	return false
}

type MediaRequest_QueueItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL               string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	RequesterID       string `protobuf:"bytes,2,opt,name=RequesterID,proto3" json:"RequesterID,omitempty"`
	RequesterName     string `protobuf:"bytes,3,opt,name=RequesterName,proto3" json:"RequesterName,omitempty"`
	RequesterProvider string `protobuf:"bytes,4,opt,name=RequesterProvider,proto3" json:"RequesterProvider,omitempty"`
	Timestamp         int64  `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	IsApproved        bool   `protobuf:"varint,6,opt,name=IsApproved,proto3" json:"IsApproved,omitempty"`
	Progress          int32  `protobuf:"varint,7,opt,name=Progress,proto3" json:"Progress,omitempty"`
}

func (x *MediaRequest_QueueItem) Reset() {
	*x = MediaRequest_QueueItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_media_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaRequest_QueueItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaRequest_QueueItem) ProtoMessage() {}

func (x *MediaRequest_QueueItem) ProtoReflect() protoreflect.Message {
	mi := &file_pb_media_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaRequest_QueueItem.ProtoReflect.Descriptor instead.
func (*MediaRequest_QueueItem) Descriptor() ([]byte, []int) {
	return file_pb_media_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MediaRequest_QueueItem) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *MediaRequest_QueueItem) GetRequesterID() string {
	if x != nil {
		return x.RequesterID
	}
	return ""
}

func (x *MediaRequest_QueueItem) GetRequesterName() string {
	if x != nil {
		return x.RequesterName
	}
	return ""
}

func (x *MediaRequest_QueueItem) GetRequesterProvider() string {
	if x != nil {
		return x.RequesterProvider
	}
	return ""
}

func (x *MediaRequest_QueueItem) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MediaRequest_QueueItem) GetIsApproved() bool {
	if x != nil {
		return x.IsApproved
	}
	return false
}

func (x *MediaRequest_QueueItem) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

var File_pb_media_request_proto protoreflect.FileDescriptor

var file_pb_media_request_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f,
	0x6e, 0x6f, 0x1a, 0x12, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x50, 0x6c, 0x61, 0x79,
	0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f,
	0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x36, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x11, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x1a,
	0xed, 0x01, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12,
	0x20, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xa2, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x4d, 0x69, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x07, 0x4d, 0x69, 0x6e, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x69, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4d, 0x69, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x4d, 0x69, 0x6e, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x4d, 0x69, 0x6e, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6b, 0x75, 0x6c, 0x69, 0x6b, 0x30, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x64,
	0x6f, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_media_request_proto_rawDescOnce sync.Once
	file_pb_media_request_proto_rawDescData = file_pb_media_request_proto_rawDesc
)

func file_pb_media_request_proto_rawDescGZIP() []byte {
	file_pb_media_request_proto_rawDescOnce.Do(func() {
		file_pb_media_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_media_request_proto_rawDescData)
	})
	return file_pb_media_request_proto_rawDescData
}

var file_pb_media_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_media_request_proto_goTypes = []interface{}{
	(*MediaRequest)(nil),           // 0: stredono.MediaRequest
	(*MediaRequestSettings)(nil),   // 1: stredono.MediaRequestSettings
	(*MediaRequest_QueueItem)(nil), // 2: stredono.MediaRequest.QueueItem
	(Role)(0),                      // 3: stredono.Role
}
var file_pb_media_request_proto_depIdxs = []int32{
	1, // 0: stredono.MediaRequest.Settings:type_name -> stredono.MediaRequestSettings
	2, // 1: stredono.MediaRequest.Queue:type_name -> stredono.MediaRequest.QueueItem
	3, // 2: stredono.MediaRequestSettings.MinRole:type_name -> stredono.Role
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_media_request_proto_init() }
func file_pb_media_request_proto_init() {
	if File_pb_media_request_proto != nil {
		return
	}
	file_pb_chat_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_media_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaRequest); i {
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
		file_pb_media_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaRequestSettings); i {
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
		file_pb_media_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaRequest_QueueItem); i {
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
			RawDescriptor: file_pb_media_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_media_request_proto_goTypes,
		DependencyIndexes: file_pb_media_request_proto_depIdxs,
		MessageInfos:      file_pb_media_request_proto_msgTypes,
	}.Build()
	File_pb_media_request_proto = out.File
	file_pb_media_request_proto_rawDesc = nil
	file_pb_media_request_proto_goTypes = nil
	file_pb_media_request_proto_depIdxs = nil
}
