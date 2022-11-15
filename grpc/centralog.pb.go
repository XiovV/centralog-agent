// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: grpc/centralog.proto

package centralog_agent

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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Container string `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *Log) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FollowLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers []string `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	ShowAll    bool     `protobuf:"varint,2,opt,name=showAll,proto3" json:"showAll,omitempty"`
}

func (x *FollowLogsRequest) Reset() {
	*x = FollowLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowLogsRequest) ProtoMessage() {}

func (x *FollowLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowLogsRequest.ProtoReflect.Descriptor instead.
func (*FollowLogsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{1}
}

func (x *FollowLogsRequest) GetContainers() []string {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *FollowLogsRequest) GetShowAll() bool {
	if x != nil {
		return x.ShowAll
	}
	return false
}

type ContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers []*Container `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
}

func (x *ContainerResponse) Reset() {
	*x = ContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerResponse) ProtoMessage() {}

func (x *ContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerResponse.ProtoReflect.Descriptor instead.
func (*ContainerResponse) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerResponse) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

type Containers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers []string `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
}

func (x *Containers) Reset() {
	*x = Containers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Containers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Containers) ProtoMessage() {}

func (x *Containers) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Containers.ProtoReflect.Descriptor instead.
func (*Containers) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{3}
}

func (x *Containers) GetContainers() []string {
	if x != nil {
		return x.Containers
	}
	return nil
}

type GetContainersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetContainersRequest) Reset() {
	*x = GetContainersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainersRequest) ProtoMessage() {}

func (x *GetContainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainersRequest.ProtoReflect.Descriptor instead.
func (*GetContainersRequest) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{4}
}

type CheckAPIKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CheckAPIKeyRequest) Reset() {
	*x = CheckAPIKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAPIKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAPIKeyRequest) ProtoMessage() {}

func (x *CheckAPIKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAPIKeyRequest.ProtoReflect.Descriptor instead.
func (*CheckAPIKeyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{5}
}

func (x *CheckAPIKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type CheckAPIKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *CheckAPIKeyResponse) Reset() {
	*x = CheckAPIKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAPIKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAPIKeyResponse) ProtoMessage() {}

func (x *CheckAPIKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAPIKeyResponse.ProtoReflect.Descriptor instead.
func (*CheckAPIKeyResponse) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{6}
}

func (x *CheckAPIKeyResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{7}
}

func (x *Container) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{8}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_centralog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_centralog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_grpc_centralog_proto_rawDescGZIP(), []int{9}
}

func (x *HealthCheckResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_grpc_centralog_proto protoreflect.FileDescriptor

var file_grpc_centralog_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x77,
	0x41, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x77, 0x41,
	0x6c, 0x6c, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x22, 0x2c, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x12, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x45,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd4, 0x02, 0x0a, 0x09, 0x43,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x35, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x13, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x13,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x50, 0x49, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0a, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x12, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x04, 0x2e,
	0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x1a, 0x0b, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x58, 0x69, 0x6f, 0x76, 0x56, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x2d,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_centralog_proto_rawDescOnce sync.Once
	file_grpc_centralog_proto_rawDescData = file_grpc_centralog_proto_rawDesc
)

func file_grpc_centralog_proto_rawDescGZIP() []byte {
	file_grpc_centralog_proto_rawDescOnce.Do(func() {
		file_grpc_centralog_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_centralog_proto_rawDescData)
	})
	return file_grpc_centralog_proto_rawDescData
}

var file_grpc_centralog_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_grpc_centralog_proto_goTypes = []interface{}{
	(*Log)(nil),                  // 0: Log
	(*FollowLogsRequest)(nil),    // 1: FollowLogsRequest
	(*ContainerResponse)(nil),    // 2: ContainerResponse
	(*Containers)(nil),           // 3: Containers
	(*GetContainersRequest)(nil), // 4: GetContainersRequest
	(*CheckAPIKeyRequest)(nil),   // 5: CheckAPIKeyRequest
	(*CheckAPIKeyResponse)(nil),  // 6: CheckAPIKeyResponse
	(*Container)(nil),            // 7: Container
	(*HealthCheckRequest)(nil),   // 8: HealthCheckRequest
	(*HealthCheckResponse)(nil),  // 9: HealthCheckResponse
}
var file_grpc_centralog_proto_depIdxs = []int32{
	7, // 0: ContainerResponse.containers:type_name -> Container
	8, // 1: Centralog.Health:input_type -> HealthCheckRequest
	5, // 2: Centralog.CheckAPIKey:input_type -> CheckAPIKeyRequest
	1, // 3: Centralog.FollowLogs:input_type -> FollowLogsRequest
	4, // 4: Centralog.GetContainers:input_type -> GetContainersRequest
	3, // 5: Centralog.GetRunningContainers:input_type -> Containers
	3, // 6: Centralog.GetContainersInfo:input_type -> Containers
	9, // 7: Centralog.Health:output_type -> HealthCheckResponse
	6, // 8: Centralog.CheckAPIKey:output_type -> CheckAPIKeyResponse
	0, // 9: Centralog.FollowLogs:output_type -> Log
	2, // 10: Centralog.GetContainers:output_type -> ContainerResponse
	3, // 11: Centralog.GetRunningContainers:output_type -> Containers
	2, // 12: Centralog.GetContainersInfo:output_type -> ContainerResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_centralog_proto_init() }
func file_grpc_centralog_proto_init() {
	if File_grpc_centralog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_centralog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_grpc_centralog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowLogsRequest); i {
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
		file_grpc_centralog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerResponse); i {
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
		file_grpc_centralog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Containers); i {
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
		file_grpc_centralog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContainersRequest); i {
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
		file_grpc_centralog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAPIKeyRequest); i {
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
		file_grpc_centralog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAPIKeyResponse); i {
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
		file_grpc_centralog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_grpc_centralog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_grpc_centralog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_grpc_centralog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_centralog_proto_goTypes,
		DependencyIndexes: file_grpc_centralog_proto_depIdxs,
		MessageInfos:      file_grpc_centralog_proto_msgTypes,
	}.Build()
	File_grpc_centralog_proto = out.File
	file_grpc_centralog_proto_rawDesc = nil
	file_grpc_centralog_proto_goTypes = nil
	file_grpc_centralog_proto_depIdxs = nil
}
