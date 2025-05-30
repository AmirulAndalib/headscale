// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: headscale/v1/preauthkey.proto

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

type PreAuthKey struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Id            uint64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Key           string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Reusable      bool                   `protobuf:"varint,4,opt,name=reusable,proto3" json:"reusable,omitempty"`
	Ephemeral     bool                   `protobuf:"varint,5,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	Used          bool                   `protobuf:"varint,6,opt,name=used,proto3" json:"used,omitempty"`
	Expiration    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expiration,proto3" json:"expiration,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AclTags       []string               `protobuf:"bytes,9,rep,name=acl_tags,json=aclTags,proto3" json:"acl_tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PreAuthKey) Reset() {
	*x = PreAuthKey{}
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PreAuthKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthKey) ProtoMessage() {}

func (x *PreAuthKey) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthKey.ProtoReflect.Descriptor instead.
func (*PreAuthKey) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{0}
}

func (x *PreAuthKey) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PreAuthKey) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PreAuthKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PreAuthKey) GetReusable() bool {
	if x != nil {
		return x.Reusable
	}
	return false
}

func (x *PreAuthKey) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

func (x *PreAuthKey) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *PreAuthKey) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *PreAuthKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PreAuthKey) GetAclTags() []string {
	if x != nil {
		return x.AclTags
	}
	return nil
}

type CreatePreAuthKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          uint64                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Reusable      bool                   `protobuf:"varint,2,opt,name=reusable,proto3" json:"reusable,omitempty"`
	Ephemeral     bool                   `protobuf:"varint,3,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	Expiration    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	AclTags       []string               `protobuf:"bytes,5,rep,name=acl_tags,json=aclTags,proto3" json:"acl_tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePreAuthKeyRequest) Reset() {
	*x = CreatePreAuthKeyRequest{}
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePreAuthKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePreAuthKeyRequest) ProtoMessage() {}

func (x *CreatePreAuthKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePreAuthKeyRequest.ProtoReflect.Descriptor instead.
func (*CreatePreAuthKeyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePreAuthKeyRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreatePreAuthKeyRequest) GetReusable() bool {
	if x != nil {
		return x.Reusable
	}
	return false
}

func (x *CreatePreAuthKeyRequest) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

func (x *CreatePreAuthKeyRequest) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *CreatePreAuthKeyRequest) GetAclTags() []string {
	if x != nil {
		return x.AclTags
	}
	return nil
}

type CreatePreAuthKeyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PreAuthKey    *PreAuthKey            `protobuf:"bytes,1,opt,name=pre_auth_key,json=preAuthKey,proto3" json:"pre_auth_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePreAuthKeyResponse) Reset() {
	*x = CreatePreAuthKeyResponse{}
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePreAuthKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePreAuthKeyResponse) ProtoMessage() {}

func (x *CreatePreAuthKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePreAuthKeyResponse.ProtoReflect.Descriptor instead.
func (*CreatePreAuthKeyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePreAuthKeyResponse) GetPreAuthKey() *PreAuthKey {
	if x != nil {
		return x.PreAuthKey
	}
	return nil
}

type ExpirePreAuthKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          uint64                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpirePreAuthKeyRequest) Reset() {
	*x = ExpirePreAuthKeyRequest{}
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpirePreAuthKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpirePreAuthKeyRequest) ProtoMessage() {}

func (x *ExpirePreAuthKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpirePreAuthKeyRequest.ProtoReflect.Descriptor instead.
func (*ExpirePreAuthKeyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{3}
}

func (x *ExpirePreAuthKeyRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *ExpirePreAuthKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ExpirePreAuthKeyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpirePreAuthKeyResponse) Reset() {
	*x = ExpirePreAuthKeyResponse{}
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpirePreAuthKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpirePreAuthKeyResponse) ProtoMessage() {}

func (x *ExpirePreAuthKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpirePreAuthKeyResponse.ProtoReflect.Descriptor instead.
func (*ExpirePreAuthKeyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{4}
}

type ListPreAuthKeysRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          uint64                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPreAuthKeysRequest) Reset() {
	*x = ListPreAuthKeysRequest{}
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPreAuthKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPreAuthKeysRequest) ProtoMessage() {}

func (x *ListPreAuthKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPreAuthKeysRequest.ProtoReflect.Descriptor instead.
func (*ListPreAuthKeysRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{5}
}

func (x *ListPreAuthKeysRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

type ListPreAuthKeysResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PreAuthKeys   []*PreAuthKey          `protobuf:"bytes,1,rep,name=pre_auth_keys,json=preAuthKeys,proto3" json:"pre_auth_keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPreAuthKeysResponse) Reset() {
	*x = ListPreAuthKeysResponse{}
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPreAuthKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPreAuthKeysResponse) ProtoMessage() {}

