package ring_test

import (
	"testing"

	"github.com/nevercatchmeow/penumbra/core/tools/container/ring"
)

func TestRing(t *testing.T) {

	r := ring.New[int64](10)

	for i := int64(0); i < 100; i++ {
		r.Push(i)
	}

	for i := int64(0); i < 100; i++ {
		v, ok := r.Pop()
		if !ok {
			t.Fatal("pop failed")
		}
		if v != i {
			t.Fatal("pop failed")
		}
	}
}
