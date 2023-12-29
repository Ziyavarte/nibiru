// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nibiru/oracle/v1/event.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Emitted when a price is posted
type EventPriceUpdate struct {
	Pair        string                                 `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Price       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	TimestampMs int64                                  `protobuf:"varint,3,opt,name=timestamp_ms,json=timestampMs,proto3" json:"timestamp_ms,omitempty"`
}

func (m *EventPriceUpdate) Reset()         { *m = EventPriceUpdate{} }
func (m *EventPriceUpdate) String() string { return proto.CompactTextString(m) }
func (*EventPriceUpdate) ProtoMessage()    {}
func (*EventPriceUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ec441b793fc0ea, []int{0}
}
func (m *EventPriceUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPriceUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPriceUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPriceUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPriceUpdate.Merge(m, src)
}
func (m *EventPriceUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventPriceUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPriceUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventPriceUpdate proto.InternalMessageInfo

func (m *EventPriceUpdate) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func (m *EventPriceUpdate) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

// Emitted when a valoper delegates oracle voting rights to a feeder address.
type EventDelegateFeederConsent struct {
	// Validator is the Bech32 address that is delegating voting rights.
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	// Feeder is the delegate or representative that will be able to send
	// vote and prevote transaction messages.
	Feeder string `protobuf:"bytes,2,opt,name=feeder,proto3" json:"feeder,omitempty"`
}

func (m *EventDelegateFeederConsent) Reset()         { *m = EventDelegateFeederConsent{} }
func (m *EventDelegateFeederConsent) String() string { return proto.CompactTextString(m) }
func (*EventDelegateFeederConsent) ProtoMessage()    {}
func (*EventDelegateFeederConsent) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ec441b793fc0ea, []int{1}
}
func (m *EventDelegateFeederConsent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDelegateFeederConsent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDelegateFeederConsent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDelegateFeederConsent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDelegateFeederConsent.Merge(m, src)
}
func (m *EventDelegateFeederConsent) XXX_Size() int {
	return m.Size()
}
func (m *EventDelegateFeederConsent) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDelegateFeederConsent.DiscardUnknown(m)
}

var xxx_messageInfo_EventDelegateFeederConsent proto.InternalMessageInfo

func (m *EventDelegateFeederConsent) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *EventDelegateFeederConsent) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

// Emitted by MsgAggregateExchangeVote when an aggregate vote is added to state
type EventAggregateVote struct {
	// Validator is the Bech32 address to which the vote will be credited.
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	// Feeder is the delegate or representative that will send vote and prevote
	// transaction messages on behalf of the voting validator.
	Feeder string             `protobuf:"bytes,2,opt,name=feeder,proto3" json:"feeder,omitempty"`
	Prices ExchangeRateTuples `protobuf:"bytes,3,rep,name=prices,proto3,castrepeated=ExchangeRateTuples" json:"prices"`
}

func (m *EventAggregateVote) Reset()         { *m = EventAggregateVote{} }
func (m *EventAggregateVote) String() string { return proto.CompactTextString(m) }
func (*EventAggregateVote) ProtoMessage()    {}
func (*EventAggregateVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ec441b793fc0ea, []int{2}
}
func (m *EventAggregateVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAggregateVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAggregateVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAggregateVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAggregateVote.Merge(m, src)
}
func (m *EventAggregateVote) XXX_Size() int {
	return m.Size()
}
func (m *EventAggregateVote) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAggregateVote.DiscardUnknown(m)
}

var xxx_messageInfo_EventAggregateVote proto.InternalMessageInfo

func (m *EventAggregateVote) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *EventAggregateVote) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

func (m *EventAggregateVote) GetPrices() ExchangeRateTuples {
	if m != nil {
		return m.Prices
	}
	return nil
}

// Emitted by MsgAggregateExchangePrevote when an aggregate prevote is added
// to state
type EventAggregatePrevote struct {
	// Validator is the Bech32 address to which the vote will be credited.
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	// Feeder is the delegate or representative that will send vote and prevote
	// transaction messages on behalf of the voting validator.
	Feeder string `protobuf:"bytes,2,opt,name=feeder,proto3" json:"feeder,omitempty"`
}

func (m *EventAggregatePrevote) Reset()         { *m = EventAggregatePrevote{} }
func (m *EventAggregatePrevote) String() string { return proto.CompactTextString(m) }
func (*EventAggregatePrevote) ProtoMessage()    {}
func (*EventAggregatePrevote) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ec441b793fc0ea, []int{3}
}
func (m *EventAggregatePrevote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAggregatePrevote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAggregatePrevote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAggregatePrevote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAggregatePrevote.Merge(m, src)
}
func (m *EventAggregatePrevote) XXX_Size() int {
	return m.Size()
}
func (m *EventAggregatePrevote) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAggregatePrevote.DiscardUnknown(m)
}

var xxx_messageInfo_EventAggregatePrevote proto.InternalMessageInfo

func (m *EventAggregatePrevote) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *EventAggregatePrevote) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

type EventValidatorPerformance struct {
	// Validator is the Bech32 address to which the vote will be credited.
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	// Tendermint consensus voting power
	VotingPower int64 `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	// RewardWeight: Weight of rewards the validator should receive in units of
	// consensus power.
	RewardWeight int64 `protobuf:"varint,3,opt,name=reward_weight,json=rewardWeight,proto3" json:"reward_weight,omitempty"`
	// Number of valid votes for which the validator will be rewarded
	WinCount int64 `protobuf:"varint,4,opt,name=win_count,json=winCount,proto3" json:"win_count,omitempty"`
	// Number of abstained votes for which there will be no reward or punishment
	AbstainCount int64 `protobuf:"varint,5,opt,name=abstain_count,json=abstainCount,proto3" json:"abstain_count,omitempty"`
	// Number of invalid/punishable votes
	MissCount int64 `protobuf:"varint,6,opt,name=miss_count,json=missCount,proto3" json:"miss_count,omitempty"`
}

