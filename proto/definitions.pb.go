// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: definitions.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Log_Level int32

const (
	Log_NONE    Log_Level = 0
	Log_INFO    Log_Level = 1
	Log_DEBUG   Log_Level = 2
	Log_WARNING Log_Level = 3
	Log_ERROR   Log_Level = 4
	Log_FATAL   Log_Level = 5
)

// Enum value maps for Log_Level.
var (
	Log_Level_name = map[int32]string{
		0: "NONE",
		1: "INFO",
		2: "DEBUG",
		3: "WARNING",
		4: "ERROR",
		5: "FATAL",
	}
	Log_Level_value = map[string]int32{
		"NONE":    0,
		"INFO":    1,
		"DEBUG":   2,
		"WARNING": 3,
		"ERROR":   4,
		"FATAL":   5,
	}
)

func (x Log_Level) Enum() *Log_Level {
	p := new(Log_Level)
	*p = x
	return p
}

func (x Log_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Log_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_definitions_proto_enumTypes[0].Descriptor()
}

func (Log_Level) Type() protoreflect.EnumType {
	return &file_definitions_proto_enumTypes[0]
}

func (x Log_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Log_Level.Descriptor instead.
func (Log_Level) EnumDescriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{0, 0}
}

type Claim_Type int32

const (
	Claim_STRING    Claim_Type = 0
	Claim_NUMBER    Claim_Type = 1
	Claim_BOOLEAN   Claim_Type = 2
	Claim_TIMESTAMP Claim_Type = 3
)

// Enum value maps for Claim_Type.
var (
	Claim_Type_name = map[int32]string{
		0: "STRING",
		1: "NUMBER",
		2: "BOOLEAN",
		3: "TIMESTAMP",
	}
	Claim_Type_value = map[string]int32{
		"STRING":    0,
		"NUMBER":    1,
		"BOOLEAN":   2,
		"TIMESTAMP": 3,
	}
)

func (x Claim_Type) Enum() *Claim_Type {
	p := new(Claim_Type)
	*p = x
	return p
}

func (x Claim_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Claim_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_definitions_proto_enumTypes[1].Descriptor()
}

func (Claim_Type) Type() protoreflect.EnumType {
	return &file_definitions_proto_enumTypes[1]
}

func (x Claim_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Claim_Type.Descriptor instead.
func (Claim_Type) EnumDescriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{1, 0}
}

type Result_Status int32

const (
	Result_ACKNOWLEDGED   Result_Status = 0
	Result_UNAUTHORIZED   Result_Status = 1
	Result_ACCESS_DENIED  Result_Status = 2
	Result_INTERNAL_ERROR Result_Status = 3
)

// Enum value maps for Result_Status.
var (
	Result_Status_name = map[int32]string{
		0: "ACKNOWLEDGED",
		1: "UNAUTHORIZED",
		2: "ACCESS_DENIED",
		3: "INTERNAL_ERROR",
	}
	Result_Status_value = map[string]int32{
		"ACKNOWLEDGED":   0,
		"UNAUTHORIZED":   1,
		"ACCESS_DENIED":  2,
		"INTERNAL_ERROR": 3,
	}
)

func (x Result_Status) Enum() *Result_Status {
	p := new(Result_Status)
	*p = x
	return p
}

func (x Result_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_definitions_proto_enumTypes[2].Descriptor()
}

func (Result_Status) Type() protoreflect.EnumType {
	return &file_definitions_proto_enumTypes[2]
}

func (x Result_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result_Status.Descriptor instead.
func (Result_Status) EnumDescriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{2, 0}
}

type Condition_Comparison int32

const (
	Condition_EQUAL            Condition_Comparison = 0
	Condition_NOT_EQUAL        Condition_Comparison = 1
	Condition_GREATER_THAN     Condition_Comparison = 2
	Condition_GREATER_OR_EQUAL Condition_Comparison = 3
	Condition_LESS_THAN        Condition_Comparison = 4
	Condition_LESS_OR_EQUAL    Condition_Comparison = 5
)

