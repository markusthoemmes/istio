// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/redisquota/config/config.proto

// The `redisquota` adapter can be used to support Istio's quota management
// system. It depends on a Redis server to store quota values.
//
// This adapter supports the [quota template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/quota/).

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Algorithms for rate-limiting:
type Params_QuotaAlgorithm int32

const (
	// FIXED_WINDOW The fixed window approach can allow 2x peak specified rate, whereas the rolling-window doesn't.
	FIXED_WINDOW Params_QuotaAlgorithm = 0
	// ROLLING_WINDOW The rolling window algorithm's additional precision comes at the cost of increased redis resource usage.
	ROLLING_WINDOW Params_QuotaAlgorithm = 1
)

var Params_QuotaAlgorithm_name = map[int32]string{
	0: "FIXED_WINDOW",
	1: "ROLLING_WINDOW",
}

var Params_QuotaAlgorithm_value = map[string]int32{
	"FIXED_WINDOW":   0,
	"ROLLING_WINDOW": 1,
}

func (Params_QuotaAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4ec77e3e2f5a044, []int{0, 0}
}

// redisquota adapter supports the rate limit quota using either fixed or
// rolling window algorithm. And it is using Redis as a shared data storage.
//
// Example configuration:
//
// ```yaml
// redisServerUrl: localhost:6379
// connectionPoolSize: 10
// quotas:
//   - name: requestcount.quota.istio-system
//     maxAmount: 50
//     validDuration: 60s
//     bucketDuration: 1s
//     rateLimitAlgorithm: ROLLING_WINDOW
//     overrides:
//       - dimensions:
//           destination: ratings
//           source: reviews
//         maxAmount: 12
//       - dimensions:
//           destination: reviews
//         maxAmount: 5
// ```
type Params struct {
	// The set of known quotas. At least one quota configuration is required
	Quotas []Params_Quota `protobuf:"bytes,1,rep,name=quotas,proto3" json:"quotas"`
	// Redis connection string <hostname>:<port number>
	// ex) localhost:6379
	RedisServerUrl string `protobuf:"bytes,2,opt,name=redis_server_url,json=redisServerUrl,proto3" json:"redis_server_url,omitempty"`
	// Maximum number of idle connections to redis
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	ConnectionPoolSize int64 `protobuf:"varint,3,opt,name=connection_pool_size,json=connectionPoolSize,proto3" json:"connection_pool_size,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ec77e3e2f5a044, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type Params_Override struct {
	// The specific dimensions for which this override applies.
	// String representation of instance dimensions is used to check against configured dimensions.
	// dimensions should not be empty
	Dimensions map[string]string `protobuf:"bytes,1,rep,name=dimensions,proto3" json:"dimensions" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The upper limit for this quota override.
	// This value should be bigger than 0
	MaxAmount int64 `protobuf:"varint,2,opt,name=max_amount,json=maxAmount,proto3" json:"max_amount,omitempty"`
}

func (m *Params_Override) Reset()      { *m = Params_Override{} }
func (*Params_Override) ProtoMessage() {}
func (*Params_Override) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ec77e3e2f5a044, []int{0, 0}
}
func (m *Params_Override) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params_Override) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params_Override.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params_Override) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params_Override.Merge(m, src)
}
func (m *Params_Override) XXX_Size() int {
	return m.Size()
}
func (m *Params_Override) XXX_DiscardUnknown() {
	xxx_messageInfo_Params_Override.DiscardUnknown(m)
}

var xxx_messageInfo_Params_Override proto.InternalMessageInfo

func (m *Params_Override) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Params_Override) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

type Params_Quota struct {
	// The name of the quota
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The upper limit for this quota. max_amount should be bigger than 0
	MaxAmount int64 `protobuf:"varint,2,opt,name=max_amount,json=maxAmount,proto3" json:"max_amount,omitempty"`
	// The amount of time allocated quota remains valid before it is
	// automatically released. This is only meaningful for rate limit quotas.
	// value should be 0 < valid_duration
	ValidDuration time.Duration `protobuf:"bytes,3,opt,name=valid_duration,json=validDuration,proto3,stdduration" json:"valid_duration"`
	// bucket_duration will be ignored if rate_limit_algorithm is FIXED_WINDOW
	// value should be 0 < bucket_duration < valid_duration
	BucketDuration time.Duration `protobuf:"bytes,4,opt,name=bucket_duration,json=bucketDuration,proto3,stdduration" json:"bucket_duration"`
	// Quota management algorithm. The default value is FIXED_WINDOW
	RateLimitAlgorithm Params_QuotaAlgorithm `protobuf:"varint,5,opt,name=rate_limit_algorithm,json=rateLimitAlgorithm,proto3,enum=adapter.redisquota.config.Params_QuotaAlgorithm" json:"rate_limit_algorithm,omitempty"`
	// Overrides associated with this quota.
	// The first matching override is applied.
	Overrides []*Params_Override `protobuf:"bytes,6,rep,name=overrides,proto3" json:"overrides,omitempty"`
}

