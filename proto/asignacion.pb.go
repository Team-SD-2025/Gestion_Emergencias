// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/asignacion.proto

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

type EmergenciaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nombre        string                 `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Latitude      float32                `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float32                `protobuf:"fixed32,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Magnitud      int32                  `protobuf:"varint,4,opt,name=magnitud,proto3" json:"magnitud,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmergenciaRequest) Reset() {
	*x = EmergenciaRequest{}
	mi := &file_proto_asignacion_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmergenciaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmergenciaRequest) ProtoMessage() {}

func (x *EmergenciaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asignacion_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmergenciaRequest.ProtoReflect.Descriptor instead.
func (*EmergenciaRequest) Descriptor() ([]byte, []int) {
	return file_proto_asignacion_proto_rawDescGZIP(), []int{0}
}

func (x *EmergenciaRequest) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *EmergenciaRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *EmergenciaRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *EmergenciaRequest) GetMagnitud() int32 {
	if x != nil {
		return x.Magnitud
	}
	return 0
}

type EmergenciaReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Estado        string                 `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
	DronAsignado  string                 `protobuf:"bytes,2,opt,name=dronAsignado,proto3" json:"dronAsignado,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmergenciaReply) Reset() {
	*x = EmergenciaReply{}
	mi := &file_proto_asignacion_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmergenciaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmergenciaReply) ProtoMessage() {}

func (x *EmergenciaReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asignacion_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmergenciaReply.ProtoReflect.Descriptor instead.
func (*EmergenciaReply) Descriptor() ([]byte, []int) {
	return file_proto_asignacion_proto_rawDescGZIP(), []int{1}
}

func (x *EmergenciaReply) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *EmergenciaReply) GetDronAsignado() string {
	if x != nil {
		return x.DronAsignado
	}
	return ""
}

var File_proto_asignacion_proto protoreflect.FileDescriptor

const file_proto_asignacion_proto_rawDesc = "" +
	"\n" +
	"\x16proto/asignacion.proto\x12\n" +
	"asignacion\"\x81\x01\n" +
	"\x11EmergenciaRequest\x12\x16\n" +
	"\x06nombre\x18\x01 \x01(\tR\x06nombre\x12\x1a\n" +
	"\blatitude\x18\x02 \x01(\x02R\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x03 \x01(\x02R\tlongitude\x12\x1a\n" +
	"\bmagnitud\x18\x04 \x01(\x05R\bmagnitud\"M\n" +
	"\x0fEmergenciaReply\x12\x16\n" +
	"\x06estado\x18\x01 \x01(\tR\x06estado\x12\"\n" +
	"\fdronAsignado\x18\x02 \x01(\tR\fdronAsignado2d\n" +
	"\x11AsignacionService\x12O\n" +
	"\x11AsignarEmergencia\x12\x1d.asignacion.EmergenciaRequest\x1a\x1b.asignacion.EmergenciaReplyB\tZ\a./protob\x06proto3"

var (
	file_proto_asignacion_proto_rawDescOnce sync.Once
	file_proto_asignacion_proto_rawDescData []byte
)

func file_proto_asignacion_proto_rawDescGZIP() []byte {
	file_proto_asignacion_proto_rawDescOnce.Do(func() {
		file_proto_asignacion_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_asignacion_proto_rawDesc), len(file_proto_asignacion_proto_rawDesc)))
	})
	return file_proto_asignacion_proto_rawDescData
}

var file_proto_asignacion_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_asignacion_proto_goTypes = []any{
	(*EmergenciaRequest)(nil), // 0: asignacion.EmergenciaRequest
	(*EmergenciaReply)(nil),   // 1: asignacion.EmergenciaReply
}
var file_proto_asignacion_proto_depIdxs = []int32{
	0, // 0: asignacion.AsignacionService.AsignarEmergencia:input_type -> asignacion.EmergenciaRequest
	1, // 1: asignacion.AsignacionService.AsignarEmergencia:output_type -> asignacion.EmergenciaReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_asignacion_proto_init() }
func file_proto_asignacion_proto_init() {
	if File_proto_asignacion_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_asignacion_proto_rawDesc), len(file_proto_asignacion_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_asignacion_proto_goTypes,
		DependencyIndexes: file_proto_asignacion_proto_depIdxs,
		MessageInfos:      file_proto_asignacion_proto_msgTypes,
	}.Build()
	File_proto_asignacion_proto = out.File
	file_proto_asignacion_proto_goTypes = nil
	file_proto_asignacion_proto_depIdxs = nil
}
