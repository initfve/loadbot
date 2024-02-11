// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: lbot/proto/progress.proto

package proto

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

type ProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshInterval string `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
}

func (x *ProgressRequest) Reset() {
	*x = ProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lbot_proto_progress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressRequest) ProtoMessage() {}

func (x *ProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lbot_proto_progress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressRequest.ProtoReflect.Descriptor instead.
func (*ProgressRequest) Descriptor() ([]byte, []int) {
	return file_lbot_proto_progress_proto_rawDescGZIP(), []int{0}
}

func (x *ProgressRequest) GetRefreshInterval() string {
	if x != nil {
		return x.RefreshInterval
	}
	return ""
}

type ProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// change reqeust total
	Requests  uint64  `protobuf:"varint,1,opt,name=requests,proto3" json:"requests,omitempty"`
	Duration  uint64  `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Rps       uint64  `protobuf:"varint,3,opt,name=rps,proto3" json:"rps,omitempty"`
	ErrorRate float32 `protobuf:"fixed32,4,opt,name=error_rate,json=errorRate,proto3" json:"error_rate,omitempty"`
	// job type -
	// duration, operation, nolimit
	JobName           string `protobuf:"bytes,5,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	RequestDuration   uint64 `protobuf:"varint,6,opt,name=request_duration,json=requestDuration,proto3" json:"request_duration,omitempty"`
	RequestOperations uint64 `protobuf:"varint,7,opt,name=request_operations,json=requestOperations,proto3" json:"request_operations,omitempty"`
}

func (x *ProgressResponse) Reset() {
	*x = ProgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lbot_proto_progress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressResponse) ProtoMessage() {}

func (x *ProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lbot_proto_progress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressResponse.ProtoReflect.Descriptor instead.
func (*ProgressResponse) Descriptor() ([]byte, []int) {
	return file_lbot_proto_progress_proto_rawDescGZIP(), []int{1}
}

func (x *ProgressResponse) GetRequests() uint64 {
	if x != nil {
		return x.Requests
	}
	return 0
}

func (x *ProgressResponse) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *ProgressResponse) GetRps() uint64 {
	if x != nil {
		return x.Rps
	}
	return 0
}

func (x *ProgressResponse) GetErrorRate() float32 {
	if x != nil {
		return x.ErrorRate
	}
	return 0
}

func (x *ProgressResponse) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *ProgressResponse) GetRequestDuration() uint64 {
	if x != nil {
		return x.RequestDuration
	}
	return 0
}

func (x *ProgressResponse) GetRequestOperations() uint64 {
	if x != nil {
		return x.RequestOperations
	}
	return 0
}

var File_lbot_proto_progress_proto protoreflect.FileDescriptor

var file_lbot_proto_progress_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x62, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x22, 0xf0, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x70, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x4d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lbot_proto_progress_proto_rawDescOnce sync.Once
	file_lbot_proto_progress_proto_rawDescData = file_lbot_proto_progress_proto_rawDesc
)

func file_lbot_proto_progress_proto_rawDescGZIP() []byte {
	file_lbot_proto_progress_proto_rawDescOnce.Do(func() {
		file_lbot_proto_progress_proto_rawDescData = protoimpl.X.CompressGZIP(file_lbot_proto_progress_proto_rawDescData)
	})
	return file_lbot_proto_progress_proto_rawDescData
}

var file_lbot_proto_progress_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lbot_proto_progress_proto_goTypes = []interface{}{
	(*ProgressRequest)(nil),  // 0: proto.ProgressRequest
	(*ProgressResponse)(nil), // 1: proto.ProgressResponse
}
var file_lbot_proto_progress_proto_depIdxs = []int32{
	0, // 0: proto.ProgressProcess.Run:input_type -> proto.ProgressRequest
	1, // 1: proto.ProgressProcess.Run:output_type -> proto.ProgressResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lbot_proto_progress_proto_init() }
func file_lbot_proto_progress_proto_init() {
	if File_lbot_proto_progress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lbot_proto_progress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressRequest); i {
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
		file_lbot_proto_progress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressResponse); i {
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
			RawDescriptor: file_lbot_proto_progress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lbot_proto_progress_proto_goTypes,
		DependencyIndexes: file_lbot_proto_progress_proto_depIdxs,
		MessageInfos:      file_lbot_proto_progress_proto_msgTypes,
	}.Build()
	File_lbot_proto_progress_proto = out.File
	file_lbot_proto_progress_proto_rawDesc = nil
	file_lbot_proto_progress_proto_goTypes = nil
	file_lbot_proto_progress_proto_depIdxs = nil
}
