// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: credential.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectType string `protobuf:"bytes,1,opt,name=subjectType,proto3" json:"subjectType,omitempty"`
	Subject     string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	IsSystem    bool   `protobuf:"varint,3,opt,name=isSystem,proto3" json:"isSystem,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateAccessKeyRequest) Reset() {
	*x = CreateAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessKeyRequest) ProtoMessage() {}

func (x *CreateAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccessKeyRequest) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *CreateAccessKeyRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *CreateAccessKeyRequest) GetIsSystem() bool {
	if x != nil {
		return x.IsSystem
	}
	return false
}

func (x *CreateAccessKeyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *AccessKey `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateAccessKeyResponse) Reset() {
	*x = CreateAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessKeyResponse) ProtoMessage() {}

func (x *CreateAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccessKeyResponse) GetData() *AccessKey {
	if x != nil {
		return x.Data
	}
	return nil
}

type AccessKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AccessKeyId string                 `protobuf:"bytes,2,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
	SecretKey   string                 `protobuf:"bytes,3,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	IsSystem    bool                   `protobuf:"varint,4,opt,name=isSystem,proto3" json:"isSystem,omitempty"`
	Status      string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	SubjectType string                 `protobuf:"bytes,6,opt,name=subjectType,proto3" json:"subjectType,omitempty"`
	Subject     string                 `protobuf:"bytes,7,opt,name=subject,proto3" json:"subject,omitempty"`
	Description string                 `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *AccessKey) Reset() {
	*x = AccessKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKey) ProtoMessage() {}

func (x *AccessKey) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKey.ProtoReflect.Descriptor instead.
func (*AccessKey) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{2}
}

func (x *AccessKey) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AccessKey) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *AccessKey) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *AccessKey) GetIsSystem() bool {
	if x != nil {
		return x.IsSystem
	}
	return false
}

func (x *AccessKey) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AccessKey) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *AccessKey) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *AccessKey) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AccessKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AccessKey) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DeleteAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKeyId string `protobuf:"bytes,1,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
}

func (x *DeleteAccessKeyRequest) Reset() {
	*x = DeleteAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessKeyRequest) ProtoMessage() {}

func (x *DeleteAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAccessKeyRequest) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

type DeleteAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccessKeyResponse) Reset() {
	*x = DeleteAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessKeyResponse) ProtoMessage() {}

func (x *DeleteAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{4}
}

type GetAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKeyId string `protobuf:"bytes,1,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
}

func (x *GetAccessKeyRequest) Reset() {
	*x = GetAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessKeyRequest) ProtoMessage() {}

func (x *GetAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*GetAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccessKeyRequest) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

type GetAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *AccessKey `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAccessKeyResponse) Reset() {
	*x = GetAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessKeyResponse) ProtoMessage() {}

func (x *GetAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*GetAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{6}
}

