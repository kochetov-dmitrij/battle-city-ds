// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: connection/pb/coms.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message_Direction int32

const (
	Message_STEADY Message_Direction = 0
	Message_UP     Message_Direction = 1
	Message_DOWN   Message_Direction = 2
	Message_LEFT   Message_Direction = 3
	Message_RIGHT  Message_Direction = 4
)

// Enum value maps for Message_Direction.
var (
	Message_Direction_name = map[int32]string{
		0: "STEADY",
		1: "UP",
		2: "DOWN",
		3: "LEFT",
		4: "RIGHT",
	}
	Message_Direction_value = map[string]int32{
		"STEADY": 0,
		"UP":     1,
		"DOWN":   2,
		"LEFT":   3,
		"RIGHT":  4,
	}
)

func (x Message_Direction) Enum() *Message_Direction {
	p := new(Message_Direction)
	*p = x
	return p
}

func (x Message_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_connection_pb_coms_proto_enumTypes[0].Descriptor()
}

func (Message_Direction) Type() protoreflect.EnumType {
	return &file_connection_pb_coms_proto_enumTypes[0]
}

func (x Message_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_Direction.Descriptor instead.
func (Message_Direction) EnumDescriptor() ([]byte, []int) {
	return file_connection_pb_coms_proto_rawDescGZIP(), []int{0, 0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TankState       uint32            `protobuf:"varint,1,opt,name=tankState,proto3" json:"tankState,omitempty"`
	TankPosition    *Message_Position `protobuf:"bytes,2,opt,name=tankPosition,proto3" json:"tankPosition,omitempty"`
	BulletState     uint32            `protobuf:"varint,3,opt,name=bulletState,proto3" json:"bulletState,omitempty"`
	BulletPosition  *Message_Position `protobuf:"bytes,4,opt,name=bulletPosition,proto3" json:"bulletPosition,omitempty"`
	BulletDirection Message_Direction `protobuf:"varint,5,opt,name=bulletDirection,proto3,enum=battle_city_ds.Message_Direction" json:"bulletDirection,omitempty"`
	TankDirection   Message_Direction `protobuf:"varint,6,opt,name=tankDirection,proto3,enum=battle_city_ds.Message_Direction" json:"tankDirection,omitempty"`
	Host            string            `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	AllPeers        []string          `protobuf:"bytes,8,rep,name=allPeers,proto3" json:"allPeers,omitempty"`
	LevelState      [][]byte          `protobuf:"bytes,9,rep,name=levelState,proto3" json:"levelState,omitempty"`
	Score           uint32            `protobuf:"varint,10,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_pb_coms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_connection_pb_coms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_connection_pb_coms_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetTankState() uint32 {
	if x != nil {
		return x.TankState
	}
	return 0
}

func (x *Message) GetTankPosition() *Message_Position {
	if x != nil {
		return x.TankPosition
	}
	return nil
}

func (x *Message) GetBulletState() uint32 {
	if x != nil {
		return x.BulletState
	}
	return 0
}

func (x *Message) GetBulletPosition() *Message_Position {
	if x != nil {
		return x.BulletPosition
	}
	return nil
}

func (x *Message) GetBulletDirection() Message_Direction {
	if x != nil {
		return x.BulletDirection
	}
	return Message_STEADY
}

func (x *Message) GetTankDirection() Message_Direction {
	if x != nil {
		return x.TankDirection
	}
	return Message_STEADY
}

func (x *Message) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Message) GetAllPeers() []string {
	if x != nil {
		return x.AllPeers
	}
	return nil
}

func (x *Message) GetLevelState() [][]byte {
	if x != nil {
		return x.LevelState
	}
	return nil
}

func (x *Message) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Message_Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y uint32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Message_Position) Reset() {
	*x = Message_Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_pb_coms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Position) ProtoMessage() {}

func (x *Message_Position) ProtoReflect() protoreflect.Message {
	mi := &file_connection_pb_coms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Position.ProtoReflect.Descriptor instead.
func (*Message_Position) Descriptor() ([]byte, []int) {
	return file_connection_pb_coms_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Message_Position) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Message_Position) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_connection_pb_coms_proto protoreflect.FileDescriptor

var file_connection_pb_coms_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x04, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x44, 0x0a, 0x0c, 0x74, 0x61, 0x6e, 0x6b, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x61, 0x6e, 0x6b, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x75, 0x6c, 0x6c, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x75,
	0x6c, 0x6c, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x62, 0x75, 0x6c,
	0x6c, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x64, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0f, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x47, 0x0a, 0x0d, 0x74, 0x61, 0x6e, 0x6b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x74, 0x61, 0x6e, 0x6b,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x6c, 0x6c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x6c, 0x6c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x26, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79, 0x22, 0x3e, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x45, 0x41, 0x44, 0x59, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x04, 0x32, 0x45, 0x0a, 0x04, 0x43, 0x6f, 0x6d, 0x73, 0x12,
	0x3d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x07,
	0x5a, 0x05, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connection_pb_coms_proto_rawDescOnce sync.Once
	file_connection_pb_coms_proto_rawDescData = file_connection_pb_coms_proto_rawDesc
)

func file_connection_pb_coms_proto_rawDescGZIP() []byte {
	file_connection_pb_coms_proto_rawDescOnce.Do(func() {
		file_connection_pb_coms_proto_rawDescData = protoimpl.X.CompressGZIP(file_connection_pb_coms_proto_rawDescData)
	})
	return file_connection_pb_coms_proto_rawDescData
}

var file_connection_pb_coms_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_connection_pb_coms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_connection_pb_coms_proto_goTypes = []interface{}{
	(Message_Direction)(0),   // 0: battle_city_ds.Message.Direction
	(*Message)(nil),          // 1: battle_city_ds.Message
	(*Message_Position)(nil), // 2: battle_city_ds.Message.Position
	(*empty.Empty)(nil),      // 3: google.protobuf.Empty
}
var file_connection_pb_coms_proto_depIdxs = []int32{
	2, // 0: battle_city_ds.Message.tankPosition:type_name -> battle_city_ds.Message.Position
	2, // 1: battle_city_ds.Message.bulletPosition:type_name -> battle_city_ds.Message.Position
	0, // 2: battle_city_ds.Message.bulletDirection:type_name -> battle_city_ds.Message.Direction
	0, // 3: battle_city_ds.Message.tankDirection:type_name -> battle_city_ds.Message.Direction
	1, // 4: battle_city_ds.Coms.addMessage:input_type -> battle_city_ds.Message
	3, // 5: battle_city_ds.Coms.addMessage:output_type -> google.protobuf.Empty
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_connection_pb_coms_proto_init() }
func file_connection_pb_coms_proto_init() {
	if File_connection_pb_coms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connection_pb_coms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_connection_pb_coms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_Position); i {
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
			RawDescriptor: file_connection_pb_coms_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connection_pb_coms_proto_goTypes,
		DependencyIndexes: file_connection_pb_coms_proto_depIdxs,
		EnumInfos:         file_connection_pb_coms_proto_enumTypes,
		MessageInfos:      file_connection_pb_coms_proto_msgTypes,
	}.Build()
	File_connection_pb_coms_proto = out.File
	file_connection_pb_coms_proto_rawDesc = nil
	file_connection_pb_coms_proto_goTypes = nil
	file_connection_pb_coms_proto_depIdxs = nil
}
