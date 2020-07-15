// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/worker.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StartExtractingLinksRequest struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartExtractingLinksRequest) Reset()         { *m = StartExtractingLinksRequest{} }
func (m *StartExtractingLinksRequest) String() string { return proto.CompactTextString(m) }
func (*StartExtractingLinksRequest) ProtoMessage()    {}
func (*StartExtractingLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a01179bc4c672ae, []int{0}
}

func (m *StartExtractingLinksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartExtractingLinksRequest.Unmarshal(m, b)
}
func (m *StartExtractingLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartExtractingLinksRequest.Marshal(b, m, deterministic)
}
func (m *StartExtractingLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartExtractingLinksRequest.Merge(m, src)
}
func (m *StartExtractingLinksRequest) XXX_Size() int {
	return xxx_messageInfo_StartExtractingLinksRequest.Size(m)
}
func (m *StartExtractingLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartExtractingLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartExtractingLinksRequest proto.InternalMessageInfo

func (m *StartExtractingLinksRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *StartExtractingLinksRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*StartExtractingLinksRequest)(nil), "proto.StartExtractingLinksRequest")
}

func init() { proto.RegisterFile("proto/worker.proto", fileDescriptor_6a01179bc4c672ae) }

var fileDescriptor_6a01179bc4c672ae = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xcf, 0x2f, 0xca, 0x4e, 0x2d, 0xd2, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x94,
	0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x9a, 0x5b,
	0x50, 0x52, 0x09, 0x51, 0xa3, 0xe4, 0xce, 0x25, 0x1d, 0x5c, 0x92, 0x58, 0x54, 0xe2, 0x5a, 0x51,
	0x52, 0x94, 0x98, 0x5c, 0x92, 0x99, 0x97, 0xee, 0x93, 0x99, 0x97, 0x5d, 0x1c, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc6, 0xc5, 0x56, 0x9c, 0x9a, 0x5c, 0x94, 0x5a, 0x22, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0x09, 0x70, 0x31, 0x97, 0x16, 0xe5, 0x48, 0x30, 0x81,
	0x05, 0x41, 0x4c, 0xa3, 0x38, 0x2e, 0xb6, 0x70, 0xb0, 0xe5, 0x42, 0x21, 0x5c, 0x22, 0xd8, 0x8c,
	0x14, 0x52, 0x82, 0x58, 0xa9, 0x87, 0xc7, 0x3e, 0x29, 0x31, 0x3d, 0x88, 0x63, 0xf5, 0x60, 0x8e,
	0xd5, 0x73, 0x05, 0x39, 0x56, 0x89, 0x21, 0x89, 0x0d, 0x2c, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x26, 0x1b, 0x5c, 0x6c, 0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	StartExtractingLinks(ctx context.Context, in *StartExtractingLinksRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) StartExtractingLinks(ctx context.Context, in *StartExtractingLinksRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Worker/StartExtractingLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	StartExtractingLinks(context.Context, *StartExtractingLinksRequest) (*empty.Empty, error)
}

// UnimplementedWorkerServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (*UnimplementedWorkerServer) StartExtractingLinks(ctx context.Context, req *StartExtractingLinksRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartExtractingLinks not implemented")
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_StartExtractingLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartExtractingLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).StartExtractingLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Worker/StartExtractingLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).StartExtractingLinks(ctx, req.(*StartExtractingLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartExtractingLinks",
			Handler:    _Worker_StartExtractingLinks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/worker.proto",
}
