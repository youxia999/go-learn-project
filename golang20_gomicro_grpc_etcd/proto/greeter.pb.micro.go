// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: greetercenter/greeter.greetercenter

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the greetercenter package it is being compiled against.
// A compilation error at this line likely means your copy of the
// greetercenter package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the greetercenter package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterService interface {
	Helloservice(ctx context.Context, in *RequestMessage, opts ...client.CallOption) (*ResponseMessage, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "protocol"
	}
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Helloservice(ctx context.Context, in *RequestMessage, opts ...client.CallOption) (*ResponseMessage, error) {
	req := c.c.NewRequest(c.name, "Greeter.Helloservice", in)
	out := new(ResponseMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Helloservice(context.Context, *RequestMessage, *ResponseMessage) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Helloservice(ctx context.Context, in *RequestMessage, out *ResponseMessage) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Helloservice(ctx context.Context, in *RequestMessage, out *ResponseMessage) error {
	return h.GreeterHandler.Helloservice(ctx, in, out)
}
