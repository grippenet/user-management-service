// Code generated by protoc-gen-go. DO NOT EDIT.
// source: study.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Study struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string          `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	SecretKey            string          `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Status               string          `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Props                *Study_Props    `protobuf:"bytes,5,opt,name=props,proto3" json:"props,omitempty"`
	Rules                []*Expression   `protobuf:"bytes,6,rep,name=rules,proto3" json:"rules,omitempty"`
	Members              []*Study_Member `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Study) Reset()         { *m = Study{} }
func (m *Study) String() string { return proto.CompactTextString(m) }
func (*Study) ProtoMessage()    {}
func (*Study) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59e246e4c4c09f0, []int{0}
}

func (m *Study) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Study.Unmarshal(m, b)
}
func (m *Study) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Study.Marshal(b, m, deterministic)
}
func (m *Study) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Study.Merge(m, src)
}
func (m *Study) XXX_Size() int {
	return xxx_messageInfo_Study.Size(m)
}
func (m *Study) XXX_DiscardUnknown() {
	xxx_messageInfo_Study.DiscardUnknown(m)
}

var xxx_messageInfo_Study proto.InternalMessageInfo

func (m *Study) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Study) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Study) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *Study) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Study) GetProps() *Study_Props {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *Study) GetRules() []*Expression {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Study) GetMembers() []*Study_Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type Study_Props struct {
	Name                 []*LocalisedObject `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
	Description          []*LocalisedObject `protobuf:"bytes,2,rep,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Study_Props) Reset()         { *m = Study_Props{} }
func (m *Study_Props) String() string { return proto.CompactTextString(m) }
func (*Study_Props) ProtoMessage()    {}
func (*Study_Props) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59e246e4c4c09f0, []int{0, 0}
}

func (m *Study_Props) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Study_Props.Unmarshal(m, b)
}
func (m *Study_Props) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Study_Props.Marshal(b, m, deterministic)
}
func (m *Study_Props) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Study_Props.Merge(m, src)
}
func (m *Study_Props) XXX_Size() int {
	return xxx_messageInfo_Study_Props.Size(m)
}
func (m *Study_Props) XXX_DiscardUnknown() {
	xxx_messageInfo_Study_Props.DiscardUnknown(m)
}

var xxx_messageInfo_Study_Props proto.InternalMessageInfo

func (m *Study_Props) GetName() []*LocalisedObject {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Study_Props) GetDescription() []*LocalisedObject {
	if m != nil {
		return m.Description
	}
	return nil
}

type Study_Member struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Study_Member) Reset()         { *m = Study_Member{} }
func (m *Study_Member) String() string { return proto.CompactTextString(m) }
func (*Study_Member) ProtoMessage()    {}
func (*Study_Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59e246e4c4c09f0, []int{0, 1}
}

func (m *Study_Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Study_Member.Unmarshal(m, b)
}
func (m *Study_Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Study_Member.Marshal(b, m, deterministic)
}
func (m *Study_Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Study_Member.Merge(m, src)
}
func (m *Study_Member) XXX_Size() int {
	return xxx_messageInfo_Study_Member.Size(m)
}
func (m *Study_Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Study_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Study_Member proto.InternalMessageInfo

func (m *Study_Member) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Study_Member) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Study_Member) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AssignedSurvey struct {
	SurveyKey            string   `protobuf:"bytes,1,opt,name=survey_key,json=surveyKey,proto3" json:"survey_key,omitempty"`
	ValidFrom            int64    `protobuf:"varint,2,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ValidUntil           int64    `protobuf:"varint,3,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	StudyKey             string   `protobuf:"bytes,4,opt,name=study_key,json=studyKey,proto3" json:"study_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignedSurvey) Reset()         { *m = AssignedSurvey{} }
func (m *AssignedSurvey) String() string { return proto.CompactTextString(m) }
func (*AssignedSurvey) ProtoMessage()    {}
func (*AssignedSurvey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59e246e4c4c09f0, []int{1}
}

func (m *AssignedSurvey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignedSurvey.Unmarshal(m, b)
}
func (m *AssignedSurvey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignedSurvey.Marshal(b, m, deterministic)
}
func (m *AssignedSurvey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignedSurvey.Merge(m, src)
}
func (m *AssignedSurvey) XXX_Size() int {
	return xxx_messageInfo_AssignedSurvey.Size(m)
}
func (m *AssignedSurvey) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignedSurvey.DiscardUnknown(m)
}

var xxx_messageInfo_AssignedSurvey proto.InternalMessageInfo

func (m *AssignedSurvey) GetSurveyKey() string {
	if m != nil {
		return m.SurveyKey
	}
	return ""
}

func (m *AssignedSurvey) GetValidFrom() int64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *AssignedSurvey) GetValidUntil() int64 {
	if m != nil {
		return m.ValidUntil
	}
	return 0
}

func (m *AssignedSurvey) GetStudyKey() string {
	if m != nil {
		return m.StudyKey
	}
	return ""
}

type AssignedSurveys struct {
	Surveys              []*AssignedSurvey `protobuf:"bytes,1,rep,name=surveys,proto3" json:"surveys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AssignedSurveys) Reset()         { *m = AssignedSurveys{} }