func (m *Params_Quota) Reset()      { *m = Params_Quota{} }
func (*Params_Quota) ProtoMessage() {}
func (*Params_Quota) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ec77e3e2f5a044, []int{0, 1}
}
func (m *Params_Quota) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params_Quota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params_Quota.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params_Quota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params_Quota.Merge(m, src)
}
func (m *Params_Quota) XXX_Size() int {
	return m.Size()
}
func (m *Params_Quota) XXX_DiscardUnknown() {
	xxx_messageInfo_Params_Quota.DiscardUnknown(m)
}

var xxx_messageInfo_Params_Quota proto.InternalMessageInfo

func (m *Params_Quota) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Params_Quota) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *Params_Quota) GetValidDuration() time.Duration {
	if m != nil {
		return m.ValidDuration
	}
	return 0
}

func (m *Params_Quota) GetBucketDuration() time.Duration {
	if m != nil {
		return m.BucketDuration
	}
	return 0
}

func (m *Params_Quota) GetRateLimitAlgorithm() Params_QuotaAlgorithm {
	if m != nil {
		return m.RateLimitAlgorithm
	}
	return FIXED_WINDOW
}

func (m *Params_Quota) GetOverrides() []*Params_Override {
	if m != nil {
		return m.Overrides
	}
	return nil
}

func init() {
	proto.RegisterEnum("adapter.redisquota.config.Params_QuotaAlgorithm", Params_QuotaAlgorithm_name, Params_QuotaAlgorithm_value)
	proto.RegisterType((*Params)(nil), "adapter.redisquota.config.Params")
	proto.RegisterType((*Params_Override)(nil), "adapter.redisquota.config.Params.Override")
	proto.RegisterMapType((map[string]string)(nil), "adapter.redisquota.config.Params.Override.DimensionsEntry")
	proto.RegisterType((*Params_Quota)(nil), "adapter.redisquota.config.Params.Quota")
}

func init() {
	proto.RegisterFile("mixer/adapter/redisquota/config/config.proto", fileDescriptor_b4ec77e3e2f5a044)
}

