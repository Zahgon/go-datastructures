package hilbert

import (
	"github.com/Workiva/go-datastructures/queue"
	"github.com/Workiva/go-datastructures/rtree"
)

type operation int

const (
	get operation = iota
	add
	remove
)

const multiThreadAt = 1000

type keyBundle struct {
	key         hilbert
	left, right rtree.Rectangle
}

type tree struct {
	root            *node
	_               [8]uint64
	number          uint64
	_               [8]uint64
	ary, bufferSize uint64
	actions         *queue.RingBuffer
	cache           []interface{}
	_               [8]uint64
	disposed        uint64
	_               [8]uint64
	running         uint64
}

func (tree *tree) checkAndRun(action action) { _ = "STUB: not implemented"; return }

func (tree *tree) init(bufferSize, ary uint64) {
	tree.bufferSize = bufferSize
	tree.ary = ary
	tree.cache = make([]interface{}, 0, bufferSize)
	tree.root = newNode(true, newKeys(ary), newNodes(ary))
	tree.root.mbr = &rectangle{}
	tree.actions = queue.NewRingBuffer(tree.bufferSize)
}

func (tree *tree) operationRunner(xns interfaces, threaded bool) { _ = "STUB: not implemented"; return }

func (tree *tree) fetchKeys(xns interfaces, inParallel bool) (map[*node][]*keyBundle, map[*node][]*keyBundle, actions) {
	_ = "STUB: not implemented"
	return nil, nil, *new(actions)
}

func (tree *tree) fetchKeysInSerial(xns interfaces) { _ = "STUB: not implemented"; return }

func (tree *tree) reset() { _ = "STUB: not implemented"; return }

func (tree *tree) fetchKeysInParallel(xns []interface{}) { _ = "STUB: not implemented"; return }

func (tree *tree) splitNode(n, parent *node, nodes *[]*node, keys *hilberts) {
	_ = "STUB: not implemented"
	return
}

func (tree *tree) applyNode(n *node, adds, deletes []*keyBundle) { _ = "STUB: not implemented"; return }

func (tree *tree) recursiveMutate(adds, deletes map[*node][]*keyBundle, setRoot, inParallel bool) {
	_ = "STUB: not implemented"
	return
}

func (tree *tree) Insert(rects ...rtree.Rectangle) { _ = "STUB: not implemented"; return }

func (tree *tree) Delete(rects ...rtree.Rectangle) { _ = "STUB: not implemented"; return }

func (tree *tree) search(r *rectangle) rtree.Rectangles {
	_ = "STUB: not implemented"
	return *new(rtree.Rectangles)
}

func (tree *tree) Search(rect rtree.Rectangle) rtree.Rectangles {
	_ = "STUB: not implemented"
	return *new(rtree.Rectangles)
}

func (tree *tree) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (tree *tree) Dispose() { _ = "STUB: not implemented"; return }

func newTree(bufferSize, ary uint64) *tree { _ = "STUB: not implemented"; return nil }

func New(bufferSize, ary uint64) rtree.RTree { _ = "STUB: not implemented"; return *new(rtree.RTree) }
