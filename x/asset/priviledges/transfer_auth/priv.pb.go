// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: realionetwork/asset/priviledges/transfer_auth/priv.proto

package transfer_auth

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

type AllowAddrs struct {
	Addrs map[string]bool `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *AllowAddrs) Reset()         { *m = AllowAddrs{} }
func (m *AllowAddrs) String() string { return proto.CompactTextString(m) }
func (*AllowAddrs) ProtoMessage()    {}
func (*AllowAddrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f16378177ffba5d1, []int{0}
}
func (m *AllowAddrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowAddrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowAddrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowAddrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowAddrs.Merge(m, src)
}
func (m *AllowAddrs) XXX_Size() int {
	return m.Size()
}
func (m *AllowAddrs) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowAddrs.DiscardUnknown(m)
}

var xxx_messageInfo_AllowAddrs proto.InternalMessageInfo

func (m *AllowAddrs) GetAddrs() map[string]bool {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func init() {
	proto.RegisterType((*AllowAddrs)(nil), "realionetwork.asset.v1.AllowAddrs")
	proto.RegisterMapType((map[string]bool)(nil), "realionetwork.asset.v1.AllowAddrs.AddrsEntry")
}

func init() {
	proto.RegisterFile("realionetwork/asset/priviledges/transfer_auth/priv.proto", fileDescriptor_f16378177ffba5d1)
}

var fileDescriptor_f16378177ffba5d1 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x28, 0x4a, 0x4d, 0xcc,
	0xc9, 0xcc, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1,
	0x2f, 0x28, 0xca, 0x2c, 0xcb, 0xcc, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd6, 0x2f, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0x4b, 0x2d, 0x8a, 0x4f, 0x2c, 0x2d, 0xc9, 0x00, 0xcb, 0xe8, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0x89, 0xa1, 0xe8, 0xd4, 0x03, 0xeb, 0xd4, 0x2b, 0x33, 0x54, 0xea, 0x66, 0xe4, 0xe2,
	0x72, 0xcc, 0xc9, 0xc9, 0x2f, 0x77, 0x4c, 0x49, 0x29, 0x2a, 0x16, 0x72, 0xe6, 0x62, 0x4d, 0x04,
	0x31, 0x24, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x74, 0xf5, 0xb0, 0x6b, 0xd3, 0x43, 0x68, 0xd1,
	0x03, 0x93, 0xae, 0x79, 0x25, 0x45, 0x95, 0x41, 0x10, 0xbd, 0x52, 0x16, 0x5c, 0x5c, 0x08, 0x41,
	0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53,
	0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x23, 0x08,
	0xc2, 0xb1, 0x62, 0xb2, 0x60, 0x74, 0x4a, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0xb7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0x9b,
	0x4a, 0x52, 0x93, 0x33, 0xa0, 0x4c, 0x5d, 0x58, 0x80, 0x54, 0x10, 0x0a, 0x92, 0x24, 0x36, 0x70,
	0x70, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x62, 0xbe, 0x22, 0x6a, 0x4a, 0x01, 0x00, 0x00,
}

func (m *AllowAddrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowAddrs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowAddrs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addrs) > 0 {
		for k := range m.Addrs {
			v := m.Addrs[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintPriv(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintPriv(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPriv(dAtA []byte, offset int, v uint64) int {
	offset -= sovPriv(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllowAddrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addrs) > 0 {
		for k, v := range m.Addrs {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovPriv(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovPriv(uint64(mapEntrySize))
		}
	}
	return n
}

func sovPriv(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPriv(x uint64) (n int) {
	return sovPriv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllowAddrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPriv
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
			return fmt.Errorf("proto: AllowAddrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowAddrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriv
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
				return ErrInvalidLengthPriv
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPriv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Addrs == nil {
				m.Addrs = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPriv
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
							return ErrIntOverflowPriv
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
						return ErrInvalidLengthPriv
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthPriv
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPriv
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPriv(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthPriv
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Addrs[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPriv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPriv
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
func skipPriv(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPriv
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
					return 0, ErrIntOverflowPriv
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
					return 0, ErrIntOverflowPriv
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
				return 0, ErrInvalidLengthPriv
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPriv
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPriv
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPriv        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPriv          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPriv = fmt.Errorf("proto: unexpected end of group")
)