var fileDescriptor_b4ec77e3e2f5a044 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x6f, 0xd3, 0x4e,
	0x18, 0xf5, 0xd5, 0xa9, 0xd5, 0x5e, 0x7f, 0x3f, 0x37, 0x3a, 0x65, 0x70, 0x23, 0x71, 0x8d, 0xba,
	0x10, 0x21, 0x64, 0x57, 0x45, 0x42, 0xa8, 0x12, 0x12, 0xad, 0x52, 0xa0, 0x28, 0x6a, 0x8a, 0x2b,
	0x54, 0xc4, 0x62, 0x2e, 0xf1, 0xd5, 0x9c, 0x6a, 0xfb, 0xc2, 0xd9, 0x8e, 0xd2, 0x4e, 0x8c, 0x1d,
	0x19, 0x19, 0x19, 0xf9, 0x37, 0xd8, 0x32, 0x66, 0xec, 0x04, 0xc4, 0x59, 0x18, 0xfb, 0x27, 0x20,
	0x9f, 0xed, 0x04, 0x90, 0x50, 0x33, 0xf9, 0xf3, 0xbb, 0xf7, 0xde, 0x7d, 0xdf, 0xf7, 0x0e, 0xde,
	0x0f, 0xd8, 0x90, 0x0a, 0x8b, 0xb8, 0xa4, 0x1f, 0x53, 0x61, 0x09, 0xea, 0xb2, 0xe8, 0x7d, 0xc2,
	0x63, 0x62, 0xf5, 0x78, 0x78, 0xc6, 0xbc, 0xe2, 0x63, 0xf6, 0x05, 0x8f, 0x39, 0xda, 0x28, 0x78,
	0xe6, 0x9c, 0x67, 0xe6, 0x84, 0x3a, 0xf6, 0x38, 0xf7, 0x7c, 0x6a, 0x49, 0x62, 0x37, 0x39, 0xb3,
	0xdc, 0x44, 0x90, 0x98, 0xf1, 0x30, 0x97, 0xd6, 0x6b, 0x1e, 0xf7, 0xb8, 0x2c, 0xad, 0xac, 0xca,
	0xd1, 0xad, 0xaf, 0x1a, 0xd4, 0x8e, 0x89, 0x20, 0x41, 0x84, 0x0e, 0xa0, 0x26, 0x0d, 0x23, 0x03,
	0x34, 0xd4, 0xe6, 0xda, 0xce, 0x5d, 0xf3, 0x9f, 0x97, 0x99, 0xb9, 0xc4, 0x7c, 0x99, 0x61, 0xfb,
	0x95, 0xd1, 0xb7, 0x4d, 0xc5, 0x2e, 0xc4, 0xa8, 0x09, 0xab, 0x92, 0xef, 0x44, 0x54, 0x0c, 0xa8,
	0x70, 0x12, 0xe1, 0x1b, 0x4b, 0x0d, 0xd0, 0x5c, 0xb5, 0x75, 0x89, 0x9f, 0x48, 0xf8, 0x95, 0xf0,
	0xd1, 0x36, 0xac, 0xf5, 0x78, 0x18, 0xd2, 0x5e, 0xd6, 0xa5, 0xd3, 0xe7, 0xdc, 0x77, 0x22, 0x76,
	0x49, 0x0d, 0xb5, 0x01, 0x9a, 0xaa, 0x8d, 0xe6, 0x67, 0xc7, 0x9c, 0xfb, 0x27, 0xec, 0x92, 0xd6,
	0xc7, 0x00, 0xae, 0x74, 0x06, 0x54, 0x08, 0xe6, 0x52, 0xf4, 0x16, 0x42, 0x97, 0x05, 0x34, 0x8c,
	0x18, 0x0f, 0xcb, 0x9e, 0x77, 0x6f, 0xef, 0xb9, 0xd4, 0x9b, 0xad, 0x99, 0xf8, 0x20, 0x8c, 0xc5,
	0x45, 0x31, 0xc6, 0x6f, 0x9e, 0xe8, 0x0e, 0x84, 0x01, 0x19, 0x3a, 0x24, 0xe0, 0x49, 0x18, 0xcb,
	0x21, 0x54, 0x7b, 0x35, 0x20, 0xc3, 0x3d, 0x09, 0xd4, 0x1f, 0xc3, 0xf5, 0xbf, 0x3c, 0x50, 0x15,
	0xaa, 0xe7, 0xf4, 0xc2, 0x00, 0x72, 0xde, 0xac, 0x44, 0x35, 0xb8, 0x3c, 0x20, 0x7e, 0x42, 0x8b,
	0x1d, 0xe4, 0x3f, 0xbb, 0x4b, 0x8f, 0xc0, 0x6e, 0xe5, 0xea, 0xf3, 0x26, 0xa8, 0x5f, 0xa9, 0x70,
	0x59, 0xae, 0x11, 0x21, 0x58, 0x09, 0x49, 0x40, 0x0b, 0xb1, 0xac, 0x6f, 0xe9, 0x00, 0xbd, 0x80,
	0xfa, 0x80, 0xf8, 0xcc, 0x75, 0xca, 0xac, 0xe5, 0xee, 0xd6, 0x76, 0x36, 0xcc, 0xfc, 0x31, 0x98,
	0xe5, 0x63, 0x30, 0x5b, 0x05, 0x61, 0x7f, 0x25, 0x9b, 0xf2, 0xd3, 0xf7, 0x4d, 0x60, 0xff, 0x2f,
	0xa5, 0xe5, 0x01, 0x6a, 0xc3, 0xf5, 0x6e, 0xd2, 0x3b, 0xa7, 0xf1, 0xdc, 0xac, 0xb2, 0xb8, 0x99,
	0x9e, 0x6b, 0x67, 0x6e, 0x5d, 0x58, 0x13, 0x24, 0xa6, 0x8e, 0xcf, 0x02, 0x16, 0x3b, 0xc4, 0xf7,
	0xb8, 0x60, 0xf1, 0xbb, 0xc0, 0x58, 0x6e, 0x80, 0xa6, 0xbe, 0xb3, 0xbd, 0xe0, 0xd3, 0xda, 0x2b,
	0x75, 0x36, 0xca, 0xdc, 0xda, 0x99, 0xd9, 0x0c, 0x43, 0xcf, 0xe1, 0x2a, 0x2f, 0xc2, 0x8c, 0x0c,
	0x4d, 0xe6, 0x7f, 0x6f, 0xf1, 0xfc, 0xed, 0xb9, 0x38, 0x8f, 0x62, 0xeb, 0x21, 0xd4, 0xff, 0xbc,
	0x15, 0x55, 0xe1, 0x7f, 0x4f, 0x0f, 0x5f, 0x1f, 0xb4, 0x9c, 0xd3, 0xc3, 0xa3, 0x56, 0xe7, 0xb4,
	0xaa, 0x20, 0x04, 0x75, 0xbb, 0xd3, 0x6e, 0x1f, 0x1e, 0x3d, 0x2b, 0x31, 0xb0, 0xff, 0x64, 0x34,
	0xc1, 0xca, 0x78, 0x82, 0x95, 0xeb, 0x09, 0x56, 0x6e, 0x26, 0x58, 0xf9, 0x90, 0x62, 0xf0, 0x25,
	0xc5, 0xca, 0x28, 0xc5, 0x60, 0x9c, 0x62, 0xf0, 0x23, 0xc5, 0xe0, 0x67, 0x8a, 0x95, 0x9b, 0x14,
	0x83, 0x8f, 0x53, 0xac, 0x8c, 0xa7, 0x58, 0xb9, 0x9e, 0x62, 0xe5, 0x8d, 0x96, 0xb7, 0xd6, 0xd5,
	0xe4, 0x6a, 0x1f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xea, 0x07, 0x8f, 0x2c, 0x0d, 0x04, 0x00,
	0x00,
}