func (m *EventValidatorPerformance) Reset()         { *m = EventValidatorPerformance{} }
func (m *EventValidatorPerformance) String() string { return proto.CompactTextString(m) }
func (*EventValidatorPerformance) ProtoMessage()    {}
func (*EventValidatorPerformance) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ec441b793fc0ea, []int{4}
}
func (m *EventValidatorPerformance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventValidatorPerformance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventValidatorPerformance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventValidatorPerformance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventValidatorPerformance.Merge(m, src)
}
func (m *EventValidatorPerformance) XXX_Size() int {
	return m.Size()
}
func (m *EventValidatorPerformance) XXX_DiscardUnknown() {
	xxx_messageInfo_EventValidatorPerformance.DiscardUnknown(m)
}

var xxx_messageInfo_EventValidatorPerformance proto.InternalMessageInfo

func (m *EventValidatorPerformance) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *EventValidatorPerformance) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *EventValidatorPerformance) GetRewardWeight() int64 {
	if m != nil {
		return m.RewardWeight
	}
	return 0
}

func (m *EventValidatorPerformance) GetWinCount() int64 {
	if m != nil {
		return m.WinCount
	}
	return 0
}

func (m *EventValidatorPerformance) GetAbstainCount() int64 {
	if m != nil {
		return m.AbstainCount
	}
	return 0
}

func (m *EventValidatorPerformance) GetMissCount() int64 {
	if m != nil {
		return m.MissCount
	}
	return 0
}

func init() {
	proto.RegisterType((*EventPriceUpdate)(nil), "nibiru.oracle.v1.EventPriceUpdate")
	proto.RegisterType((*EventDelegateFeederConsent)(nil), "nibiru.oracle.v1.EventDelegateFeederConsent")
	proto.RegisterType((*EventAggregateVote)(nil), "nibiru.oracle.v1.EventAggregateVote")
	proto.RegisterType((*EventAggregatePrevote)(nil), "nibiru.oracle.v1.EventAggregatePrevote")
	proto.RegisterType((*EventValidatorPerformance)(nil), "nibiru.oracle.v1.EventValidatorPerformance")
}

func init() { proto.RegisterFile("nibiru/oracle/v1/event.proto", fileDescriptor_94ec441b793fc0ea) }

