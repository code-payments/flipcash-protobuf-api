// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: profile/v1/model.proto

package profilepb

import (
	v11 "github.com/code-payments/flipcash-protobuf-api/generated/go/email/v1"
	v1 "github.com/code-payments/flipcash-protobuf-api/generated/go/phone/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type XProfile_VerifiedType int32

const (
	XProfile_NONE       XProfile_VerifiedType = 0
	XProfile_BLUE       XProfile_VerifiedType = 1
	XProfile_BUSINESS   XProfile_VerifiedType = 2
	XProfile_GOVERNMENT XProfile_VerifiedType = 3
)

// Enum value maps for XProfile_VerifiedType.
var (
	XProfile_VerifiedType_name = map[int32]string{
		0: "NONE",
		1: "BLUE",
		2: "BUSINESS",
		3: "GOVERNMENT",
	}
	XProfile_VerifiedType_value = map[string]int32{
		"NONE":       0,
		"BLUE":       1,
		"BUSINESS":   2,
		"GOVERNMENT": 3,
	}
)

func (x XProfile_VerifiedType) Enum() *XProfile_VerifiedType {
	p := new(XProfile_VerifiedType)
	*p = x
	return p
}

func (x XProfile_VerifiedType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XProfile_VerifiedType) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_v1_model_proto_enumTypes[0].Descriptor()
}

func (XProfile_VerifiedType) Type() protoreflect.EnumType {
	return &file_profile_v1_model_proto_enumTypes[0]
}

func (x XProfile_VerifiedType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use XProfile_VerifiedType.Descriptor instead.
func (XProfile_VerifiedType) EnumDescriptor() ([]byte, []int) {
	return file_profile_v1_model_proto_rawDescGZIP(), []int{2, 0}
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display name is the display name of the user (if found).
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Social profiles are links to external social accounts
	SocialProfiles []*SocialProfile `protobuf:"bytes,2,rep,name=social_profiles,json=socialProfiles,proto3" json:"social_profiles,omitempty"`
	// Phone number linked to this user. This is private and will only be returned
	// when the requesting user asks for their own profile
	PhoneNumber *v1.PhoneNumber `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// Email address linked to this user. This is private and will only be returned
	// when the requesting user asks for their own profile
	EmailAddress *v11.EmailAddress `protobuf:"bytes,4,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_profile_v1_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_profile_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *UserProfile) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UserProfile) GetSocialProfiles() []*SocialProfile {
	if x != nil {
		return x.SocialProfiles
	}
	return nil
}

func (x *UserProfile) GetPhoneNumber() *v1.PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *UserProfile) GetEmailAddress() *v11.EmailAddress {
	if x != nil {
		return x.EmailAddress
	}
	return nil
}

type SocialProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*SocialProfile_X
	Type isSocialProfile_Type `protobuf_oneof:"type"`
}

func (x *SocialProfile) Reset() {
	*x = SocialProfile{}
	mi := &file_profile_v1_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SocialProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialProfile) ProtoMessage() {}

func (x *SocialProfile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialProfile.ProtoReflect.Descriptor instead.
func (*SocialProfile) Descriptor() ([]byte, []int) {
	return file_profile_v1_model_proto_rawDescGZIP(), []int{1}
}

func (m *SocialProfile) GetType() isSocialProfile_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *SocialProfile) GetX() *XProfile {
	if x, ok := x.GetType().(*SocialProfile_X); ok {
		return x.X
	}
	return nil
}

type isSocialProfile_Type interface {
	isSocialProfile_Type()
}

type SocialProfile_X struct {
	X *XProfile `protobuf:"bytes,1,opt,name=x,proto3,oneof"`
}

func (*SocialProfile_X) isSocialProfile_Type() {}

type XProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user's ID on X
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The user's username on X
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The user's friendly name on X
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The user's description on X
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// URL to the user's X profile picture
	ProfilePicUrl string `protobuf:"bytes,5,opt,name=profile_pic_url,json=profilePicUrl,proto3" json:"profile_pic_url,omitempty"`
	// The type of X verification associated with the user
	VerifiedType XProfile_VerifiedType `protobuf:"varint,6,opt,name=verified_type,json=verifiedType,proto3,enum=flipcash.profile.v1.XProfile_VerifiedType" json:"verified_type,omitempty"`
	// The number of followers the user has on X
	FollowerCount uint32 `protobuf:"varint,7,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
}

func (x *XProfile) Reset() {
	*x = XProfile{}
	mi := &file_profile_v1_model_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *XProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XProfile) ProtoMessage() {}

func (x *XProfile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_model_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XProfile.ProtoReflect.Descriptor instead.
func (*XProfile) Descriptor() ([]byte, []int) {
	return file_profile_v1_model_proto_rawDescGZIP(), []int{2}
}

func (x *XProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *XProfile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *XProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XProfile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *XProfile) GetProfilePicUrl() string {
	if x != nil {
		return x.ProfilePicUrl
	}
	return ""
}

func (x *XProfile) GetVerifiedType() XProfile_VerifiedType {
	if x != nil {
		return x.VerifiedType
	}
	return XProfile_NONE
}

func (x *XProfile) GetFollowerCount() uint32 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

var File_profile_v1_model_proto protoreflect.FileDescriptor

var file_profile_v1_model_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x00, 0x18, 0x40, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x57, 0x0a, 0x0f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x69, 0x70,
	0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x92, 0x01, 0x04, 0x08, 0x00, 0x10, 0x01, 0x52, 0x0e, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x4b, 0x0a, 0x0d, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x01, 0x78, 0x42, 0x0b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22,
	0x84, 0x03, 0x0a, 0x08, 0x58, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x01, 0x18, 0x20, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x01, 0x18, 0x0f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0x18, 0x80, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x20, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x10, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x4f, 0x0a, 0x0d,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x53, 0x49,
	0x4e, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x4f, 0x56, 0x45, 0x52, 0x4e,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x42, 0x86, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x69, 0x6e, 0x63, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73, 0x68, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x61, 0x73,
	0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62,
	0xa2, 0x02, 0x0c, 0x46, 0x50, 0x42, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_v1_model_proto_rawDescOnce sync.Once
	file_profile_v1_model_proto_rawDescData = file_profile_v1_model_proto_rawDesc
)

func file_profile_v1_model_proto_rawDescGZIP() []byte {
	file_profile_v1_model_proto_rawDescOnce.Do(func() {
		file_profile_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_v1_model_proto_rawDescData)
	})
	return file_profile_v1_model_proto_rawDescData
}

var file_profile_v1_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_profile_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_profile_v1_model_proto_goTypes = []any{
	(XProfile_VerifiedType)(0), // 0: flipcash.profile.v1.XProfile.VerifiedType
	(*UserProfile)(nil),        // 1: flipcash.profile.v1.UserProfile
	(*SocialProfile)(nil),      // 2: flipcash.profile.v1.SocialProfile
	(*XProfile)(nil),           // 3: flipcash.profile.v1.XProfile
	(*v1.PhoneNumber)(nil),     // 4: flipcash.phone.v1.PhoneNumber
	(*v11.EmailAddress)(nil),   // 5: flipcash.email.v1.EmailAddress
}
var file_profile_v1_model_proto_depIdxs = []int32{
	2, // 0: flipcash.profile.v1.UserProfile.social_profiles:type_name -> flipcash.profile.v1.SocialProfile
	4, // 1: flipcash.profile.v1.UserProfile.phone_number:type_name -> flipcash.phone.v1.PhoneNumber
	5, // 2: flipcash.profile.v1.UserProfile.email_address:type_name -> flipcash.email.v1.EmailAddress
	3, // 3: flipcash.profile.v1.SocialProfile.x:type_name -> flipcash.profile.v1.XProfile
	0, // 4: flipcash.profile.v1.XProfile.verified_type:type_name -> flipcash.profile.v1.XProfile.VerifiedType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_profile_v1_model_proto_init() }
func file_profile_v1_model_proto_init() {
	if File_profile_v1_model_proto != nil {
		return
	}
	file_profile_v1_model_proto_msgTypes[1].OneofWrappers = []any{
		(*SocialProfile_X)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_v1_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_profile_v1_model_proto_goTypes,
		DependencyIndexes: file_profile_v1_model_proto_depIdxs,
		EnumInfos:         file_profile_v1_model_proto_enumTypes,
		MessageInfos:      file_profile_v1_model_proto_msgTypes,
	}.Build()
	File_profile_v1_model_proto = out.File
	file_profile_v1_model_proto_rawDesc = nil
	file_profile_v1_model_proto_goTypes = nil
	file_profile_v1_model_proto_depIdxs = nil
}
