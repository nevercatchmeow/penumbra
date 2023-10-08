package geometry

import "github.com/nevercatchmeow/penumbra/core/tools/constraints"

type Vector2[T constraints.Signed | constraints.Float] struct {
	x T
	y T
}

func (slf *Vector2[T]) SetX(x T) {
	slf.x = x
}

func (slf *Vector2[T]) SetY(y T) {
	slf.y = y
}

func (slf *Vector2[T]) GetX() T {
	return slf.x
}

func (slf *Vector2[T]) GetY() T {
	return slf.y
}