func (m *AssignedSurveys) String() string { return proto.CompactTextString(m) }
func (*AssignedSurveys) ProtoMessage()    {}
func (*AssignedSurveys) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59e246e4c4c09f0, []int{2}
}

func (m *AssignedSurveys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignedSurveys.Unmarshal(m, b)
}
func (m *AssignedSurveys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignedSurveys.Marshal(b, m, deterministic)
}
func (m *AssignedSurveys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignedSurveys.Merge(m, src)
}
func (m *AssignedSurveys) XXX_Size() int {
	return xxx_messageInfo_AssignedSurveys.Size(m)
}
func (m *AssignedSurveys) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignedSurveys.DiscardUnknown(m)
}

var xxx_messageInfo_AssignedSurveys proto.InternalMessageInfo

func (m *AssignedSurveys) GetSurveys() []*AssignedSurvey {
	if m != nil {
		return m.Surveys
	}
	return nil
}

func init() {
	proto.RegisterType((*Study)(nil), "inf.study.Study")
	proto.RegisterType((*Study_Props)(nil), "inf.study.Study.Props")
	proto.RegisterType((*Study_Member)(nil), "inf.study.Study.Member")
	proto.RegisterType((*AssignedSurvey)(nil), "inf.study.AssignedSurvey")
	proto.RegisterType((*AssignedSurveys)(nil), "inf.study.AssignedSurveys")
}

func init() {
	proto.RegisterFile("study.proto", fileDescriptor_b59e246e4c4c09f0)
}

