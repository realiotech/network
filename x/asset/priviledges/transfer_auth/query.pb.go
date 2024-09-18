// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: realionetwork/asset/priviledges/transfer_auth/query.proto

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

type QueryWhitelistedAddressesRequest struct {
	TokenId string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (m *QueryWhitelistedAddressesRequest) Reset()         { *m = QueryWhitelistedAddressesRequest{} }
func (m *QueryWhitelistedAddressesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWhitelistedAddressesRequest) ProtoMessage()    {}
func (*QueryWhitelistedAddressesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a20f3df665e91f36, []int{0}
}
func (m *QueryWhitelistedAddressesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhitelistedAddressesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhitelistedAddressesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhitelistedAddressesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhitelistedAddressesRequest.Merge(m, src)
}
func (m *QueryWhitelistedAddressesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhitelistedAddressesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhitelistedAddressesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhitelistedAddressesRequest proto.InternalMessageInfo

func (m *QueryWhitelistedAddressesRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

type QueryWhitelistedAddressesResponse struct {
	WhitelistedAddrs []string `protobuf:"bytes,1,rep,name=whitelisted_addrs,json=whitelistedAddrs,proto3" json:"whitelisted_addrs,omitempty"`
}

func (m *QueryWhitelistedAddressesResponse) Reset()         { *m = QueryWhitelistedAddressesResponse{} }
func (m *QueryWhitelistedAddressesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWhitelistedAddressesResponse) ProtoMessage()    {}
func (*QueryWhitelistedAddressesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a20f3df665e91f36, []int{1}
}
func (m *QueryWhitelistedAddressesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhitelistedAddressesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhitelistedAddressesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhitelistedAddressesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhitelistedAddressesResponse.Merge(m, src)
}
func (m *QueryWhitelistedAddressesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhitelistedAddressesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhitelistedAddressesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhitelistedAddressesResponse proto.InternalMessageInfo

func (m *QueryWhitelistedAddressesResponse) GetWhitelistedAddrs() []string {
	if m != nil {
		return m.WhitelistedAddrs
	}
	return nil
}

type QueryIsAddressWhitelistedRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryIsAddressWhitelistedRequest) Reset()         { *m = QueryIsAddressWhitelistedRequest{} }
func (m *QueryIsAddressWhitelistedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIsAddressWhitelistedRequest) ProtoMessage()    {}
func (*QueryIsAddressWhitelistedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a20f3df665e91f36, []int{2}
}
func (m *QueryIsAddressWhitelistedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsAddressWhitelistedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsAddressWhitelistedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsAddressWhitelistedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsAddressWhitelistedRequest.Merge(m, src)
}
func (m *QueryIsAddressWhitelistedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsAddressWhitelistedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsAddressWhitelistedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsAddressWhitelistedRequest proto.InternalMessageInfo

func (m *QueryIsAddressWhitelistedRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryIsAddressWhitelistedRespones struct {
	IsWhitelisted bool `protobuf:"varint,1,opt,name=is_whitelisted,json=isWhitelisted,proto3" json:"is_whitelisted,omitempty"`
}

func (m *QueryIsAddressWhitelistedRespones) Reset()         { *m = QueryIsAddressWhitelistedRespones{} }
func (m *QueryIsAddressWhitelistedRespones) String() string { return proto.CompactTextString(m) }
func (*QueryIsAddressWhitelistedRespones) ProtoMessage()    {}
func (*QueryIsAddressWhitelistedRespones) Descriptor() ([]byte, []int) {
	return fileDescriptor_a20f3df665e91f36, []int{3}
}
func (m *QueryIsAddressWhitelistedRespones) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsAddressWhitelistedRespones) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsAddressWhitelistedRespones.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsAddressWhitelistedRespones) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsAddressWhitelistedRespones.Merge(m, src)
}
func (m *QueryIsAddressWhitelistedRespones) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsAddressWhitelistedRespones) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsAddressWhitelistedRespones.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsAddressWhitelistedRespones proto.InternalMessageInfo

func (m *QueryIsAddressWhitelistedRespones) GetIsWhitelisted() bool {
	if m != nil {
		return m.IsWhitelisted
	}
	return false
}

func init() {
	proto.RegisterType((*QueryWhitelistedAddressesRequest)(nil), "realionetwork.asset.v1.QueryWhitelistedAddressesRequest")
	proto.RegisterType((*QueryWhitelistedAddressesResponse)(nil), "realionetwork.asset.v1.QueryWhitelistedAddressesResponse")
	proto.RegisterType((*QueryIsAddressWhitelistedRequest)(nil), "realionetwork.asset.v1.QueryIsAddressWhitelistedRequest")
	proto.RegisterType((*QueryIsAddressWhitelistedRespones)(nil), "realionetwork.asset.v1.QueryIsAddressWhitelistedRespones")
}

func init() {
	proto.RegisterFile("realionetwork/asset/priviledges/transfer_auth/query.proto", fileDescriptor_a20f3df665e91f36)
}

var fileDescriptor_a20f3df665e91f36 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4b, 0x3a, 0x41,
	0x14, 0xc7, 0x5d, 0x7e, 0xf0, 0x53, 0x07, 0x8a, 0x5a, 0x22, 0xb4, 0xc3, 0x62, 0x0b, 0x81, 0x10,
	0xee, 0x50, 0x9d, 0x3a, 0x74, 0xc8, 0x43, 0x60, 0xa7, 0xda, 0xa0, 0xa0, 0xcb, 0x36, 0x3a, 0x2f,
	0x77, 0x50, 0x77, 0x74, 0xde, 0x5b, 0xad, 0xff, 0xa2, 0x3f, 0xa6, 0x3f, 0xa2, 0xa3, 0x74, 0xea,
	0x18, 0xfa, 0x8f, 0x84, 0xbb, 0x23, 0x29, 0x44, 0xdd, 0xe6, 0xf1, 0xde, 0xe7, 0xcb, 0x67, 0xf8,
	0xb2, 0x53, 0x03, 0xa2, 0xaf, 0x74, 0x02, 0x34, 0xd1, 0xa6, 0xc7, 0x05, 0x22, 0x10, 0x1f, 0x1a,
	0x35, 0x56, 0x7d, 0x90, 0x5d, 0x40, 0x4e, 0x46, 0x24, 0xf8, 0x08, 0x26, 0x12, 0x29, 0xc5, 0x7c,
	0x94, 0x82, 0x79, 0x0e, 0x86, 0x46, 0x93, 0x76, 0x77, 0xd7, 0xd0, 0x20, 0x43, 0x83, 0xf1, 0xd1,
	0x5e, 0xb5, 0xa3, 0x71, 0xa0, 0x31, 0xca, 0xae, 0x78, 0x3e, 0xe4, 0x88, 0x7f, 0xc6, 0x6a, 0xd7,
	0x8b, 0x84, 0xbb, 0x58, 0x11, 0xf4, 0x15, 0x12, 0xc8, 0x73, 0x29, 0x0d, 0x20, 0x02, 0x86, 0x30,
	0x4a, 0x01, 0xc9, 0xad, 0xb2, 0x12, 0xe9, 0x1e, 0x24, 0x91, 0x92, 0x15, 0xa7, 0xe6, 0xd4, 0xcb,
	0x61, 0x31, 0x9b, 0x5b, 0xd2, 0xbf, 0x62, 0xfb, 0xbf, 0xe0, 0x38, 0xd4, 0x09, 0x82, 0x7b, 0xc8,
	0xb6, 0x27, 0xdf, 0xfb, 0x48, 0x48, 0x69, 0xb0, 0xe2, 0xd4, 0xfe, 0xd5, 0xcb, 0xe1, 0xd6, 0x64,
	0x1d, 0x44, 0xff, 0xd6, 0x0a, 0xb5, 0xd0, 0x06, 0xad, 0x44, 0x2f, 0x85, 0x8e, 0x59, 0x51, 0xe4,
	0xcb, 0xdc, 0xa7, 0x59, 0x79, 0x7f, 0x6d, 0xec, 0xd8, 0x7f, 0x59, 0xec, 0x86, 0x8c, 0x4a, 0xba,
	0xe1, 0xf2, 0xd0, 0xbf, 0xb4, 0xa6, 0x3f, 0xe7, 0x2e, 0x4c, 0x01, 0xdd, 0x03, 0xb6, 0xa9, 0x30,
	0x5a, 0x71, 0xca, 0xf2, 0x4b, 0xe1, 0x86, 0x5a, 0x3d, 0x6f, 0x3e, 0xbc, 0xcd, 0x3c, 0x67, 0x3a,
	0xf3, 0x9c, 0xcf, 0x99, 0xe7, 0xbc, 0xcc, 0xbd, 0xc2, 0x74, 0xee, 0x15, 0x3e, 0xe6, 0x5e, 0xe1,
	0xfe, 0xa2, 0xab, 0x28, 0x4e, 0xdb, 0x41, 0x47, 0x0f, 0x78, 0x5e, 0x06, 0x41, 0x27, 0xb6, 0xcf,
	0xc6, 0xb2, 0xd3, 0xa7, 0xbf, 0x5a, 0x6d, 0xff, 0xcf, 0xda, 0x39, 0xf9, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xb2, 0x87, 0x2d, 0x1a, 0x0d, 0x02, 0x00, 0x00,
}

func (m *QueryWhitelistedAddressesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhitelistedAddressesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhitelistedAddressesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryWhitelistedAddressesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhitelistedAddressesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhitelistedAddressesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WhitelistedAddrs) > 0 {
		for iNdEx := len(m.WhitelistedAddrs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WhitelistedAddrs[iNdEx])
			copy(dAtA[i:], m.WhitelistedAddrs[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.WhitelistedAddrs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryIsAddressWhitelistedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsAddressWhitelistedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsAddressWhitelistedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIsAddressWhitelistedRespones) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsAddressWhitelistedRespones) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsAddressWhitelistedRespones) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsWhitelisted {
		i--
		if m.IsWhitelisted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryWhitelistedAddressesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryWhitelistedAddressesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WhitelistedAddrs) > 0 {
		for _, s := range m.WhitelistedAddrs {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryIsAddressWhitelistedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIsAddressWhitelistedRespones) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsWhitelisted {
		n += 2
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryWhitelistedAddressesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryWhitelistedAddressesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhitelistedAddressesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryWhitelistedAddressesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryWhitelistedAddressesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhitelistedAddressesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedAddrs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedAddrs = append(m.WhitelistedAddrs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryIsAddressWhitelistedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIsAddressWhitelistedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsAddressWhitelistedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryIsAddressWhitelistedRespones) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIsAddressWhitelistedRespones: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsAddressWhitelistedRespones: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsWhitelisted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
			m.IsWhitelisted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)