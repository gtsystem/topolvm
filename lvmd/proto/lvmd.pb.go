// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lvmd/proto/lvmd.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type LogicalVolume struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeGb               uint64   `protobuf:"varint,2,opt,name=size_gb,json=sizeGb,proto3" json:"size_gb,omitempty"`
	DevMajor             uint32   `protobuf:"varint,3,opt,name=dev_major,json=devMajor,proto3" json:"dev_major,omitempty"`
	DevMinor             uint32   `protobuf:"varint,4,opt,name=dev_minor,json=devMinor,proto3" json:"dev_minor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogicalVolume) Reset()         { *m = LogicalVolume{} }
func (m *LogicalVolume) String() string { return proto.CompactTextString(m) }
func (*LogicalVolume) ProtoMessage()    {}
func (*LogicalVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{1}
}

func (m *LogicalVolume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalVolume.Unmarshal(m, b)
}
func (m *LogicalVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalVolume.Marshal(b, m, deterministic)
}
func (m *LogicalVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalVolume.Merge(m, src)
}
func (m *LogicalVolume) XXX_Size() int {
	return xxx_messageInfo_LogicalVolume.Size(m)
}
func (m *LogicalVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalVolume.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalVolume proto.InternalMessageInfo

func (m *LogicalVolume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicalVolume) GetSizeGb() uint64 {
	if m != nil {
		return m.SizeGb
	}
	return 0
}

func (m *LogicalVolume) GetDevMajor() uint32 {
	if m != nil {
		return m.DevMajor
	}
	return 0
}

func (m *LogicalVolume) GetDevMinor() uint32 {
	if m != nil {
		return m.DevMinor
	}
	return 0
}

type CreateLVRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeGb               uint64   `protobuf:"varint,2,opt,name=size_gb,json=sizeGb,proto3" json:"size_gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLVRequest) Reset()         { *m = CreateLVRequest{} }
func (m *CreateLVRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLVRequest) ProtoMessage()    {}
func (*CreateLVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{2}
}

