// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Common.proto

package Common

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

type ErrorCodes int32

const (
	ErrorCodes_ErrorCodes_NONE ErrorCodes = 0
	// 權限
	ErrorCodes_SUCCESS           ErrorCodes = 10000
	ErrorCodes_INVAILD_TOKEN     ErrorCodes = 10001
	ErrorCodes_ACCOUNT_DELETED   ErrorCodes = 10002
	ErrorCodes_LOGIN_INVALID     ErrorCodes = 10003
	ErrorCodes_PERMISSION_REJECT ErrorCodes = 10004
	// API
	ErrorCodes_INVAILD_PARAM  ErrorCodes = 20001
	ErrorCodes_REPEATED_ERROR ErrorCodes = 20003
	ErrorCodes_DATA_NOT_FOUND ErrorCodes = 20004
	ErrorCodes_INTERNAL_ERROR ErrorCodes = 20005
	ErrorCodes_DB_ERROR       ErrorCodes = 20006
	// File
	ErrorCodes_UPLOAD_FILE_SIZE_INVALID     ErrorCodes = 40001
	ErrorCodes_UPLOAD_FILE_TYPE_NOT_SUPPORT ErrorCodes = 40002
	// Email
	ErrorCodes_EMAIL_LIMIT_REACH ErrorCodes = 50001
	ErrorCodes_EMAIL_TEST        ErrorCodes = 50002
	// WebSocket
	ErrorCodes_SEND_COMMAND_TIMEOUT                  ErrorCodes = 60001
	ErrorCodes_RECEIVER_NOT_EXIST_IN_CONNECTION_POOL ErrorCodes = 60002
)

// Enum value maps for ErrorCodes.
var (
	ErrorCodes_name = map[int32]string{
		0:     "ErrorCodes_NONE",
		10000: "SUCCESS",
		10001: "INVAILD_TOKEN",
		10002: "ACCOUNT_DELETED",
		10003: "LOGIN_INVALID",
		10004: "PERMISSION_REJECT",
		20001: "INVAILD_PARAM",
		20003: "REPEATED_ERROR",
		20004: "DATA_NOT_FOUND",
		20005: "INTERNAL_ERROR",
		20006: "DB_ERROR",
		40001: "UPLOAD_FILE_SIZE_INVALID",
		40002: "UPLOAD_FILE_TYPE_NOT_SUPPORT",
		50001: "EMAIL_LIMIT_REACH",
		50002: "EMAIL_TEST",
		60001: "SEND_COMMAND_TIMEOUT",
		60002: "RECEIVER_NOT_EXIST_IN_CONNECTION_POOL",
	}
	ErrorCodes_value = map[string]int32{
		"ErrorCodes_NONE":                       0,
		"SUCCESS":                               10000,
		"INVAILD_TOKEN":                         10001,
		"ACCOUNT_DELETED":                       10002,
		"LOGIN_INVALID":                         10003,
		"PERMISSION_REJECT":                     10004,
		"INVAILD_PARAM":                         20001,
		"REPEATED_ERROR":                        20003,
		"DATA_NOT_FOUND":                        20004,
		"INTERNAL_ERROR":                        20005,
		"DB_ERROR":                              20006,
		"UPLOAD_FILE_SIZE_INVALID":              40001,
		"UPLOAD_FILE_TYPE_NOT_SUPPORT":          40002,
		"EMAIL_LIMIT_REACH":                     50001,
		"EMAIL_TEST":                            50002,
		"SEND_COMMAND_TIMEOUT":                  60001,
		"RECEIVER_NOT_EXIST_IN_CONNECTION_POOL": 60002,
	}
)

func (x ErrorCodes) Enum() *ErrorCodes {
	p := new(ErrorCodes)
	*p = x
	return p
}

func (x ErrorCodes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCodes) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[0].Descriptor()
}

func (ErrorCodes) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[0]
}

func (x ErrorCodes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCodes.Descriptor instead.
func (ErrorCodes) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{0}
}

type CommonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommonRequest) Reset() {
	*x = CommonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRequest) ProtoMessage() {}

func (x *CommonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRequest.ProtoReflect.Descriptor instead.
func (*CommonRequest) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{0}
}

type CommonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommonReply) Reset() {
	*x = CommonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReply) ProtoMessage() {}

func (x *CommonReply) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReply.ProtoReflect.Descriptor instead.
func (*CommonReply) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{1}
}

type PageInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageItemCount  int32 `protobuf:"varint,1,opt,name=page_item_count,json=pageItemCount,proto3" json:"page_item_count,omitempty"`
	CurrentPageNum int32 `protobuf:"varint,2,opt,name=current_page_num,json=currentPageNum,proto3" json:"current_page_num,omitempty"`
}

func (x *PageInfoRequest) Reset() {
	*x = PageInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoRequest) ProtoMessage() {}

func (x *PageInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoRequest.ProtoReflect.Descriptor instead.
func (*PageInfoRequest) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{2}
}

func (x *PageInfoRequest) GetPageItemCount() int32 {
	if x != nil {
		return x.PageItemCount
	}
	return 0
}

func (x *PageInfoRequest) GetCurrentPageNum() int32 {
	if x != nil {
		return x.CurrentPageNum
	}
	return 0
}

type PageInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *PageInfoReply) Reset() {
	*x = PageInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoReply) ProtoMessage() {}

