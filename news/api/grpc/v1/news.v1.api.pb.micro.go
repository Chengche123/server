// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: news.v1.api.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for NewsService service

func NewNewsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NewsService service

type NewsService interface {
	ListNewsCategoryDetail(ctx context.Context, in *ListNewsCategoryDetailRequest, opts ...client.CallOption) (*ListNewsCategoryDetailResponse, error)
	ListNewsCategory(ctx context.Context, in *ListNewsCategoryRequest, opts ...client.CallOption) (*ListNewsCategoryResponse, error)
}

type newsService struct {
	c    client.Client
	name string
}

func NewNewsService(name string, c client.Client) NewsService {
	return &newsService{
		c:    c,
		name: name,
	}
}

func (c *newsService) ListNewsCategoryDetail(ctx context.Context, in *ListNewsCategoryDetailRequest, opts ...client.CallOption) (*ListNewsCategoryDetailResponse, error) {
	req := c.c.NewRequest(c.name, "NewsService.ListNewsCategoryDetail", in)
	out := new(ListNewsCategoryDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsService) ListNewsCategory(ctx context.Context, in *ListNewsCategoryRequest, opts ...client.CallOption) (*ListNewsCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "NewsService.ListNewsCategory", in)
	out := new(ListNewsCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NewsService service

type NewsServiceHandler interface {
	ListNewsCategoryDetail(context.Context, *ListNewsCategoryDetailRequest, *ListNewsCategoryDetailResponse) error
	ListNewsCategory(context.Context, *ListNewsCategoryRequest, *ListNewsCategoryResponse) error
}

func RegisterNewsServiceHandler(s server.Server, hdlr NewsServiceHandler, opts ...server.HandlerOption) error {
	type newsService interface {
		ListNewsCategoryDetail(ctx context.Context, in *ListNewsCategoryDetailRequest, out *ListNewsCategoryDetailResponse) error
		ListNewsCategory(ctx context.Context, in *ListNewsCategoryRequest, out *ListNewsCategoryResponse) error
	}
	type NewsService struct {
		newsService
	}
	h := &newsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&NewsService{h}, opts...))
}

type newsServiceHandler struct {
	NewsServiceHandler
}

func (h *newsServiceHandler) ListNewsCategoryDetail(ctx context.Context, in *ListNewsCategoryDetailRequest, out *ListNewsCategoryDetailResponse) error {
	return h.NewsServiceHandler.ListNewsCategoryDetail(ctx, in, out)
}

func (h *newsServiceHandler) ListNewsCategory(ctx context.Context, in *ListNewsCategoryRequest, out *ListNewsCategoryResponse) error {
	return h.NewsServiceHandler.ListNewsCategory(ctx, in, out)
}
