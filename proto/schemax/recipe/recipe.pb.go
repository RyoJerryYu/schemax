// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: schemax/recipe/recipe.proto

package recipe

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Recipe struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dsn           string                 `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty"`
	Schema        string                 `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"` // optional, if not set, use schema from dsn
	TableSets     []*TableSet            `protobuf:"bytes,3,rep,name=table_sets,json=tableSets,proto3" json:"table_sets,omitempty"`
	Plugins       []*PluginSetting       `protobuf:"bytes,4,rep,name=plugins,proto3" json:"plugins,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	mi := &file_schemax_recipe_recipe_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_schemax_recipe_recipe_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_schemax_recipe_recipe_proto_rawDescGZIP(), []int{0}
}

func (x *Recipe) GetDsn() string {
	if x != nil {
		return x.Dsn
	}
	return ""
}

func (x *Recipe) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Recipe) GetTableSets() []*TableSet {
	if x != nil {
		return x.TableSets
	}
	return nil
}

func (x *Recipe) GetPlugins() []*PluginSetting {
	if x != nil {
		return x.Plugins
	}
	return nil
}

type TableSet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SetName       string                 `protobuf:"bytes,1,opt,name=set_name,json=setName,proto3" json:"set_name,omitempty"`
	Tables        []string               `protobuf:"bytes,2,rep,name=tables,proto3" json:"tables,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TableSet) Reset() {
	*x = TableSet{}
	mi := &file_schemax_recipe_recipe_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableSet) ProtoMessage() {}

func (x *TableSet) ProtoReflect() protoreflect.Message {
	mi := &file_schemax_recipe_recipe_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableSet.ProtoReflect.Descriptor instead.
func (*TableSet) Descriptor() ([]byte, []int) {
	return file_schemax_recipe_recipe_proto_rawDescGZIP(), []int{1}
}

func (x *TableSet) GetSetName() string {
	if x != nil {
		return x.SetName
	}
	return ""
}

func (x *TableSet) GetTables() []string {
	if x != nil {
		return x.Tables
	}
	return nil
}

type PluginSetting struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Plugin        string                 `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Out           string                 `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`   // output directory path template, can use {set_name} , all generated files will be in this directory
	Opts          []string               `protobuf:"bytes,3,rep,name=opts,proto3" json:"opts,omitempty"` // plugin options
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginSetting) Reset() {
	*x = PluginSetting{}
	mi := &file_schemax_recipe_recipe_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginSetting) ProtoMessage() {}

func (x *PluginSetting) ProtoReflect() protoreflect.Message {
	mi := &file_schemax_recipe_recipe_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginSetting.ProtoReflect.Descriptor instead.
func (*PluginSetting) Descriptor() ([]byte, []int) {
	return file_schemax_recipe_recipe_proto_rawDescGZIP(), []int{2}
}

func (x *PluginSetting) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *PluginSetting) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

func (x *PluginSetting) GetOpts() []string {
	if x != nil {
		return x.Opts
	}
	return nil
}

var File_schemax_recipe_recipe_proto protoreflect.FileDescriptor

var file_schemax_recipe_recipe_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x22, 0xa4, 0x01,
	0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x78,
	0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x74,
	0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x78, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x22, 0x3d, 0x0a, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70,
	0x74, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x52, 0x79, 0x6f, 0x4a, 0x65, 0x72, 0x72, 0x79, 0x59, 0x75, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x78, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_schemax_recipe_recipe_proto_rawDescOnce sync.Once
	file_schemax_recipe_recipe_proto_rawDescData []byte
)

func file_schemax_recipe_recipe_proto_rawDescGZIP() []byte {
	file_schemax_recipe_recipe_proto_rawDescOnce.Do(func() {
		file_schemax_recipe_recipe_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_schemax_recipe_recipe_proto_rawDesc), len(file_schemax_recipe_recipe_proto_rawDesc)))
	})
	return file_schemax_recipe_recipe_proto_rawDescData
}

var file_schemax_recipe_recipe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schemax_recipe_recipe_proto_goTypes = []any{
	(*Recipe)(nil),        // 0: schemax.recipe.Recipe
	(*TableSet)(nil),      // 1: schemax.recipe.TableSet
	(*PluginSetting)(nil), // 2: schemax.recipe.PluginSetting
}
var file_schemax_recipe_recipe_proto_depIdxs = []int32{
	1, // 0: schemax.recipe.Recipe.table_sets:type_name -> schemax.recipe.TableSet
	2, // 1: schemax.recipe.Recipe.plugins:type_name -> schemax.recipe.PluginSetting
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schemax_recipe_recipe_proto_init() }
func file_schemax_recipe_recipe_proto_init() {
	if File_schemax_recipe_recipe_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_schemax_recipe_recipe_proto_rawDesc), len(file_schemax_recipe_recipe_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schemax_recipe_recipe_proto_goTypes,
		DependencyIndexes: file_schemax_recipe_recipe_proto_depIdxs,
		MessageInfos:      file_schemax_recipe_recipe_proto_msgTypes,
	}.Build()
	File_schemax_recipe_recipe_proto = out.File
	file_schemax_recipe_recipe_proto_goTypes = nil
	file_schemax_recipe_recipe_proto_depIdxs = nil
}
