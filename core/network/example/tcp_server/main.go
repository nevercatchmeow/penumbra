package main

import (
	"github.com/nevercatchmeow/penumbra/core/network"
	"github.com/nevercatchmeow/penumbra/core/network/tcp"
	"github.com/nevercatchmeow/penumbra/core/tools/log"
)

func main() {
	srv := tcp.NewServer(":2359", tcp.WithRunMode(network.RunModeDev), tcp.WithMaxConnNum(1024), tcp.WithMaxMsgLen(1024))
	srv.RegStartEvent(func(srv network.Server) {
		log.Info("server started", log.String("protocol", srv.Protocol()), log.String("address", srv.Address()))
	})
	srv.RegStopEvent(func(srv network.Server) {
		log.Info("server stopped", log.String("protocol", srv.Protocol()), log.String("address", srv.Address()))
	})
	srv.RegConnectEvent(func(conn network.Conn) {
		log.Info("client connected", log.String("address", conn.RemoteAddr().String()))
	})
	srv.RegDisconnectEvent(func(conn network.Conn) {
		log.Info("client disconnected", log.String("address", conn.RemoteAddr().String()))
	})
	srv.RegReceiveEvent(func(conn network.Conn, data []byte) {
		log.Info("received data", log.String("address", conn.RemoteAddr().String()), log.Int("length", len(data)))
		codec := tcp.NewCodec(true, 0x1234, 1024)
		bytes, _ := codec.Encode(data)
		conn.Write(bytes)
	})

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
