package hilbert

import "github.com/Workiva/go-datastructures/rtree"

type rectangle struct {
	xlow, xhigh, ylow, yhigh int32
}

func (r *rectangle) adjust(rect rtree.Rectangle) { _ = "STUB: not implemented"; return }

func equal(r1, r2 rtree.Rectangle) bool { _ = "STUB: not implemented"; return false }

func intersect(rect1 *rectangle, rect2 rtree.Rectangle) bool {
	_ = "STUB: not implemented"
	return false
}

func newRectangeFromRect(rect rtree.Rectangle) *rectangle { _ = "STUB: not implemented"; return nil }

func newRectangleFromRects(rects rtree.Rectangles) *rectangle {
	_ = "STUB: not implemented"
	return nil
}
