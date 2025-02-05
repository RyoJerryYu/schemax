// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: schemax/plugin/plugin.proto

package plugin

import (
	schema "github.com/RyoJerryYu/schemax/proto/schemax/schema"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pluginpb "google.golang.org/protobuf/types/pluginpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TableGeneratorRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	TablesToGenerate []string               `protobuf:"bytes,1,rep,name=tables_to_generate,json=tablesToGenerate,proto3" json:"tables_to_generate,omitempty"`
	Parameter        *string                `protobuf:"bytes,2,opt,name=parameter,proto3,oneof" json:"parameter,omitempty"`
	CompilerVersion  *pluginpb.Version      `protobuf:"bytes,3,opt,name=compiler_version,json=compilerVersion,proto3" json:"compiler_version,omitempty"`
	// tables are the tables to generate.
	Tables []*schema.Table `protobuf:"bytes,4,rep,name=tables,proto3" json:"tables,omitempty"`
	// source_schema is the source schema that the tables are from.
	SourceSchema  *schema.Schema `protobuf:"bytes,5,opt,name=source_schema,json=sourceSchema,proto3" json:"source_schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TableGeneratorRequest) Reset() {
	*x = TableGeneratorRequest{}
	mi := &file_schemax_plugin_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableGeneratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableGeneratorRequest) ProtoMessage() {}

func (x *TableGeneratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schemax_plugin_plugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableGeneratorRequest.ProtoReflect.Descriptor instead.
func (*TableGeneratorRequest) Descriptor() ([]byte, []int) {
	return file_schemax_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *TableGeneratorRequest) GetTablesToGenerate() []string {
	if x != nil {
		return x.TablesToGenerate
	}
	return nil
}

func (x *TableGeneratorRequest) GetParameter() string {
	if x != nil && x.Parameter != nil {
		return *x.Parameter
	}
	return ""
}

func (x *TableGeneratorRequest) GetCompilerVersion() *pluginpb.Version {
	if x != nil {
		return x.CompilerVersion
	}
	return nil
}

func (x *TableGeneratorRequest) GetTables() []*schema.Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

func (x *TableGeneratorRequest) GetSourceSchema() *schema.Schema {
	if x != nil {
		return x.SourceSchema
	}
	return nil
}

type TableGeneratorResponse struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Error         *string                        `protobuf:"bytes,1,opt,name=error,proto3,oneof" json:"error,omitempty"`
	Files         []*TableGeneratorResponse_File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TableGeneratorResponse) Reset() {
	*x = TableGeneratorResponse{}
	mi := &file_schemax_plugin_plugin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableGeneratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableGeneratorResponse) ProtoMessage() {}

func (x *TableGeneratorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schemax_plugin_plugin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableGeneratorResponse.ProtoReflect.Descriptor instead.
func (*TableGeneratorResponse) Descriptor() ([]byte, []int) {
	return file_schemax_plugin_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *TableGeneratorResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *TableGeneratorResponse) GetFiles() []*TableGeneratorResponse_File {
	if x != nil {
		return x.Files
	}
	return nil
}

type TableGeneratorResponse_File struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the file name, relative to the output directory.
	// should not contain . or .. as path elements, so the generated file
	// will be in the output directory or a subdirectory.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content       string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // the file content
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TableGeneratorResponse_File) Reset() {
	*x = TableGeneratorResponse_File{}
	mi := &file_schemax_plugin_plugin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableGeneratorResponse_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableGeneratorResponse_File) ProtoMessage() {}

func (x *TableGeneratorResponse_File) ProtoReflect() protoreflect.Message {
	mi := &file_schemax_plugin_plugin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableGeneratorResponse_File.ProtoReflect.Descriptor instead.
func (*TableGeneratorResponse_File) Descriptor() ([]byte, []int) {
	return file_schemax_plugin_plugin_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TableGeneratorResponse_File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableGeneratorResponse_File) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_schemax_plugin_plugin_proto protoreflect.FileDescriptor

var file_schemax_plugin_plugin_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x25, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb0, 0x02, 0x0a, 0x15, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x54,
	0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x22, 0xb6, 0x01, 0x0a, 0x16, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x34, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x79, 0x6f, 0x4a,
	0x65, 0x72, 0x72, 0x79, 0x59, 0x75, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_schemax_plugin_plugin_proto_rawDescOnce sync.Once
	file_schemax_plugin_plugin_proto_rawDescData []byte
)

func file_schemax_plugin_plugin_proto_rawDescGZIP() []byte {
	file_schemax_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_schemax_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_schemax_plugin_plugin_proto_rawDesc), len(file_schemax_plugin_plugin_proto_rawDesc)))
	})
	return file_schemax_plugin_plugin_proto_rawDescData
}

var file_schemax_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schemax_plugin_plugin_proto_goTypes = []any{
	(*TableGeneratorRequest)(nil),       // 0: schemax.plugin.TableGeneratorRequest
	(*TableGeneratorResponse)(nil),      // 1: schemax.plugin.TableGeneratorResponse
	(*TableGeneratorResponse_File)(nil), // 2: schemax.plugin.TableGeneratorResponse.File
	(*pluginpb.Version)(nil),            // 3: google.protobuf.compiler.Version
	(*schema.Table)(nil),                // 4: schemax.schema.Table
	(*schema.Schema)(nil),               // 5: schemax.schema.Schema
}
var file_schemax_plugin_plugin_proto_depIdxs = []int32{
	3, // 0: schemax.plugin.TableGeneratorRequest.compiler_version:type_name -> google.protobuf.compiler.Version
	4, // 1: schemax.plugin.TableGeneratorRequest.tables:type_name -> schemax.schema.Table
	5, // 2: schemax.plugin.TableGeneratorRequest.source_schema:type_name -> schemax.schema.Schema
	2, // 3: schemax.plugin.TableGeneratorResponse.files:type_name -> schemax.plugin.TableGeneratorResponse.File
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_schemax_plugin_plugin_proto_init() }
func file_schemax_plugin_plugin_proto_init() {
	if File_schemax_plugin_plugin_proto != nil {
		return
	}
	file_schemax_plugin_plugin_proto_msgTypes[0].OneofWrappers = []any{}
	file_schemax_plugin_plugin_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_schemax_plugin_plugin_proto_rawDesc), len(file_schemax_plugin_plugin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schemax_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_schemax_plugin_plugin_proto_depIdxs,
		MessageInfos:      file_schemax_plugin_plugin_proto_msgTypes,
	}.Build()
	File_schemax_plugin_plugin_proto = out.File
	file_schemax_plugin_plugin_proto_goTypes = nil
	file_schemax_plugin_plugin_proto_depIdxs = nil
}
