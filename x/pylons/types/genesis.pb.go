// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the pylons module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	TradeList                    []*Trade                   `protobuf:"bytes,12,rep,name=tradeList,proto3" json:"tradeList,omitempty"`
	TradeCount                   uint64                     `protobuf:"varint,13,opt,name=tradeCount,proto3" json:"tradeCount,omitempty"`
	EntityCount                  uint64                     `protobuf:"varint,11,opt,name=entityCount,proto3" json:"entityCount,omitempty"`
	Params                       Params                     `protobuf:"bytes,10,opt,name=params,proto3" json:"params"`
	GoogleInAppPurchaseOrderList []GoogleInAppPurchaseOrder `protobuf:"bytes,8,rep,name=googleInAppPurchaseOrderList,proto3" json:"googleInAppPurchaseOrderList"`
	GoogleIAPOrderCount          uint64                     `protobuf:"varint,9,opt,name=googleIAPOrderCount,proto3" json:"googleIAPOrderCount,omitempty"`
	ExecutionList                []Execution                `protobuf:"bytes,7,rep,name=executionList,proto3" json:"executionList"`
	ExecutionCount               uint64                     `protobuf:"varint,6,opt,name=executionCount,proto3" json:"executionCount,omitempty"`
	PendingExecutionList         []Execution                `protobuf:"bytes,5,rep,name=pendingExecutionList,proto3" json:"pendingExecutionList"`
	PendingExecutionCount        uint64                     `protobuf:"varint,4,opt,name=pendingExecutionCount,proto3" json:"pendingExecutionCount,omitempty"`
	ItemList                     []Item                     `protobuf:"bytes,3,rep,name=itemList,proto3" json:"itemList"`
	RecipeList                   []Recipe                   `protobuf:"bytes,2,rep,name=recipeList,proto3" json:"recipeList"`
	CookbookList                 []Cookbook                 `protobuf:"bytes,1,rep,name=cookbookList,proto3" json:"cookbookList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6de94af8a698f1f8, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetTradeList() []*Trade {
	if m != nil {
		return m.TradeList
	}
	return nil
}

func (m *GenesisState) GetTradeCount() uint64 {
	if m != nil {
		return m.TradeCount
	}
	return 0
}

func (m *GenesisState) GetEntityCount() uint64 {
	if m != nil {
		return m.EntityCount
	}
	return 0
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGoogleInAppPurchaseOrderList() []GoogleInAppPurchaseOrder {
	if m != nil {
		return m.GoogleInAppPurchaseOrderList
	}
	return nil
}

func (m *GenesisState) GetGoogleIAPOrderCount() uint64 {
	if m != nil {
		return m.GoogleIAPOrderCount
	}
	return 0
}

func (m *GenesisState) GetExecutionList() []Execution {
	if m != nil {
		return m.ExecutionList
	}
	return nil
}

func (m *GenesisState) GetExecutionCount() uint64 {
	if m != nil {
		return m.ExecutionCount
	}
	return 0
}

func (m *GenesisState) GetPendingExecutionList() []Execution {
	if m != nil {
		return m.PendingExecutionList
	}
	return nil
}

func (m *GenesisState) GetPendingExecutionCount() uint64 {
	if m != nil {
		return m.PendingExecutionCount
	}
	return 0
}

func (m *GenesisState) GetItemList() []Item {
	if m != nil {
		return m.ItemList
	}
	return nil
}

func (m *GenesisState) GetRecipeList() []Recipe {
	if m != nil {
		return m.RecipeList
	}
	return nil
}

func (m *GenesisState) GetCookbookList() []Cookbook {
	if m != nil {
		return m.CookbookList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "Pylonstech.pylons.pylons.GenesisState")
}

func init() { proto.RegisterFile("pylons/genesis.proto", fileDescriptor_6de94af8a698f1f8) }

var fileDescriptor_6de94af8a698f1f8 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0xce, 0xd9, 0x18, 0xdb, 0x49, 0x2a, 0xb8, 0x4d, 0x25, 0x04, 0xbd, 0x1e, 0x15, 0x24, 0x0f,
	0x7a, 0x91, 0xe8, 0xab, 0x62, 0x5b, 0xda, 0x52, 0x28, 0x34, 0x44, 0x9f, 0x04, 0x29, 0x97, 0xcb,
	0x70, 0x59, 0xda, 0xdc, 0x2e, 0x77, 0x1b, 0x68, 0xc0, 0x1f, 0xe1, 0xcf, 0xea, 0x63, 0x1f, 0x7d,
	0x12, 0x49, 0xf0, 0x7f, 0xc8, 0xcd, 0xce, 0xc5, 0x24, 0xe4, 0x0a, 0x7d, 0xca, 0xe6, 0x9b, 0x6f,
	0xbe, 0xf9, 0xbe, 0xbd, 0x1d, 0xa8, 0xeb, 0xc9, 0xb5, 0x8a, 0xd3, 0x76, 0x84, 0x31, 0xa6, 0x32,
	0xf5, 0x75, 0xa2, 0x8c, 0x12, 0x8d, 0x2e, 0xa1, 0x06, 0xc3, 0xa1, 0x6f, 0x09, 0xfc, 0xd3, 0x14,
	0xcc, 0x37, 0x49, 0x30, 0x40, 0xcb, 0x6e, 0xbe, 0xcc, 0x35, 0x94, 0x8a, 0xae, 0xf1, 0x52, 0x06,
	0xfa, 0x52, 0x25, 0x03, 0x4c, 0xb8, 0xfc, 0x9c, 0xcb, 0x78, 0x83, 0xe1, 0xd8, 0x48, 0x15, 0x33,
	0xfe, 0x8c, 0x71, 0x69, 0x70, 0xc4, 0xd0, 0x0e, 0x43, 0x09, 0x86, 0x52, 0xe7, 0xf2, 0xbb, 0x0c,
	0x86, 0x4a, 0x5d, 0xf5, 0x95, 0xba, 0x5a, 0xe1, 0xea, 0x20, 0x09, 0x46, 0x6c, 0xbc, 0x59, 0x8f,
	0x54, 0xa4, 0xe8, 0xd8, 0xce, 0x4e, 0x16, 0xdd, 0xff, 0x5b, 0x81, 0xda, 0xa9, 0x0d, 0xf8, 0xc5,
	0x04, 0x06, 0xc5, 0x47, 0xd8, 0xa2, 0x00, 0xe7, 0x32, 0x35, 0x8d, 0x9a, 0xb7, 0xd1, 0xaa, 0x76,
	0xf6, 0xfc, 0xa2, 0xcc, 0xfe, 0xd7, 0x8c, 0xda, 0xfb, 0xdf, 0x21, 0x5c, 0x00, 0xfa, 0x73, 0xa4,
	0xc6, 0xb1, 0x69, 0x6c, 0x7b, 0x4e, 0xab, 0xdc, 0x5b, 0x40, 0x84, 0x07, 0x55, 0x8c, 0x8d, 0x34,
	0x13, 0x4b, 0xa8, 0x12, 0x61, 0x11, 0x12, 0x9f, 0xa0, 0x62, 0x7d, 0x37, 0xc0, 0x73, 0x5a, 0xd5,
	0x8e, 0x57, 0x3c, 0xbd, 0x4b, 0xbc, 0xc3, 0xf2, 0xed, 0xef, 0xbd, 0x52, 0x8f, 0xbb, 0xc4, 0x0f,
	0x78, 0x61, 0x6f, 0xfb, 0x2c, 0x3e, 0xd0, 0xba, 0x3b, 0x4e, 0xc2, 0x61, 0x90, 0xe2, 0x45, 0x76,
	0xeb, 0x94, 0x69, 0x93, 0x32, 0x75, 0x8a, 0x55, 0x4f, 0x0b, 0xba, 0x79, 0xce, 0xbd, 0xea, 0xe2,
	0x1d, 0xec, 0x70, 0xfd, 0xa0, 0x4b, 0xa8, 0xcd, 0xb9, 0x45, 0x39, 0xd7, 0x95, 0xc4, 0x05, 0x6c,
	0xcf, 0x3f, 0x3f, 0x19, 0x7c, 0x42, 0x06, 0x5f, 0x15, 0x1b, 0x3c, 0xce, 0xe9, 0xec, 0x68, 0xb9,
	0x5f, 0xbc, 0x86, 0xa7, 0x73, 0xc0, 0x4e, 0xaf, 0xd0, 0xf4, 0x15, 0x54, 0x7c, 0x87, 0xba, 0xc6,
	0x78, 0x20, 0xe3, 0xe8, 0x78, 0x69, 0xfe, 0xe3, 0x87, 0xce, 0x5f, 0x2b, 0x23, 0x3e, 0xc0, 0xee,
	0x2a, 0x6e, 0xdd, 0x94, 0xc9, 0xcd, 0xfa, 0xa2, 0xf8, 0x0c, 0x9b, 0xd9, 0xa3, 0x27, 0x23, 0x1b,
	0x64, 0xc4, 0x2d, 0x36, 0x72, 0x66, 0x70, 0xc4, 0x1e, 0xe6, 0x5d, 0xe2, 0x04, 0xc0, 0xee, 0x08,
	0x69, 0x3c, 0x22, 0x8d, 0x7b, 0xde, 0x50, 0x8f, 0xb8, 0xac, 0xb2, 0xd0, 0x29, 0xce, 0xa1, 0x96,
	0xaf, 0x15, 0x29, 0x39, 0xa4, 0xb4, 0x5f, 0xac, 0x74, 0xc4, 0x6c, 0xd6, 0x5a, 0xea, 0x3e, 0x3c,
	0xb9, 0x9d, 0xba, 0xce, 0xdd, 0xd4, 0x75, 0xfe, 0x4c, 0x5d, 0xe7, 0xe7, 0xcc, 0x2d, 0xdd, 0xcd,
	0xdc, 0xd2, 0xaf, 0x99, 0x5b, 0xfa, 0xf6, 0x26, 0x92, 0x66, 0x38, 0xee, 0xfb, 0xa1, 0x1a, 0xb5,
	0xad, 0xf6, 0xdb, 0x4c, 0xbc, 0xcd, 0x3b, 0x7c, 0x93, 0x1f, 0xcc, 0x44, 0x63, 0xda, 0xaf, 0xd0,
	0xda, 0xbe, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x36, 0x81, 0x45, 0x49, 0x9d, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TradeCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TradeCount))
		i--
		dAtA[i] = 0x68
	}
	if len(m.TradeList) > 0 {
		for iNdEx := len(m.TradeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TradeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.EntityCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EntityCount))
		i--
		dAtA[i] = 0x58
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.GoogleIAPOrderCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GoogleIAPOrderCount))
		i--
		dAtA[i] = 0x48
	}
	if len(m.GoogleInAppPurchaseOrderList) > 0 {
		for iNdEx := len(m.GoogleInAppPurchaseOrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GoogleInAppPurchaseOrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ExecutionList) > 0 {
		for iNdEx := len(m.ExecutionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExecutionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ExecutionCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ExecutionCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PendingExecutionList) > 0 {
		for iNdEx := len(m.PendingExecutionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingExecutionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PendingExecutionCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PendingExecutionCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ItemList) > 0 {
		for iNdEx := len(m.ItemList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ItemList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RecipeList) > 0 {
		for iNdEx := len(m.RecipeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecipeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CookbookList) > 0 {
		for iNdEx := len(m.CookbookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CookbookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CookbookList) > 0 {
		for _, e := range m.CookbookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RecipeList) > 0 {
		for _, e := range m.RecipeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ItemList) > 0 {
		for _, e := range m.ItemList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PendingExecutionCount != 0 {
		n += 1 + sovGenesis(uint64(m.PendingExecutionCount))
	}
	if len(m.PendingExecutionList) > 0 {
		for _, e := range m.PendingExecutionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ExecutionCount != 0 {
		n += 1 + sovGenesis(uint64(m.ExecutionCount))
	}
	if len(m.ExecutionList) > 0 {
		for _, e := range m.ExecutionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GoogleInAppPurchaseOrderList) > 0 {
		for _, e := range m.GoogleInAppPurchaseOrderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.GoogleIAPOrderCount != 0 {
		n += 1 + sovGenesis(uint64(m.GoogleIAPOrderCount))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EntityCount != 0 {
		n += 1 + sovGenesis(uint64(m.EntityCount))
	}
	if len(m.TradeList) > 0 {
		for _, e := range m.TradeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TradeCount != 0 {
		n += 1 + sovGenesis(uint64(m.TradeCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CookbookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CookbookList = append(m.CookbookList, Cookbook{})
			if err := m.CookbookList[len(m.CookbookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipeList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipeList = append(m.RecipeList, Recipe{})
			if err := m.RecipeList[len(m.RecipeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemList = append(m.ItemList, Item{})
			if err := m.ItemList[len(m.ItemList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingExecutionCount", wireType)
			}
			m.PendingExecutionCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PendingExecutionCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingExecutionList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingExecutionList = append(m.PendingExecutionList, Execution{})
			if err := m.PendingExecutionList[len(m.PendingExecutionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionCount", wireType)
			}
			m.ExecutionCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutionList = append(m.ExecutionList, Execution{})
			if err := m.ExecutionList[len(m.ExecutionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleInAppPurchaseOrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleInAppPurchaseOrderList = append(m.GoogleInAppPurchaseOrderList, GoogleInAppPurchaseOrder{})
			if err := m.GoogleInAppPurchaseOrderList[len(m.GoogleInAppPurchaseOrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleIAPOrderCount", wireType)
			}
			m.GoogleIAPOrderCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GoogleIAPOrderCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityCount", wireType)
			}
			m.EntityCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EntityCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeList = append(m.TradeList, &Trade{})
			if err := m.TradeList[len(m.TradeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeCount", wireType)
			}
			m.TradeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradeCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
