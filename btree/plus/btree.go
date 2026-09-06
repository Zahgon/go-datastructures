package plus

func keySearch(keys keys, key Key) int { _ = "STUB: not implemented"; return 0 }

type btree struct {
	root             node
	nodeSize, number uint64
}

func (tree *btree) insert(key Key) { _ = "STUB: not implemented"; return }

func (tree *btree) Insert(keys ...Key) { _ = "STUB: not implemented"; return }

func (tree *btree) Iter(key Key) Iterator { _ = "STUB: not implemented"; return *new(Iterator) }

func (tree *btree) get(key Key) Key { _ = "STUB: not implemented"; return *new(Key) }

func (tree *btree) Get(keys ...Key) Keys { _ = "STUB: not implemented"; return *new(Keys) }

func (tree *btree) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func newBTree(nodeSize uint64) *btree { _ = "STUB: not implemented"; return nil }