// Enum value maps for Condition_Comparison.
var (
	Condition_Comparison_name = map[int32]string{
		0: "EQUAL",
		1: "NOT_EQUAL",
		2: "GREATER_THAN",
		3: "GREATER_OR_EQUAL",
		4: "LESS_THAN",
		5: "LESS_OR_EQUAL",
	}
	Condition_Comparison_value = map[string]int32{
		"EQUAL":            0,
		"NOT_EQUAL":        1,
		"GREATER_THAN":     2,
		"GREATER_OR_EQUAL": 3,
		"LESS_THAN":        4,
		"LESS_OR_EQUAL":    5,
	}
)

func (x Condition_Comparison) Enum() *Condition_Comparison {
	p := new(Condition_Comparison)
	*p = x
	return p
}

func (x Condition_Comparison) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_Comparison) Descriptor() protoreflect.EnumDescriptor {
	return file_definitions_proto_enumTypes[3].Descriptor()
}

func (Condition_Comparison) Type() protoreflect.EnumType {
	return &file_definitions_proto_enumTypes[3]
}

func (x Condition_Comparison) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_Comparison.Descriptor instead.
func (Condition_Comparison) EnumDescriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{4, 0}
}

type QueryResult_QueryStatus int32

const (
	QueryResult_SUCCESS        QueryResult_QueryStatus = 0
	QueryResult_UNAUTHORIZED   QueryResult_QueryStatus = 1
	QueryResult_ACCESS_DENIED  QueryResult_QueryStatus = 3
	QueryResult_QUERY_ERROR    QueryResult_QueryStatus = 4
	QueryResult_INTERNAL_ERROR QueryResult_QueryStatus = 5
)

// Enum value maps for QueryResult_QueryStatus.
var (
	QueryResult_QueryStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "UNAUTHORIZED",
		3: "ACCESS_DENIED",
		4: "QUERY_ERROR",
		5: "INTERNAL_ERROR",
	}
	QueryResult_QueryStatus_value = map[string]int32{
		"SUCCESS":        0,
		"UNAUTHORIZED":   1,
		"ACCESS_DENIED":  3,
		"QUERY_ERROR":    4,
		"INTERNAL_ERROR": 5,
	}
)

func (x QueryResult_QueryStatus) Enum() *QueryResult_QueryStatus {
	p := new(QueryResult_QueryStatus)
	*p = x
	return p
}

func (x QueryResult_QueryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryResult_QueryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_definitions_proto_enumTypes[4].Descriptor()
}

func (QueryResult_QueryStatus) Type() protoreflect.EnumType {
	return &file_definitions_proto_enumTypes[4]
}

func (x QueryResult_QueryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryResult_QueryStatus.Descriptor instead.
func (QueryResult_QueryStatus) EnumDescriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{6, 0}
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source    string    `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Thread    string    `protobuf:"bytes,2,opt,name=thread,proto3" json:"thread,omitempty"`
	Message   string    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Claims    []*Claim  `protobuf:"bytes,4,rep,name=claims,proto3" json:"claims,omitempty"`
	File      string    `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty"`
	Line      uint32    `protobuf:"varint,6,opt,name=line,proto3" json:"line,omitempty"`
	Function  string    `protobuf:"bytes,7,opt,name=function,proto3" json:"function,omitempty"`
	Level     Log_Level `protobuf:"varint,8,opt,name=level,proto3,enum=proto.Log_Level" json:"level,omitempty"`
	Timestamp uint64    `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Group     bool      `protobuf:"varint,10,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definitions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_definitions_proto_msgTypes[0]
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
	return file_definitions_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Log) GetThread() string {
	if x != nil {
		return x.Thread
	}
	return ""
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Log) GetClaims() []*Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *Log) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Log) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Log) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *Log) GetLevel() Log_Level {
	if x != nil {
		return x.Level
	}
	return Log_NONE
}

func (x *Log) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Log) GetGroup() bool {
	if x != nil {
		return x.Group
	}
	return false
}

type Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  Claim_Type `protobuf:"varint,1,opt,name=type,proto3,enum=proto.Claim_Type" json:"type,omitempty"`
	Name  string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value string     `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Claim) Reset() {
	*x = Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definitions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claim) ProtoMessage() {}

func (x *Claim) ProtoReflect() protoreflect.Message {
	mi := &file_definitions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claim.ProtoReflect.Descriptor instead.
func (*Claim) Descriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{1}
}

func (x *Claim) GetType() Claim_Type {
	if x != nil {
		return x.Type
	}
	return Claim_STRING
}

func (x *Claim) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Claim) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Result_Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Result_Status" json:"status,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definitions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_definitions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetStatus() Result_Status {
	if x != nil {
		return x.Status
	}
	return Result_ACKNOWLEDGED
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AuditLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Claims    []*Claim `protobuf:"bytes,2,rep,name=claims,proto3" json:"claims,omitempty"`
	Topic     string   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Timestamp uint64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definitions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_definitions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{3}
}

func (x *AuditLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuditLog) GetClaims() []*Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *AuditLog) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *AuditLog) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field      string               `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Comparison Condition_Comparison `protobuf:"varint,2,opt,name=comparison,proto3,enum=proto.Condition_Comparison" json:"comparison,omitempty"`
	Value      string               `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definitions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_definitions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{4}
}

func (x *Condition) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Condition) GetComparison() Condition_Comparison {
	if x != nil {
		return x.Comparison
	}
	return Condition_EQUAL
}

func (x *Condition) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic      string       `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Conditions []*Condition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
	Cursor     string       `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definitions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_definitions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{5}
}

func (x *Query) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Query) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Query) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type QueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *QueryResult `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Logs    []*AuditLog  `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty"`
	Cursor  string       `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definitions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_definitions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_definitions_proto_rawDescGZIP(), []int{6}
}

func (x *QueryResult) GetStatus() *QueryResult {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *QueryResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryResult) GetLogs() []*AuditLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *QueryResult) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

var File_definitions_proto protoreflect.FileDescriptor

var file_definitions_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x06,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f,
	0x67, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x49, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41,
	0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x05, 0x22, 0x94, 0x01,
	0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c,
	0x45, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41,
	0x4d, 0x50, 0x10, 0x03, 0x22, 0xa5, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x53, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x4b, 0x4e, 0x4f, 0x57, 0x4c, 0x45, 0x44, 0x47, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49,
	0x5a, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x7e, 0x0a, 0x08,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xe6, 0x01, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x70, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f,
	0x6e, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x47,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41,
	0x4c, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e,
	0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51,
	0x55, 0x41, 0x4c, 0x10, 0x05, 0x22, 0x67, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xf6,
	0x01, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x22, 0x64, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x32, 0x97, 0x01, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x1a, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_definitions_proto_rawDescOnce sync.Once
	file_definitions_proto_rawDescData = file_definitions_proto_rawDesc
)

func file_definitions_proto_rawDescGZIP() []byte {
	file_definitions_proto_rawDescOnce.Do(func() {
		file_definitions_proto_rawDescData = protoimpl.X.CompressGZIP(file_definitions_proto_rawDescData)
	})
	return file_definitions_proto_rawDescData
}

var file_definitions_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_definitions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_definitions_proto_goTypes = []interface{}{
	(Log_Level)(0),               // 0: proto.Log.Level
	(Claim_Type)(0),              // 1: proto.Claim.Type
	(Result_Status)(0),           // 2: proto.Result.Status
	(Condition_Comparison)(0),    // 3: proto.Condition.Comparison
	(QueryResult_QueryStatus)(0), // 4: proto.QueryResult.QueryStatus
	(*Log)(nil),                  // 5: proto.Log
	(*Claim)(nil),                // 6: proto.Claim
	(*Result)(nil),               // 7: proto.Result
	(*AuditLog)(nil),             // 8: proto.AuditLog
	(*Condition)(nil),            // 9: proto.Condition
	(*Query)(nil),                // 10: proto.Query
	(*QueryResult)(nil),          // 11: proto.QueryResult
}
var file_definitions_proto_depIdxs = []int32{
	6,  // 0: proto.Log.claims:type_name -> proto.Claim
	0,  // 1: proto.Log.level:type_name -> proto.Log.Level
	1,  // 2: proto.Claim.type:type_name -> proto.Claim.Type
	2,  // 3: proto.Result.status:type_name -> proto.Result.Status
	6,  // 4: proto.AuditLog.claims:type_name -> proto.Claim
	3,  // 5: proto.Condition.comparison:type_name -> proto.Condition.Comparison
	9,  // 6: proto.Query.conditions:type_name -> proto.Condition
	11, // 7: proto.QueryResult.status:type_name -> proto.QueryResult
	8,  // 8: proto.QueryResult.logs:type_name -> proto.AuditLog
	5,  // 9: proto.Logging.WriteLog:input_type -> proto.Log
	8,  // 10: proto.Logging.WriteAuditLog:input_type -> proto.AuditLog
	10, // 11: proto.Logging.AuditQuery:input_type -> proto.Query
	7,  // 12: proto.Logging.WriteLog:output_type -> proto.Result
	7,  // 13: proto.Logging.WriteAuditLog:output_type -> proto.Result
	11, // 14: proto.Logging.AuditQuery:output_type -> proto.QueryResult
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_definitions_proto_init() }
func file_definitions_proto_init() {
	if File_definitions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_definitions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_definitions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claim); i {
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
		file_definitions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_definitions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLog); i {
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
		file_definitions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_definitions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_definitions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResult); i {
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
			RawDescriptor: file_definitions_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_definitions_proto_goTypes,
		DependencyIndexes: file_definitions_proto_depIdxs,
		EnumInfos:         file_definitions_proto_enumTypes,
		MessageInfos:      file_definitions_proto_msgTypes,
	}.Build()
	File_definitions_proto = out.File
	file_definitions_proto_rawDesc = nil
	file_definitions_proto_goTypes = nil
	file_definitions_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoggingClient is the client API for Logging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggingClient interface {
	// Creates a log
	WriteLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Result, error)
	// Creates an audit log
	WriteAuditLog(ctx context.Context, in *AuditLog, opts ...grpc.CallOption) (*Result, error)
	// Query audit logs for reporting or something
	AuditQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*QueryResult, error)
}

type loggingClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggingClient(cc grpc.ClientConnInterface) LoggingClient {
	return &loggingClient{cc}
}

func (c *loggingClient) WriteLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/proto.Logging/WriteLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingClient) WriteAuditLog(ctx context.Context, in *AuditLog, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/proto.Logging/WriteAuditLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingClient) AuditQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*QueryResult, error) {
	out := new(QueryResult)
	err := c.cc.Invoke(ctx, "/proto.Logging/AuditQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggingServer is the server API for Logging service.
type LoggingServer interface {
	// Creates a log
	WriteLog(context.Context, *Log) (*Result, error)
	// Creates an audit log
	WriteAuditLog(context.Context, *AuditLog) (*Result, error)
	// Query audit logs for reporting or something
	AuditQuery(context.Context, *Query) (*QueryResult, error)
}

// UnimplementedLoggingServer can be embedded to have forward compatible implementations.
type UnimplementedLoggingServer struct {
}

func (*UnimplementedLoggingServer) WriteLog(context.Context, *Log) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLog not implemented")
}
func (*UnimplementedLoggingServer) WriteAuditLog(context.Context, *AuditLog) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteAuditLog not implemented")
}
func (*UnimplementedLoggingServer) AuditQuery(context.Context, *Query) (*QueryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditQuery not implemented")
}

func RegisterLoggingServer(s *grpc.Server, srv LoggingServer) {
	s.RegisterService(&_Logging_serviceDesc, srv)
}

func _Logging_WriteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServer).WriteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Logging/WriteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServer).WriteLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logging_WriteAuditLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServer).WriteAuditLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Logging/WriteAuditLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServer).WriteAuditLog(ctx, req.(*AuditLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logging_AuditQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServer).AuditQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Logging/AuditQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServer).AuditQuery(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logging_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Logging",
	HandlerType: (*LoggingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteLog",
			Handler:    _Logging_WriteLog_Handler,
		},
		{
			MethodName: "WriteAuditLog",
			Handler:    _Logging_WriteAuditLog_Handler,
		},
		{
			MethodName: "AuditQuery",
			Handler:    _Logging_AuditQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "definitions.proto",
}
