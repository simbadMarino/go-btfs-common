// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/shared/shared.proto

package shared

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RuntimeInfoReport_HealthStatus int32

const (
	RuntimeInfoReport_SICK         RuntimeInfoReport_HealthStatus = 0
	RuntimeInfoReport_RUNNING      RuntimeInfoReport_HealthStatus = 1
	RuntimeInfoReport_BOOTSTRAP    RuntimeInfoReport_HealthStatus = 2
	RuntimeInfoReport_PARTIAL_STOP RuntimeInfoReport_HealthStatus = 3
)

var RuntimeInfoReport_HealthStatus_name = map[int32]string{
	0: "SICK",
	1: "RUNNING",
	2: "BOOTSTRAP",
	3: "PARTIAL_STOP",
}

var RuntimeInfoReport_HealthStatus_value = map[string]int32{
	"SICK":         0,
	"RUNNING":      1,
	"BOOTSTRAP":    2,
	"PARTIAL_STOP": 3,
}

func (x RuntimeInfoReport_HealthStatus) String() string {
	return proto.EnumName(RuntimeInfoReport_HealthStatus_name, int32(x))
}

func (RuntimeInfoReport_HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_40efd7a2bcef91b6, []int{1, 0}
}

type RuntimeInfoRequest struct {
	RequestAddress       []byte           `protobuf:"bytes,1,opt,name=request_address,json=requestAddress,proto3" json:"request_address,omitempty"`
	CurentTime           *types.Timestamp `protobuf:"bytes,2,opt,name=curent_time,json=curentTime,proto3" json:"curent_time,omitempty"`
	Signature            []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RuntimeInfoRequest) Reset()         { *m = RuntimeInfoRequest{} }
