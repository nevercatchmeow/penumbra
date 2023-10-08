package geometry

import (
	"fmt"
	"github.com/nevercatchmeow/penumbra/core/tools/constraints"
)

type Point2D[T constraints.Signed | constraints.Float] struct {
	Vector2[T]
}

func NewPoint2D[T constraints.Signed | constraints.Float](x, y T) *Point2D[T] {
	return &Point2D[T]{Vector2[T]{x, y}}
}

func (slf *Point2D[T]) Coordinates() (T, T) {
	return slf.x, slf.y
}

func (slf *Point2D[T]) String() string {
	return fmt.Sprintf("(%v, %v)", slf.x, slf.y)
}
