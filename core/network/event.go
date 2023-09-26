package network

type (
	StartEventHandler      func(srv Server)
	StopEventHandler       func(srv Server)
	ConnectEventHandler    func(conn Conn)
	DisconnectEventHandler func(conn Conn)
	ReceiveEventHandler    func(conn Conn, data []byte)
)

type Event interface {
	RegStartEvent(handler StartEventHandler)
	OnStartEvent()
	RegStopEvent(handler StopEventHandler)
	OnStopEvent()
	RegConnectEvent(handler ConnectEventHandler)
	OnConnectEvent(conn Conn)
	RegDisconnectEvent(handler DisconnectEventHandler)
	OnDisconnectEvent(conn Conn)
	RegReceiveEvent(handler ReceiveEventHandler)
	OnReceiveEvent(conn Conn, data []byte)
}

type BaseEvent struct {
	Server
	StartEvent      StartEventHandler
	StopEvent       StopEventHandler
	ConnectEvent    ConnectEventHandler
	DisconnectEvent DisconnectEventHandler
	ReceiveEvent    ReceiveEventHandler
}

func (slf *BaseEvent) RegStartEvent(handler StartEventHandler) {
	slf.StartEvent = handler
}

func (slf *BaseEvent) OnStartEvent() {
	if slf.StartEvent != nil {
		slf.StartEvent(slf.Server)
	}
}

func (slf *BaseEvent) RegStopEvent(handler StopEventHandler) {
	slf.StopEvent = handler
}

func (slf *BaseEvent) OnStopEvent() {
	if slf.StopEvent != nil {
		slf.StopEvent(slf.Server)
	}
}

func (slf *BaseEvent) RegConnectEvent(handler ConnectEventHandler) {
	slf.ConnectEvent = handler
}

func (slf *BaseEvent) OnConnectEvent(conn Conn) {
	if slf.ConnectEvent != nil {
		slf.ConnectEvent(conn)
	}
}

func (slf *BaseEvent) RegDisconnectEvent(handler DisconnectEventHandler) {
	slf.DisconnectEvent = handler
}

func (slf *BaseEvent) OnDisconnectEvent(conn Conn) {
	if slf.DisconnectEvent != nil {
		slf.DisconnectEvent(conn)
	}
}

func (slf *BaseEvent) RegReceiveEvent(handler ReceiveEventHandler) {
	slf.ReceiveEvent = handler
}

func (slf *BaseEvent) OnReceiveEvent(conn Conn, data []byte) {
	if slf.ReceiveEvent != nil {
		slf.ReceiveEvent(conn, data)
	}
}
