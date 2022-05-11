// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vpool/v1/vpool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Direction int32

const (
	Direction_ADD_TO_POOL      Direction = 0
	Direction_REMOVE_FROM_POOL Direction = 1
)

var Direction_name = map[int32]string{
	0: "ADD_TO_POOL",
	1: "REMOVE_FROM_POOL",
}

var Direction_value = map[string]int32{
	"ADD_TO_POOL":      0,
	"REMOVE_FROM_POOL": 1,
}

func (x Direction) String() string {
	return proto.EnumName(Direction_name, int32(x))
}

func (Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb4ac691d1b54d04, []int{0}
}

// Enumerates different options of calculating twap.
type TwapCalcOption int32

const (
	// Spot price from quote asset reserve / base asset reserve
	TwapCalcOption_SPOT TwapCalcOption = 0
	// Swapping with quote assets, output denominated in base assets
	TwapCalcOption_QUOTE_ASSET_SWAP TwapCalcOption = 1
	// Swapping with base assets, output denominated in quote assets
	TwapCalcOption_BASE_ASSET_SWAP TwapCalcOption = 2
)

var TwapCalcOption_name = map[int32]string{
	0: "SPOT",
	1: "QUOTE_ASSET_SWAP",
	2: "BASE_ASSET_SWAP",
}

var TwapCalcOption_value = map[string]int32{
	"SPOT":             0,
	"QUOTE_ASSET_SWAP": 1,
	"BASE_ASSET_SWAP":  2,
}

func (x TwapCalcOption) String() string {
	return proto.EnumName(TwapCalcOption_name, int32(x))
}

func (TwapCalcOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb4ac691d1b54d04, []int{1}
}

