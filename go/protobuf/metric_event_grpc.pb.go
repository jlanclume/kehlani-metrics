// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: metric_event.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MetricEventService_SaveMetricEvent_FullMethodName = "/protobuf.MetricEventService/SaveMetricEvent"
)

// MetricEventServiceClient is the client API for MetricEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricEventServiceClient interface {
	SaveMetricEvent(ctx context.Context, in *MetricEventRequest, opts ...grpc.CallOption) (*MetricEventResponse, error)
}

type metricEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricEventServiceClient(cc grpc.ClientConnInterface) MetricEventServiceClient {
	return &metricEventServiceClient{cc}
}

func (c *metricEventServiceClient) SaveMetricEvent(ctx context.Context, in *MetricEventRequest, opts ...grpc.CallOption) (*MetricEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetricEventResponse)
	err := c.cc.Invoke(ctx, MetricEventService_SaveMetricEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricEventServiceServer is the server API for MetricEventService service.
// All implementations must embed UnimplementedMetricEventServiceServer
// for forward compatibility.
type MetricEventServiceServer interface {
	SaveMetricEvent(context.Context, *MetricEventRequest) (*MetricEventResponse, error)
	mustEmbedUnimplementedMetricEventServiceServer()
}

// UnimplementedMetricEventServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetricEventServiceServer struct{}

func (UnimplementedMetricEventServiceServer) SaveMetricEvent(context.Context, *MetricEventRequest) (*MetricEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMetricEvent not implemented")
}
func (UnimplementedMetricEventServiceServer) mustEmbedUnimplementedMetricEventServiceServer() {}
func (UnimplementedMetricEventServiceServer) testEmbeddedByValue()                            {}

// UnsafeMetricEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricEventServiceServer will
// result in compilation errors.
type UnsafeMetricEventServiceServer interface {
	mustEmbedUnimplementedMetricEventServiceServer()
}

func RegisterMetricEventServiceServer(s grpc.ServiceRegistrar, srv MetricEventServiceServer) {
	// If the following call pancis, it indicates UnimplementedMetricEventServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MetricEventService_ServiceDesc, srv)
}

func _MetricEventService_SaveMetricEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricEventServiceServer).SaveMetricEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricEventService_SaveMetricEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricEventServiceServer).SaveMetricEvent(ctx, req.(*MetricEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricEventService_ServiceDesc is the grpc.ServiceDesc for MetricEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.MetricEventService",
	HandlerType: (*MetricEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveMetricEvent",
			Handler:    _MetricEventService_SaveMetricEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metric_event.proto",
}
