// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: repeated.proto

package testdata

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

type RepeatedProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampList []int64                `protobuf:"varint,1,rep,packed,name=timestamp_list,json=timestampList,proto3" json:"timestamp_list,omitempty"`
	StrList       []string               `protobuf:"bytes,2,rep,name=str_list,json=strList,proto3" json:"str_list,omitempty"`
	BooleanList   []bool                 `protobuf:"varint,3,rep,packed,name=boolean_list,json=booleanList,proto3" json:"boolean_list,omitempty"`
	IntegerList   []int64                `protobuf:"varint,4,rep,packed,name=integer_list,json=integerList,proto3" json:"integer_list,omitempty"`
	FloatList     []float64              `protobuf:"fixed64,5,rep,packed,name=float_list,json=floatList,proto3" json:"float_list,omitempty"`
	RecordList    []*RepeatedRecordProto `protobuf:"bytes,6,rep,name=record_list,json=recordList,proto3" json:"record_list,omitempty"`
}

func (x *RepeatedProto) Reset() {
	*x = RepeatedProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repeated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedProto) ProtoMessage() {}

func (x *RepeatedProto) ProtoReflect() protoreflect.Message {
	mi := &file_repeated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedProto.ProtoReflect.Descriptor instead.
func (*RepeatedProto) Descriptor() ([]byte, []int) {
	return file_repeated_proto_rawDescGZIP(), []int{0}
}

func (x *RepeatedProto) GetTimestampList() []int64 {
	if x != nil {
		return x.TimestampList
	}
	return nil
}

func (x *RepeatedProto) GetStrList() []string {
	if x != nil {
		return x.StrList
	}
	return nil
}

func (x *RepeatedProto) GetBooleanList() []bool {
	if x != nil {
		return x.BooleanList
	}
	return nil
}

func (x *RepeatedProto) GetIntegerList() []int64 {
	if x != nil {
		return x.IntegerList
	}
	return nil
}

func (x *RepeatedProto) GetFloatList() []float64 {
	if x != nil {
		return x.FloatList
	}
	return nil
}

func (x *RepeatedProto) GetRecordList() []*RepeatedRecordProto {
	if x != nil {
		return x.RecordList
	}
	return nil
}

type RepeatedRecordProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Len         int64   `protobuf:"varint,1,opt,name=len,proto3" json:"len,omitempty"`
	IntegerList []int64 `protobuf:"varint,2,rep,packed,name=integer_list,json=integerList,proto3" json:"integer_list,omitempty"`
}

func (x *RepeatedRecordProto) Reset() {
	*x = RepeatedRecordProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repeated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedRecordProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedRecordProto) ProtoMessage() {}

func (x *RepeatedRecordProto) ProtoReflect() protoreflect.Message {
	mi := &file_repeated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedRecordProto.ProtoReflect.Descriptor instead.
func (*RepeatedRecordProto) Descriptor() ([]byte, []int) {
	return file_repeated_proto_rawDescGZIP(), []int{1}
}

func (x *RepeatedRecordProto) GetLen() int64 {
	if x != nil {
		return x.Len
	}
	return 0
}

func (x *RepeatedRecordProto) GetIntegerList() []int64 {
	if x != nil {
		return x.IntegerList
	}
	return nil
}

var File_repeated_proto protoreflect.FileDescriptor

var file_repeated_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xed, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x08, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6c,
	0x65, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x01, 0x52, 0x09,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x4a, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x61,
	0x76, 0x2f, 0x62, 0x71, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_repeated_proto_rawDescOnce sync.Once
	file_repeated_proto_rawDescData = file_repeated_proto_rawDesc
)

func file_repeated_proto_rawDescGZIP() []byte {
	file_repeated_proto_rawDescOnce.Do(func() {
		file_repeated_proto_rawDescData = protoimpl.X.CompressGZIP(file_repeated_proto_rawDescData)
	})
	return file_repeated_proto_rawDescData
}

var file_repeated_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_repeated_proto_goTypes = []interface{}{
	(*RepeatedProto)(nil),       // 0: RepeatedProto
	(*RepeatedRecordProto)(nil), // 1: RepeatedRecordProto
}
var file_repeated_proto_depIdxs = []int32{
	1, // 0: RepeatedProto.record_list:type_name -> RepeatedRecordProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_repeated_proto_init() }
func file_repeated_proto_init() {
	if File_repeated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repeated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedProto); i {
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
		file_repeated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedRecordProto); i {
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
			RawDescriptor: file_repeated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_repeated_proto_goTypes,
		DependencyIndexes: file_repeated_proto_depIdxs,
		MessageInfos:      file_repeated_proto_msgTypes,
	}.Build()
	File_repeated_proto = out.File
	file_repeated_proto_rawDesc = nil
	file_repeated_proto_goTypes = nil
	file_repeated_proto_depIdxs = nil
}
