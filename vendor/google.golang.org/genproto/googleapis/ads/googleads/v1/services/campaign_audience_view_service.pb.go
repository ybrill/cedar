// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_audience_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [CampaignAudienceViewService.GetCampaignAudienceView][google.ads.googleads.v1.services.CampaignAudienceViewService.GetCampaignAudienceView].
type GetCampaignAudienceViewRequest struct {
	// The resource name of the campaign audience view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignAudienceViewRequest) Reset()         { *m = GetCampaignAudienceViewRequest{} }
func (m *GetCampaignAudienceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignAudienceViewRequest) ProtoMessage()    {}
func (*GetCampaignAudienceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae1606ecaa0347f1, []int{0}
}

func (m *GetCampaignAudienceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Unmarshal(m, b)
}
func (m *GetCampaignAudienceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignAudienceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignAudienceViewRequest.Merge(m, src)
}
func (m *GetCampaignAudienceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Size(m)
}
func (m *GetCampaignAudienceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignAudienceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignAudienceViewRequest proto.InternalMessageInfo

func (m *GetCampaignAudienceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignAudienceViewRequest)(nil), "google.ads.googleads.v1.services.GetCampaignAudienceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_audience_view_service.proto", fileDescriptor_ae1606ecaa0347f1)
}

var fileDescriptor_ae1606ecaa0347f1 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x25, 0x11, 0x04, 0x83, 0x6e, 0xb2, 0x69, 0x49, 0x45, 0x42, 0xed, 0x42, 0xba, 0x98, 0x21,
	0xba, 0x28, 0x8e, 0xf8, 0x91, 0x4a, 0xa9, 0x2b, 0x29, 0x15, 0xb2, 0x90, 0x40, 0x18, 0x93, 0x4b,
	0x18, 0x68, 0x66, 0x62, 0x26, 0x49, 0x17, 0xe2, 0xc6, 0x8d, 0x7b, 0xfd, 0x07, 0x2e, 0xfd, 0x29,
	0xdd, 0x3e, 0xde, 0x3f, 0x78, 0xab, 0xf7, 0x2b, 0x1e, 0xe9, 0x64, 0xd2, 0x3e, 0x68, 0xda, 0xdd,
	0x61, 0xee, 0xb9, 0xe7, 0xdc, 0x7b, 0xee, 0x58, 0x8b, 0x54, 0x88, 0x74, 0x03, 0x98, 0x26, 0x12,
	0x2b, 0xd8, 0xa0, 0xda, 0xc3, 0x12, 0x8a, 0x9a, 0xc5, 0x20, 0x71, 0x4c, 0xb3, 0x9c, 0xb2, 0x94,
	0x47, 0xb4, 0x4a, 0x18, 0xf0, 0x18, 0xa2, 0x9a, 0xc1, 0x36, 0x6a, 0xeb, 0x28, 0x2f, 0x44, 0x29,
	0x6c, 0x57, 0xf5, 0x22, 0x9a, 0x48, 0xd4, 0xc9, 0xa0, 0xda, 0x43, 0x5a, 0xc6, 0x79, 0xd7, 0x67,
	0x54, 0x80, 0x14, 0x55, 0xd1, 0xef, 0xa4, 0x1c, 0x9c, 0xa7, 0xba, 0x3f, 0x67, 0x98, 0x72, 0x2e,
	0x4a, 0x5a, 0x32, 0xc1, 0x65, 0x5b, 0x1d, 0x1c, 0x55, 0xe3, 0x0d, 0x03, 0x5e, 0xaa, 0xc2, 0x78,
	0x61, 0x3d, 0x5b, 0x42, 0xf9, 0xb1, 0x55, 0xf6, 0x5b, 0xe1, 0x80, 0xc1, 0x76, 0x0d, 0xdf, 0x2b,
	0x90, 0xa5, 0xfd, 0xdc, 0x7a, 0xa2, 0x47, 0x88, 0x38, 0xcd, 0x60, 0x68, 0xb8, 0xc6, 0x8b, 0x47,
	0xeb, 0xc7, 0xfa, 0xf1, 0x33, 0xcd, 0xe0, 0xe5, 0x1f, 0xd3, 0x1a, 0x9d, 0x12, 0xf9, 0xa2, 0xd6,
	0xb3, 0xaf, 0x0d, 0x6b, 0xd0, 0xe3, 0x63, 0x7f, 0x40, 0x97, 0xc2, 0x41, 0xe7, 0x47, 0x74, 0x66,
	0xbd, 0x0a, 0x5d, 0x78, 0xe8, 0x54, 0xff, 0xf8, 0xfd, 0xaf, 0xab, 0x9b, 0xbf, 0xe6, 0x6b, 0x7b,
	0xd6, 0x04, 0xfd, 0xe3, 0xde, 0x9a, 0x6f, 0xe3, 0x4a, 0x96, 0x22, 0x83, 0x42, 0xe2, 0x69, 0x97,
	0xfc, 0x71, 0xb3, 0xc4, 0xd3, 0x9f, 0xce, 0x68, 0xe7, 0x0f, 0x0f, 0x86, 0x2d, 0xca, 0x99, 0x44,
	0xb1, 0xc8, 0xe6, 0xbf, 0x4d, 0x6b, 0x12, 0x8b, 0xec, 0xe2, 0x7a, 0x73, 0xf7, 0x4c, 0x74, 0xab,
	0xe6, 0x4c, 0x2b, 0xe3, 0xeb, 0xa7, 0x56, 0x25, 0x15, 0x1b, 0xca, 0x53, 0x24, 0x8a, 0x14, 0xa7,
	0xc0, 0xf7, 0x47, 0xc4, 0x07, 0xdf, 0xfe, 0x7f, 0xfa, 0x46, 0x83, 0x7f, 0xe6, 0x83, 0xa5, 0xef,
	0xff, 0x37, 0xdd, 0xa5, 0x12, 0xf4, 0x13, 0x89, 0x14, 0x6c, 0x50, 0xe0, 0xa1, 0xd6, 0x58, 0xee,
	0x34, 0x25, 0xf4, 0x13, 0x19, 0x76, 0x94, 0x30, 0xf0, 0x42, 0x4d, 0xb9, 0x35, 0x27, 0xea, 0x9d,
	0x10, 0x3f, 0x91, 0x84, 0x74, 0x24, 0x42, 0x02, 0x8f, 0x10, 0x4d, 0xfb, 0xf6, 0x70, 0x3f, 0xe7,
	0xab, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x22, 0x25, 0x65, 0x4e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignAudienceViewServiceClient is the client API for CampaignAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignAudienceViewServiceClient interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error)
}

type campaignAudienceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignAudienceViewServiceClient(cc grpc.ClientConnInterface) CampaignAudienceViewServiceClient {
	return &campaignAudienceViewServiceClient{cc}
}

func (c *campaignAudienceViewServiceClient) GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error) {
	out := new(resources.CampaignAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignAudienceViewService/GetCampaignAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAudienceViewServiceServer is the server API for CampaignAudienceViewService service.
type CampaignAudienceViewServiceServer interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(context.Context, *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error)
}

// UnimplementedCampaignAudienceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignAudienceViewServiceServer struct {
}

func (*UnimplementedCampaignAudienceViewServiceServer) GetCampaignAudienceView(ctx context.Context, req *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignAudienceView not implemented")
}

func RegisterCampaignAudienceViewServiceServer(s *grpc.Server, srv CampaignAudienceViewServiceServer) {
	s.RegisterService(&_CampaignAudienceViewService_serviceDesc, srv)
}

func _CampaignAudienceViewService_GetCampaignAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignAudienceViewService/GetCampaignAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, req.(*GetCampaignAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignAudienceViewService",
	HandlerType: (*CampaignAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignAudienceView",
			Handler:    _CampaignAudienceViewService_GetCampaignAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_audience_view_service.proto",
}