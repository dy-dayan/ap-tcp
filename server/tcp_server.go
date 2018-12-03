package server

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Options struct {
	writeTimeout    int
	readTimeout     int
	deadlineTimeout int
	srvId           string //当前服务实例id
	handleRequest   func(context.Context, *Session, []byte) error
	addr            string
}

var (
	DefaultAddr = "0.0.0.0:9999"
)

type Option func(*Options)

func DeadlineTimeOut(t int) Option {
	return func(o *Options) {
		o.deadlineTimeout = t
	}
}

func ReadTimeout(t int) Option {
	return func(o *Options) {
		o.readTimeout = t
	}
}

func WriteTimeout(t int) Option {
	return func(o *Options) {
		o.writeTimeout = t
	}
}

func SrvId(id string) Option {
	return func(o *Options) {
		o.srvId = id
	}
}

func HandleRequest(f func(context.Context, *Session, []byte) error) Option {
	return func(o *Options) {
		o.handleRequest = f
	}
}

func Addr(addr string) Option {
	return func(o *Options) {
		o.addr = addr
	}
}

type TcpServer struct {
	uid2Sid    sync.Map
	opt        Options
	listener   net.Listener
	wg         sync.WaitGroup
	sessionHub SessionHub
	sockCnt    uint32 //当前socket 计数
	baseValue  uint64 //会话id的基础值
	closeCh    chan struct{}
}

func NewTcpServer(op ...Option) *TcpServer {
	s := &TcpServer{
		sockCnt: 0,
	}

	for _, o := range op {
		o(&(s.opt))
	}

	if s.opt.addr == "" {
		s.opt.addr = DefaultAddr
	}
	id, _ := strconv.ParseUint(s.opt.srvId, 10, 64)
	s.baseValue = id << 32
	return s
}

// ErrListenClosed listener is closed error.
var ErrListenClosed = errors.New("listener is closed")

func (s *TcpServer) Run() error {
	var err error

	s.listener, err = net.Listen("tcp", s.opt.addr)
	if err != nil {
		log.Errorf("listen failed:%v", err)
	}
	var (
		tempDelay time.Duration // how long to sleep on accept failure
		closeCh   = s.closeCh
	)
	for {
		conn, e := s.listener.Accept()
		if e != nil {
			select {
			case <-closeCh:
				return ErrListenClosed
			default:
			}
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		go s.acceptCon(conn)
	}

	return nil
}

//接受客户端连接
func (s *TcpServer) acceptCon(con net.Conn) {
	id := atomic.AddUint32(&s.sockCnt, 1)
	tmp := s.baseValue + uint64(id)
	var ses = NewSession(s, tmp, con)
	if ses == nil{
		log.Debug("create session failed")
	}
	log.Debugf("accept new con from",con.RemoteAddr().String())
	s.sessionHub.Add(ses)
	ses.StartReadAndHandle()
}

func (s *TcpServer) Stop() error {
	var (
		count int
	)
	s.sessionHub.Range(func(ses *Session) bool {
		count++
		ses.Close()
		return true
	})
	return nil
}

//向指定uid发送消息
func (s *TcpServer) SendMsgByUid(uid uint64, msg []byte) error {
	sid, loaded := s.uid2Sid.Load(uid)
	if loaded {
		return s.SendMsgBySid(sid.(uint64), msg)
	}
	return errors.New("not find uid :" + strconv.FormatUint(uid, 10))
}

//向指定会话id发送消息
func (s *TcpServer) SendMsgBySid(sid uint64, msg []byte) error {
	_ss, loaded := s.sessionHub.Get(sid)
	if loaded {
		return _ss.WriteMsg(msg)
	}
	return errors.New("not find fid" + strconv.FormatUint(sid, 10))
}
