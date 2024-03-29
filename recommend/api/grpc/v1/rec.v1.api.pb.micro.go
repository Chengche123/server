// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rec.v1.api.proto

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

// Api Endpoints for RecService service

func NewRecServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RecService service

type RecService interface {
	Rec(ctx context.Context, in *RecReq, opts ...client.CallOption) (*RecResponse, error)
	RelatedRec(ctx context.Context, in *RecReq, opts ...client.CallOption) (*RecResponse, error)
	AuthorRec(ctx context.Context, in *RecReq, opts ...client.CallOption) (*RecResponse, error)
}

type recService struct {
	c    client.Client
	name string
}

func NewRecService(name string, c client.Client) RecService {
	return &recService{
		c:    c,
		name: name,
	}
}

func (c *recService) Rec(ctx context.Context, in *RecReq, opts ...client.CallOption) (*RecResponse, error) {
	req := c.c.NewRequest(c.name, "RecService.Rec", in)
	out := new(RecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recService) RelatedRec(ctx context.Context, in *RecReq, opts ...client.CallOption) (*RecResponse, error) {
	req := c.c.NewRequest(c.name, "RecService.RelatedRec", in)
	out := new(RecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recService) AuthorRec(ctx context.Context, in *RecReq, opts ...client.CallOption) (*RecResponse, error) {
	req := c.c.NewRequest(c.name, "RecService.AuthorRec", in)
	out := new(RecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RecService service

type RecServiceHandler interface {
	Rec(context.Context, *RecReq, *RecResponse) error
	RelatedRec(context.Context, *RecReq, *RecResponse) error
	AuthorRec(context.Context, *RecReq, *RecResponse) error
}

func RegisterRecServiceHandler(s server.Server, hdlr RecServiceHandler, opts ...server.HandlerOption) error {
	type recService interface {
		Rec(ctx context.Context, in *RecReq, out *RecResponse) error
		RelatedRec(ctx context.Context, in *RecReq, out *RecResponse) error
		AuthorRec(ctx context.Context, in *RecReq, out *RecResponse) error
	}
	type RecService struct {
		recService
	}
	h := &recServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RecService{h}, opts...))
}

type recServiceHandler struct {
	RecServiceHandler
}

func (h *recServiceHandler) Rec(ctx context.Context, in *RecReq, out *RecResponse) error {
	return h.RecServiceHandler.Rec(ctx, in, out)
}

func (h *recServiceHandler) RelatedRec(ctx context.Context, in *RecReq, out *RecResponse) error {
	return h.RecServiceHandler.RelatedRec(ctx, in, out)
}

func (h *recServiceHandler) AuthorRec(ctx context.Context, in *RecReq, out *RecResponse) error {
	return h.RecServiceHandler.AuthorRec(ctx, in, out)
}
