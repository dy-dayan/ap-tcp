package socket

import "sync"

type SocketHub struct {
	fid2Socket sync.Map
}

func NewSocketHub() *SocketHub {
	return &SocketHub{}
}

//添加一个socket
func (sh *SocketHub) Add(socket Socket) {
	//查看是否已经存在socket
	_socket, loaded := sh.fid2Socket.LoadOrStore(socket.GetFid(), socket)
	//不存在，加加进去，并返回

	if !loaded {
		return
	}
	//存入新的socket
	sh.fid2Socket.Store(socket.GetFid(), socket)
	//新旧socket不一样的时候，关闭老的socket
	if oldSocket := _socket.(Socket); socket != oldSocket {
		oldSocket.Close()
	}

}

//删掉一个socket
func (sh *SocketHub) Delete(fid uint64) {
	//查看是否已经存在socket
	_socket, loaded := sh.fid2Socket.Load(fid)
	if !loaded {
		return
	} else {
		_socket.(Socket).Close()
		sh.fid2Socket.Delete(fid)
	}
}

//获得一个socket
func (sh *SocketHub) Get(fid uint64) (Socket, bool) {
	_socket, loaded := sh.fid2Socket.Load(fid)
	if !loaded {
		return nil, false
	}
	return _socket.(Socket), true
}

func (sh *SocketHub) Range(f func(Socket) bool) {
	sh.fid2Socket.Range(func(key, value interface{}) bool {
		return f(value.(Socket))
	})
}
