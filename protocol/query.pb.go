// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: query.proto

package protocol

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

type AlertType int32

const (
	AlertType_UNKNOWN_ALERT_TYPE AlertType = 0
	AlertType_TRANSACTION        AlertType = 1
	AlertType_BLOCK              AlertType = 2
)

// Enum value maps for AlertType.
var (
	AlertType_name = map[int32]string{
		0: "UNKNOWN_ALERT_TYPE",
		1: "TRANSACTION",
		2: "BLOCK",
	}
	AlertType_value = map[string]int32{
		"UNKNOWN_ALERT_TYPE": 0,
		"TRANSACTION":        1,
		"BLOCK":              2,
	}
)

func (x AlertType) Enum() *AlertType {
	p := new(AlertType)
	*p = x
	return p
}

func (x AlertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertType) Descriptor() protoreflect.EnumDescriptor {
	return file_query_proto_enumTypes[0].Descriptor()
}

func (AlertType) Type() protoreflect.EnumType {
	return &file_query_proto_enumTypes[0]
}

func (x AlertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertType.Descriptor instead.
func (AlertType) EnumDescriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

type AlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts        []*SignedAlert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	NextPageToken string         `protobuf:"bytes,2,opt,name=nextPageToken,proto3" json:"nextPageToken,omitempty"`
}

func (x *AlertResponse) Reset() {
	*x = AlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertResponse) ProtoMessage() {}

func (x *AlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertResponse.ProtoReflect.Descriptor instead.
func (*AlertResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

func (x *AlertResponse) GetAlerts() []*SignedAlert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *AlertResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ScannerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ScannerInfo) Reset() {
	*x = ScannerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScannerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScannerInfo) ProtoMessage() {}

func (x *ScannerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScannerInfo.ProtoReflect.Descriptor instead.
func (*ScannerInfo) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{1}
}

func (x *ScannerInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image     string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	ImageHash string `protobuf:"bytes,2,opt,name=imageHash,proto3" json:"imageHash,omitempty"`
	Id        string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	IsTest    bool   `protobuf:"varint,4,opt,name=isTest,proto3" json:"isTest,omitempty"`
	Manifest  string `protobuf:"bytes,5,opt,name=manifest,proto3" json:"manifest,omitempty"`
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{2}
}

func (x *AgentInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *AgentInfo) GetImageHash() string {
	if x != nil {
		return x.ImageHash
	}
	return ""
}

func (x *AgentInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgentInfo) GetIsTest() bool {
	if x != nil {
		return x.IsTest
	}
	return false
}

func (x *AgentInfo) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Algorithm string `protobuf:"bytes,2,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Signer    string `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{3}
}

func (x *Signature) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Signature) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *Signature) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      AlertType         `protobuf:"varint,2,opt,name=type,proto3,enum=network.forta.AlertType" json:"type,omitempty"`
	Finding   *Finding          `protobuf:"bytes,3,opt,name=finding,proto3" json:"finding,omitempty"`
	Timestamp string            `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Agent     *AgentInfo        `protobuf:"bytes,6,opt,name=agent,proto3" json:"agent,omitempty"`
	Tags      map[string]string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Scanner   *ScannerInfo      `protobuf:"bytes,8,opt,name=scanner,proto3" json:"scanner,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{4}
}

func (x *Alert) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Alert) GetType() AlertType {
	if x != nil {
		return x.Type
	}
	return AlertType_UNKNOWN_ALERT_TYPE
}

func (x *Alert) GetFinding() *Finding {
	if x != nil {
		return x.Finding
	}
	return nil
}

func (x *Alert) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Alert) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Alert) GetAgent() *AgentInfo {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *Alert) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Alert) GetScanner() *ScannerInfo {
	if x != nil {
		return x.Scanner
	}
	return nil
}

type SignedAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert           *Alert     `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
	Signature       *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	ChainId         string     `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	BlockNumber     string     `protobuf:"bytes,4,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	PublishedWithTx string     `protobuf:"bytes,5,opt,name=publishedWithTx,proto3" json:"publishedWithTx,omitempty"`
}

func (x *SignedAlert) Reset() {
	*x = SignedAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedAlert) ProtoMessage() {}

func (x *SignedAlert) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedAlert.ProtoReflect.Descriptor instead.
func (*SignedAlert) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{5}
}

func (x *SignedAlert) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

func (x *SignedAlert) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedAlert) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *SignedAlert) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *SignedAlert) GetPublishedWithTx() string {
	if x != nil {
		return x.PublishedWithTx
	}
	return ""
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedAlert       *SignedAlert           `protobuf:"bytes,1,opt,name=signedAlert,proto3" json:"signedAlert,omitempty"`
	EvalTxRequest     *EvaluateTxRequest     `protobuf:"bytes,2,opt,name=evalTxRequest,proto3" json:"evalTxRequest,omitempty"`
	EvalTxResponse    *EvaluateTxResponse    `protobuf:"bytes,3,opt,name=evalTxResponse,proto3" json:"evalTxResponse,omitempty"`
	EvalBlockRequest  *EvaluateBlockRequest  `protobuf:"bytes,4,opt,name=evalBlockRequest,proto3" json:"evalBlockRequest,omitempty"`
	EvalBlockResponse *EvaluateBlockResponse `protobuf:"bytes,5,opt,name=evalBlockResponse,proto3" json:"evalBlockResponse,omitempty"`
	AgentInfo         *AgentInfo             `protobuf:"bytes,6,opt,name=agentInfo,proto3" json:"agentInfo,omitempty"`
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{6}
}

func (x *NotifyRequest) GetSignedAlert() *SignedAlert {
	if x != nil {
		return x.SignedAlert
	}
	return nil
}

func (x *NotifyRequest) GetEvalTxRequest() *EvaluateTxRequest {
	if x != nil {
		return x.EvalTxRequest
	}
	return nil
}

func (x *NotifyRequest) GetEvalTxResponse() *EvaluateTxResponse {
	if x != nil {
		return x.EvalTxResponse
	}
	return nil
}

func (x *NotifyRequest) GetEvalBlockRequest() *EvaluateBlockRequest {
	if x != nil {
		return x.EvalBlockRequest
	}
	return nil
}

func (x *NotifyRequest) GetEvalBlockResponse() *EvaluateBlockResponse {
	if x != nil {
		return x.EvalBlockResponse
	}
	return nil
}

func (x *NotifyRequest) GetAgentInfo() *AgentInfo {
	if x != nil {
		return x.AgentInfo
	}
	return nil
}

type NotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyResponse) Reset() {
	*x = NotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyResponse) ProtoMessage() {}

func (x *NotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyResponse.ProtoReflect.Descriptor instead.
func (*NotifyResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{7}
}

var File_query_proto protoreflect.FileDescriptor

var file_query_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x1a, 0x0b, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0d, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x83, 0x01,
	0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x22, 0xe5, 0x03, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x07,
	0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3e, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x34, 0x0a, 0x07, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74,
	0x61, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd7, 0x01, 0x0a,
	0x0b, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x05,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x54, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x54, 0x78, 0x22, 0xbd, 0x03, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x46, 0x0a, 0x0d, 0x65, 0x76, 0x61, 0x6c, 0x54, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0d, 0x65, 0x76, 0x61, 0x6c, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49,
	0x0a, 0x0e, 0x65, 0x76, 0x61, 0x6c, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x65, 0x76, 0x61, 0x6c, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x65, 0x76, 0x61,
	0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f,
	0x72, 0x74, 0x61, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x10, 0x65, 0x76, 0x61, 0x6c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x11, 0x65, 0x76,
	0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x11, 0x65, 0x76, 0x61,
	0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74,
	0x61, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x3f, 0x0a, 0x09, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x02, 0x32, 0x54, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_proto_rawDescOnce sync.Once
	file_query_proto_rawDescData = file_query_proto_rawDesc
)

func file_query_proto_rawDescGZIP() []byte {
	file_query_proto_rawDescOnce.Do(func() {
		file_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_proto_rawDescData)
	})
	return file_query_proto_rawDescData
}

var file_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_query_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_query_proto_goTypes = []interface{}{
	(AlertType)(0),                // 0: network.forta.AlertType
	(*AlertResponse)(nil),         // 1: network.forta.AlertResponse
	(*ScannerInfo)(nil),           // 2: network.forta.ScannerInfo
	(*AgentInfo)(nil),             // 3: network.forta.AgentInfo
	(*Signature)(nil),             // 4: network.forta.Signature
	(*Alert)(nil),                 // 5: network.forta.Alert
	(*SignedAlert)(nil),           // 6: network.forta.SignedAlert
	(*NotifyRequest)(nil),         // 7: network.forta.NotifyRequest
	(*NotifyResponse)(nil),        // 8: network.forta.NotifyResponse
	nil,                           // 9: network.forta.Alert.MetadataEntry
	nil,                           // 10: network.forta.Alert.TagsEntry
	(*Finding)(nil),               // 11: network.forta.Finding
	(*EvaluateTxRequest)(nil),     // 12: network.forta.EvaluateTxRequest
	(*EvaluateTxResponse)(nil),    // 13: network.forta.EvaluateTxResponse
	(*EvaluateBlockRequest)(nil),  // 14: network.forta.EvaluateBlockRequest
	(*EvaluateBlockResponse)(nil), // 15: network.forta.EvaluateBlockResponse
}
var file_query_proto_depIdxs = []int32{
	6,  // 0: network.forta.AlertResponse.alerts:type_name -> network.forta.SignedAlert
	0,  // 1: network.forta.Alert.type:type_name -> network.forta.AlertType
	11, // 2: network.forta.Alert.finding:type_name -> network.forta.Finding
	9,  // 3: network.forta.Alert.metadata:type_name -> network.forta.Alert.MetadataEntry
	3,  // 4: network.forta.Alert.agent:type_name -> network.forta.AgentInfo
	10, // 5: network.forta.Alert.tags:type_name -> network.forta.Alert.TagsEntry
	2,  // 6: network.forta.Alert.scanner:type_name -> network.forta.ScannerInfo
	5,  // 7: network.forta.SignedAlert.alert:type_name -> network.forta.Alert
	4,  // 8: network.forta.SignedAlert.signature:type_name -> network.forta.Signature
	6,  // 9: network.forta.NotifyRequest.signedAlert:type_name -> network.forta.SignedAlert
	12, // 10: network.forta.NotifyRequest.evalTxRequest:type_name -> network.forta.EvaluateTxRequest
	13, // 11: network.forta.NotifyRequest.evalTxResponse:type_name -> network.forta.EvaluateTxResponse
	14, // 12: network.forta.NotifyRequest.evalBlockRequest:type_name -> network.forta.EvaluateBlockRequest
	15, // 13: network.forta.NotifyRequest.evalBlockResponse:type_name -> network.forta.EvaluateBlockResponse
	3,  // 14: network.forta.NotifyRequest.agentInfo:type_name -> network.forta.AgentInfo
	7,  // 15: network.forta.QueryNode.Notify:input_type -> network.forta.NotifyRequest
	8,  // 16: network.forta.QueryNode.Notify:output_type -> network.forta.NotifyResponse
	16, // [16:17] is the sub-list for method output_type
	15, // [15:16] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_query_proto_init() }
func file_query_proto_init() {
	if File_query_proto != nil {
		return
	}
	file_agent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertResponse); i {
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
		file_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScannerInfo); i {
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
		file_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfo); i {
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
		file_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedAlert); i {
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
		file_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRequest); i {
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
		file_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyResponse); i {
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
			RawDescriptor: file_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_proto_goTypes,
		DependencyIndexes: file_query_proto_depIdxs,
		EnumInfos:         file_query_proto_enumTypes,
		MessageInfos:      file_query_proto_msgTypes,
	}.Build()
	File_query_proto = out.File
	file_query_proto_rawDesc = nil
	file_query_proto_goTypes = nil
	file_query_proto_depIdxs = nil
}
