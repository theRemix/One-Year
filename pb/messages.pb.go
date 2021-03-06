// protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

// protoc --go_out=$DST_DIR $SRC_DIR/addressbook.proto

// protoc --go_out=. --proto_path=. messages.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: pb/messages.proto

package pb

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

type Status_Code int32

const (
	Status_Success     Status_Code = 0
	Status_ClientError Status_Code = 1
	Status_ServerError Status_Code = 2
)

// Enum value maps for Status_Code.
var (
	Status_Code_name = map[int32]string{
		0: "Success",
		1: "ClientError",
		2: "ServerError",
	}
	Status_Code_value = map[string]int32{
		"Success":     0,
		"ClientError": 1,
		"ServerError": 2,
	}
)

func (x Status_Code) Enum() *Status_Code {
	p := new(Status_Code)
	*p = x
	return p
}

func (x Status_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_messages_proto_enumTypes[0].Descriptor()
}

func (Status_Code) Type() protoreflect.EnumType {
	return &file_pb_messages_proto_enumTypes[0]
}

func (x Status_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_Code.Descriptor instead.
func (Status_Code) EnumDescriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{1, 0}
}

type HostCreateStatus_Code int32

const (
	HostCreateStatus_Success     HostCreateStatus_Code = 0
	HostCreateStatus_ClientError HostCreateStatus_Code = 1 // none of these are really used
	HostCreateStatus_ServerError HostCreateStatus_Code = 2
)

// Enum value maps for HostCreateStatus_Code.
var (
	HostCreateStatus_Code_name = map[int32]string{
		0: "Success",
		1: "ClientError",
		2: "ServerError",
	}
	HostCreateStatus_Code_value = map[string]int32{
		"Success":     0,
		"ClientError": 1,
		"ServerError": 2,
	}
)

func (x HostCreateStatus_Code) Enum() *HostCreateStatus_Code {
	p := new(HostCreateStatus_Code)
	*p = x
	return p
}

func (x HostCreateStatus_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HostCreateStatus_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_messages_proto_enumTypes[1].Descriptor()
}

func (HostCreateStatus_Code) Type() protoreflect.EnumType {
	return &file_pb_messages_proto_enumTypes[1]
}

func (x HostCreateStatus_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HostCreateStatus_Code.Descriptor instead.
func (HostCreateStatus_Code) EnumDescriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{2, 0}
}

type Ready_Status int32

const (
	Ready_Success     Ready_Status = 0
	Ready_ClientError Ready_Status = 1 // none of these are really used
	Ready_ServerError Ready_Status = 2
)

// Enum value maps for Ready_Status.
var (
	Ready_Status_name = map[int32]string{
		0: "Success",
		1: "ClientError",
		2: "ServerError",
	}
	Ready_Status_value = map[string]int32{
		"Success":     0,
		"ClientError": 1,
		"ServerError": 2,
	}
)

func (x Ready_Status) Enum() *Ready_Status {
	p := new(Ready_Status)
	*p = x
	return p
}

func (x Ready_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ready_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_messages_proto_enumTypes[2].Descriptor()
}

func (Ready_Status) Type() protoreflect.EnumType {
	return &file_pb_messages_proto_enumTypes[2]
}

func (x Ready_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ready_Status.Descriptor instead.
func (Ready_Status) EnumDescriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{3, 0}
}

type PlayerEvent_Op int32

const (
	PlayerEvent_None     PlayerEvent_Op = 0 // default value gets omitted in json
	PlayerEvent_Ready    PlayerEvent_Op = 1
	PlayerEvent_Prompt   PlayerEvent_Op = 2
	PlayerEvent_Answered PlayerEvent_Op = 3
	PlayerEvent_Resolved PlayerEvent_Op = 4
)

// Enum value maps for PlayerEvent_Op.
var (
	PlayerEvent_Op_name = map[int32]string{
		0: "None",
		1: "Ready",
		2: "Prompt",
		3: "Answered",
		4: "Resolved",
	}
	PlayerEvent_Op_value = map[string]int32{
		"None":     0,
		"Ready":    1,
		"Prompt":   2,
		"Answered": 3,
		"Resolved": 4,
	}
)

func (x PlayerEvent_Op) Enum() *PlayerEvent_Op {
	p := new(PlayerEvent_Op)
	*p = x
	return p
}

func (x PlayerEvent_Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerEvent_Op) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_messages_proto_enumTypes[3].Descriptor()
}

func (PlayerEvent_Op) Type() protoreflect.EnumType {
	return &file_pb_messages_proto_enumTypes[3]
}

func (x PlayerEvent_Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerEvent_Op.Descriptor instead.
func (PlayerEvent_Op) EnumDescriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{5, 0}
}

type Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	PlayerName string `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
}

func (x *Join) Reset() {
	*x = Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join) ProtoMessage() {}

func (x *Join) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join.ProtoReflect.Descriptor instead.
func (*Join) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Join) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Join) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         Status_Code `protobuf:"varint,1,opt,name=code,proto3,enum=Status_Code" json:"code,omitempty"`
	ErrorMessage string      `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetCode() Status_Code {
	if x != nil {
		return x.Code
	}
	return Status_Success
}