var fileDescriptor_94ec441b793fc0ea = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0xe8, 0x56, 0x51, 0xb7, 0x48, 0x93, 0x05, 0xa8, 0x94, 0x2e, 0xdd, 0x3a, 0x09, 0xf5,
	0x02, 0x12, 0x0d, 0x9e, 0x80, 0xb6, 0xdb, 0xdd, 0x50, 0x15, 0xc1, 0x26, 0x71, 0x53, 0xb9, 0xc9,
	0x99, 0x6b, 0x91, 0xd8, 0x91, 0xed, 0x26, 0xe3, 0x29, 0xe0, 0x1d, 0xb8, 0xe3, 0x49, 0x76, 0xb9,
	0x4b, 0xc4, 0xc5, 0x40, 0xed, 0x8b, 0x20, 0x3b, 0xee, 0xf8, 0xd9, 0x05, 0xd2, 0xae, 0x72, 0xfc,
	0x7d, 0x9f, 0xbf, 0x7c, 0xc9, 0x39, 0x07, 0xf5, 0x38, 0x9b, 0x33, 0xb9, 0x0c, 0x85, 0x24, 0x71,
	0x0a, 0x61, 0x71, 0x18, 0x42, 0x01, 0x5c, 0x07, 0xb9, 0x14, 0x5a, 0xe0, 0x9d, 0x8a, 0x0d, 0x2a,
	0x36, 0x28, 0x0e, 0xbb, 0xbb, 0xb7, 0xf4, 0x8e, 0xb3, 0x17, 0xba, 0x0f, 0xa9, 0xa0, 0xc2, 0x96,
	0xa1, 0xa9, 0x1c, 0xda, 0xa3, 0x42, 0xd0, 0x14, 0x42, 0x92, 0xb3, 0x90, 0x70, 0x2e, 0x34, 0xd1,
	0x4c, 0x70, 0x55, 0xb1, 0x83, 0x4f, 0x1e, 0xda, 0x39, 0x32, 0x2f, 0x9d, 0x4a, 0x16, 0xc3, 0xbb,
	0x3c, 0x21, 0x1a, 0x30, 0x46, 0x5b, 0x39, 0x61, 0xb2, 0xe3, 0xed, 0x79, 0xc3, 0x66, 0x64, 0x6b,
	0x3c, 0x41, 0xdb, 0xb9, 0x91, 0x74, 0xee, 0x19, 0x70, 0x14, 0x5c, 0x5e, 0xf7, 0x6b, 0xdf, 0xaf,
	0xfb, 0xcf, 0x28, 0xd3, 0x8b, 0xe5, 0x3c, 0x88, 0x45, 0x16, 0xc6, 0x42, 0x65, 0x42, 0xb9, 0xc7,
	0x0b, 0x95, 0x7c, 0x08, 0xf5, 0xc7, 0x1c, 0x54, 0x30, 0x81, 0x38, 0xaa, 0x2e, 0xe3, 0x7d, 0xd4,
	0xd6, 0x2c, 0x03, 0xa5, 0x49, 0x96, 0xcf, 0x32, 0xd5, 0xa9, 0xef, 0x79, 0xc3, 0x7a, 0xd4, 0xba,
	0xc1, 0x4e, 0xd4, 0x20, 0x42, 0x5d, 0x1b, 0x68, 0x02, 0x29, 0x50, 0xa2, 0xe1, 0x18, 0x20, 0x01,
	0x39, 0x16, 0x5c, 0x01, 0xd7, 0xb8, 0x87, 0x9a, 0x05, 0x49, 0x59, 0x42, 0xb4, 0xd8, 0xe4, 0xfb,
	0x0d, 0xe0, 0xc7, 0xa8, 0x71, 0x6e, 0xe5, 0x55, 0xca, 0xc8, 0x9d, 0x06, 0x5f, 0x3c, 0x84, 0xad,
	0xe9, 0x6b, 0x4a, 0xa5, 0x75, 0x3d, 0x15, 0x1a, 0xee, 0x66, 0x86, 0xcf, 0x50, 0xc3, 0x7e, 0x8c,
	0x49, 0x5f, 0x1f, 0xb6, 0x5e, 0x1e, 0x04, 0xff, 0x36, 0x2a, 0x38, 0xba, 0x88, 0x17, 0x84, 0x53,
	0x88, 0x88, 0x86, 0xb7, 0xcb, 0x3c, 0x85, 0x51, 0xd7, 0xfc, 0xaf, 0xaf, 0x3f, 0xfa, 0xf8, 0x16,
	0xa5, 0x22, 0x67, 0x37, 0x38, 0x41, 0x8f, 0xfe, 0x0e, 0x39, 0x95, 0x50, 0xdc, 0x39, 0xe7, 0x60,
	0xe5, 0xa1, 0x27, 0xd6, 0xef, 0x74, 0x23, 0x9d, 0x82, 0x3c, 0x17, 0x32, 0x23, 0x3c, 0xfe, 0x9f,
	0xe7, 0x3e, 0x6a, 0x17, 0x42, 0x33, 0x4e, 0x67, 0xb9, 0x28, 0x9d, 0x73, 0x3d, 0x6a, 0x55, 0xd8,
	0xd4, 0x40, 0xf8, 0x00, 0x3d, 0x90, 0x50, 0x12, 0x99, 0xcc, 0x4a, 0x60, 0x74, 0xa1, 0x5d, 0x2f,
	0xdb, 0x15, 0x78, 0x66, 0x31, 0xfc, 0x14, 0x35, 0x4b, 0xc6, 0x67, 0xb1, 0x58, 0x72, 0xdd, 0xd9,
	0xb2, 0x82, 0xfb, 0x25, 0xe3, 0x63, 0x73, 0x36, 0x0e, 0x64, 0xae, 0x34, 0xb9, 0x11, 0x6c, 0x57,
	0x0e, 0x0e, 0xac, 0x44, 0xbb, 0x08, 0x65, 0x4c, 0x29, 0xa7, 0x68, 0x58, 0x45, 0xd3, 0x20, 0x96,
	0x1e, 0x1d, 0x5f, 0xae, 0x7c, 0xef, 0x6a, 0xe5, 0x7b, 0x3f, 0x57, 0xbe, 0xf7, 0x79, 0xed, 0xd7,
	0xae, 0xd6, 0x7e, 0xed, 0xdb, 0xda, 0xaf, 0xbd, 0x7f, 0xfe, 0xc7, 0x64, 0xbe, 0xb1, 0x0d, 0x1a,
	0x2f, 0x08, 0xe3, 0xa1, 0xdb, 0xa1, 0x8b, 0xcd, 0x16, 0xd9, 0x19, 0x9d, 0x37, 0xec, 0x3a, 0xbc,
	0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xd7, 0x9a, 0x89, 0x93, 0x03, 0x00, 0x00,
}

