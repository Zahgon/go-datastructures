package augmentedtree

func intervalOverlaps(n *node, low, high int64, interval Interval, maxDimension uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func overlaps(high, otherHigh, low, otherLow int64) bool { _ = "STUB: not implemented"; return false }

func compare(nodeLow, ivLow int64, nodeID, ivID uint64) int { _ = "STUB: not implemented"; return 0 }

type node struct {
	interval Interval
	max, min int64
	children [2]*node
	red      bool
	id       uint64
}

func (n *node) query(low, high int64, interval Interval, maxDimension uint64, fn func(node *node)) {
	_ = "STUB: not implemented"
	return
}

func (n *node) adjustRanges() { _ = "STUB: not implemented"; return }

func (n *node) adjustRange() { _ = "STUB: not implemented"; return }

func newDummy() node { _ = "STUB: not implemented"; return *new(node) }

func newNode(interval Interval, min, max int64, dimension uint64) *node {
	_ = "STUB: not implemented"
	return nil
}

type tree struct {
	root                 *node
	maxDimension, number uint64
	dummy                node
}

func (t *tree) Traverse(fn func(id Interval)) { _ = "STUB: not implemented"; return }

func (tree *tree) resetDummy() { _ = "STUB: not implemented"; return }

func (tree *tree) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (tree *tree) add(iv Interval) { _ = "STUB: not implemented"; return }

func (tree *tree) Add(intervals ...Interval) { _ = "STUB: not implemented"; return }

func (tree *tree) delete(iv Interval) { _ = "STUB: not implemented"; return }

func (tree *tree) Delete(intervals ...Interval) { _ = "STUB: not implemented"; return }

func (tree *tree) Query(interval Interval) Intervals {
	_ = "STUB: not implemented"
	return *new(Intervals)
}

func isRed(node *node) bool { _ = "STUB: not implemented"; return false }

func setMax(parent *node) { _ = "STUB: not implemented"; return }

func setMin(parent *node) { _ = "STUB: not implemented"; return }

func rotate(parent *node, dir int) *node { _ = "STUB: not implemented"; return nil }

func doubleRotate(parent *node, dir int) *node { _ = "STUB: not implemented"; return nil }

func intFromBool(value bool) int { _ = "STUB: not implemented"; return 0 }

func takeOpposite(value int) int { _ = "STUB: not implemented"; return 0 }

func newTree(maxDimension uint64) *tree { _ = "STUB: not implemented"; return nil }

func New(dimensions uint64) Tree { _ = "STUB: not implemented"; return *new(Tree) }
