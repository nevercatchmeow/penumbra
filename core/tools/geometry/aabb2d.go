package geometry

import "github.com/nevercatchmeow/penumbra/core/tools/constraints"

// AABB2D Axis-Aligned Bounding Box in 2D space.
type AABB2D[T constraints.Signed | constraints.Float] struct {
	Width, Height    T
	Center, Min, Max *Point2D[T]
}

// NewAABB2D creates an axis aligned bounding box in 2D space. It takes the Center and half point.
func NewAABB2D[T constraints.Signed | constraints.Float](center *Point2D[T], width, height T) *AABB2D[T] {
	return &AABB2D[T]{
		Center: center, Width: width, Height: height,
		Min: NewPoint2D[T](center.x-width/2, center.y-height/2),
		Max: NewPoint2D[T](center.x+width/2, center.y+height/2),
	}
}

// Contains checks whether the point provided resides within the axis
// aligned bounding box.
func (slf *AABB2D[T]) Contains(point *Point2D[T]) bool {
	if point == nil {
		return false
	}
	if point.GetX() < slf.Min.GetX() || point.GetY() < slf.Min.GetY() ||
		point.GetX() > slf.Max.GetX() || point.GetY() > slf.Max.GetY() {
		return false
	}
	return true
}

// Intersect checks whether two axis aligned bounding boxes overlap.
func (slf *AABB2D[T]) Intersect(aabb *AABB2D[T]) bool {
	if aabb == nil {
		return false
	}
	if aabb.Max.GetX() < slf.Min.GetX() || aabb.Max.GetY() < slf.Min.GetY() ||
		aabb.Min.GetX() > slf.Max.GetX() || aabb.Min.GetY() > slf.Max.GetY() {
		return false
	}
	return true
}
