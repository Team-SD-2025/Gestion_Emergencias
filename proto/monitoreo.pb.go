// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/monitoreo.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActualizacionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClienteId     string                 `protobuf:"bytes,1,opt,name=cliente_id,json=clienteId,proto3" json:"cliente_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActualizacionRequest) Reset() {
	*x = ActualizacionRequest{}
	mi := &file_proto_monitoreo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActualizacionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActualizacionRequest) ProtoMessage() {}

func (x *ActualizacionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoreo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActualizacionRequest.ProtoReflect.Descriptor instead.
func (*ActualizacionRequest) Descriptor() ([]byte, []int) {
	return file_proto_monitoreo_proto_rawDescGZIP(), []int{0}
}

func (x *ActualizacionRequest) GetClienteId() string {
	if x != nil {
		return x.ClienteId
	}
	return ""
}

type ActualizacionReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DronId        string                 `protobuf:"bytes,1,opt,name=dron_id,json=dronId,proto3" json:"dron_id,omitempty"`
	Estado        string                 `protobuf:"bytes,2,opt,name=estado,proto3" json:"estado,omitempty"`
	Ubicacion     string                 `protobuf:"bytes,3,opt,name=ubicacion,proto3" json:"ubicacion,omitempty"`
	Timestamp     string                 `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActualizacionReply) Reset() {
	*x = ActualizacionReply{}
	mi := &file_proto_monitoreo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActualizacionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActualizacionReply) ProtoMessage() {}

func (x *ActualizacionReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoreo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActualizacionReply.ProtoReflect.Descriptor instead.
func (*ActualizacionReply) Descriptor() ([]byte, []int) {
	return file_proto_monitoreo_proto_rawDescGZIP(), []int{1}
}

func (x *ActualizacionReply) GetDronId() string {
	if x != nil {
		return x.DronId
	}
	return ""
}

func (x *ActualizacionReply) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *ActualizacionReply) GetUbicacion() string {
	if x != nil {
		return x.Ubicacion
	}
	return ""
}

func (x *ActualizacionReply) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_proto_monitoreo_proto protoreflect.FileDescriptor

const file_proto_monitoreo_proto_rawDesc = "" +
	"\n" +
	"\x15proto/monitoreo.proto\x12\tmonitoreo\"5\n" +
	"\x14ActualizacionRequest\x12\x1d\n" +
	"\n" +
	"cliente_id\x18\x01 \x01(\tR\tclienteId\"\x81\x01\n" +
	"\x12ActualizacionReply\x12\x17\n" +
	"\adron_id\x18\x01 \x01(\tR\x06dronId\x12\x16\n" +
	"\x06estado\x18\x02 \x01(\tR\x06estado\x12\x1c\n" +
	"\tubicacion\x18\x03 \x01(\tR\tubicacion\x12\x1c\n" +
	"\ttimestamp\x18\x04 \x01(\tR\ttimestamp2n\n" +
	"\x10MonitoreoService\x12Z\n" +
	"\x16RecibirActualizaciones\x12\x1f.monitoreo.ActualizacionRequest\x1a\x1d.monitoreo.ActualizacionReply0\x01B\tZ\a./protob\x06proto3"

var (
	file_proto_monitoreo_proto_rawDescOnce sync.Once
	file_proto_monitoreo_proto_rawDescData []byte
)

func file_proto_monitoreo_proto_rawDescGZIP() []byte {
	file_proto_monitoreo_proto_rawDescOnce.Do(func() {
		file_proto_monitoreo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_monitoreo_proto_rawDesc), len(file_proto_monitoreo_proto_rawDesc)))
	})
	return file_proto_monitoreo_proto_rawDescData
}

var file_proto_monitoreo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_monitoreo_proto_goTypes = []any{
	(*ActualizacionRequest)(nil), // 0: monitoreo.ActualizacionRequest
	(*ActualizacionReply)(nil),   // 1: monitoreo.ActualizacionReply
}
var file_proto_monitoreo_proto_depIdxs = []int32{
	0, // 0: monitoreo.MonitoreoService.RecibirActualizaciones:input_type -> monitoreo.ActualizacionRequest
	1, // 1: monitoreo.MonitoreoService.RecibirActualizaciones:output_type -> monitoreo.ActualizacionReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_monitoreo_proto_init() }
func file_proto_monitoreo_proto_init() {
	if File_proto_monitoreo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_monitoreo_proto_rawDesc), len(file_proto_monitoreo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_monitoreo_proto_goTypes,
		DependencyIndexes: file_proto_monitoreo_proto_depIdxs,
		MessageInfos:      file_proto_monitoreo_proto_msgTypes,
	}.Build()
	File_proto_monitoreo_proto = out.File
	file_proto_monitoreo_proto_goTypes = nil
	file_proto_monitoreo_proto_depIdxs = nil
}
