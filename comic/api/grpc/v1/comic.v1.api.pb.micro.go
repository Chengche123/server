// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: comic.v1.api.proto

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

// Api Endpoints for ComicService service

func NewComicServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ComicService service

type ComicService interface {
	ListComicDetail(ctx context.Context, in *ListComicDetailRequest, opts ...client.CallOption) (*ListComicDetailResponse, error)
	ListCategoryComicDetail(ctx context.Context, in *ListCategoryComicDetailRequest, opts ...client.CallOption) (*ListCategoryComicDetailResponse, error)
	ListComicCategoryDetail(ctx context.Context, in *ListComicCategoryDetailRequest, opts ...client.CallOption) (*ListComicCategoryDetailResponse, error)
	ListComicCategoryFilter(ctx context.Context, in *ListComicCategoryFilterRequest, opts ...client.CallOption) (*ListComicCategoryFilterResponse, error)
	ListComicSpecial(ctx context.Context, in *ListComicSpecialRequest, opts ...client.CallOption) (*ListComicSpecialResponse, error)
	ListComicChapter(ctx context.Context, in *ListComicChapterRequest, opts ...client.CallOption) (*ListComicChapterResponse, error)
}

type comicService struct {
	c    client.Client
	name string
}

func NewComicService(name string, c client.Client) ComicService {
	return &comicService{
		c:    c,
		name: name,
	}
}

func (c *comicService) ListComicDetail(ctx context.Context, in *ListComicDetailRequest, opts ...client.CallOption) (*ListComicDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ComicService.ListComicDetail", in)
	out := new(ListComicDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicService) ListCategoryComicDetail(ctx context.Context, in *ListCategoryComicDetailRequest, opts ...client.CallOption) (*ListCategoryComicDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ComicService.ListCategoryComicDetail", in)
	out := new(ListCategoryComicDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicService) ListComicCategoryDetail(ctx context.Context, in *ListComicCategoryDetailRequest, opts ...client.CallOption) (*ListComicCategoryDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ComicService.ListComicCategoryDetail", in)
	out := new(ListComicCategoryDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicService) ListComicCategoryFilter(ctx context.Context, in *ListComicCategoryFilterRequest, opts ...client.CallOption) (*ListComicCategoryFilterResponse, error) {
	req := c.c.NewRequest(c.name, "ComicService.ListComicCategoryFilter", in)
	out := new(ListComicCategoryFilterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicService) ListComicSpecial(ctx context.Context, in *ListComicSpecialRequest, opts ...client.CallOption) (*ListComicSpecialResponse, error) {
	req := c.c.NewRequest(c.name, "ComicService.ListComicSpecial", in)
	out := new(ListComicSpecialResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicService) ListComicChapter(ctx context.Context, in *ListComicChapterRequest, opts ...client.CallOption) (*ListComicChapterResponse, error) {
	req := c.c.NewRequest(c.name, "ComicService.ListComicChapter", in)
	out := new(ListComicChapterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ComicService service

type ComicServiceHandler interface {
	ListComicDetail(context.Context, *ListComicDetailRequest, *ListComicDetailResponse) error
	ListCategoryComicDetail(context.Context, *ListCategoryComicDetailRequest, *ListCategoryComicDetailResponse) error
	ListComicCategoryDetail(context.Context, *ListComicCategoryDetailRequest, *ListComicCategoryDetailResponse) error
	ListComicCategoryFilter(context.Context, *ListComicCategoryFilterRequest, *ListComicCategoryFilterResponse) error
	ListComicSpecial(context.Context, *ListComicSpecialRequest, *ListComicSpecialResponse) error
	ListComicChapter(context.Context, *ListComicChapterRequest, *ListComicChapterResponse) error
}

func RegisterComicServiceHandler(s server.Server, hdlr ComicServiceHandler, opts ...server.HandlerOption) error {
	type comicService interface {
		ListComicDetail(ctx context.Context, in *ListComicDetailRequest, out *ListComicDetailResponse) error
		ListCategoryComicDetail(ctx context.Context, in *ListCategoryComicDetailRequest, out *ListCategoryComicDetailResponse) error
		ListComicCategoryDetail(ctx context.Context, in *ListComicCategoryDetailRequest, out *ListComicCategoryDetailResponse) error
		ListComicCategoryFilter(ctx context.Context, in *ListComicCategoryFilterRequest, out *ListComicCategoryFilterResponse) error
		ListComicSpecial(ctx context.Context, in *ListComicSpecialRequest, out *ListComicSpecialResponse) error
		ListComicChapter(ctx context.Context, in *ListComicChapterRequest, out *ListComicChapterResponse) error
	}
	type ComicService struct {
		comicService
	}
	h := &comicServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ComicService{h}, opts...))
}

type comicServiceHandler struct {
	ComicServiceHandler
}

func (h *comicServiceHandler) ListComicDetail(ctx context.Context, in *ListComicDetailRequest, out *ListComicDetailResponse) error {
	return h.ComicServiceHandler.ListComicDetail(ctx, in, out)
}

func (h *comicServiceHandler) ListCategoryComicDetail(ctx context.Context, in *ListCategoryComicDetailRequest, out *ListCategoryComicDetailResponse) error {
	return h.ComicServiceHandler.ListCategoryComicDetail(ctx, in, out)
}

func (h *comicServiceHandler) ListComicCategoryDetail(ctx context.Context, in *ListComicCategoryDetailRequest, out *ListComicCategoryDetailResponse) error {
	return h.ComicServiceHandler.ListComicCategoryDetail(ctx, in, out)
}

func (h *comicServiceHandler) ListComicCategoryFilter(ctx context.Context, in *ListComicCategoryFilterRequest, out *ListComicCategoryFilterResponse) error {
	return h.ComicServiceHandler.ListComicCategoryFilter(ctx, in, out)
}

func (h *comicServiceHandler) ListComicSpecial(ctx context.Context, in *ListComicSpecialRequest, out *ListComicSpecialResponse) error {
	return h.ComicServiceHandler.ListComicSpecial(ctx, in, out)
}

func (h *comicServiceHandler) ListComicChapter(ctx context.Context, in *ListComicChapterRequest, out *ListComicChapterResponse) error {
	return h.ComicServiceHandler.ListComicChapter(ctx, in, out)
}
