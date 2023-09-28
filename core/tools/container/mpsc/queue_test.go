package mpsc_test

import (
	"sync/atomic"
	"testing"

	"github.com/nevercatchmeow/penumbra/core/tools/container/mpsc"
)

func TestQueue(t *testing.T) {

	type Message struct {
		Number int
	}

	queue := mpsc.New[*Message]()
	for i := 0; i < 3; i++ {
		base := i + 1
		go func() {
			for j := (base - 1) * 10; j < base*10; j++ {
				queue.Push(&Message{Number: j})
			}
		}()
	}

	var count int32
	for {
		val, success := queue.Pop()
		if !success {
			continue
		}
		t.Log(val)
		if atomic.AddInt32(&count, 1) == 30 {
			break
		}
	}
}
