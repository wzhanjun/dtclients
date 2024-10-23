// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.4
// source: captcha.proto

package captcha

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

type SendCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Driver  *string `protobuf:"bytes,2,opt,name=driver,proto3,oneof" json:"driver,omitempty"`
	Source  *string `protobuf:"bytes,3,opt,name=source,proto3,oneof" json:"source,omitempty"`
}

func (x *SendCaptchaRequest) Reset() {
	*x = SendCaptchaRequest{}
	mi := &file_captcha_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendCaptchaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCaptchaRequest) ProtoMessage() {}

func (x *SendCaptchaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCaptchaRequest.ProtoReflect.Descriptor instead.
func (*SendCaptchaRequest) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *SendCaptchaRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SendCaptchaRequest) GetDriver() string {
	if x != nil && x.Driver != nil {
		return *x.Driver
	}
	return ""
}

func (x *SendCaptchaRequest) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

type SendCaptchaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendCaptchaResponse) Reset() {
	*x = SendCaptchaResponse{}
	mi := &file_captcha_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendCaptchaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCaptchaResponse) ProtoMessage() {}

func (x *SendCaptchaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCaptchaResponse.ProtoReflect.Descriptor instead.
func (*SendCaptchaResponse) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{1}
}

func (x *SendCaptchaResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VerifyCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Code    string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Driver  *string `protobuf:"bytes,3,opt,name=driver,proto3,oneof" json:"driver,omitempty"`
	Source  *string `protobuf:"bytes,4,opt,name=source,proto3,oneof" json:"source,omitempty"`
}

func (x *VerifyCaptchaRequest) Reset() {
	*x = VerifyCaptchaRequest{}
	mi := &file_captcha_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyCaptchaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCaptchaRequest) ProtoMessage() {}

func (x *VerifyCaptchaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCaptchaRequest.ProtoReflect.Descriptor instead.
func (*VerifyCaptchaRequest) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyCaptchaRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *VerifyCaptchaRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyCaptchaRequest) GetDriver() string {
	if x != nil && x.Driver != nil {
		return *x.Driver
	}
	return ""
}

func (x *VerifyCaptchaRequest) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

type VerifyCaptchaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *VerifyCaptchaResponse) Reset() {
	*x = VerifyCaptchaResponse{}
	mi := &file_captcha_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyCaptchaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCaptchaResponse) ProtoMessage() {}

func (x *VerifyCaptchaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCaptchaResponse.ProtoReflect.Descriptor instead.
func (*VerifyCaptchaResponse) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyCaptchaResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_captcha_proto protoreflect.FileDescriptor

var file_captcha_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x7e, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1b, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x31, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0x95, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12,
	0x41, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x2e, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_captcha_proto_rawDescOnce sync.Once
	file_captcha_proto_rawDescData = file_captcha_proto_rawDesc
)

func file_captcha_proto_rawDescGZIP() []byte {
	file_captcha_proto_rawDescOnce.Do(func() {
		file_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_captcha_proto_rawDescData)
	})
	return file_captcha_proto_rawDescData
}

var file_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_captcha_proto_goTypes = []any{
	(*SendCaptchaRequest)(nil),    // 0: captcha.SendCaptchaRequest
	(*SendCaptchaResponse)(nil),   // 1: captcha.SendCaptchaResponse
	(*VerifyCaptchaRequest)(nil),  // 2: captcha.VerifyCaptchaRequest
	(*VerifyCaptchaResponse)(nil), // 3: captcha.VerifyCaptchaResponse
}
var file_captcha_proto_depIdxs = []int32{
	0, // 0: captcha.Captcha.Send:input_type -> captcha.SendCaptchaRequest
	2, // 1: captcha.Captcha.Verify:input_type -> captcha.VerifyCaptchaRequest
	1, // 2: captcha.Captcha.Send:output_type -> captcha.SendCaptchaResponse
	3, // 3: captcha.Captcha.Verify:output_type -> captcha.VerifyCaptchaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_captcha_proto_init() }
func file_captcha_proto_init() {
	if File_captcha_proto != nil {
		return
	}
	file_captcha_proto_msgTypes[0].OneofWrappers = []any{}
	file_captcha_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_captcha_proto_goTypes,
		DependencyIndexes: file_captcha_proto_depIdxs,
		MessageInfos:      file_captcha_proto_msgTypes,
	}.Build()
	File_captcha_proto = out.File
	file_captcha_proto_rawDesc = nil
	file_captcha_proto_goTypes = nil
	file_captcha_proto_depIdxs = nil
}