// a snapshot of the vpool's reserves at a given point in time
type ReserveSnapshot struct {
	BaseAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=base_asset_reserve,json=baseAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_asset_reserve"`
	// quote asset is usually the margin asset, e.g. NUSD
	QuoteAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=quote_asset_reserve,json=quoteAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quote_asset_reserve"`
	// seconds since unix epoch
	Timestamp   int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BlockNumber int64 `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (m *ReserveSnapshot) Reset()         { *m = ReserveSnapshot{} }
func (m *ReserveSnapshot) String() string { return proto.CompactTextString(m) }
func (*ReserveSnapshot) ProtoMessage()    {}
func (*ReserveSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb4ac691d1b54d04, []int{0}
}
func (m *ReserveSnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReserveSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReserveSnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReserveSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveSnapshot.Merge(m, src)
}
func (m *ReserveSnapshot) XXX_Size() int {
	return m.Size()
}
func (m *ReserveSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveSnapshot proto.InternalMessageInfo

func (m *ReserveSnapshot) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ReserveSnapshot) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

// A virtual pool used only for price discovery of perpetual futures contracts.
// No real liquidity exists in this pool.
type Pool struct {
	// always BASE:QUOTE, e.g. BTC:NUSD or ETH:NUSD
	Pair string `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	// base asset is the crypto asset, e.g. BTC or ETH
	BaseAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=base_asset_reserve,json=baseAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_asset_reserve"`
	// quote asset is usually stablecoin, in our case NUSD
	QuoteAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=quote_asset_reserve,json=quoteAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quote_asset_reserve"`
	// ratio applied to reserves in order not to over trade
	TradeLimitRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=trade_limit_ratio,json=tradeLimitRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"trade_limit_ratio"`
	// percentage that a single open or close position can alter the reserve amounts
	FluctuationLimitRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=fluctuation_limit_ratio,json=fluctuationLimitRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fluctuation_limit_ratio"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb4ac691d1b54d04, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func init() {
	proto.RegisterEnum("nibiru.vpool.v1.Direction", Direction_name, Direction_value)
	proto.RegisterEnum("nibiru.vpool.v1.TwapCalcOption", TwapCalcOption_name, TwapCalcOption_value)
	proto.RegisterType((*ReserveSnapshot)(nil), "nibiru.vpool.v1.ReserveSnapshot")
	proto.RegisterType((*Pool)(nil), "nibiru.vpool.v1.Pool")
}

func init() { proto.RegisterFile("vpool/v1/vpool.proto", fileDescriptor_eb4ac691d1b54d04) }

var fileDescriptor_eb4ac691d1b54d04 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x93, 0xb6, 0x20, 0xea, 0x21, 0xd2, 0x79, 0x45, 0x94, 0x09, 0x65, 0x63, 0x07, 0x34,
	0x0d, 0x91, 0x68, 0xf0, 0x04, 0xe9, 0x1a, 0x4e, 0xdb, 0x52, 0x92, 0x00, 0xd2, 0x84, 0xb0, 0x9c,
	0xcc, 0x6b, 0xad, 0x25, 0x71, 0xb0, 0x9d, 0x00, 0x6f, 0xb1, 0x07, 0xe0, 0x81, 0x76, 0xdc, 0x11,
	0x71, 0x98, 0x50, 0xfb, 0x22, 0x28, 0x76, 0xa5, 0xad, 0x48, 0x1c, 0xa8, 0xe0, 0x94, 0x2f, 0x3f,
	0xdb, 0xbf, 0xbf, 0xbe, 0x4f, 0xfa, 0x40, 0xbf, 0x2e, 0x19, 0xcb, 0xdc, 0x7a, 0xdf, 0x55, 0x85,
	0x53, 0x72, 0x26, 0x19, 0xb4, 0x0a, 0x9a, 0x50, 0x5e, 0x39, 0x9a, 0xd5, 0xfb, 0x9b, 0x8f, 0x53,
	0x26, 0x72, 0x26, 0x90, 0x3a, 0x76, 0xf5, 0x8f, 0xbe, 0xbb, 0xd9, 0x9f, 0xb0, 0x09, 0xd3, 0xbc,
	0xa9, 0x34, 0xdd, 0xb9, 0x68, 0x01, 0x2b, 0x24, 0x82, 0xf0, 0x9a, 0x44, 0x05, 0x2e, 0xc5, 0x94,
	0x49, 0xf8, 0x01, 0xc0, 0x04, 0x0b, 0x82, 0xb0, 0x10, 0x44, 0x22, 0xae, 0x4f, 0x07, 0xe6, 0xb6,
	0xb9, 0xdb, 0x1d, 0x3a, 0x97, 0xd7, 0x5b, 0xc6, 0x8f, 0xeb, 0xad, 0x67, 0x13, 0x2a, 0xa7, 0x55,
	0xe2, 0xa4, 0x2c, 0x5f, 0xc4, 0x2c, 0x3e, 0x2f, 0xc4, 0xe9, 0xb9, 0x2b, 0xbf, 0x96, 0x44, 0x38,
	0x23, 0x92, 0x86, 0xbd, 0xc6, 0xe4, 0x35, 0xa2, 0x45, 0x0a, 0xfc, 0x08, 0x36, 0x3e, 0x55, 0x4c,
	0xfe, 0xae, 0x6f, 0xad, 0xa4, 0x5f, 0x57, 0xaa, 0x25, 0xff, 0x13, 0xd0, 0x95, 0x34, 0x27, 0x42,
	0xe2, 0xbc, 0x1c, 0xb4, 0xb7, 0xcd, 0xdd, 0x76, 0x78, 0x03, 0xe0, 0x53, 0x70, 0x3f, 0xc9, 0x58,
	0x7a, 0x8e, 0x8a, 0x2a, 0x4f, 0x08, 0x1f, 0x74, 0xd4, 0x85, 0x35, 0xc5, 0x8e, 0x15, 0xda, 0xf9,
	0xd6, 0x06, 0x9d, 0x31, 0x63, 0x19, 0x84, 0xa0, 0x53, 0x62, 0xca, 0x75, 0xe7, 0xa1, 0xaa, 0xff,
	0x30, 0x9b, 0xd6, 0xff, 0x9d, 0x4d, 0xfb, 0x5f, 0xcd, 0xe6, 0x04, 0xac, 0x4b, 0x8e, 0x4f, 0x09,
	0xca, 0x68, 0x4e, 0x25, 0xe2, 0x58, 0x52, 0xa6, 0x46, 0xf0, 0xf7, 0x76, 0x4b, 0x89, 0x0e, 0x1b,
	0x4f, 0xd8, 0x68, 0xe0, 0x19, 0x78, 0x74, 0x96, 0x55, 0xa9, 0xac, 0x9a, 0xbf, 0x62, 0x29, 0xe1,
	0xce, 0x4a, 0x09, 0x0f, 0x6f, 0xe9, 0x6e, 0x72, 0xf6, 0x5e, 0x82, 0xee, 0x88, 0x72, 0x92, 0x36,
	0x18, 0x5a, 0x60, 0xcd, 0x1b, 0x8d, 0x50, 0x1c, 0xa0, 0x71, 0x10, 0x1c, 0xf6, 0x0c, 0xd8, 0x07,
	0xbd, 0xd0, 0x3f, 0x0a, 0xde, 0xf9, 0xe8, 0x75, 0x18, 0x1c, 0x69, 0x6a, 0xee, 0xf9, 0xe0, 0x41,
	0xfc, 0x19, 0x97, 0x07, 0x38, 0x4b, 0x83, 0x52, 0x3d, 0xbc, 0x07, 0x3a, 0xd1, 0x38, 0x88, 0xf5,
	0x8b, 0x37, 0x6f, 0x83, 0xd8, 0x47, 0x5e, 0x14, 0xf9, 0x31, 0x8a, 0xde, 0x7b, 0xe3, 0x9e, 0x09,
	0x37, 0x80, 0x35, 0xf4, 0xa2, 0x25, 0xd8, 0x1a, 0xfa, 0x97, 0x33, 0xdb, 0xbc, 0x9a, 0xd9, 0xe6,
	0xcf, 0x99, 0x6d, 0x5e, 0xcc, 0x6d, 0xe3, 0x6a, 0x6e, 0x1b, 0xdf, 0xe7, 0xb6, 0x71, 0xf2, 0xfc,
	0x56, 0x4f, 0xc7, 0x6a, 0x27, 0x0f, 0xa6, 0x98, 0x16, 0xae, 0xde, 0x4f, 0xf7, 0x8b, 0xde, 0x5a,
	0xdd, 0x5c, 0x72, 0x57, 0xad, 0xde, 0xab, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x5d, 0x6a,
	0x72, 0xd4, 0x03, 0x00, 0x00,
}

func (m *ReserveSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReserveSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReserveSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockNumber != 0 {
		i = encodeVarintVpool(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x20
	}
	if m.Timestamp != 0 {
		i = encodeVarintVpool(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.QuoteAssetReserve.Size()
		i -= size
		if _, err := m.QuoteAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.BaseAssetReserve.Size()
		i -= size
		if _, err := m.BaseAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FluctuationLimitRatio.Size()
		i -= size
		if _, err := m.FluctuationLimitRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TradeLimitRatio.Size()
		i -= size
		if _, err := m.TradeLimitRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.QuoteAssetReserve.Size()
		i -= size
		if _, err := m.QuoteAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BaseAssetReserve.Size()
		i -= size
		if _, err := m.BaseAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintVpool(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVpool(dAtA []byte, offset int, v uint64) int {
	offset -= sovVpool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReserveSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseAssetReserve.Size()
	n += 1 + l + sovVpool(uint64(l))
	l = m.QuoteAssetReserve.Size()
	n += 1 + l + sovVpool(uint64(l))
	if m.Timestamp != 0 {
		n += 1 + sovVpool(uint64(m.Timestamp))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovVpool(uint64(m.BlockNumber))
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovVpool(uint64(l))
	}
	l = m.BaseAssetReserve.Size()
	n += 1 + l + sovVpool(uint64(l))
	l = m.QuoteAssetReserve.Size()
	n += 1 + l + sovVpool(uint64(l))
	l = m.TradeLimitRatio.Size()
	n += 1 + l + sovVpool(uint64(l))
	l = m.FluctuationLimitRatio.Size()
	n += 1 + l + sovVpool(uint64(l))
	return n
}

func sovVpool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVpool(x uint64) (n int) {
	return sovVpool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReserveSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVpool
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
			return fmt.Errorf("proto: ReserveSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReserveSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
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
				return ErrInvalidLengthVpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
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
				return ErrInvalidLengthVpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVpool
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVpool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
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
				return ErrInvalidLengthVpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
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
				return ErrInvalidLengthVpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
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
				return ErrInvalidLengthVpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeLimitRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
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
				return ErrInvalidLengthVpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradeLimitRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FluctuationLimitRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVpool
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
				return ErrInvalidLengthVpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FluctuationLimitRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVpool
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
func skipVpool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVpool
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
					return 0, ErrIntOverflowVpool
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
					return 0, ErrIntOverflowVpool
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
				return 0, ErrInvalidLengthVpool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVpool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVpool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVpool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVpool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVpool = fmt.Errorf("proto: unexpected end of group")
)
