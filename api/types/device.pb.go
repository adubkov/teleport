// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: teleport/legacy/types/device.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DeviceV1 is the resource representation of teleport.devicetrust.v1.Device.
type DeviceV1 struct {
	// Header is the common resource header.
	//
	// - Kind is always "device".
	// - SubKind is unused.
	// - Version is equivalent to teleport.devicetrust.v1.Device.api_version.
	// - Metadata.Name is equivalent to teleport.devicetrust.v1.Device.Id.
	ResourceHeader `protobuf:"bytes,1,opt,name=Header,proto3,embedded=Header" json:""`
	// Specification of the device.
	Spec                 *DeviceSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeviceV1) Reset()         { *m = DeviceV1{} }
func (m *DeviceV1) String() string { return proto.CompactTextString(m) }
func (*DeviceV1) ProtoMessage()    {}
func (*DeviceV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_aceaef1b58496e7d, []int{0}
}
func (m *DeviceV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceV1.Merge(m, src)
}
func (m *DeviceV1) XXX_Size() int {
	return m.Size()
}
func (m *DeviceV1) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceV1.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceV1 proto.InternalMessageInfo

// DeviceSpec is a device specification.
// Roughly matches teleport.devicetrust.v1.Device, with some fields changed for
// better UX.
type DeviceSpec struct {
	OsType               string                 `protobuf:"bytes,1,opt,name=os_type,json=osType,proto3" json:"os_type"`
	AssetTag             string                 `protobuf:"bytes,2,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag"`
	CreateTime           *time.Time             `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3,stdtime" json:"create_time"`
	UpdateTime           *time.Time             `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3,stdtime" json:"update_time"`
	EnrollStatus         string                 `protobuf:"bytes,5,opt,name=enroll_status,json=enrollStatus,proto3" json:"enroll_status"`
	Credential           *DeviceCredential      `protobuf:"bytes,6,opt,name=credential,proto3" json:"credential,omitempty"`
	CollectedData        []*DeviceCollectedData `protobuf:"bytes,7,rep,name=collected_data,json=collectedData,proto3" json:"collected_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DeviceSpec) Reset()         { *m = DeviceSpec{} }
func (m *DeviceSpec) String() string { return proto.CompactTextString(m) }
func (*DeviceSpec) ProtoMessage()    {}
func (*DeviceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_aceaef1b58496e7d, []int{1}
}
func (m *DeviceSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSpec.Merge(m, src)
}
func (m *DeviceSpec) XXX_Size() int {
	return m.Size()
}
func (m *DeviceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSpec proto.InternalMessageInfo

// DeviceCredential is the resource representation of
// teleport.devicetrust.v1.DeviceCredential.
type DeviceCredential struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	PublicKeyDer         []byte   `protobuf:"bytes,2,opt,name=public_key_der,json=publicKeyDer,proto3" json:"public_key_der"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceCredential) Reset()         { *m = DeviceCredential{} }
func (m *DeviceCredential) String() string { return proto.CompactTextString(m) }
func (*DeviceCredential) ProtoMessage()    {}
func (*DeviceCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_aceaef1b58496e7d, []int{2}
}
func (m *DeviceCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCredential.Merge(m, src)
}
func (m *DeviceCredential) XXX_Size() int {
	return m.Size()
}
func (m *DeviceCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCredential.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCredential proto.InternalMessageInfo

// DeviceCollectedData is the resource representation of
// teleport.devicetrust.v1.DeviceCollectedData.
type DeviceCollectedData struct {
	CollectTime          *time.Time `protobuf:"bytes,1,opt,name=collect_time,json=collectTime,proto3,stdtime" json:"collect_time"`
	RecordTime           *time.Time `protobuf:"bytes,2,opt,name=record_time,json=recordTime,proto3,stdtime" json:"record_time"`
	OsType               string     `protobuf:"bytes,3,opt,name=os_type,json=osType,proto3" json:"os_type"`
	SerialNumber         string     `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeviceCollectedData) Reset()         { *m = DeviceCollectedData{} }
func (m *DeviceCollectedData) String() string { return proto.CompactTextString(m) }
func (*DeviceCollectedData) ProtoMessage()    {}
func (*DeviceCollectedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_aceaef1b58496e7d, []int{3}
}
func (m *DeviceCollectedData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceCollectedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceCollectedData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceCollectedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCollectedData.Merge(m, src)
}
func (m *DeviceCollectedData) XXX_Size() int {
	return m.Size()
}
func (m *DeviceCollectedData) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCollectedData.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCollectedData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeviceV1)(nil), "types.DeviceV1")
	proto.RegisterType((*DeviceSpec)(nil), "types.DeviceSpec")
	proto.RegisterType((*DeviceCredential)(nil), "types.DeviceCredential")
	proto.RegisterType((*DeviceCollectedData)(nil), "types.DeviceCollectedData")
}

func init() {
	proto.RegisterFile("teleport/legacy/types/device.proto", fileDescriptor_aceaef1b58496e7d)
}

var fileDescriptor_aceaef1b58496e7d = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0xda, 0xad, 0x5b, 0xdd, 0x76, 0x62, 0xde, 0x60, 0xd1, 0x40, 0xf5, 0xa8, 0x78, 0x40,
	0x80, 0x1a, 0x01, 0x12, 0x42, 0x42, 0x48, 0x28, 0xec, 0x01, 0x09, 0x09, 0x41, 0x36, 0xf1, 0xc0,
	0x4b, 0xe4, 0x3a, 0x97, 0x60, 0x91, 0xd6, 0x91, 0xe3, 0x4c, 0xf4, 0x2f, 0xf8, 0x08, 0x3e, 0x66,
	0x8f, 0xfb, 0x02, 0x03, 0x7b, 0xf4, 0x03, 0xdf, 0x80, 0x62, 0xb7, 0x5b, 0x32, 0x0d, 0xb1, 0x97,
	0x28, 0xf7, 0xdc, 0x73, 0x8e, 0x6f, 0xee, 0xbd, 0x31, 0x1a, 0x29, 0xc8, 0x20, 0x17, 0x52, 0x05,
	0x19, 0xa4, 0x94, 0xcd, 0x03, 0x35, 0xcf, 0xa1, 0x08, 0x12, 0x38, 0xe6, 0x0c, 0xc6, 0xb9, 0x14,
	0x4a, 0xe0, 0x35, 0x8b, 0xed, 0xed, 0xa4, 0x22, 0x15, 0x16, 0x09, 0xaa, 0x37, 0x97, 0xdc, 0x23,
	0xa9, 0x10, 0x69, 0x06, 0x81, 0x8d, 0x26, 0xe5, 0xe7, 0x40, 0xf1, 0x29, 0x14, 0x8a, 0x4e, 0xf3,
	0x05, 0xe1, 0xee, 0xd5, 0x27, 0xd8, 0xa7, 0xa3, 0x8c, 0xbe, 0xa1, 0x8d, 0x03, 0x7b, 0xe0, 0xc7,
	0xc7, 0xf8, 0x05, 0xea, 0xbc, 0x01, 0x9a, 0x80, 0xf4, 0xbd, 0x7d, 0xef, 0x7e, 0xef, 0xc9, 0xcd,
	0xb1, 0x63, 0x46, 0x50, 0x88, 0x52, 0x32, 0x70, 0xc9, 0xb0, 0x7f, 0xa2, 0xc9, 0xca, 0xa9, 0x26,
	0x9e, 0xd1, 0x64, 0x25, 0x5a, 0x48, 0x70, 0x80, 0x56, 0x8b, 0x1c, 0x98, 0xbf, 0x66, 0xa5, 0x5b,
	0x0b, 0xa9, 0xf3, 0x3e, 0xcc, 0x81, 0x85, 0x1b, 0x46, 0x13, 0x4b, 0x89, 0xec, 0x73, 0xf4, 0xa7,
	0x8d, 0xd0, 0x45, 0x1a, 0xdf, 0x43, 0xeb, 0xa2, 0x88, 0x2b, 0x95, 0x3d, 0xbd, 0x1b, 0xf6, 0x8c,
	0x26, 0x4b, 0x28, 0xea, 0x88, 0xe2, 0x68, 0x9e, 0x03, 0x7e, 0x80, 0xba, 0xb4, 0x28, 0x40, 0xc5,
	0x8a, 0xa6, 0x7e, 0xcb, 0xf2, 0x06, 0x46, 0x93, 0x0b, 0x30, 0xda, 0xb0, 0xaf, 0x47, 0x34, 0xc5,
	0xef, 0x51, 0x8f, 0x49, 0xa0, 0x0a, 0xe2, 0xaa, 0x2f, 0x7e, 0xdb, 0x16, 0xb6, 0x37, 0x76, 0x4d,
	0x1b, 0x2f, 0x9b, 0x36, 0x3e, 0x5a, 0x36, 0x2d, 0xdc, 0x36, 0x9a, 0xd4, 0x25, 0xdf, 0x7f, 0x12,
	0x2f, 0x42, 0x0e, 0xa8, 0x58, 0x95, 0x63, 0x99, 0x27, 0xe7, 0x8e, 0xab, 0xd7, 0x73, 0xac, 0x49,
	0x9c, 0xa3, 0x03, 0xac, 0xe3, 0x33, 0x34, 0x80, 0x99, 0x14, 0x59, 0x16, 0x17, 0x8a, 0xaa, 0xb2,
	0xb0, 0xed, 0xeb, 0x86, 0x5b, 0x46, 0x93, 0x66, 0x22, 0xea, 0xbb, 0xf0, 0xd0, 0x46, 0xf8, 0x03,
	0xaa, 0xea, 0x4a, 0x60, 0xa6, 0x38, 0xcd, 0xfc, 0x8e, 0x2d, 0x64, 0xb7, 0xd1, 0xf3, 0xd7, 0xe7,
	0xe9, 0xd0, 0x37, 0x9a, 0xec, 0x5c, 0xd0, 0x1f, 0x89, 0x29, 0x57, 0x30, 0xcd, 0xd5, 0x3c, 0xaa,
	0x99, 0xe0, 0x18, 0x6d, 0x32, 0x91, 0x65, 0xc0, 0x14, 0x24, 0x71, 0x42, 0x15, 0xf5, 0xd7, 0xf7,
	0xdb, 0xf6, 0xfb, 0x1a, 0xb6, 0x4b, 0xca, 0x01, 0x55, 0x34, 0xbc, 0x63, 0x34, 0xf1, 0x9b, 0xaa,
	0x9a, 0xfb, 0x80, 0xd5, 0xc9, 0xa3, 0x04, 0xdd, 0xb8, 0x5c, 0x1a, 0xbe, 0x85, 0x5a, 0x3c, 0x59,
	0x0c, 0xbc, 0x63, 0x34, 0x69, 0xf1, 0x24, 0x6a, 0xf1, 0x04, 0x3f, 0x47, 0x9b, 0x79, 0x39, 0xc9,
	0x38, 0x8b, 0xbf, 0xc2, 0x3c, 0xae, 0x56, 0xb2, 0x1a, 0x76, 0x3f, 0xc4, 0x46, 0x93, 0x4b, 0x99,
	0xa8, 0xef, 0xe2, 0xb7, 0x30, 0x3f, 0x00, 0x39, 0xfa, 0xd1, 0x42, 0xdb, 0x57, 0x94, 0x8a, 0x0f,
	0x51, 0x7f, 0x51, 0x8e, 0x1b, 0x9e, 0xf7, 0xdf, 0xe1, 0xed, 0x18, 0x4d, 0x1a, 0x1a, 0x3b, 0xbd,
	0xde, 0x02, 0x59, 0x2e, 0x84, 0x04, 0x26, 0x64, 0xe2, 0x3c, 0x5b, 0xd7, 0x5b, 0x88, 0x9a, 0xc4,
	0x2d, 0x84, 0x03, 0xac, 0x63, 0xed, 0x37, 0x68, 0xff, 0xfb, 0x37, 0x78, 0x85, 0x06, 0x05, 0x48,
	0x4e, 0xb3, 0x78, 0x56, 0x4e, 0x27, 0x20, 0xed, 0x2a, 0x76, 0xc3, 0xdb, 0x46, 0x93, 0xdd, 0x46,
	0xa2, 0x36, 0x8d, 0xbe, 0x4b, 0xbc, 0xb3, 0x78, 0xf8, 0xf2, 0xe4, 0xf7, 0x70, 0xe5, 0xe4, 0x6c,
	0xe8, 0x9d, 0x9e, 0x0d, 0xbd, 0x5f, 0x67, 0x43, 0xef, 0xd3, 0xc3, 0x94, 0xab, 0x2f, 0xe5, 0x64,
	0xcc, 0xc4, 0x34, 0x48, 0x25, 0x3d, 0xe6, 0x8a, 0x2a, 0x2e, 0x66, 0x34, 0x0b, 0xce, 0x6f, 0x11,
	0x9a, 0x73, 0x77, 0x79, 0x4c, 0x3a, 0xf6, 0xdb, 0x9e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0x1a, 0x20, 0x63, 0xc4, 0x04, 0x00, 0x00,
}

func (m *DeviceV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceV1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceV1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Spec != nil {
		{
			size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDevice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.ResourceHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDevice(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DeviceSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CollectedData) > 0 {
		for iNdEx := len(m.CollectedData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollectedData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDevice(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Credential != nil {
		{
			size, err := m.Credential.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDevice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.EnrollStatus) > 0 {
		i -= len(m.EnrollStatus)
		copy(dAtA[i:], m.EnrollStatus)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.EnrollStatus)))
		i--
		dAtA[i] = 0x2a
	}
	if m.UpdateTime != nil {
		n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.UpdateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdateTime):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintDevice(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0x22
	}
	if m.CreateTime != nil {
		n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateTime):])
		if err5 != nil {
			return 0, err5
		}
		i -= n5
		i = encodeVarintDevice(dAtA, i, uint64(n5))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AssetTag) > 0 {
		i -= len(m.AssetTag)
		copy(dAtA[i:], m.AssetTag)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.AssetTag)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OsType) > 0 {
		i -= len(m.OsType)
		copy(dAtA[i:], m.OsType)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.OsType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeviceCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PublicKeyDer) > 0 {
		i -= len(m.PublicKeyDer)
		copy(dAtA[i:], m.PublicKeyDer)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.PublicKeyDer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeviceCollectedData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceCollectedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceCollectedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OsType) > 0 {
		i -= len(m.OsType)
		copy(dAtA[i:], m.OsType)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.OsType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RecordTime != nil {
		n6, err6 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.RecordTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.RecordTime):])
		if err6 != nil {
			return 0, err6
		}
		i -= n6
		i = encodeVarintDevice(dAtA, i, uint64(n6))
		i--
		dAtA[i] = 0x12
	}
	if m.CollectTime != nil {
		n7, err7 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CollectTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CollectTime):])
		if err7 != nil {
			return 0, err7
		}
		i -= n7
		i = encodeVarintDevice(dAtA, i, uint64(n7))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDevice(dAtA []byte, offset int, v uint64) int {
	offset -= sovDevice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ResourceHeader.Size()
	n += 1 + l + sovDevice(uint64(l))
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeviceSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OsType)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.AssetTag)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.CreateTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateTime)
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.UpdateTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdateTime)
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.EnrollStatus)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.Credential != nil {
		l = m.Credential.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if len(m.CollectedData) > 0 {
		for _, e := range m.CollectedData {
			l = e.Size()
			n += 1 + l + sovDevice(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeviceCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.PublicKeyDer)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeviceCollectedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollectTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CollectTime)
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.RecordTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.RecordTime)
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.OsType)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDevice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDevice(x uint64) (n int) {
	return sovDevice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ResourceHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &DeviceSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDevice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeviceSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OsType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreateTime == nil {
				m.CreateTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdateTime == nil {
				m.UpdateTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.UpdateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnrollStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnrollStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credential", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Credential == nil {
				m.Credential = &DeviceCredential{}
			}
			if err := m.Credential.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectedData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectedData = append(m.CollectedData, &DeviceCollectedData{})
			if err := m.CollectedData[len(m.CollectedData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDevice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeviceCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeyDer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeyDer = append(m.PublicKeyDer[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKeyDer == nil {
				m.PublicKeyDer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDevice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeviceCollectedData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceCollectedData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceCollectedData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CollectTime == nil {
				m.CollectTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CollectTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RecordTime == nil {
				m.RecordTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.RecordTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OsType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDevice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDevice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDevice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDevice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDevice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDevice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDevice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDevice = fmt.Errorf("proto: unexpected end of group")
)
