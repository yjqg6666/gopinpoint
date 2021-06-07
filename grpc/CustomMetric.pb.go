// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: v1/CustomMetric.proto

package grpc

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

type PCustomMetricMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp       []int64          `protobuf:"varint,1,rep,packed,name=timestamp,proto3" json:"timestamp,omitempty"`
	CollectInterval []int64          `protobuf:"varint,2,rep,packed,name=collectInterval,proto3" json:"collectInterval,omitempty"`
	CustomMetrics   []*PCustomMetric `protobuf:"bytes,3,rep,name=customMetrics,proto3" json:"customMetrics,omitempty"`
}

func (x *PCustomMetricMessage) Reset() {
	*x = PCustomMetricMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCustomMetricMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCustomMetricMessage) ProtoMessage() {}

func (x *PCustomMetricMessage) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCustomMetricMessage.ProtoReflect.Descriptor instead.
func (*PCustomMetricMessage) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{0}
}

func (x *PCustomMetricMessage) GetTimestamp() []int64 {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *PCustomMetricMessage) GetCollectInterval() []int64 {
	if x != nil {
		return x.CollectInterval
	}
	return nil
}

func (x *PCustomMetricMessage) GetCustomMetrics() []*PCustomMetric {
	if x != nil {
		return x.CustomMetrics
	}
	return nil
}

type PCustomMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Field:
	//	*PCustomMetric_IntCountMetric
	//	*PCustomMetric_LongCountMetric
	//	*PCustomMetric_IntGaugeMetric
	//	*PCustomMetric_LongGaugeMetric
	//	*PCustomMetric_DoubleGaugeMetric
	Field isPCustomMetric_Field `protobuf_oneof:"field"`
}

func (x *PCustomMetric) Reset() {
	*x = PCustomMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCustomMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCustomMetric) ProtoMessage() {}

func (x *PCustomMetric) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCustomMetric.ProtoReflect.Descriptor instead.
func (*PCustomMetric) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{1}
}

func (m *PCustomMetric) GetField() isPCustomMetric_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (x *PCustomMetric) GetIntCountMetric() *PIntCountMetric {
	if x, ok := x.GetField().(*PCustomMetric_IntCountMetric); ok {
		return x.IntCountMetric
	}
	return nil
}

func (x *PCustomMetric) GetLongCountMetric() *PLongCountMetric {
	if x, ok := x.GetField().(*PCustomMetric_LongCountMetric); ok {
		return x.LongCountMetric
	}
	return nil
}

func (x *PCustomMetric) GetIntGaugeMetric() *PIntGaugeMetric {
	if x, ok := x.GetField().(*PCustomMetric_IntGaugeMetric); ok {
		return x.IntGaugeMetric
	}
	return nil
}

func (x *PCustomMetric) GetLongGaugeMetric() *PLongGaugeMetric {
	if x, ok := x.GetField().(*PCustomMetric_LongGaugeMetric); ok {
		return x.LongGaugeMetric
	}
	return nil
}

func (x *PCustomMetric) GetDoubleGaugeMetric() *PDouleGaugeMetric {
	if x, ok := x.GetField().(*PCustomMetric_DoubleGaugeMetric); ok {
		return x.DoubleGaugeMetric
	}
	return nil
}

type isPCustomMetric_Field interface {
	isPCustomMetric_Field()
}

type PCustomMetric_IntCountMetric struct {
	IntCountMetric *PIntCountMetric `protobuf:"bytes,1,opt,name=intCountMetric,proto3,oneof"`
}

type PCustomMetric_LongCountMetric struct {
	LongCountMetric *PLongCountMetric `protobuf:"bytes,2,opt,name=longCountMetric,proto3,oneof"`
}

type PCustomMetric_IntGaugeMetric struct {
	IntGaugeMetric *PIntGaugeMetric `protobuf:"bytes,3,opt,name=intGaugeMetric,proto3,oneof"`
}

type PCustomMetric_LongGaugeMetric struct {
	LongGaugeMetric *PLongGaugeMetric `protobuf:"bytes,4,opt,name=longGaugeMetric,proto3,oneof"`
}

type PCustomMetric_DoubleGaugeMetric struct {
	DoubleGaugeMetric *PDouleGaugeMetric `protobuf:"bytes,5,opt,name=doubleGaugeMetric,proto3,oneof"`
}

func (*PCustomMetric_IntCountMetric) isPCustomMetric_Field() {}

func (*PCustomMetric_LongCountMetric) isPCustomMetric_Field() {}

func (*PCustomMetric_IntGaugeMetric) isPCustomMetric_Field() {}

func (*PCustomMetric_LongGaugeMetric) isPCustomMetric_Field() {}

func (*PCustomMetric_DoubleGaugeMetric) isPCustomMetric_Field() {}

type PIntCountMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values []*PIntValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *PIntCountMetric) Reset() {
	*x = PIntCountMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIntCountMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIntCountMetric) ProtoMessage() {}

