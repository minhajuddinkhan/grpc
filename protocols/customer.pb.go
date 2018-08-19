// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocols/customer.proto

package customerprotocol

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Id struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_f64da8b4509d7512, []int{0}
}
func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (dst *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(dst, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type CustomerProto struct {
	Id                   *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address              *string  `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Email                *string  `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerProto) Reset()         { *m = CustomerProto{} }
func (m *CustomerProto) String() string { return proto.CompactTextString(m) }
func (*CustomerProto) ProtoMessage()    {}
func (*CustomerProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_f64da8b4509d7512, []int{1}
}
func (m *CustomerProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerProto.Unmarshal(m, b)
}
func (m *CustomerProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerProto.Marshal(b, m, deterministic)
}
func (dst *CustomerProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerProto.Merge(dst, src)
}
func (m *CustomerProto) XXX_Size() int {
	return xxx_messageInfo_CustomerProto.Size(m)
}
func (m *CustomerProto) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerProto.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerProto proto.InternalMessageInfo

func (m *CustomerProto) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *CustomerProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CustomerProto) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *CustomerProto) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

type NothingFancy struct {
	Toldyou              *string  `protobuf:"bytes,1,opt,name=toldyou" json:"toldyou,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NothingFancy) Reset()         { *m = NothingFancy{} }
func (m *NothingFancy) String() string { return proto.CompactTextString(m) }
func (*NothingFancy) ProtoMessage()    {}
func (*NothingFancy) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_f64da8b4509d7512, []int{2}
}
func (m *NothingFancy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NothingFancy.Unmarshal(m, b)
}
func (m *NothingFancy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NothingFancy.Marshal(b, m, deterministic)
}
func (dst *NothingFancy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NothingFancy.Merge(dst, src)
}
func (m *NothingFancy) XXX_Size() int {
	return xxx_messageInfo_NothingFancy.Size(m)
}
func (m *NothingFancy) XXX_DiscardUnknown() {
	xxx_messageInfo_NothingFancy.DiscardUnknown(m)
}

var xxx_messageInfo_NothingFancy proto.InternalMessageInfo

func (m *NothingFancy) GetToldyou() string {
	if m != nil && m.Toldyou != nil {
		return *m.Toldyou
	}
	return ""
}

type CustomersProto struct {
	Customers            []*CustomerProto `protobuf:"bytes,1,rep,name=customers" json:"customers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomersProto) Reset()         { *m = CustomersProto{} }
func (m *CustomersProto) String() string { return proto.CompactTextString(m) }
func (*CustomersProto) ProtoMessage()    {}
func (*CustomersProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_f64da8b4509d7512, []int{3}
}
func (m *CustomersProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomersProto.Unmarshal(m, b)
}
func (m *CustomersProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomersProto.Marshal(b, m, deterministic)
}
func (dst *CustomersProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomersProto.Merge(dst, src)
}
func (m *CustomersProto) XXX_Size() int {
	return xxx_messageInfo_CustomersProto.Size(m)
}
func (m *CustomersProto) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomersProto.DiscardUnknown(m)
}

var xxx_messageInfo_CustomersProto proto.InternalMessageInfo

func (m *CustomersProto) GetCustomers() []*CustomerProto {
	if m != nil {
		return m.Customers
	}
	return nil
}

