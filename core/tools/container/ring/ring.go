package ring

import (
	"sync"
	"sync/atomic"
)

type buffer[T any] struct {
	items           []T
	head, tail, mod int64
}

type Ring[T any] struct {
	length  int64
	content *buffer[T]
	mutex   sync.Mutex
}

func New[T any](size int64) *Ring[T] {
	return &Ring[T]{
		content: &buffer[T]{
			items: make([]T, size),
			mod:   size,
		},
	}
}

func (slf *Ring[T]) Push(item T) {
	slf.mutex.Lock()
	slf.content.tail = (slf.content.tail + 1) % slf.content.mod
	if slf.content.tail == slf.content.head {
		size := slf.content.mod * 2
		newBuff := make([]T, size)
		for i := int64(0); i < slf.content.mod; i++ {
			idx := (slf.content.tail + i) % slf.content.mod
			newBuff[i] = slf.content.items[idx]
		}
		content := &buffer[T]{
			items: newBuff,
			head:  0,
			tail:  slf.content.mod,
			mod:   size,
		}
		slf.content = content
	}
	atomic.AddInt64(&slf.length, 1)
	slf.content.items[slf.content.tail] = item
	slf.mutex.Unlock()
}

func (slf *Ring[T]) Length() int64 {
	return atomic.LoadInt64(&slf.length)
}

func (slf *Ring[T]) Empty() bool {
	return slf.Length() == 0
}

func (slf *Ring[T]) Pop() (T, bool) {
	if slf.Empty() {
		var t T
		return t, false
	}
	slf.mutex.Lock()
	slf.content.head = (slf.content.head + 1) % slf.content.mod
	item := slf.content.items[slf.content.head]
	var t T
	slf.content.items[slf.content.head] = t
	atomic.AddInt64(&slf.length, -1)
	slf.mutex.Unlock()
	return item, true
}

func (slf *Ring[T]) PopN(n int64) ([]T, bool) {
	if slf.Empty() {
		return nil, false
	}
	slf.mutex.Lock()
	content := slf.content
	if n >= slf.length {
		n = slf.length
	}
	atomic.AddInt64(&slf.length, -n)
	items := make([]T, n)
	for i := int64(0); i < n; i++ {
		pos := (content.head + 1 + i) % content.mod
		items[i] = content.items[pos]
		var t T
		content.items[pos] = t
	}
	content.head = (content.head + n) % content.mod
	slf.mutex.Unlock()
	return items, true
}
