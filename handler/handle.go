package handler

import (
	"context"
	"errors"
	"github.com/dy-dayan/ap-tcp/help"
	"github.com/dy-dayan/ap-tcp/idl"
	"github.com/dy-dayan/ap-tcp/server"
	"github.com/dy-gopkg/kit"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"os"
)

type Handler struct {
	tcpSrv *server.TcpServer
}

func NewHandler() *Handler {
	h := &Handler{}
	log.Debug("external listen addr is :", kit.ServiceExternAddr())
	h.tcpSrv = server.NewTcpServer(
		server.Addr(kit.ServiceExternAddr()),
		server.SrvId(kit.ServiceMetadata("id", "0")),
		server.HandleRequest(h.HandleRequest),
	)
	return h
}

func (h *Handler) Start() {
	if err := h.tcpSrv.Run(); err != nil {
		log.Fatalf("run tcp server failed(err:%v)", err)
	}
}

func (h *Handler) Push(ctx context.Context, req *access.PushReq, rsp *access.PushRsp) error {
	rsp.Code = 0
	// TODO: check uid:session is valid first?

	msg := &access.PkgRsp{
		Head: &access.PkgRspHead{Seq: req.Seq},
		Body: &access.PkgRspBody{
			Head: &access.RspHead{
				Uid:  req.Uid,
				Code: 0,
			},
			Bodys: []*access.RspBody{
				&access.RspBody{
					Service: req.Service,
					Method:  req.Method,
					Content: req.Content,
				},
			},
		},
	}

	err := h.PushMsg(req.Uid, msg)
	if err != nil {
		rsp.Code = 1
		return err
	}
	return nil
}

func (h *Handler) HandleRequest(ctx context.Context, ses *server.Session, body []byte) error {
	req := &access.PkgReq{}
	err := proto.Unmarshal(body, req)
	if err != nil {
		log.Errorf("gid:[%d] PkgReq Unmarshal failed(err:%v), %v", helpFunc.GetGID(), err, body)
		os.Exit(-1)
		return err
	}

	// TODO: need close session?
	if req.Head == nil || req.Body == nil {
		log.Errorf("invalid request")
		return errors.New("invalid request")
	}
	rsp := &access.PkgRsp{
		Head: &access.PkgRspHead{Seq: req.Head.Seq},
		Body: &access.PkgRspBody{
			Head: &access.RspHead{},
		},
	}
	defer h.Response(ctx, ses, rsp)

	c := kit.Client()

	if ses.Authed {
		// session has authenticated
		for _, subReq := range req.Body.Bodys {
			//out, err := h.RawCallMicroService(ctx, subReq.Service, subReq.Method, subReq.Content)
			reqTmp := &access.Message{}
			rspTmp := &access.Message{}
			if len(subReq.Content) > 0 {
				reqTmp = access.NewMessage(subReq.Content)
			}
			request := c.NewRequest(subReq.Service, subReq.Method, reqTmp)
			if err = c.Call(ctx, request, rspTmp); err != nil {
				continue
			}
			rspByte, _ := rspTmp.Marshal()
			rspBody := &access.RspBody{
				Service: subReq.Service,
				Method:  subReq.Method,
				Content: rspByte,
				Code:    0,
			}
			if err != nil {
				rspBody.Code = 1
			}
			rsp.Body.Bodys = append(rsp.Body.Bodys, rspBody)
		}

		return nil
	} else {
		if req.Body.Head.Account != nil {
			// TODO: authenticate account
		} else {
			rsp.Body.Head.Code = 1 // failed
			return nil
		}
	}
	return nil
}

func (h *Handler) Response(ctx context.Context, ses *server.Session, msg *access.PkgRsp) error {
	byt, err := proto.Marshal(msg)
	if err != nil {
		log.Errorf("PkgRsp Marshal failed(err:%v)", err)
		return err
	}

	return ses.WriteMsg(byt)
}

func (h *Handler) PushMsg(uid uint64, msg *access.PkgRsp) error {
	byt, err := proto.Marshal(msg)
	if err != nil {
		log.Errorf("PkgRsp Marshal failed(err:%v)", err)
		return err
	}

	return h.tcpSrv.SendMsgByUid(uid, byt)
}

//func (h *Handler) RawCallMicroService(ctx context.Context, service, method string, in []byte, opts ...client.CallOption) (out []byte, err error) {
//	c := client.NewClient(client.Codec("byte-rpc", byterpc.NewCodec))
//	req := c.NewRequest(service, method, in, client.WithContentType("byte-rpc"))
//	out = []byte{}
//
//	err = c.Call(ctx, req, &out, opts...)
//	if err != nil {
//		return nil, err
//	}
//	return out, nil
//}
