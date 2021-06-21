// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package echo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "github.com/fabric-creed/grpc"
	codes "github.com/fabric-creed/grpc/codes"
	status "github.com/fabric-creed/grpc/status"
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

// EchoRequest is the request for echo.
type EchoRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// EchoResponse is the response for echo.
type EchoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "grpc.examples.echo.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "grpc.examples.echo.EchoResponse")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0x87, 0x3d, 0x11, 0xa5, 0x4f, 0xa7, 0xb8, 0x94, 0x2e, 0x96, 0x5b, 0xbc, 0x29, 0x29, 0x16,
	0xff, 0x81, 0x8a, 0xbb, 0xb4, 0xb8, 0x88, 0x4b, 0x3c, 0x7f, 0xa6, 0x81, 0x5c, 0xde, 0xf9, 0x92,
	0x8a, 0xfe, 0xed, 0x2e, 0x92, 0x2b, 0x05, 0x41, 0xba, 0xd5, 0x2d, 0x8f, 0x7c, 0xef, 0xfb, 0x96,
	0x47, 0x84, 0x76, 0xcd, 0xba, 0x17, 0xce, 0xac, 0x94, 0x93, 0xbe, 0xd5, 0xf8, 0xb4, 0x5d, 0x1f,
	0x90, 0x74, 0xf9, 0xa9, 0xaf, 0xe9, 0xfc, 0xbe, 0x5d, 0xf3, 0x12, 0xef, 0x1b, 0xa4, 0xac, 0xc6,
	0x74, 0xd6, 0x21, 0x25, 0xeb, 0x30, 0xae, 0xa6, 0x55, 0x33, 0x5a, 0xee, 0xc6, 0xba, 0xa1, 0x8b,
	0x2d, 0x98, 0x7a, 0x8e, 0x09, 0xfb, 0xc9, 0x9b, 0xef, 0x63, 0x3a, 0x29, 0xa8, 0x7a, 0xa0, 0xd1,
	0x63, 0xb4, 0xf2, 0x35, 0x0c, 0x57, 0xfa, 0x6f, 0x5d, 0xff, 0x4a, 0x4f, 0xa6, 0xfb, 0x81, 0x6d,
	0xb2, 0x3e, 0x52, 0xcf, 0x74, 0xb9, 0x82, 0x7c, 0x40, 0x56, 0x59, 0x60, 0x3b, 0x1f, 0xdd, 0xc1,
	0xdc, 0xb3, 0xaa, 0xd8, 0xef, 0x82, 0x47, 0xcc, 0x87, 0xb7, 0x37, 0x95, 0x02, 0x4d, 0x16, 0xfe,
	0xd5, 0x0b, 0xda, 0xec, 0x39, 0xda, 0xf0, 0x1f, 0x91, 0x59, 0xb5, 0xb8, 0x7d, 0x9a, 0x3b, 0x66,
	0x17, 0xa0, 0x1d, 0x07, 0x1b, 0x9d, 0x66, 0x71, 0xa6, 0xac, 0x9a, 0xdd, 0xaa, 0x79, 0x83, 0xcd,
	0x1b, 0x41, 0x32, 0xc3, 0x59, 0x98, 0x62, 0x7a, 0x39, 0x1d, 0xde, 0xf3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x23, 0x14, 0x26, 0x96, 0x30, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/fabric-creed/grpc#ClientConn.NewStream.
type EchoClient interface {
	// UnaryEcho is unary echo.
	UnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// ServerStreamingEcho is server side streaming.
	ServerStreamingEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (Echo_ServerStreamingEchoClient, error)
	// ClientStreamingEcho is client side streaming.
	ClientStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (Echo_ClientStreamingEchoClient, error)
	// BidirectionalStreamingEcho is bidi streaming.
	BidirectionalStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (Echo_BidirectionalStreamingEchoClient, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) UnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/grpc.examples.echo.Echo/UnaryEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) ServerStreamingEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (Echo_ServerStreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[0], "/grpc.examples.echo.Echo/ServerStreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServerStreamingEchoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Echo_ServerStreamingEchoClient interface {
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServerStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoServerStreamingEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) ClientStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (Echo_ClientStreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[1], "/grpc.examples.echo.Echo/ClientStreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoClientStreamingEchoClient{stream}
	return x, nil
}

type Echo_ClientStreamingEchoClient interface {
	Send(*EchoRequest) error
	CloseAndRecv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoClientStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoClientStreamingEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoClientStreamingEchoClient) CloseAndRecv() (*EchoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) BidirectionalStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (Echo_BidirectionalStreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[2], "/grpc.examples.echo.Echo/BidirectionalStreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoBidirectionalStreamingEchoClient{stream}
	return x, nil
}

type Echo_BidirectionalStreamingEchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoBidirectionalStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoBidirectionalStreamingEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoBidirectionalStreamingEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	// UnaryEcho is unary echo.
	UnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)
	// ServerStreamingEcho is server side streaming.
	ServerStreamingEcho(*EchoRequest, Echo_ServerStreamingEchoServer) error
	// ClientStreamingEcho is client side streaming.
	ClientStreamingEcho(Echo_ClientStreamingEchoServer) error
	// BidirectionalStreamingEcho is bidi streaming.
	BidirectionalStreamingEcho(Echo_BidirectionalStreamingEchoServer) error
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) UnaryEcho(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryEcho not implemented")
}
func (*UnimplementedEchoServer) ServerStreamingEcho(req *EchoRequest, srv Echo_ServerStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingEcho not implemented")
}
func (*UnimplementedEchoServer) ClientStreamingEcho(srv Echo_ClientStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingEcho not implemented")
}
func (*UnimplementedEchoServer) BidirectionalStreamingEcho(srv Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingEcho not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_UnaryEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).UnaryEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.examples.echo.Echo/UnaryEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).UnaryEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_ServerStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServer).ServerStreamingEcho(m, &echoServerStreamingEchoServer{stream})
}

type Echo_ServerStreamingEchoServer interface {
	Send(*EchoResponse) error
	grpc.ServerStream
}

type echoServerStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoServerStreamingEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Echo_ClientStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).ClientStreamingEcho(&echoClientStreamingEchoServer{stream})
}

type Echo_ClientStreamingEchoServer interface {
	SendAndClose(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoClientStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoClientStreamingEchoServer) SendAndClose(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoClientStreamingEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Echo_BidirectionalStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).BidirectionalStreamingEcho(&echoBidirectionalStreamingEchoServer{stream})
}

type Echo_BidirectionalStreamingEchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoBidirectionalStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoBidirectionalStreamingEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoBidirectionalStreamingEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.examples.echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryEcho",
			Handler:    _Echo_UnaryEcho_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingEcho",
			Handler:       _Echo_ServerStreamingEcho_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingEcho",
			Handler:       _Echo_ClientStreamingEcho_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingEcho",
			Handler:       _Echo_BidirectionalStreamingEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}
