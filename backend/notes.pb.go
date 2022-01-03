// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: notes.proto

package main

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{0}
}

type Str struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Str) Reset() {
	*x = Str{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Str) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Str) ProtoMessage() {}

func (x *Str) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Str.ProtoReflect.Descriptor instead.
func (*Str) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{1}
}

func (x *Str) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type Cache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Cache) Reset() {
	*x = Cache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cache) ProtoMessage() {}

func (x *Cache) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cache.ProtoReflect.Descriptor instead.
func (*Cache) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{2}
}

func (x *Cache) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Cache) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_notes_proto protoreflect.FileDescriptor

var file_notes_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22,
	0x2f, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x32, 0x5d, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x04, 0x2e, 0x53, 0x74, 0x72,
	0x1a, 0x04, 0x2e, 0x53, 0x74, 0x72, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x06, 0x73, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x06, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x1a, 0x06, 0x2e, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x22, 0x00, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_notes_proto_rawDescOnce sync.Once
	file_notes_proto_rawDescData = file_notes_proto_rawDesc
)

func file_notes_proto_rawDescGZIP() []byte {
	file_notes_proto_rawDescOnce.Do(func() {
		file_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_notes_proto_rawDescData)
	})
	return file_notes_proto_rawDescData
}

var file_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_notes_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: Empty
	(*Str)(nil),   // 1: Str
	(*Cache)(nil), // 2: Cache
}
var file_notes_proto_depIdxs = []int32{
	1, // 0: CacheService.getKey:input_type -> Str
	2, // 1: CacheService.setKey:input_type -> Cache
	0, // 2: CacheService.clear:input_type -> Empty
	1, // 3: CacheService.getKey:output_type -> Str
	2, // 4: CacheService.setKey:output_type -> Cache
	0, // 5: CacheService.clear:output_type -> Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notes_proto_init() }
func file_notes_proto_init() {
	if File_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Str); i {
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
		file_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cache); i {
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
			RawDescriptor: file_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_proto_goTypes,
		DependencyIndexes: file_notes_proto_depIdxs,
		MessageInfos:      file_notes_proto_msgTypes,
	}.Build()
	File_notes_proto = out.File
	file_notes_proto_rawDesc = nil
	file_notes_proto_goTypes = nil
	file_notes_proto_depIdxs = nil
}