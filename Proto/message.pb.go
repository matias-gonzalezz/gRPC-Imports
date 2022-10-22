// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: Proto/message.proto

package __

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[0]
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
	return file_Proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipo string `protobuf:"bytes,1,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Data) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Data) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ConsultaTipo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipo string `protobuf:"bytes,1,opt,name=tipo,proto3" json:"tipo,omitempty"`
}

func (x *ConsultaTipo) Reset() {
	*x = ConsultaTipo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultaTipo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultaTipo) ProtoMessage() {}

func (x *ConsultaTipo) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultaTipo.ProtoReflect.Descriptor instead.
func (*ConsultaTipo) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *ConsultaTipo) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

type DataSolicitada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataSolicitada) Reset() {
	*x = DataSolicitada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSolicitada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSolicitada) ProtoMessage() {}

func (x *DataSolicitada) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSolicitada.ProtoReflect.Descriptor instead.
func (*DataSolicitada) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{3}
}

func (x *DataSolicitada) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DataSolicitada) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DataID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DataID) Reset() {
	*x = DataID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataID) ProtoMessage() {}

func (x *DataID) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataID.ProtoReflect.Descriptor instead.
func (*DataID) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{4}
}

func (x *DataID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_Proto_message_proto protoreflect.FileDescriptor

var file_Proto_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x23, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x54, 0x69, 0x70, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x70, 0x6f, 0x22, 0x34, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x61, 0x64, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x18, 0x0a, 0x06, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x76, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x61, 0x6d,
	0x62, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x0d, 0x53,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0a, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x44, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x3e, 0x0a, 0x12,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x43, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x45, 0x6e, 0x76, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0x53, 0x0a, 0x12,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x52, 0x65, 0x62, 0x65, 0x6c,
	0x64, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x74, 0x61, 0x54, 0x69, 0x70, 0x6f, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x64, 0x61, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_message_proto_rawDescOnce sync.Once
	file_Proto_message_proto_rawDescData = file_Proto_message_proto_rawDesc
)

func file_Proto_message_proto_rawDescGZIP() []byte {
	file_Proto_message_proto_rawDescOnce.Do(func() {
		file_Proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_message_proto_rawDescData)
	})
	return file_Proto_message_proto_rawDescData
}

var file_Proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Proto_message_proto_goTypes = []interface{}{
	(*Message)(nil),        // 0: grpc.Message
	(*Data)(nil),           // 1: grpc.Data
	(*ConsultaTipo)(nil),   // 2: grpc.ConsultaTipo
	(*DataSolicitada)(nil), // 3: grpc.DataSolicitada
	(*DataID)(nil),         // 4: grpc.DataID
}
var file_Proto_message_proto_depIdxs = []int32{
	0, // 0: grpc.IntercambioDataNode.GuardarData:input_type -> grpc.Message
	1, // 1: grpc.IntercambioDataNode.SolicitarData:input_type -> grpc.Data
	1, // 2: grpc.IntercambioCombine.EnvioData:input_type -> grpc.Data
	2, // 3: grpc.IntercambioRebelde.ConsultarData:input_type -> grpc.ConsultaTipo
	1, // 4: grpc.IntercambioDataNode.GuardarData:output_type -> grpc.Data
	4, // 5: grpc.IntercambioDataNode.SolicitarData:output_type -> grpc.DataID
	0, // 6: grpc.IntercambioCombine.EnvioData:output_type -> grpc.Message
	3, // 7: grpc.IntercambioRebelde.ConsultarData:output_type -> grpc.DataSolicitada
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Proto_message_proto_init() }
func file_Proto_message_proto_init() {
	if File_Proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_Proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultaTipo); i {
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
		file_Proto_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSolicitada); i {
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
		file_Proto_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataID); i {
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
			RawDescriptor: file_Proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_Proto_message_proto_goTypes,
		DependencyIndexes: file_Proto_message_proto_depIdxs,
		MessageInfos:      file_Proto_message_proto_msgTypes,
	}.Build()
	File_Proto_message_proto = out.File
	file_Proto_message_proto_rawDesc = nil
	file_Proto_message_proto_goTypes = nil
	file_Proto_message_proto_depIdxs = nil
}
