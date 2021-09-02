// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: livekit_recording.proto

package livekit

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

type RecordingPreset int32

const (
	RecordingPreset_NONE       RecordingPreset = 0
	RecordingPreset_HD_30      RecordingPreset = 1 // 720p, 30fps
	RecordingPreset_HD_60      RecordingPreset = 2 // 720p, 60fps
	RecordingPreset_FULL_HD_30 RecordingPreset = 3 // 1080p, 30fps
	RecordingPreset_FULL_HD_60 RecordingPreset = 4 // 1080p, 60fps
)

// Enum value maps for RecordingPreset.
var (
	RecordingPreset_name = map[int32]string{
		0: "NONE",
		1: "HD_30",
		2: "HD_60",
		3: "FULL_HD_30",
		4: "FULL_HD_60",
	}
	RecordingPreset_value = map[string]int32{
		"NONE":       0,
		"HD_30":      1,
		"HD_60":      2,
		"FULL_HD_30": 3,
		"FULL_HD_60": 4,
	}
)

func (x RecordingPreset) Enum() *RecordingPreset {
	p := new(RecordingPreset)
	*p = x
	return p
}

func (x RecordingPreset) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordingPreset) Descriptor() protoreflect.EnumDescriptor {
	return file_livekit_recording_proto_enumTypes[0].Descriptor()
}

func (RecordingPreset) Type() protoreflect.EnumType {
	return &file_livekit_recording_proto_enumTypes[0]
}

func (x RecordingPreset) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordingPreset.Descriptor instead.
func (RecordingPreset) EnumDescriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{0}
}

type StartRecordingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input   *RecordingInput   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output  *RecordingOutput  `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Options *RecordingOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *StartRecordingRequest) Reset() {
	*x = StartRecordingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRecordingRequest) ProtoMessage() {}

func (x *StartRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRecordingRequest.ProtoReflect.Descriptor instead.
func (*StartRecordingRequest) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{0}
}

func (x *StartRecordingRequest) GetInput() *RecordingInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *StartRecordingRequest) GetOutput() *RecordingOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *StartRecordingRequest) GetOptions() *RecordingOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type EndRecordingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordingId string `protobuf:"bytes,1,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
}

func (x *EndRecordingRequest) Reset() {
	*x = EndRecordingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndRecordingRequest) ProtoMessage() {}

func (x *EndRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndRecordingRequest.ProtoReflect.Descriptor instead.
func (*EndRecordingRequest) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{1}
}

func (x *EndRecordingRequest) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

type RecordingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordingId string `protobuf:"bytes,1,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
}

func (x *RecordingResponse) Reset() {
	*x = RecordingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingResponse) ProtoMessage() {}

func (x *RecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingResponse.ProtoReflect.Descriptor instead.
func (*RecordingResponse) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{2}
}

func (x *RecordingResponse) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

type RecordingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string             `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Template *RecordingTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *RecordingInput) Reset() {
	*x = RecordingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingInput) ProtoMessage() {}

func (x *RecordingInput) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingInput.ProtoReflect.Descriptor instead.
func (*RecordingInput) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{3}
}

func (x *RecordingInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RecordingInput) GetTemplate() *RecordingTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

type RecordingTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layout string `protobuf:"bytes,1,opt,name=layout,proto3" json:"layout,omitempty"`
	WsUrl  string `protobuf:"bytes,2,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
	// either token or room name required
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	RoomName string `protobuf:"bytes,4,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
}

func (x *RecordingTemplate) Reset() {
	*x = RecordingTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingTemplate) ProtoMessage() {}

func (x *RecordingTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingTemplate.ProtoReflect.Descriptor instead.
func (*RecordingTemplate) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{4}
}

func (x *RecordingTemplate) GetLayout() string {
	if x != nil {
		return x.Layout
	}
	return ""
}

func (x *RecordingTemplate) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

func (x *RecordingTemplate) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RecordingTemplate) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

type RecordingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File string             `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	S3   *RecordingS3Output `protobuf:"bytes,2,opt,name=s3,proto3" json:"s3,omitempty"`
	Rtmp string             `protobuf:"bytes,3,opt,name=rtmp,proto3" json:"rtmp,omitempty"`
}

func (x *RecordingOutput) Reset() {
	*x = RecordingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingOutput) ProtoMessage() {}

func (x *RecordingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingOutput.ProtoReflect.Descriptor instead.
func (*RecordingOutput) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{5}
}

func (x *RecordingOutput) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *RecordingOutput) GetS3() *RecordingS3Output {
	if x != nil {
		return x.S3
	}
	return nil
}

func (x *RecordingOutput) GetRtmp() string {
	if x != nil {
		return x.Rtmp
	}
	return ""
}

type RecordingS3Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// optional
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	Secret    string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *RecordingS3Output) Reset() {
	*x = RecordingS3Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingS3Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingS3Output) ProtoMessage() {}

func (x *RecordingS3Output) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingS3Output.ProtoReflect.Descriptor instead.
func (*RecordingS3Output) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{6}
}

func (x *RecordingS3Output) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *RecordingS3Output) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RecordingS3Output) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *RecordingS3Output) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type RecordingOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Preset         RecordingPreset `protobuf:"varint,1,opt,name=preset,proto3,enum=livekit.RecordingPreset" json:"preset,omitempty"`
	InputWidth     int32           `protobuf:"varint,2,opt,name=input_width,json=inputWidth,proto3" json:"input_width,omitempty"`
	InputHeight    int32           `protobuf:"varint,3,opt,name=input_height,json=inputHeight,proto3" json:"input_height,omitempty"`
	OutputWidth    int32           `protobuf:"varint,4,opt,name=output_width,json=outputWidth,proto3" json:"output_width,omitempty"`
	OutputHeight   int32           `protobuf:"varint,5,opt,name=output_height,json=outputHeight,proto3" json:"output_height,omitempty"`
	Depth          int32           `protobuf:"varint,6,opt,name=depth,proto3" json:"depth,omitempty"`
	Framerate      int32           `protobuf:"varint,7,opt,name=framerate,proto3" json:"framerate,omitempty"`
	AudioBitrate   int32           `protobuf:"varint,8,opt,name=audio_bitrate,json=audioBitrate,proto3" json:"audio_bitrate,omitempty"`
	AudioFrequency int32           `protobuf:"varint,9,opt,name=audio_frequency,json=audioFrequency,proto3" json:"audio_frequency,omitempty"`
	VideoBitrate   int32           `protobuf:"varint,10,opt,name=video_bitrate,json=videoBitrate,proto3" json:"video_bitrate,omitempty"`
}

func (x *RecordingOptions) Reset() {
	*x = RecordingOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_recording_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingOptions) ProtoMessage() {}

func (x *RecordingOptions) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_recording_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingOptions.ProtoReflect.Descriptor instead.
func (*RecordingOptions) Descriptor() ([]byte, []int) {
	return file_livekit_recording_proto_rawDescGZIP(), []int{7}
}

func (x *RecordingOptions) GetPreset() RecordingPreset {
	if x != nil {
		return x.Preset
	}
	return RecordingPreset_NONE
}

func (x *RecordingOptions) GetInputWidth() int32 {
	if x != nil {
		return x.InputWidth
	}
	return 0
}

func (x *RecordingOptions) GetInputHeight() int32 {
	if x != nil {
		return x.InputHeight
	}
	return 0
}

func (x *RecordingOptions) GetOutputWidth() int32 {
	if x != nil {
		return x.OutputWidth
	}
	return 0
}

func (x *RecordingOptions) GetOutputHeight() int32 {
	if x != nil {
		return x.OutputHeight
	}
	return 0
}

func (x *RecordingOptions) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *RecordingOptions) GetFramerate() int32 {
	if x != nil {
		return x.Framerate
	}
	return 0
}

func (x *RecordingOptions) GetAudioBitrate() int32 {
	if x != nil {
		return x.AudioBitrate
	}
	return 0
}

func (x *RecordingOptions) GetAudioFrequency() int32 {
	if x != nil {
		return x.AudioFrequency
	}
	return 0
}

func (x *RecordingOptions) GetVideoBitrate() int32 {
	if x != nil {
		return x.VideoBitrate
	}
	return 0
}

var File_livekit_recording_proto protoreflect.FileDescriptor

var file_livekit_recording_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x33, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x38, 0x0a, 0x13, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x11,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x22, 0x75, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77,
	0x73, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2a,
	0x0a, 0x02, 0x73, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x33,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x02, 0x73, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x74,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x74, 0x6d, 0x70, 0x22, 0x74,
	0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x33, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0xf7, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x5f, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x2a, 0x51,
	0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x48,
	0x44, 0x5f, 0x33, 0x30, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x44, 0x5f, 0x36, 0x30, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x48, 0x44, 0x5f, 0x33, 0x30, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x48, 0x44, 0x5f, 0x36, 0x30, 0x10,
	0x04, 0x32, 0xaa, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45,
	0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_livekit_recording_proto_rawDescOnce sync.Once
	file_livekit_recording_proto_rawDescData = file_livekit_recording_proto_rawDesc
)

func file_livekit_recording_proto_rawDescGZIP() []byte {
	file_livekit_recording_proto_rawDescOnce.Do(func() {
		file_livekit_recording_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_recording_proto_rawDescData)
	})
	return file_livekit_recording_proto_rawDescData
}

var file_livekit_recording_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_livekit_recording_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_livekit_recording_proto_goTypes = []interface{}{
	(RecordingPreset)(0),          // 0: livekit.RecordingPreset
	(*StartRecordingRequest)(nil), // 1: livekit.StartRecordingRequest
	(*EndRecordingRequest)(nil),   // 2: livekit.EndRecordingRequest
	(*RecordingResponse)(nil),     // 3: livekit.RecordingResponse
	(*RecordingInput)(nil),        // 4: livekit.RecordingInput
	(*RecordingTemplate)(nil),     // 5: livekit.RecordingTemplate
	(*RecordingOutput)(nil),       // 6: livekit.RecordingOutput
	(*RecordingS3Output)(nil),     // 7: livekit.RecordingS3Output
	(*RecordingOptions)(nil),      // 8: livekit.RecordingOptions
}
var file_livekit_recording_proto_depIdxs = []int32{
	4, // 0: livekit.StartRecordingRequest.input:type_name -> livekit.RecordingInput
	6, // 1: livekit.StartRecordingRequest.output:type_name -> livekit.RecordingOutput
	8, // 2: livekit.StartRecordingRequest.options:type_name -> livekit.RecordingOptions
	5, // 3: livekit.RecordingInput.template:type_name -> livekit.RecordingTemplate
	7, // 4: livekit.RecordingOutput.s3:type_name -> livekit.RecordingS3Output
	0, // 5: livekit.RecordingOptions.preset:type_name -> livekit.RecordingPreset
	1, // 6: livekit.RecordingService.StartRecording:input_type -> livekit.StartRecordingRequest
	2, // 7: livekit.RecordingService.EndRecording:input_type -> livekit.EndRecordingRequest
	3, // 8: livekit.RecordingService.StartRecording:output_type -> livekit.RecordingResponse
	3, // 9: livekit.RecordingService.EndRecording:output_type -> livekit.RecordingResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_livekit_recording_proto_init() }
func file_livekit_recording_proto_init() {
	if File_livekit_recording_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_livekit_recording_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRecordingRequest); i {
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
		file_livekit_recording_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndRecordingRequest); i {
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
		file_livekit_recording_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingResponse); i {
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
		file_livekit_recording_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingInput); i {
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
		file_livekit_recording_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingTemplate); i {
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
		file_livekit_recording_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingOutput); i {
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
		file_livekit_recording_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingS3Output); i {
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
		file_livekit_recording_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingOptions); i {
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
			RawDescriptor: file_livekit_recording_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_livekit_recording_proto_goTypes,
		DependencyIndexes: file_livekit_recording_proto_depIdxs,
		EnumInfos:         file_livekit_recording_proto_enumTypes,
		MessageInfos:      file_livekit_recording_proto_msgTypes,
	}.Build()
	File_livekit_recording_proto = out.File
	file_livekit_recording_proto_rawDesc = nil
	file_livekit_recording_proto_goTypes = nil
	file_livekit_recording_proto_depIdxs = nil
}
