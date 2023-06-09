// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/routes.proto

package routes

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

type RegisterInput struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" json:"username" form:"username" binding:"required,min=3,max=20"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" json:"email" form:"email" binding:"required,email"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" json:"password" form:"password" binding:"required,min=8,max=20"`
}

func (m *RegisterInput) Reset()         { *m = RegisterInput{} }
func (m *RegisterInput) String() string { return proto.CompactTextString(m) }
func (*RegisterInput) ProtoMessage()    {}
func (*RegisterInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2350d19bac5b124, []int{0}
}
func (m *RegisterInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterInput.Merge(m, src)
}
func (m *RegisterInput) XXX_Size() int {
	return m.Size()
}
func (m *RegisterInput) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterInput.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterInput proto.InternalMessageInfo

func (m *RegisterInput) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterInput) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterInput)(nil), "sso.routes.RegisterInput")
}

func init() { proto.RegisterFile("protos/routes.proto", fileDescriptor_c2350d19bac5b124) }

var fileDescriptor_c2350d19bac5b124 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0xca, 0x2f, 0x2d, 0x49, 0x2d, 0xd6, 0x03, 0xf3, 0x84, 0xb8, 0x8a, 0x8b,
	0xf3, 0xf5, 0x20, 0x22, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9,
	0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0x25, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66,
	0x41, 0xb4, 0x2a, 0xcd, 0x67, 0xe2, 0xe2, 0x0d, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0xf2,
	0xcc, 0x2b, 0x28, 0x2d, 0x11, 0x8a, 0xe7, 0xe2, 0x28, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x72, 0xfe, 0x74, 0x4f, 0xde, 0x3e, 0xab, 0x38, 0x3f,
	0xcf, 0x4a, 0x09, 0x26, 0xa3, 0xa4, 0x90, 0x96, 0x5f, 0x94, 0x8b, 0xcc, 0x4f, 0xca, 0xcc, 0x4b,
	0xc9, 0xcc, 0x4b, 0xb7, 0x52, 0x2a, 0x4a, 0x2d, 0x2c, 0xcd, 0x2c, 0x4a, 0x4d, 0xd1, 0xc9, 0xcd,
	0xcc, 0xb3, 0x35, 0xd6, 0xc9, 0x4d, 0xac, 0xb0, 0x35, 0x32, 0x50, 0x0a, 0x82, 0x1b, 0x2a, 0xe4,
	0xc3, 0xc5, 0x9a, 0x9a, 0x9b, 0x98, 0x99, 0x23, 0xc1, 0x04, 0x36, 0xdd, 0xec, 0xd3, 0x3d, 0x79,
	0x23, 0x88, 0xe9, 0x60, 0x61, 0x98, 0xd1, 0x50, 0x0e, 0xa6, 0xb9, 0x10, 0x89, 0x20, 0x88, 0x21,
	0x20, 0xe7, 0x16, 0x24, 0x16, 0x17, 0x97, 0xe7, 0x17, 0xa5, 0x48, 0x30, 0xa3, 0x3b, 0x17, 0x26,
	0x03, 0x33, 0x13, 0xc1, 0xc7, 0xee, 0x5c, 0x0b, 0x84, 0x73, 0x61, 0x4a, 0x9d, 0xd4, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x17, 0x14, 0xba, 0xd9, 0xe9, 0xd0, 0xb8,
	0x48, 0x62, 0x03, 0x87, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xc1, 0x56, 0x1c, 0xa3,
	0x01, 0x00, 0x00,
}

func (m *RegisterInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintRoutes(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintRoutes(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintRoutes(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRoutes(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoutes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovRoutes(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovRoutes(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRoutes(uint64(l))
	}
	return n
}

func sovRoutes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoutes(x uint64) (n int) {
	return sovRoutes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoutes
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
			return fmt.Errorf("proto: RegisterInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
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
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoutes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
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
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoutes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
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
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoutes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoutes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoutes
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
func skipRoutes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoutes
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
					return 0, ErrIntOverflowRoutes
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
					return 0, ErrIntOverflowRoutes
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
				return 0, ErrInvalidLengthRoutes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoutes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoutes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoutes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoutes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoutes = fmt.Errorf("proto: unexpected end of group")
)
