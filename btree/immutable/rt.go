//go:generate msgp -tests=false -io=false

package btree

import "sync"

type context struct {
	lock      sync.RWMutex
	seenNodes map[string]*Node
}

func (c *context) nodeExists(id ID) bool { _ = "STUB: not implemented"; return false }

func (c *context) addNode(n *Node) { _ = "STUB: not implemented"; return }

func (c *context) getNode(id ID) *Node { _ = "STUB: not implemented"; return nil }

func newContext() *context { _ = "STUB: not implemented"; return nil }

type Tr struct {
	UUID      ID  `msg:"u"`
	Count     int `msg:"c"`
	config    Config
	Root      ID `msg:"r"`
	cacher    *cacher
	context   *context
	NodeWidth int `msg:"nw"`
	mutable   bool
}

func (t *Tr) createRoot() *Node { _ = "STUB: not implemented"; return nil }

func (t *Tr) contextOrCachedNode(id ID, cache bool) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tr) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (t *Tr) toBytes() []byte { _ = "STUB: not implemented"; return nil }

func (t *Tr) reset() { _ = "STUB: not implemented"; return }

func (t *Tr) commit() []*Payload { _ = "STUB: not implemented"; return nil }

func (t *Tr) copyNode(n *Node) *Node { _ = "STUB: not implemented"; return nil }

func (t *Tr) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *Tr) AsMutable() MutableTree { _ = "STUB: not implemented"; return *new(MutableTree) }

func (t *Tr) Commit() (ReadableTree, error) {
	_ = "STUB: not implemented"
	return *new(ReadableTree), nil
}

func treeFromBytes(p Persister, data []byte, comparator Comparator) (*Tr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTree(cfg Config) *Tr { _ = "STUB: not implemented"; return nil }

func New(cfg Config) ReadableTree { _ = "STUB: not implemented"; return *new(ReadableTree) }

func Load(p Persister, id []byte, comparator Comparator) (ReadableTree, error) {
	_ = "STUB: not implemented"
	return *new(ReadableTree), nil
}
