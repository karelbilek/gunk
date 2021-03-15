// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: testdata.tld/util/imported/all.proto

package imported

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Message is a Echo message.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Msg holds a message.
	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_tld_util_imported_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_tld_util_imported_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_testdata_tld_util_imported_all_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_testdata_tld_util_imported_all_proto protoreflect.FileDescriptor

var file_testdata_tld_util_imported_all_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x6c, 0x64, 0x2f, 0x75,
	0x74, 0x69, 0x6c, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x22, 0x1a, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x3a, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_testdata_tld_util_imported_all_proto_rawDescOnce sync.Once
	file_testdata_tld_util_imported_all_proto_rawDescData = file_testdata_tld_util_imported_all_proto_rawDesc
)

func file_testdata_tld_util_imported_all_proto_rawDescGZIP() []byte {
	file_testdata_tld_util_imported_all_proto_rawDescOnce.Do(func() {
		file_testdata_tld_util_imported_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_tld_util_imported_all_proto_rawDescData)
	})
	return file_testdata_tld_util_imported_all_proto_rawDescData
}

var (
	file_testdata_tld_util_imported_all_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
	file_testdata_tld_util_imported_all_proto_goTypes  = []interface{}{
		(*Message)(nil), // 0: testdata.v1.util.imported.Message
	}
)

var file_testdata_tld_util_imported_all_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testdata_tld_util_imported_all_proto_init() }
func file_testdata_tld_util_imported_all_proto_init() {
	if File_testdata_tld_util_imported_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_tld_util_imported_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_testdata_tld_util_imported_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_tld_util_imported_all_proto_goTypes,
		DependencyIndexes: file_testdata_tld_util_imported_all_proto_depIdxs,
		MessageInfos:      file_testdata_tld_util_imported_all_proto_msgTypes,
	}.Build()
	File_testdata_tld_util_imported_all_proto = out.File
	file_testdata_tld_util_imported_all_proto_rawDesc = nil
	file_testdata_tld_util_imported_all_proto_goTypes = nil
	file_testdata_tld_util_imported_all_proto_depIdxs = nil
}
