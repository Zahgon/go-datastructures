package set

import "sync"

var pool = sync.Pool{}

type Set struct {
	items     map[interface{}]struct{}
	lock      sync.RWMutex
	flattened []interface{}
}

func (set *Set) Add(items ...interface{}) { _ = "STUB: not implemented"; return }

func (set *Set) Remove(items ...interface{}) { _ = "STUB: not implemented"; return }

func (set *Set) Exists(item interface{}) bool { _ = "STUB: not implemented"; return false }

func (set *Set) Flatten() []interface{} { _ = "STUB: not implemented"; return nil }

func (set *Set) Len() int64 { _ = "STUB: not implemented"; return 0 }

func (set *Set) Clear() { _ = "STUB: not implemented"; return }

func (set *Set) All(items ...interface{}) bool { _ = "STUB: not implemented"; return false }

func (set *Set) Dispose() { _ = "STUB: not implemented"; return }

func New(items ...interface{}) *Set { _ = "STUB: not implemented"; return nil }

func init() {
	pool.New = func() interface{} {
		return &Set{
			items: make(map[interface{}]struct{}, 10),
		}
	}
}
