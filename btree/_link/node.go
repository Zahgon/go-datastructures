package link

import (
	"log"
	"sync"
)

func search(parent *node, key Key) Key { _ = "STUB: not implemented"; return *new(Key) }

func getParent(parent *node, stack *nodes, key Key) *node { _ = "STUB: not implemented"; return nil }

func insert(tree *blink, parent *node, stack *nodes, key Key) Key {
	_ = "STUB: not implemented"
	return *new(Key)
}

func split(tree *blink, n *node, stack *nodes) { _ = "STUB: not implemented"; return }

func moveRight(n *node, key Key, getLock bool) *node { _ = "STUB: not implemented"; return nil }

type nodes []*node

func (ns *nodes) reset() { _ = "STUB: not implemented"; return }

func (ns *nodes) push(n *node) { _ = "STUB: not implemented"; return }

func (ns *nodes) pop() *node { _ = "STUB: not implemented"; return nil }

func (ns *nodes) insertAt(n *node, i int) { _ = "STUB: not implemented"; return }

func (ns *nodes) splitAt(i int) (nodes, nodes) {
	_ = "STUB: not implemented"
	return *new(nodes), *new(nodes)
}

type node struct {
	keys    Keys
	nodes   nodes
	right   *node
	lock    sync.RWMutex
	isLeaf  bool
	maxSeen Key
}

func (n *node) key() Key { _ = "STUB: not implemented"; return *new(Key) }

func (n *node) insert(key Key) Key { _ = "STUB: not implemented"; return *new(Key) }

func (n *node) insertNode(other *node) { _ = "STUB: not implemented"; return }

func (n *node) needsSplit() bool { _ = "STUB: not implemented"; return false }

func (n *node) max() Key { _ = "STUB: not implemented"; return *new(Key) }

func (n *node) splitLeaf() (Key, *node, *node) {
	_ = "STUB: not implemented"
	return *new(Key), nil, nil
}

func (n *node) splitInternal() (Key, *node, *node) {
	_ = "STUB: not implemented"
	return *new(Key), nil, nil
}

func (n *node) split() (Key, *node, *node) { _ = "STUB: not implemented"; return *new(Key), nil, nil }

func (n *node) search(key Key) int { _ = "STUB: not implemented"; return 0 }

func (n *node) searchNode(key Key) *node { _ = "STUB: not implemented"; return nil }

func (n *node) print(output *log.Logger) { _ = "STUB: not implemented"; return }

func newNode(isLeaf bool, keys Keys, ns nodes) *node { _ = "STUB: not implemented"; return nil }
