// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: isolation/usagepb/usage.proto

package usagepb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

/// Resources allows to report basic resource usage statistics for request.
/// It's meant to be used as tags in tracing spans, logged on demand or for billing purposes.
/// We hold all information that otherwise cannot be deduced from tracing.
type Resources struct {
	// Major allocations only, both heap and stack.
	MemoryAllocBytes uint64 `protobuf:"varint,1,opt,name=memory_alloc_bytes,json=memoryAllocBytes,proto3" json:"memoryAllocBytes"`
	// From upsteam perspective.
	NetworkSentBytes     uint64 `protobuf:"varint,2,opt,name=network_sent_bytes,json=networkSentBytes,proto3" json:"networkSentBytes"`
	NetworkRecvBytes     uint64 `protobuf:"varint,3,opt,name=network_recv_bytes,json=networkRecvBytes,proto3" json:"networkRecvBytes"`
	GoRoutinesSpawned    uint64 `protobuf:"varint,4,opt,name=go_routines_spawned,json=goRoutinesSpawned,proto3" json:"goRoutinesSpawned"`
	GoRoutinesMaxSpawned uint64 `protobuf:"varint,5,opt,name=go_routines_max_spawned,json=goRoutinesMaxSpawned,proto3" json:"goRoutinesMaxSpawned"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_d64eded179b13829, []int{0}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(m, src)
}
func (m *Resources) XXX_Size() int {
	return m.Size()
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Resources)(nil), "usagepb.Resources")
}

func init() { proto.RegisterFile("isolation/usagepb/usage.proto", fileDescriptor_d64eded179b13829) }

var fileDescriptor_d64eded179b13829 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xbb, 0x4e, 0x84, 0x40,
	0x14, 0xc6, 0x71, 0x58, 0x57, 0x8d, 0x54, 0x8a, 0x6b, 0x24, 0x26, 0x0e, 0xc6, 0xc4, 0xc4, 0xca,
	0x2d, 0x7c, 0x02, 0x49, 0x2c, 0x8d, 0x09, 0xdb, 0xd9, 0x90, 0x01, 0x4f, 0x26, 0x44, 0x98, 0x43,
	0x66, 0x86, 0xbd, 0xbc, 0x85, 0x8f, 0xb5, 0xe5, 0x96, 0x56, 0x44, 0xa1, 0xa3, 0xf1, 0x15, 0x0c,
	0x30, 0x7b, 0x51, 0xac, 0x20, 0xdf, 0x3f, 0xf3, 0x6b, 0x8e, 0x75, 0x19, 0x4b, 0x4c, 0xa8, 0x8a,
	0x91, 0x8f, 0x73, 0x49, 0x19, 0x64, 0x61, 0xf7, 0xbd, 0xcb, 0x04, 0x2a, 0xb4, 0x0f, 0xf5, 0x78,
	0x31, 0x62, 0xc8, 0xb0, 0xdd, 0xc6, 0xcd, 0x5f, 0x97, 0xaf, 0xbf, 0x07, 0xd6, 0x91, 0x0f, 0x12,
	0x73, 0x11, 0x81, 0xb4, 0x3d, 0xcb, 0x4e, 0x21, 0x45, 0xb1, 0x08, 0x68, 0x92, 0x60, 0x14, 0x84,
	0x0b, 0x05, 0xd2, 0x31, 0xaf, 0xcc, 0xdb, 0xa1, 0x37, 0xaa, 0x0b, 0xf7, 0xb8, 0xab, 0x0f, 0x4d,
	0xf4, 0x9a, 0xe6, 0xf7, 0x96, 0xc6, 0xe0, 0xa0, 0x66, 0x28, 0xde, 0x02, 0x09, 0x5c, 0x69, 0x63,
	0xb0, 0x35, 0x74, 0x9d, 0x00, 0x57, 0xda, 0xf8, 0xbb, 0xec, 0x1a, 0x02, 0xa2, 0xa9, 0x36, 0xf6,
	0x7a, 0x86, 0x0f, 0xd1, 0xf4, 0xb7, 0xb1, 0x59, 0xec, 0x47, 0xeb, 0x94, 0x61, 0x20, 0x30, 0x57,
	0x31, 0x07, 0x19, 0xc8, 0x8c, 0xce, 0x38, 0xbc, 0x3a, 0xc3, 0x16, 0x39, 0xab, 0x0b, 0xf7, 0x84,
	0xa1, 0xaf, 0xeb, 0xa4, 0x8b, 0x7e, 0x7f, 0xb2, 0x9f, 0xad, 0xf3, 0x5d, 0x26, 0xa5, 0xf3, 0x0d,
	0xb5, 0xdf, 0x52, 0x4e, 0x5d, 0xb8, 0xa3, 0xed, 0xbb, 0x27, 0x3a, 0x5f, 0x6b, 0xff, 0xae, 0xde,
	0xcd, 0xf2, 0x8b, 0x18, 0xcb, 0x92, 0x98, 0xab, 0x92, 0x98, 0x9f, 0x25, 0x31, 0xdf, 0x2b, 0x62,
	0xac, 0x2a, 0x62, 0x7c, 0x54, 0xc4, 0x78, 0x59, 0x9f, 0x2b, 0x3c, 0x68, 0xef, 0x73, 0xff, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x88, 0xff, 0xe3, 0xe8, 0xdf, 0x01, 0x00, 0x00,
}

func (m *Resources) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resources) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resources) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GoRoutinesMaxSpawned != 0 {
		i = encodeVarintUsage(dAtA, i, uint64(m.GoRoutinesMaxSpawned))
		i--
		dAtA[i] = 0x28
	}
	if m.GoRoutinesSpawned != 0 {
		i = encodeVarintUsage(dAtA, i, uint64(m.GoRoutinesSpawned))
		i--
		dAtA[i] = 0x20
	}
	if m.NetworkRecvBytes != 0 {
		i = encodeVarintUsage(dAtA, i, uint64(m.NetworkRecvBytes))
		i--
		dAtA[i] = 0x18
	}
	if m.NetworkSentBytes != 0 {
		i = encodeVarintUsage(dAtA, i, uint64(m.NetworkSentBytes))
		i--
		dAtA[i] = 0x10
	}
	if m.MemoryAllocBytes != 0 {
		i = encodeVarintUsage(dAtA, i, uint64(m.MemoryAllocBytes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUsage(dAtA []byte, offset int, v uint64) int {
	offset -= sovUsage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resources) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MemoryAllocBytes != 0 {
		n += 1 + sovUsage(uint64(m.MemoryAllocBytes))
	}
	if m.NetworkSentBytes != 0 {
		n += 1 + sovUsage(uint64(m.NetworkSentBytes))
	}
	if m.NetworkRecvBytes != 0 {
		n += 1 + sovUsage(uint64(m.NetworkRecvBytes))
	}
	if m.GoRoutinesSpawned != 0 {
		n += 1 + sovUsage(uint64(m.GoRoutinesSpawned))
	}
	if m.GoRoutinesMaxSpawned != 0 {
		n += 1 + sovUsage(uint64(m.GoRoutinesMaxSpawned))
	}
	return n
}

func sovUsage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUsage(x uint64) (n int) {
	return sovUsage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resources) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsage
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
			return fmt.Errorf("proto: Resources: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resources: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryAllocBytes", wireType)
			}
			m.MemoryAllocBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryAllocBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkSentBytes", wireType)
			}
			m.NetworkSentBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkSentBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkRecvBytes", wireType)
			}
			m.NetworkRecvBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkRecvBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoRoutinesSpawned", wireType)
			}
			m.GoRoutinesSpawned = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GoRoutinesSpawned |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoRoutinesMaxSpawned", wireType)
			}
			m.GoRoutinesMaxSpawned = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GoRoutinesMaxSpawned |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUsage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUsage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUsage
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
func skipUsage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUsage
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
					return 0, ErrIntOverflowUsage
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
					return 0, ErrIntOverflowUsage
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
				return 0, ErrInvalidLengthUsage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUsage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUsage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUsage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUsage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUsage = fmt.Errorf("proto: unexpected end of group")
)