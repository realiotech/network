// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: realionetwork/asset/v1/token.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Token represents an asset in the module
type Token struct {
	TokenId     string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol      string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimal     string `protobuf:"bytes,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f83138fc60a3176, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *Token) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Token) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Token) GetDecimal() string {
	if m != nil {
		return m.Decimal
	}
	return ""
}

func (m *Token) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type TokenManagement struct {
	Manager            string   `protobuf:"bytes,1,opt,name=manager,proto3" json:"manager,omitempty"`
	AddNewPrivilege    bool     `protobuf:"varint,2,opt,name=add_new_privilege,json=addNewPrivilege,proto3" json:"add_new_privilege,omitempty"`
	ExcludedPrivileges []string `protobuf:"bytes,3,rep,name=excluded_privileges,json=excludedPrivileges,proto3" json:"excluded_privileges,omitempty"`
}

func (m *TokenManagement) Reset()         { *m = TokenManagement{} }
func (m *TokenManagement) String() string { return proto.CompactTextString(m) }
func (*TokenManagement) ProtoMessage()    {}
func (*TokenManagement) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f83138fc60a3176, []int{1}
}
func (m *TokenManagement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenManagement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenManagement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenManagement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenManagement.Merge(m, src)
}
func (m *TokenManagement) XXX_Size() int {
	return m.Size()
}
func (m *TokenManagement) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenManagement.DiscardUnknown(m)
}

var xxx_messageInfo_TokenManagement proto.InternalMessageInfo

func (m *TokenManagement) GetManager() string {
	if m != nil {
		return m.Manager
	}
	return ""
}

func (m *TokenManagement) GetAddNewPrivilege() bool {
	if m != nil {
		return m.AddNewPrivilege
	}
	return false
}

func (m *TokenManagement) GetExcludedPrivileges() []string {
	if m != nil {
		return m.ExcludedPrivileges
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "realionetwork.asset.v1.Token")
	proto.RegisterType((*TokenManagement)(nil), "realionetwork.asset.v1.TokenManagement")
}

func init() {
	proto.RegisterFile("realionetwork/asset/v1/token.proto", fileDescriptor_2f83138fc60a3176)
}

var fileDescriptor_2f83138fc60a3176 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0xee, 0xd3, 0x30,
	0x14, 0xc5, 0x13, 0xfa, 0x6d, 0x86, 0x0a, 0x53, 0x55, 0x6e, 0x87, 0x50, 0x75, 0xaa, 0x90, 0x9a,
	0xa8, 0x65, 0x63, 0xa3, 0x1b, 0x12, 0x20, 0x14, 0x98, 0x58, 0x22, 0x37, 0xbe, 0x4a, 0xad, 0x26,
	0x76, 0x64, 0xbb, 0x5f, 0x6f, 0xd1, 0x47, 0x40, 0x3c, 0x03, 0x0f, 0xc1, 0x58, 0x31, 0x31, 0xa2,
	0x76, 0xe1, 0x31, 0x50, 0xed, 0x04, 0xfe, 0xff, 0xed, 0x9e, 0x73, 0x7e, 0xb2, 0xef, 0xd5, 0x41,
	0x53, 0x05, 0x34, 0xe7, 0x52, 0x80, 0x39, 0x48, 0xb5, 0x8d, 0xa8, 0xd6, 0x60, 0xa2, 0xfd, 0x22,
	0x32, 0x72, 0x0b, 0x22, 0x2c, 0x95, 0x34, 0x12, 0x0f, 0x1f, 0x31, 0xa1, 0x65, 0xc2, 0xfd, 0x62,
	0x3c, 0xc8, 0x64, 0x26, 0x2d, 0x12, 0xdd, 0x27, 0x47, 0x8f, 0x47, 0xa9, 0xd4, 0x85, 0xd4, 0x89,
	0x0b, 0x9c, 0x70, 0xd1, 0xf4, 0xec, 0xa3, 0xd6, 0xe7, 0xfb, 0xc3, 0x78, 0x84, 0xba, 0xf6, 0x87,
	0x84, 0x33, 0xe2, 0x4f, 0xfc, 0x59, 0x2f, 0xee, 0x58, 0xfd, 0x96, 0x61, 0x8c, 0x9a, 0x82, 0x16,
	0x40, 0x9e, 0x58, 0xdb, 0xce, 0x78, 0x88, 0xda, 0xfa, 0x54, 0xac, 0x65, 0x4e, 0x1a, 0xd6, 0xad,
	0x14, 0x26, 0xa8, 0xc3, 0x20, 0xe5, 0x05, 0xcd, 0x49, 0xd3, 0xbd, 0x52, 0x49, 0x3c, 0x41, 0x4f,
	0x19, 0xe8, 0x54, 0xf1, 0xd2, 0x70, 0x29, 0x48, 0xcb, 0xa6, 0x0f, 0xad, 0xd7, 0xcd, 0x3f, 0x5f,
	0x5f, 0x78, 0xd3, 0x6f, 0x3e, 0xea, 0xdb, 0x95, 0xde, 0x53, 0x41, 0x33, 0x28, 0x40, 0x18, 0xbc,
	0x44, 0x9d, 0xc2, 0x2a, 0xe5, 0x76, 0x5b, 0x91, 0x9f, 0xdf, 0xe7, 0x83, 0xea, 0x92, 0x37, 0x8c,
	0x29, 0xd0, 0xfa, 0x93, 0x51, 0x5c, 0x64, 0x71, 0x0d, 0xe2, 0x97, 0xe8, 0x19, 0x65, 0x2c, 0x11,
	0x70, 0x48, 0x4a, 0xc5, 0xf7, 0x3c, 0x87, 0xcc, 0x9d, 0xd0, 0x8d, 0xfb, 0x94, 0xb1, 0x0f, 0x70,
	0xf8, 0x58, 0xdb, 0x38, 0x42, 0xcf, 0xe1, 0x98, 0xe6, 0x3b, 0x06, 0xec, 0x3f, 0xac, 0x49, 0x63,
	0xd2, 0x98, 0xf5, 0x62, 0x5c, 0x47, 0xff, 0x78, 0xbd, 0x7a, 0xf7, 0xe3, 0x1a, 0xf8, 0x97, 0x6b,
	0xe0, 0xff, 0xbe, 0x06, 0xfe, 0xf9, 0x16, 0x78, 0x97, 0x5b, 0xe0, 0xfd, 0xba, 0x05, 0xde, 0x97,
	0x65, 0xc6, 0xcd, 0x66, 0xb7, 0x0e, 0x53, 0x59, 0x44, 0xae, 0x25, 0x03, 0xe9, 0xa6, 0x1a, 0xe7,
	0x75, 0xab, 0xc7, 0xaa, 0x57, 0x73, 0x2a, 0x41, 0xaf, 0xdb, 0xb6, 0x8c, 0x57, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0xe5, 0x0f, 0xb9, 0xfb, 0x01, 0x00, 0x00,
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Decimal) > 0 {
		i -= len(m.Decimal)
		copy(dAtA[i:], m.Decimal)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Decimal)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintToken(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenManagement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenManagement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenManagement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExcludedPrivileges) > 0 {
		for iNdEx := len(m.ExcludedPrivileges) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ExcludedPrivileges[iNdEx])
			copy(dAtA[i:], m.ExcludedPrivileges[iNdEx])
			i = encodeVarintToken(dAtA, i, uint64(len(m.ExcludedPrivileges[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.AddNewPrivilege {
		i--
		if m.AddNewPrivilege {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Manager) > 0 {
		i -= len(m.Manager)
		copy(dAtA[i:], m.Manager)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Manager)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Decimal)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func (m *TokenManagement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Manager)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.AddNewPrivilege {
		n += 2
	}
	if len(m.ExcludedPrivileges) > 0 {
		for _, s := range m.ExcludedPrivileges {
			l = len(s)
			n += 1 + l + sovToken(uint64(l))
		}
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Decimal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func (m *TokenManagement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: TokenManagement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenManagement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manager", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Manager = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddNewPrivilege", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddNewPrivilege = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExcludedPrivileges", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExcludedPrivileges = append(m.ExcludedPrivileges, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