func (x *Status) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type HostCreateStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         HostCreateStatus_Code `protobuf:"varint,1,opt,name=code,proto3,enum=HostCreateStatus_Code" json:"code,omitempty"`
	Key          string                `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	ErrorMessage string                `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *HostCreateStatus) Reset() {
	*x = HostCreateStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCreateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCreateStatus) ProtoMessage() {}

func (x *HostCreateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCreateStatus.ProtoReflect.Descriptor instead.
func (*HostCreateStatus) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{2}
}

func (x *HostCreateStatus) GetCode() HostCreateStatus_Code {
	if x != nil {
		return x.Code
	}
	return HostCreateStatus_Success
}

func (x *HostCreateStatus) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HostCreateStatus) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type Ready struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Ready_Status `protobuf:"varint,1,opt,name=status,proto3,enum=Ready_Status" json:"status,omitempty"`
}

func (x *Ready) Reset() {
	*x = Ready{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ready) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ready) ProtoMessage() {}

func (x *Ready) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ready.ProtoReflect.Descriptor instead.
func (*Ready) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{3}
}

func (x *Ready) GetStatus() Ready_Status {
	if x != nil {
		return x.Status
	}
	return Ready_Success
}

type PlayerScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ScoreChange int32  `protobuf:"varint,2,opt,name=scoreChange,proto3" json:"scoreChange,omitempty"`
	NewScore    int32  `protobuf:"varint,3,opt,name=newScore,proto3" json:"newScore,omitempty"`
	Comment     string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *PlayerScore) Reset() {
	*x = PlayerScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerScore) ProtoMessage() {}

func (x *PlayerScore) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerScore.ProtoReflect.Descriptor instead.
func (*PlayerScore) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerScore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerScore) GetScoreChange() int32 {
	if x != nil {
		return x.ScoreChange
	}
	return 0
}

func (x *PlayerScore) GetNewScore() int32 {
	if x != nil {
		return x.NewScore
	}
	return 0
}

func (x *PlayerScore) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type PlayerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op PlayerEvent_Op `protobuf:"varint,1,opt,name=op,proto3,enum=PlayerEvent_Op" json:"op,omitempty"`
	// Ready
	PlayerName string `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
	// Prompt
	Prompt *Prompt `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// Answered
	Answer *PromptAnswer `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
	// Resolved
	PlayerScores []*PlayerScore `protobuf:"bytes,5,rep,name=playerScores,proto3" json:"playerScores,omitempty"`
}

func (x *PlayerEvent) Reset() {
	*x = PlayerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerEvent) ProtoMessage() {}

func (x *PlayerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerEvent.ProtoReflect.Descriptor instead.
func (*PlayerEvent) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerEvent) GetOp() PlayerEvent_Op {
	if x != nil {
		return x.Op
	}
	return PlayerEvent_None
}

func (x *PlayerEvent) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerEvent) GetPrompt() *Prompt {
	if x != nil {
		return x.Prompt
	}
	return nil
}

func (x *PlayerEvent) GetAnswer() *PromptAnswer {
	if x != nil {
		return x.Answer
	}
	return nil
}

func (x *PlayerEvent) GetPlayerScores() []*PlayerScore {
	if x != nil {
		return x.PlayerScores
	}
	return nil
}

type Prompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code        string               `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Key         string               `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Contestants []*Prompt_Contestant `protobuf:"bytes,4,rep,name=contestants,proto3" json:"contestants,omitempty"`
	Options     []string             `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Prompt) Reset() {
	*x = Prompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt) ProtoMessage() {}

func (x *Prompt) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt.ProtoReflect.Descriptor instead.
func (*Prompt) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{6}
}

func (x *Prompt) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Prompt) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Prompt) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Prompt) GetContestants() []*Prompt_Contestant {
	if x != nil {
		return x.Contestants
	}
	return nil
}

func (x *Prompt) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

type PromptAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	PlayerName string                 `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
	Answers    []*PromptAnswer_Answer `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
	Comment    string                 `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *PromptAnswer) Reset() {
	*x = PromptAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptAnswer) ProtoMessage() {}

func (x *PromptAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptAnswer.ProtoReflect.Descriptor instead.
func (*PromptAnswer) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{7}
}

func (x *PromptAnswer) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PromptAnswer) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PromptAnswer) GetAnswers() []*PromptAnswer_Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

func (x *PromptAnswer) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type Resolve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         string         `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Key          string         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	PlayerScores []*PlayerScore `protobuf:"bytes,3,rep,name=playerScores,proto3" json:"playerScores,omitempty"`
}

func (x *Resolve) Reset() {
	*x = Resolve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resolve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resolve) ProtoMessage() {}

func (x *Resolve) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resolve.ProtoReflect.Descriptor instead.
func (*Resolve) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{8}
}

func (x *Resolve) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Resolve) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Resolve) GetPlayerScores() []*PlayerScore {
	if x != nil {
		return x.PlayerScores
	}
	return nil
}

