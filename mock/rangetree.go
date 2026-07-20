package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/Workiva/go-datastructures/rangetree"
)

type RangeTree struct {
	mock.Mock
}

var _ rangetree.RangeTree = new(RangeTree)

func (m *RangeTree) Add(entries ...rangetree.Entry) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}

func (m *RangeTree) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (m *RangeTree) Delete(entries ...rangetree.Entry) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}

func (m *RangeTree) Query(interval rangetree.Interval) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}

func (m *RangeTree) InsertAtDimension(dimension uint64, index,
	number int64) (rangetree.Entries, rangetree.Entries) {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries), *new(rangetree.Entries)
}

func (m *RangeTree) Apply(interval rangetree.Interval, fn func(rangetree.Entry) bool) {
	_ = "STUB: not implemented"
	return
}

func (m *RangeTree) Get(entries ...rangetree.Entry) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}
