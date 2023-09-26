package network

import "net"

type Conn interface {
	ID() string
	RemoteAddr() net.Addr
	IP() string
	Server() Server
	Close()
	Write([]byte)
}
