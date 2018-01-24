// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	CityNameCountryCodeRequest
	ForecastResponse
*/
package api

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

type CityNameCountryCodeRequest struct {
	CityName    string `protobuf:"bytes,1,opt,name=CityName" json:"CityName,omitempty"`
	CountryCode string `protobuf:"bytes,2,opt,name=CountryCode" json:"CountryCode,omitempty"`
}

func (m *CityNameCountryCodeRequest) Reset()                    { *m = CityNameCountryCodeRequest{} }
func (m *CityNameCountryCodeRequest) String() string            { return proto.CompactTextString(m) }
func (*CityNameCountryCodeRequest) ProtoMessage()               {}
func (*CityNameCountryCodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CityNameCountryCodeRequest) GetCityName() string {
	if m != nil {
		return m.CityName
	}
	return ""
}

func (m *CityNameCountryCodeRequest) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

type ForecastResponse struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *ForecastResponse) Reset()                    { *m = ForecastResponse{} }
func (m *ForecastResponse) String() string            { return proto.CompactTextString(m) }
func (*ForecastResponse) ProtoMessage()               {}
func (*ForecastResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ForecastResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*CityNameCountryCodeRequest)(nil), "api.CityNameCountryCodeRequest")
	proto.RegisterType((*ForecastResponse)(nil), "api.ForecastResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WeatherAppService service

type WeatherAppServiceClient interface {
	GetWeatherForecast(ctx context.Context, in *CityNameCountryCodeRequest, opts ...grpc.CallOption) (*ForecastResponse, error)
}

type weatherAppServiceClient struct {
	cc *grpc.ClientConn
}

func NewWeatherAppServiceClient(cc *grpc.ClientConn) WeatherAppServiceClient {
	return &weatherAppServiceClient{cc}
}

func (c *weatherAppServiceClient) GetWeatherForecast(ctx context.Context, in *CityNameCountryCodeRequest, opts ...grpc.CallOption) (*ForecastResponse, error) {
	out := new(ForecastResponse)
	err := grpc.Invoke(ctx, "/api.WeatherAppService/GetWeatherForecast", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WeatherAppService service

type WeatherAppServiceServer interface {
	GetWeatherForecast(context.Context, *CityNameCountryCodeRequest) (*ForecastResponse, error)
}

func RegisterWeatherAppServiceServer(s *grpc.Server, srv WeatherAppServiceServer) {
	s.RegisterService(&_WeatherAppService_serviceDesc, srv)
}

func _WeatherAppService_GetWeatherForecast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityNameCountryCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherAppServiceServer).GetWeatherForecast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WeatherAppService/GetWeatherForecast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherAppServiceServer).GetWeatherForecast(ctx, req.(*CityNameCountryCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeatherAppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WeatherAppService",
	HandlerType: (*WeatherAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeatherForecast",
			Handler:    _WeatherAppService_GetWeatherForecast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x8a, 0xe2, 0x92, 0x72, 0xce,
	0x2c, 0xa9, 0xf4, 0x4b, 0xcc, 0x4d, 0x75, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xaa, 0x74, 0xce, 0x4f,
	0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x80, 0xc9, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x0a, 0x5c, 0xdc, 0x48, 0x3a, 0x24, 0x98, 0xc0,
	0xd2, 0xc8, 0x42, 0x4a, 0x6a, 0x5c, 0x02, 0x6e, 0xf9, 0x45, 0xa9, 0xc9, 0x89, 0xc5, 0x25, 0x41,
	0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0x21, 0xa9, 0x15, 0x25, 0x50,
	0xd3, 0xc0, 0x6c, 0xa3, 0x64, 0x2e, 0xc1, 0xf0, 0xd4, 0xc4, 0x92, 0x8c, 0xd4, 0x22, 0xc7, 0x82,
	0x82, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x3f, 0x2e, 0x21, 0xf7, 0xd4, 0x12, 0xa8,
	0x38, 0xcc, 0x18, 0x21, 0x79, 0x3d, 0x90, 0xfb, 0x71, 0xbb, 0x58, 0x4a, 0x14, 0xac, 0x00, 0xdd,
	0x5a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xa7, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x77,
	0xe6, 0xb9, 0x01, 0x01, 0x00, 0x00,
}