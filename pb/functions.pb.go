// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pb/functions.proto

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

type SendDonateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    string  `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Email     string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Message   string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Amount    float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency  string  `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Recipient string  `protobuf:"bytes,6,opt,name=recipient,proto3" json:"recipient,omitempty"`
}

func (x *SendDonateRequest) Reset() {
	*x = SendDonateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_functions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDonateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDonateRequest) ProtoMessage() {}

func (x *SendDonateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_functions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDonateRequest.ProtoReflect.Descriptor instead.
func (*SendDonateRequest) Descriptor() ([]byte, []int) {
	return file_pb_functions_proto_rawDescGZIP(), []int{0}
}

func (x *SendDonateRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *SendDonateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendDonateRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendDonateRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SendDonateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SendDonateRequest) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

type SendDonateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirectUrl,proto3" json:"redirectUrl,omitempty"`
}

func (x *SendDonateResponse) Reset() {
	*x = SendDonateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_functions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDonateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDonateResponse) ProtoMessage() {}

func (x *SendDonateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_functions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDonateResponse.ProtoReflect.Descriptor instead.
func (*SendDonateResponse) Descriptor() ([]byte, []int) {
	return file_pb_functions_proto_rawDescGZIP(), []int{1}
}

func (x *SendDonateResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_functions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_functions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_pb_functions_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pb_functions_proto protoreflect.FileDescriptor

var file_pb_functions_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x22, 0xad,
	0x01, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x36,
	0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x25, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x20, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6b, 0x75, 0x6c,
	0x69, 0x6b, 0x30, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_functions_proto_rawDescOnce sync.Once
	file_pb_functions_proto_rawDescData = file_pb_functions_proto_rawDesc
)

func file_pb_functions_proto_rawDescGZIP() []byte {
	file_pb_functions_proto_rawDescOnce.Do(func() {
		file_pb_functions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_functions_proto_rawDescData)
	})
	return file_pb_functions_proto_rawDescData
}

var file_pb_functions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_functions_proto_goTypes = []interface{}{
	(*SendDonateRequest)(nil),  // 0: stredono.SendDonateRequest
	(*SendDonateResponse)(nil), // 1: stredono.SendDonateResponse
	(*ErrorResponse)(nil),      // 2: stredono.ErrorResponse
}
var file_pb_functions_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_functions_proto_init() }
func file_pb_functions_proto_init() {
	if File_pb_functions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_functions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDonateRequest); i {
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
		file_pb_functions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDonateResponse); i {
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
		file_pb_functions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_pb_functions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_functions_proto_goTypes,
		DependencyIndexes: file_pb_functions_proto_depIdxs,
		MessageInfos:      file_pb_functions_proto_msgTypes,
	}.Build()
	File_pb_functions_proto = out.File
	file_pb_functions_proto_rawDesc = nil
	file_pb_functions_proto_goTypes = nil
	file_pb_functions_proto_depIdxs = nil
}