func (m *EventPriceUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPriceUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPriceUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimestampMs != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.TimestampMs))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventDelegateFeederConsent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDelegateFeederConsent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDelegateFeederConsent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAggregateVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAggregateVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAggregateVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAggregatePrevote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAggregatePrevote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAggregatePrevote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventValidatorPerformance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventValidatorPerformance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventValidatorPerformance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissCount != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.MissCount))
		i--
		dAtA[i] = 0x30
	}
	if m.AbstainCount != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.AbstainCount))
		i--
		dAtA[i] = 0x28
	}
	if m.WinCount != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.WinCount))
		i--
		dAtA[i] = 0x20
	}
	if m.RewardWeight != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.RewardWeight))
		i--
		dAtA[i] = 0x18
	}
	if m.VotingPower != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventPriceUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovEvent(uint64(l))
	if m.TimestampMs != 0 {
		n += 1 + sovEvent(uint64(m.TimestampMs))
	}
	return n
}

func (m *EventDelegateFeederConsent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventAggregateVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *EventAggregatePrevote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventValidatorPerformance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovEvent(uint64(m.VotingPower))
	}
	if m.RewardWeight != 0 {
		n += 1 + sovEvent(uint64(m.RewardWeight))
	}
	if m.WinCount != 0 {
		n += 1 + sovEvent(uint64(m.WinCount))
	}
	if m.AbstainCount != 0 {
		n += 1 + sovEvent(uint64(m.AbstainCount))
	}
	if m.MissCount != 0 {
		n += 1 + sovEvent(uint64(m.MissCount))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventPriceUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventPriceUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPriceUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampMs", wireType)
			}
			m.TimestampMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimestampMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventDelegateFeederConsent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventDelegateFeederConsent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDelegateFeederConsent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventAggregateVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventAggregateVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAggregateVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prices = append(m.Prices, ExchangeRateTuple{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventAggregatePrevote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventAggregatePrevote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAggregatePrevote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventValidatorPerformance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventValidatorPerformance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventValidatorPerformance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardWeight", wireType)
			}
			m.RewardWeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardWeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinCount", wireType)
			}
			m.WinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WinCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbstainCount", wireType)
			}
			m.AbstainCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AbstainCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCount", wireType)
			}
			m.MissCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
