// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: amm/amm.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
	Direction_UNDEFINED       Direction = 0
	Direction_ADD_TO_AMM      Direction = 1
	Direction_REMOVE_FROM_AMM Direction = 2
)

var Direction_name = map[int32]string{
	0: "UNDEFINED",
	1: "ADD_TO_AMM",
	2: "REMOVE_FROM_AMM",
}

var Direction_value = map[string]int32{
	"UNDEFINED":       0,
	"ADD_TO_AMM":      1,
	"REMOVE_FROM_AMM": 2,
}

func (x Direction) String() string {
	return proto.EnumName(Direction_name, int32(x))
}

func (Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b07d20e6d863e9f, []int{0}
}

type Pool struct {
	Pair string `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	// ratio applied to reserves in order not to over trade
	TradeLimitRatio   string `protobuf:"bytes,2,opt,name=trade_limit_ratio,json=tradeLimitRatio,proto3" json:"trade_limit_ratio,omitempty"`
	QuoteAssetReserve string `protobuf:"bytes,3,opt,name=quote_asset_reserve,json=quoteAssetReserve,proto3" json:"quote_asset_reserve,omitempty"`
	BaseAssetReserve  string `protobuf:"bytes,4,opt,name=base_asset_reserve,json=baseAssetReserve,proto3" json:"base_asset_reserve,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b07d20e6d863e9f, []int{0}
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

func (m *Pool) GetTradeLimitRatio() string {
	if m != nil {
		return m.TradeLimitRatio
	}
	return ""
}

func (m *Pool) GetQuoteAssetReserve() string {
	if m != nil {
		return m.QuoteAssetReserve
	}
	return ""
}

func (m *Pool) GetBaseAssetReserve() string {
	if m != nil {
		return m.BaseAssetReserve
	}
	return ""
}

func init() {
	proto.RegisterEnum("matrix.amm.v1.Direction", Direction_name, Direction_value)
	proto.RegisterType((*Pool)(nil), "matrix.amm.v1.Pool")
}

func init() { proto.RegisterFile("amm/amm.proto", fileDescriptor_4b07d20e6d863e9f) }

var fileDescriptor_4b07d20e6d863e9f = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xfb, 0x30,
	0x1c, 0xc7, 0x9b, 0xfd, 0xc7, 0x1f, 0xf6, 0x83, 0xba, 0x2e, 0xbb, 0xf4, 0x14, 0xc4, 0x8b, 0x32,
	0xa4, 0x45, 0x7c, 0x00, 0xa9, 0xa4, 0x03, 0xc1, 0xae, 0x52, 0xd4, 0x83, 0x97, 0x90, 0xce, 0xa0,
	0x81, 0xc5, 0xd4, 0x34, 0x1b, 0xf3, 0x2d, 0x7c, 0x08, 0x1f, 0xc6, 0xe3, 0x8e, 0x1e, 0xa5, 0x7d,
	0x11, 0x49, 0x76, 0x11, 0x6f, 0x5f, 0xbe, 0x9f, 0xcf, 0xe9, 0x03, 0x21, 0x57, 0x2a, 0xe5, 0x4a,
	0x25, 0x8d, 0xd1, 0x56, 0xe3, 0x50, 0x71, 0x6b, 0xe4, 0x36, 0x71, 0xcf, 0xe6, 0xec, 0xe8, 0x03,
	0xc1, 0xf0, 0x46, 0xeb, 0x15, 0xc6, 0x30, 0x6c, 0xb8, 0x34, 0x31, 0x3a, 0x44, 0x27, 0xa3, 0xca,
	0x6f, 0x3c, 0x83, 0x89, 0x35, 0xfc, 0x51, 0xb0, 0x95, 0x54, 0xd2, 0x32, 0xc3, 0xad, 0xd4, 0xf1,
	0xc0, 0x0b, 0x63, 0x0f, 0xae, 0xdd, 0x5f, 0xb9, 0x1b, 0x27, 0x30, 0x7d, 0x5d, 0x6b, 0x2b, 0x18,
	0x6f, 0x5b, 0x61, 0x99, 0x11, 0xad, 0x30, 0x1b, 0x11, 0xff, 0xf3, 0xf6, 0xc4, 0xa3, 0xcc, 0x91,
	0x6a, 0x0f, 0xf0, 0x29, 0xe0, 0x9a, 0xb7, 0x7f, 0xf5, 0xa1, 0xd7, 0x23, 0x47, 0x7e, 0xdb, 0xb3,
	0x0b, 0x18, 0x51, 0x69, 0xc4, 0xd2, 0x4a, 0xfd, 0x82, 0x43, 0x18, 0xdd, 0x2d, 0x68, 0x3e, 0xbf,
	0x5a, 0xe4, 0x34, 0x0a, 0xf0, 0x01, 0x40, 0x46, 0x29, 0xbb, 0x2d, 0x59, 0x56, 0x14, 0x11, 0xc2,
	0x53, 0x18, 0x57, 0x79, 0x51, 0xde, 0xe7, 0x6c, 0x5e, 0x95, 0x85, 0x3f, 0x07, 0x97, 0xd9, 0x67,
	0x47, 0xd0, 0xae, 0x23, 0xe8, 0xbb, 0x23, 0xe8, 0xbd, 0x27, 0xc1, 0xae, 0x27, 0xc1, 0x57, 0x4f,
	0x82, 0x87, 0xe3, 0x27, 0x69, 0x9f, 0xd7, 0x75, 0xb2, 0xd4, 0x2a, 0x2d, 0x7c, 0x1b, 0xca, 0x75,
	0xba, 0xaf, 0x94, 0x6e, 0x5d, 0xb9, 0xd4, 0xbe, 0x35, 0xa2, 0xad, 0xff, 0xfb, 0x80, 0xe7, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x23, 0x9f, 0x43, 0x51, 0x01, 0x00, 0x00,
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
	if len(m.BaseAssetReserve) > 0 {
		i -= len(m.BaseAssetReserve)
		copy(dAtA[i:], m.BaseAssetReserve)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.BaseAssetReserve)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.QuoteAssetReserve) > 0 {
		i -= len(m.QuoteAssetReserve)
		copy(dAtA[i:], m.QuoteAssetReserve)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.QuoteAssetReserve)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TradeLimitRatio) > 0 {
		i -= len(m.TradeLimitRatio)
		copy(dAtA[i:], m.TradeLimitRatio)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.TradeLimitRatio)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAmm(dAtA []byte, offset int, v uint64) int {
	offset -= sovAmm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	l = len(m.TradeLimitRatio)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	l = len(m.QuoteAssetReserve)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	l = len(m.BaseAssetReserve)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	return n
}

func sovAmm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAmm(x uint64) (n int) {
	return sovAmm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAmm
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
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeLimitRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeLimitRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAssetReserve = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAssetReserve = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAmm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAmm
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
func skipAmm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAmm
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
					return 0, ErrIntOverflowAmm
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
					return 0, ErrIntOverflowAmm
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
				return 0, ErrInvalidLengthAmm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAmm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAmm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAmm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAmm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAmm = fmt.Errorf("proto: unexpected end of group")
)
