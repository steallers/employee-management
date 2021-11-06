// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: pb/employee.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	AddNewEmployee(ctx context.Context, in *AddNewEmployeeRequest, opts ...grpc.CallOption) (*AddNewEmployeeResponses, error)
	RemoveEmployee(ctx context.Context, in *RemoveEmployeeRequest, opts ...grpc.CallOption) (*RemoveEmployeeResponse, error)
	UpdateEmployeeData(ctx context.Context, in *UpdateEmployeeDataRequest, opts ...grpc.CallOption) (*UpdateEmployeeDataResponses, error)
	UpdateEmployeeProfilePicture(ctx context.Context, in *UpdateEmployeeProfilePictureRequest, opts ...grpc.CallOption) (*UpdateEmployeeProfilePictureResponses, error)
	// employee contacts
	AddEmployeeContact(ctx context.Context, in *AddEmployeeContactRequest, opts ...grpc.CallOption) (*AddEmployeeContactResponses, error)
	RemoveEmployeeContact(ctx context.Context, in *RemoveEmployeeContactRequest, opts ...grpc.CallOption) (*RemoveEmployeeContactResponses, error)
	UpdateEmployeeContact(ctx context.Context, in *UpdateEmployeeContactRequest, opts ...grpc.CallOption) (*UpdateEmployeeContactResponses, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) AddNewEmployee(ctx context.Context, in *AddNewEmployeeRequest, opts ...grpc.CallOption) (*AddNewEmployeeResponses, error) {
	out := new(AddNewEmployeeResponses)
	err := c.cc.Invoke(ctx, "/connectors.EmployeeService/AddNewEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) RemoveEmployee(ctx context.Context, in *RemoveEmployeeRequest, opts ...grpc.CallOption) (*RemoveEmployeeResponse, error) {
	out := new(RemoveEmployeeResponse)
	err := c.cc.Invoke(ctx, "/connectors.EmployeeService/RemoveEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployeeData(ctx context.Context, in *UpdateEmployeeDataRequest, opts ...grpc.CallOption) (*UpdateEmployeeDataResponses, error) {
	out := new(UpdateEmployeeDataResponses)
	err := c.cc.Invoke(ctx, "/connectors.EmployeeService/UpdateEmployeeData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployeeProfilePicture(ctx context.Context, in *UpdateEmployeeProfilePictureRequest, opts ...grpc.CallOption) (*UpdateEmployeeProfilePictureResponses, error) {
	out := new(UpdateEmployeeProfilePictureResponses)
	err := c.cc.Invoke(ctx, "/connectors.EmployeeService/UpdateEmployeeProfilePicture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) AddEmployeeContact(ctx context.Context, in *AddEmployeeContactRequest, opts ...grpc.CallOption) (*AddEmployeeContactResponses, error) {
	out := new(AddEmployeeContactResponses)
	err := c.cc.Invoke(ctx, "/connectors.EmployeeService/AddEmployeeContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) RemoveEmployeeContact(ctx context.Context, in *RemoveEmployeeContactRequest, opts ...grpc.CallOption) (*RemoveEmployeeContactResponses, error) {
	out := new(RemoveEmployeeContactResponses)
	err := c.cc.Invoke(ctx, "/connectors.EmployeeService/RemoveEmployeeContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployeeContact(ctx context.Context, in *UpdateEmployeeContactRequest, opts ...grpc.CallOption) (*UpdateEmployeeContactResponses, error) {
	out := new(UpdateEmployeeContactResponses)
	err := c.cc.Invoke(ctx, "/connectors.EmployeeService/UpdateEmployeeContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	AddNewEmployee(context.Context, *AddNewEmployeeRequest) (*AddNewEmployeeResponses, error)
	RemoveEmployee(context.Context, *RemoveEmployeeRequest) (*RemoveEmployeeResponse, error)
	UpdateEmployeeData(context.Context, *UpdateEmployeeDataRequest) (*UpdateEmployeeDataResponses, error)
	UpdateEmployeeProfilePicture(context.Context, *UpdateEmployeeProfilePictureRequest) (*UpdateEmployeeProfilePictureResponses, error)
	// employee contacts
	AddEmployeeContact(context.Context, *AddEmployeeContactRequest) (*AddEmployeeContactResponses, error)
	RemoveEmployeeContact(context.Context, *RemoveEmployeeContactRequest) (*RemoveEmployeeContactResponses, error)
	UpdateEmployeeContact(context.Context, *UpdateEmployeeContactRequest) (*UpdateEmployeeContactResponses, error)
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) AddNewEmployee(context.Context, *AddNewEmployeeRequest) (*AddNewEmployeeResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) RemoveEmployee(context.Context, *RemoveEmployeeRequest) (*RemoveEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) UpdateEmployeeData(context.Context, *UpdateEmployeeDataRequest) (*UpdateEmployeeDataResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployeeData not implemented")
}
func (UnimplementedEmployeeServiceServer) UpdateEmployeeProfilePicture(context.Context, *UpdateEmployeeProfilePictureRequest) (*UpdateEmployeeProfilePictureResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployeeProfilePicture not implemented")
}
func (UnimplementedEmployeeServiceServer) AddEmployeeContact(context.Context, *AddEmployeeContactRequest) (*AddEmployeeContactResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployeeContact not implemented")
}
func (UnimplementedEmployeeServiceServer) RemoveEmployeeContact(context.Context, *RemoveEmployeeContactRequest) (*RemoveEmployeeContactResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEmployeeContact not implemented")
}
func (UnimplementedEmployeeServiceServer) UpdateEmployeeContact(context.Context, *UpdateEmployeeContactRequest) (*UpdateEmployeeContactResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployeeContact not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_AddNewEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).AddNewEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.EmployeeService/AddNewEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).AddNewEmployee(ctx, req.(*AddNewEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_RemoveEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).RemoveEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.EmployeeService/RemoveEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).RemoveEmployee(ctx, req.(*RemoveEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployeeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployeeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.EmployeeService/UpdateEmployeeData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployeeData(ctx, req.(*UpdateEmployeeDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployeeProfilePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeProfilePictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployeeProfilePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.EmployeeService/UpdateEmployeeProfilePicture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployeeProfilePicture(ctx, req.(*UpdateEmployeeProfilePictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_AddEmployeeContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEmployeeContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).AddEmployeeContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.EmployeeService/AddEmployeeContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).AddEmployeeContact(ctx, req.(*AddEmployeeContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_RemoveEmployeeContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEmployeeContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).RemoveEmployeeContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.EmployeeService/RemoveEmployeeContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).RemoveEmployeeContact(ctx, req.(*RemoveEmployeeContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployeeContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployeeContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.EmployeeService/UpdateEmployeeContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployeeContact(ctx, req.(*UpdateEmployeeContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connectors.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewEmployee",
			Handler:    _EmployeeService_AddNewEmployee_Handler,
		},
		{
			MethodName: "RemoveEmployee",
			Handler:    _EmployeeService_RemoveEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployeeData",
			Handler:    _EmployeeService_UpdateEmployeeData_Handler,
		},
		{
			MethodName: "UpdateEmployeeProfilePicture",
			Handler:    _EmployeeService_UpdateEmployeeProfilePicture_Handler,
		},
		{
			MethodName: "AddEmployeeContact",
			Handler:    _EmployeeService_AddEmployeeContact_Handler,
		},
		{
			MethodName: "RemoveEmployeeContact",
			Handler:    _EmployeeService_RemoveEmployeeContact_Handler,
		},
		{
			MethodName: "UpdateEmployeeContact",
			Handler:    _EmployeeService_UpdateEmployeeContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/employee.proto",
}
