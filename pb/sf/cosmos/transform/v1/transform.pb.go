// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: sf/cosmos/transform/v1/transform.proto

package pbtransform

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

type CombinedFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypeFilters   []*EventTypeFilter   `protobuf:"bytes,1,rep,name=event_type_filters,json=eventTypeFilters,proto3" json:"event_type_filters,omitempty"`
	EventOriginFilters []*EventOriginFilter `protobuf:"bytes,2,rep,name=event_origin_filters,json=eventOriginFilters,proto3" json:"event_origin_filters,omitempty"`
	MessageTypeFilters []*MessageTypeFilter `protobuf:"bytes,3,rep,name=message_type_filters,json=messageTypeFilters,proto3" json:"message_type_filters,omitempty"`
}

func (x *CombinedFilter) Reset() {
	*x = CombinedFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombinedFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombinedFilter) ProtoMessage() {}

func (x *CombinedFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombinedFilter.ProtoReflect.Descriptor instead.
func (*CombinedFilter) Descriptor() ([]byte, []int) {
	return file_sf_cosmos_transform_v1_transform_proto_rawDescGZIP(), []int{0}
}

func (x *CombinedFilter) GetEventTypeFilters() []*EventTypeFilter {
	if x != nil {
		return x.EventTypeFilters
	}
	return nil
}

func (x *CombinedFilter) GetEventOriginFilters() []*EventOriginFilter {
	if x != nil {
		return x.EventOriginFilters
	}
	return nil
}

func (x *CombinedFilter) GetMessageTypeFilters() []*MessageTypeFilter {
	if x != nil {
		return x.MessageTypeFilters
	}
	return nil
}

type EventTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypes []string `protobuf:"bytes,1,rep,name=event_types,json=eventTypes,proto3" json:"event_types,omitempty"`
}

func (x *EventTypeFilter) Reset() {
	*x = EventTypeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeFilter) ProtoMessage() {}

func (x *EventTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeFilter.ProtoReflect.Descriptor instead.
func (*EventTypeFilter) Descriptor() ([]byte, []int) {
	return file_sf_cosmos_transform_v1_transform_proto_rawDescGZIP(), []int{1}
}

func (x *EventTypeFilter) GetEventTypes() []string {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

type EventOriginFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventOrigins []string `protobuf:"bytes,1,rep,name=event_origins,json=eventOrigins,proto3" json:"event_origins,omitempty"`
}

func (x *EventOriginFilter) Reset() {
	*x = EventOriginFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventOriginFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventOriginFilter) ProtoMessage() {}

func (x *EventOriginFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventOriginFilter.ProtoReflect.Descriptor instead.
func (*EventOriginFilter) Descriptor() ([]byte, []int) {
	return file_sf_cosmos_transform_v1_transform_proto_rawDescGZIP(), []int{2}
}

func (x *EventOriginFilter) GetEventOrigins() []string {
	if x != nil {
		return x.EventOrigins
	}
	return nil
}

type MessageTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTypes []string `protobuf:"bytes,1,rep,name=message_types,json=messageTypes,proto3" json:"message_types,omitempty"`
}

func (x *MessageTypeFilter) Reset() {
	*x = MessageTypeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTypeFilter) ProtoMessage() {}

func (x *MessageTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sf_cosmos_transform_v1_transform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTypeFilter.ProtoReflect.Descriptor instead.
func (*MessageTypeFilter) Descriptor() ([]byte, []int) {
	return file_sf_cosmos_transform_v1_transform_proto_rawDescGZIP(), []int{3}
}

func (x *MessageTypeFilter) GetMessageTypes() []string {
	if x != nil {
		return x.MessageTypes
	}
	return nil
}

var File_sf_cosmos_transform_v1_transform_proto protoreflect.FileDescriptor

var file_sf_cosmos_transform_v1_transform_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x66, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x66, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x22, 0xa1, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x73, 0x66, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x5b, 0x0a, 0x14, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x66, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x5b, 0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x66, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x22, 0x32, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x50, 0x5a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x76, 0x31, 0x3b, 0x70, 0x62, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_cosmos_transform_v1_transform_proto_rawDescOnce sync.Once
	file_sf_cosmos_transform_v1_transform_proto_rawDescData = file_sf_cosmos_transform_v1_transform_proto_rawDesc
)

func file_sf_cosmos_transform_v1_transform_proto_rawDescGZIP() []byte {
	file_sf_cosmos_transform_v1_transform_proto_rawDescOnce.Do(func() {
		file_sf_cosmos_transform_v1_transform_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_cosmos_transform_v1_transform_proto_rawDescData)
	})
	return file_sf_cosmos_transform_v1_transform_proto_rawDescData
}

var file_sf_cosmos_transform_v1_transform_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sf_cosmos_transform_v1_transform_proto_goTypes = []interface{}{
	(*CombinedFilter)(nil),    // 0: sf.cosmos.transform.v1.CombinedFilter
	(*EventTypeFilter)(nil),   // 1: sf.cosmos.transform.v1.EventTypeFilter
	(*EventOriginFilter)(nil), // 2: sf.cosmos.transform.v1.EventOriginFilter
	(*MessageTypeFilter)(nil), // 3: sf.cosmos.transform.v1.MessageTypeFilter
}
var file_sf_cosmos_transform_v1_transform_proto_depIdxs = []int32{
	1, // 0: sf.cosmos.transform.v1.CombinedFilter.event_type_filters:type_name -> sf.cosmos.transform.v1.EventTypeFilter
	2, // 1: sf.cosmos.transform.v1.CombinedFilter.event_origin_filters:type_name -> sf.cosmos.transform.v1.EventOriginFilter
	3, // 2: sf.cosmos.transform.v1.CombinedFilter.message_type_filters:type_name -> sf.cosmos.transform.v1.MessageTypeFilter
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sf_cosmos_transform_v1_transform_proto_init() }
func file_sf_cosmos_transform_v1_transform_proto_init() {
	if File_sf_cosmos_transform_v1_transform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_cosmos_transform_v1_transform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombinedFilter); i {
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
		file_sf_cosmos_transform_v1_transform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeFilter); i {
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
		file_sf_cosmos_transform_v1_transform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventOriginFilter); i {
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
		file_sf_cosmos_transform_v1_transform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageTypeFilter); i {
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
			RawDescriptor: file_sf_cosmos_transform_v1_transform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_cosmos_transform_v1_transform_proto_goTypes,
		DependencyIndexes: file_sf_cosmos_transform_v1_transform_proto_depIdxs,
		MessageInfos:      file_sf_cosmos_transform_v1_transform_proto_msgTypes,
	}.Build()
	File_sf_cosmos_transform_v1_transform_proto = out.File
	file_sf_cosmos_transform_v1_transform_proto_rawDesc = nil
	file_sf_cosmos_transform_v1_transform_proto_goTypes = nil
	file_sf_cosmos_transform_v1_transform_proto_depIdxs = nil
}
