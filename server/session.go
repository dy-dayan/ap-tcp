package server

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/dy-dayan/ap-tcp/server/socket"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io"
	"net"
	"sync"
)

const (
	MSG_READ_SIZE   = 4096
	MSG_BUFFER_SIZE = 10240
)

type Session struct {
	socket        socket.Socket
	uid           uint64
	Authed        bool
	colseNotifyCh chan struct{}
	status        int32
	srv           *TcpServer
}

func NewSession(srv *TcpServer, id uint64, con net.Conn) *Session {
	ss := &Session{
		socket: socket.NewSocket(con),
		uid:    0,
		Authed: false,
		srv:    srv,
	}

	/*
		if err := ss.socket.SetDeadline(time.Now().Add(time.Duration(10 * time.Second)));err != nil{
			return nil
		}
		if err := ss.socket.SetReadDeadline(time.Now().Add(time.Duration(6 * time.Second))); err != nil{
			return nil
		}
		if err := ss.socket.SetWriteDeadline(time.Now().Add(time.Duration(6 * time.Second))); err != nil{
			return nil
		}*/

	ss.socket.SetFid(id)
	return ss
}

func (ss *Session) Id() uint64 {
	return ss.socket.GetFid()
}

func (ss *Session) Close() error {
	return ss.socket.Close()
}

func (ss *Session) StartReadAndHandle() {
	ctx := context.Background()
	msgBuf := bytes.NewBuffer(make([]byte, 0, MSG_BUFFER_SIZE))
	// 数据缓冲
	dataBuf := make([]byte, MSG_READ_SIZE)
	// 消息长度
	length := 0
	// 消息长度uint32
	uLen := uint32(0)
	msgFlag := ""

	for {
		// 读取数据
		n, err := ss.socket.Read(dataBuf)
		if err == io.EOF {
			log.Errorf("Client exit: %s", ss.socket.RemoteAddr())
			goto FAILED
		}

		if err != nil {
			log.Errorf("Read error: %s", err)
			goto FAILED
		}
		// 数据添加到消息缓冲
		n, err = msgBuf.Write(dataBuf[:n])
		if err != nil {
			log.Errorf("Buffer write error: %s", err)
			goto FAILED
		}

		// 消息分割循环
		for {
			// 消息头
			if length == 0 && msgBuf.Len() >= 6 {
				msgFlag = string(msgBuf.Next(2))
				if msgFlag != "DY" {
					log.Error("invalid message")
					goto FAILED
				}
				lengthByte := msgBuf.Next(4)
				uLen = binary.BigEndian.Uint32(lengthByte)
				length = int(uLen)
				// 检查超长消息
				if length > MSG_BUFFER_SIZE {
					log.Errorf("Message too length: %d", length)
					goto FAILED
				}
			}
			// 消息体
			if length > 0 && msgBuf.Len() >= length {
				msg := msgBuf.Next(length)
				length = 0
				ss.HandleMsg(ctx, msg)
			} else {
				break
			}
		}
	}

FAILED:
	ss.srv.sessionHub.Delete(ss.socket.GetFid())
	return

}

func (ss *Session) HandleMsg(ctx context.Context, msg []byte) {
	ss.srv.opt.handleRequest(ctx, ss, msg)
}

func (ss *Session) WriteMsg(msg []byte) error {
	writeBuff := bufio.NewWriter(ss.socket)
	msgLenByte := make([]byte, 4)
	msgLen := len(msg)
	binary.BigEndian.PutUint32(msgLenByte, uint32(msgLen))
	writeBuff.Write([]byte("DY"))
	writeBuff.Write(msgLenByte)
	writeBuff.Write(msg)
	writeBuff.Flush()
	return nil
}

type SessionHub struct {
	sessions sync.Map
}

func NewSessionHub() *SessionHub {
	return &SessionHub{}
}

//添加一个socket
func (sh *SessionHub) Add(ss *Session) {
	_session, loaded := sh.sessions.LoadOrStore(ss.Id(), ss)
	if !loaded {
		return
	}
	sh.sessions.Store(ss.Id(), ss)
	if oldSession := _session.(*Session); ss != oldSession {
		oldSession.Close()
	}

}

func (sh *SessionHub) Delete(id uint64) {
	_ss, loaded := sh.sessions.Load(id)
	if !loaded {
		return
	} else {
		_ss.(*Session).Close()
		sh.sessions.Delete(id)
	}
}

func (sh *SessionHub) Get(id uint64) (*Session, bool) {
	_ss, loaded := sh.sessions.Load(id)
	if !loaded {
		return nil, false
	}
	return _ss.(*Session), true
}

func (sh *SessionHub) Range(f func(ss *Session) bool) {
	sh.sessions.Range(func(key, value interface{}) bool {
		return f(value.(*Session))
	})
}
