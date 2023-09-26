package network

type Server interface {
	Event
	Address() string
	Protocol() string
	Start() error
	Stop()
}
