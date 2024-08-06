// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/v1/auth_service.proto

package apiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAuthStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthStatusRequest) Reset() {
	*x = GetAuthStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthStatusRequest) ProtoMessage() {}

func (x *GetAuthStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthStatusRequest.ProtoReflect.Descriptor instead.
func (*GetAuthStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{0}
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInWithSSORequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the SSO provider.
	IdpName int32 `protobuf:"varint,1,opt,name=idp_name,json=idpName,proto3" json:"idp_name,omitempty"`
	// The code to sign in with.
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// The redirect URI.
	RedirectUri string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *SignInWithSSORequest) Reset() {
	*x = SignInWithSSORequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInWithSSORequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInWithSSORequest) ProtoMessage() {}

func (x *SignInWithSSORequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInWithSSORequest.ProtoReflect.Descriptor instead.
func (*SignInWithSSORequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *SignInWithSSORequest) GetIdpName() int32 {
	if x != nil {
		return x.IdpName
	}
	return 0
}

func (x *SignInWithSSORequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SignInWithSSORequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type SignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutRequest) Reset() {
	*x = SignOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRequest) ProtoMessage() {}

func (x *SignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRequest.ProtoReflect.Descriptor instead.
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{4}
}

var File_api_v1_auth_service_proto protoreflect.FileDescriptor

var file_api_v1_auth_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5d, 0x0a, 0x0d, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x68, 0x0a, 0x14, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x64, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x69, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xec, 0x03, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x56, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x12, 0x68, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74,
	0x68, 0x53, 0x53, 0x4f, 0x12, 0x22, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x53, 0x53,
	0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x2f, 0x73, 0x73, 0x6f, 0x12, 0x56, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x5d, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74,
	0x12, 0x1c, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x14,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x6f, 0x75, 0x74, 0x42, 0xae, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x6c, 0x66, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x53, 0x6c, 0x61,
	0x73, 0x68, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_auth_service_proto_rawDescOnce sync.Once
	file_api_v1_auth_service_proto_rawDescData = file_api_v1_auth_service_proto_rawDesc
)

func file_api_v1_auth_service_proto_rawDescGZIP() []byte {
	file_api_v1_auth_service_proto_rawDescOnce.Do(func() {
		file_api_v1_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_auth_service_proto_rawDescData)
	})
	return file_api_v1_auth_service_proto_rawDescData
}

var file_api_v1_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_auth_service_proto_goTypes = []any{
	(*GetAuthStatusRequest)(nil), // 0: slash.api.v1.GetAuthStatusRequest
	(*SignInRequest)(nil),        // 1: slash.api.v1.SignInRequest
	(*SignUpRequest)(nil),        // 2: slash.api.v1.SignUpRequest
	(*SignInWithSSORequest)(nil), // 3: slash.api.v1.SignInWithSSORequest
	(*SignOutRequest)(nil),       // 4: slash.api.v1.SignOutRequest
	(*User)(nil),                 // 5: slash.api.v1.User
	(*emptypb.Empty)(nil),        // 6: google.protobuf.Empty
}
var file_api_v1_auth_service_proto_depIdxs = []int32{
	0, // 0: slash.api.v1.AuthService.GetAuthStatus:input_type -> slash.api.v1.GetAuthStatusRequest
	1, // 1: slash.api.v1.AuthService.SignIn:input_type -> slash.api.v1.SignInRequest
	3, // 2: slash.api.v1.AuthService.SignInWithSSO:input_type -> slash.api.v1.SignInWithSSORequest
	2, // 3: slash.api.v1.AuthService.SignUp:input_type -> slash.api.v1.SignUpRequest
	4, // 4: slash.api.v1.AuthService.SignOut:input_type -> slash.api.v1.SignOutRequest
	5, // 5: slash.api.v1.AuthService.GetAuthStatus:output_type -> slash.api.v1.User
	5, // 6: slash.api.v1.AuthService.SignIn:output_type -> slash.api.v1.User
	5, // 7: slash.api.v1.AuthService.SignInWithSSO:output_type -> slash.api.v1.User
	5, // 8: slash.api.v1.AuthService.SignUp:output_type -> slash.api.v1.User
	6, // 9: slash.api.v1.AuthService.SignOut:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_auth_service_proto_init() }
func file_api_v1_auth_service_proto_init() {
	if File_api_v1_auth_service_proto != nil {
		return
	}
	file_api_v1_user_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_auth_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuthStatusRequest); i {
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
		file_api_v1_auth_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SignInRequest); i {
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
		file_api_v1_auth_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SignUpRequest); i {
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
		file_api_v1_auth_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SignInWithSSORequest); i {
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
		file_api_v1_auth_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SignOutRequest); i {
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
			RawDescriptor: file_api_v1_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_auth_service_proto_goTypes,
		DependencyIndexes: file_api_v1_auth_service_proto_depIdxs,
		MessageInfos:      file_api_v1_auth_service_proto_msgTypes,
	}.Build()
	File_api_v1_auth_service_proto = out.File
	file_api_v1_auth_service_proto_rawDesc = nil
	file_api_v1_auth_service_proto_goTypes = nil
	file_api_v1_auth_service_proto_depIdxs = nil
}
