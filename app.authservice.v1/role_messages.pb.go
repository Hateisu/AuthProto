// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: role_messages.proto

package authservicev1

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

type RoleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleCreateRequest) Reset() {
	*x = RoleCreateRequest{}
	mi := &file_role_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateRequest) ProtoMessage() {}

func (x *RoleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateRequest.ProtoReflect.Descriptor instead.
func (*RoleCreateRequest) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{0}
}

func (x *RoleCreateRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type RoleCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleCreateResponse) Reset() {
	*x = RoleCreateResponse{}
	mi := &file_role_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateResponse) ProtoMessage() {}

func (x *RoleCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateResponse.ProtoReflect.Descriptor instead.
func (*RoleCreateResponse) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{1}
}

func (x *RoleCreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleReadRequest) Reset() {
	*x = RoleReadRequest{}
	mi := &file_role_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleReadRequest) ProtoMessage() {}

func (x *RoleReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleReadRequest.ProtoReflect.Descriptor instead.
func (*RoleReadRequest) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{2}
}

func (x *RoleReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleReadResponse) Reset() {
	*x = RoleReadResponse{}
	mi := &file_role_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleReadResponse) ProtoMessage() {}

func (x *RoleReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleReadResponse.ProtoReflect.Descriptor instead.
func (*RoleReadResponse) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{3}
}

func (x *RoleReadResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type RoleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role *Role  `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleUpdateRequest) Reset() {
	*x = RoleUpdateRequest{}
	mi := &file_role_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateRequest) ProtoMessage() {}

func (x *RoleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RoleUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleUpdateRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type RoleUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleUpdateResponse) Reset() {
	*x = RoleUpdateResponse{}
	mi := &file_role_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateResponse) ProtoMessage() {}

func (x *RoleUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateResponse.ProtoReflect.Descriptor instead.
func (*RoleUpdateResponse) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{5}
}

func (x *RoleUpdateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleDeleteRequest) Reset() {
	*x = RoleDeleteRequest{}
	mi := &file_role_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteRequest) ProtoMessage() {}

func (x *RoleDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteRequest.ProtoReflect.Descriptor instead.
func (*RoleDeleteRequest) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{6}
}

func (x *RoleDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoleDeleteResponse) Reset() {
	*x = RoleDeleteResponse{}
	mi := &file_role_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteResponse) ProtoMessage() {}

func (x *RoleDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteResponse.ProtoReflect.Descriptor instead.
func (*RoleDeleteResponse) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{7}
}

var File_role_messages_proto protoreflect.FileDescriptor

var file_role_messages_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x10, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4a, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12,
	0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_messages_proto_rawDescOnce sync.Once
	file_role_messages_proto_rawDescData = file_role_messages_proto_rawDesc
)

func file_role_messages_proto_rawDescGZIP() []byte {
	file_role_messages_proto_rawDescOnce.Do(func() {
		file_role_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_messages_proto_rawDescData)
	})
	return file_role_messages_proto_rawDescData
}

var file_role_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_role_messages_proto_goTypes = []any{
	(*RoleCreateRequest)(nil),  // 0: authservice.RoleCreateRequest
	(*RoleCreateResponse)(nil), // 1: authservice.RoleCreateResponse
	(*RoleReadRequest)(nil),    // 2: authservice.RoleReadRequest
	(*RoleReadResponse)(nil),   // 3: authservice.RoleReadResponse
	(*RoleUpdateRequest)(nil),  // 4: authservice.RoleUpdateRequest
	(*RoleUpdateResponse)(nil), // 5: authservice.RoleUpdateResponse
	(*RoleDeleteRequest)(nil),  // 6: authservice.RoleDeleteRequest
	(*RoleDeleteResponse)(nil), // 7: authservice.RoleDeleteResponse
	(*Role)(nil),               // 8: authservice.Role
}
var file_role_messages_proto_depIdxs = []int32{
	8, // 0: authservice.RoleCreateRequest.role:type_name -> authservice.Role
	8, // 1: authservice.RoleReadResponse.role:type_name -> authservice.Role
	8, // 2: authservice.RoleUpdateRequest.role:type_name -> authservice.Role
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_role_messages_proto_init() }
func file_role_messages_proto_init() {
	if File_role_messages_proto != nil {
		return
	}
	file_auth_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_role_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_messages_proto_goTypes,
		DependencyIndexes: file_role_messages_proto_depIdxs,
		MessageInfos:      file_role_messages_proto_msgTypes,
	}.Build()
	File_role_messages_proto = out.File
	file_role_messages_proto_rawDesc = nil
	file_role_messages_proto_goTypes = nil
	file_role_messages_proto_depIdxs = nil
}
