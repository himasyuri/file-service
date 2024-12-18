// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: file_service.proto

package fileservicepb

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
	FileService_UploadFile_FullMethodName   = "/fileservice.FileService/UploadFile"
	FileService_ListFiles_FullMethodName    = "/fileservice.FileService/ListFiles"
	FileService_DownloadFile_FullMethodName = "/fileservice.FileService/DownloadFile"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	// Uploads a file in chunks
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadFileRequest, UploadFileResponse], error)
	// Lists all files with their metadata
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	// Downloads a file in chunks
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileChunk], error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadFileRequest, UploadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], FileService_UploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadFileRequest, UploadFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadFileClient = grpc.ClientStreamingClient[UploadFileRequest, UploadFileResponse]

func (c *fileServiceClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, FileService_ListFiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[1], FileService_DownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DownloadFileRequest, FileChunk]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_DownloadFileClient = grpc.ServerStreamingClient[FileChunk]

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
type FileServiceServer interface {
	// Uploads a file in chunks
	UploadFile(grpc.ClientStreamingServer[UploadFileRequest, UploadFileResponse]) error
	// Lists all files with their metadata
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	// Downloads a file in chunks
	DownloadFile(*DownloadFileRequest, grpc.ServerStreamingServer[FileChunk]) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) UploadFile(grpc.ClientStreamingServer[UploadFileRequest, UploadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileServiceServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileServiceServer) DownloadFile(*DownloadFileRequest, grpc.ServerStreamingServer[FileChunk]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadFile(&grpc.GenericServerStream[UploadFileRequest, UploadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadFileServer = grpc.ClientStreamingServer[UploadFileRequest, UploadFileResponse]

func _FileService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_ListFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadFile(m, &grpc.GenericServerStream[DownloadFileRequest, FileChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_DownloadFileServer = grpc.ServerStreamingServer[FileChunk]

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fileservice.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FileService_ListFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileService_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file_service.proto",
}