package avl

type Immutable struct {
	root   *node
	number uint64
	dummy  node
}

func (immutable *Immutable) copy() *Immutable { _ = "STUB: not implemented"; return nil }

func (immutable *Immutable) resetDummy() { _ = "STUB: not implemented"; return }

func (immutable *Immutable) init() {
	immutable.dummy = node{
		children: [2]*node{},
	}
}

func (immutable *Immutable) get(entry Entry) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (immutable *Immutable) Get(entries ...Entry) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (immutable *Immutable) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (immutable *Immutable) insert(entry Entry) Entry {
	_ = "STUB: not implemented"
	return *new(Entry)
}

func (immutable *Immutable) Insert(entries ...Entry) (*Immutable, Entries) {
	_ = "STUB: not implemented"
	return nil, *new(Entries)
}

func (immutable *Immutable) delete(entry Entry) Entry {
	_ = "STUB: not implemented"
	return *new(Entry)
}

func (immutable *Immutable) Delete(entries ...Entry) (*Immutable, Entries) {
	_ = "STUB: not implemented"
	return nil, *new(Entries)
}

func insertBalance(root *node, dir int) *node { _ = "STUB: not implemented"; return nil }

func removeBalance(root *node, dir int, done *int) *node { _ = "STUB: not implemented"; return nil }

func intFromBool(value bool) int { _ = "STUB: not implemented"; return 0 }

func takeOpposite(value int) int { _ = "STUB: not implemented"; return 0 }

func adjustBalance(root *node, dir, bal int) { _ = "STUB: not implemented"; return }

func rotate(parent *node, dir int) *node { _ = "STUB: not implemented"; return nil }

func doubleRotate(parent *node, dir int) *node { _ = "STUB: not implemented"; return nil }

func normalizeComparison(i int) int { _ = "STUB: not implemented"; return 0 }

func NewImmutable() *Immutable { _ = "STUB: not implemented"; return nil }
