package hilbert

import (
	"github.com/Workiva/go-datastructures/rtree"
)

func getCenter(rect rtree.Rectangle) (int32, int32) { _ = "STUB: not implemented"; return 0, 0 }

type hilbertBundle struct {
	hilbert hilbert
	rect    rtree.Rectangle
}

func bundlesFromRects(rects ...rtree.Rectangle) []*hilbertBundle {
	_ = "STUB: not implemented"
	return nil
}

func chunkRectangles(slice rtree.Rectangles, numParts int64) []rtree.Rectangles {
	_ = "STUB: not implemented"
	return nil
}
