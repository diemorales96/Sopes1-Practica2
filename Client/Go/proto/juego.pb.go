// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: proto/juego.proto

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

type JuegoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId       int32  `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Players      int32  `protobuf:"varint,2,opt,name=players,proto3" json:"players,omitempty"`
	GameName     string `protobuf:"bytes,3,opt,name=game_name,json=gameName,proto3" json:"game_name,omitempty"`
	WinnerNumber string `protobuf:"bytes,4,opt,name=winner_number,json=winnerNumber,proto3" json:"winner_number,omitempty"`
	Queue        string `protobuf:"bytes,5,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *JuegoRequest) Reset() {
	*x = JuegoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_juego_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JuegoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegoRequest) ProtoMessage() {}

func (x *JuegoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juego_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegoRequest.ProtoReflect.Descriptor instead.
func (*JuegoRequest) Descriptor() ([]byte, []int) {
	return file_proto_juego_proto_rawDescGZIP(), []int{0}
}

func (x *JuegoRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *JuegoRequest) GetPlayers() int32 {
	if x != nil {
		return x.Players
	}
	return 0
}

func (x *JuegoRequest) GetGameName() string {
	if x != nil {
		return x.GameName
	}
	return ""
}

func (x *JuegoRequest) GetWinnerNumber() string {
	if x != nil {
		return x.WinnerNumber
	}
	return ""
}

func (x *JuegoRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type RequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestReply) Reset() {
	*x = RequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_juego_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReply) ProtoMessage() {}

func (x *RequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juego_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReply.ProtoReflect.Descriptor instead.
func (*RequestReply) Descriptor() ([]byte, []int) {
	return file_proto_juego_proto_rawDescGZIP(), []int{1}
}

func (x *RequestReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_juego_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juego_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_juego_proto_rawDescGZIP(), []int{2}
}

type JuegoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId       int32  `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Players      int32  `protobuf:"varint,2,opt,name=players,proto3" json:"players,omitempty"`
	GameName     string `protobuf:"bytes,3,opt,name=game_name,json=gameName,proto3" json:"game_name,omitempty"`
	WinnerNumber string `protobuf:"bytes,4,opt,name=winner_number,json=winnerNumber,proto3" json:"winner_number,omitempty"`
	Queue        string `protobuf:"bytes,5,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *JuegoReply) Reset() {
	*x = JuegoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_juego_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JuegoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegoReply) ProtoMessage() {}

func (x *JuegoReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juego_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegoReply.ProtoReflect.Descriptor instead.
func (*JuegoReply) Descriptor() ([]byte, []int) {
	return file_proto_juego_proto_rawDescGZIP(), []int{3}
}

func (x *JuegoReply) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *JuegoReply) GetPlayers() int32 {
	if x != nil {
		return x.Players
	}
	return 0
}

func (x *JuegoReply) GetGameName() string {
	if x != nil {
		return x.GameName
	}
	return ""
}

func (x *JuegoReply) GetWinnerNumber() string {
	if x != nil {
		return x.WinnerNumber
	}
	return ""
}

func (x *JuegoReply) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

var File_proto_juego_proto protoreflect.FileDescriptor

var file_proto_juego_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x0c, 0x4a,
	0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x97, 0x01, 0x0a, 0x0a, 0x4a, 0x75,
	0x65, 0x67, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x32, 0x6e, 0x0a, 0x06, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x12, 0x34, 0x0a,
	0x06, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x2e,
	0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6a,
	0x75, 0x65, 0x67, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x0c,
	0x2e, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x6a,
	0x75, 0x65, 0x67, 0x6f, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x65, 0x6d, 0x6f, 0x72, 0x61, 0x6c, 0x65, 0x73, 0x39, 0x36, 0x2f, 0x53,
	0x6f, 0x70, 0x65, 0x73, 0x31, 0x5f, 0x50, 0x72, 0x6f, 0x79, 0x65, 0x63, 0x74, 0x6f, 0x46, 0x32,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_juego_proto_rawDescOnce sync.Once
	file_proto_juego_proto_rawDescData = file_proto_juego_proto_rawDesc
)

func file_proto_juego_proto_rawDescGZIP() []byte {
	file_proto_juego_proto_rawDescOnce.Do(func() {
		file_proto_juego_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_juego_proto_rawDescData)
	})
	return file_proto_juego_proto_rawDescData
}

var file_proto_juego_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_juego_proto_goTypes = []interface{}{
	(*JuegoRequest)(nil), // 0: juego.JuegoRequest
	(*RequestReply)(nil), // 1: juego.RequestReply
	(*Empty)(nil),        // 2: juego.Empty
	(*JuegoReply)(nil),   // 3: juego.JuegoReply
}
var file_proto_juego_proto_depIdxs = []int32{
	0, // 0: juego.Juegos.AddLog:input_type -> juego.JuegoRequest
	2, // 1: juego.Juegos.getLogs:input_type -> juego.Empty
	1, // 2: juego.Juegos.AddLog:output_type -> juego.RequestReply
	3, // 3: juego.Juegos.getLogs:output_type -> juego.JuegoReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_juego_proto_init() }
func file_proto_juego_proto_init() {
	if File_proto_juego_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_juego_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JuegoRequest); i {
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
		file_proto_juego_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReply); i {
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
		file_proto_juego_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_juego_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JuegoReply); i {
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
			RawDescriptor: file_proto_juego_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_juego_proto_goTypes,
		DependencyIndexes: file_proto_juego_proto_depIdxs,
		MessageInfos:      file_proto_juego_proto_msgTypes,
	}.Build()
	File_proto_juego_proto = out.File
	file_proto_juego_proto_rawDesc = nil
	file_proto_juego_proto_goTypes = nil
	file_proto_juego_proto_depIdxs = nil
}
