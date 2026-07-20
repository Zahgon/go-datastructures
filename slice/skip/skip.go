package skip

import (
	"math/rand"
	"sync"
	"time"

	"github.com/Workiva/go-datastructures/common"
)

const p = .5

type lockedSource struct {
	mu  sync.Mutex
	src rand.Source
}

func (ls *lockedSource) Int63() (n int64) { _ = "STUB: not implemented"; return 0 }

func (ls *lockedSource) Seed(seed int64) { _ = "STUB: not implemented"; return }

var generator = rand.New(&lockedSource{src: rand.NewSource(time.Now().UnixNano())})

func generateLevel(maxLevel uint8) uint8 { _ = "STUB: not implemented"; return 0 }

func insertNode(sl *SkipList, n *node, cmp common.Comparator, pos uint64, cache nodes, posCache widths, allowDuplicate bool) common.Comparator {
	_ = "STUB: not implemented"
	return *new(common.Comparator)
}

func splitAt(sl *SkipList, index uint64) (*SkipList, *SkipList) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SkipList struct {
	maxLevel, level uint8
	head            *node
	num             uint64

	cache    nodes
	posCache widths
}

func (sl *SkipList) init(ifc interface{}) {
	switch ifc.(type) {
	case uint8:
		sl.maxLevel = 8
	case uint16:
		sl.maxLevel = 16
	case uint32:
		sl.maxLevel = 32
	case uint64, uint:
		sl.maxLevel = 64
	}
	sl.cache = make(nodes, sl.maxLevel)
	sl.posCache = make(widths, sl.maxLevel)
	sl.head = newNode(nil, sl.maxLevel)
}

func (sl *SkipList) search(cmp common.Comparator, update nodes, widths widths) (*node, uint64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (sl *SkipList) resetMaxLevel() { _ = "STUB: not implemented"; return }

func (sl *SkipList) searchByPosition(position uint64, update nodes, widths widths) (*node, uint64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (sl *SkipList) Get(comparators ...common.Comparator) common.Comparators {
	_ = "STUB: not implemented"
	return *new(common.Comparators)
}

func (sl *SkipList) GetWithPosition(cmp common.Comparator) (common.Comparator, uint64) {
	_ = "STUB: not implemented"
	return *new(common.Comparator), 0
}

func (sl *SkipList) ByPosition(position uint64) common.Comparator {
	_ = "STUB: not implemented"
	return *new(common.Comparator)
}

func (sl *SkipList) insert(cmp common.Comparator) common.Comparator {
	_ = "STUB: not implemented"
	return *new(common.Comparator)
}

func (sl *SkipList) Insert(comparators ...common.Comparator) common.Comparators {
	_ = "STUB: not implemented"
	return *new(common.Comparators)
}

func (sl *SkipList) insertAtPosition(position uint64, cmp common.Comparator) {
	_ = "STUB: not implemented"
	return
}

func (sl *SkipList) InsertAtPosition(position uint64, cmp common.Comparator) {
	_ = "STUB: not implemented"
	return
}

func (sl *SkipList) replaceAtPosition(position uint64, cmp common.Comparator) {
	_ = "STUB: not implemented"
	return
}

func (sl *SkipList) ReplaceAtPosition(position uint64, cmp common.Comparator) {
	_ = "STUB: not implemented"
	return
}

func (sl *SkipList) delete(cmp common.Comparator) common.Comparator {
	_ = "STUB: not implemented"
	return *new(common.Comparator)
}

func (sl *SkipList) Delete(comparators ...common.Comparator) common.Comparators {
	_ = "STUB: not implemented"
	return *new(common.Comparators)
}

func (sl *SkipList) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (sl *SkipList) iterAtPosition(pos uint64) *iterator { _ = "STUB: not implemented"; return nil }

func (sl *SkipList) IterAtPosition(pos uint64) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

func (sl *SkipList) iter(cmp common.Comparator) *iterator { _ = "STUB: not implemented"; return nil }

func (sl *SkipList) Iter(cmp common.Comparator) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

func (sl *SkipList) SplitAt(index uint64) (*SkipList, *SkipList) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ifc interface{}) *SkipList { _ = "STUB: not implemented"; return nil }
