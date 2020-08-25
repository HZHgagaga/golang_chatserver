// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: chat.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type MSG int32

const (
	MSG_M_CreatePlayer  MSG = 0
	MSG_M_BroadcastChat MSG = 1
)

// Enum value maps for MSG.
var (
	MSG_name = map[int32]string{
		0: "M_CreatePlayer",
		1: "M_BroadcastChat",
	}
	MSG_value = map[string]int32{
		"M_CreatePlayer":  0,
		"M_BroadcastChat": 1,
	}
)

func (x MSG) Enum() *MSG {
	p := new(MSG)
	*p = x
	return p
}

func (x MSG) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[0].Descriptor()
}

func (MSG) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[0]
}

func (x MSG) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG.Descriptor instead.
func (MSG) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type CreatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreatePlayer) Reset() {
	*x = CreatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayer) ProtoMessage() {}

func (x *CreatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayer.ProtoReflect.Descriptor instead.
func (*CreatePlayer) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BroadcastChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Chatdata string `protobuf:"bytes,2,opt,name=chatdata,proto3" json:"chatdata,omitempty"`
}

func (x *BroadcastChat) Reset() {
	*x = BroadcastChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastChat) ProtoMessage() {}

func (x *BroadcastChat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastChat.ProtoReflect.Descriptor instead.
func (*BroadcastChat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *BroadcastChat) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *BroadcastChat) GetChatdata() string {
	if x != nil {
		return x.Chatdata
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x22, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2a, 0x2e, 0x0a, 0x03, 0x4d, 0x53, 0x47, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x5f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x4d, 0x5f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_proto_goTypes = []interface{}{
	(MSG)(0),              // 0: pb.MSG
	(*CreatePlayer)(nil),  // 1: pb.CreatePlayer
	(*BroadcastChat)(nil), // 2: pb.BroadcastChat
}
var file_chat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayer); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastChat); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		EnumInfos:         file_chat_proto_enumTypes,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
