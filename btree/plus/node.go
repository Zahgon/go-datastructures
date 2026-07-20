package plus

func split(tree *btree, parent, child node) node { _ = "STUB: not implemented"; return *new(node) }

type node interface {
	insert(tree *btree, key Key) bool
	needsSplit(nodeSize uint64) bool

	split() (Key, node, node)
	search(key Key) int
	find(key Key) *iterator
}

type nodes []node

func (nodes *nodes) insertAt(i int, node node) { _ = "STUB: not implemented"; return }

func (ns nodes) splitAt(i int) (nodes, nodes) {
	_ = "STUB: not implemented"
	return *new(nodes), *new(nodes)
}

type inode struct {
	keys  keys
	nodes nodes
}

func (node *inode) search(key Key) int { _ = "STUB: not implemented"; return 0 }

func (node *inode) find(key Key) *iterator { _ = "STUB: not implemented"; return nil }

func (n *inode) insert(tree *btree, key Key) bool { _ = "STUB: not implemented"; return false }

func (n *inode) needsSplit(nodeSize uint64) bool { _ = "STUB: not implemented"; return false }

func (n *inode) split() (Key, node, node) {
	_ = "STUB: not implemented"
	return *new(Key), *new(node), *new(node)
}

func newInternalNode(size uint64) *inode { _ = "STUB: not implemented"; return nil }

type lnode struct {
	pointer *lnode
	keys    keys
}

func (node *lnode) search(key Key) int { _ = "STUB: not implemented"; return 0 }

func (lnode *lnode) insert(tree *btree, key Key) bool { _ = "STUB: not implemented"; return false }

func (node *lnode) find(key Key) *iterator { _ = "STUB: not implemented"; return nil }

func (node *lnode) split() (Key, node, node) {
	_ = "STUB: not implemented"
	return *new(Key), *new(node), *new(node)
}

func (lnode *lnode) needsSplit(nodeSize uint64) bool { _ = "STUB: not implemented"; return false }

func newLeafNode(size uint64) *lnode { _ = "STUB: not implemented"; return nil }

type keys []Key

func (keys keys) search(key Key) int { _ = "STUB: not implemented"; return 0 }

func (keys *keys) insertAt(i int, key Key) { _ = "STUB: not implemented"; return }

func (keys keys) reverse() { _ = "STUB: not implemented"; return }
