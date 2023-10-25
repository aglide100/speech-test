// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: pb/svc/audio/audio.proto

package audio

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

type Audio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Audio) Reset() {
	*x = Audio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audio) ProtoMessage() {}

func (x *Audio) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audio.ProtoReflect.Descriptor instead.
func (*Audio) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{0}
}

func (x *Audio) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Who   string `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{2}
}

func (x *Auth) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Auth) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Speaker string `protobuf:"bytes,2,opt,name=speaker,proto3" json:"speaker,omitempty"`
	Id      string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	No      int32  `protobuf:"varint,4,opt,name=no,proto3" json:"no,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{3}
}

func (x *Job) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Job) GetSpeaker() string {
	if x != nil {
		return x.Speaker
	}
	return ""
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetNo() int32 {
	if x != nil {
		return x.No
	}
	return 0
}

type CheckingJobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *Auth `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *CheckingJobReq) Reset() {
	*x = CheckingJobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckingJobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckingJobReq) ProtoMessage() {}

func (x *CheckingJobReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckingJobReq.ProtoReflect.Descriptor instead.
func (*CheckingJobReq) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{4}
}

func (x *CheckingJobReq) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type CheckingJobRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job   *Job   `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CheckingJobRes) Reset() {
	*x = CheckingJobRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckingJobRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckingJobRes) ProtoMessage() {}

func (x *CheckingJobRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckingJobRes.ProtoReflect.Descriptor instead.
func (*CheckingJobRes) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{5}
}

func (x *CheckingJobRes) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *CheckingJobRes) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type SendingResultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audio *Audio `protobuf:"bytes,1,opt,name=audio,proto3" json:"audio,omitempty"`
	Auth  *Auth  `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
	Job   *Job   `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *SendingResultReq) Reset() {
	*x = SendingResultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendingResultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendingResultReq) ProtoMessage() {}

func (x *SendingResultReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendingResultReq.ProtoReflect.Descriptor instead.
func (*SendingResultReq) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{6}
}

func (x *SendingResultReq) GetAudio() *Audio {
	if x != nil {
		return x.Audio
	}
	return nil
}

func (x *SendingResultReq) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *SendingResultReq) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type MakingNewJobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth    *Auth  `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Speaker string `protobuf:"bytes,3,opt,name=speaker,proto3" json:"speaker,omitempty"`
	Path    string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *MakingNewJobReq) Reset() {
	*x = MakingNewJobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakingNewJobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakingNewJobReq) ProtoMessage() {}

func (x *MakingNewJobReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakingNewJobReq.ProtoReflect.Descriptor instead.
func (*MakingNewJobReq) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{7}
}

func (x *MakingNewJobReq) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *MakingNewJobReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MakingNewJobReq) GetSpeaker() string {
	if x != nil {
		return x.Speaker
	}
	return ""
}

func (x *MakingNewJobReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetAudioReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *Auth `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Job  *Job  `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *GetAudioReq) Reset() {
	*x = GetAudioReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_audio_audio_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAudioReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAudioReq) ProtoMessage() {}

func (x *GetAudioReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_audio_audio_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAudioReq.ProtoReflect.Descriptor instead.
func (*GetAudioReq) Descriptor() ([]byte, []int) {
	return file_pb_svc_audio_audio_proto_rawDescGZIP(), []int{8}
}

func (x *GetAudioReq) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *GetAudioReq) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

var File_pb_svc_audio_audio_proto protoreflect.FileDescriptor

var file_pb_svc_audio_audio_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x05, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x19, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x2e, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77,
	0x68, 0x6f, 0x22, 0x59, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e, 0x6f, 0x22, 0x2b, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x46, 0x0a, 0x0e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x03,
	0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x03, 0x6a, 0x6f, 0x62, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x63, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x12, 0x19, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4a,
	0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x74, 0x0a, 0x0f, 0x4d, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x40, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x32,
	0xb7, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x0c, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x12, 0x28, 0x0a, 0x0c, 0x4d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x77, 0x4a,
	0x6f, 0x62, 0x12, 0x10, 0x2e, 0x4d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x77, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x0b,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x0f, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x0d, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x11,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x6c, 0x69, 0x64, 0x65, 0x31, 0x30,
	0x30, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_svc_audio_audio_proto_rawDescOnce sync.Once
	file_pb_svc_audio_audio_proto_rawDescData = file_pb_svc_audio_audio_proto_rawDesc
)

func file_pb_svc_audio_audio_proto_rawDescGZIP() []byte {
	file_pb_svc_audio_audio_proto_rawDescOnce.Do(func() {
		file_pb_svc_audio_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_svc_audio_audio_proto_rawDescData)
	})
	return file_pb_svc_audio_audio_proto_rawDescData
}

var file_pb_svc_audio_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_svc_audio_audio_proto_goTypes = []interface{}{
	(*Audio)(nil),            // 0: Audio
	(*Error)(nil),            // 1: Error
	(*Auth)(nil),             // 2: Auth
	(*Job)(nil),              // 3: Job
	(*CheckingJobReq)(nil),   // 4: CheckingJobReq
	(*CheckingJobRes)(nil),   // 5: CheckingJobRes
	(*SendingResultReq)(nil), // 6: SendingResultReq
	(*MakingNewJobReq)(nil),  // 7: MakingNewJobReq
	(*GetAudioReq)(nil),      // 8: GetAudioReq
}
var file_pb_svc_audio_audio_proto_depIdxs = []int32{
	2,  // 0: CheckingJobReq.auth:type_name -> Auth
	3,  // 1: CheckingJobRes.job:type_name -> Job
	1,  // 2: CheckingJobRes.error:type_name -> Error
	0,  // 3: SendingResultReq.audio:type_name -> Audio
	2,  // 4: SendingResultReq.auth:type_name -> Auth
	3,  // 5: SendingResultReq.job:type_name -> Job
	2,  // 6: MakingNewJobReq.auth:type_name -> Auth
	2,  // 7: GetAudioReq.auth:type_name -> Auth
	3,  // 8: GetAudioReq.job:type_name -> Job
	8,  // 9: AudioService.GetAudio:input_type -> GetAudioReq
	7,  // 10: AudioService.MakingNewJob:input_type -> MakingNewJobReq
	4,  // 11: AudioService.CheckingJob:input_type -> CheckingJobReq
	6,  // 12: AudioService.SendingResult:input_type -> SendingResultReq
	0,  // 13: AudioService.GetAudio:output_type -> Audio
	1,  // 14: AudioService.MakingNewJob:output_type -> Error
	5,  // 15: AudioService.CheckingJob:output_type -> CheckingJobRes
	1,  // 16: AudioService.SendingResult:output_type -> Error
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pb_svc_audio_audio_proto_init() }
func file_pb_svc_audio_audio_proto_init() {
	if File_pb_svc_audio_audio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_svc_audio_audio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audio); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckingJobReq); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckingJobRes); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendingResultReq); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakingNewJobReq); i {
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
		file_pb_svc_audio_audio_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAudioReq); i {
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
			RawDescriptor: file_pb_svc_audio_audio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_svc_audio_audio_proto_goTypes,
		DependencyIndexes: file_pb_svc_audio_audio_proto_depIdxs,
		MessageInfos:      file_pb_svc_audio_audio_proto_msgTypes,
	}.Build()
	File_pb_svc_audio_audio_proto = out.File
	file_pb_svc_audio_audio_proto_rawDesc = nil
	file_pb_svc_audio_audio_proto_goTypes = nil
	file_pb_svc_audio_audio_proto_depIdxs = nil
}