func (x *ListPreAuthKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPreAuthKeysResponse.ProtoReflect.Descriptor instead.
func (*ListPreAuthKeysResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{6}
}

func (x *ListPreAuthKeysResponse) GetPreAuthKeys() []*PreAuthKey {
	if x != nil {
		return x.PreAuthKeys
	}
	return nil
}

var File_headscale_v1_preauthkey_proto protoreflect.FileDescriptor

const file_headscale_v1_preauthkey_proto_rawDesc = "" +
	"\n" +
	"\x1dheadscale/v1/preauthkey.proto\x12\fheadscale.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17headscale/v1/user.proto\"\xb6\x02\n" +
	"\n" +
	"PreAuthKey\x12&\n" +
	"\x04user\x18\x01 \x01(\v2\x12.headscale.v1.UserR\x04user\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\x04R\x02id\x12\x10\n" +
	"\x03key\x18\x03 \x01(\tR\x03key\x12\x1a\n" +
	"\breusable\x18\x04 \x01(\bR\breusable\x12\x1c\n" +
	"\tephemeral\x18\x05 \x01(\bR\tephemeral\x12\x12\n" +
	"\x04used\x18\x06 \x01(\bR\x04used\x12:\n" +
	"\n" +
	"expiration\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"expiration\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x19\n" +
	"\bacl_tags\x18\t \x03(\tR\aaclTags\"\xbe\x01\n" +
	"\x17CreatePreAuthKeyRequest\x12\x12\n" +
	"\x04user\x18\x01 \x01(\x04R\x04user\x12\x1a\n" +
	"\breusable\x18\x02 \x01(\bR\breusable\x12\x1c\n" +
	"\tephemeral\x18\x03 \x01(\bR\tephemeral\x12:\n" +
	"\n" +
	"expiration\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"expiration\x12\x19\n" +
	"\bacl_tags\x18\x05 \x03(\tR\aaclTags\"V\n" +
	"\x18CreatePreAuthKeyResponse\x12:\n" +
	"\fpre_auth_key\x18\x01 \x01(\v2\x18.headscale.v1.PreAuthKeyR\n" +
	"preAuthKey\"?\n" +
	"\x17ExpirePreAuthKeyRequest\x12\x12\n" +
	"\x04user\x18\x01 \x01(\x04R\x04user\x12\x10\n" +
	"\x03key\x18\x02 \x01(\tR\x03key\"\x1a\n" +
	"\x18ExpirePreAuthKeyResponse\",\n" +
	"\x16ListPreAuthKeysRequest\x12\x12\n" +
	"\x04user\x18\x01 \x01(\x04R\x04user\"W\n" +
	"\x17ListPreAuthKeysResponse\x12<\n" +
	"\rpre_auth_keys\x18\x01 \x03(\v2\x18.headscale.v1.PreAuthKeyR\vpreAuthKeysB)Z'github.com/juanfont/headscale/gen/go/v1b\x06proto3"

var (
	file_headscale_v1_preauthkey_proto_rawDescOnce sync.Once
	file_headscale_v1_preauthkey_proto_rawDescData []byte
)

func file_headscale_v1_preauthkey_proto_rawDescGZIP() []byte {
	file_headscale_v1_preauthkey_proto_rawDescOnce.Do(func() {
		file_headscale_v1_preauthkey_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_headscale_v1_preauthkey_proto_rawDesc), len(file_headscale_v1_preauthkey_proto_rawDesc)))
	})
	return file_headscale_v1_preauthkey_proto_rawDescData
}

var file_headscale_v1_preauthkey_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_headscale_v1_preauthkey_proto_goTypes = []any{
	(*PreAuthKey)(nil),               // 0: headscale.v1.PreAuthKey
	(*CreatePreAuthKeyRequest)(nil),  // 1: headscale.v1.CreatePreAuthKeyRequest
	(*CreatePreAuthKeyResponse)(nil), // 2: headscale.v1.CreatePreAuthKeyResponse
	(*ExpirePreAuthKeyRequest)(nil),  // 3: headscale.v1.ExpirePreAuthKeyRequest
	(*ExpirePreAuthKeyResponse)(nil), // 4: headscale.v1.ExpirePreAuthKeyResponse
	(*ListPreAuthKeysRequest)(nil),   // 5: headscale.v1.ListPreAuthKeysRequest
	(*ListPreAuthKeysResponse)(nil),  // 6: headscale.v1.ListPreAuthKeysResponse
	(*User)(nil),                     // 7: headscale.v1.User
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_headscale_v1_preauthkey_proto_depIdxs = []int32{
	7, // 0: headscale.v1.PreAuthKey.user:type_name -> headscale.v1.User
	8, // 1: headscale.v1.PreAuthKey.expiration:type_name -> google.protobuf.Timestamp
	8, // 2: headscale.v1.PreAuthKey.created_at:type_name -> google.protobuf.Timestamp
	8, // 3: headscale.v1.CreatePreAuthKeyRequest.expiration:type_name -> google.protobuf.Timestamp
	0, // 4: headscale.v1.CreatePreAuthKeyResponse.pre_auth_key:type_name -> headscale.v1.PreAuthKey
	0, // 5: headscale.v1.ListPreAuthKeysResponse.pre_auth_keys:type_name -> headscale.v1.PreAuthKey
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_headscale_v1_preauthkey_proto_init() }
func file_headscale_v1_preauthkey_proto_init() {
	if File_headscale_v1_preauthkey_proto != nil {
		return
	}
	file_headscale_v1_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_headscale_v1_preauthkey_proto_rawDesc), len(file_headscale_v1_preauthkey_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_headscale_v1_preauthkey_proto_goTypes,
		DependencyIndexes: file_headscale_v1_preauthkey_proto_depIdxs,
		MessageInfos:      file_headscale_v1_preauthkey_proto_msgTypes,
	}.Build()
	File_headscale_v1_preauthkey_proto = out.File
	file_headscale_v1_preauthkey_proto_goTypes = nil
	file_headscale_v1_preauthkey_proto_depIdxs = nil
}
