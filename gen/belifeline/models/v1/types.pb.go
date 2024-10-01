// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: belifeline/models/v1/types.proto

package mainv1

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

type DataType int32

const (
	DataType_DATA_TYPE_UNSPECIFIED DataType = 0
	DataType_DATA_TYPE_IMAGE       DataType = 1
	DataType_DATA_TYPE_CSV         DataType = 2
	DataType_DATA_TYPE_JSON        DataType = 3
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "DATA_TYPE_UNSPECIFIED",
		1: "DATA_TYPE_IMAGE",
		2: "DATA_TYPE_CSV",
		3: "DATA_TYPE_JSON",
	}
	DataType_value = map[string]int32{
		"DATA_TYPE_UNSPECIFIED": 0,
		"DATA_TYPE_IMAGE":       1,
		"DATA_TYPE_CSV":         2,
		"DATA_TYPE_JSON":        3,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_belifeline_models_v1_types_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_belifeline_models_v1_types_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{0}
}

// Common Types
type DateUntil struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before string `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  string `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *DateUntil) Reset() {
	*x = DateUntil{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateUntil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateUntil) ProtoMessage() {}

func (x *DateUntil) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateUntil.ProtoReflect.Descriptor instead.
func (*DateUntil) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *DateUntil) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

func (x *DateUntil) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Version) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Auth Types
type ApiKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ApiKey) Reset() {
	*x = ApiKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKey) ProtoMessage() {}

func (x *ApiKey) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKey.ProtoReflect.Descriptor instead.
func (*ApiKey) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *ApiKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// GeoJSON Like Types
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates []float64 `protobuf:"fixed64,1,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{4}
}

func (x *Point) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type MultiPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates []*Point `protobuf:"bytes,1,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *MultiPoint) Reset() {
	*x = MultiPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiPoint) ProtoMessage() {}

func (x *MultiPoint) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiPoint.ProtoReflect.Descriptor instead.
func (*MultiPoint) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{5}
}

func (x *MultiPoint) GetCoordinates() []*Point {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type Polygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates []*Point `protobuf:"bytes,1,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *Polygon) Reset() {
	*x = Polygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Polygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Polygon) ProtoMessage() {}

func (x *Polygon) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{6}
}

func (x *Polygon) GetCoordinates() []*Point {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type MultiPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates []*Polygon `protobuf:"bytes,2,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *MultiPolygon) Reset() {
	*x = MultiPolygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_belifeline_models_v1_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiPolygon) ProtoMessage() {}

func (x *MultiPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiPolygon.ProtoReflect.Descriptor instead.
func (*MultiPolygon) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_types_proto_rawDescGZIP(), []int{7}
}

func (x *MultiPolygon) GetCoordinates() []*Polygon {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

var File_belifeline_models_v1_types_proto protoreflect.FileDescriptor

var file_belifeline_models_v1_types_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x39, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x1f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x1a, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x29,
	0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x0a, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x22, 0x4f, 0x0a, 0x0c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c,
	0x79, 0x67, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x2a, 0x61, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x53, 0x56, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x53,
	0x4f, 0x4e, 0x10, 0x03, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6c, 0x63, 0x79, 0x6f, 0x6e, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x6b,
	0x69, 0x7a, 0x75, 0x6e, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_belifeline_models_v1_types_proto_rawDescOnce sync.Once
	file_belifeline_models_v1_types_proto_rawDescData = file_belifeline_models_v1_types_proto_rawDesc
)

func file_belifeline_models_v1_types_proto_rawDescGZIP() []byte {
	file_belifeline_models_v1_types_proto_rawDescOnce.Do(func() {
		file_belifeline_models_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_belifeline_models_v1_types_proto_rawDescData)
	})
	return file_belifeline_models_v1_types_proto_rawDescData
}

var file_belifeline_models_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_belifeline_models_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_belifeline_models_v1_types_proto_goTypes = []any{
	(DataType)(0),        // 0: belifeline.models.v1.DataType
	(*DateUntil)(nil),    // 1: belifeline.models.v1.DateUntil
	(*UUID)(nil),         // 2: belifeline.models.v1.UUID
	(*Version)(nil),      // 3: belifeline.models.v1.Version
	(*ApiKey)(nil),       // 4: belifeline.models.v1.ApiKey
	(*Point)(nil),        // 5: belifeline.models.v1.Point
	(*MultiPoint)(nil),   // 6: belifeline.models.v1.MultiPoint
	(*Polygon)(nil),      // 7: belifeline.models.v1.Polygon
	(*MultiPolygon)(nil), // 8: belifeline.models.v1.MultiPolygon
}
var file_belifeline_models_v1_types_proto_depIdxs = []int32{
	5, // 0: belifeline.models.v1.MultiPoint.coordinates:type_name -> belifeline.models.v1.Point
	5, // 1: belifeline.models.v1.Polygon.coordinates:type_name -> belifeline.models.v1.Point
	7, // 2: belifeline.models.v1.MultiPolygon.coordinates:type_name -> belifeline.models.v1.Polygon
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_belifeline_models_v1_types_proto_init() }
func file_belifeline_models_v1_types_proto_init() {
	if File_belifeline_models_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_belifeline_models_v1_types_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DateUntil); i {
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
		file_belifeline_models_v1_types_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UUID); i {
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
		file_belifeline_models_v1_types_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Version); i {
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
		file_belifeline_models_v1_types_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ApiKey); i {
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
		file_belifeline_models_v1_types_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Point); i {
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
		file_belifeline_models_v1_types_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MultiPoint); i {
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
		file_belifeline_models_v1_types_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Polygon); i {
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
		file_belifeline_models_v1_types_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*MultiPolygon); i {
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
			RawDescriptor: file_belifeline_models_v1_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_belifeline_models_v1_types_proto_goTypes,
		DependencyIndexes: file_belifeline_models_v1_types_proto_depIdxs,
		EnumInfos:         file_belifeline_models_v1_types_proto_enumTypes,
		MessageInfos:      file_belifeline_models_v1_types_proto_msgTypes,
	}.Build()
	File_belifeline_models_v1_types_proto = out.File
	file_belifeline_models_v1_types_proto_rawDesc = nil
	file_belifeline_models_v1_types_proto_goTypes = nil
	file_belifeline_models_v1_types_proto_depIdxs = nil
}
