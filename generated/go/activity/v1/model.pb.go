// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: activity/v1/model.proto

package activitypb

import (
	v1 "github.com/code-payments/flipcash-protobuf-api/generated/go/common/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ActivityFeedType enables multiple activity feeds, where notifications may be
// split across different parts of the app
type ActivityFeedType int32

const (
	ActivityFeedType_UNKNOWN             ActivityFeedType = 0
	ActivityFeedType_TRANSACTION_HISTORY ActivityFeedType = 1 // Activity feed displayed under the Balance tab
)

// Enum value maps for ActivityFeedType.
var (
	ActivityFeedType_name = map[int32]string{
		0: "UNKNOWN",
		1: "TRANSACTION_HISTORY",
	}
	ActivityFeedType_value = map[string]int32{
		"UNKNOWN":             0,
		"TRANSACTION_HISTORY": 1,
	}
)

func (x ActivityFeedType) Enum() *ActivityFeedType {
	p := new(ActivityFeedType)
	*p = x
	return p
}

func (x ActivityFeedType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityFeedType) Descriptor() protoreflect.EnumDescriptor {
	return file_activity_v1_model_proto_enumTypes[0].Descriptor()
}

func (ActivityFeedType) Type() protoreflect.EnumType {
	return &file_activity_v1_model_proto_enumTypes[0]
}

func (x ActivityFeedType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityFeedType.Descriptor instead.
func (ActivityFeedType) EnumDescriptor() ([]byte, []int) {
	return file_activity_v1_model_proto_rawDescGZIP(), []int{0}
}

// The ID of the notification, which is guaranteed to be consistent for grouped
// events. Updates to a notification with the same ID should result in re-ordering
// within the activity feed using the latest content.
type NotificationId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NotificationId) Reset() {
	*x = NotificationId{}
	mi := &file_activity_v1_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationId) ProtoMessage() {}

func (x *NotificationId) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationId.ProtoReflect.Descriptor instead.
func (*NotificationId) Descriptor() ([]byte, []int) {
	return file_activity_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationId) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Notification is a message that is displayed in an activity feed
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of this notification
	Id *NotificationId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The localized title text for the notification
	LocalizedText string `protobuf:"bytes,2,opt,name=localized_text,json=localizedText,proto3" json:"localized_text,omitempty"`
	// If a payment applies, the amount that was paid
	Amount *v1.PaymentAmount `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// The timestamp of this notification
	Ts *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=ts,proto3" json:"ts,omitempty"`
	// Additional metadata for this notification specific to the notification
	//
	// Types that are assignable to AdditionalMetadata:
	//
	//	*Notification_WelcomeBonus
	//	*Notification_GaveUsdc
	//	*Notification_ReceivedUsdc
	AdditionalMetadata isNotification_AdditionalMetadata `protobuf_oneof:"additional_metadata"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_activity_v1_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_activity_v1_model_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetId() *NotificationId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Notification) GetLocalizedText() string {
	if x != nil {
		return x.LocalizedText
	}
	return ""
}

func (x *Notification) GetAmount() *v1.PaymentAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Notification) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (m *Notification) GetAdditionalMetadata() isNotification_AdditionalMetadata {
	if m != nil {
		return m.AdditionalMetadata
	}
	return nil
}

func (x *Notification) GetWelcomeBonus() *WelcomeBonusNotificationMetadata {
	if x, ok := x.GetAdditionalMetadata().(*Notification_WelcomeBonus); ok {
		return x.WelcomeBonus
	}
	return nil
}

func (x *Notification) GetGaveUsdc() *GaveUsdcNotificationMetadata {
	if x, ok := x.GetAdditionalMetadata().(*Notification_GaveUsdc); ok {
		return x.GaveUsdc
	}
	return nil
}

func (x *Notification) GetReceivedUsdc() *ReceivedUsdcNotificationMetadata {
	if x, ok := x.GetAdditionalMetadata().(*Notification_ReceivedUsdc); ok {
		return x.ReceivedUsdc
	}
	return nil
}

type isNotification_AdditionalMetadata interface {
	isNotification_AdditionalMetadata()
}

type Notification_WelcomeBonus struct {
	WelcomeBonus *WelcomeBonusNotificationMetadata `protobuf:"bytes,5,opt,name=welcome_bonus,json=welcomeBonus,proto3,oneof"`
}

type Notification_GaveUsdc struct {
	GaveUsdc *GaveUsdcNotificationMetadata `protobuf:"bytes,6,opt,name=gave_usdc,json=gaveUsdc,proto3,oneof"`
}

type Notification_ReceivedUsdc struct {
	ReceivedUsdc *ReceivedUsdcNotificationMetadata `protobuf:"bytes,7,opt,name=received_usdc,json=receivedUsdc,proto3,oneof"`
}

func (*Notification_WelcomeBonus) isNotification_AdditionalMetadata() {}

func (*Notification_GaveUsdc) isNotification_AdditionalMetadata() {}

func (*Notification_ReceivedUsdc) isNotification_AdditionalMetadata() {}

type WelcomeBonusNotificationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WelcomeBonusNotificationMetadata) Reset() {
	*x = WelcomeBonusNotificationMetadata{}
	mi := &file_activity_v1_model_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WelcomeBonusNotificationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeBonusNotificationMetadata) ProtoMessage() {}

func (x *WelcomeBonusNotificationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_model_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeBonusNotificationMetadata.ProtoReflect.Descriptor instead.
func (*WelcomeBonusNotificationMetadata) Descriptor() ([]byte, []int) {
	return file_activity_v1_model_proto_rawDescGZIP(), []int{2}
}

type GaveUsdcNotificationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GaveUsdcNotificationMetadata) Reset() {
	*x = GaveUsdcNotificationMetadata{}
	mi := &file_activity_v1_model_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GaveUsdcNotificationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaveUsdcNotificationMetadata) ProtoMessage() {}

func (x *GaveUsdcNotificationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_model_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaveUsdcNotificationMetadata.ProtoReflect.Descriptor instead.
func (*GaveUsdcNotificationMetadata) Descriptor() ([]byte, []int) {
	return file_activity_v1_model_proto_rawDescGZIP(), []int{3}
}

type ReceivedUsdcNotificationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceivedUsdcNotificationMetadata) Reset() {
	*x = ReceivedUsdcNotificationMetadata{}
	mi := &file_activity_v1_model_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceivedUsdcNotificationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedUsdcNotificationMetadata) ProtoMessage() {}

func (x *ReceivedUsdcNotificationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_model_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedUsdcNotificationMetadata.ProtoReflect.Descriptor instead.
func (*ReceivedUsdcNotificationMetadata) Descriptor() ([]byte, []int) {
	return file_activity_v1_model_proto_rawDescGZIP(), []int{4}
}

var File_activity_v1_model_proto protoreflect.FileDescriptor

var file_activity_v1_model_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66, 0x6c, 0x69, 0x70, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x31, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x7a, 0x04, 0x10, 0x20, 0x18, 0x20, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x9a, 0x04, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x02, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0xb2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x02, 0x74, 0x73, 0x12, 0x5d, 0x0a, 0x0d, 0x77, 0x65, 0x6c,
	0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x42,
	0x6f, 0x6e, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0c, 0x77, 0x65, 0x6c, 0x63,
	0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x51, 0x0a, 0x09, 0x67, 0x61, 0x76, 0x65,
	0x5f, 0x75, 0x73, 0x64, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x66, 0x6c,
	0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x61, 0x76, 0x65, 0x55, 0x73, 0x64, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x08, 0x67, 0x61, 0x76, 0x65, 0x55, 0x73, 0x64, 0x63, 0x12, 0x5d, 0x0a, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x64, 0x63, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x55, 0x73, 0x64, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x55, 0x73, 0x64, 0x63, 0x42, 0x15, 0x0a, 0x13, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x22, 0x0a, 0x20, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x61, 0x76, 0x65, 0x55, 0x73, 0x64,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x20, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x55, 0x73, 0x64, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x38, 0x0a, 0x10, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52,
	0x59, 0x10, 0x01, 0x42, 0x8b, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x69, 0x6e, 0x63, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x52, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62,
	0xa2, 0x02, 0x0e, 0x46, 0x43, 0x50, 0x42, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_v1_model_proto_rawDescOnce sync.Once
	file_activity_v1_model_proto_rawDescData = file_activity_v1_model_proto_rawDesc
)

func file_activity_v1_model_proto_rawDescGZIP() []byte {
	file_activity_v1_model_proto_rawDescOnce.Do(func() {
		file_activity_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_v1_model_proto_rawDescData)
	})
	return file_activity_v1_model_proto_rawDescData
}

var file_activity_v1_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_activity_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_activity_v1_model_proto_goTypes = []any{
	(ActivityFeedType)(0),                    // 0: flipcash.activity.v1.ActivityFeedType
	(*NotificationId)(nil),                   // 1: flipcash.activity.v1.NotificationId
	(*Notification)(nil),                     // 2: flipcash.activity.v1.Notification
	(*WelcomeBonusNotificationMetadata)(nil), // 3: flipcash.activity.v1.WelcomeBonusNotificationMetadata
	(*GaveUsdcNotificationMetadata)(nil),     // 4: flipcash.activity.v1.GaveUsdcNotificationMetadata
	(*ReceivedUsdcNotificationMetadata)(nil), // 5: flipcash.activity.v1.ReceivedUsdcNotificationMetadata
	(*v1.PaymentAmount)(nil),                 // 6: flipcash.common.v1.PaymentAmount
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
}
var file_activity_v1_model_proto_depIdxs = []int32{
	1, // 0: flipcash.activity.v1.Notification.id:type_name -> flipcash.activity.v1.NotificationId
	6, // 1: flipcash.activity.v1.Notification.amount:type_name -> flipcash.common.v1.PaymentAmount
	7, // 2: flipcash.activity.v1.Notification.ts:type_name -> google.protobuf.Timestamp
	3, // 3: flipcash.activity.v1.Notification.welcome_bonus:type_name -> flipcash.activity.v1.WelcomeBonusNotificationMetadata
	4, // 4: flipcash.activity.v1.Notification.gave_usdc:type_name -> flipcash.activity.v1.GaveUsdcNotificationMetadata
	5, // 5: flipcash.activity.v1.Notification.received_usdc:type_name -> flipcash.activity.v1.ReceivedUsdcNotificationMetadata
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_activity_v1_model_proto_init() }
func file_activity_v1_model_proto_init() {
	if File_activity_v1_model_proto != nil {
		return
	}
	file_activity_v1_model_proto_msgTypes[1].OneofWrappers = []any{
		(*Notification_WelcomeBonus)(nil),
		(*Notification_GaveUsdc)(nil),
		(*Notification_ReceivedUsdc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_activity_v1_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_activity_v1_model_proto_goTypes,
		DependencyIndexes: file_activity_v1_model_proto_depIdxs,
		EnumInfos:         file_activity_v1_model_proto_enumTypes,
		MessageInfos:      file_activity_v1_model_proto_msgTypes,
	}.Build()
	File_activity_v1_model_proto = out.File
	file_activity_v1_model_proto_rawDesc = nil
	file_activity_v1_model_proto_goTypes = nil
	file_activity_v1_model_proto_depIdxs = nil
}
