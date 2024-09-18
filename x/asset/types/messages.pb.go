// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: realionetwork/asset/priviledges/transfer_auth/messages.proto

package types

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

type ActionType int32

const (
	ActionType_ACTION_TYPE_UNSPECIFIED ActionType = 0
	// ACTION_TYPE_ADD defines the type for adding addresses.
	ActionType_ACTION_TYPE_ADD ActionType = 1
	// ACTION_TYPE_REMOVE defines the type for remove addresses.
	ActionType_ACTION_TYPE_REMOVE ActionType = 2
)

var ActionType_name = map[int32]string{
	0: "ACTION_TYPE_UNSPECIFIED",
	1: "ACTION_TYPE_ADD",
	2: "ACTION_TYPE_REMOVE",
}

var ActionType_value = map[string]int32{
	"ACTION_TYPE_UNSPECIFIED": 0,
	"ACTION_TYPE_ADD":         1,
	"ACTION_TYPE_REMOVE":      2,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_815f42c56d94596e, []int{0}
}

type MsgUpdateAllowList struct {
	Address    string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	TokenId    string     `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	ActionType ActionType `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3,enum=realionetwork.asset.v1.ActionType" json:"action_type,omitempty"`
}

func (m *MsgUpdateAllowList) Reset()         { *m = MsgUpdateAllowList{} }
func (m *MsgUpdateAllowList) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAllowList) ProtoMessage()    {}
func (*MsgUpdateAllowList) Descriptor() ([]byte, []int) {
	return fileDescriptor_815f42c56d94596e, []int{0}
}
func (m *MsgUpdateAllowList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAllowList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAllowList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAllowList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAllowList.Merge(m, src)
}
func (m *MsgUpdateAllowList) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAllowList) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAllowList.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAllowList proto.InternalMessageInfo

func (m *MsgUpdateAllowList) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgUpdateAllowList) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *MsgUpdateAllowList) GetActionType() ActionType {
	if m != nil {
		return m.ActionType
	}
	return ActionType_ACTION_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("realionetwork.asset.v1.ActionType", ActionType_name, ActionType_value)
	proto.RegisterType((*MsgUpdateAllowList)(nil), "realionetwork.asset.v1.MsgUpdateAllowList")
}

func init() {
	proto.RegisterFile("realionetwork/asset/priviledges/transfer_auth/messages.proto", fileDescriptor_815f42c56d94596e)
}

var fileDescriptor_815f42c56d94596e = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x18, 0xc5, 0x33, 0x5e, 0xb8, 0xde, 0x3b, 0x17, 0x6e, 0x65, 0x5a, 0x6c, 0x6c, 0x21, 0x88, 0x2b,
	0x29, 0x98, 0x50, 0xbb, 0xed, 0x26, 0x6a, 0x0a, 0x01, 0xff, 0x11, 0xff, 0x40, 0xbb, 0x09, 0x63,
	0x32, 0x8d, 0x83, 0x31, 0x13, 0x66, 0x46, 0xad, 0x6f, 0xd1, 0xc7, 0xe8, 0x03, 0xf4, 0x21, 0xba,
	0x94, 0xae, 0xba, 0x2c, 0xfa, 0x22, 0xc5, 0x44, 0x5b, 0x0b, 0xdd, 0x7d, 0xdf, 0x9c, 0xf3, 0x83,
	0x33, 0xdf, 0x81, 0xd7, 0x9c, 0xe0, 0x90, 0xb2, 0x88, 0xc8, 0x05, 0xe3, 0x13, 0x03, 0x0b, 0x41,
	0xa4, 0x11, 0x73, 0x3a, 0xa7, 0x21, 0xf1, 0x03, 0x22, 0x0c, 0xc9, 0x71, 0x24, 0xee, 0x09, 0x77,
	0xf1, 0x4c, 0x8e, 0x8d, 0x29, 0x11, 0x02, 0x07, 0x44, 0xe8, 0x31, 0x67, 0x92, 0xa1, 0xfc, 0x37,
	0x5a, 0x4f, 0x68, 0x7d, 0x7e, 0x79, 0x56, 0xf0, 0x98, 0x98, 0x32, 0xe1, 0x26, 0x2e, 0x23, 0x5d,
	0x52, 0xa4, 0xf4, 0x04, 0x20, 0x6a, 0x89, 0x60, 0x10, 0xfb, 0x58, 0x12, 0x33, 0x0c, 0xd9, 0xa2,
	0x49, 0x85, 0x44, 0x55, 0x98, 0xc5, 0xbe, 0xcf, 0x89, 0x10, 0x2a, 0x28, 0x82, 0xf2, 0xdf, 0x9a,
	0xfa, 0xfa, 0x5c, 0x39, 0xd9, 0x91, 0x66, 0xaa, 0xf4, 0x24, 0xa7, 0x51, 0xe0, 0xec, 0x8d, 0xa8,
	0x00, 0xff, 0x48, 0x36, 0x21, 0x91, 0x4b, 0x7d, 0x35, 0xb3, 0x85, 0x9c, 0x6c, 0xb2, 0xdb, 0x3e,
	0xaa, 0xc3, 0x7f, 0xd8, 0x93, 0x94, 0x45, 0xae, 0x5c, 0xc6, 0x44, 0xfd, 0x55, 0x04, 0xe5, 0xff,
	0xd5, 0x92, 0xfe, 0x73, 0x5c, 0xdd, 0x4c, 0xac, 0xfd, 0x65, 0x4c, 0x1c, 0x88, 0x3f, 0xe7, 0x8b,
	0x21, 0x84, 0x5f, 0x0a, 0x3a, 0x87, 0xa7, 0x66, 0xbd, 0x6f, 0x77, 0xda, 0x6e, 0xff, 0xb6, 0x6b,
	0xb9, 0x83, 0x76, 0xaf, 0x6b, 0xd5, 0xed, 0x1b, 0xdb, 0x6a, 0xe4, 0x14, 0x74, 0x0c, 0x8f, 0x0e,
	0x45, 0xb3, 0xd1, 0xc8, 0x01, 0x94, 0x87, 0xe8, 0xf0, 0xd1, 0xb1, 0x5a, 0x9d, 0xa1, 0x95, 0xcb,
	0xd4, 0x9a, 0x2f, 0x6b, 0x0d, 0xac, 0xd6, 0x1a, 0x78, 0x5f, 0x6b, 0xe0, 0x71, 0xa3, 0x29, 0xab,
	0x8d, 0xa6, 0xbc, 0x6d, 0x34, 0xe5, 0xae, 0x1a, 0x50, 0x39, 0x9e, 0x8d, 0x74, 0x8f, 0x4d, 0x8d,
	0x34, 0xab, 0x24, 0xde, 0x78, 0x37, 0x56, 0xf6, 0x25, 0x3d, 0xec, 0x6a, 0xda, 0x7e, 0x4d, 0x8c,
	0x7e, 0x27, 0x77, 0xbd, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x19, 0x3a, 0x40, 0xca, 0x01,
	0x00, 0x00,
}

func (m *MsgUpdateAllowList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAllowList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAllowList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ActionType != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.ActionType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *MsgUpdateAllowList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.ActionType != 0 {
		n += 1 + sovMessages(uint64(m.ActionType))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateAllowList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateAllowList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAllowList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
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
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionType", wireType)
			}
			m.ActionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActionType |= ActionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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