func (m *RuntimeInfoRequest) String() string { return proto.CompactTextString(m) }
func (*RuntimeInfoRequest) ProtoMessage()    {}
func (*RuntimeInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40efd7a2bcef91b6, []int{0}
}
func (m *RuntimeInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuntimeInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuntimeInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeInfoRequest.Merge(m, src)
}
func (m *RuntimeInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeInfoRequest proto.InternalMessageInfo

func (m *RuntimeInfoRequest) GetRequestAddress() []byte {
	if m != nil {
		return m.RequestAddress
	}
	return nil
}

func (m *RuntimeInfoRequest) GetCurentTime() *types.Timestamp {
	if m != nil {
		return m.CurentTime
	}
	return nil
}

func (m *RuntimeInfoRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (*RuntimeInfoRequest) XXX_MessageName() string {
	return "shared.RuntimeInfoRequest"
}

type RuntimeInfoReport struct {
	PeerId               []byte                         `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Address              []byte                         `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ServiceName          []byte                         `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Status               RuntimeInfoReport_HealthStatus `protobuf:"varint,4,opt,name=status,proto3,enum=shared.RuntimeInfoReport_HealthStatus" json:"status,omitempty"`
	CurentTime           *time.Time                     `protobuf:"bytes,5,opt,name=curent_time,json=curentTime,proto3,stdtime" json:"curent_time,omitempty"`
	StartTime            *time.Time                     `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time,omitempty"`
	GitHash              []byte                         `protobuf:"bytes,7,opt,name=git_hash,json=gitHash,proto3" json:"git_hash,omitempty"`
	Version              []byte                         `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`
	DbStatusExtra        []byte                         `protobuf:"bytes,9,opt,name=db_status_extra,json=dbStatusExtra,proto3" json:"db_status_extra,omitempty"`
	QueueStatusExtra     []byte                         `protobuf:"bytes,10,opt,name=queue_status_extra,json=queueStatusExtra,proto3" json:"queue_status_extra,omitempty"`
	ChainStatusExtra     []byte                         `protobuf:"bytes,11,opt,name=chain_status_extra,json=chainStatusExtra,proto3" json:"chain_status_extra,omitempty"`
	CacheStatusExtra     []byte                         `protobuf:"bytes,12,opt,name=cache_status_extra,json=cacheStatusExtra,proto3" json:"cache_status_extra,omitempty"`
	Extra                []byte                         `protobuf:"bytes,13,opt,name=extra,proto3" json:"extra,omitempty"`
	Signature            []byte                         `protobuf:"bytes,14,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RuntimeInfoReport) Reset()         { *m = RuntimeInfoReport{} }
func (m *RuntimeInfoReport) String() string { return proto.CompactTextString(m) }
func (*RuntimeInfoReport) ProtoMessage()    {}
func (*RuntimeInfoReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_40efd7a2bcef91b6, []int{1}
}
func (m *RuntimeInfoReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeInfoReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuntimeInfoReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuntimeInfoReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeInfoReport.Merge(m, src)
}
func (m *RuntimeInfoReport) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeInfoReport) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeInfoReport.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeInfoReport proto.InternalMessageInfo

func (m *RuntimeInfoReport) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

func (m *RuntimeInfoReport) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *RuntimeInfoReport) GetServiceName() []byte {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *RuntimeInfoReport) GetStatus() RuntimeInfoReport_HealthStatus {
	if m != nil {
		return m.Status
	}
	return RuntimeInfoReport_SICK
}

func (m *RuntimeInfoReport) GetCurentTime() *time.Time {
	if m != nil {
		return m.CurentTime
	}
	return nil
}

func (m *RuntimeInfoReport) GetStartTime() *time.Time {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *RuntimeInfoReport) GetGitHash() []byte {
	if m != nil {
		return m.GitHash
	}
	return nil
}

func (m *RuntimeInfoReport) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *RuntimeInfoReport) GetDbStatusExtra() []byte {
	if m != nil {
		return m.DbStatusExtra
	}
	return nil
}

func (m *RuntimeInfoReport) GetQueueStatusExtra() []byte {
	if m != nil {
		return m.QueueStatusExtra
	}
	return nil
}

func (m *RuntimeInfoReport) GetChainStatusExtra() []byte {
	if m != nil {
		return m.ChainStatusExtra
	}
	return nil
}

func (m *RuntimeInfoReport) GetCacheStatusExtra() []byte {
	if m != nil {
		return m.CacheStatusExtra
	}
	return nil
}

func (m *RuntimeInfoReport) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *RuntimeInfoReport) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (*RuntimeInfoReport) XXX_MessageName() string {
	return "shared.RuntimeInfoReport"
}
func init() {
	proto.RegisterEnum("shared.RuntimeInfoReport_HealthStatus", RuntimeInfoReport_HealthStatus_name, RuntimeInfoReport_HealthStatus_value)
	golang_proto.RegisterEnum("shared.RuntimeInfoReport_HealthStatus", RuntimeInfoReport_HealthStatus_name, RuntimeInfoReport_HealthStatus_value)
	proto.RegisterType((*RuntimeInfoRequest)(nil), "shared.RuntimeInfoRequest")
	golang_proto.RegisterType((*RuntimeInfoRequest)(nil), "shared.RuntimeInfoRequest")
	proto.RegisterType((*RuntimeInfoReport)(nil), "shared.RuntimeInfoReport")
	golang_proto.RegisterType((*RuntimeInfoReport)(nil), "shared.RuntimeInfoReport")
}

func init() { proto.RegisterFile("protos/shared/shared.proto", fileDescriptor_40efd7a2bcef91b6) }
func init() { golang_proto.RegisterFile("protos/shared/shared.proto", fileDescriptor_40efd7a2bcef91b6) }

var fileDescriptor_40efd7a2bcef91b6 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xf1, 0x7e, 0xf4, 0xc7, 0x4b, 0xdb, 0x05, 0x0b, 0x89, 0x50, 0xa1, 0xac, 0xec, 0x30,
	0x7a, 0x80, 0x54, 0x1a, 0x47, 0x24, 0xa6, 0x16, 0x01, 0xab, 0x40, 0x6d, 0x95, 0x96, 0x0b, 0x97,
	0xc8, 0x69, 0xdc, 0x24, 0xd2, 0x1a, 0x77, 0xb6, 0x33, 0xf1, 0x67, 0xec, 0xb8, 0x3f, 0x87, 0xe3,
	0x8e, 0xfc, 0x07, 0xa0, 0xf6, 0x1f, 0x41, 0xb1, 0x5d, 0x6d, 0x29, 0x07, 0x38, 0xc5, 0xdf, 0xef,
	0xfb, 0xf8, 0xbd, 0xe7, 0x67, 0x07, 0xda, 0x2b, 0xce, 0x24, 0x13, 0x3d, 0x91, 0x10, 0x4e, 0x23,
	0xf3, 0xf1, 0x94, 0x89, 0x2b, 0x5a, 0xb5, 0x5f, 0xc7, 0xa9, 0x4c, 0xf2, 0xd0, 0x9b, 0xb3, 0x65,
	0x2f, 0x66, 0x31, 0xeb, 0xa9, 0x70, 0x98, 0x2f, 0x94, 0x52, 0x42, 0xad, 0xf4, 0xb6, 0xf6, 0x71,
	0xcc, 0x58, 0x7c, 0x49, 0xef, 0x29, 0x99, 0x2e, 0xa9, 0x90, 0x64, 0xb9, 0xd2, 0xc0, 0xc9, 0x2d,
	0x02, 0xec, 0xe7, 0x59, 0x61, 0x0f, 0xb3, 0x05, 0xf3, 0xe9, 0x55, 0x4e, 0x85, 0xc4, 0x2f, 0xe1,
	0x88, 0xeb, 0x65, 0x40, 0xa2, 0x88, 0x53, 0x21, 0x1c, 0xd4, 0x41, 0xdd, 0x86, 0xdf, 0x32, 0x76,
	0x5f, 0xbb, 0xf8, 0x2d, 0x58, 0xf3, 0x9c, 0xd3, 0x4c, 0x06, 0x45, 0x0a, 0x67, 0xaf, 0x83, 0xba,
	0xd6, 0x59, 0xdb, 0xd3, 0x65, 0xbd, 0x6d, 0x59, 0x6f, 0xb6, 0x2d, 0xeb, 0x83, 0xc6, 0x0b, 0x03,
	0x3f, 0x87, 0xba, 0x48, 0xe3, 0x8c, 0xc8, 0x9c, 0x53, 0x67, 0x5f, 0xe5, 0xbf, 0x37, 0x4e, 0x6e,
	0x0e, 0xe1, 0x71, 0xa9, 0xb5, 0x15, 0xe3, 0x12, 0x3f, 0x85, 0xea, 0x8a, 0x52, 0x1e, 0xa4, 0x91,
	0xe9, 0xa8, 0x52, 0xc8, 0x61, 0x84, 0x1d, 0xa8, 0x6e, 0x5b, 0xdd, 0x53, 0x81, 0xad, 0xc4, 0x2f,
	0xa0, 0x21, 0x28, 0xbf, 0x4e, 0xe7, 0x34, 0xc8, 0xc8, 0x72, 0x5b, 0xc9, 0x32, 0xde, 0x88, 0x2c,
	0x29, 0x7e, 0x07, 0x15, 0x21, 0x89, 0xcc, 0x85, 0x73, 0xd0, 0x41, 0xdd, 0xd6, 0xd9, 0xa9, 0x67,
	0xa6, 0xff, 0x57, 0x03, 0xde, 0x05, 0x25, 0x97, 0x32, 0x99, 0x2a, 0xda, 0x37, 0xbb, 0x70, 0xbf,
	0x3c, 0x86, 0xc3, 0x7f, 0x8d, 0x61, 0x70, 0x70, 0xf3, 0xeb, 0x18, 0x95, 0x86, 0x71, 0x0e, 0x20,
	0x24, 0xe1, 0x26, 0x43, 0xe5, 0x3f, 0x33, 0xd4, 0xd5, 0x1e, 0x95, 0xe0, 0x19, 0xd4, 0xe2, 0x54,
	0x06, 0x09, 0x11, 0x89, 0x53, 0xd5, 0x13, 0x88, 0x53, 0x79, 0x41, 0x44, 0x52, 0xcc, 0xe6, 0x9a,
	0x72, 0x91, 0xb2, 0xcc, 0xa9, 0xe9, 0x88, 0x91, 0xf8, 0x14, 0x8e, 0xa2, 0x30, 0xd0, 0xa7, 0x08,
	0xe8, 0x77, 0xc9, 0x89, 0x53, 0x57, 0x44, 0x33, 0x0a, 0xf5, 0x19, 0x3f, 0x14, 0x26, 0x7e, 0x05,
	0xf8, 0x2a, 0xa7, 0x39, 0x2d, 0xa3, 0xa0, 0x50, 0x5b, 0x45, 0x76, 0xe8, 0x79, 0x42, 0xd2, 0xac,
	0x4c, 0x5b, 0x9a, 0x56, 0x91, 0x5d, 0x9a, 0xcc, 0x93, 0x9d, 0xdc, 0x0d, 0x43, 0x17, 0x91, 0x87,
	0xf4, 0x13, 0x38, 0xd4, 0x40, 0x53, 0x01, 0x5a, 0x94, 0x9f, 0x52, 0x6b, 0xf7, 0x29, 0x7d, 0x84,
	0xc6, 0xc3, 0x6b, 0xc3, 0x35, 0x38, 0x98, 0x0e, 0xdf, 0x7f, 0xb6, 0x1f, 0x61, 0x0b, 0xaa, 0xfe,
	0xd7, 0xd1, 0x68, 0x38, 0xfa, 0x64, 0x23, 0xdc, 0x84, 0xfa, 0x60, 0x3c, 0x9e, 0x4d, 0x67, 0x7e,
	0x7f, 0x62, 0xef, 0x61, 0x1b, 0x1a, 0x93, 0xbe, 0x3f, 0x1b, 0xf6, 0xbf, 0x04, 0xd3, 0xd9, 0x78,
	0x62, 0xef, 0x0f, 0xce, 0xef, 0xd6, 0x2e, 0xfa, 0xb9, 0x76, 0xd1, 0xef, 0xb5, 0x8b, 0x6e, 0x37,
	0x2e, 0xfa, 0xb1, 0x71, 0xd1, 0xdd, 0xc6, 0x45, 0xd0, 0x4a, 0x99, 0x17, 0xca, 0x85, 0x30, 0xef,
	0x66, 0x60, 0x4d, 0xd5, 0x77, 0x52, 0xdc, 0xdb, 0x04, 0x7d, 0x33, 0xbf, 0x6f, 0x58, 0x51, 0x17,
	0xf9, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xb9, 0x38, 0xc3, 0xeb, 0x03, 0x00, 0x00,
}

func (m *RuntimeInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuntimeInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintShared(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CurentTime != nil {
		{
			size, err := m.CurentTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintShared(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RequestAddress) > 0 {
		i -= len(m.RequestAddress)
		copy(dAtA[i:], m.RequestAddress)
		i = encodeVarintShared(dAtA, i, uint64(len(m.RequestAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RuntimeInfoReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeInfoReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuntimeInfoReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintShared(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarintShared(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.CacheStatusExtra) > 0 {
		i -= len(m.CacheStatusExtra)
		copy(dAtA[i:], m.CacheStatusExtra)
		i = encodeVarintShared(dAtA, i, uint64(len(m.CacheStatusExtra)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ChainStatusExtra) > 0 {
		i -= len(m.ChainStatusExtra)
		copy(dAtA[i:], m.ChainStatusExtra)
		i = encodeVarintShared(dAtA, i, uint64(len(m.ChainStatusExtra)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.QueueStatusExtra) > 0 {
		i -= len(m.QueueStatusExtra)
		copy(dAtA[i:], m.QueueStatusExtra)
		i = encodeVarintShared(dAtA, i, uint64(len(m.QueueStatusExtra)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.DbStatusExtra) > 0 {
		i -= len(m.DbStatusExtra)
		copy(dAtA[i:], m.DbStatusExtra)
		i = encodeVarintShared(dAtA, i, uint64(len(m.DbStatusExtra)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintShared(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.GitHash) > 0 {
		i -= len(m.GitHash)
		copy(dAtA[i:], m.GitHash)
		i = encodeVarintShared(dAtA, i, uint64(len(m.GitHash)))
		i--
		dAtA[i] = 0x3a
	}
	if m.StartTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintShared(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x32
	}
	if m.CurentTime != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CurentTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CurentTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintShared(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintShared(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintShared(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PeerId) > 0 {
		i -= len(m.PeerId)
		copy(dAtA[i:], m.PeerId)
		i = encodeVarintShared(dAtA, i, uint64(len(m.PeerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintShared(dAtA []byte, offset int, v uint64) int {
	offset -= sovShared(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RuntimeInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestAddress)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	if m.CurentTime != nil {
		l = m.CurentTime.Size()
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RuntimeInfoReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovShared(uint64(m.Status))
	}
	if m.CurentTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CurentTime)
		n += 1 + l + sovShared(uint64(l))
	}
	if m.StartTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartTime)
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.GitHash)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.DbStatusExtra)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.QueueStatusExtra)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.ChainStatusExtra)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.CacheStatusExtra)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.Extra)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovShared(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShared(x uint64) (n int) {
	return sovShared(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RuntimeInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShared
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
			return fmt.Errorf("proto: RuntimeInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestAddress = append(m.RequestAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestAddress == nil {
				m.RequestAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurentTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurentTime == nil {
				m.CurentTime = &types.Timestamp{}
			}
			if err := m.CurentTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShared(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShared
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShared
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
func (m *RuntimeInfoReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShared
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
			return fmt.Errorf("proto: RuntimeInfoReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeInfoReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = append(m.PeerId[:0], dAtA[iNdEx:postIndex]...)
			if m.PeerId == nil {
				m.PeerId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = append(m.ServiceName[:0], dAtA[iNdEx:postIndex]...)
			if m.ServiceName == nil {
				m.ServiceName = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RuntimeInfoReport_HealthStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurentTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurentTime == nil {
				m.CurentTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CurentTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitHash = append(m.GitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.GitHash == nil {
				m.GitHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = append(m.Version[:0], dAtA[iNdEx:postIndex]...)
			if m.Version == nil {
				m.Version = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbStatusExtra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DbStatusExtra = append(m.DbStatusExtra[:0], dAtA[iNdEx:postIndex]...)
			if m.DbStatusExtra == nil {
				m.DbStatusExtra = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueStatusExtra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueueStatusExtra = append(m.QueueStatusExtra[:0], dAtA[iNdEx:postIndex]...)
			if m.QueueStatusExtra == nil {
				m.QueueStatusExtra = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainStatusExtra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainStatusExtra = append(m.ChainStatusExtra[:0], dAtA[iNdEx:postIndex]...)
			if m.ChainStatusExtra == nil {
				m.ChainStatusExtra = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheStatusExtra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CacheStatusExtra = append(m.CacheStatusExtra[:0], dAtA[iNdEx:postIndex]...)
			if m.CacheStatusExtra == nil {
				m.CacheStatusExtra = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShared(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShared
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShared
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
func skipShared(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShared
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
					return 0, ErrIntOverflowShared
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
					return 0, ErrIntOverflowShared
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
				return 0, ErrInvalidLengthShared
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShared
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShared
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShared        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShared          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShared = fmt.Errorf("proto: unexpected end of group")
)
