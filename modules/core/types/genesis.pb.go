// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/types/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/reapchain/ibc-go/modules/core/02-client/types"
	types1 "github.com/reapchain/ibc-go/modules/core/03-connection/types"
	types2 "github.com/reapchain/ibc-go/modules/core/04-channel/types"
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

// GenesisState defines the ibc module's genesis state.
type GenesisState struct {
	// ICS002 - Clients genesis state
	ClientGenesis types.GenesisState `protobuf:"bytes,1,opt,name=client_genesis,json=clientGenesis,proto3" json:"client_genesis" yaml:"client_genesis"`
	// ICS003 - Connections genesis state
	ConnectionGenesis types1.GenesisState `protobuf:"bytes,2,opt,name=connection_genesis,json=connectionGenesis,proto3" json:"connection_genesis" yaml:"connection_genesis"`
	// ICS004 - Channel genesis state
	ChannelGenesis types2.GenesisState `protobuf:"bytes,3,opt,name=channel_genesis,json=channelGenesis,proto3" json:"channel_genesis" yaml:"channel_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a49c5663e6fc59, []int{0}
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

func (m *GenesisState) GetClientGenesis() types.GenesisState {
	if m != nil {
		return m.ClientGenesis
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetConnectionGenesis() types1.GenesisState {
	if m != nil {
		return m.ConnectionGenesis
	}
	return types1.GenesisState{}
}

func (m *GenesisState) GetChannelGenesis() types2.GenesisState {
	if m != nil {
		return m.ChannelGenesis
	}
	return types2.GenesisState{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.types.v1.GenesisState")
}

func init() { proto.RegisterFile("ibc/core/types/v1/genesis.proto", fileDescriptor_b9a49c5663e6fc59) }

var fileDescriptor_b9a49c5663e6fc59 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x93, 0x5e, 0xe9, 0x0e, 0x01, 0x8a, 0x1a, 0x01, 0x82, 0x4a, 0xb8, 0x6d, 0xd4, 0x81,
	0x05, 0x5b, 0xa5, 0x1b, 0x63, 0x17, 0x98, 0xc3, 0xc6, 0x82, 0x12, 0x63, 0x52, 0xa3, 0xc4, 0xa7,
	0xaa, 0xdd, 0x48, 0x7d, 0x0b, 0x1e, 0xab, 0x63, 0x47, 0xc4, 0x50, 0xa1, 0xe4, 0x0d, 0x78, 0x02,
	0xd4, 0xd8, 0x24, 0xa9, 0xbc, 0x45, 0xff, 0xf9, 0xce, 0xff, 0x1d, 0x25, 0xf1, 0x06, 0x3c, 0xa6,
	0x84, 0xc2, 0x92, 0x11, 0xb5, 0x5e, 0x30, 0x49, 0xf2, 0x09, 0x49, 0x98, 0x60, 0x92, 0x4b, 0xbc,
	0x58, 0x82, 0x02, 0xbf, 0xc7, 0x63, 0x8a, 0xf7, 0x00, 0xae, 0x00, 0x9c, 0x4f, 0xfa, 0x67, 0x09,
	0x24, 0x50, 0x4d, 0xc9, 0xfe, 0x49, 0x83, 0xfd, 0x61, 0xdd, 0x44, 0x53, 0xce, 0x84, 0xb2, 0xaa,
	0xfa, 0xe3, 0x86, 0x00, 0x21, 0x18, 0x55, 0x1c, 0x84, 0x4d, 0x8d, 0x1a, 0x6a, 0x1e, 0x09, 0xc1,
	0x52, 0x0b, 0x09, 0xbe, 0x3a, 0xde, 0xf1, 0x83, 0x4e, 0x9e, 0x54, 0xa4, 0x98, 0xff, 0xe6, 0x75,
	0xb5, 0xf4, 0xc5, 0x80, 0x97, 0xee, 0xd0, 0xbd, 0x39, 0xba, 0x1b, 0xe2, 0xfa, 0x7a, 0x3d, 0xc7,
	0xf9, 0x04, 0xb7, 0x37, 0x67, 0xd7, 0x9b, 0xdd, 0xc0, 0xf9, 0xd9, 0x0d, 0xce, 0xd7, 0x51, 0x96,
	0xde, 0x07, 0x87, 0x2d, 0x41, 0x78, 0xa2, 0x03, 0xb3, 0xe2, 0xe7, 0x9e, 0xdf, 0x9c, 0x5e, 0xbb,
	0x3a, 0x95, 0x6b, 0xdc, 0x72, 0xd5, 0x8c, 0xe5, 0x1b, 0x19, 0xdf, 0x95, 0xf1, 0x59, 0x6d, 0x41,
	0xd8, 0x6b, 0xc2, 0x3f, 0xef, 0xbb, 0x77, 0x6a, 0x5e, 0x46, 0x2d, 0xfd, 0x57, 0x49, 0x47, 0x2d,
	0xa9, 0x06, 0x2c, 0x23, 0x32, 0xc6, 0x0b, 0x63, 0x3c, 0xec, 0x09, 0xc2, 0xae, 0x49, 0xcc, 0xd2,
	0xec, 0x71, 0x53, 0x20, 0x77, 0x5b, 0x20, 0xf7, 0xbb, 0x40, 0xee, 0x47, 0x89, 0x9c, 0x6d, 0x89,
	0x9c, 0xcf, 0x12, 0x39, 0xcf, 0x38, 0xe1, 0x6a, 0xbe, 0x8a, 0x31, 0x85, 0x8c, 0x50, 0x90, 0x19,
	0x48, 0xc2, 0x63, 0x7a, 0x9b, 0x00, 0xc9, 0xa7, 0x24, 0x83, 0xd7, 0x55, 0xca, 0x64, 0xeb, 0x5f,
	0x8a, 0xff, 0x57, 0x5f, 0x6b, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x81, 0x4d, 0x5d, 0x1d, 0x64,
	0x02, 0x00, 0x00,
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
	{
		size, err := m.ChannelGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ConnectionGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ClientGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.ClientGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ConnectionGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ChannelGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClientGenesis", wireType)
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
			if err := m.ClientGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionGenesis", wireType)
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
			if err := m.ConnectionGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelGenesis", wireType)
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
			if err := m.ChannelGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
