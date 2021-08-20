// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*Article, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error)
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error)
	SearchArticle(ctx context.Context, in *SearchArticleRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type articlesClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesClient(cc grpc.ClientConnInterface) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/protofiles.Articles/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/protofiles.Articles/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error) {
	out := new(ListArticlesResponse)
	err := c.cc.Invoke(ctx, "/protofiles.Articles/ListArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) SearchArticle(ctx context.Context, in *SearchArticleRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error) {
	out := new(ListArticlesResponse)
	err := c.cc.Invoke(ctx, "/protofiles.Articles/SearchArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protofiles.Articles/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
// All implementations must embed UnimplementedArticlesServer
// for forward compatibility
type ArticlesServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*Article, error)
	GetArticle(context.Context, *GetArticleRequest) (*Article, error)
	ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error)
	SearchArticle(context.Context, *SearchArticleRequest) (*ListArticlesResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedArticlesServer()
}

// UnimplementedArticlesServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesServer struct {
}

func (UnimplementedArticlesServer) CreateArticle(context.Context, *CreateArticleRequest) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticlesServer) GetArticle(context.Context, *GetArticleRequest) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticlesServer) ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedArticlesServer) SearchArticle(context.Context, *SearchArticleRequest) (*ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArticle not implemented")
}
func (UnimplementedArticlesServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticlesServer) mustEmbedUnimplementedArticlesServer() {}

// UnsafeArticlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesServer will
// result in compilation errors.
type UnsafeArticlesServer interface {
	mustEmbedUnimplementedArticlesServer()
}

func RegisterArticlesServer(s grpc.ServiceRegistrar, srv ArticlesServer) {
	s.RegisterService(&Articles_ServiceDesc, srv)
}

func _Articles_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Articles/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Articles/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Articles/ListArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_SearchArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).SearchArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Articles/SearchArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).SearchArticle(ctx, req.(*SearchArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Articles/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Articles_ServiceDesc is the grpc.ServiceDesc for Articles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Articles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _Articles_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _Articles_GetArticle_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _Articles_ListArticles_Handler,
		},
		{
			MethodName: "SearchArticle",
			Handler:    _Articles_SearchArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Articles_DeleteArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/inshorts.proto",
}