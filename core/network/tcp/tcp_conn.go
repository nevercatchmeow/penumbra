package tcp

import (
	"net"
	"strings"
	"sync"

	"github.com/nevercatchmeow/penumbra/core/tools/log"

	"github.com/panjf2000/gnet/v2"

	"github.com/nevercatchmeow/penumbra/core/network"
)

type tcpConn struct {
	server     network.Server
	gn         gnet.Conn
	remoteAddr net.Addr
	ip         string
	mu         sync.Mutex
	isClosed   bool
}

func newTcpConn(server network.Server, gn gnet.Conn) network.Conn {
	conn := &tcpConn{
		server:     server,
		gn:         gn,
		remoteAddr: gn.RemoteAddr(),
		ip:         gn.RemoteAddr().String(),
	}
	if index := strings.LastIndex(conn.ip, ":"); index != -1 {
		conn.ip = conn.ip[0:index]
	}

	return conn
}

func (slf *tcpConn) ID() string {
	return slf.remoteAddr.String()
}

func (slf *tcpConn) RemoteAddr() net.Addr {
	return slf.remoteAddr
}

func (slf *tcpConn) IP() string {
	return slf.ip
}

func (slf *tcpConn) Server() network.Server {
	return slf.server
}

func (slf *tcpConn) Close() {
	slf.mu.Lock()
	if slf.isClosed {
		slf.mu.Unlock()
		return
	}
	slf.isClosed = true
	_ = slf.gn.Close()
	slf.mu.Unlock()
}

func (slf *tcpConn) Write(data []byte) {
	if _, err := slf.gn.Write(data); err != nil {
		log.Warn("tcp conn write failed", log.Err(err))
	}
}
