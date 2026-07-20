package rangetree

type orderedNodes nodes

func (nodes orderedNodes) search(value int64) int { _ = "STUB: not implemented"; return 0 }

func (nodes *orderedNodes) addAt(i int, node *node) *node { _ = "STUB: not implemented"; return nil }

func (nodes *orderedNodes) add(node *node) *node { _ = "STUB: not implemented"; return nil }

func (nodes *orderedNodes) deleteAt(i int) *node { _ = "STUB: not implemented"; return nil }

func (nodes *orderedNodes) delete(value int64) *node { _ = "STUB: not implemented"; return nil }

func (nodes orderedNodes) apply(low, high int64, fn func(*node) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (nodes orderedNodes) get(value int64) (*node, int) { _ = "STUB: not implemented"; return nil, 0 }

func (nodes *orderedNodes) getOrAdd(entry Entry,
	dimension, lastDimension uint64) (*node, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (nodes orderedNodes) flatten(entries *Entries) { _ = "STUB: not implemented"; return }

func (nodes *orderedNodes) insert(insertDimension, dimension, maxDimension uint64,
	index, number int64, modified, deleted *Entries) {
	_ = "STUB: not implemented"
	return
}

func (nodes orderedNodes) immutableInsert(insertDimension, dimension, maxDimension uint64,
	index, number int64, modified, deleted *Entries) orderedNodes {
	_ = "STUB: not implemented"
	return *new(orderedNodes)
}
