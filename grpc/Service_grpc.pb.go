// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SpanClient is the client API for Span service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpanClient interface {
	SendSpan(ctx context.Context, opts ...grpc.CallOption) (Span_SendSpanClient, error)
}

type spanClient struct {
	cc grpc.ClientConnInterface
}

func NewSpanClient(cc grpc.ClientConnInterface) SpanClient {
	return &spanClient{cc}
}

func (c *spanClient) SendSpan(ctx context.Context, opts ...grpc.CallOption) (Span_SendSpanClient, error) {
	stream, err := c.cc.NewStream(ctx, &Span_ServiceDesc.Streams[0], "/v1.Span/SendSpan", opts...)
	if err != nil {
		return nil, err
	}
	x := &spanSendSpanClient{stream}
	return x, nil
}

type Span_SendSpanClient interface {
	Send(*PSpanMessage) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type spanSendSpanClient struct {
	grpc.ClientStream
}

func (x *spanSendSpanClient) Send(m *PSpanMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *spanSendSpanClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpanServer is the server API for Span service.
// All implementations must embed UnimplementedSpanServer
// for forward compatibility
type SpanServer interface {
	SendSpan(Span_SendSpanServer) error
	mustEmbedUnimplementedSpanServer()
}

// UnimplementedSpanServer must be embedded to have forward compatible implementations.
type UnimplementedSpanServer struct {
}

func (UnimplementedSpanServer) SendSpan(Span_SendSpanServer) error {
	return status.Errorf(codes.Unimplemented, "method SendSpan not implemented")
}
func (UnimplementedSpanServer) mustEmbedUnimplementedSpanServer() {}

// UnsafeSpanServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpanServer will
// result in compilation errors.
type UnsafeSpanServer interface {
	mustEmbedUnimplementedSpanServer()
}

func RegisterSpanServer(s grpc.ServiceRegistrar, srv SpanServer) {
	s.RegisterService(&Span_ServiceDesc, srv)
}

func _Span_SendSpan_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SpanServer).SendSpan(&spanSendSpanServer{stream})
}

type Span_SendSpanServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PSpanMessage, error)
	grpc.ServerStream
}

type spanSendSpanServer struct {
	grpc.ServerStream
}

