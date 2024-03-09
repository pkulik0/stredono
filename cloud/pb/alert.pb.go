// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pb/alert.proto

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

type AnimationType int32

const (
	AnimationType_PULSE              AnimationType = 0
	AnimationType_HEART_BEAT         AnimationType = 2
	AnimationType_SHAKE_VERTICALLY   AnimationType = 4
	AnimationType_SHAKE_HORIZONTALLY AnimationType = 5
	AnimationType_TADA               AnimationType = 8
	AnimationType_JELLO              AnimationType = 9
	AnimationType_BOUNCE             AnimationType = 10
)

// Enum value maps for AnimationType.
var (
	AnimationType_name = map[int32]string{
		0:  "PULSE",
		2:  "HEART_BEAT",
		4:  "SHAKE_VERTICALLY",
		5:  "SHAKE_HORIZONTALLY",
		8:  "TADA",
		9:  "JELLO",
		10: "BOUNCE",
	}
	AnimationType_value = map[string]int32{
		"PULSE":              0,
		"HEART_BEAT":         2,
		"SHAKE_VERTICALLY":   4,
		"SHAKE_HORIZONTALLY": 5,
		"TADA":               8,
		"JELLO":              9,
		"BOUNCE":             10,
	}
)

func (x AnimationType) Enum() *AnimationType {
	p := new(AnimationType)
	*p = x
	return p
}

func (x AnimationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnimationType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_alert_proto_enumTypes[0].Descriptor()
}

func (AnimationType) Type() protoreflect.EnumType {
	return &file_pb_alert_proto_enumTypes[0]
}

func (x AnimationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnimationType.Descriptor instead.
func (AnimationType) EnumDescriptor() ([]byte, []int) {
	return file_pb_alert_proto_rawDescGZIP(), []int{0}
}

type Alignment int32

const (
	Alignment_START   Alignment = 0
	Alignment_CENTER  Alignment = 1
	Alignment_END     Alignment = 2
	Alignment_JUSTIFY Alignment = 3
)

// Enum value maps for Alignment.
var (
	Alignment_name = map[int32]string{
		0: "START",
		1: "CENTER",
		2: "END",
		3: "JUSTIFY",
	}
	Alignment_value = map[string]int32{
		"START":   0,
		"CENTER":  1,
		"END":     2,
		"JUSTIFY": 3,
	}
)

func (x Alignment) Enum() *Alignment {
	p := new(Alignment)
	*p = x
	return p
}

func (x Alignment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Alignment) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_alert_proto_enumTypes[1].Descriptor()
}

func (Alignment) Type() protoreflect.EnumType {
	return &file_pb_alert_proto_enumTypes[1]
}

func (x Alignment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Alignment.Descriptor instead.
func (Alignment) EnumDescriptor() ([]byte, []int) {
	return file_pb_alert_proto_rawDescGZIP(), []int{1}
}

type Position int32

const (
	Position_TOP    Position = 0
	Position_LEFT   Position = 1
	Position_RIGHT  Position = 2
	Position_BOTTOM Position = 3
)

// Enum value maps for Position.
var (
	Position_name = map[int32]string{
		0: "TOP",
		1: "LEFT",
		2: "RIGHT",
		3: "BOTTOM",
	}
	Position_value = map[string]int32{
		"TOP":    0,
		"LEFT":   1,
		"RIGHT":  2,
		"BOTTOM": 3,
	}
)

func (x Position) Enum() *Position {
	p := new(Position)
	*p = x
	return p
}

func (x Position) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Position) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_alert_proto_enumTypes[2].Descriptor()
}

func (Position) Type() protoreflect.EnumType {
	return &file_pb_alert_proto_enumTypes[2]
}

func (x Position) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Position.Descriptor instead.
func (Position) EnumDescriptor() ([]byte, []int) {
	return file_pb_alert_proto_rawDescGZIP(), []int{2}
}

type Speed int32

const (
	Speed_OFF    Speed = 0
	Speed_SLOW   Speed = 1
	Speed_MEDIUM Speed = 2
	Speed_FAST   Speed = 3
	Speed_FASTER Speed = 4
)