func (m *CreateLVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLVRequest.Unmarshal(m, b)
}
func (m *CreateLVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLVRequest.Marshal(b, m, deterministic)
}
func (m *CreateLVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLVRequest.Merge(m, src)
}
func (m *CreateLVRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLVRequest.Size(m)
}
func (m *CreateLVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLVRequest proto.InternalMessageInfo

func (m *CreateLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateLVRequest) GetSizeGb() uint64 {
	if m != nil {
		return m.SizeGb
	}
	return 0
}

type CreateLVResponse struct {
	Volume               *LogicalVolume `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateLVResponse) Reset()         { *m = CreateLVResponse{} }
func (m *CreateLVResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLVResponse) ProtoMessage()    {}
func (*CreateLVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{3}
}

func (m *CreateLVResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLVResponse.Unmarshal(m, b)
}
func (m *CreateLVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLVResponse.Marshal(b, m, deterministic)
}
func (m *CreateLVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLVResponse.Merge(m, src)
}
func (m *CreateLVResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLVResponse.Size(m)
}
func (m *CreateLVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLVResponse proto.InternalMessageInfo

func (m *CreateLVResponse) GetVolume() *LogicalVolume {
	if m != nil {
		return m.Volume
	}
	return nil
}

type RemoveLVRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveLVRequest) Reset()         { *m = RemoveLVRequest{} }
func (m *RemoveLVRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveLVRequest) ProtoMessage()    {}
func (*RemoveLVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{4}
}

func (m *RemoveLVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveLVRequest.Unmarshal(m, b)
}
func (m *RemoveLVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveLVRequest.Marshal(b, m, deterministic)
}
func (m *RemoveLVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveLVRequest.Merge(m, src)
}
func (m *RemoveLVRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveLVRequest.Size(m)
}
func (m *RemoveLVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveLVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveLVRequest proto.InternalMessageInfo

func (m *RemoveLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ResizeLVRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeGb               uint64   `protobuf:"varint,2,opt,name=size_gb,json=sizeGb,proto3" json:"size_gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeLVRequest) Reset()         { *m = ResizeLVRequest{} }
func (m *ResizeLVRequest) String() string { return proto.CompactTextString(m) }
func (*ResizeLVRequest) ProtoMessage()    {}
func (*ResizeLVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{5}
}

func (m *ResizeLVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeLVRequest.Unmarshal(m, b)
}
func (m *ResizeLVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeLVRequest.Marshal(b, m, deterministic)
}
func (m *ResizeLVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeLVRequest.Merge(m, src)
}
func (m *ResizeLVRequest) XXX_Size() int {
	return xxx_messageInfo_ResizeLVRequest.Size(m)
}
func (m *ResizeLVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeLVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeLVRequest proto.InternalMessageInfo

func (m *ResizeLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResizeLVRequest) GetSizeGb() uint64 {
	if m != nil {
		return m.SizeGb
	}
	return 0
}

type GetLVListResponse struct {
	Volumes              []*LogicalVolume `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetLVListResponse) Reset()         { *m = GetLVListResponse{} }
func (m *GetLVListResponse) String() string { return proto.CompactTextString(m) }
func (*GetLVListResponse) ProtoMessage()    {}
func (*GetLVListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{6}
}

func (m *GetLVListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLVListResponse.Unmarshal(m, b)
}
func (m *GetLVListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLVListResponse.Marshal(b, m, deterministic)
}
func (m *GetLVListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLVListResponse.Merge(m, src)
}
func (m *GetLVListResponse) XXX_Size() int {
	return xxx_messageInfo_GetLVListResponse.Size(m)
}
func (m *GetLVListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLVListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLVListResponse proto.InternalMessageInfo

func (m *GetLVListResponse) GetVolumes() []*LogicalVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type GetFreeBytesResponse struct {
	FreeBytes            uint64   `protobuf:"varint,1,opt,name=free_bytes,json=freeBytes,proto3" json:"free_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFreeBytesResponse) Reset()         { *m = GetFreeBytesResponse{} }
func (m *GetFreeBytesResponse) String() string { return proto.CompactTextString(m) }
func (*GetFreeBytesResponse) ProtoMessage()    {}
func (*GetFreeBytesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{7}
}

func (m *GetFreeBytesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFreeBytesResponse.Unmarshal(m, b)
}
func (m *GetFreeBytesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFreeBytesResponse.Marshal(b, m, deterministic)
}
func (m *GetFreeBytesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFreeBytesResponse.Merge(m, src)
}
func (m *GetFreeBytesResponse) XXX_Size() int {
	return xxx_messageInfo_GetFreeBytesResponse.Size(m)
}
func (m *GetFreeBytesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFreeBytesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFreeBytesResponse proto.InternalMessageInfo

func (m *GetFreeBytesResponse) GetFreeBytes() uint64 {
	if m != nil {
		return m.FreeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*LogicalVolume)(nil), "proto.LogicalVolume")
	proto.RegisterType((*CreateLVRequest)(nil), "proto.CreateLVRequest")
	proto.RegisterType((*CreateLVResponse)(nil), "proto.CreateLVResponse")
	proto.RegisterType((*RemoveLVRequest)(nil), "proto.RemoveLVRequest")
	proto.RegisterType((*ResizeLVRequest)(nil), "proto.ResizeLVRequest")
	proto.RegisterType((*GetLVListResponse)(nil), "proto.GetLVListResponse")
	proto.RegisterType((*GetFreeBytesResponse)(nil), "proto.GetFreeBytesResponse")
}

func init() { proto.RegisterFile("lvmd/proto/lvmd.proto", fileDescriptor_cb3c510e545f3bbd) }

var fileDescriptor_cb3c510e545f3bbd = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4b, 0x6b, 0xea, 0x40,
	0x14, 0xc7, 0xc9, 0x35, 0x3e, 0x72, 0xae, 0xe2, 0xbd, 0x83, 0xad, 0x41, 0x29, 0x84, 0x40, 0xc1,
	0x45, 0xd1, 0xa2, 0x74, 0x51, 0x0a, 0xa5, 0x54, 0xda, 0x6c, 0xd2, 0xcd, 0x14, 0xb2, 0x95, 0x44,
	0x8f, 0x92, 0x62, 0x1c, 0x9b, 0x8c, 0x01, 0xdb, 0xaf, 0xd4, 0x0f, 0x59, 0x66, 0xf2, 0xd0, 0x28,
	0x2d, 0xb4, 0x2b, 0xcf, 0xeb, 0x3f, 0xe7, 0x7f, 0x7e, 0x06, 0x4e, 0x96, 0x71, 0x30, 0x1b, 0xac,
	0x43, 0xc6, 0xd9, 0x40, 0x84, 0x7d, 0x19, 0x92, 0xb2, 0xfc, 0x31, 0xab, 0x50, 0x7e, 0x08, 0xd6,
	0x7c, 0x6b, 0xc6, 0xd0, 0xb0, 0xd9, 0xc2, 0x9f, 0xba, 0x4b, 0x87, 0x2d, 0x37, 0x01, 0x12, 0x02,
	0xea, 0xca, 0x0d, 0x50, 0x57, 0x0c, 0xa5, 0xa7, 0x51, 0x19, 0x93, 0x36, 0x54, 0x23, 0xff, 0x0d,
	0x27, 0x0b, 0x4f, 0xff, 0x63, 0x28, 0x3d, 0x95, 0x56, 0x44, 0x6a, 0x79, 0xa4, 0x0b, 0xda, 0x0c,
	0xe3, 0x49, 0xe0, 0xbe, 0xb0, 0x50, 0x2f, 0x19, 0x4a, 0xaf, 0x41, 0x6b, 0x33, 0x8c, 0x9f, 0x44,
	0x9e, 0x37, 0xfd, 0x15, 0x0b, 0x75, 0x75, 0xd7, 0x14, 0xb9, 0x79, 0x0b, 0xcd, 0x71, 0x88, 0x2e,
	0x47, 0xdb, 0xa1, 0xf8, 0xba, 0xc1, 0x88, 0xff, 0x68, 0xb3, 0x79, 0x07, 0xff, 0x76, 0xfa, 0x68,
	0xcd, 0x56, 0x11, 0x92, 0x0b, 0xa8, 0xc4, 0xf2, 0x08, 0xf9, 0xc4, 0xdf, 0x61, 0x2b, 0xb9, 0xb9,
	0x5f, 0x38, 0x90, 0xa6, 0x33, 0xe6, 0x39, 0x34, 0x29, 0x06, 0x2c, 0xfe, 0xde, 0x81, 0x30, 0x4a,
	0x51, 0x2c, 0xfd, 0xa5, 0xd1, 0x31, 0xfc, 0xb7, 0x90, 0xdb, 0x8e, 0xed, 0x47, 0x3c, 0x77, 0xda,
	0x87, 0x6a, 0xe2, 0x22, 0xd2, 0x15, 0xa3, 0xf4, 0xa5, 0xd5, 0x6c, 0xc8, 0xbc, 0x82, 0x96, 0x85,
	0xfc, 0x31, 0x44, 0xbc, 0xdf, 0x72, 0x8c, 0xf2, 0x77, 0xce, 0x00, 0xe6, 0x21, 0xe2, 0xc4, 0x13,
	0x55, 0xe9, 0x47, 0xa5, 0xda, 0x3c, 0x1b, 0x1b, 0x7e, 0x28, 0xa0, 0xd9, 0xce, 0x33, 0x86, 0xb1,
	0x3f, 0x45, 0x72, 0x03, 0xb5, 0x0c, 0x19, 0x39, 0x4d, 0xf7, 0x1d, 0xfc, 0x07, 0x9d, 0xf6, 0x51,
	0x3d, 0xdd, 0x74, 0x09, 0xb5, 0x8c, 0x56, 0x2e, 0x3e, 0xc0, 0xd7, 0xa9, 0xa7, 0x75, 0xf9, 0x65,
	0x25, 0x8a, 0x04, 0xdc, 0x9e, 0xa2, 0x40, 0xb2, 0xa8, 0x18, 0xbe, 0x83, 0xe6, 0x58, 0x99, 0xdb,
	0x11, 0x68, 0x39, 0x37, 0x52, 0x98, 0xeb, 0xe8, 0x69, 0x76, 0xcc, 0xf5, 0x1a, 0xea, 0xfb, 0x9c,
	0x0e, 0x74, 0xdd, 0x9d, 0xee, 0x08, 0xa5, 0x57, 0x91, 0xbd, 0xd1, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc3, 0xef, 0xc9, 0x1d, 0x38, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LVServiceClient is the client API for LVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LVServiceClient interface {
	CreateLV(ctx context.Context, in *CreateLVRequest, opts ...grpc.CallOption) (*CreateLVResponse, error)
	RemoveLV(ctx context.Context, in *RemoveLVRequest, opts ...grpc.CallOption) (*Empty, error)
	ResizeLV(ctx context.Context, in *ResizeLVRequest, opts ...grpc.CallOption) (*Empty, error)
}

type lVServiceClient struct {
	cc *grpc.ClientConn
}

func NewLVServiceClient(cc *grpc.ClientConn) LVServiceClient {
	return &lVServiceClient{cc}
}

func (c *lVServiceClient) CreateLV(ctx context.Context, in *CreateLVRequest, opts ...grpc.CallOption) (*CreateLVResponse, error) {
	out := new(CreateLVResponse)
	err := c.cc.Invoke(ctx, "/proto.LVService/CreateLV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lVServiceClient) RemoveLV(ctx context.Context, in *RemoveLVRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.LVService/RemoveLV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lVServiceClient) ResizeLV(ctx context.Context, in *ResizeLVRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.LVService/ResizeLV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LVServiceServer is the server API for LVService service.
type LVServiceServer interface {
	CreateLV(context.Context, *CreateLVRequest) (*CreateLVResponse, error)
	RemoveLV(context.Context, *RemoveLVRequest) (*Empty, error)
	ResizeLV(context.Context, *ResizeLVRequest) (*Empty, error)
}

func RegisterLVServiceServer(s *grpc.Server, srv LVServiceServer) {
	s.RegisterService(&_LVService_serviceDesc, srv)
}

func _LVService_CreateLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVServiceServer).CreateLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LVService/CreateLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVServiceServer).CreateLV(ctx, req.(*CreateLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LVService_RemoveLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVServiceServer).RemoveLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LVService/RemoveLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVServiceServer).RemoveLV(ctx, req.(*RemoveLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LVService_ResizeLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVServiceServer).ResizeLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LVService/ResizeLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVServiceServer).ResizeLV(ctx, req.(*ResizeLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LVService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LVService",
	HandlerType: (*LVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLV",
			Handler:    _LVService_CreateLV_Handler,
		},
		{
			MethodName: "RemoveLV",
			Handler:    _LVService_RemoveLV_Handler,
		},
		{
			MethodName: "ResizeLV",
			Handler:    _LVService_ResizeLV_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lvmd/proto/lvmd.proto",
}

// VGServiceClient is the client API for VGService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VGServiceClient interface {
	GetLVList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetLVListResponse, error)
	GetFreeBytes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFreeBytesResponse, error)
}

type vGServiceClient struct {
	cc *grpc.ClientConn
}

func NewVGServiceClient(cc *grpc.ClientConn) VGServiceClient {
	return &vGServiceClient{cc}
}

func (c *vGServiceClient) GetLVList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetLVListResponse, error) {
	out := new(GetLVListResponse)
	err := c.cc.Invoke(ctx, "/proto.VGService/GetLVList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vGServiceClient) GetFreeBytes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFreeBytesResponse, error) {
	out := new(GetFreeBytesResponse)
	err := c.cc.Invoke(ctx, "/proto.VGService/GetFreeBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VGServiceServer is the server API for VGService service.
type VGServiceServer interface {
	GetLVList(context.Context, *Empty) (*GetLVListResponse, error)
	GetFreeBytes(context.Context, *Empty) (*GetFreeBytesResponse, error)
}

func RegisterVGServiceServer(s *grpc.Server, srv VGServiceServer) {
	s.RegisterService(&_VGService_serviceDesc, srv)
}

func _VGService_GetLVList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VGServiceServer).GetLVList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VGService/GetLVList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VGServiceServer).GetLVList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VGService_GetFreeBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VGServiceServer).GetFreeBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VGService/GetFreeBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VGServiceServer).GetFreeBytes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VGService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VGService",
	HandlerType: (*VGServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLVList",
			Handler:    _VGService_GetLVList_Handler,
		},
		{
			MethodName: "GetFreeBytes",
			Handler:    _VGService_GetFreeBytes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lvmd/proto/lvmd.proto",
}