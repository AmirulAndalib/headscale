// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: headscale/v1/policy.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SetPolicyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policy        string                 `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPolicyRequest) Reset() {
	*x = SetPolicyRequest{}
	mi := &file_headscale_v1_policy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPolicyRequest) ProtoMessage() {}

func (x *SetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_policy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPolicyRequest.ProtoReflect.Descriptor instead.
func (*SetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *SetPolicyRequest) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

type SetPolicyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policy        string                 `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPolicyResponse) Reset() {
	*x = SetPolicyResponse{}
	mi := &file_headscale_v1_policy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPolicyResponse) ProtoMessage() {}

func (x *SetPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_policy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPolicyResponse.ProtoReflect.Descriptor instead.
func (*SetPolicyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *SetPolicyResponse) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *SetPolicyResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetPolicyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPolicyRequest) Reset() {
	*x = GetPolicyRequest{}
	mi := &file_headscale_v1_policy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyRequest) ProtoMessage() {}

func (x *GetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_policy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_policy_proto_rawDescGZIP(), []int{2}
}

type GetPolicyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policy        string                 `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPolicyResponse) Reset() {
	*x = GetPolicyResponse{}
	mi := &file_headscale_v1_policy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyResponse) ProtoMessage() {}

func (x *GetPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_policy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyResponse.ProtoReflect.Descriptor instead.
func (*GetPolicyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_policy_proto_rawDescGZIP(), []int{3}
}

func (x *GetPolicyResponse) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *GetPolicyResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_headscale_v1_policy_proto protoreflect.FileDescriptor

var file_headscale_v1_policy_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x66, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x66, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x66, 0x6f, 0x6e,
	0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_headscale_v1_policy_proto_rawDescOnce sync.Once
	file_headscale_v1_policy_proto_rawDescData []byte
)

func file_headscale_v1_policy_proto_rawDescGZIP() []byte {
	file_headscale_v1_policy_proto_rawDescOnce.Do(func() {
		file_headscale_v1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_headscale_v1_policy_proto_rawDesc), len(file_headscale_v1_policy_proto_rawDesc)))
	})
	return file_headscale_v1_policy_proto_rawDescData
}

var file_headscale_v1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_headscale_v1_policy_proto_goTypes = []any{
	(*SetPolicyRequest)(nil),      // 0: headscale.v1.SetPolicyRequest
	(*SetPolicyResponse)(nil),     // 1: headscale.v1.SetPolicyResponse
	(*GetPolicyRequest)(nil),      // 2: headscale.v1.GetPolicyRequest
	(*GetPolicyResponse)(nil),     // 3: headscale.v1.GetPolicyResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_headscale_v1_policy_proto_depIdxs = []int32{
	4, // 0: headscale.v1.SetPolicyResponse.updated_at:type_name -> google.protobuf.Timestamp
	4, // 1: headscale.v1.GetPolicyResponse.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_headscale_v1_policy_proto_init() }
func file_headscale_v1_policy_proto_init() {
	if File_headscale_v1_policy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_headscale_v1_policy_proto_rawDesc), len(file_headscale_v1_policy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_headscale_v1_policy_proto_goTypes,
		DependencyIndexes: file_headscale_v1_policy_proto_depIdxs,
		MessageInfos:      file_headscale_v1_policy_proto_msgTypes,
	}.Build()
	File_headscale_v1_policy_proto = out.File
	file_headscale_v1_policy_proto_goTypes = nil
	file_headscale_v1_policy_proto_depIdxs = nil
}
