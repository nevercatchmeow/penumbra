package main

import (
	"net"
	"time"

	"github.com/nevercatchmeow/penumbra/core/tools/log"

	"github.com/nevercatchmeow/penumbra/core/network/tcp"
)

func main() {
	conn, err := net.Dial("tcp", ":2359")
	if err != nil {
		panic(err)
	}

	closeCh := make(chan struct{})
	codec := tcp.NewCodec(true, 0x1234, 1024)

	go func() {
		for {
			buf := make([]byte, 1030)
			n, err := conn.Read(buf)
			if err != nil {
				closeCh <- struct{}{}
				break
			}
			data, err := codec.Unpack(buf[:n])
			if err != nil {
				log.Error("unpack error", log.Err(err))
				continue
			}
			log.Info("received data", log.ByteString("data", data))
		}
	}()
	data, _ := codec.Encode([]byte("hello"))
	ticker := time.NewTicker(5 * time.Second)
loop:
	for {
		select {
		case <-ticker.C:
			if _, err := conn.Write(data); err != nil {
				closeCh <- struct{}{}
			}
			log.Info("sent data", log.ByteString("data", data))
		case <-closeCh:
			break loop
		}
	}
}
