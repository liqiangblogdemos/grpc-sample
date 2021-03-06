// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TestRequest struct {
	FileNames            []string     `protobuf:"bytes,1,rep,name=file_names,json=fileNames,proto3" json:"file_names,omitempty"`
	Extensions           []*Extension `protobuf:"bytes,2,rep,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetFileNames() []string {
	if m != nil {
		return m.FileNames
	}
	return nil
}

func (m *TestRequest) GetExtensions() []*Extension {
	if m != nil {
		return m.Extensions
	}
	return nil
}

type Extension struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *any.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extension) Reset()         { *m = Extension{} }
func (m *Extension) String() string { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()    {}
func (*Extension) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *Extension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extension.Unmarshal(m, b)
}
func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
}
func (m *Extension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extension.Merge(m, src)
}
func (m *Extension) XXX_Size() int {
	return xxx_messageInfo_Extension.Size(m)
}
func (m *Extension) XXX_DiscardUnknown() {
	xxx_messageInfo_Extension.DiscardUnknown(m)
}

var xxx_messageInfo_Extension proto.InternalMessageInfo

func (m *Extension) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Extension) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type TestResponse struct {
	FileProtos           map[string]*descriptor.FileDescriptorProto `protobuf:"bytes,1,rep,name=file_protos,json=fileProtos,proto3" json:"file_protos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastUpdateDate       *timestamp.Timestamp                       `protobuf:"bytes,2,opt,name=last_update_date,json=lastUpdateDate,proto3" json:"last_update_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetFileProtos() map[string]*descriptor.FileDescriptorProto {
	if m != nil {
		return m.FileProtos
	}
	return nil
}

func (m *TestResponse) GetLastUpdateDate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdateDate
	}
	return nil
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "TestRequest")
	proto.RegisterType((*Extension)(nil), "Extension")
	proto.RegisterType((*TestResponse)(nil), "TestResponse")
	proto.RegisterMapType((map[string]*descriptor.FileDescriptorProto)(nil), "TestResponse.FileProtosEntry")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xdf, 0xab, 0xd3, 0x30,
	0x14, 0xa6, 0xdd, 0x14, 0x7b, 0x7a, 0xef, 0x75, 0x04, 0x91, 0x5a, 0x19, 0x96, 0x21, 0x58, 0x7c,
	0xc8, 0xa0, 0x82, 0xc8, 0x1e, 0x04, 0x61, 0xd5, 0x37, 0x19, 0x71, 0x82, 0x6f, 0x23, 0x5b, 0xcf,
	0x46, 0xb0, 0x4d, 0x6b, 0x93, 0x8e, 0xf5, 0x4f, 0xf7, 0x4d, 0x92, 0xae, 0x63, 0x76, 0xdc, 0x97,
	0xd2, 0x7c, 0xbf, 0xf2, 0x9d, 0x9e, 0xc2, 0x3d, 0x9e, 0x78, 0x51, 0xe5, 0x48, 0xab, 0xba, 0xd4,
	0x65, 0x18, 0x1d, 0xca, 0xf2, 0x90, 0xe3, 0xdc, 0x9e, 0xb6, 0xcd, 0x7e, 0x9e, 0xa1, 0xda, 0xd5,
	0xa2, 0xd2, 0x65, 0x7d, 0x56, 0xbc, 0x1a, 0x2a, 0xb8, 0x6c, 0xcf, 0xd4, 0xeb, 0x21, 0x85, 0x45,
	0xa5, 0x7b, 0xf2, 0xcd, 0x90, 0xd4, 0xa2, 0x40, 0xa5, 0x79, 0x51, 0x75, 0x82, 0xd9, 0x2f, 0xf0,
	0xd7, 0xa8, 0x34, 0xc3, 0x3f, 0x0d, 0x2a, 0x4d, 0xa6, 0x00, 0x7b, 0x91, 0xe3, 0x46, 0xf2, 0x02,
	0x55, 0xe0, 0x44, 0xa3, 0xd8, 0x63, 0x9e, 0x41, 0xbe, 0x1b, 0x80, 0xbc, 0x07, 0xc0, 0x93, 0x46,
	0xa9, 0x44, 0x29, 0x55, 0xe0, 0x46, 0xa3, 0xd8, 0x4f, 0x80, 0xa6, 0x3d, 0xc4, 0xae, 0xd8, 0x59,
	0x0a, 0xde, 0x85, 0x20, 0x0f, 0xe0, 0x8a, 0x2c, 0x70, 0x22, 0x27, 0x1e, 0x33, 0x57, 0x64, 0x24,
	0x86, 0x71, 0xc6, 0x35, 0x0f, 0xdc, 0xc8, 0x89, 0xfd, 0xe4, 0x05, 0xed, 0x6a, 0xd2, 0xbe, 0x26,
	0xfd, 0x22, 0x5b, 0x66, 0x15, 0xb3, 0xbf, 0x0e, 0xdc, 0x75, 0x0d, 0x55, 0x55, 0x4a, 0x85, 0xe4,
	0x33, 0xf8, 0xb6, 0xa2, 0xd5, 0x76, 0x1d, 0xfd, 0x64, 0x4a, 0xaf, 0x35, 0xf4, 0xab, 0xc8, 0x71,
	0x65, 0xf9, 0x54, 0xea, 0xba, 0x65, 0x76, 0xa8, 0x0e, 0x20, 0x4b, 0x98, 0xe4, 0x5c, 0xe9, 0x4d,
	0x53, 0x65, 0x5c, 0xe3, 0xc6, 0x3c, 0xce, 0x35, 0xc2, 0x9b, 0x1a, 0xeb, 0xfe, 0x6b, 0xb1, 0x07,
	0xe3, 0xf9, 0x69, 0x2d, 0x4b, 0xae, 0x31, 0xdc, 0xc1, 0xf3, 0xc1, 0x25, 0x64, 0x02, 0xa3, 0xdf,
	0xd8, 0xda, 0x21, 0x3d, 0x66, 0x5e, 0xc9, 0x02, 0x9e, 0x1c, 0x79, 0xde, 0xf4, 0xf9, 0x6f, 0x6f,
	0xf2, 0x4d, 0xc4, 0xf2, 0xb2, 0x6b, 0x1b, 0xc6, 0x3a, 0xcb, 0xc2, 0xfd, 0xe4, 0x24, 0xb2, 0x5b,
	0xce, 0x0f, 0xac, 0x8f, 0x62, 0x87, 0xe4, 0x1d, 0x3c, 0xfb, 0x86, 0xda, 0x78, 0x14, 0xb9, 0xa3,
	0x57, 0x6b, 0x0b, 0xef, 0xff, 0x1b, 0x9f, 0x7c, 0x84, 0xf1, 0x4a, 0xc8, 0x03, 0x79, 0x79, 0x73,
	0x61, 0x6a, 0xfe, 0x8d, 0xf0, 0x11, 0x7c, 0xfb, 0xd4, 0x9e, 0x3f, 0xfc, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x63, 0xf7, 0x51, 0x76, 0x9f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	GetFiles(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) GetFiles(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/TestService/GetFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/TestService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	GetFiles(context.Context, *TestRequest) (*TestResponse, error)
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) GetFiles(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (*UnimplementedTestServiceServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestService/GetFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetFiles(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFiles",
			Handler:    _TestService_GetFiles_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
