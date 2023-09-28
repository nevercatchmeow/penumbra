package tcp

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/nevercatchmeow/penumbra/core/network"
	"github.com/nevercatchmeow/penumbra/core/tools/container/dictionary"
	"github.com/nevercatchmeow/penumbra/core/tools/log"
)

type tcpServer struct {
	*options
	*network.BaseEvent
	eng       gnet.Engine
	online    *dictionary.Dictionary[string, network.Conn]
	exitCh    chan os.Signal
	isRunning bool
	codec     *Codec
}

func NewServer(addr string, opts ...Option) network.Server {
	srv := &tcpServer{
		online: dictionary.NewDictionary[string, network.Conn](),
		exitCh: make(chan os.Signal),
	}
	srv.BaseEvent = &network.BaseEvent{Server: srv}
	defaultOpts := defaultOptions(addr)
	for _, opt := range opts {
		opt(defaultOpts)
	}
	srv.options = defaultOpts
	srv.codec = NewCodec(true, 0x1234, srv.maxMsgLen)
	return srv
}

func (slf *tcpServer) Address() string {
	return slf.address
}

func (slf *tcpServer) Protocol() string {
	return slf.protocol
}

func (slf *tcpServer) protoAddr() string {
	return fmt.Sprintf("%s://%s", slf.protocol, slf.address)
}

func (slf *tcpServer) Start() error {
	go func() {
		slf.isRunning = true
		if err := gnet.Run(slf, slf.protoAddr(),
			gnet.WithLogger(log.GetLogger()),
			gnet.WithLogLevel(lo.Ternary(slf.runMode == network.RunModeProd, logging.ErrorLevel, logging.DebugLevel)),
			gnet.WithTicker(true),
			gnet.WithMulticore(true),
		); err != nil {
			slf.isRunning = false
			log.Error("gnet tcp server stopped")
			slf.Stop()
		}
	}()
	slf.OnStartEvent()
	signal.Notify(slf.exitCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-slf.exitCh:
		log.Warn("system exit signal received")
		slf.stop()
	}
	return nil
}

func (slf *tcpServer) Stop() {
	slf.exitCh <- syscall.SIGQUIT
}

func (slf *tcpServer) stop() {
	if slf.isRunning {
		if err := slf.eng.Stop(context.Background()); err != nil {
			log.Error("gnet eng stop failed: ", log.Err(err))
		}
		slf.isRunning = false
	}
}

func (slf *tcpServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	slf.eng = eng
	return
}

func (slf *tcpServer) OnShutdown(eng gnet.Engine) {}

func (slf *tcpServer) OnOpen(gn gnet.Conn) (out []byte, action gnet.Action) {
	if slf.maxConnNum > 0 && slf.online.Len() >= slf.maxConnNum {
		log.Warn("tcp server is full, reject new connection",
			log.Int("current_connections", slf.online.Len()), log.String("remote_addr", gn.RemoteAddr().String()))
		return nil, gnet.Close
	}
	conn := newTcpConn(slf, gn)
	slf.online.Set(gn.RemoteAddr().String(), conn)
	gn.SetContext(conn)
	slf.OnConnectEvent(conn)
	return
}

func (slf *tcpServer) OnClose(gn gnet.Conn, err error) (action gnet.Action) {
	conn, ok := gn.Context().(network.Conn)
	if ok {
		slf.OnDisconnectEvent(conn)
		conn.Close()
		slf.online.Delete(conn.ID())
	}
	return
}

func (slf *tcpServer) OnTraffic(gn gnet.Conn) (action gnet.Action) {
	conn := gn.Context().(network.Conn)
	var packet []byte
loop:
	for {
		data, err := slf.codec.Decode(gn)
		switch {
		case errors.Is(err, ErrIncompletePacket):
			break loop
		case errors.Is(err, ErrTooLargePacket):
			log.Warn("tcp codec decode failed", log.Err(err))
			return
		case err != nil:
			log.Error("tcp codec decode failed", log.Err(err))
			return gnet.Close
		default:
		}
		packet = append(packet, bytes.Clone(data)...)
	}
	slf.OnReceiveEvent(conn, packet)
	return
}

func (slf *tcpServer) OnTick() (delay time.Duration, action gnet.Action) {
	delay = time.Second
	return
}
