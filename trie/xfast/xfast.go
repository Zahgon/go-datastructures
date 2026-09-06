package xfast

func isInternal(n *node) bool { _ = "STUB: not implemented"; return false }

func hasInternal(n *node) bool { _ = "STUB: not implemented"; return false }

func isLeaf(n *node) bool { _ = "STUB: not implemented"; return false }

type Entry interface {
	Key() uint64
}

var masks = func() [64]uint64 {
	masks := [64]uint64{}
	mask := uint64(0)
	for i := uint64(0); i < 64; i++ {
		mask = mask | 1<<(63-i)
		masks[i] = mask
	}
	return masks
}()

var positions = func() [64]uint64 {
	positions := [64]uint64{}
	for i := uint64(0); i < 64; i++ {
		positions[i] = uint64(1 << (63 - i))
	}
	return positions
}()

type node struct {
	entry Entry

	children [2]*node

	parent *node
}

func newNode(parent *node, entry Entry) *node { _ = "STUB: not implemented"; return nil }

func binarySearchHashMaps(layers []map[uint64]*node, key uint64) (int, *node) {
	_ = "STUB: not implemented"
	return 0, nil
}

func whichSide(n, parent *node) int { _ = "STUB: not implemented"; return 0 }

type XFastTrie struct {
	layers []map[uint64]*node

	root *node

	num uint64

	bits, diff uint8

	min, max *node
}

func (xft *XFastTrie) init(intType interface{}) {
	bits := uint8(0)
	switch intType.(type) {
	case uint8:
		bits = 8
	case uint16:
		bits = 16
	case uint32:
		bits = 32
	case uint, uint64:
		bits = 64
	default:

		panic(`Invalid universe size provided.`)
	}

	xft.layers = make([]map[uint64]*node, bits)
	xft.bits = bits
	xft.diff = 64 - bits
	for i := uint8(0); i < bits; i++ {
		xft.layers[i] = make(map[uint64]*node, 50)
	}
	xft.num = 0
	xft.root = newNode(nil, nil)
}

func (xft *XFastTrie) Exists(key uint64) bool { _ = "STUB: not implemented"; return false }

func (xft *XFastTrie) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (xft *XFastTrie) Max() Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (xft *XFastTrie) Min() Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (xft *XFastTrie) insert(entry Entry) { _ = "STUB: not implemented"; return }

func (xft *XFastTrie) walkUpSuccessor(root, node, successor *node) {
	_ = "STUB: not implemented"
	return
}

func (xft *XFastTrie) walkUpPredecessor(root, node, predecessor *node) {
	_ = "STUB: not implemented"
	return
}

func (xft *XFastTrie) walkUpNode(root, node, predecessor, successor *node) {
	_ = "STUB: not implemented"
	return
}

func (xft *XFastTrie) Insert(entries ...Entry) { _ = "STUB: not implemented"; return }

func (xft *XFastTrie) delete(key uint64) { _ = "STUB: not implemented"; return }

func (xft *XFastTrie) Delete(keys ...uint64) { _ = "STUB: not implemented"; return }

func (xft *XFastTrie) predecessor(key uint64) *node { _ = "STUB: not implemented"; return nil }

func (xft *XFastTrie) successor(key uint64) *node { _ = "STUB: not implemented"; return nil }

func (xft *XFastTrie) Successor(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (xft *XFastTrie) Predecessor(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (xft *XFastTrie) Iter(key uint64) *Iterator { _ = "STUB: not implemented"; return nil }

func (xft *XFastTrie) Get(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func New(ifc interface{}) *XFastTrie { _ = "STUB: not implemented"; return nil }
