// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: realionetwork/asset/priviledges/freeze/messages.proto

package freeze

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type MsgFreezeToken struct {
	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (m *MsgFreezeToken) Reset()         { *m = MsgFreezeToken{} }
func (m *MsgFreezeToken) String() string { return proto.CompactTextString(m) }
func (*MsgFreezeToken) ProtoMessage()    {}
func (*MsgFreezeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_537c9a40c5210f44, []int{0}
}
func (m *MsgFreezeToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFreezeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFreezeToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFreezeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFreezeToken.Merge(m, src)
}
func (m *MsgFreezeToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgFreezeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFreezeToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFreezeToken proto.InternalMessageInfo

func (m *MsgFreezeToken) GetAccounts() []string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type MsgUnfreezeToken struct {
	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (m *MsgUnfreezeToken) Reset()         { *m = MsgUnfreezeToken{} }
func (m *MsgUnfreezeToken) String() string { return proto.CompactTextString(m) }
func (*MsgUnfreezeToken) ProtoMessage()    {}
func (*MsgUnfreezeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_537c9a40c5210f44, []int{1}
}
func (m *MsgUnfreezeToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnfreezeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnfreezeToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnfreezeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnfreezeToken.Merge(m, src)
}
func (m *MsgUnfreezeToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnfreezeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnfreezeToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnfreezeToken proto.InternalMessageInfo

func (m *MsgUnfreezeToken) GetAccounts() []string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgFreezeToken)(nil), "realionetwork.asset.v1.MsgFreezeToken")
	proto.RegisterType((*MsgUnfreezeToken)(nil), "realionetwork.asset.v1.MsgUnfreezeToken")
}

func init() {
	proto.RegisterFile("realionetwork/asset/priviledges/freeze/messages.proto", fileDescriptor_537c9a40c5210f44)
}

var fileDescriptor_537c9a40c5210f44 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2d, 0x4a, 0x4d, 0xcc,
	0xc9, 0xcc, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1,
	0x2f, 0x28, 0xca, 0x2c, 0xcb, 0xcc, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd6, 0x4f, 0x2b, 0x4a, 0x4d,
	0xad, 0x4a, 0xd5, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x43, 0xd1, 0xa6, 0x07, 0xd6, 0xa6, 0x57, 0x66, 0x28, 0x25, 0x99, 0x9c, 0x5f,
	0x9c, 0x9b, 0x5f, 0x1c, 0x0f, 0x56, 0xa5, 0x0f, 0xe1, 0x40, 0xb4, 0x28, 0xb9, 0x71, 0xf1, 0xf9,
	0x16, 0xa7, 0xbb, 0x81, 0x8d, 0x0b, 0xc9, 0xcf, 0x4e, 0xcd, 0x13, 0x32, 0xe1, 0xe2, 0x48, 0x4c,
	0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x74, 0x92, 0xb8, 0xb4,
	0x45, 0x57, 0x04, 0xaa, 0xcb, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x38, 0xb8, 0xa4, 0x28, 0x33,
	0x2f, 0x3d, 0x08, 0xae, 0x52, 0xc9, 0x83, 0x4b, 0xc0, 0xb7, 0x38, 0x3d, 0x34, 0x2f, 0x8d, 0x52,
	0x93, 0x9c, 0x22, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x3e, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xe2, 0xd3, 0x92, 0xd4, 0xe4, 0x0c,
	0x28, 0x53, 0x17, 0x16, 0x58, 0x15, 0x38, 0x83, 0x2b, 0x89, 0x0d, 0xec, 0x67, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xef, 0xa6, 0x54, 0x2b, 0x5f, 0x01, 0x00, 0x00,
}

func (m *MsgFreezeToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFreezeToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFreezeToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintMessages(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnfreezeToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnfreezeToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnfreezeToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintMessages(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgFreezeToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	return n
}

func (m *MsgUnfreezeToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgFreezeToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgFreezeToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFreezeToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgUnfreezeToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgUnfreezeToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnfreezeToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