type Prompt_Contestant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Thumbnail string `protobuf:"bytes,2,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *Prompt_Contestant) Reset() {
	*x = Prompt_Contestant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt_Contestant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt_Contestant) ProtoMessage() {}

func (x *Prompt_Contestant) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt_Contestant.ProtoReflect.Descriptor instead.
func (*Prompt_Contestant) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{6, 0}
}

func (x *Prompt_Contestant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Prompt_Contestant) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

type PromptAnswer_Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contestant string `protobuf:"bytes,1,opt,name=contestant,proto3" json:"contestant,omitempty"`
	Option     string `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *PromptAnswer_Answer) Reset() {
	*x = PromptAnswer_Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptAnswer_Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptAnswer_Answer) ProtoMessage() {}

func (x *PromptAnswer_Answer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptAnswer_Answer.ProtoReflect.Descriptor instead.
func (*PromptAnswer_Answer) Descriptor() ([]byte, []int) {
	return file_pb_messages_proto_rawDescGZIP(), []int{7, 0}
}

func (x *PromptAnswer_Answer) GetContestant() string {
	if x != nil {
		return x.Contestant
	}
	return ""
}

func (x *PromptAnswer_Answer) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

var File_pb_messages_proto protoreflect.FileDescriptor

var file_pb_messages_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x85, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x35, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x22, 0xab, 0x01, 0x0a, 0x10, 0x48, 0x6f, 0x73, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x35,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x02, 0x22, 0x67, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x25,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x37, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x22, 0x79,
	0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x8b, 0x02, 0x0a, 0x0b, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x10, 0x04, 0x22, 0xd2, 0x01, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x22, 0xce, 0x01, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x40, 0x0a, 0x06, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a,
	0x07, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30,
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_messages_proto_rawDescOnce sync.Once
	file_pb_messages_proto_rawDescData = file_pb_messages_proto_rawDesc
)

func file_pb_messages_proto_rawDescGZIP() []byte {
	file_pb_messages_proto_rawDescOnce.Do(func() {
		file_pb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_messages_proto_rawDescData)
	})
	return file_pb_messages_proto_rawDescData
}

var file_pb_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_pb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pb_messages_proto_goTypes = []interface{}{
	(Status_Code)(0),            // 0: Status.Code
	(HostCreateStatus_Code)(0),  // 1: HostCreateStatus.Code
	(Ready_Status)(0),           // 2: Ready.Status
	(PlayerEvent_Op)(0),         // 3: PlayerEvent.Op
	(*Join)(nil),                // 4: Join
	(*Status)(nil),              // 5: Status
	(*HostCreateStatus)(nil),    // 6: HostCreateStatus
	(*Ready)(nil),               // 7: Ready
	(*PlayerScore)(nil),         // 8: PlayerScore
	(*PlayerEvent)(nil),         // 9: PlayerEvent
	(*Prompt)(nil),              // 10: Prompt
	(*PromptAnswer)(nil),        // 11: PromptAnswer
	(*Resolve)(nil),             // 12: Resolve
	(*Prompt_Contestant)(nil),   // 13: Prompt.Contestant
	(*PromptAnswer_Answer)(nil), // 14: PromptAnswer.Answer
}
var file_pb_messages_proto_depIdxs = []int32{
	0,  // 0: Status.code:type_name -> Status.Code
	1,  // 1: HostCreateStatus.code:type_name -> HostCreateStatus.Code
	2,  // 2: Ready.status:type_name -> Ready.Status
	3,  // 3: PlayerEvent.op:type_name -> PlayerEvent.Op
	10, // 4: PlayerEvent.prompt:type_name -> Prompt
	11, // 5: PlayerEvent.answer:type_name -> PromptAnswer
	8,  // 6: PlayerEvent.playerScores:type_name -> PlayerScore
	13, // 7: Prompt.contestants:type_name -> Prompt.Contestant
	14, // 8: PromptAnswer.answers:type_name -> PromptAnswer.Answer
	8,  // 9: Resolve.playerScores:type_name -> PlayerScore
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pb_messages_proto_init() }
func file_pb_messages_proto_init() {
	if File_pb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join); i {
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
		file_pb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_pb_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCreateStatus); i {
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
		file_pb_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ready); i {
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
		file_pb_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerScore); i {
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
		file_pb_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerEvent); i {
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
		file_pb_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prompt); i {
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
		file_pb_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptAnswer); i {
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
		file_pb_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resolve); i {
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
		file_pb_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prompt_Contestant); i {
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
		file_pb_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptAnswer_Answer); i {
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
			RawDescriptor: file_pb_messages_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_messages_proto_goTypes,
		DependencyIndexes: file_pb_messages_proto_depIdxs,
		EnumInfos:         file_pb_messages_proto_enumTypes,
		MessageInfos:      file_pb_messages_proto_msgTypes,
	}.Build()
	File_pb_messages_proto = out.File
	file_pb_messages_proto_rawDesc = nil
	file_pb_messages_proto_goTypes = nil
	file_pb_messages_proto_depIdxs = nil
}
