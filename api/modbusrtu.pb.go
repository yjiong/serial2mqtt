// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modbusrtu.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type ModbusRtuUpdateRequest struct {
	Devid                string   `protobuf:"bytes,1,opt,name=devid,proto3" json:"devid,omitempty"`
	Devaddr              string   `protobuf:"bytes,3,opt,name=devaddr,proto3" json:"devaddr,omitempty"`
	Commif               string   `protobuf:"bytes,4,opt,name=commif,proto3" json:"commif,omitempty"`
	BaudRate             uint32   `protobuf:"varint,5,opt,name=baudRate,proto3" json:"baudRate,omitempty"`
	DataBits             uint32   `protobuf:"varint,6,opt,name=dataBits,proto3" json:"dataBits,omitempty"`
	StopBits             uint32   `protobuf:"varint,7,opt,name=stopBits,proto3" json:"stopBits,omitempty"`
	Parity               string   `protobuf:"bytes,8,opt,name=parity,proto3" json:"parity,omitempty"`
	FunctionCode         uint32   `protobuf:"varint,9,opt,name=functionCode,proto3" json:"functionCode,omitempty"`
	StartingAddress      uint32   `protobuf:"varint,10,opt,name=startingAddress,proto3" json:"startingAddress,omitempty"`
	Quantity             uint32   `protobuf:"varint,11,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Dname                string   `protobuf:"bytes,12,opt,name=dname,proto3" json:"dname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModbusRtuUpdateRequest) Reset()         { *m = ModbusRtuUpdateRequest{} }
func (m *ModbusRtuUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ModbusRtuUpdateRequest) ProtoMessage()    {}
func (*ModbusRtuUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modbusrtu_67506c34c879ff43, []int{0}
}
func (m *ModbusRtuUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModbusRtuUpdateRequest.Unmarshal(m, b)
}
func (m *ModbusRtuUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModbusRtuUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *ModbusRtuUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModbusRtuUpdateRequest.Merge(dst, src)
}
func (m *ModbusRtuUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ModbusRtuUpdateRequest.Size(m)
}
func (m *ModbusRtuUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModbusRtuUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModbusRtuUpdateRequest proto.InternalMessageInfo

func (m *ModbusRtuUpdateRequest) GetDevid() string {
	if m != nil {
		return m.Devid
	}
	return ""
}

func (m *ModbusRtuUpdateRequest) GetDevaddr() string {
	if m != nil {
		return m.Devaddr
	}
	return ""
}

func (m *ModbusRtuUpdateRequest) GetCommif() string {
	if m != nil {
		return m.Commif
	}
	return ""
}

func (m *ModbusRtuUpdateRequest) GetBaudRate() uint32 {
	if m != nil {
		return m.BaudRate
	}
	return 0
}

func (m *ModbusRtuUpdateRequest) GetDataBits() uint32 {
	if m != nil {
		return m.DataBits
	}
	return 0
}

func (m *ModbusRtuUpdateRequest) GetStopBits() uint32 {
	if m != nil {
		return m.StopBits
	}
	return 0
}

func (m *ModbusRtuUpdateRequest) GetParity() string {
	if m != nil {
		return m.Parity
	}
	return ""
}

func (m *ModbusRtuUpdateRequest) GetFunctionCode() uint32 {
	if m != nil {
		return m.FunctionCode
	}
	return 0
}

func (m *ModbusRtuUpdateRequest) GetStartingAddress() uint32 {
	if m != nil {
		return m.StartingAddress
	}
	return 0
}

func (m *ModbusRtuUpdateRequest) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *ModbusRtuUpdateRequest) GetDname() string {
	if m != nil {
		return m.Dname
	}
	return ""
}

type ModbusRtuUpdateResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModbusRtuUpdateResponse) Reset()         { *m = ModbusRtuUpdateResponse{} }
func (m *ModbusRtuUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*ModbusRtuUpdateResponse) ProtoMessage()    {}
func (*ModbusRtuUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modbusrtu_67506c34c879ff43, []int{1}
}
func (m *ModbusRtuUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModbusRtuUpdateResponse.Unmarshal(m, b)
}
func (m *ModbusRtuUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModbusRtuUpdateResponse.Marshal(b, m, deterministic)
}
func (dst *ModbusRtuUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModbusRtuUpdateResponse.Merge(dst, src)
}
func (m *ModbusRtuUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_ModbusRtuUpdateResponse.Size(m)
}
func (m *ModbusRtuUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModbusRtuUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModbusRtuUpdateResponse proto.InternalMessageInfo

func (m *ModbusRtuUpdateResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*ModbusRtuUpdateRequest)(nil), "api.ModbusRtuUpdateRequest")
	proto.RegisterType((*ModbusRtuUpdateResponse)(nil), "api.ModbusRtuUpdateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModbusRtuClient is the client API for ModbusRtu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModbusRtuClient interface {
	// 添加或更新ModbusRtu设备
	ModbusRtuUpdate(ctx context.Context, in *ModbusRtuUpdateRequest, opts ...grpc.CallOption) (*ModbusRtuUpdateResponse, error)
}

type modbusRtuClient struct {
	cc *grpc.ClientConn
}

func NewModbusRtuClient(cc *grpc.ClientConn) ModbusRtuClient {
	return &modbusRtuClient{cc}
}

