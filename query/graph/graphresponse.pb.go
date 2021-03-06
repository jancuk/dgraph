// Code generated by protoc-gen-go.
// source: graphresponse.proto
// DO NOT EDIT!

/*
Package graph is a generated protocol buffer package.

It is generated from these files:
	graphresponse.proto

It has these top-level messages:
	Request
	Latency
	Property
	Node
	Response
*/
package graph

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Request struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Latency struct {
	Parsing    string `protobuf:"bytes,1,opt,name=parsing" json:"parsing,omitempty"`
	Processing string `protobuf:"bytes,2,opt,name=processing" json:"processing,omitempty"`
	Pb         string `protobuf:"bytes,3,opt,name=pb" json:"pb,omitempty"`
}

func (m *Latency) Reset()                    { *m = Latency{} }
func (m *Latency) String() string            { return proto.CompactTextString(m) }
func (*Latency) ProtoMessage()               {}
func (*Latency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Property struct {
	Prop string `protobuf:"bytes,1,opt,name=prop" json:"prop,omitempty"`
	Val  []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (m *Property) Reset()                    { *m = Property{} }
func (m *Property) String() string            { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()               {}
func (*Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Node struct {
	Uid        uint64      `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Xid        string      `protobuf:"bytes,2,opt,name=xid" json:"xid,omitempty"`
	Attribute  string      `protobuf:"bytes,3,opt,name=attribute" json:"attribute,omitempty"`
	Properties []*Property `protobuf:"bytes,4,rep,name=properties" json:"properties,omitempty"`
	Children   []*Node     `protobuf:"bytes,5,rep,name=children" json:"children,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Node) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Node) GetChildren() []*Node {
	if m != nil {
		return m.Children
	}
	return nil
}

type Response struct {
	N *Node    `protobuf:"bytes,1,opt,name=n" json:"n,omitempty"`
	L *Latency `protobuf:"bytes,2,opt,name=l" json:"l,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Response) GetN() *Node {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *Response) GetL() *Latency {
	if m != nil {
		return m.L
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "graph.Request")
	proto.RegisterType((*Latency)(nil), "graph.Latency")
	proto.RegisterType((*Property)(nil), "graph.Property")
	proto.RegisterType((*Node)(nil), "graph.Node")
	proto.RegisterType((*Response)(nil), "graph.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Dgraph service

type DgraphClient interface {
	Query(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type dgraphClient struct {
	cc *grpc.ClientConn
}

func NewDgraphClient(cc *grpc.ClientConn) DgraphClient {
	return &dgraphClient{cc}
}

func (c *dgraphClient) Query(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/graph.Dgraph/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dgraph service

type DgraphServer interface {
	Query(context.Context, *Request) (*Response, error)
}

func RegisterDgraphServer(s *grpc.Server, srv DgraphServer) {
	s.RegisterService(&_Dgraph_serviceDesc, srv)
}

func _Dgraph_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.Dgraph/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphServer).Query(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dgraph_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graph.Dgraph",
	HandlerType: (*DgraphServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Dgraph_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x51, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0xc4, 0x6d, 0xd2, 0xc7, 0x16, 0xb5, 0x68, 0xe1, 0x10, 0x10, 0x82, 0xca, 0x17, 0x2a, 0x0e,
	0x05, 0x15, 0xfe, 0x00, 0x8e, 0x08, 0x81, 0xf9, 0x82, 0xb4, 0xb5, 0xda, 0x48, 0x95, 0x63, 0x6c,
	0x07, 0x91, 0xdf, 0xe1, 0x4b, 0xb1, 0x37, 0x4e, 0xa9, 0xb8, 0x8d, 0x67, 0x46, 0x3b, 0xe3, 0x5d,
	0x38, 0xdd, 0x98, 0x5c, 0x6f, 0x8d, 0xb4, 0xba, 0x54, 0x56, 0xce, 0xb5, 0x29, 0x5d, 0x89, 0x29,
	0x91, 0xfc, 0x1a, 0xfa, 0x42, 0x7e, 0x56, 0xd2, 0x3a, 0x3c, 0x83, 0xd4, 0x03, 0x53, 0x67, 0x6c,
	0xca, 0x66, 0x43, 0xd1, 0x3c, 0xf8, 0x07, 0xf4, 0x5f, 0x72, 0x27, 0xd5, 0xaa, 0xc6, 0x0c, 0xfa,
	0x3a, 0x37, 0xb6, 0x50, 0x9b, 0x68, 0x69, 0x9f, 0x78, 0x05, 0xe0, 0xa7, 0xae, 0xa4, 0x25, 0xb1,
	0x43, 0xe2, 0x01, 0x83, 0x63, 0xe8, 0xe8, 0x65, 0xd6, 0x25, 0xde, 0x23, 0x7e, 0x0f, 0x83, 0x37,
	0x53, 0x6a, 0x69, 0x5c, 0x8d, 0x08, 0x89, 0x77, 0xea, 0x38, 0x92, 0x30, 0x9e, 0x40, 0xf7, 0x2b,
	0xdf, 0xd1, 0xa0, 0x63, 0x11, 0x20, 0xff, 0x61, 0x90, 0xbc, 0x96, 0x6b, 0x19, 0xa4, 0xaa, 0x58,
	0x93, 0x3b, 0x11, 0x01, 0x06, 0xe6, 0xdb, 0x33, 0x4d, 0x6a, 0x80, 0x78, 0x09, 0xc3, 0xdc, 0x39,
	0x53, 0x2c, 0x2b, 0x27, 0x63, 0xea, 0x1f, 0x81, 0x77, 0x54, 0x36, 0x84, 0x17, 0xd2, 0x66, 0xc9,
	0xb4, 0x3b, 0x1b, 0x2d, 0x26, 0x73, 0x5a, 0xc7, 0xbc, 0x6d, 0x25, 0x0e, 0x2c, 0x78, 0x03, 0x83,
	0xd5, 0xb6, 0xd8, 0xad, 0x8d, 0x54, 0x59, 0x4a, 0xf6, 0x51, 0xb4, 0x87, 0x46, 0x62, 0x2f, 0xf2,
	0x27, 0x18, 0x88, 0xb8, 0x65, 0x3c, 0x07, 0xa6, 0xa8, 0xe5, 0x3f, 0x37, 0x53, 0xbe, 0x1e, 0x6b,
	0xfe, 0x36, 0x5a, 0x8c, 0xa3, 0x14, 0x57, 0x2c, 0xd8, 0x6e, 0xf1, 0x08, 0xbd, 0x67, 0x22, 0xf1,
	0x16, 0xd2, 0xf7, 0x70, 0x03, 0x6c, 0x5d, 0xf1, 0x52, 0x17, 0x93, 0xfd, 0xbb, 0x09, 0xe3, 0x47,
	0xcb, 0x1e, 0x5d, 0xf5, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x60, 0x8c, 0x1b, 0xb0, 0xec, 0x01,
	0x00, 0x00,
}
