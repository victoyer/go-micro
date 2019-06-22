// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessPKG

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=Capacity,proto3" json:"Capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwerId               string   `protobuf:"bytes,6,opt,name=ower_id,json=owerId,proto3" json:"ower_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwerId() string {
	if m != nil {
		return m.OwerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=Created,proto3" json:"Created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessPKG.Vessel")
	proto.RegisterType((*Specification)(nil), "vessPKG.Specification")
	proto.RegisterType((*Response)(nil), "vessPKG.Response")
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716) }

var fileDescriptor_04ef66862bb50716 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x14, 0xfc, 0x92, 0xb6, 0x9b, 0xf4, 0x7d, 0x54, 0xf1, 0x1d, 0x74, 0x2d, 0x0a, 0x21, 0x17, 0xeb,
	0x25, 0x42, 0xbd, 0x7a, 0x91, 0x82, 0xa2, 0x5e, 0x64, 0x0b, 0x7a, 0x2c, 0xdb, 0xdd, 0xa7, 0x2e,
	0xb4, 0x49, 0x48, 0x42, 0x5b, 0x0f, 0xfe, 0x15, 0x7f, 0xab, 0x74, 0xb7, 0x89, 0x58, 0x72, 0xda,
	0x37, 0x33, 0x6f, 0x87, 0xd9, 0x59, 0x38, 0xcd, 0x8b, 0xac, 0xca, 0xae, 0x56, 0x54, 0x96, 0xb4,
	0xd8, 0x1d, 0x89, 0xe5, 0x30, 0xd8, 0xa2, 0xe7, 0xa7, 0xfb, 0xf8, 0xdb, 0x03, 0xf6, 0x62, 0x15,
	0x3c, 0x00, 0xdf, 0x68, 0xee, 0x45, 0xde, 0xa8, 0x2f, 0x7c, 0xa3, 0x71, 0x08, 0xe1, 0x44, 0xe6,
	0x52, 0x99, 0xea, 0x93, 0xfb, 0x91, 0x37, 0xea, 0x89, 0x06, 0xe3, 0x39, 0xc0, 0x52, 0x6e, 0x66,
	0x6b, 0x32, 0xef, 0x1f, 0x15, 0xef, 0x58, 0xb5, 0xbf, 0x94, 0x9b, 0x57, 0x4b, 0x20, 0x42, 0x37,
	0x95, 0x4b, 0xe2, 0x5d, 0x6b, 0x66, 0x67, 0x3c, 0x83, 0xbe, 0x5c, 0x49, 0xb3, 0x90, 0xf3, 0x05,
	0xf1, 0x5e, 0xe4, 0x8d, 0x42, 0xf1, 0x4b, 0xe0, 0x09, 0x04, 0xd9, 0x9a, 0x8a, 0x99, 0xd1, 0x9c,
	0xd9, 0x4b, 0x6c, 0x0b, 0x1f, 0x74, 0xfc, 0x08, 0x83, 0x69, 0x4e, 0xca, 0xbc, 0x19, 0x25, 0x2b,
	0x93, 0xa5, 0xdb, 0x58, 0xaa, 0x8e, 0xe5, 0xb9, 0x58, 0xaa, 0x3d, 0x96, 0xbf, 0x17, 0x2b, 0xde,
	0x40, 0x28, 0xa8, 0xcc, 0xb3, 0xb4, 0x24, 0xbc, 0x00, 0xe6, 0x1a, 0xb1, 0x26, 0xff, 0xc7, 0x87,
	0xc9, 0xae, 0x92, 0xc4, 0xd5, 0x21, 0x76, 0x32, 0x5e, 0x42, 0xe0, 0xa6, 0x92, 0xfb, 0x51, 0xa7,
	0x6d, 0xb3, 0xd6, 0x91, 0x43, 0x30, 0x29, 0x48, 0x56, 0xa4, 0x6d, 0x25, 0xa1, 0xa8, 0xe1, 0xf8,
	0x0b, 0x06, 0x6e, 0x79, 0x4a, 0xc5, 0xca, 0x28, 0xc2, 0x1b, 0x18, 0xdc, 0x99, 0x54, 0xdf, 0x36,
	0x05, 0x1c, 0x37, 0xae, 0x7f, 0x9e, 0x3b, 0x3c, 0x6a, 0xf8, 0x3a, 0x7a, 0xfc, 0x0f, 0x13, 0x60,
	0xce, 0x19, 0xf7, 0xc3, 0xb4, 0xee, 0xcf, 0x99, 0xfd, 0xf5, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7f, 0xb2, 0x03, 0xb1, 0x12, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessPKG"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}