func (x *PIntCountMetric) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIntCountMetric.ProtoReflect.Descriptor instead.
func (*PIntCountMetric) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{2}
}

func (x *PIntCountMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PIntCountMetric) GetValues() []*PIntValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type PLongCountMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values []*PLongValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *PLongCountMetric) Reset() {
	*x = PLongCountMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLongCountMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLongCountMetric) ProtoMessage() {}

func (x *PLongCountMetric) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLongCountMetric.ProtoReflect.Descriptor instead.
func (*PLongCountMetric) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{3}
}

func (x *PLongCountMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PLongCountMetric) GetValues() []*PLongValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type PIntGaugeMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values []*PIntValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *PIntGaugeMetric) Reset() {
	*x = PIntGaugeMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIntGaugeMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIntGaugeMetric) ProtoMessage() {}

func (x *PIntGaugeMetric) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIntGaugeMetric.ProtoReflect.Descriptor instead.
func (*PIntGaugeMetric) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{4}
}

func (x *PIntGaugeMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PIntGaugeMetric) GetValues() []*PIntValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type PLongGaugeMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values []*PLongValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *PLongGaugeMetric) Reset() {
	*x = PLongGaugeMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLongGaugeMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLongGaugeMetric) ProtoMessage() {}

func (x *PLongGaugeMetric) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLongGaugeMetric.ProtoReflect.Descriptor instead.
func (*PLongGaugeMetric) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{5}
}

func (x *PLongGaugeMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PLongGaugeMetric) GetValues() []*PLongValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type PDouleGaugeMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values []*PDoubleValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *PDouleGaugeMetric) Reset() {
	*x = PDouleGaugeMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDouleGaugeMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDouleGaugeMetric) ProtoMessage() {}

func (x *PDouleGaugeMetric) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDouleGaugeMetric.ProtoReflect.Descriptor instead.
func (*PDouleGaugeMetric) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{6}
}

func (x *PDouleGaugeMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PDouleGaugeMetric) GetValues() []*PDoubleValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type PIntValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	IsNotSet bool  `protobuf:"varint,2,opt,name=isNotSet,proto3" json:"isNotSet,omitempty"`
}

func (x *PIntValue) Reset() {
	*x = PIntValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIntValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIntValue) ProtoMessage() {}

func (x *PIntValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIntValue.ProtoReflect.Descriptor instead.
func (*PIntValue) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{7}
}

func (x *PIntValue) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PIntValue) GetIsNotSet() bool {
	if x != nil {
		return x.IsNotSet
	}
	return false
}

type PLongValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	IsNotSet bool  `protobuf:"varint,2,opt,name=isNotSet,proto3" json:"isNotSet,omitempty"`
}

func (x *PLongValue) Reset() {
	*x = PLongValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLongValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLongValue) ProtoMessage() {}

func (x *PLongValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLongValue.ProtoReflect.Descriptor instead.
func (*PLongValue) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{8}
}

func (x *PLongValue) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PLongValue) GetIsNotSet() bool {
	if x != nil {
		return x.IsNotSet
	}
	return false
}

type PDoubleValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	IsNotSet bool    `protobuf:"varint,2,opt,name=isNotSet,proto3" json:"isNotSet,omitempty"`
}

func (x *PDoubleValue) Reset() {
	*x = PDoubleValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_CustomMetric_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDoubleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDoubleValue) ProtoMessage() {}

func (x *PDoubleValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_CustomMetric_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDoubleValue.ProtoReflect.Descriptor instead.
func (*PDoubleValue) Descriptor() ([]byte, []int) {
	return file_v1_CustomMetric_proto_rawDescGZIP(), []int{9}
}

func (x *PDoubleValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PDoubleValue) GetIsNotSet() bool {
	if x != nil {
		return x.IsNotSet
	}
	return false
}

var File_v1_CustomMetric_proto protoreflect.FileDescriptor

var file_v1_CustomMetric_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x97, 0x01, 0x0a, 0x14,
	0x50, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xe1, 0x02, 0x0a, 0x0d, 0x50, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x3d, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x49, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x40, 0x0a, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x3d, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x49, 0x6e, 0x74, 0x47, 0x61, 0x75, 0x67, 0x65, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x47, 0x61, 0x75, 0x67,
	0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x40, 0x0a, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x61, 0x75, 0x67, 0x65,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x47, 0x61,
	0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x47, 0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x44, 0x6f, 0x75, 0x6c, 0x65,
	0x47, 0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x11, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x47, 0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x42, 0x07, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x4c, 0x0a, 0x0f, 0x50, 0x49, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x10, 0x50, 0x4c, 0x6f, 0x6e, 0x67,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x4c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x0f, 0x50, 0x49, 0x6e, 0x74, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x10, 0x50, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x61,
	0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x4c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x11, 0x50, 0x44, 0x6f, 0x75, 0x6c, 0x65, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x09, 0x50, 0x49, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x4e, 0x6f, 0x74, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x4e, 0x6f, 0x74, 0x53, 0x65, 0x74, 0x22, 0x3e, 0x0a, 0x0a, 0x50, 0x4c, 0x6f, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x4e, 0x6f, 0x74, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x4e, 0x6f, 0x74, 0x53, 0x65, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x50, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x4e, 0x6f, 0x74, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x4e, 0x6f, 0x74, 0x53, 0x65, 0x74, 0x42, 0x63, 0x0a, 0x21, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x61, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x11,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x6a, 0x71, 0x67, 0x36, 0x36, 0x36, 0x36, 0x2f, 0x67, 0x6f, 0x70, 0x69, 0x6e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x03, 0x50, 0x49, 0x4e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_CustomMetric_proto_rawDescOnce sync.Once
	file_v1_CustomMetric_proto_rawDescData = file_v1_CustomMetric_proto_rawDesc
)