// Enum value maps for Speed.
var (
	Speed_name = map[int32]string{
		0: "OFF",
		1: "SLOW",
		2: "MEDIUM",
		3: "FAST",
		4: "FASTER",
	}
	Speed_value = map[string]int32{
		"OFF":    0,
		"SLOW":   1,
		"MEDIUM": 2,
		"FAST":   3,
		"FASTER": 4,
	}
)

func (x Speed) Enum() *Speed {
	p := new(Speed)
	*p = x
	return p
}

func (x Speed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Speed) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_alert_proto_enumTypes[3].Descriptor()
}

func (Speed) Type() protoreflect.EnumType {
	return &file_pb_alert_proto_enumTypes[3]
}

func (x Speed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Speed.Descriptor instead.
func (Speed) EnumDescriptor() ([]byte, []int) {
	return file_pb_alert_proto_rawDescGZIP(), []int{3}
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string        `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	EventType      EventType     `protobuf:"varint,2,opt,name=EventType,proto3,enum=stredono.EventType" json:"EventType,omitempty"`
	Message        string        `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Min            float64       `protobuf:"fixed64,4,opt,name=Min,proto3" json:"Min,omitempty"`
	Max            *float64      `protobuf:"fixed64,5,opt,name=Max,proto3,oneof" json:"Max,omitempty"`
	GifUrl         string        `protobuf:"bytes,6,opt,name=GifUrl,proto3" json:"GifUrl,omitempty"`
	SoundUrl       string        `protobuf:"bytes,7,opt,name=SoundUrl,proto3" json:"SoundUrl,omitempty"`
	Animation      AnimationType `protobuf:"varint,8,opt,name=Animation,proto3,enum=stredono.AnimationType" json:"Animation,omitempty"`
	AnimationSpeed Speed         `protobuf:"varint,9,opt,name=AnimationSpeed,proto3,enum=stredono.Speed" json:"AnimationSpeed,omitempty"`
	TextColor      string        `protobuf:"bytes,10,opt,name=TextColor,proto3" json:"TextColor,omitempty"`
	AccentColor    string        `protobuf:"bytes,11,opt,name=AccentColor,proto3" json:"AccentColor,omitempty"`
	Alignment      Alignment     `protobuf:"varint,12,opt,name=Alignment,proto3,enum=stredono.Alignment" json:"Alignment,omitempty"`
	TextPosition   Position      `protobuf:"varint,13,opt,name=TextPosition,proto3,enum=stredono.Position" json:"TextPosition,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_pb_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_pb_alert_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Alert) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_TIP
}

func (x *Alert) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Alert) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Alert) GetMax() float64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

func (x *Alert) GetGifUrl() string {
	if x != nil {
		return x.GifUrl
	}
	return ""
}

func (x *Alert) GetSoundUrl() string {
	if x != nil {
		return x.SoundUrl
	}
	return ""
}

func (x *Alert) GetAnimation() AnimationType {
	if x != nil {
		return x.Animation
	}
	return AnimationType_PULSE
}

func (x *Alert) GetAnimationSpeed() Speed {
	if x != nil {
		return x.AnimationSpeed
	}
	return Speed_OFF
}

func (x *Alert) GetTextColor() string {
	if x != nil {
		return x.TextColor
	}
	return ""
}

func (x *Alert) GetAccentColor() string {
	if x != nil {
		return x.AccentColor
	}
	return ""
}

func (x *Alert) GetAlignment() Alignment {
	if x != nil {
		return x.Alignment
	}
	return Alignment_START
}

func (x *Alert) GetTextPosition() Position {
	if x != nil {
		return x.TextPosition
	}
	return Position_TOP
}

type UsersAlerts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*Alert `protobuf:"bytes,1,rep,name=Alerts,proto3" json:"Alerts,omitempty"`
}

func (x *UsersAlerts) Reset() {
	*x = UsersAlerts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersAlerts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersAlerts) ProtoMessage() {}

func (x *UsersAlerts) ProtoReflect() protoreflect.Message {
	mi := &file_pb_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersAlerts.ProtoReflect.Descriptor instead.
func (*UsersAlerts) Descriptor() ([]byte, []int) {
	return file_pb_alert_proto_rawDescGZIP(), []int{1}
}

func (x *UsersAlerts) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_pb_alert_proto protoreflect.FileDescriptor

var file_pb_alert_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x1a, 0x0e, 0x70, 0x62, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x62, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x03, 0x0a, 0x05, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f,
	0x6e, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x4d, 0x69, 0x6e, 0x12, 0x15, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x00, 0x52, 0x03, 0x4d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x69,
	0x66, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x69, 0x66, 0x55,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x35,
	0x0a, 0x09, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x41, 0x6e, 0x69, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x0e,
	0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x31,
	0x0a, 0x09, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x41, 0x6c, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x0c, 0x54, 0x65, 0x78, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f,
	0x6e, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x54, 0x65, 0x78,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x4d, 0x61,
	0x78, 0x22, 0x36, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x12, 0x27, 0x0a, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e, 0x6f, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2a, 0x79, 0x0a, 0x0d, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x55,
	0x4c, 0x53, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x42,
	0x45, 0x41, 0x54, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x48, 0x41, 0x4b, 0x45, 0x5f, 0x56,
	0x45, 0x52, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x4c, 0x59, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x48, 0x41, 0x4b, 0x45, 0x5f, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x4f, 0x4e, 0x54, 0x41, 0x4c, 0x4c,
	0x59, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x41, 0x44, 0x41, 0x10, 0x08, 0x12, 0x09, 0x0a,
	0x05, 0x4a, 0x45, 0x4c, 0x4c, 0x4f, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4f, 0x55, 0x4e,
	0x43, 0x45, 0x10, 0x0a, 0x2a, 0x38, 0x0a, 0x09, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x44, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4a, 0x55, 0x53, 0x54, 0x49, 0x46, 0x59, 0x10, 0x03, 0x2a, 0x34,
	0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x4f,
	0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4f, 0x54, 0x54,
	0x4f, 0x4d, 0x10, 0x03, 0x2a, 0x3c, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x07, 0x0a,
	0x03, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4c, 0x4f, 0x57, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x41, 0x53, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52,
	0x10, 0x04, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6b, 0x75, 0x6c, 0x69, 0x6b, 0x30, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x6e,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_alert_proto_rawDescOnce sync.Once
	file_pb_alert_proto_rawDescData = file_pb_alert_proto_rawDesc
)

func file_pb_alert_proto_rawDescGZIP() []byte {
	file_pb_alert_proto_rawDescOnce.Do(func() {
		file_pb_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_alert_proto_rawDescData)
	})
	return file_pb_alert_proto_rawDescData
}

var file_pb_alert_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_pb_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_alert_proto_goTypes = []interface{}{
	(AnimationType)(0),  // 0: stredono.AnimationType
	(Alignment)(0),      // 1: stredono.Alignment
	(Position)(0),       // 2: stredono.Position
	(Speed)(0),          // 3: stredono.Speed
	(*Alert)(nil),       // 4: stredono.Alert
	(*UsersAlerts)(nil), // 5: stredono.UsersAlerts
	(EventType)(0),      // 6: stredono.EventType
}
var file_pb_alert_proto_depIdxs = []int32{
	6, // 0: stredono.Alert.EventType:type_name -> stredono.EventType
	0, // 1: stredono.Alert.Animation:type_name -> stredono.AnimationType
	3, // 2: stredono.Alert.AnimationSpeed:type_name -> stredono.Speed
	1, // 3: stredono.Alert.Alignment:type_name -> stredono.Alignment
	2, // 4: stredono.Alert.TextPosition:type_name -> stredono.Position
	4, // 5: stredono.UsersAlerts.Alerts:type_name -> stredono.Alert
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_alert_proto_init() }
func file_pb_alert_proto_init() {
	if File_pb_alert_proto != nil {
		return
	}
	file_pb_enums_proto_init()
	file_pb_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_pb_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersAlerts); i {
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
	file_pb_alert_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_alert_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_alert_proto_goTypes,
		DependencyIndexes: file_pb_alert_proto_depIdxs,
		EnumInfos:         file_pb_alert_proto_enumTypes,
		MessageInfos:      file_pb_alert_proto_msgTypes,
	}.Build()
	File_pb_alert_proto = out.File
	file_pb_alert_proto_rawDesc = nil
	file_pb_alert_proto_goTypes = nil
	file_pb_alert_proto_depIdxs = nil
}