// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: protos/vrfy.proto

package generated

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

type VrfyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VrfyRequest) Reset() {
	*x = VrfyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vrfy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VrfyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VrfyRequest) ProtoMessage() {}

func (x *VrfyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vrfy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VrfyRequest.ProtoReflect.Descriptor instead.
func (*VrfyRequest) Descriptor() ([]byte, []int) {
	return file_protos_vrfy_proto_rawDescGZIP(), []int{0}
}

func (x *VrfyRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VrfyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ErrorMessage *string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
}

func (x *VrfyResponse) Reset() {
	*x = VrfyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vrfy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VrfyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VrfyResponse) ProtoMessage() {}

func (x *VrfyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vrfy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VrfyResponse.ProtoReflect.Descriptor instead.
func (*VrfyResponse) Descriptor() ([]byte, []int) {
	return file_protos_vrfy_proto_rawDescGZIP(), []int{1}
}

func (x *VrfyResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *VrfyResponse) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

var File_protos_vrfy_proto protoreflect.FileDescriptor

var file_protos_vrfy_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x72, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0b, 0x56, 0x72, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6b, 0x0a, 0x0c, 0x56, 0x72, 0x66, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x36, 0x0a, 0x0b, 0x56, 0x72, 0x66, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0c,
	0x2e, 0x56, 0x72, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x56,
	0x72, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a,
	0x14, 0x2e, 0x68, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_vrfy_proto_rawDescOnce sync.Once
	file_protos_vrfy_proto_rawDescData = file_protos_vrfy_proto_rawDesc
)

func file_protos_vrfy_proto_rawDescGZIP() []byte {
	file_protos_vrfy_proto_rawDescOnce.Do(func() {
		file_protos_vrfy_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vrfy_proto_rawDescData)
	})
	return file_protos_vrfy_proto_rawDescData
}

var file_protos_vrfy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_vrfy_proto_goTypes = []interface{}{
	(*VrfyRequest)(nil),  // 0: VrfyRequest
	(*VrfyResponse)(nil), // 1: VrfyResponse
}
var file_protos_vrfy_proto_depIdxs = []int32{
	0, // 0: VrfyService.Verify:input_type -> VrfyRequest
	1, // 1: VrfyService.Verify:output_type -> VrfyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_vrfy_proto_init() }
func file_protos_vrfy_proto_init() {
	if File_protos_vrfy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_vrfy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VrfyRequest); i {
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
		file_protos_vrfy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VrfyResponse); i {
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
	file_protos_vrfy_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_vrfy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_vrfy_proto_goTypes,
		DependencyIndexes: file_protos_vrfy_proto_depIdxs,
		MessageInfos:      file_protos_vrfy_proto_msgTypes,
	}.Build()
	File_protos_vrfy_proto = out.File
	file_protos_vrfy_proto_rawDesc = nil
	file_protos_vrfy_proto_goTypes = nil
	file_protos_vrfy_proto_depIdxs = nil
}
