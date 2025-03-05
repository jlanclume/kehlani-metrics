// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: metric_event.proto

package proto

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

type MetricType int32

const (
	MetricType_GAUGE     MetricType = 0
	MetricType_COUNTER   MetricType = 1
	MetricType_HISTOGRAM MetricType = 2
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "GAUGE",
		1: "COUNTER",
		2: "HISTOGRAM",
	}
	MetricType_value = map[string]int32{
		"GAUGE":     0,
		"COUNTER":   1,
		"HISTOGRAM": 2,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_metric_event_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_metric_event_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_metric_event_proto_rawDescGZIP(), []int{0}
}

type MetricEvent struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	TenantId string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	ClientId string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Timestamp of the event in Unix format (milliseconds)
	EventTimestamp int64 `protobuf:"varint,3,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
	// The name of the metric being reported
	MetricName string `protobuf:"bytes,4,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// The value of the metric
	MetricValue float64 `protobuf:"fixed64,5,opt,name=metric_value,json=metricValue,proto3" json:"metric_value,omitempty"`
	// The type of the metric
	MetricType    MetricType `protobuf:"varint,6,opt,name=metric_type,json=metricType,proto3,enum=proto.MetricType" json:"metric_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricEvent) Reset() {
	*x = MetricEvent{}
	mi := &file_metric_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricEvent) ProtoMessage() {}

func (x *MetricEvent) ProtoReflect() protoreflect.Message {
	mi := &file_metric_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricEvent.ProtoReflect.Descriptor instead.
func (*MetricEvent) Descriptor() ([]byte, []int) {
	return file_metric_event_proto_rawDescGZIP(), []int{0}
}

func (x *MetricEvent) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *MetricEvent) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *MetricEvent) GetEventTimestamp() int64 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

func (x *MetricEvent) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *MetricEvent) GetMetricValue() float64 {
	if x != nil {
		return x.MetricValue
	}
	return 0
}

func (x *MetricEvent) GetMetricType() MetricType {
	if x != nil {
		return x.MetricType
	}
	return MetricType_GAUGE
}

type MetricEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Event         *MetricEvent           `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricEventRequest) Reset() {
	*x = MetricEventRequest{}
	mi := &file_metric_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricEventRequest) ProtoMessage() {}

func (x *MetricEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metric_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricEventRequest.ProtoReflect.Descriptor instead.
func (*MetricEventRequest) Descriptor() ([]byte, []int) {
	return file_metric_event_proto_rawDescGZIP(), []int{1}
}

func (x *MetricEventRequest) GetEvent() *MetricEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type MetricEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricEventResponse) Reset() {
	*x = MetricEventResponse{}
	mi := &file_metric_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricEventResponse) ProtoMessage() {}

func (x *MetricEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metric_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricEventResponse.ProtoReflect.Descriptor instead.
func (*MetricEventResponse) Descriptor() ([]byte, []int) {
	return file_metric_event_proto_rawDescGZIP(), []int{2}
}

func (x *MetricEventResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_metric_event_proto protoreflect.FileDescriptor

var file_metric_event_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x0b,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x33, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41, 0x55, 0x47, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48,
	0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x02, 0x32, 0x5e, 0x0a, 0x12, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x48, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x6b, 0x65, 0x68, 0x6c, 0x61,
	0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_metric_event_proto_rawDescOnce sync.Once
	file_metric_event_proto_rawDescData []byte
)

func file_metric_event_proto_rawDescGZIP() []byte {
	file_metric_event_proto_rawDescOnce.Do(func() {
		file_metric_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metric_event_proto_rawDesc), len(file_metric_event_proto_rawDesc)))
	})
	return file_metric_event_proto_rawDescData
}

var file_metric_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metric_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_metric_event_proto_goTypes = []any{
	(MetricType)(0),             // 0: proto.MetricType
	(*MetricEvent)(nil),         // 1: proto.MetricEvent
	(*MetricEventRequest)(nil),  // 2: proto.MetricEventRequest
	(*MetricEventResponse)(nil), // 3: proto.MetricEventResponse
}
var file_metric_event_proto_depIdxs = []int32{
	0, // 0: proto.MetricEvent.metric_type:type_name -> proto.MetricType
	1, // 1: proto.MetricEventRequest.event:type_name -> proto.MetricEvent
	2, // 2: proto.MetricEventService.SaveMetricEvent:input_type -> proto.MetricEventRequest
	3, // 3: proto.MetricEventService.SaveMetricEvent:output_type -> proto.MetricEventResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_metric_event_proto_init() }
func file_metric_event_proto_init() {
	if File_metric_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metric_event_proto_rawDesc), len(file_metric_event_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metric_event_proto_goTypes,
		DependencyIndexes: file_metric_event_proto_depIdxs,
		EnumInfos:         file_metric_event_proto_enumTypes,
		MessageInfos:      file_metric_event_proto_msgTypes,
	}.Build()
	File_metric_event_proto = out.File
	file_metric_event_proto_goTypes = nil
	file_metric_event_proto_depIdxs = nil
}
