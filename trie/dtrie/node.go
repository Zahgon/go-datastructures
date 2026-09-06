package dtrie

import (
	"github.com/Workiva/go-datastructures/bitarray"
)

type node struct {
	entries []Entry
	nodeMap bitarray.Bitmap32
	dataMap bitarray.Bitmap32
	level   uint8
}

func (n *node) KeyHash() uint32    { _ = "STUB: not implemented"; return 0 }
func (n *node) Key() interface{}   { _ = "STUB: not implemented"; return nil }
func (n *node) Value() interface{} { _ = "STUB: not implemented"; return nil }

func (n *node) String() string { _ = "STUB: not implemented"; return "" }

type collisionNode struct {
	entries []Entry
}

func (n *collisionNode) KeyHash() uint32    { _ = "STUB: not implemented"; return 0 }
func (n *collisionNode) Key() interface{}   { _ = "STUB: not implemented"; return nil }
func (n *collisionNode) Value() interface{} { _ = "STUB: not implemented"; return nil }

func (n *collisionNode) String() string { _ = "STUB: not implemented"; return "" }

type Entry interface {
	KeyHash() uint32
	Key() interface{}
	Value() interface{}
}

func emptyNode(level uint8, capacity int) *node { _ = "STUB: not implemented"; return nil }

func insert(n *node, entry Entry) *node { _ = "STUB: not implemented"; return nil }

func get(n *node, keyHash uint32, key interface{}) Entry {
	_ = "STUB: not implemented"
	return *new(Entry)
}

func remove(n *node, keyHash uint32, key interface{}) *node { _ = "STUB: not implemented"; return nil }

func iterate(n *node, stop <-chan struct{}) <-chan Entry { _ = "STUB: not implemented"; return nil }

func pushEntries(n *node, stop <-chan struct{}, out chan Entry) { _ = "STUB: not implemented"; return }
