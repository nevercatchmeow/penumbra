package tcp

import "github.com/nevercatchmeow/penumbra/core/network"

type options struct {
	runMode    network.RunMode
	address    string
	protocol   string
	maxMsgLen  uint32
	maxConnNum int
}

type Option func(*options)

func defaultOptions(address string) *options {
	return &options{
		runMode:  network.RunModeDev,
		address:  address,
		protocol: "tcp",
	}
}

func WithRunMode(runMode network.RunMode) Option {
	return func(opts *options) {
		opts.runMode = runMode
	}
}

func WithMaxMsgLen(maxMsgLen uint32) Option {
	return func(opts *options) {
		opts.maxMsgLen = maxMsgLen
	}
}

func WithMaxConnNum(maxConnNum int) Option {
	return func(opts *options) {
		opts.maxConnNum = maxConnNum
	}
}
