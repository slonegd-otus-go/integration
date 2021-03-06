// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

package grpc_api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SumArgs struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumArgs) Reset()         { *m = SumArgs{} }
func (m *SumArgs) String() string { return proto.CompactTextString(m) }
func (*SumArgs) ProtoMessage()    {}
func (*SumArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{0}
}

func (m *SumArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumArgs.Unmarshal(m, b)
}
func (m *SumArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumArgs.Marshal(b, m, deterministic)
}
func (m *SumArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumArgs.Merge(m, src)
}
func (m *SumArgs) XXX_Size() int {
	return xxx_messageInfo_SumArgs.Size(m)
}
func (m *SumArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_SumArgs.DiscardUnknown(m)
}

var xxx_messageInfo_SumArgs proto.InternalMessageInfo

func (m *SumArgs) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumArgs) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResult struct {
	C                    int32    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResult) Reset()         { *m = SumResult{} }
func (m *SumResult) String() string { return proto.CompactTextString(m) }
func (*SumResult) ProtoMessage()    {}
func (*SumResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{1}
}

func (m *SumResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResult.Unmarshal(m, b)
}
func (m *SumResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResult.Marshal(b, m, deterministic)
}
func (m *SumResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResult.Merge(m, src)
}
func (m *SumResult) XXX_Size() int {
	return xxx_messageInfo_SumResult.Size(m)
}
func (m *SumResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResult.DiscardUnknown(m)
}

var xxx_messageInfo_SumResult proto.InternalMessageInfo

func (m *SumResult) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

func init() {
	proto.RegisterType((*SumArgs)(nil), "SumArgs")
	proto.RegisterType((*SumResult)(nil), "SumResult")
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor_a2b9900dc883ea68) }

var fileDescriptor_a2b9900dc883ea68 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe5, 0x62, 0x0f, 0x2e, 0xcd, 0x75, 0x2c, 0x4a,
	0x2f, 0x16, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x4c, 0x04,
	0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x49, 0x2e, 0xce, 0xe0, 0xd2, 0xdc, 0xa0, 0xd4,
	0xe2, 0xd2, 0x9c, 0x12, 0x90, 0x54, 0x32, 0x4c, 0x61, 0xb2, 0x91, 0x32, 0x17, 0x8b, 0x73, 0x62,
	0x4e, 0xb2, 0x90, 0x34, 0x17, 0x73, 0x70, 0x69, 0xae, 0x10, 0x87, 0x1e, 0xd4, 0x3c, 0x29, 0x2e,
	0x3d, 0xb8, 0x16, 0x27, 0xae, 0x28, 0x8e, 0xf4, 0xa2, 0x82, 0xe4, 0xf8, 0xc4, 0x82, 0xcc, 0x24,
	0x36, 0xb0, 0xcd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x10, 0x1f, 0xdc, 0x87, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcClient interface {
	Sum(ctx context.Context, in *SumArgs, opts ...grpc.CallOption) (*SumResult, error)
}

type calcClient struct {
	cc *grpc.ClientConn
}

func NewCalcClient(cc *grpc.ClientConn) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Sum(ctx context.Context, in *SumArgs, opts ...grpc.CallOption) (*SumResult, error) {
	out := new(SumResult)
	err := c.cc.Invoke(ctx, "/Calc/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
type CalcServer interface {
	Sum(context.Context, *SumArgs) (*SumResult, error)
}

// UnimplementedCalcServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (*UnimplementedCalcServer) Sum(ctx context.Context, req *SumArgs) (*SumResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterCalcServer(s *grpc.Server, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calc/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Sum(ctx, req.(*SumArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Calc_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}
