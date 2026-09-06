package yfast

import "github.com/Workiva/go-datastructures/trie/xfast"

type YFastTrie struct {
	num   uint64
	xfast *xfast.XFastTrie
	bits  uint8
}

func (yfast *YFastTrie) init(intType interface{}) {
	switch intType.(type) {
	case uint8:
		yfast.bits = 8
	case uint16:
		yfast.bits = 16
	case uint32:
		yfast.bits = 32
	case uint, uint64:
		yfast.bits = 64
	default:

		panic(`Invalid universe size provided.`)
	}

	yfast.xfast = xfast.New(intType)
}

func (yfast *YFastTrie) getBucketKey(key uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (yfast *YFastTrie) insert(entry Entry) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (yfast *YFastTrie) Insert(entries ...Entry) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (yfast *YFastTrie) delete(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (yfast *YFastTrie) Delete(keys ...uint64) Entries {
	_ = "STUB: not implemented"
	return *new(Entries)
}

func (yfast *YFastTrie) get(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (yfast *YFastTrie) Get(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (yfast *YFastTrie) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (yfast *YFastTrie) successor(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (yfast *YFastTrie) Successor(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (yfast *YFastTrie) predecessor(key uint64) Entry {
	_ = "STUB: not implemented"
	return *new(Entry)
}

func (yfast *YFastTrie) Predecessor(key uint64) Entry {
	_ = "STUB: not implemented"
	return *new(Entry)
}

func (yfast *YFastTrie) iter(key uint64) *Iterator { _ = "STUB: not implemented"; return nil }

func (yfast *YFastTrie) Iter(key uint64) *Iterator { _ = "STUB: not implemented"; return nil }

func New(ifc interface{}) *YFastTrie { _ = "STUB: not implemented"; return nil }
