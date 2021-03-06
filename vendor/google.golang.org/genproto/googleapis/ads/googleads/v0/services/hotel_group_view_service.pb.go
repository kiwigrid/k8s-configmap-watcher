// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/hotel_group_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
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

// Request message for [HotelGroupViewService.GetHotelGroupView][google.ads.googleads.v0.services.HotelGroupViewService.GetHotelGroupView].
type GetHotelGroupViewRequest struct {
	// Resource name of the Hotel Group View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotelGroupViewRequest) Reset()         { *m = GetHotelGroupViewRequest{} }
func (m *GetHotelGroupViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotelGroupViewRequest) ProtoMessage()    {}
func (*GetHotelGroupViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_group_view_service_d553180658033695, []int{0}
}
func (m *GetHotelGroupViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotelGroupViewRequest.Unmarshal(m, b)
}
func (m *GetHotelGroupViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotelGroupViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetHotelGroupViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotelGroupViewRequest.Merge(dst, src)
}
func (m *GetHotelGroupViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotelGroupViewRequest.Size(m)
}
func (m *GetHotelGroupViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotelGroupViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotelGroupViewRequest proto.InternalMessageInfo

func (m *GetHotelGroupViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHotelGroupViewRequest)(nil), "google.ads.googleads.v0.services.GetHotelGroupViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotelGroupViewServiceClient is the client API for HotelGroupViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotelGroupViewServiceClient interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error)
}

type hotelGroupViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewHotelGroupViewServiceClient(cc *grpc.ClientConn) HotelGroupViewServiceClient {
	return &hotelGroupViewServiceClient{cc}
}

func (c *hotelGroupViewServiceClient) GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error) {
	out := new(resources.HotelGroupView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.HotelGroupViewService/GetHotelGroupView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelGroupViewServiceServer is the server API for HotelGroupViewService service.
type HotelGroupViewServiceServer interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(context.Context, *GetHotelGroupViewRequest) (*resources.HotelGroupView, error)
}

func RegisterHotelGroupViewServiceServer(s *grpc.Server, srv HotelGroupViewServiceServer) {
	s.RegisterService(&_HotelGroupViewService_serviceDesc, srv)
}

func _HotelGroupViewService_GetHotelGroupView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelGroupViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.HotelGroupViewService/GetHotelGroupView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, req.(*GetHotelGroupViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelGroupViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.HotelGroupViewService",
	HandlerType: (*HotelGroupViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelGroupView",
			Handler:    _HotelGroupViewService_GetHotelGroupView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/hotel_group_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/hotel_group_view_service.proto", fileDescriptor_hotel_group_view_service_d553180658033695)
}

var fileDescriptor_hotel_group_view_service_d553180658033695 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xe2, 0xd4,
	0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x62, 0xfd, 0x8c, 0xfc, 0x92, 0xd4, 0x9c, 0xf8, 0xf4, 0xa2, 0xfc,
	0xd2, 0x82, 0xf8, 0xb2, 0xcc, 0xd4, 0xf2, 0x78, 0xa8, 0x8c, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe,
	0x90, 0x02, 0x44, 0x97, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0xdc, 0x00, 0xbd, 0x32, 0x03, 0x3d, 0x98,
	0x01, 0x52, 0x16, 0xb8, 0xac, 0x28, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0xc2, 0x66, 0x07, 0xc4, 0x6c,
	0x29, 0x19, 0x98, 0xce, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc,
	0xbc, 0x62, 0x88, 0xac, 0x92, 0x3d, 0x97, 0x84, 0x7b, 0x6a, 0x89, 0x07, 0x48, 0xab, 0x3b, 0x48,
	0x67, 0x58, 0x66, 0x6a, 0x79, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x32, 0x17, 0x2f,
	0xcc, 0xf4, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x1e, 0x98,
	0xa0, 0x5f, 0x62, 0x6e, 0xaa, 0xd1, 0x75, 0x46, 0x2e, 0x51, 0x54, 0xed, 0xc1, 0x10, 0x37, 0x0b,
	0xed, 0x65, 0xe4, 0x12, 0xc4, 0x30, 0x5b, 0xc8, 0x4a, 0x8f, 0x90, 0x5f, 0xf5, 0x70, 0x39, 0x48,
	0xca, 0x10, 0xa7, 0x5e, 0x78, 0x28, 0xe8, 0xa1, 0xea, 0x54, 0xb2, 0x6c, 0xba, 0xfc, 0x64, 0x32,
	0x93, 0xb1, 0x90, 0x21, 0x28, 0xac, 0xaa, 0x51, 0xbc, 0x63, 0x9b, 0x5c, 0x5a, 0x5c, 0x92, 0x9f,
	0x9b, 0x5a, 0x54, 0xac, 0xaf, 0x05, 0x09, 0x3c, 0xb8, 0xb6, 0x62, 0x7d, 0xad, 0x5a, 0xa7, 0x5b,
	0x8c, 0x5c, 0x2a, 0xc9, 0xf9, 0xb9, 0x04, 0xdd, 0xeb, 0x24, 0x85, 0xd5, 0xff, 0x01, 0xa0, 0xf0,
	0x0d, 0x60, 0x8c, 0xf2, 0x80, 0xea, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a,
	0xd7, 0x4f, 0x4f, 0xcd, 0x03, 0x87, 0x3e, 0x2c, 0x26, 0x0b, 0x32, 0x8b, 0x71, 0xa7, 0x1d, 0x6b,
	0x18, 0x63, 0x11, 0x13, 0xb3, 0xbb, 0xa3, 0xe3, 0x2a, 0x26, 0x05, 0x77, 0x88, 0x81, 0x8e, 0x29,
	0xc5, 0x7a, 0x10, 0x26, 0x88, 0x15, 0x66, 0xa0, 0x07, 0xb5, 0xb8, 0xf8, 0x14, 0x4c, 0x49, 0x8c,
	0x63, 0x4a, 0x71, 0x0c, 0x5c, 0x49, 0x4c, 0x98, 0x41, 0x0c, 0x4c, 0x49, 0x12, 0x1b, 0xd8, 0x01,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xe6, 0xf0, 0xd1, 0xbb, 0x02, 0x00, 0x00,
}