func (x *PageInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoReply.ProtoReflect.Descriptor instead.
func (*PageInfoReply) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{3}
}

func (x *PageInfoReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type ErrorReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorCodes `protobuf:"varint,1,opt,name=code,proto3,enum=Common.ErrorCodes" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorReply) Reset() {
	*x = ErrorReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorReply) ProtoMessage() {}

func (x *ErrorReply) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorReply.ProtoReflect.Descriptor instead.
func (*ErrorReply) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorReply) GetCode() ErrorCodes {
	if x != nil {
		return x.Code
	}
	return ErrorCodes_ErrorCodes_NONE
}

func (x *ErrorReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ErrorReplyGin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     ErrorCodes `protobuf:"varint,1,opt,name=code,proto3,enum=Common.ErrorCodes" json:"code,omitempty"`
	Message  string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileName string     `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath string     `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *ErrorReplyGin) Reset() {
	*x = ErrorReplyGin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorReplyGin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorReplyGin) ProtoMessage() {}

func (x *ErrorReplyGin) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorReplyGin.ProtoReflect.Descriptor instead.
func (*ErrorReplyGin) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorReplyGin) GetCode() ErrorCodes {
	if x != nil {
		return x.Code
	}
	return ErrorCodes_ErrorCodes_NONE
}

func (x *ErrorReplyGin) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorReplyGin) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ErrorReplyGin) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_Common_proto protoreflect.FileDescriptor

var file_Common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x63, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x30, 0x0a, 0x0d, 0x50,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4e, 0x0a,
	0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01,
	0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x69, 0x6e, 0x12,
	0x26, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x2a, 0xa4, 0x03, 0x0a, 0x0a,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x90, 0x4e, 0x12, 0x12, 0x0a,
	0x0d, 0x49, 0x4e, 0x56, 0x41, 0x49, 0x4c, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x91,
	0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x92, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x93, 0x4e, 0x12, 0x16, 0x0a, 0x11, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x10, 0x94, 0x4e, 0x12, 0x13, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x49, 0x4c, 0x44, 0x5f, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x10, 0xa1, 0x9c, 0x01, 0x12, 0x14, 0x0a, 0x0e, 0x52, 0x45, 0x50, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa3, 0x9c, 0x01, 0x12, 0x14,
	0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0xa4, 0x9c, 0x01, 0x12, 0x14, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa5, 0x9c, 0x01, 0x12, 0x0e, 0x0a, 0x08, 0x44, 0x42,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa6, 0x9c, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x55, 0x50,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xc1, 0xb8, 0x02, 0x12, 0x22, 0x0a, 0x1c, 0x55, 0x50,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0xc2, 0xb8, 0x02, 0x12, 0x17,
	0x0a, 0x11, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x52, 0x45,
	0x41, 0x43, 0x48, 0x10, 0xd1, 0x86, 0x03, 0x12, 0x10, 0x0a, 0x0a, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0xd2, 0x86, 0x03, 0x12, 0x1a, 0x0a, 0x14, 0x53, 0x45, 0x4e,
	0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0xe1, 0xd4, 0x03, 0x12, 0x2b, 0x0a, 0x25, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0xe2,
	0xd4, 0x03, 0x42, 0x1e, 0x5a, 0x1c, 0x70, 0x69, 0x63, 0x62, 0x6f, 0x74, 0x2d, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_proto_rawDescOnce sync.Once
	file_Common_proto_rawDescData = file_Common_proto_rawDesc
)

func file_Common_proto_rawDescGZIP() []byte {
	file_Common_proto_rawDescOnce.Do(func() {
		file_Common_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_proto_rawDescData)
	})
	return file_Common_proto_rawDescData
}

var file_Common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Common_proto_goTypes = []interface{}{
	(ErrorCodes)(0),         // 0: Common.ErrorCodes
	(*CommonRequest)(nil),   // 1: Common.CommonRequest
	(*CommonReply)(nil),     // 2: Common.CommonReply
	(*PageInfoRequest)(nil), // 3: Common.PageInfoRequest
	(*PageInfoReply)(nil),   // 4: Common.PageInfoReply
	(*ErrorReply)(nil),      // 5: Common.ErrorReply
	(*ErrorReplyGin)(nil),   // 6: Common.ErrorReplyGin
}
var file_Common_proto_depIdxs = []int32{
	0, // 0: Common.ErrorReply.code:type_name -> Common.ErrorCodes
	0, // 1: Common.ErrorReplyGin.code:type_name -> Common.ErrorCodes
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Common_proto_init() }
func file_Common_proto_init() {
	if File_Common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRequest); i {
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
		file_Common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReply); i {
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
		file_Common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoRequest); i {
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
		file_Common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoReply); i {
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
		file_Common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorReply); i {
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
		file_Common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorReplyGin); i {
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
			RawDescriptor: file_Common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_proto_goTypes,
		DependencyIndexes: file_Common_proto_depIdxs,
		EnumInfos:         file_Common_proto_enumTypes,
		MessageInfos:      file_Common_proto_msgTypes,
	}.Build()
	File_Common_proto = out.File
	file_Common_proto_rawDesc = nil
	file_Common_proto_goTypes = nil
	file_Common_proto_depIdxs = nil
}
