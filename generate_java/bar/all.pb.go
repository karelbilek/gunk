// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: testdata.tld/util/bar/all.proto

package bar

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bar is a string
	Bar string `protobuf:"bytes,1,opt,name=Bar,json=bar,proto3" json:"Bar,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_tld_util_bar_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_tld_util_bar_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_testdata_tld_util_bar_all_proto_rawDescGZIP(), []int{0}
}

func (x *Foo) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

var File_testdata_tld_util_bar_all_proto protoreflect.FileDescriptor

var file_testdata_tld_util_bar_all_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x6c, 0x64, 0x2f, 0x75,
	0x74, 0x69, 0x6c, 0x2f, 0x62, 0x61, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x62, 0x61, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2b, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x1c, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52,
	0x03, 0x62, 0x61, 0x72, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32, 0x71, 0x0a, 0x0a,
	0x46, 0x6f, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x62,
	0x61, 0x72, 0x2e, 0x46, 0x6f, 0x6f, 0x22, 0x30, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92, 0x41,
	0x18, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x08, 0x47, 0x65, 0x74, 0x73, 0x20, 0x66, 0x6f, 0x6f,
	0x4a, 0x07, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12,
	0x07, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42,
	0x30, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x74, 0x6c, 0x64, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x62, 0x61, 0x72, 0x80, 0x01, 0x00, 0x88,
	0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_tld_util_bar_all_proto_rawDescOnce sync.Once
	file_testdata_tld_util_bar_all_proto_rawDescData = file_testdata_tld_util_bar_all_proto_rawDesc
)

func file_testdata_tld_util_bar_all_proto_rawDescGZIP() []byte {
	file_testdata_tld_util_bar_all_proto_rawDescOnce.Do(func() {
		file_testdata_tld_util_bar_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_tld_util_bar_all_proto_rawDescData)
	})
	return file_testdata_tld_util_bar_all_proto_rawDescData
}

var file_testdata_tld_util_bar_all_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_testdata_tld_util_bar_all_proto_goTypes = []interface{}{
	(*Foo)(nil),         // 0: bar.Foo
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_testdata_tld_util_bar_all_proto_depIdxs = []int32{
	1, // 0: bar.FooService.GetFoo:input_type -> google.protobuf.Empty
	0, // 1: bar.FooService.GetFoo:output_type -> bar.Foo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testdata_tld_util_bar_all_proto_init() }
func file_testdata_tld_util_bar_all_proto_init() {
	if File_testdata_tld_util_bar_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_tld_util_bar_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
			RawDescriptor: file_testdata_tld_util_bar_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testdata_tld_util_bar_all_proto_goTypes,
		DependencyIndexes: file_testdata_tld_util_bar_all_proto_depIdxs,
		MessageInfos:      file_testdata_tld_util_bar_all_proto_msgTypes,
	}.Build()
	File_testdata_tld_util_bar_all_proto = out.File
	file_testdata_tld_util_bar_all_proto_rawDesc = nil
	file_testdata_tld_util_bar_all_proto_goTypes = nil
	file_testdata_tld_util_bar_all_proto_depIdxs = nil
}
