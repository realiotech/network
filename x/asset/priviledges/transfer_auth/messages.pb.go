// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: realionetwork/asset/priviledges/transfer_auth/messages.proto

package transfer_auth

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
	Addresses  []string   `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
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

func (m *MsgUpdateAllowList) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
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
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x6a, 0xea, 0x40,
	0x1c, 0xc5, 0x33, 0x0a, 0xf7, 0x5e, 0xe7, 0xc2, 0xbd, 0x32, 0xf7, 0x62, 0x63, 0x0b, 0x41, 0x5c,
	0x49, 0xc1, 0x84, 0xb6, 0xd0, 0x55, 0x37, 0x51, 0x23, 0x04, 0xea, 0x07, 0xf1, 0x03, 0xda, 0x4d,
	0x3a, 0x26, 0xd3, 0x38, 0x18, 0x33, 0x61, 0x66, 0xd4, 0xfa, 0x16, 0x7d, 0x90, 0x2e, 0xfb, 0x10,
	0x5d, 0x4a, 0x57, 0x5d, 0x16, 0x7d, 0x91, 0x42, 0x62, 0xab, 0x85, 0x42, 0x77, 0xff, 0x33, 0xe7,
	0xfc, 0xe0, 0x30, 0x07, 0x5e, 0x70, 0x82, 0x43, 0xca, 0x22, 0x22, 0x17, 0x8c, 0x4f, 0x0c, 0x2c,
	0x04, 0x91, 0x46, 0xcc, 0xe9, 0x9c, 0x86, 0xc4, 0x0f, 0x88, 0x30, 0x24, 0xc7, 0x91, 0xb8, 0x25,
	0xdc, 0xc5, 0x33, 0x39, 0x36, 0xa6, 0x44, 0x08, 0x1c, 0x10, 0xa1, 0xc7, 0x9c, 0x49, 0x86, 0x0a,
	0x9f, 0x68, 0x3d, 0xa1, 0xf5, 0xf9, 0xc9, 0x61, 0xd1, 0x63, 0x62, 0xca, 0x84, 0x9b, 0xa4, 0x8c,
	0x54, 0xa4, 0x48, 0xf9, 0x01, 0x40, 0xd4, 0x12, 0xc1, 0x20, 0xf6, 0xb1, 0x24, 0x66, 0x18, 0xb2,
	0xc5, 0x25, 0x15, 0x12, 0x9d, 0xc3, 0x1c, 0xf6, 0x7d, 0x4e, 0x84, 0x20, 0x42, 0x05, 0xa5, 0x6c,
	0x25, 0x57, 0x53, 0x9f, 0x1f, 0xab, 0xff, 0xb7, 0xac, 0x99, 0x7a, 0x3d, 0xc9, 0x69, 0x14, 0x38,
	0xbb, 0x28, 0x2a, 0xc2, 0x5f, 0x92, 0x4d, 0x48, 0xe4, 0x52, 0x5f, 0xcd, 0x94, 0x40, 0x25, 0xe7,
	0xfc, 0x4c, 0xb4, 0xed, 0xa3, 0x3a, 0xfc, 0x8d, 0x3d, 0x49, 0x59, 0xe4, 0xca, 0x65, 0x4c, 0xd4,
	0x6c, 0x09, 0x54, 0xfe, 0x9c, 0x96, 0xf5, 0xaf, 0x2b, 0xeb, 0x66, 0x12, 0xed, 0x2f, 0x63, 0xe2,
	0x40, 0xfc, 0x71, 0x1f, 0x0f, 0x21, 0xdc, 0x39, 0xe8, 0x08, 0x1e, 0x98, 0xf5, 0xbe, 0xdd, 0x69,
	0xbb, 0xfd, 0xab, 0xae, 0xe5, 0x0e, 0xda, 0xbd, 0xae, 0x55, 0xb7, 0x9b, 0xb6, 0xd5, 0xc8, 0x2b,
	0xe8, 0x1f, 0xfc, 0xbb, 0x6f, 0x9a, 0x8d, 0x46, 0x1e, 0xa0, 0x02, 0x44, 0xfb, 0x8f, 0x8e, 0xd5,
	0xea, 0x0c, 0xad, 0x7c, 0xa6, 0x76, 0xf3, 0xb4, 0xd6, 0xc0, 0x6a, 0xad, 0x81, 0xd7, 0xb5, 0x06,
	0xee, 0x37, 0x9a, 0xb2, 0xda, 0x68, 0xca, 0xcb, 0x46, 0x53, 0xae, 0x9b, 0x01, 0x95, 0xe3, 0xd9,
	0x48, 0xf7, 0xd8, 0xd4, 0x48, 0xbb, 0x4a, 0xe2, 0x8d, 0xb7, 0x67, 0xf5, 0x7d, 0xa8, 0xbb, 0xef,
	0xa6, 0x1a, 0xfd, 0x48, 0xfe, 0xfb, 0xec, 0x2d, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x4d, 0x31, 0x9d,
	0xe2, 0x01, 0x00, 0x00,
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
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintMessages(dAtA, i, uint64(len(m.Addresses[iNdEx])))
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
func (m *MsgUpdateAllowList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
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
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
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
