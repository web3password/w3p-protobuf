// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: storage.proto

package storage

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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Params    string `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UploadRequest) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *UploadRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *UploadReply_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadReply) Reset() {
	*x = UploadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReply) ProtoMessage() {}

func (x *UploadReply) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReply.ProtoReflect.Descriptor instead.
func (*UploadReply) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *UploadReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UploadReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UploadReply) GetData() *UploadReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Params    string `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *DownloadRequest) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

type DownloadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32               `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *DownloadReply_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DownloadReply) Reset() {
	*x = DownloadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReply) ProtoMessage() {}

func (x *DownloadReply) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadReply.ProtoReflect.Descriptor instead.
func (*DownloadReply) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DownloadReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DownloadReply) GetData() *DownloadReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type AttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Params    string `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AttachmentRequest) Reset() {
	*x = AttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentRequest) ProtoMessage() {}

func (x *AttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentRequest.ProtoReflect.Descriptor instead.
func (*AttachmentRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{4}
}

func (x *AttachmentRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AttachmentRequest) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *AttachmentRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AttachmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *AttachmentReply_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AttachmentReply) Reset() {
	*x = AttachmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentReply) ProtoMessage() {}

func (x *AttachmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentReply.ProtoReflect.Descriptor instead.
func (*AttachmentReply) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{5}
}

func (x *AttachmentReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AttachmentReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AttachmentReply) GetData() *AttachmentReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Params    string `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{6}
}

func (x *ReportRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ReportRequest) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *ReportRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ReportReply) Reset() {
	*x = ReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReply) ProtoMessage() {}

func (x *ReportReply) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReply.ProtoReflect.Descriptor instead.
func (*ReportReply) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{7}
}

func (x *ReportReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UploadReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *UploadReply_Data) Reset() {
	*x = UploadReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReply_Data) ProtoMessage() {}

func (x *UploadReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReply_Data.ProtoReflect.Descriptor instead.
func (*UploadReply_Data) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1, 0}
}

func (x *UploadReply_Data) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type DownloadReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid     string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *DownloadReply_Data) Reset() {
	*x = DownloadReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReply_Data) ProtoMessage() {}

func (x *DownloadReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadReply_Data.ProtoReflect.Descriptor instead.
func (*DownloadReply_Data) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DownloadReply_Data) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *DownloadReply_Data) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type AttachmentReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid     string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AttachmentReply_Data) Reset() {
	*x = AttachmentReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentReply_Data) ProtoMessage() {}

func (x *AttachmentReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentReply_Data.ProtoReflect.Descriptor instead.
func (*AttachmentReply_Data) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{5, 0}
}

func (x *AttachmentReply_Data) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *AttachmentReply_Data) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x7c, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x18, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69,
	0x64, 0x22, 0x47, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x32, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x32, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x27, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x83, 0x02, 0x0a, 0x07,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x61, 0x6d, 0x68, 0x65, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData = file_storage_proto_rawDesc
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_proto_rawDescData)
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_storage_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),        // 0: storage.UploadRequest
	(*UploadReply)(nil),          // 1: storage.UploadReply
	(*DownloadRequest)(nil),      // 2: storage.DownloadRequest
	(*DownloadReply)(nil),        // 3: storage.DownloadReply
	(*AttachmentRequest)(nil),    // 4: storage.AttachmentRequest
	(*AttachmentReply)(nil),      // 5: storage.AttachmentReply
	(*ReportRequest)(nil),        // 6: storage.ReportRequest
	(*ReportReply)(nil),          // 7: storage.ReportReply
	(*UploadReply_Data)(nil),     // 8: storage.UploadReply.Data
	(*DownloadReply_Data)(nil),   // 9: storage.DownloadReply.Data
	(*AttachmentReply_Data)(nil), // 10: storage.AttachmentReply.Data
}
var file_storage_proto_depIdxs = []int32{
	8,  // 0: storage.UploadReply.data:type_name -> storage.UploadReply.Data
	9,  // 1: storage.DownloadReply.data:type_name -> storage.DownloadReply.Data
	10, // 2: storage.AttachmentReply.data:type_name -> storage.AttachmentReply.Data
	0,  // 3: storage.Storage.Upload:input_type -> storage.UploadRequest
	2,  // 4: storage.Storage.Download:input_type -> storage.DownloadRequest
	4,  // 5: storage.Storage.Attachment:input_type -> storage.AttachmentRequest
	6,  // 6: storage.Storage.Report:input_type -> storage.ReportRequest
	1,  // 7: storage.Storage.Upload:output_type -> storage.UploadReply
	3,  // 8: storage.Storage.Download:output_type -> storage.DownloadReply
	5,  // 9: storage.Storage.Attachment:output_type -> storage.AttachmentReply
	7,  // 10: storage.Storage.Report:output_type -> storage.ReportReply
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReply); i {
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
		file_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadReply); i {
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
		file_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentRequest); i {
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
		file_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentReply); i {
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
		file_storage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_storage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReply); i {
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
		file_storage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReply_Data); i {
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
		file_storage_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadReply_Data); i {
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
		file_storage_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentReply_Data); i {
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
			RawDescriptor: file_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_rawDesc = nil
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}