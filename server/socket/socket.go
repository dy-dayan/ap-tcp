package socket

import (
	"bufio"
	"net"
	"sync"
	"time"
)

type (
	Socket interface {
		//本地地址
		LocalAddr() net.Addr
		//远程地址
		RemoteAddr() net.Addr
		//设置socket保活时间
		SetDeadline(t time.Time) error
		//设置socket读超时时间
		SetReadDeadline(t time.Time) error
		//设置socket写超时时间
		SetWriteDeadline(t time.Time) error
		//读数据
		Read(b []byte) (n int, err error)
		//写数据
		Write(b []byte) (n int, err error)
		//关闭socket
		Close() error

		//设置链路ID
		SetFid(fid uint64)

		GetFid() uint64
	}

	//tips:fid的读写没有加锁，由一个地方写，其他地方读
	socket struct {
		con              net.Conn
		fid              uint64
		uid              uint64
		readerWithBuffer *bufio.Reader
		mu               sync.RWMutex
	}
)

const READ_SIZE = 4096

func NewSocket(con net.Conn) Socket {
	return newSocket(con)
}

func newSocket(c net.Conn) *socket {
	c.(*net.TCPConn).SetKeepAlive(true)
	return &socket{
		con:              c,
		readerWithBuffer: bufio.NewReaderSize(c, READ_SIZE),
	}
}

//读socket数据·
func (s *socket) Read(b []byte) (int, error) {
	return s.readerWithBuffer.Read(b)
}

//socket写数据
func (s *socket) Write(b []byte) (int, error) {
	return s.con.Write(b)
}

//设置socketlinkid
func (s *socket) SetFid(fid uint64) {
	s.fid = fid
}

//
func (s *socket) GetFid() uint64 {
	return s.fid
}

//设置socketuid 信息
func (s *socket) SetUid(uid uint64) {
	s.uid = uid
}

//获得用户id
func (s *socket) GetUid() uint64 {
	return s.uid
}

//关闭socket
func (s *socket) Close() error {
	return s.con.Close()
}

//远程地址
func (s *socket) RemoteAddr() net.Addr {
	return s.con.RemoteAddr()
}

//本地地址
func (s *socket) LocalAddr() net.Addr {
	return s.con.LocalAddr()
}

//设置socket保活时间
func (s *socket) SetDeadline(t time.Time) error {
	return s.con.SetDeadline(t)
}

//设置socket读超时时间
func (s *socket) SetReadDeadline(t time.Time) error {
	return s.con.SetReadDeadline(t)
}

//设置socket写超时时间
func (s *socket) SetWriteDeadline(t time.Time) error {
	return s.con.SetWriteDeadline(t)
}
