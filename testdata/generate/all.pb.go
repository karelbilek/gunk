// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: testdata.tld/util/all.proto

// package util contains a simple Echo service.

package util

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	imported "testdata.tld/util/imported"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status is a server health status.
type Status int32

const (
	// Status_Unknown is the default, unset status value.
	Status_Unknown Status = 0
	// Status_Error is a status value that implies something went wrong.
	Status_Error Status = 1
	// Status_OK is a status value used when all went well.
	Status_OK Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "Error",
		2: "OK",
	}
	Status_value = map[string]int32{
		"Unknown": 0,
		"Error":   1,
		"OK":      2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_tld_util_all_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_testdata_tld_util_all_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_testdata_tld_util_all_proto_rawDescGZIP(), []int{0}
}

// CheckStatusResponse is the response for a check status.
type CheckStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=Status,json=status,proto3,enum=testdata.v1.util.Status" json:"status,omitempty"`
}

func (x *CheckStatusResponse) Reset() {
	*x = CheckStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_tld_util_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStatusResponse) ProtoMessage() {}

func (x *CheckStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_tld_util_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStatusResponse.ProtoReflect.Descriptor instead.
func (*CheckStatusResponse) Descriptor() ([]byte, []int) {
	return file_testdata_tld_util_all_proto_rawDescGZIP(), []int{0}
}

func (x *CheckStatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Unknown
}

type UtilTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ints       []int32                        `protobuf:"varint,1,rep,packed,name=Ints,proto3" json:"Ints,omitempty"`
	Structs    []*imported.Message            `protobuf:"bytes,2,rep,name=Structs,proto3" json:"Structs,omitempty"`
	Bool       bool                           `protobuf:"varint,3,opt,name=Bool,proto3" json:"Bool,omitempty"`
	Int32      int32                          `protobuf:"varint,4,opt,name=Int32,proto3" json:"Int32,omitempty"`
	Timestamp  *timestamppb.Timestamp         `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Duration   *durationpb.Duration           `protobuf:"bytes,6,opt,name=Duration,proto3" json:"Duration,omitempty"`
	MapSimple  map[string]int32               `protobuf:"bytes,7,rep,name=MapSimple,proto3" json:"MapSimple,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapComplex map[int32]*CheckStatusResponse `protobuf:"bytes,8,rep,name=MapComplex,proto3" json:"MapComplex,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Bytes      []byte                         `protobuf:"bytes,9,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
}

func (x *UtilTestRequest) Reset() {
	*x = UtilTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_tld_util_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtilTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtilTestRequest) ProtoMessage() {}

func (x *UtilTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_tld_util_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtilTestRequest.ProtoReflect.Descriptor instead.
func (*UtilTestRequest) Descriptor() ([]byte, []int) {
	return file_testdata_tld_util_all_proto_rawDescGZIP(), []int{1}
}

func (x *UtilTestRequest) GetInts() []int32 {
	if x != nil {
		return x.Ints
	}
	return nil
}

func (x *UtilTestRequest) GetStructs() []*imported.Message {
	if x != nil {
		return x.Structs
	}
	return nil
}

func (x *UtilTestRequest) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *UtilTestRequest) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *UtilTestRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *UtilTestRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *UtilTestRequest) GetMapSimple() map[string]int32 {
	if x != nil {
		return x.MapSimple
	}
	return nil
}

func (x *UtilTestRequest) GetMapComplex() map[int32]*CheckStatusResponse {
	if x != nil {
		return x.MapComplex
	}
	return nil
}

func (x *UtilTestRequest) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

var File_testdata_tld_util_all_proto protoreflect.FileDescriptor

var file_testdata_tld_util_all_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x6c, 0x64, 0x2f, 0x75,
	0x74, 0x69, 0x6c, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x6c, 0x64, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5b, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0xe8,
	0x04, 0x0a, 0x0f, 0x55, 0x74, 0x69, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x49, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x3f, 0x0a, 0x07,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c,
	0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x18, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x37, 0x0a,
	0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x4f, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x55, 0x74, 0x69,
	0x6c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x55,
	0x74, 0x69, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d,
	0x61, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x19, 0x0a, 0x05, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x1a, 0x30, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x12, 0x0d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x58, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0b, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x2a, 0x38, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x1a, 0x02, 0x08, 0x00, 0x12, 0x0d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x1a,
	0x02, 0x08, 0x00, 0x12, 0x0a, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x02, 0x1a, 0x02, 0x08, 0x00, 0x1a,
	0x02, 0x18, 0x00, 0x32, 0xe2, 0x01, 0x0a, 0x04, 0x55, 0x74, 0x69, 0x6c, 0x12, 0x6d, 0x0a, 0x04,
	0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x19, 0x88, 0x02,
	0x00, 0x90, 0x02, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x28, 0x00, 0x30, 0x00, 0x12, 0x66, 0x0a, 0x0b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x88, 0x02, 0x00, 0x90, 0x02,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x32, 0x72, 0x0a, 0x09, 0x55, 0x74, 0x69, 0x6c,
	0x54, 0x65, 0x73, 0x74, 0x73, 0x12, 0x60, 0x0a, 0x08, 0x55, 0x74, 0x69, 0x6c, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x74, 0x69, 0x6c, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00,
	0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42, 0x31, 0x48, 0x01,
	0x50, 0x00, 0x5a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x6c, 0x64,
	0x2f, 0x75, 0x74, 0x69, 0x6c, 0x3b, 0x75, 0x74, 0x69, 0x6c, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00,
	0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_tld_util_all_proto_rawDescOnce sync.Once
	file_testdata_tld_util_all_proto_rawDescData = file_testdata_tld_util_all_proto_rawDesc
)

func file_testdata_tld_util_all_proto_rawDescGZIP() []byte {
	file_testdata_tld_util_all_proto_rawDescOnce.Do(func() {
		file_testdata_tld_util_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_tld_util_all_proto_rawDescData)
	})
	return file_testdata_tld_util_all_proto_rawDescData
}

var (
	file_testdata_tld_util_all_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_testdata_tld_util_all_proto_msgTypes  = make([]protoimpl.MessageInfo, 4)
	file_testdata_tld_util_all_proto_goTypes   = []interface{}{
		(Status)(0),                   // 0: testdata.v1.util.Status
		(*CheckStatusResponse)(nil),   // 1: testdata.v1.util.CheckStatusResponse
		(*UtilTestRequest)(nil),       // 2: testdata.v1.util.UtilTestRequest
		nil,                           // 3: testdata.v1.util.UtilTestRequest.MapSimpleEntry
		nil,                           // 4: testdata.v1.util.UtilTestRequest.MapComplexEntry
		(*imported.Message)(nil),      // 5: testdata.v1.util.imported.Message
		(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
		(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
		(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
	}
)

var file_testdata_tld_util_all_proto_depIdxs = []int32{
	0,  // 0: testdata.v1.util.CheckStatusResponse.Status:type_name -> testdata.v1.util.Status
	5,  // 1: testdata.v1.util.UtilTestRequest.Structs:type_name -> testdata.v1.util.imported.Message
	6,  // 2: testdata.v1.util.UtilTestRequest.Timestamp:type_name -> google.protobuf.Timestamp
	7,  // 3: testdata.v1.util.UtilTestRequest.Duration:type_name -> google.protobuf.Duration
	3,  // 4: testdata.v1.util.UtilTestRequest.MapSimple:type_name -> testdata.v1.util.UtilTestRequest.MapSimpleEntry
	4,  // 5: testdata.v1.util.UtilTestRequest.MapComplex:type_name -> testdata.v1.util.UtilTestRequest.MapComplexEntry
	1,  // 6: testdata.v1.util.UtilTestRequest.MapComplexEntry.value:type_name -> testdata.v1.util.CheckStatusResponse
	5,  // 7: testdata.v1.util.Util.Echo:input_type -> testdata.v1.util.imported.Message
	8,  // 8: testdata.v1.util.Util.CheckStatus:input_type -> google.protobuf.Empty
	2,  // 9: testdata.v1.util.UtilTests.UtilTest:input_type -> testdata.v1.util.UtilTestRequest
	5,  // 10: testdata.v1.util.Util.Echo:output_type -> testdata.v1.util.imported.Message
	1,  // 11: testdata.v1.util.Util.CheckStatus:output_type -> testdata.v1.util.CheckStatusResponse
	1,  // 12: testdata.v1.util.UtilTests.UtilTest:output_type -> testdata.v1.util.CheckStatusResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_testdata_tld_util_all_proto_init() }
func file_testdata_tld_util_all_proto_init() {
	if File_testdata_tld_util_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_tld_util_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStatusResponse); i {
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
		file_testdata_tld_util_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtilTestRequest); i {
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
			RawDescriptor: file_testdata_tld_util_all_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_testdata_tld_util_all_proto_goTypes,
		DependencyIndexes: file_testdata_tld_util_all_proto_depIdxs,
		EnumInfos:         file_testdata_tld_util_all_proto_enumTypes,
		MessageInfos:      file_testdata_tld_util_all_proto_msgTypes,
	}.Build()
	File_testdata_tld_util_all_proto = out.File
	file_testdata_tld_util_all_proto_rawDesc = nil
	file_testdata_tld_util_all_proto_goTypes = nil
	file_testdata_tld_util_all_proto_depIdxs = nil
}
