package link

import (
	"log"
	"sync"
)

const numberOfItemsBeforeMultithread = 10

type blink struct {
	root                     *node
	lock                     sync.RWMutex
	number, ary, numRoutines uint64
}

func (blink *blink) insert(key Key, stack *nodes) Key { _ = "STUB: not implemented"; return *new(Key) }

func (blink *blink) multithreadedInsert(keys Keys) Keys {
	_ = "STUB: not implemented"
	return *new(Keys)
}

func (blink *blink) Insert(keys ...Key) Keys { _ = "STUB: not implemented"; return *new(Keys) }

func (blink *blink) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (blink *blink) get(key Key) Key { _ = "STUB: not implemented"; return *new(Key) }

func (blink *blink) Get(keys ...Key) Keys { _ = "STUB: not implemented"; return *new(Keys) }

func (blink *blink) print(output *log.Logger) { _ = "STUB: not implemented"; return }

func newTree(ary, numRoutines uint64) *blink { _ = "STUB: not implemented"; return nil }
