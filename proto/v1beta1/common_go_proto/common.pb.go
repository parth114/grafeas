// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1beta1/common.proto

package common_go_proto

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

// Kind represents the kinds of notes supported.
type NoteKind int32

const (
	// Unknown.
	NoteKind_NOTE_KIND_UNSPECIFIED NoteKind = 0
	// The note and occurrence represent a package vulnerability.
	NoteKind_VULNERABILITY NoteKind = 1
	// The note and occurrence assert build provenance.
	NoteKind_BUILD NoteKind = 2
	// This represents an image basis relationship.
	NoteKind_IMAGE NoteKind = 3
	// This represents a package installed via a package manager.
	NoteKind_PACKAGE NoteKind = 4
	// The note and occurrence track deployment events.
	NoteKind_DEPLOYMENT NoteKind = 5
	// The note and occurrence track the initial discovery status of a resource.
	NoteKind_DISCOVERY NoteKind = 6
	// This represents a logical "role" that can attest to artifacts.
	NoteKind_ATTESTATION NoteKind = 7
)

var NoteKind_name = map[int32]string{
	0: "NOTE_KIND_UNSPECIFIED",
	1: "VULNERABILITY",
	2: "BUILD",
	3: "IMAGE",
	4: "PACKAGE",
	5: "DEPLOYMENT",
	6: "DISCOVERY",
	7: "ATTESTATION",
}

var NoteKind_value = map[string]int32{
	"NOTE_KIND_UNSPECIFIED": 0,
	"VULNERABILITY":         1,
	"BUILD":                 2,
	"IMAGE":                 3,
	"PACKAGE":               4,
	"DEPLOYMENT":            5,
	"DISCOVERY":             6,
	"ATTESTATION":           7,
}

func (x NoteKind) String() string {
	return proto.EnumName(NoteKind_name, int32(x))
}

func (NoteKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a10ef3af5a300e59, []int{0}
}

// Metadata for any related URL information.
type RelatedUrl struct {
	// Specific URL associated with the resource.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Label to describe usage of the URL.
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelatedUrl) Reset()         { *m = RelatedUrl{} }
func (m *RelatedUrl) String() string { return proto.CompactTextString(m) }
func (*RelatedUrl) ProtoMessage()    {}
func (*RelatedUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10ef3af5a300e59, []int{0}
}

func (m *RelatedUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelatedUrl.Unmarshal(m, b)
}
func (m *RelatedUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelatedUrl.Marshal(b, m, deterministic)
}
func (m *RelatedUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelatedUrl.Merge(m, src)
}
func (m *RelatedUrl) XXX_Size() int {
	return xxx_messageInfo_RelatedUrl.Size(m)
}
func (m *RelatedUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_RelatedUrl.DiscardUnknown(m)
}

var xxx_messageInfo_RelatedUrl proto.InternalMessageInfo

func (m *RelatedUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RelatedUrl) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.NoteKind", NoteKind_name, NoteKind_value)
	proto.RegisterType((*RelatedUrl)(nil), "grafeas.v1beta1.RelatedUrl")
}

func init() { proto.RegisterFile("proto/v1beta1/common.proto", fileDescriptor_a10ef3af5a300e59) }

var fileDescriptor_a10ef3af5a300e59 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x14, 0x45, 0xa5, 0xb5, 0xad, 0x7d, 0x4d, 0xed, 0x38, 0xd1, 0xa4, 0xba, 0x32, 0xae, 0x8c, 0x0b,
	0x9a, 0x46, 0x17, 0x6e, 0x07, 0x18, 0x9b, 0x09, 0x74, 0x20, 0x30, 0x90, 0xe0, 0x86, 0x40, 0x3b,
	0x22, 0x09, 0x74, 0x0c, 0x52, 0x7f, 0xc2, 0xbf, 0xf0, 0x4b, 0x0d, 0xd8, 0xb8, 0x30, 0xae, 0xde,
	0x3d, 0xe7, 0xae, 0xde, 0x85, 0xab, 0xb7, 0x5a, 0x35, 0x6a, 0xf1, 0xb1, 0xcc, 0x64, 0x93, 0x2e,
	0x17, 0x1b, 0x55, 0x55, 0x6a, 0xa7, 0x77, 0x12, 0xcf, 0xf2, 0x3a, 0x7d, 0x91, 0xe9, 0xbb, 0x7e,
	0x68, 0x6f, 0x1e, 0x00, 0x7c, 0x59, 0xa6, 0x8d, 0xdc, 0x86, 0x75, 0x89, 0x11, 0xf4, 0xf7, 0x75,
	0x39, 0xd7, 0xae, 0xb5, 0xdb, 0xb1, 0xdf, 0x46, 0x7c, 0x0e, 0x83, 0x32, 0xcd, 0x64, 0x39, 0xef,
	0x75, 0xee, 0x07, 0xee, 0x3e, 0x35, 0x38, 0xe1, 0xaa, 0x91, 0x76, 0xb1, 0xdb, 0xe2, 0x4b, 0xb8,
	0xe0, 0xae, 0xa0, 0x89, 0xcd, 0xb8, 0x95, 0x84, 0x3c, 0xf0, 0xa8, 0xc9, 0x9e, 0x18, 0xb5, 0xd0,
	0x11, 0x3e, 0x83, 0x69, 0x14, 0x3a, 0x9c, 0xfa, 0xc4, 0x60, 0x0e, 0x13, 0x31, 0xd2, 0xf0, 0x18,
	0x06, 0x46, 0xc8, 0x1c, 0x0b, 0xf5, 0xda, 0xc8, 0xd6, 0x64, 0x45, 0x51, 0x1f, 0x4f, 0x60, 0xe4,
	0x11, 0xd3, 0x6e, 0xe1, 0x18, 0x9f, 0x02, 0x58, 0xd4, 0x73, 0xdc, 0x78, 0x4d, 0xb9, 0x40, 0x03,
	0x3c, 0x85, 0xb1, 0xc5, 0x02, 0xd3, 0x8d, 0xa8, 0x1f, 0xa3, 0x21, 0x9e, 0xc1, 0x84, 0x08, 0x41,
	0x03, 0x41, 0x04, 0x73, 0x39, 0x1a, 0x19, 0x11, 0xe0, 0x42, 0xe9, 0x7f, 0x3e, 0xf3, 0xb4, 0xe7,
	0xc7, 0xbc, 0x68, 0x5e, 0xf7, 0x99, 0xbe, 0x51, 0xd5, 0xe2, 0xd0, 0xfe, 0xde, 0xff, 0x36, 0x4a,
	0x72, 0x95, 0x74, 0xfe, 0xab, 0xd7, 0x5f, 0xf9, 0x24, 0x1b, 0x76, 0x70, 0xff, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xbf, 0x1a, 0xf1, 0x9a, 0x51, 0x01, 0x00, 0x00,
}
