package skiplist

import (
	"github.com/Workiva/go-datastructures/common"
	"github.com/Workiva/go-datastructures/rangetree"
	"github.com/Workiva/go-datastructures/slice/skip"
)

type keyed interface {
	key() uint64
}

type skipEntry uint64

func (se skipEntry) Compare(other common.Comparator) int { _ = "STUB: not implemented"; return 0 }

func (se skipEntry) key() uint64 { _ = "STUB: not implemented"; return 0 }

func isLastDimension(dimension, lastDimension uint64) bool { _ = "STUB: not implemented"; return false }

func needsDeletion(value, index, number int64) bool { _ = "STUB: not implemented"; return false }

type dimensionalBundle struct {
	id uint64
	sl *skip.SkipList
}

func (db *dimensionalBundle) Compare(e common.Comparator) int { _ = "STUB: not implemented"; return 0 }

func (db *dimensionalBundle) key() uint64 { _ = "STUB: not implemented"; return 0 }

type lastBundle struct {
	id    uint64
	entry rangetree.Entry
}

func (lb *lastBundle) Compare(e common.Comparator) int { _ = "STUB: not implemented"; return 0 }

func (lb *lastBundle) key() uint64 { _ = "STUB: not implemented"; return 0 }

type skipListRT struct {
	top                *skip.SkipList
	dimensions, number uint64
}

func (rt *skipListRT) init(dimensions uint64) {
	rt.dimensions = dimensions
	rt.top = skip.New(uint64(0))
}

func (rt *skipListRT) add(entry rangetree.Entry) rangetree.Entry {
	_ = "STUB: not implemented"
	return *new(rangetree.Entry)
}

func (rt *skipListRT) Add(entries ...rangetree.Entry) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}

func (rt *skipListRT) get(entry rangetree.Entry) rangetree.Entry {
	_ = "STUB: not implemented"
	return *new(rangetree.Entry)
}

func (rt *skipListRT) Get(entries ...rangetree.Entry) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}

func (rt *skipListRT) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (rt *skipListRT) deleteRecursive(sl *skip.SkipList, dimension uint64,
	entry rangetree.Entry) rangetree.Entry {
	_ = "STUB: not implemented"
	return *new(rangetree.Entry)
}

func (rt *skipListRT) delete(entry rangetree.Entry) rangetree.Entry {
	_ = "STUB: not implemented"
	return *new(rangetree.Entry)
}

func (rt *skipListRT) Delete(entries ...rangetree.Entry) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}

func (rt *skipListRT) apply(sl *skip.SkipList, dimension uint64,
	interval rangetree.Interval, fn func(rangetree.Entry) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (rt *skipListRT) Apply(interval rangetree.Interval, fn func(rangetree.Entry) bool) {
	_ = "STUB: not implemented"
	return
}

func (rt *skipListRT) Query(interval rangetree.Interval) rangetree.Entries {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries)
}

func (rt *skipListRT) flatten(sl *skip.SkipList, dimension uint64, entries *rangetree.Entries) {
	_ = "STUB: not implemented"
	return
}

func (rt *skipListRT) insert(sl *skip.SkipList, dimension, insertDimension uint64,
	index, number int64, deleted, affected *rangetree.Entries) {
	_ = "STUB: not implemented"
	return
}

func (rt *skipListRT) InsertAtDimension(dimension uint64,
	index, number int64) (rangetree.Entries, rangetree.Entries) {
	_ = "STUB: not implemented"
	return *new(rangetree.Entries), *new(rangetree.Entries)
}

func new(dimensions uint64) *skipListRT { _ = "STUB: not implemented"; return nil }

func New(dimensions uint64) rangetree.RangeTree {
	_ = "STUB: not implemented"
	return *new(rangetree.RangeTree)
}