var fileDescriptor_b59e246e4c4c09f0 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0xd8, 0xa9, 0xc7, 0xa8, 0x54, 0x73, 0x68, 0x8d, 0x2b, 0x44, 0x94, 0x53, 0x0e,
	0xc8, 0xa8, 0xed, 0x99, 0x03, 0x48, 0x54, 0x42, 0x01, 0x01, 0xae, 0xb8, 0x70, 0x89, 0x9c, 0xec,
	0x14, 0x2d, 0xd8, 0x5e, 0x6b, 0x67, 0x5d, 0xc8, 0x3f, 0xf0, 0x51, 0x7c, 0x1a, 0xf2, 0xac, 0x1b,
	0x52, 0x71, 0xe0, 0x36, 0xf3, 0xde, 0x1b, 0xbf, 0xf1, 0x9b, 0x85, 0x94, 0x5d, 0xaf, 0x76, 0x45,
	0x67, 0x8d, 0x33, 0x98, 0xe8, 0xf6, 0xb6, 0x10, 0x20, 0x3f, 0xa1, 0x9f, 0x9d, 0x25, 0x66, 0x6d,
	0x5a, 0x4f, 0xe6, 0x8f, 0xb8, 0xb7, 0x77, 0x34, 0x4a, 0x17, 0xbf, 0x43, 0x88, 0x6e, 0x06, 0x25,
	0x1e, 0xc3, 0x44, 0xab, 0x2c, 0x98, 0x07, 0xcb, 0xa4, 0x9c, 0x68, 0x85, 0x27, 0x10, 0x7e, 0xa7,
	0x5d, 0x36, 0x11, 0x60, 0x28, 0xf1, 0x29, 0x00, 0xd3, 0xd6, 0x92, 0x5b, 0x0f, 0x44, 0x28, 0x44,
	0xe2, 0x91, 0x15, 0xed, 0xf0, 0x14, 0x62, 0x76, 0x95, 0xeb, 0x39, 0x9b, 0x0a, 0x35, 0x76, 0xf8,
	0x1c, 0xa2, 0xce, 0x9a, 0x8e, 0xb3, 0x68, 0x1e, 0x2c, 0xd3, 0xcb, 0xd3, 0x62, 0xbf, 0x5d, 0x21,
	0xce, 0xc5, 0xc7, 0x81, 0x2d, 0xbd, 0x08, 0x2f, 0x20, 0xb2, 0x7d, 0x4d, 0x9c, 0xc5, 0xf3, 0x70,
	0x99, 0x5e, 0x9e, 0x8b, 0xfa, 0xef, 0x4f, 0x70, 0xf1, 0x66, 0x5f, 0x97, 0x5e, 0x89, 0x17, 0x30,
	0x6b, 0xa8, 0xd9, 0x90, 0xe5, 0x6c, 0x26, 0x43, 0x67, 0xff, 0x58, 0xbc, 0x17, 0xbe, 0xbc, 0xd7,
	0xe5, 0x3f, 0x20, 0x12, 0x57, 0x7c, 0x01, 0xd3, 0xb6, 0x6a, 0x28, 0x0b, 0x0e, 0xdc, 0xc6, 0x80,
	0xde, 0x99, 0x6d, 0x55, 0x6b, 0x26, 0xf5, 0x61, 0xf3, 0x8d, 0xb6, 0xae, 0x14, 0x21, 0xbe, 0x84,
	0x54, 0x11, 0x6f, 0xad, 0xee, 0x9c, 0x36, 0x6d, 0x36, 0xf9, 0xff, 0xdc, 0xa1, 0x3e, 0xff, 0x04,
	0xb1, 0xdf, 0x05, 0xcf, 0x60, 0xd6, 0x33, 0xd9, 0xf5, 0x3e, 0xf4, 0x78, 0x68, 0xdf, 0x2a, 0x44,
	0x98, 0x5a, 0x53, 0xd3, 0x98, 0xbc, 0xd4, 0x98, 0xc3, 0xd1, 0xc0, 0xca, 0xaa, 0x3e, 0xf8, 0x7d,
	0xbf, 0xf8, 0x15, 0xc0, 0xf1, 0x2b, 0x66, 0xfd, 0xb5, 0x25, 0x75, 0x23, 0x2b, 0xc8, 0xa5, 0xa4,
	0x92, 0x4b, 0x05, 0xe3, 0xa5, 0x04, 0x59, 0x79, 0xfa, 0xae, 0xaa, 0xb5, 0x5a, 0xdf, 0x5a, 0xd3,
	0x88, 0x4f, 0x58, 0x26, 0x82, 0x5c, 0x5b, 0xd3, 0xe0, 0x33, 0x48, 0x3d, 0xdd, 0xb7, 0x4e, 0xd7,
	0xe2, 0x17, 0x96, 0x7e, 0xe2, 0xf3, 0x80, 0xe0, 0x39, 0x24, 0x12, 0xae, 0x7c, 0xdd, 0x1f, 0xfb,
	0x48, 0x80, 0x15, 0xed, 0x16, 0xd7, 0xf0, 0xf8, 0xe1, 0x36, 0x8c, 0x57, 0x30, 0xf3, 0xe6, 0x3c,
	0xe6, 0xfc, 0xe4, 0xe0, 0x40, 0x0f, 0xc5, 0xe5, 0xbd, 0xf2, 0x75, 0xf4, 0x25, 0xac, 0x3a, 0xbd,
	0x89, 0xe5, 0x9d, 0x5e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x71, 0x46, 0x25, 0x3d, 0xe1, 0x02,
	0x00, 0x00,
}
