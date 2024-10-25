// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.4
// source: log.proto

package eslog

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

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId         string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Label         string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Level         string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Content       string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Caller        string `protobuf:"bytes,5,opt,name=caller,proto3" json:"caller,omitempty"`
	Datatime      string `protobuf:"bytes,6,opt,name=datatime,proto3" json:"datatime,omitempty"`
	RequestId     string `protobuf:"bytes,7,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	EsIndexPrefix string `protobuf:"bytes,8,opt,name=es_index_prefix,json=esIndexPrefix,proto3" json:"es_index_prefix,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	mi := &file_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *LogRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *LogRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *LogRequest) GetCaller() string {
	if x != nil {
		return x.Caller
	}
	return ""
}

func (x *LogRequest) GetDatatime() string {
	if x != nil {
		return x.Datatime
	}
	return ""
}

func (x *LogRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *LogRequest) GetEsIndexPrefix() string {
	if x != nil {
		return x.EsIndexPrefix
	}
	return ""
}

type LogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LogReply) Reset() {
	*x = LogReply{}
	mi := &file_log_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogReply) ProtoMessage() {}

func (x *LogReply) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogReply.ProtoReflect.Descriptor instead.
func (*LogReply) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LogReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LogReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_log_proto protoreflect.FileDescriptor

var file_log_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x73, 0x6c,
	0x6f, 0x67, 0x22, 0xe4, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x73, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0xee, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x2d, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x12, 0x11, 0x2e, 0x65, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11,
	0x2e, 0x65, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x57, 0x61, 0x72, 0x6e, 0x12, 0x11, 0x2e, 0x65,
	0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x65, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x2e, 0x65, 0x73,
	0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x65, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x2d, 0x0a, 0x05, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x12, 0x11, 0x2e, 0x65, 0x73, 0x6c,
	0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x65, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x65, 0x73, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_log_proto_rawDescOnce sync.Once
	file_log_proto_rawDescData = file_log_proto_rawDesc
)

func file_log_proto_rawDescGZIP() []byte {
	file_log_proto_rawDescOnce.Do(func() {
		file_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_proto_rawDescData)
	})
	return file_log_proto_rawDescData
}

var file_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_log_proto_goTypes = []any{
	(*LogRequest)(nil), // 0: eslog.LogRequest
	(*LogReply)(nil),   // 1: eslog.LogReply
}
var file_log_proto_depIdxs = []int32{
	0, // 0: eslog.Log.Debug:input_type -> eslog.LogRequest
	0, // 1: eslog.Log.Info:input_type -> eslog.LogRequest
	0, // 2: eslog.Log.Warn:input_type -> eslog.LogRequest
	0, // 3: eslog.Log.Error:input_type -> eslog.LogRequest
	0, // 4: eslog.Log.Fatal:input_type -> eslog.LogRequest
	1, // 5: eslog.Log.Debug:output_type -> eslog.LogReply
	1, // 6: eslog.Log.Info:output_type -> eslog.LogReply
	1, // 7: eslog.Log.Warn:output_type -> eslog.LogReply
	1, // 8: eslog.Log.Error:output_type -> eslog.LogReply
	1, // 9: eslog.Log.Fatal:output_type -> eslog.LogReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_log_proto_init() }
func file_log_proto_init() {
	if File_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_proto_goTypes,
		DependencyIndexes: file_log_proto_depIdxs,
		MessageInfos:      file_log_proto_msgTypes,
	}.Build()
	File_log_proto = out.File
	file_log_proto_rawDesc = nil
	file_log_proto_goTypes = nil
	file_log_proto_depIdxs = nil
}