func file_v1_CustomMetric_proto_rawDescGZIP() []byte {
	file_v1_CustomMetric_proto_rawDescOnce.Do(func() {
		file_v1_CustomMetric_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_CustomMetric_proto_rawDescData)
	})
	return file_v1_CustomMetric_proto_rawDescData
}

var file_v1_CustomMetric_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_CustomMetric_proto_goTypes = []interface{}{
	(*PCustomMetricMessage)(nil), // 0: v1.PCustomMetricMessage
	(*PCustomMetric)(nil),        // 1: v1.PCustomMetric
	(*PIntCountMetric)(nil),      // 2: v1.PIntCountMetric
	(*PLongCountMetric)(nil),     // 3: v1.PLongCountMetric
	(*PIntGaugeMetric)(nil),      // 4: v1.PIntGaugeMetric
	(*PLongGaugeMetric)(nil),     // 5: v1.PLongGaugeMetric
	(*PDouleGaugeMetric)(nil),    // 6: v1.PDouleGaugeMetric
	(*PIntValue)(nil),            // 7: v1.PIntValue
	(*PLongValue)(nil),           // 8: v1.PLongValue
	(*PDoubleValue)(nil),         // 9: v1.PDoubleValue
}
var file_v1_CustomMetric_proto_depIdxs = []int32{
	1,  // 0: v1.PCustomMetricMessage.customMetrics:type_name -> v1.PCustomMetric
	2,  // 1: v1.PCustomMetric.intCountMetric:type_name -> v1.PIntCountMetric
	3,  // 2: v1.PCustomMetric.longCountMetric:type_name -> v1.PLongCountMetric
	4,  // 3: v1.PCustomMetric.intGaugeMetric:type_name -> v1.PIntGaugeMetric
	5,  // 4: v1.PCustomMetric.longGaugeMetric:type_name -> v1.PLongGaugeMetric
	6,  // 5: v1.PCustomMetric.doubleGaugeMetric:type_name -> v1.PDouleGaugeMetric
	7,  // 6: v1.PIntCountMetric.values:type_name -> v1.PIntValue
	8,  // 7: v1.PLongCountMetric.values:type_name -> v1.PLongValue
	7,  // 8: v1.PIntGaugeMetric.values:type_name -> v1.PIntValue
	8,  // 9: v1.PLongGaugeMetric.values:type_name -> v1.PLongValue
	9,  // 10: v1.PDouleGaugeMetric.values:type_name -> v1.PDoubleValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_v1_CustomMetric_proto_init() }
func file_v1_CustomMetric_proto_init() {
	if File_v1_CustomMetric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_CustomMetric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCustomMetricMessage); i {
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
		file_v1_CustomMetric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCustomMetric); i {
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
		file_v1_CustomMetric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIntCountMetric); i {
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
		file_v1_CustomMetric_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLongCountMetric); i {
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
		file_v1_CustomMetric_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIntGaugeMetric); i {
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
		file_v1_CustomMetric_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLongGaugeMetric); i {
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
		file_v1_CustomMetric_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDouleGaugeMetric); i {
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
		file_v1_CustomMetric_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIntValue); i {
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
		file_v1_CustomMetric_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLongValue); i {
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
		file_v1_CustomMetric_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDoubleValue); i {
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
	file_v1_CustomMetric_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PCustomMetric_IntCountMetric)(nil),
		(*PCustomMetric_LongCountMetric)(nil),
		(*PCustomMetric_IntGaugeMetric)(nil),
		(*PCustomMetric_LongGaugeMetric)(nil),
		(*PCustomMetric_DoubleGaugeMetric)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_CustomMetric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_CustomMetric_proto_goTypes,
		DependencyIndexes: file_v1_CustomMetric_proto_depIdxs,
		MessageInfos:      file_v1_CustomMetric_proto_msgTypes,
	}.Build()
	File_v1_CustomMetric_proto = out.File
	file_v1_CustomMetric_proto_rawDesc = nil
	file_v1_CustomMetric_proto_goTypes = nil
	file_v1_CustomMetric_proto_depIdxs = nil
}