func (c *modbusRtuClient) ModbusRtuUpdate(ctx context.Context, in *ModbusRtuUpdateRequest, opts ...grpc.CallOption) (*ModbusRtuUpdateResponse, error) {
	out := new(ModbusRtuUpdateResponse)
	err := c.cc.Invoke(ctx, "/api.ModbusRtu/ModbusRtuUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModbusRtuServer is the server API for ModbusRtu service.
type ModbusRtuServer interface {
	// 添加或更新ModbusRtu设备
	ModbusRtuUpdate(context.Context, *ModbusRtuUpdateRequest) (*ModbusRtuUpdateResponse, error)
}

func RegisterModbusRtuServer(s *grpc.Server, srv ModbusRtuServer) {
	s.RegisterService(&_ModbusRtu_serviceDesc, srv)
}

func _ModbusRtu_ModbusRtuUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModbusRtuUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModbusRtuServer).ModbusRtuUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ModbusRtu/ModbusRtuUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModbusRtuServer).ModbusRtuUpdate(ctx, req.(*ModbusRtuUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModbusRtu_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ModbusRtu",
	HandlerType: (*ModbusRtuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModbusRtuUpdate",
			Handler:    _ModbusRtu_ModbusRtuUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modbusrtu.proto",
}

func init() { proto.RegisterFile("modbusrtu.proto", fileDescriptor_modbusrtu_67506c34c879ff43) }

var fileDescriptor_modbusrtu_67506c34c879ff43 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xd9, 0xf6, 0xed, 0xbf, 0xbc, 0x95, 0x42, 0x90, 0x1a, 0xd6, 0x0a, 0x65, 0x4f, 0x8b,
	0x87, 0x16, 0xf5, 0xe6, 0x4d, 0x3d, 0x7b, 0x59, 0xf0, 0x5c, 0xa6, 0x4d, 0x5a, 0x02, 0xdd, 0x24,
	0xdd, 0x4c, 0x0a, 0x5e, 0xf5, 0x23, 0xf8, 0xb5, 0xbc, 0xf9, 0x15, 0xfc, 0x20, 0x92, 0x64, 0xb7,
	0x60, 0xed, 0xf1, 0x37, 0xcf, 0xcc, 0x3e, 0xb3, 0x4f, 0x86, 0x8c, 0x4a, 0xcd, 0x97, 0xce, 0x56,
	0xe8, 0x66, 0xa6, 0xd2, 0xa8, 0x69, 0x1b, 0x8c, 0x4c, 0x27, 0x1b, 0xad, 0x37, 0x5b, 0x31, 0x07,
	0x23, 0xe7, 0xa0, 0x94, 0x46, 0x40, 0xa9, 0x95, 0x8d, 0x2d, 0xd9, 0x67, 0x8b, 0x8c, 0x9f, 0xc3,
	0x58, 0x81, 0xee, 0xc5, 0x70, 0x40, 0x51, 0x88, 0x9d, 0x13, 0x16, 0xe9, 0x39, 0xe9, 0x70, 0xb1,
	0x97, 0x9c, 0x25, 0xd3, 0x24, 0x1f, 0x14, 0x11, 0x28, 0x23, 0x3d, 0x2e, 0xf6, 0xc0, 0x79, 0xc5,
	0xda, 0xa1, 0xde, 0x20, 0x1d, 0x93, 0xee, 0x4a, 0x97, 0xa5, 0x5c, 0xb3, 0x7f, 0x41, 0xa8, 0x89,
	0xa6, 0xa4, 0xbf, 0x04, 0xc7, 0x0b, 0x40, 0xc1, 0x3a, 0xd3, 0x24, 0x3f, 0x2b, 0x0e, 0xec, 0x35,
	0x0e, 0x08, 0x8f, 0x12, 0x2d, 0xeb, 0x46, 0xad, 0x61, 0xaf, 0x59, 0xd4, 0x26, 0x68, 0xbd, 0xa8,
	0x35, 0xec, 0xbd, 0x0c, 0x54, 0x12, 0x5f, 0x59, 0x3f, 0x7a, 0x45, 0xa2, 0x19, 0x19, 0xae, 0x9d,
	0x5a, 0xf9, 0x3f, 0x7c, 0xd2, 0x5c, 0xb0, 0x41, 0x98, 0xfb, 0x55, 0xa3, 0x39, 0x19, 0x59, 0x84,
	0x0a, 0xa5, 0xda, 0x3c, 0x70, 0x5e, 0x09, 0x6b, 0x19, 0x09, 0x6d, 0xc7, 0x65, 0xbf, 0xc1, 0xce,
	0x81, 0x42, 0xef, 0xf3, 0x3f, 0x6e, 0xd0, 0x70, 0x48, 0x47, 0x41, 0x29, 0xd8, 0xb0, 0x4e, 0xc7,
	0x43, 0x76, 0x43, 0x2e, 0xfe, 0xa4, 0x69, 0x8d, 0x56, 0x56, 0xf8, 0x95, 0x0b, 0x61, 0xdd, 0x16,
	0xeb, 0x3c, 0x6b, 0xba, 0x7d, 0x4f, 0xc8, 0xe0, 0x30, 0x43, 0xf7, 0x64, 0x74, 0xf4, 0x01, 0x7a,
	0x39, 0x03, 0x23, 0x67, 0xa7, 0x1f, 0x29, 0x9d, 0x9c, 0x16, 0xa3, 0x67, 0x96, 0xbf, 0x7d, 0x7d,
	0x7f, 0xb4, 0xb2, 0xf4, 0x2a, 0xbc, 0xfe, 0xe1, 0x3c, 0x16, 0xc0, 0xf9, 0x42, 0x57, 0x0b, 0x67,
	0x42, 0xfb, 0x7d, 0x72, 0xbd, 0xec, 0x86, 0x73, 0xb8, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x71, 0x15, 0xde, 0x44, 0x02, 0x00, 0x00,
}