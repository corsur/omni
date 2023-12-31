// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethstate/ethstate/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgCreateStorageSlot struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Storage string `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (m *MsgCreateStorageSlot) Reset()         { *m = MsgCreateStorageSlot{} }
func (m *MsgCreateStorageSlot) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStorageSlot) ProtoMessage()    {}
func (*MsgCreateStorageSlot) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f05befdf5e8ee0, []int{0}
}
func (m *MsgCreateStorageSlot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStorageSlot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStorageSlot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStorageSlot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStorageSlot.Merge(m, src)
}
func (m *MsgCreateStorageSlot) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStorageSlot) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStorageSlot.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStorageSlot proto.InternalMessageInfo

func (m *MsgCreateStorageSlot) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateStorageSlot) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgCreateStorageSlot) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type MsgCreateStorageSlotResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateStorageSlotResponse) Reset()         { *m = MsgCreateStorageSlotResponse{} }
func (m *MsgCreateStorageSlotResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStorageSlotResponse) ProtoMessage()    {}
func (*MsgCreateStorageSlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f05befdf5e8ee0, []int{1}
}
func (m *MsgCreateStorageSlotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStorageSlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStorageSlotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStorageSlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStorageSlotResponse.Merge(m, src)
}
func (m *MsgCreateStorageSlotResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStorageSlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStorageSlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStorageSlotResponse proto.InternalMessageInfo

func (m *MsgCreateStorageSlotResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgUpdateStorageSlot struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Storage string `protobuf:"bytes,4,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (m *MsgUpdateStorageSlot) Reset()         { *m = MsgUpdateStorageSlot{} }
func (m *MsgUpdateStorageSlot) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateStorageSlot) ProtoMessage()    {}
func (*MsgUpdateStorageSlot) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f05befdf5e8ee0, []int{2}
}
func (m *MsgUpdateStorageSlot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateStorageSlot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateStorageSlot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateStorageSlot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateStorageSlot.Merge(m, src)
}
func (m *MsgUpdateStorageSlot) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateStorageSlot) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateStorageSlot.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateStorageSlot proto.InternalMessageInfo

func (m *MsgUpdateStorageSlot) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateStorageSlot) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgUpdateStorageSlot) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgUpdateStorageSlot) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type MsgUpdateStorageSlotResponse struct {
}

func (m *MsgUpdateStorageSlotResponse) Reset()         { *m = MsgUpdateStorageSlotResponse{} }
func (m *MsgUpdateStorageSlotResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateStorageSlotResponse) ProtoMessage()    {}
func (*MsgUpdateStorageSlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f05befdf5e8ee0, []int{3}
}
func (m *MsgUpdateStorageSlotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateStorageSlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateStorageSlotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateStorageSlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateStorageSlotResponse.Merge(m, src)
}
func (m *MsgUpdateStorageSlotResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateStorageSlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateStorageSlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateStorageSlotResponse proto.InternalMessageInfo

type MsgDeleteStorageSlot struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteStorageSlot) Reset()         { *m = MsgDeleteStorageSlot{} }
func (m *MsgDeleteStorageSlot) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteStorageSlot) ProtoMessage()    {}
func (*MsgDeleteStorageSlot) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f05befdf5e8ee0, []int{4}
}
func (m *MsgDeleteStorageSlot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteStorageSlot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteStorageSlot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteStorageSlot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteStorageSlot.Merge(m, src)
}
func (m *MsgDeleteStorageSlot) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteStorageSlot) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteStorageSlot.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteStorageSlot proto.InternalMessageInfo

func (m *MsgDeleteStorageSlot) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteStorageSlot) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgDeleteStorageSlotResponse struct {
}

func (m *MsgDeleteStorageSlotResponse) Reset()         { *m = MsgDeleteStorageSlotResponse{} }
func (m *MsgDeleteStorageSlotResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteStorageSlotResponse) ProtoMessage()    {}
func (*MsgDeleteStorageSlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f05befdf5e8ee0, []int{5}
}
func (m *MsgDeleteStorageSlotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteStorageSlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteStorageSlotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteStorageSlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteStorageSlotResponse.Merge(m, src)
}
func (m *MsgDeleteStorageSlotResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteStorageSlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteStorageSlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteStorageSlotResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateStorageSlot)(nil), "ethstate.ethstate.MsgCreateStorageSlot")
	proto.RegisterType((*MsgCreateStorageSlotResponse)(nil), "ethstate.ethstate.MsgCreateStorageSlotResponse")
	proto.RegisterType((*MsgUpdateStorageSlot)(nil), "ethstate.ethstate.MsgUpdateStorageSlot")
	proto.RegisterType((*MsgUpdateStorageSlotResponse)(nil), "ethstate.ethstate.MsgUpdateStorageSlotResponse")
	proto.RegisterType((*MsgDeleteStorageSlot)(nil), "ethstate.ethstate.MsgDeleteStorageSlot")
	proto.RegisterType((*MsgDeleteStorageSlotResponse)(nil), "ethstate.ethstate.MsgDeleteStorageSlotResponse")
}

func init() { proto.RegisterFile("ethstate/ethstate/tx.proto", fileDescriptor_90f05befdf5e8ee0) }

var fileDescriptor_90f05befdf5e8ee0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xf3, 0xa7, 0xbc, 0x2f, 0xee, 0x41, 0x48, 0xf0, 0x10, 0x83, 0x2c, 0x12, 0x04, 0x3d,
	0xa5, 0x60, 0xbf, 0x80, 0xa8, 0xd7, 0x5e, 0x12, 0xbc, 0x78, 0x91, 0xe8, 0x0e, 0x69, 0x21, 0x75,
	0x43, 0x66, 0x0e, 0xf5, 0x5b, 0xf8, 0xb1, 0x3c, 0x16, 0x4f, 0x1e, 0x25, 0xf9, 0x22, 0x92, 0x4d,
	0xb6, 0x54, 0xb6, 0x0b, 0xc5, 0xdb, 0xce, 0x3e, 0xb3, 0xf3, 0x7b, 0x78, 0x96, 0x61, 0x31, 0xd0,
	0x02, 0xa9, 0x20, 0x98, 0x6e, 0x0f, 0xb4, 0x4e, 0xeb, 0x46, 0x92, 0x0c, 0x03, 0x7d, 0x95, 0xea,
	0x43, 0x7c, 0x61, 0xb6, 0x23, 0xc9, 0xa6, 0x28, 0xe1, 0x09, 0x2b, 0x49, 0xc3, 0xc3, 0x44, 0xb0,
	0x93, 0x39, 0x96, 0x77, 0x0d, 0x14, 0x04, 0xf9, 0x20, 0xe7, 0x95, 0xa4, 0x30, 0x62, 0xff, 0x5f,
	0xfa, 0x4b, 0xd9, 0x44, 0xee, 0xb9, 0x7b, 0x75, 0x94, 0xe9, 0xb2, 0x57, 0x0a, 0x21, 0x1a, 0x40,
	0x8c, 0xbc, 0x41, 0x19, 0xcb, 0x5e, 0x19, 0x09, 0x91, 0x3f, 0x28, 0x63, 0x99, 0xa4, 0xec, 0x6c,
	0x1f, 0x25, 0x03, 0xac, 0xe5, 0x2b, 0x42, 0x78, 0xcc, 0xbc, 0xa5, 0x50, 0xa0, 0x49, 0xe6, 0x2d,
	0x45, 0x42, 0xca, 0xd5, 0x43, 0x2d, 0x0e, 0x76, 0x35, 0x4c, 0xf0, 0xf4, 0x84, 0x5d, 0x97, 0xbe,
	0xd5, 0xe5, 0xe4, 0xb7, 0x4b, 0xae, 0x5c, 0x1a, 0x54, 0xed, 0x32, 0xb9, 0x51, 0xae, 0xee, 0xa1,
	0x82, 0x3f, 0xba, 0x1a, 0x09, 0xc6, 0x04, 0x4d, 0xb8, 0xfe, 0xf4, 0x98, 0x3f, 0xc7, 0x32, 0x5c,
	0xb1, 0xc0, 0xfc, 0x92, 0xcb, 0xd4, 0xf8, 0xe4, 0x74, 0x5f, 0xaa, 0xf1, 0xf4, 0xc0, 0xc6, 0x6d,
	0xfc, 0x2b, 0x16, 0x98, 0x59, 0x5b, 0x70, 0x46, 0xa3, 0x0d, 0x67, 0xcd, 0xb1, 0xc7, 0x99, 0x21,
	0x5a, 0x70, 0x46, 0xa3, 0x0d, 0x67, 0x0d, 0xf5, 0x76, 0xf6, 0xd1, 0x72, 0x77, 0xd3, 0x72, 0xf7,
	0xbb, 0xe5, 0xee, 0x7b, 0xc7, 0x9d, 0x4d, 0xc7, 0x9d, 0xaf, 0x8e, 0x3b, 0x8f, 0xa7, 0x40, 0x8b,
	0x5c, 0x6d, 0xc6, 0x7a, 0x67, 0xa7, 0xde, 0x6a, 0xc0, 0xe7, 0x7f, 0x6a, 0x3d, 0x66, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x28, 0x97, 0x07, 0x8e, 0x75, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateStorageSlot(ctx context.Context, in *MsgCreateStorageSlot, opts ...grpc.CallOption) (*MsgCreateStorageSlotResponse, error)
	UpdateStorageSlot(ctx context.Context, in *MsgUpdateStorageSlot, opts ...grpc.CallOption) (*MsgUpdateStorageSlotResponse, error)
	DeleteStorageSlot(ctx context.Context, in *MsgDeleteStorageSlot, opts ...grpc.CallOption) (*MsgDeleteStorageSlotResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateStorageSlot(ctx context.Context, in *MsgCreateStorageSlot, opts ...grpc.CallOption) (*MsgCreateStorageSlotResponse, error) {
	out := new(MsgCreateStorageSlotResponse)
	err := c.cc.Invoke(ctx, "/ethstate.ethstate.Msg/CreateStorageSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateStorageSlot(ctx context.Context, in *MsgUpdateStorageSlot, opts ...grpc.CallOption) (*MsgUpdateStorageSlotResponse, error) {
	out := new(MsgUpdateStorageSlotResponse)
	err := c.cc.Invoke(ctx, "/ethstate.ethstate.Msg/UpdateStorageSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteStorageSlot(ctx context.Context, in *MsgDeleteStorageSlot, opts ...grpc.CallOption) (*MsgDeleteStorageSlotResponse, error) {
	out := new(MsgDeleteStorageSlotResponse)
	err := c.cc.Invoke(ctx, "/ethstate.ethstate.Msg/DeleteStorageSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateStorageSlot(context.Context, *MsgCreateStorageSlot) (*MsgCreateStorageSlotResponse, error)
	UpdateStorageSlot(context.Context, *MsgUpdateStorageSlot) (*MsgUpdateStorageSlotResponse, error)
	DeleteStorageSlot(context.Context, *MsgDeleteStorageSlot) (*MsgDeleteStorageSlotResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateStorageSlot(ctx context.Context, req *MsgCreateStorageSlot) (*MsgCreateStorageSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorageSlot not implemented")
}
func (*UnimplementedMsgServer) UpdateStorageSlot(ctx context.Context, req *MsgUpdateStorageSlot) (*MsgUpdateStorageSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorageSlot not implemented")
}
func (*UnimplementedMsgServer) DeleteStorageSlot(ctx context.Context, req *MsgDeleteStorageSlot) (*MsgDeleteStorageSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorageSlot not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateStorageSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStorageSlot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStorageSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethstate.ethstate.Msg/CreateStorageSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStorageSlot(ctx, req.(*MsgCreateStorageSlot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateStorageSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateStorageSlot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateStorageSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethstate.ethstate.Msg/UpdateStorageSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateStorageSlot(ctx, req.(*MsgUpdateStorageSlot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteStorageSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteStorageSlot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteStorageSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethstate.ethstate.Msg/DeleteStorageSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteStorageSlot(ctx, req.(*MsgDeleteStorageSlot))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethstate.ethstate.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStorageSlot",
			Handler:    _Msg_CreateStorageSlot_Handler,
		},
		{
			MethodName: "UpdateStorageSlot",
			Handler:    _Msg_UpdateStorageSlot_Handler,
		},
		{
			MethodName: "DeleteStorageSlot",
			Handler:    _Msg_DeleteStorageSlot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethstate/ethstate/tx.proto",
}

func (m *MsgCreateStorageSlot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStorageSlot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStorageSlot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Storage) > 0 {
		i -= len(m.Storage)
		copy(dAtA[i:], m.Storage)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Storage)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateStorageSlotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStorageSlotResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStorageSlotResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateStorageSlot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateStorageSlot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateStorageSlot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Storage) > 0 {
		i -= len(m.Storage)
		copy(dAtA[i:], m.Storage)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Storage)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateStorageSlotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateStorageSlotResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateStorageSlotResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteStorageSlot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteStorageSlot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteStorageSlot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteStorageSlotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteStorageSlotResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteStorageSlotResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateStorageSlot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateStorageSlotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgUpdateStorageSlot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateStorageSlotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteStorageSlot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgDeleteStorageSlotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateStorageSlot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStorageSlot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStorageSlot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateStorageSlotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStorageSlotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStorageSlotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateStorageSlot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateStorageSlot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateStorageSlot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateStorageSlotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateStorageSlotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateStorageSlotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeleteStorageSlot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteStorageSlot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteStorageSlot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeleteStorageSlotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteStorageSlotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteStorageSlotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
