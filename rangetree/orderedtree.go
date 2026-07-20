package rangetree

func isLastDimension(value, test uint64) bool { _ = "STUB: not implemented"; return false }

type nodeBundle struct {
	list  *orderedNodes
	index int
}

type orderedTree struct {
	top        orderedNodes
	number     uint64
	dimensions uint64
	path       []*nodeBundle
}

func (ot *orderedTree) resetPath() { _ = "STUB: not implemented"; return }

func (ot *orderedTree) needNextDimension() bool { _ = "STUB: not implemented"; return false }

func (ot *orderedTree) add(entry Entry) *node { _ = "STUB: not implemented"; return nil }

func (ot *orderedTree) Add(entries ...Entry) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (ot *orderedTree) delete(entry Entry) *node { _ = "STUB: not implemented"; return nil }

func (ot *orderedTree) get(entry Entry) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (ot *orderedTree) Get(entries ...Entry) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (ot *orderedTree) Delete(entries ...Entry) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (ot *orderedTree) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (ot *orderedTree) apply(list orderedNodes, interval Interval,
	dimension uint64, fn func(*node) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (ot *orderedTree) Apply(interval Interval, fn func(Entry) bool) {
	_ = "STUB: not implemented"
	return
}

func (ot *orderedTree) Query(interval Interval) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (ot *orderedTree) InsertAtDimension(dimension uint64,
	index, number int64) (Entries, Entries) {
	_ = "STUB: not implemented"
	return *new(Entries), *new(Entries)
}

func newOrderedTree(dimensions uint64) *orderedTree { _ = "STUB: not implemented"; return nil }

func New(dimensions uint64) RangeTree { _ = "STUB: not implemented"; return *new(RangeTree) }
