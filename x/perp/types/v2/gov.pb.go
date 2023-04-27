// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: perp/v2/gov.proto

package v2

import (
	fmt "fmt"
	github_com_NibiruChain_nibiru_x_common_asset "github.com/NibiruChain/nibiru/x/common/asset"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type CreatePoolProposal struct {
	Title       string                                            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Pair        github_com_NibiruChain_nibiru_x_common_asset.Pair `protobuf:"bytes,3,opt,name=pair,proto3,customtype=github.com/NibiruChain/nibiru/x/common/asset.Pair" json:"pair"`
	Market      Market                                            `protobuf:"bytes,4,opt,name=market,proto3" json:"market"`
	Amm         AMM                                               `protobuf:"bytes,5,opt,name=amm,proto3" json:"amm"`
}

func (m *CreatePoolProposal) Reset()         { *m = CreatePoolProposal{} }
func (m *CreatePoolProposal) String() string { return proto.CompactTextString(m) }
func (*CreatePoolProposal) ProtoMessage()    {}
func (*CreatePoolProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9fedff114e21530, []int{0}
}
func (m *CreatePoolProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePoolProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePoolProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePoolProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePoolProposal.Merge(m, src)
}
func (m *CreatePoolProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreatePoolProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePoolProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePoolProposal proto.InternalMessageInfo

func (m *CreatePoolProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePoolProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreatePoolProposal) GetMarket() Market {
	if m != nil {
		return m.Market
	}
	return Market{}
}

func (m *CreatePoolProposal) GetAmm() AMM {
	if m != nil {
		return m.Amm
	}
	return AMM{}
}

type EditPoolConfigProposal struct {
	Title       string                                            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Pair        github_com_NibiruChain_nibiru_x_common_asset.Pair `protobuf:"bytes,3,opt,name=pair,proto3,customtype=github.com/NibiruChain/nibiru/x/common/asset.Pair" json:"pair"`
	Market      Market                                            `protobuf:"bytes,4,opt,name=market,proto3" json:"market"`
	Amm         AMM                                               `protobuf:"bytes,5,opt,name=amm,proto3" json:"amm"`
}

func (m *EditPoolConfigProposal) Reset()         { *m = EditPoolConfigProposal{} }
func (m *EditPoolConfigProposal) String() string { return proto.CompactTextString(m) }
func (*EditPoolConfigProposal) ProtoMessage()    {}
func (*EditPoolConfigProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9fedff114e21530, []int{1}
}
func (m *EditPoolConfigProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditPoolConfigProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditPoolConfigProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditPoolConfigProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditPoolConfigProposal.Merge(m, src)
}
func (m *EditPoolConfigProposal) XXX_Size() int {
	return m.Size()
}
func (m *EditPoolConfigProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_EditPoolConfigProposal.DiscardUnknown(m)
}

var xxx_messageInfo_EditPoolConfigProposal proto.InternalMessageInfo

func (m *EditPoolConfigProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EditPoolConfigProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EditPoolConfigProposal) GetMarket() Market {
	if m != nil {
		return m.Market
	}
	return Market{}
}

func (m *EditPoolConfigProposal) GetAmm() AMM {
	if m != nil {
		return m.Amm
	}
	return AMM{}
}

// EditSwapInvariantsProposal is a governance proposal to change the swap
// invariant of the virtual pool for one or more trading pairs.
type EditSwapInvariantsProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Map from pair ID to a multiple on the swap invariant. For example, a
	// proposal containing "swap_invariant_maps": [{ "uatom:unusd": "5" }, {
	// "uosmo:unusd": "0.9" }] would mutliply the swap invariant of the ATOM and
	// OSMO trading pairs by 5 and 0.9 respectively. The price at which k changes
	// is the instantaneous mark price at the time of the proposal's execution.
	SwapInvariantMaps []EditSwapInvariantsProposal_SwapInvariantMultiple `protobuf:"bytes,5,rep,name=swap_invariant_maps,json=swapInvariantMaps,proto3" json:"swap_invariant_maps"`
}

func (m *EditSwapInvariantsProposal) Reset()         { *m = EditSwapInvariantsProposal{} }
func (m *EditSwapInvariantsProposal) String() string { return proto.CompactTextString(m) }
func (*EditSwapInvariantsProposal) ProtoMessage()    {}
func (*EditSwapInvariantsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9fedff114e21530, []int{2}
}
func (m *EditSwapInvariantsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditSwapInvariantsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditSwapInvariantsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditSwapInvariantsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditSwapInvariantsProposal.Merge(m, src)
}
func (m *EditSwapInvariantsProposal) XXX_Size() int {
	return m.Size()
}
func (m *EditSwapInvariantsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_EditSwapInvariantsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_EditSwapInvariantsProposal proto.InternalMessageInfo

func (m *EditSwapInvariantsProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EditSwapInvariantsProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EditSwapInvariantsProposal) GetSwapInvariantMaps() []EditSwapInvariantsProposal_SwapInvariantMultiple {
	if m != nil {
		return m.SwapInvariantMaps
	}
	return nil
}

// A map between a trading pair and a desired multiplier for its swap
// invariant.
type EditSwapInvariantsProposal_SwapInvariantMultiple struct {
	// Pair is a string identifier for an asset pair.
	Pair github_com_NibiruChain_nibiru_x_common_asset.Pair `protobuf:"bytes,3,opt,name=pair,proto3,customtype=github.com/NibiruChain/nibiru/x/common/asset.Pair" json:"pair"`
	// Multiplier is a number representing the desired percentage change to the
	// swap invariant of the AMM pool underlying 'pair'
	Multiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=multiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"multiplier"`
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Reset() {
	*m = EditSwapInvariantsProposal_SwapInvariantMultiple{}
}
func (*EditSwapInvariantsProposal_SwapInvariantMultiple) ProtoMessage() {}
func (*EditSwapInvariantsProposal_SwapInvariantMultiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9fedff114e21530, []int{2, 0}
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple.Merge(m, src)
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Size() int {
	return m.Size()
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_DiscardUnknown() {
	xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple.DiscardUnknown(m)
}

var xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreatePoolProposal)(nil), "nibiru.perp.v2.CreatePoolProposal")
	proto.RegisterType((*EditPoolConfigProposal)(nil), "nibiru.perp.v2.EditPoolConfigProposal")
	proto.RegisterType((*EditSwapInvariantsProposal)(nil), "nibiru.perp.v2.EditSwapInvariantsProposal")
	proto.RegisterType((*EditSwapInvariantsProposal_SwapInvariantMultiple)(nil), "nibiru.perp.v2.EditSwapInvariantsProposal.SwapInvariantMultiple")
}

func init() { proto.RegisterFile("perp/v2/gov.proto", fileDescriptor_d9fedff114e21530) }

var fileDescriptor_d9fedff114e21530 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xf7, 0x35, 0x49, 0x45, 0x2f, 0x12, 0x52, 0x9d, 0x52, 0x59, 0x19, 0x9c, 0xa8, 0x03, 0x8a,
	0x84, 0x7a, 0x27, 0x0c, 0x0b, 0x4c, 0x90, 0x80, 0x10, 0x83, 0xab, 0xc8, 0x6c, 0x2c, 0xd5, 0xc5,
	0x39, 0xdc, 0x53, 0x7d, 0x7e, 0xa7, 0xbb, 0x8b, 0x0b, 0x2b, 0x13, 0x1b, 0x8c, 0x8c, 0xfd, 0x28,
	0x8c, 0x1d, 0x3b, 0x22, 0x86, 0x0a, 0x25, 0xdf, 0x03, 0x21, 0x9f, 0x5d, 0x29, 0xa9, 0x40, 0x0c,
	0xb0, 0x31, 0xd9, 0x7e, 0xbf, 0x3f, 0xfe, 0xbd, 0xf7, 0xf4, 0xf0, 0xae, 0xe2, 0x5a, 0xd1, 0x32,
	0xa2, 0x19, 0x94, 0x44, 0x69, 0xb0, 0xe0, 0xdf, 0x2e, 0xc4, 0x4c, 0xe8, 0x05, 0xa9, 0x10, 0x52,
	0x46, 0xfd, 0xbd, 0x0c, 0x32, 0x70, 0x10, 0xad, 0xde, 0x6a, 0x56, 0xbf, 0x77, 0x2d, 0x34, 0x96,
	0x59, 0x5e, 0x17, 0x0f, 0x7e, 0x20, 0xec, 0x4f, 0x34, 0x67, 0x96, 0x4f, 0x01, 0xf2, 0xa9, 0x06,
	0x05, 0x86, 0xe5, 0xfe, 0x1e, 0xee, 0x58, 0x61, 0x73, 0x1e, 0xa0, 0x21, 0x1a, 0xed, 0x24, 0xf5,
	0x87, 0x3f, 0xc4, 0xdd, 0x39, 0x37, 0xa9, 0x16, 0xca, 0x0a, 0x28, 0x82, 0x2d, 0x87, 0xad, 0x97,
	0xfc, 0x18, 0xb7, 0x15, 0x13, 0x3a, 0x68, 0x55, 0xd0, 0xf8, 0xd1, 0xc5, 0xd5, 0xc0, 0xfb, 0x76,
	0x35, 0xb8, 0x9f, 0x09, 0x7b, 0xb2, 0x98, 0x91, 0x14, 0x24, 0x3d, 0x72, 0x51, 0x27, 0x27, 0x4c,
	0x14, 0xb4, 0x8e, 0x4d, 0xdf, 0xd2, 0x14, 0xa4, 0x84, 0x82, 0x32, 0x63, 0xb8, 0x25, 0x53, 0x26,
	0x74, 0xe2, 0x6c, 0xfc, 0x87, 0x78, 0x5b, 0x32, 0x7d, 0xca, 0x6d, 0xd0, 0x1e, 0xa2, 0x51, 0x37,
	0xda, 0x27, 0x9b, 0x9d, 0x92, 0xd8, 0xa1, 0xe3, 0x76, 0xf5, 0xa3, 0xa4, 0xe1, 0xfa, 0xf7, 0x70,
	0x8b, 0x49, 0x19, 0x74, 0x9c, 0xa4, 0x77, 0x53, 0xf2, 0x34, 0x8e, 0x1b, 0x7e, 0xc5, 0x3a, 0x78,
	0xbf, 0x85, 0xf7, 0x9f, 0xcf, 0x85, 0xad, 0xda, 0x9f, 0x40, 0xf1, 0x46, 0x64, 0xff, 0xe1, 0x10,
	0x3e, 0xb6, 0x70, 0xbf, 0x1a, 0xc2, 0xab, 0x33, 0xa6, 0x5e, 0x16, 0x25, 0xd3, 0x82, 0x15, 0xd6,
	0xfc, 0xf5, 0x20, 0x4a, 0xdc, 0x33, 0x67, 0x4c, 0x1d, 0x8b, 0x6b, 0xcb, 0x63, 0xc9, 0x94, 0x09,
	0x3a, 0xc3, 0xd6, 0xa8, 0x1b, 0x3d, 0xb9, 0x99, 0xe9, 0xf7, 0x01, 0xc8, 0x46, 0x39, 0x5e, 0xe4,
	0x56, 0xa8, 0x9c, 0x37, 0x0d, 0xec, 0x9a, 0x0d, 0x90, 0x29, 0xd3, 0xff, 0x82, 0xf0, 0x9d, 0x5f,
	0x4a, 0xfe, 0xf5, 0x6a, 0x8e, 0x30, 0x96, 0xb5, 0xb5, 0xe0, 0xda, 0xad, 0x67, 0x67, 0x4c, 0x1a,
	0xd3, 0xbb, 0x6b, 0xa6, 0x29, 0x18, 0x09, 0xa6, 0x79, 0x1c, 0x9a, 0xf9, 0x29, 0xb5, 0xef, 0x14,
	0x37, 0xe4, 0x19, 0x4f, 0x93, 0x35, 0x87, 0xc7, 0xb7, 0x3e, 0x9c, 0x0f, 0xbc, 0xcf, 0xe7, 0x03,
	0x6f, 0xfc, 0xe2, 0x62, 0x19, 0xa2, 0xcb, 0x65, 0x88, 0xbe, 0x2f, 0x43, 0xf4, 0x69, 0x15, 0x7a,
	0x97, 0xab, 0xd0, 0xfb, 0xba, 0x0a, 0xbd, 0xd7, 0x87, 0x7f, 0x0a, 0xeb, 0xce, 0xdc, 0xf9, 0xd3,
	0x32, 0x9a, 0x6d, 0xbb, 0x3b, 0x7f, 0xf0, 0x33, 0x00, 0x00, 0xff, 0xff, 0x32, 0x0d, 0x03, 0xa3,
	0x37, 0x04, 0x00, 0x00,
}

func (m *CreatePoolProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePoolProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePoolProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amm.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Market.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Pair.Size()
		i -= size
		if _, err := m.Pair.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EditPoolConfigProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditPoolConfigProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditPoolConfigProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amm.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Market.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Pair.Size()
		i -= size
		if _, err := m.Pair.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EditSwapInvariantsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditSwapInvariantsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditSwapInvariantsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SwapInvariantMaps) > 0 {
		for iNdEx := len(m.SwapInvariantMaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SwapInvariantMaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Multiplier.Size()
		i -= size
		if _, err := m.Multiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Pair.Size()
		i -= size
		if _, err := m.Pair.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreatePoolProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.Pair.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Market.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Amm.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *EditPoolConfigProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.Pair.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Market.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Amm.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *EditSwapInvariantsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.SwapInvariantMaps) > 0 {
		for _, e := range m.SwapInvariantMaps {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pair.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Multiplier.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreatePoolProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: CreatePoolProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePoolProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Market.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *EditPoolConfigProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: EditPoolConfigProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EditPoolConfigProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Market.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *EditSwapInvariantsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: EditSwapInvariantsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EditSwapInvariantsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapInvariantMaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SwapInvariantMaps = append(m.SwapInvariantMaps, EditSwapInvariantsProposal_SwapInvariantMultiple{})
			if err := m.SwapInvariantMaps[len(m.SwapInvariantMaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: SwapInvariantMultiple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapInvariantMultiple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Multiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