func (x Params_QuotaAlgorithm) String() string {
	s, ok := Params_QuotaAlgorithm_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for _, msg := range m.Quotas {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.RedisServerUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RedisServerUrl)))
		i += copy(dAtA[i:], m.RedisServerUrl)
	}
	if m.ConnectionPoolSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ConnectionPoolSize))
	}
	return i, nil
}

func (m *Params_Override) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_Override) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0xa
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.MaxAmount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxAmount))
	}
	return i, nil
}

func (m *Params_Quota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_Quota) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.MaxAmount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxAmount))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)))
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ValidDuration, dAtA[i:])
	if err1 != nil {
		return 0, err1
	}
	i += n1
	dAtA[i] = 0x22
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.BucketDuration)))
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.BucketDuration, dAtA[i:])
	if err2 != nil {
		return 0, err2
	}
	i += n2
	if m.RateLimitAlgorithm != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.RateLimitAlgorithm))
	}
	if len(m.Overrides) > 0 {
		for _, msg := range m.Overrides {
			dAtA[i] = 0x32
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for _, e := range m.Quotas {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.RedisServerUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.ConnectionPoolSize != 0 {
		n += 1 + sovConfig(uint64(m.ConnectionPoolSize))
	}
	return n
}

func (m *Params_Override) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	if m.MaxAmount != 0 {
		n += 1 + sovConfig(uint64(m.MaxAmount))
	}
	return n
}

func (m *Params_Quota) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.MaxAmount != 0 {
		n += 1 + sovConfig(uint64(m.MaxAmount))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovConfig(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.BucketDuration)
	n += 1 + l + sovConfig(uint64(l))
	if m.RateLimitAlgorithm != 0 {
		n += 1 + sovConfig(uint64(m.RateLimitAlgorithm))
	}
	if len(m.Overrides) > 0 {
		for _, e := range m.Overrides {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForQuotas := "[]Params_Quota{"
	for _, f := range this.Quotas {
		repeatedStringForQuotas += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForQuotas += "}"
	s := strings.Join([]string{`&Params{`,
		`Quotas:` + repeatedStringForQuotas + `,`,
		`RedisServerUrl:` + fmt.Sprintf("%v", this.RedisServerUrl) + `,`,
		`ConnectionPoolSize:` + fmt.Sprintf("%v", this.ConnectionPoolSize) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Override) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	s := strings.Join([]string{`&Params_Override{`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`MaxAmount:` + fmt.Sprintf("%v", this.MaxAmount) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Quota) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForOverrides := "[]*Params_Override{"
	for _, f := range this.Overrides {
		repeatedStringForOverrides += strings.Replace(fmt.Sprintf("%v", f), "Params_Override", "Params_Override", 1) + ","
	}
	repeatedStringForOverrides += "}"
	s := strings.Join([]string{`&Params_Quota{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`MaxAmount:` + fmt.Sprintf("%v", this.MaxAmount) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ValidDuration), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`BucketDuration:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.BucketDuration), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`RateLimitAlgorithm:` + fmt.Sprintf("%v", this.RateLimitAlgorithm) + `,`,
		`Overrides:` + repeatedStringForOverrides + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quotas = append(m.Quotas, Params_Quota{})
			if err := m.Quotas[len(m.Quotas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedisServerUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedisServerUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionPoolSize", wireType)
			}
			m.ConnectionPoolSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectionPoolSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_Override) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Override: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Override: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthConfig
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthConfig
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Dimensions[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAmount", wireType)
			}
			m.MaxAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAmount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_Quota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Quota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAmount", wireType)
			}
			m.MaxAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAmount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ValidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.BucketDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitAlgorithm", wireType)
			}
			m.RateLimitAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RateLimitAlgorithm |= Params_QuotaAlgorithm(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overrides", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Overrides = append(m.Overrides, &Params_Override{})
			if err := m.Overrides[len(m.Overrides)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)