func (x *GetAccessKeyResponse) GetData() *AccessKey {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadAccessKeyFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKeyId string `protobuf:"bytes,1,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
	Path        string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DownloadAccessKeyFileRequest) Reset() {
	*x = DownloadAccessKeyFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadAccessKeyFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadAccessKeyFileRequest) ProtoMessage() {}

func (x *DownloadAccessKeyFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadAccessKeyFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadAccessKeyFileRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadAccessKeyFileRequest) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *DownloadAccessKeyFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DownloadAccessKeyFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DownloadAccessKeyFileResponse) Reset() {
	*x = DownloadAccessKeyFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadAccessKeyFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadAccessKeyFileResponse) ProtoMessage() {}

func (x *DownloadAccessKeyFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadAccessKeyFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadAccessKeyFileResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{8}
}

var File_credential_proto protoreflect.FileDescriptor

var file_credential_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4d, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe5, 0x02, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64,
	0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x54, 0x0a, 0x1c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa0, 0x05, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x01, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x12, 0x2b, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2d,
	0x6b, 0x65, 0x79, 0x73, 0x12, 0xa3, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x2a, 0x2d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x7d, 0x12, 0x9a, 0x01, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x73, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x7d, 0x12, 0xb0, 0x01, 0x0a, 0x15, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x31, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a,
	0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2d, 0x6b, 0x65, 0x79,
	0x73, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_credential_proto_rawDescOnce sync.Once
	file_credential_proto_rawDescData = file_credential_proto_rawDesc
)

func file_credential_proto_rawDescGZIP() []byte {
	file_credential_proto_rawDescOnce.Do(func() {
		file_credential_proto_rawDescData = protoimpl.X.CompressGZIP(file_credential_proto_rawDescData)
	})
	return file_credential_proto_rawDescData
}

var file_credential_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_credential_proto_goTypes = []interface{}{
	(*CreateAccessKeyRequest)(nil),        // 0: erda.msp.credential.CreateAccessKeyRequest
	(*CreateAccessKeyResponse)(nil),       // 1: erda.msp.credential.CreateAccessKeyResponse
	(*AccessKey)(nil),                     // 2: erda.msp.credential.AccessKey
	(*DeleteAccessKeyRequest)(nil),        // 3: erda.msp.credential.DeleteAccessKeyRequest
	(*DeleteAccessKeyResponse)(nil),       // 4: erda.msp.credential.DeleteAccessKeyResponse
	(*GetAccessKeyRequest)(nil),           // 5: erda.msp.credential.GetAccessKeyRequest
	(*GetAccessKeyResponse)(nil),          // 6: erda.msp.credential.GetAccessKeyResponse
	(*DownloadAccessKeyFileRequest)(nil),  // 7: erda.msp.credential.DownloadAccessKeyFileRequest
	(*DownloadAccessKeyFileResponse)(nil), // 8: erda.msp.credential.DownloadAccessKeyFileResponse
	(*timestamppb.Timestamp)(nil),         // 9: google.protobuf.Timestamp
}
var file_credential_proto_depIdxs = []int32{
	2, // 0: erda.msp.credential.CreateAccessKeyResponse.data:type_name -> erda.msp.credential.AccessKey
	9, // 1: erda.msp.credential.AccessKey.createdAt:type_name -> google.protobuf.Timestamp
	9, // 2: erda.msp.credential.AccessKey.updatedAt:type_name -> google.protobuf.Timestamp
	2, // 3: erda.msp.credential.GetAccessKeyResponse.data:type_name -> erda.msp.credential.AccessKey
	0, // 4: erda.msp.credential.AccessKeyService.CreateAccessKey:input_type -> erda.msp.credential.CreateAccessKeyRequest
	3, // 5: erda.msp.credential.AccessKeyService.DeleteAccessKey:input_type -> erda.msp.credential.DeleteAccessKeyRequest
	5, // 6: erda.msp.credential.AccessKeyService.GetAccessKey:input_type -> erda.msp.credential.GetAccessKeyRequest
	7, // 7: erda.msp.credential.AccessKeyService.DownloadAccessKeyFile:input_type -> erda.msp.credential.DownloadAccessKeyFileRequest
	1, // 8: erda.msp.credential.AccessKeyService.CreateAccessKey:output_type -> erda.msp.credential.CreateAccessKeyResponse
	4, // 9: erda.msp.credential.AccessKeyService.DeleteAccessKey:output_type -> erda.msp.credential.DeleteAccessKeyResponse
	6, // 10: erda.msp.credential.AccessKeyService.GetAccessKey:output_type -> erda.msp.credential.GetAccessKeyResponse
	8, // 11: erda.msp.credential.AccessKeyService.DownloadAccessKeyFile:output_type -> erda.msp.credential.DownloadAccessKeyFileResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_credential_proto_init() }
func file_credential_proto_init() {
	if File_credential_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_credential_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessKeyRequest); i {
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
		file_credential_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessKeyResponse); i {
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
		file_credential_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKey); i {
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
		file_credential_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessKeyRequest); i {
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
		file_credential_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessKeyResponse); i {
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
		file_credential_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessKeyRequest); i {
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
		file_credential_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessKeyResponse); i {
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
		file_credential_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadAccessKeyFileRequest); i {
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
		file_credential_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadAccessKeyFileResponse); i {
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
			RawDescriptor: file_credential_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_credential_proto_goTypes,
		DependencyIndexes: file_credential_proto_depIdxs,
		MessageInfos:      file_credential_proto_msgTypes,
	}.Build()
	File_credential_proto = out.File
	file_credential_proto_rawDesc = nil
	file_credential_proto_goTypes = nil
	file_credential_proto_depIdxs = nil
}