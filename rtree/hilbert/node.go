package hilbert

import (
	"github.com/Workiva/go-datastructures/rtree"
)

type hilbert int64

type hilberts []hilbert

func getParent(parent *node, key hilbert, r1 rtree.Rectangle) *node {
	_ = "STUB: not implemented"
	return nil
}

type nodes struct {
	list rtree.Rectangles
}

func (ns *nodes) push(n rtree.Rectangle) { _ = "STUB: not implemented"; return }

func (ns *nodes) splitAt(i, capacity uint64) (*nodes, *nodes) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *nodes) byPosition(pos uint64) *node { _ = "STUB: not implemented"; return nil }

func (ns *nodes) insertAt(i uint64, n rtree.Rectangle) { _ = "STUB: not implemented"; return }

func (ns *nodes) replaceAt(i uint64, n rtree.Rectangle) { _ = "STUB: not implemented"; return }

func (ns *nodes) len() uint64 { _ = "STUB: not implemented"; return 0 }

func (ns *nodes) deleteAt(i uint64) { _ = "STUB: not implemented"; return }

func newNodes(size uint64) *nodes { _ = "STUB: not implemented"; return nil }

type keys struct {
	list hilberts
}

func (ks *keys) splitAt(i, capacity uint64) (*keys, *keys) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ks *keys) len() uint64 { _ = "STUB: not implemented"; return 0 }

func (ks *keys) byPosition(i uint64) hilbert { _ = "STUB: not implemented"; return *new(hilbert) }

func (ks *keys) deleteAt(i uint64) { _ = "STUB: not implemented"; return }

func (ks *keys) delete(k hilbert) hilbert { _ = "STUB: not implemented"; return *new(hilbert) }

func (ks *keys) search(key hilbert) uint64 { _ = "STUB: not implemented"; return 0 }

func (ks *keys) insert(key hilbert) (hilbert, uint64) {
	_ = "STUB: not implemented"
	return *new(hilbert), 0
}

func (ks *keys) last() hilbert { _ = "STUB: not implemented"; return *new(hilbert) }

func (ks *keys) insertAt(i uint64, k hilbert) { _ = "STUB: not implemented"; return }

func (ks *keys) withPosition(k hilbert) (hilbert, uint64) {
	_ = "STUB: not implemented"
	return *new(hilbert), 0
}

func newKeys(size uint64) *keys { _ = "STUB: not implemented"; return nil }

type node struct {
	keys          *keys
	nodes         *nodes
	isLeaf        bool
	parent, right *node
	mbr           *rectangle
	maxHilbert    hilbert
}

func (n *node) insert(kb *keyBundle) rtree.Rectangle {
	_ = "STUB: not implemented"
	return *new(rtree.Rectangle)
}

func (n *node) delete(kb *keyBundle) rtree.Rectangle {
	_ = "STUB: not implemented"
	return *new(rtree.Rectangle)
}

func (n *node) LowerLeft() (int32, int32) { _ = "STUB: not implemented"; return 0, 0 }

func (n *node) UpperRight() (int32, int32) { _ = "STUB: not implemented"; return 0, 0 }

func (n *node) needsSplit(ary uint64) bool { _ = "STUB: not implemented"; return false }

func (n *node) splitLeaf(i, capacity uint64) (hilbert, *node, *node) {
	_ = "STUB: not implemented"
	return *new(hilbert), nil, nil
}

func (n *node) splitInternal(i, capacity uint64) (hilbert, *node, *node) {
	_ = "STUB: not implemented"
	return *new(hilbert), nil, nil
}

func (n *node) split(i, capacity uint64) (hilbert, *node, *node) {
	_ = "STUB: not implemented"
	return *new(hilbert), nil, nil
}

func (n *node) search(key hilbert) uint64 { _ = "STUB: not implemented"; return 0 }

func (n *node) searchNode(key hilbert) *node { _ = "STUB: not implemented"; return nil }

func (n *node) searchRects(r *rectangle) rtree.Rectangles {
	_ = "STUB: not implemented"
	return *new(rtree.Rectangles)
}

func (n *node) key() hilbert { _ = "STUB: not implemented"; return *new(hilbert) }

func newNode(isLeaf bool, keys *keys, ns *nodes) *node { _ = "STUB: not implemented"; return nil }
