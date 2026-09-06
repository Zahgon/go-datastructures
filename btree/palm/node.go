package palm

import (
	"log"

	"github.com/Workiva/go-datastructures/common"
)

func getParent(parent *node, key common.Comparator) *node { _ = "STUB: not implemented"; return nil }

type nodes struct {
	list []*node
}

func (ns *nodes) push(n *node) { _ = "STUB: not implemented"; return }

func (ns *nodes) splitAt(i, capacity uint64) (*nodes, *nodes) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *nodes) byPosition(pos uint64) *node { _ = "STUB: not implemented"; return nil }

func (ns *nodes) insertAt(i uint64, n *node) { _ = "STUB: not implemented"; return }

func (ns *nodes) replaceAt(i uint64, n *node) { _ = "STUB: not implemented"; return }

func (ns *nodes) len() uint64 { _ = "STUB: not implemented"; return 0 }

func newNodes(size uint64) *nodes { _ = "STUB: not implemented"; return nil }

type keys struct {
	list common.Comparators
}

func (ks *keys) splitAt(i, capacity uint64) (*keys, *keys) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ks *keys) len() uint64 { _ = "STUB: not implemented"; return 0 }

func (ks *keys) byPosition(i uint64) common.Comparator {
	_ = "STUB: not implemented"
	return *new(common.Comparator)
}

func (ks *keys) delete(k common.Comparator) common.Comparator {
	_ = "STUB: not implemented"
	return *new(common.Comparator)
}

func (ks *keys) search(key common.Comparator) uint64 { _ = "STUB: not implemented"; return 0 }

func (ks *keys) insert(key common.Comparator) (common.Comparator, uint64) {
	_ = "STUB: not implemented"
	return *new(common.Comparator), 0
}

func (ks *keys) last() common.Comparator { _ = "STUB: not implemented"; return *new(common.Comparator) }

func (ks *keys) insertAt(i uint64, k common.Comparator) { _ = "STUB: not implemented"; return }

func (ks *keys) withPosition(k common.Comparator) (common.Comparator, uint64) {
	_ = "STUB: not implemented"
	return *new(common.Comparator), 0
}

func newKeys(size uint64) *keys { _ = "STUB: not implemented"; return nil }

type node struct {
	keys          *keys
	nodes         *nodes
	isLeaf        bool
	parent, right *node
}

func (n *node) needsSplit(ary uint64) bool { _ = "STUB: not implemented"; return false }

func (n *node) splitLeaf(i, capacity uint64) (common.Comparator, *node, *node) {
	_ = "STUB: not implemented"
	return *new(common.Comparator), nil, nil
}

func (n *node) splitInternal(i, capacity uint64) (common.Comparator, *node, *node) {
	_ = "STUB: not implemented"
	return *new(common.Comparator), nil, nil
}

func (n *node) split(i, capacity uint64) (common.Comparator, *node, *node) {
	_ = "STUB: not implemented"
	return *new(common.Comparator), nil, nil
}

func (n *node) search(key common.Comparator) uint64 { _ = "STUB: not implemented"; return 0 }

func (n *node) searchNode(key common.Comparator) *node { _ = "STUB: not implemented"; return nil }

func (n *node) key() common.Comparator { _ = "STUB: not implemented"; return *new(common.Comparator) }

func (n *node) print(output *log.Logger) { _ = "STUB: not implemented"; return }

func (n *node) Compare(e common.Comparator) int { _ = "STUB: not implemented"; return 0 }

func newNode(isLeaf bool, keys *keys, ns *nodes) *node { _ = "STUB: not implemented"; return nil }
