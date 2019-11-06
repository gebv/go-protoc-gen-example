// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/bar/bar.proto

package bar

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RequestInput struct {
	Data                 *SomeModel `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestInput) Reset()         { *m = RequestInput{} }
func (m *RequestInput) String() string { return proto.CompactTextString(m) }
func (*RequestInput) ProtoMessage()    {}
func (*RequestInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_54dcc294bbf6f45a, []int{0}
}
func (m *RequestInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInput.Merge(m, src)
}
func (m *RequestInput) XXX_Size() int {
	return m.Size()
}
func (m *RequestInput) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInput.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInput proto.InternalMessageInfo

func (m *RequestInput) GetData() *SomeModel {
	if m != nil {
		return m.Data
	}
	return nil
}

type RequestOutput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestOutput) Reset()         { *m = RequestOutput{} }
func (m *RequestOutput) String() string { return proto.CompactTextString(m) }
func (*RequestOutput) ProtoMessage()    {}
func (*RequestOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_54dcc294bbf6f45a, []int{1}
}
func (m *RequestOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestOutput.Merge(m, src)
}
func (m *RequestOutput) XXX_Size() int {
	return m.Size()
}
func (m *RequestOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestOutput.DiscardUnknown(m)
}

var xxx_messageInfo_RequestOutput proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RequestInput)(nil), "bar.RequestInput")
	proto.RegisterType((*RequestOutput)(nil), "bar.RequestOutput")
}

func init() { proto.RegisterFile("api/bar/bar.proto", fileDescriptor_54dcc294bbf6f45a) }

var fileDescriptor_54dcc294bbf6f45a = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0xc8, 0xd4,
	0x4f, 0x4a, 0x2c, 0x02, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xe6, 0xa4, 0xc4, 0x22,
	0x29, 0x11, 0x98, 0x78, 0x6e, 0x7e, 0x4a, 0x6a, 0x4e, 0x31, 0x44, 0x4a, 0xc9, 0x88, 0x8b, 0x27,
	0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0xc4, 0x33, 0xaf, 0xa0, 0xb4, 0x44, 0x48, 0x89, 0x8b, 0x25,
	0x25, 0xb1, 0x24, 0x51, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x4f, 0x0f, 0x64, 0x48, 0x70,
	0x7e, 0x6e, 0xaa, 0x2f, 0x48, 0x53, 0x10, 0x58, 0x4e, 0x89, 0x9f, 0x8b, 0x17, 0xaa, 0xc7, 0xbf,
	0xb4, 0xa4, 0xa0, 0xb4, 0xc4, 0xc8, 0x9e, 0x8b, 0xd3, 0x29, 0xb1, 0x28, 0x38, 0xb5, 0xa8, 0x2c,
	0xb5, 0x48, 0xc8, 0x88, 0x8b, 0x1d, 0x2a, 0x2b, 0x24, 0x08, 0xd6, 0x8e, 0x6c, 0xbe, 0x94, 0x10,
	0xb2, 0x10, 0x44, 0xbb, 0x12, 0x83, 0x13, 0xcf, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x98, 0xc4, 0x06, 0x76, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x70,
	0xd8, 0xf7, 0x74, 0xca, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BarServerClient is the client API for BarServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BarServerClient interface {
	Request(ctx context.Context, in *RequestInput, opts ...grpc.CallOption) (*RequestOutput, error)
}

type barServerClient struct {
	cc *grpc.ClientConn
}

func NewBarServerClient(cc *grpc.ClientConn) BarServerClient {
	return &barServerClient{cc}
}

func (c *barServerClient) Request(ctx context.Context, in *RequestInput, opts ...grpc.CallOption) (*RequestOutput, error) {
	out := new(RequestOutput)
	err := c.cc.Invoke(ctx, "/bar.BarServer/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BarServerServer is the server API for BarServer service.
type BarServerServer interface {
	Request(context.Context, *RequestInput) (*RequestOutput, error)
}

// UnimplementedBarServerServer can be embedded to have forward compatible implementations.
type UnimplementedBarServerServer struct {
}

func (*UnimplementedBarServerServer) Request(ctx context.Context, req *RequestInput) (*RequestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}

func RegisterBarServerServer(s *grpc.Server, srv BarServerServer) {
	s.RegisterService(&_BarServer_serviceDesc, srv)
}

func _BarServer_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServerServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bar.BarServer/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServerServer).Request(ctx, req.(*RequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _BarServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bar.BarServer",
	HandlerType: (*BarServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _BarServer_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/bar/bar.proto",
}

func (m *RequestInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBar(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestOutput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestOutput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintBar(dAtA []byte, offset int, v uint64) int {
	offset -= sovBar(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovBar(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBar(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBar(x uint64) (n int) {
	return sovBar(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBar
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
			return fmt.Errorf("proto: RequestInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBar
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
				return ErrInvalidLengthBar
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &SomeModel{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBar
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBar
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestOutput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBar
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
			return fmt.Errorf("proto: RequestOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipBar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBar
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBar
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBar(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBar
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
					return 0, ErrIntOverflowBar
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
					return 0, ErrIntOverflowBar
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
				return 0, ErrInvalidLengthBar
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBar
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBar
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBar        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBar          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBar = fmt.Errorf("proto: unexpected end of group")
)
