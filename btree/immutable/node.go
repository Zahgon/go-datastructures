//go:generate msgp -tests=false -io=false

package btree

func newID() []byte { _ = "STUB: not implemented"; return nil }

type ID []byte

type Key struct {
	UUID    ID          `msg:"u"`
	Value   interface{} `msg:"v"`
	Payload []byte      `msg:"p"`
}

func (k Key) ID() []byte { _ = "STUB: not implemented"; return nil }

func (k Key) ToItem() *Item { _ = "STUB: not implemented"; return nil }

type Keys []*Key

func (k Keys) toItems() items { _ = "STUB: not implemented"; return *new(items) }

func (k Keys) sort(comparator Comparator) Keys { _ = "STUB: not implemented"; return *new(Keys) }

type keySortWrapper struct {
	comparator Comparator
	keys       Keys
}

func (sw *keySortWrapper) Len() int { _ = "STUB: not implemented"; return 0 }

func (sw *keySortWrapper) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (sw *keySortWrapper) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (sw *keySortWrapper) sort() Keys { _ = "STUB: not implemented"; return *new(Keys) }

func splitKeys(keys Keys, numParts int) []Keys { _ = "STUB: not implemented"; return nil }

type Node struct {
	ID ID `msg:"id"`

	IsLeaf bool `msg:"il"`

	ChildValues []interface{} `msg:"cv"`

	ChildKeys Keys `msg:"ck"`
}

func (n *Node) copy() *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) searchKey(comparator Comparator, value interface{}) (*Key, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (n *Node) insert(comparator Comparator, key *Key) *Key { _ = "STUB: not implemented"; return nil }

func (n *Node) delete(comparator Comparator, key *Key) *Key { _ = "STUB: not implemented"; return nil }

func (n *Node) multiDelete(comparator Comparator, keys ...*Key) { _ = "STUB: not implemented"; return }

func (n *Node) replaceKeyAt(key *Key, i int) { _ = "STUB: not implemented"; return }

func (n *Node) flatten() ([]interface{}, Keys) { _ = "STUB: not implemented"; return nil, *new(Keys) }

func (n *Node) iter(comparator Comparator, start, stop interface{}) iterator {
	_ = "STUB: not implemented"
	return *new(iterator)
}

func (n *Node) valueAt(i int) interface{} { _ = "STUB: not implemented"; return nil }

func (n *Node) keyAt(i int) *Key { _ = "STUB: not implemented"; return nil }

func (n *Node) needsSplit(max int) bool { _ = "STUB: not implemented"; return false }

func (n *Node) lastValue() interface{} { _ = "STUB: not implemented"; return nil }

func (n *Node) firstValue() interface{} { _ = "STUB: not implemented"; return nil }

func (n *Node) append(other *Node) { _ = "STUB: not implemented"; return }

func (n *Node) replaceValueAt(i int, value interface{}) { _ = "STUB: not implemented"; return }

func (n *Node) deleteValueAt(i int) { _ = "STUB: not implemented"; return }

func (n *Node) deleteKeyAt(i int) { _ = "STUB: not implemented"; return }

func (n *Node) splitLeafAt(i int) (interface{}, *Node) { _ = "STUB: not implemented"; return nil, nil }

func (n *Node) splitInternalAt(i int) (interface{}, *Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *Node) splitAt(i int) (interface{}, *Node) { _ = "STUB: not implemented"; return nil, nil }

func (n *Node) lenKeys() int { _ = "STUB: not implemented"; return 0 }

func (n *Node) lenValues() int { _ = "STUB: not implemented"; return 0 }

func (n *Node) appendChild(key *Key) { _ = "STUB: not implemented"; return }

func (n *Node) appendValue(value interface{}) { _ = "STUB: not implemented"; return }

func (n *Node) popFirstKey() *Key { _ = "STUB: not implemented"; return nil }

func (n *Node) popFirstValue() interface{} { _ = "STUB: not implemented"; return nil }

func (n *Node) popKey() *Key { _ = "STUB: not implemented"; return nil }

func (n *Node) popValue() interface{} { _ = "STUB: not implemented"; return nil }

func (n *Node) prependKey(key *Key) { _ = "STUB: not implemented"; return }

func (n *Node) prependValue(value interface{}) { _ = "STUB: not implemented"; return }

func (n *Node) search(comparator Comparator, value interface{}) int {
	_ = "STUB: not implemented"
	return 0
}

func nodeFromBytes(t *Tr, data []byte) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

func newNode() *Node { _ = "STUB: not implemented"; return nil }

type sliceIterator struct {
	stop       interface{}
	n          *Node
	pointer    int
	comparator Comparator
}

func (s *sliceIterator) next() bool { _ = "STUB: not implemented"; return false }

func (s *sliceIterator) value() (*Key, int) { _ = "STUB: not implemented"; return nil, 0 }

type iterator interface {
	next() bool
	value() (*Key, int)
}

type nodeBundle struct {
	path *path
	k    *Key
}

type nodeSortWrapper struct {
	values     []interface{}
	keys       Keys
	comparator Comparator
}

func (n *nodeSortWrapper) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *nodeSortWrapper) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (n *nodeSortWrapper) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func splitValues(values []interface{}, numParts int) [][]interface{} {
	_ = "STUB: not implemented"
	return nil
}
