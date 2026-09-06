package rangetree

import "github.com/Workiva/go-datastructures/slice"

type immutableRangeTree struct {
	number     uint64
	top        orderedNodes
	dimensions uint64
}

func newCache(dimensions uint64) []slice.Int64Slice { _ = "STUB: not implemented"; return nil }

func (irt *immutableRangeTree) needNextDimension() bool { _ = "STUB: not implemented"; return false }

func (irt *immutableRangeTree) add(nodes *orderedNodes, cache []slice.Int64Slice, entry Entry, added *uint64) {
	_ = "STUB: not implemented"
	return
}

func (irt *immutableRangeTree) Add(entries ...Entry) *immutableRangeTree {
	_ = "STUB: not implemented"
	return nil
}

func (irt *immutableRangeTree) InsertAtDimension(dimension uint64,
	index, number int64) (*immutableRangeTree, Entries, Entries) {
	_ = "STUB: not implemented"
	return nil, *new(Entries), *new(Entries)
}

type immutableNodeBundle struct {
	list         *orderedNodes
	index        int
	previousNode *node
	newNode      *node
}

func (irt *immutableRangeTree) Delete(entries ...Entry) *immutableRangeTree {
	_ = "STUB: not implemented"
	return nil
}

func (irt *immutableRangeTree) delete(top *orderedNodes,
	cache []slice.Int64Slice, entry Entry, deleted *uint64) {
	_ = "STUB: not implemented"
	return
}

func (irt *immutableRangeTree) apply(list orderedNodes, interval Interval,
	dimension uint64, fn func(*node) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (irt *immutableRangeTree) Query(interval Interval) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (irt *immutableRangeTree) get(entry Entry) Entry {
	_ = "STUB: not implemented"
	return *new(Entry)
}

func (irt *immutableRangeTree) Get(entries ...Entry) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (irt *immutableRangeTree) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func newImmutableRangeTree(dimensions uint64) *immutableRangeTree {
	_ = "STUB: not implemented"
	return nil
}
