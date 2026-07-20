package yfast

import "github.com/Workiva/go-datastructures/trie/xfast"

const iteratorExhausted = -2

func iterExhausted(iter *Iterator) bool { _ = "STUB: not implemented"; return false }

type Iterator struct {
	xfastIterator *xfast.Iterator
	index         int
	entries       *entriesWrapper
}

func (iter *Iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (iter *Iterator) Value() Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (iter *Iterator) exhaust() Entries { _ = "STUB: not implemented"; return *new(Entries) }

func nilIterator() *Iterator { _ = "STUB: not implemented"; return nil }