func (x *spanSendSpanServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *spanSendSpanServer) Recv() (*PSpanMessage, error) {
	m := new(PSpanMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Span_ServiceDesc is the grpc.ServiceDesc for Span service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Span_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Span",
	HandlerType: (*SpanServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendSpan",
			Handler:       _Span_SendSpan_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "v1/Service.proto",
}

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	RequestAgentInfo(ctx context.Context, in *PAgentInfo, opts ...grpc.CallOption) (*PResult, error)
	PingSession(ctx context.Context, opts ...grpc.CallOption) (Agent_PingSessionClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) RequestAgentInfo(ctx context.Context, in *PAgentInfo, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Agent/RequestAgentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PingSession(ctx context.Context, opts ...grpc.CallOption) (Agent_PingSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/v1.Agent/PingSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentPingSessionClient{stream}
	return x, nil
}

type Agent_PingSessionClient interface {
	Send(*PPing) error
	Recv() (*PPing, error)
	grpc.ClientStream
}

type agentPingSessionClient struct {
	grpc.ClientStream
}

func (x *agentPingSessionClient) Send(m *PPing) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentPingSessionClient) Recv() (*PPing, error) {
	m := new(PPing)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	RequestAgentInfo(context.Context, *PAgentInfo) (*PResult, error)
	PingSession(Agent_PingSessionServer) error
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) RequestAgentInfo(context.Context, *PAgentInfo) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAgentInfo not implemented")
}
func (UnimplementedAgentServer) PingSession(Agent_PingSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method PingSession not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_RequestAgentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PAgentInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RequestAgentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Agent/RequestAgentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RequestAgentInfo(ctx, req.(*PAgentInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_PingSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).PingSession(&agentPingSessionServer{stream})
}

type Agent_PingSessionServer interface {
	Send(*PPing) error
	Recv() (*PPing, error)
	grpc.ServerStream
}

type agentPingSessionServer struct {
	grpc.ServerStream
}

func (x *agentPingSessionServer) Send(m *PPing) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentPingSessionServer) Recv() (*PPing, error) {
	m := new(PPing)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAgentInfo",
			Handler:    _Agent_RequestAgentInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingSession",
			Handler:       _Agent_PingSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/Service.proto",
}

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataClient interface {
	RequestSqlMetaData(ctx context.Context, in *PSqlMetaData, opts ...grpc.CallOption) (*PResult, error)
	RequestApiMetaData(ctx context.Context, in *PApiMetaData, opts ...grpc.CallOption) (*PResult, error)
	RequestStringMetaData(ctx context.Context, in *PStringMetaData, opts ...grpc.CallOption) (*PResult, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) RequestSqlMetaData(ctx context.Context, in *PSqlMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestSqlMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) RequestApiMetaData(ctx context.Context, in *PApiMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestApiMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) RequestStringMetaData(ctx context.Context, in *PStringMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestStringMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
// All implementations must embed UnimplementedMetadataServer
// for forward compatibility
type MetadataServer interface {
	RequestSqlMetaData(context.Context, *PSqlMetaData) (*PResult, error)
	RequestApiMetaData(context.Context, *PApiMetaData) (*PResult, error)
	RequestStringMetaData(context.Context, *PStringMetaData) (*PResult, error)
	mustEmbedUnimplementedMetadataServer()
}

// UnimplementedMetadataServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (UnimplementedMetadataServer) RequestSqlMetaData(context.Context, *PSqlMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSqlMetaData not implemented")
}
func (UnimplementedMetadataServer) RequestApiMetaData(context.Context, *PApiMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestApiMetaData not implemented")
}
func (UnimplementedMetadataServer) RequestStringMetaData(context.Context, *PStringMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestStringMetaData not implemented")
}
func (UnimplementedMetadataServer) mustEmbedUnimplementedMetadataServer() {}

// UnsafeMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServer will
// result in compilation errors.
type UnsafeMetadataServer interface {
	mustEmbedUnimplementedMetadataServer()
}

func RegisterMetadataServer(s grpc.ServiceRegistrar, srv MetadataServer) {
	s.RegisterService(&Metadata_ServiceDesc, srv)
}

func _Metadata_RequestSqlMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PSqlMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestSqlMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestSqlMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestSqlMetaData(ctx, req.(*PSqlMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_RequestApiMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PApiMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestApiMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestApiMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestApiMetaData(ctx, req.(*PApiMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_RequestStringMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PStringMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestStringMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestStringMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestStringMetaData(ctx, req.(*PStringMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

// Metadata_ServiceDesc is the grpc.ServiceDesc for Metadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestSqlMetaData",
			Handler:    _Metadata_RequestSqlMetaData_Handler,
		},
		{
			MethodName: "RequestApiMetaData",
			Handler:    _Metadata_RequestApiMetaData_Handler,
		},
		{
			MethodName: "RequestStringMetaData",
			Handler:    _Metadata_RequestStringMetaData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/Service.proto",
}

// StatClient is the client API for Stat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatClient interface {
	SendAgentStat(ctx context.Context, opts ...grpc.CallOption) (Stat_SendAgentStatClient, error)
}

type statClient struct {
	cc grpc.ClientConnInterface
}

func NewStatClient(cc grpc.ClientConnInterface) StatClient {
	return &statClient{cc}
}

func (c *statClient) SendAgentStat(ctx context.Context, opts ...grpc.CallOption) (Stat_SendAgentStatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stat_ServiceDesc.Streams[0], "/v1.Stat/SendAgentStat", opts...)
	if err != nil {
		return nil, err
	}
	x := &statSendAgentStatClient{stream}
	return x, nil
}

type Stat_SendAgentStatClient interface {
	Send(*PStatMessage) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type statSendAgentStatClient struct {
	grpc.ClientStream
}

func (x *statSendAgentStatClient) Send(m *PStatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *statSendAgentStatClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatServer is the server API for Stat service.
// All implementations must embed UnimplementedStatServer
// for forward compatibility
type StatServer interface {
	SendAgentStat(Stat_SendAgentStatServer) error
	mustEmbedUnimplementedStatServer()
}

// UnimplementedStatServer must be embedded to have forward compatible implementations.
type UnimplementedStatServer struct {
}

func (UnimplementedStatServer) SendAgentStat(Stat_SendAgentStatServer) error {
	return status.Errorf(codes.Unimplemented, "method SendAgentStat not implemented")
}
func (UnimplementedStatServer) mustEmbedUnimplementedStatServer() {}

// UnsafeStatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatServer will
// result in compilation errors.
type UnsafeStatServer interface {
	mustEmbedUnimplementedStatServer()
}

func RegisterStatServer(s grpc.ServiceRegistrar, srv StatServer) {
	s.RegisterService(&Stat_ServiceDesc, srv)
}

func _Stat_SendAgentStat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StatServer).SendAgentStat(&statSendAgentStatServer{stream})
}

type Stat_SendAgentStatServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PStatMessage, error)
	grpc.ServerStream
}

type statSendAgentStatServer struct {
	grpc.ServerStream
}

func (x *statSendAgentStatServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *statSendAgentStatServer) Recv() (*PStatMessage, error) {
	m := new(PStatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Stat_ServiceDesc is the grpc.ServiceDesc for Stat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Stat",
	HandlerType: (*StatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendAgentStat",
			Handler:       _Stat_SendAgentStat_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "v1/Service.proto",
}

// ProfilerCommandServiceClient is the client API for ProfilerCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilerCommandServiceClient interface {
	// Deprecated: Do not use.
	// deprecated api
	HandleCommand(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandClient, error)
	HandleCommandV2(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandV2Client, error)
	CommandEcho(ctx context.Context, in *PCmdEchoResponse, opts ...grpc.CallOption) (*empty.Empty, error)
	CommandStreamActiveThreadCount(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_CommandStreamActiveThreadCountClient, error)
	CommandActiveThreadDump(ctx context.Context, in *PCmdActiveThreadDumpRes, opts ...grpc.CallOption) (*empty.Empty, error)
	CommandActiveThreadLightDump(ctx context.Context, in *PCmdActiveThreadLightDumpRes, opts ...grpc.CallOption) (*empty.Empty, error)
}

type profilerCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilerCommandServiceClient(cc grpc.ClientConnInterface) ProfilerCommandServiceClient {
	return &profilerCommandServiceClient{cc}
}

// Deprecated: Do not use.
func (c *profilerCommandServiceClient) HandleCommand(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProfilerCommandService_ServiceDesc.Streams[0], "/v1.ProfilerCommandService/HandleCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilerCommandServiceHandleCommandClient{stream}
	return x, nil
}

type ProfilerCommandService_HandleCommandClient interface {
	Send(*PCmdMessage) error
	Recv() (*PCmdRequest, error)
	grpc.ClientStream
}

type profilerCommandServiceHandleCommandClient struct {
	grpc.ClientStream
}

func (x *profilerCommandServiceHandleCommandClient) Send(m *PCmdMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandClient) Recv() (*PCmdRequest, error) {
	m := new(PCmdRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profilerCommandServiceClient) HandleCommandV2(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &ProfilerCommandService_ServiceDesc.Streams[1], "/v1.ProfilerCommandService/HandleCommandV2", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilerCommandServiceHandleCommandV2Client{stream}
	return x, nil
}

type ProfilerCommandService_HandleCommandV2Client interface {
	Send(*PCmdMessage) error
	Recv() (*PCmdRequest, error)
	grpc.ClientStream
}

type profilerCommandServiceHandleCommandV2Client struct {
	grpc.ClientStream
}

func (x *profilerCommandServiceHandleCommandV2Client) Send(m *PCmdMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandV2Client) Recv() (*PCmdRequest, error) {
	m := new(PCmdRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profilerCommandServiceClient) CommandEcho(ctx context.Context, in *PCmdEchoResponse, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerCommandServiceClient) CommandStreamActiveThreadCount(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_CommandStreamActiveThreadCountClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProfilerCommandService_ServiceDesc.Streams[2], "/v1.ProfilerCommandService/CommandStreamActiveThreadCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilerCommandServiceCommandStreamActiveThreadCountClient{stream}
	return x, nil
}

type ProfilerCommandService_CommandStreamActiveThreadCountClient interface {
	Send(*PCmdActiveThreadCountRes) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type profilerCommandServiceCommandStreamActiveThreadCountClient struct {
	grpc.ClientStream
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountClient) Send(m *PCmdActiveThreadCountRes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profilerCommandServiceClient) CommandActiveThreadDump(ctx context.Context, in *PCmdActiveThreadDumpRes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandActiveThreadDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerCommandServiceClient) CommandActiveThreadLightDump(ctx context.Context, in *PCmdActiveThreadLightDumpRes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandActiveThreadLightDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilerCommandServiceServer is the server API for ProfilerCommandService service.
// All implementations must embed UnimplementedProfilerCommandServiceServer
// for forward compatibility
type ProfilerCommandServiceServer interface {
	// Deprecated: Do not use.
	// deprecated api
	HandleCommand(ProfilerCommandService_HandleCommandServer) error
	HandleCommandV2(ProfilerCommandService_HandleCommandV2Server) error
	CommandEcho(context.Context, *PCmdEchoResponse) (*empty.Empty, error)
	CommandStreamActiveThreadCount(ProfilerCommandService_CommandStreamActiveThreadCountServer) error
	CommandActiveThreadDump(context.Context, *PCmdActiveThreadDumpRes) (*empty.Empty, error)
	CommandActiveThreadLightDump(context.Context, *PCmdActiveThreadLightDumpRes) (*empty.Empty, error)
	mustEmbedUnimplementedProfilerCommandServiceServer()
}

// UnimplementedProfilerCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfilerCommandServiceServer struct {
}

func (UnimplementedProfilerCommandServiceServer) HandleCommand(ProfilerCommandService_HandleCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleCommand not implemented")
}
func (UnimplementedProfilerCommandServiceServer) HandleCommandV2(ProfilerCommandService_HandleCommandV2Server) error {
	return status.Errorf(codes.Unimplemented, "method HandleCommandV2 not implemented")
}
func (UnimplementedProfilerCommandServiceServer) CommandEcho(context.Context, *PCmdEchoResponse) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandEcho not implemented")
}
func (UnimplementedProfilerCommandServiceServer) CommandStreamActiveThreadCount(ProfilerCommandService_CommandStreamActiveThreadCountServer) error {
	return status.Errorf(codes.Unimplemented, "method CommandStreamActiveThreadCount not implemented")
}
func (UnimplementedProfilerCommandServiceServer) CommandActiveThreadDump(context.Context, *PCmdActiveThreadDumpRes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandActiveThreadDump not implemented")
}
func (UnimplementedProfilerCommandServiceServer) CommandActiveThreadLightDump(context.Context, *PCmdActiveThreadLightDumpRes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandActiveThreadLightDump not implemented")
}
func (UnimplementedProfilerCommandServiceServer) mustEmbedUnimplementedProfilerCommandServiceServer() {
}

// UnsafeProfilerCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilerCommandServiceServer will
// result in compilation errors.
type UnsafeProfilerCommandServiceServer interface {
	mustEmbedUnimplementedProfilerCommandServiceServer()
}

func RegisterProfilerCommandServiceServer(s grpc.ServiceRegistrar, srv ProfilerCommandServiceServer) {
	s.RegisterService(&ProfilerCommandService_ServiceDesc, srv)
}

func _ProfilerCommandService_HandleCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilerCommandServiceServer).HandleCommand(&profilerCommandServiceHandleCommandServer{stream})
}

type ProfilerCommandService_HandleCommandServer interface {
	Send(*PCmdRequest) error
	Recv() (*PCmdMessage, error)
	grpc.ServerStream
}

type profilerCommandServiceHandleCommandServer struct {
	grpc.ServerStream
}

func (x *profilerCommandServiceHandleCommandServer) Send(m *PCmdRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandServer) Recv() (*PCmdMessage, error) {
	m := new(PCmdMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfilerCommandService_HandleCommandV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilerCommandServiceServer).HandleCommandV2(&profilerCommandServiceHandleCommandV2Server{stream})
}

type ProfilerCommandService_HandleCommandV2Server interface {
	Send(*PCmdRequest) error
	Recv() (*PCmdMessage, error)
	grpc.ServerStream
}

type profilerCommandServiceHandleCommandV2Server struct {
	grpc.ServerStream
}

func (x *profilerCommandServiceHandleCommandV2Server) Send(m *PCmdRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandV2Server) Recv() (*PCmdMessage, error) {
	m := new(PCmdMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfilerCommandService_CommandEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdEchoResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandEcho(ctx, req.(*PCmdEchoResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerCommandService_CommandStreamActiveThreadCount_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilerCommandServiceServer).CommandStreamActiveThreadCount(&profilerCommandServiceCommandStreamActiveThreadCountServer{stream})
}

type ProfilerCommandService_CommandStreamActiveThreadCountServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PCmdActiveThreadCountRes, error)
	grpc.ServerStream
}

type profilerCommandServiceCommandStreamActiveThreadCountServer struct {
	grpc.ServerStream
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountServer) Recv() (*PCmdActiveThreadCountRes, error) {
	m := new(PCmdActiveThreadCountRes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfilerCommandService_CommandActiveThreadDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdActiveThreadDumpRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandActiveThreadDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadDump(ctx, req.(*PCmdActiveThreadDumpRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerCommandService_CommandActiveThreadLightDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdActiveThreadLightDumpRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadLightDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandActiveThreadLightDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadLightDump(ctx, req.(*PCmdActiveThreadLightDumpRes))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfilerCommandService_ServiceDesc is the grpc.ServiceDesc for ProfilerCommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfilerCommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProfilerCommandService",
	HandlerType: (*ProfilerCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommandEcho",
			Handler:    _ProfilerCommandService_CommandEcho_Handler,
		},
		{
			MethodName: "CommandActiveThreadDump",
			Handler:    _ProfilerCommandService_CommandActiveThreadDump_Handler,
		},
		{
			MethodName: "CommandActiveThreadLightDump",
			Handler:    _ProfilerCommandService_CommandActiveThreadLightDump_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HandleCommand",
			Handler:       _ProfilerCommandService_HandleCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "HandleCommandV2",
			Handler:       _ProfilerCommandService_HandleCommandV2_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CommandStreamActiveThreadCount",
			Handler:       _ProfilerCommandService_CommandStreamActiveThreadCount_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "v1/Service.proto",
}
