package mpsc

import (
	"sync/atomic"
	"unsafe"
)

type node[T any] struct {
	next *node[T]
	val  T
}

type Queue[T any] struct {
	head, tail *node[T]
}

func New[T any]() *Queue[T] {
	queue := &Queue[T]{}
	stub := &node[T]{}
	queue.head = stub
	queue.tail = stub
	return queue
}

// Push adds x to the back of the queue
// can be safely called from multiple goroutines
func (q *Queue[T]) Push(x T) {
	n := new(node[T])
	n.val = x
	prev := (*node[T])(atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), unsafe.Pointer(n)))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&prev.next)), unsafe.Pointer(n))
}

// Pop removes the item from the front of the queue or nil if the queue is empty
// must be called from a single, consumer goroutine
func (q *Queue[T]) Pop() (T, bool) {
	var t T
	tail := q.tail
	next := (*node[T])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next))))
	if next != nil {
		q.tail = next
		v := next.val
		next.val = t
		return v, true
	}
	return t, false
}

// Empty returns true if the queue is empty
// must be called from a single, consumer goroutine
func (q *Queue[T]) Empty() bool {
	tail := q.tail
	next := (*node[T])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next))))
	return next == nil
}
