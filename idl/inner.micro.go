// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: inner.proto

/*
Package access is a generated protocol buffer package.

It is generated from these files:
	inner.proto

It has these top-level messages:
	PushReq
	PushRsp
*/
package access

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Access service

type AccessService interface {
	Push(ctx context.Context, in *PushReq, opts ...client.CallOption) (*PushRsp, error)
}

type accessService struct {
	c    client.Client
	name string
}

func NewAccessService(name string, c client.Client) AccessService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "access"
	}
	return &accessService{
		c:    c,
		name: name,
	}
}

func (c *accessService) Push(ctx context.Context, in *PushReq, opts ...client.CallOption) (*PushRsp, error) {
	req := c.c.NewRequest(c.name, "Access.Push", in)
	out := new(PushRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Access service

type AccessHandler interface {
	Push(context.Context, *PushReq, *PushRsp) error
}

func RegisterAccessHandler(s server.Server, hdlr AccessHandler, opts ...server.HandlerOption) error {
	type access interface {
		Push(ctx context.Context, in *PushReq, out *PushRsp) error
	}
	type Access struct {
		access
	}
	h := &accessHandler{hdlr}
	return s.Handle(s.NewHandler(&Access{h}, opts...))
}

type accessHandler struct {
	AccessHandler
}

func (h *accessHandler) Push(ctx context.Context, in *PushReq, out *PushRsp) error {
	return h.AccessHandler.Push(ctx, in, out)
}
