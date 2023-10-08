package inbox

import (
	"runtime"
	"sync/atomic"

	"github.com/nevercatchmeow/penumbra/core/tools/constraints"
	"github.com/nevercatchmeow/penumbra/core/tools/container/ring"
)

const (
	IDLE int32 = iota
	RUNNING
)

type Inbox[V any] struct {
	buffer    *ring.Ring[V]
	processor Processor[V]
	scheduler scheduler
	status    int32
}

func NewInbox[T constraints.Signed, V any](initialSize T) *Inbox[V] {
	return &Inbox[V]{
		buffer:    ring.New[V](int64(initialSize)),
		scheduler: scheduler(defaultThroughput),
	}
}

func (slf *Inbox[V]) Start(processor Processor[V]) {
	slf.processor = processor
}

func (slf *Inbox[V]) Stop() error {
	return nil
}

func (slf *Inbox[V]) Push(env V) {
	slf.buffer.Push(env)
	slf.schedule()
}

func (slf *Inbox[V]) schedule() {
	if atomic.CompareAndSwapInt32(&slf.status, IDLE, RUNNING) {
		slf.scheduler.Schedule(slf.process)
	}
}

func (slf *Inbox[V]) process() {
	slf.run()
	atomic.StoreInt32(&slf.status, IDLE)
}

func (slf *Inbox[V]) run() {
	count, throughput := 0, slf.scheduler.Throughput()
	for {
		if count > throughput {
			count = 0
			runtime.Gosched()
		}
		count++
		if envelope, ok := slf.buffer.Pop(); ok {
			slf.processor.Invoke([]V{envelope})
		} else {
			return
		}
	}
}