func init() {
	proto.RegisterType((*Id)(nil), "customerprotocol.Id")
	proto.RegisterType((*CustomerProto)(nil), "customerprotocol.CustomerProto")
	proto.RegisterType((*NothingFancy)(nil), "customerprotocol.NothingFancy")
	proto.RegisterType((*CustomersProto)(nil), "customerprotocol.CustomersProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	GetCustomer(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CustomerProto, error)
	CreateCustomer(ctx context.Context, in *CustomerProto, opts ...grpc.CallOption) (*Id, error)
	GetAllCustomers(ctx context.Context, in *NothingFancy, opts ...grpc.CallOption) (*CustomersProto, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomer(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CustomerProto, error) {
	out := new(CustomerProto)
	err := c.cc.Invoke(ctx, "/customerprotocol.Customer/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CustomerProto, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/customerprotocol.Customer/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetAllCustomers(ctx context.Context, in *NothingFancy, opts ...grpc.CallOption) (*CustomersProto, error) {
	out := new(CustomersProto)
	err := c.cc.Invoke(ctx, "/customerprotocol.Customer/GetAllCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	GetCustomer(context.Context, *Id) (*CustomerProto, error)
	CreateCustomer(context.Context, *CustomerProto) (*Id, error)
	GetAllCustomers(context.Context, *NothingFancy) (*CustomersProto, error)
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerprotocol.Customer/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetCustomer(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerprotocol.Customer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CustomerProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetAllCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NothingFancy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetAllCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerprotocol.Customer/GetAllCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetAllCustomers(ctx, req.(*NothingFancy))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customerprotocol.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _Customer_GetCustomer_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
		{
			MethodName: "GetAllCustomers",
			Handler:    _Customer_GetAllCustomers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocols/customer.proto",
}

func init() { proto.RegisterFile("protocols/customer.proto", fileDescriptor_customer_f64da8b4509d7512) }

var fileDescriptor_customer_f64da8b4509d7512 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xcb, 0x4a, 0xc4, 0x30,
	0x14, 0x25, 0x99, 0x11, 0xed, 0x9d, 0x27, 0xc1, 0x45, 0xe8, 0xc2, 0x29, 0x59, 0x75, 0x55, 0xa1,
	0x7f, 0x20, 0x33, 0x58, 0xba, 0x11, 0xc1, 0x2f, 0x08, 0xcd, 0x45, 0x0b, 0x69, 0x23, 0x49, 0x66,
	0x31, 0x3f, 0xed, 0x37, 0xc8, 0xa4, 0xa6, 0xa8, 0xa3, 0x74, 0x79, 0x4e, 0xce, 0x2b, 0x17, 0xf8,
	0xbb, 0x35, 0xde, 0x34, 0x46, 0xbb, 0xfb, 0xe6, 0xe8, 0xbc, 0xe9, 0xd0, 0x16, 0x81, 0x62, 0xdb,
	0x88, 0xa3, 0x42, 0x6c, 0x81, 0xd6, 0x8a, 0x01, 0xd0, 0x56, 0x71, 0x92, 0xd1, 0x3c, 0x11, 0x35,
	0xac, 0xf6, 0x5f, 0xaa, 0xe7, 0x60, 0x8a, 0x8f, 0x24, 0x4f, 0xd8, 0x12, 0xe6, 0xbd, 0xec, 0x90,
	0xd3, 0x80, 0x36, 0x70, 0x2d, 0x95, 0xb2, 0xe8, 0x1c, 0x9f, 0x05, 0x62, 0x05, 0x57, 0xd8, 0xc9,
	0x56, 0xf3, 0xf9, 0x19, 0x8a, 0x1d, 0x2c, 0x9f, 0x8c, 0x7f, 0x6b, 0xfb, 0xd7, 0x47, 0xd9, 0x37,
	0xa7, 0xb3, 0xde, 0x1b, 0xad, 0x4e, 0xe6, 0x38, 0xc4, 0x89, 0x03, 0xac, 0x63, 0x97, 0x1b, 0xca,
	0x4a, 0x48, 0xe2, 0x46, 0xc7, 0x49, 0x36, 0xcb, 0x17, 0xe5, 0xae, 0xf8, 0xbd, 0xba, 0xf8, 0x31,
	0xb0, 0xfc, 0x20, 0x70, 0x13, 0x19, 0x76, 0x80, 0x45, 0x85, 0x7e, 0x84, 0xb7, 0x97, 0xe6, 0x5a,
	0xa5, 0x53, 0x91, 0xac, 0x82, 0xf5, 0xde, 0xa2, 0xf4, 0x38, 0x06, 0x4d, 0x59, 0xd2, 0x3f, 0x9b,
	0xd8, 0x0b, 0x6c, 0x2a, 0xf4, 0x0f, 0x5a, 0x8f, 0xff, 0x64, 0x77, 0x97, 0xc2, 0xef, 0x57, 0x4a,
	0xb3, 0xff, 0x9b, 0x86, 0x23, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xf3, 0xf2, 0x33, 0xe1,
	0x01, 0x00, 0x00,
}