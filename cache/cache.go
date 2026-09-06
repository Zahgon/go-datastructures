package cache

import (
	"container/list"
	"sync"
)

type Cache interface {
	Get(keys ...string) []Item

	Put(key string, item Item)

	Remove(keys ...string)

	Size() uint64
}

type Item interface {
	Size() uint64
}

type cached struct {
	item    Item
	element *list.Element
}

func (c *cached) setElementIfNotNil(element *list.Element) { _ = "STUB: not implemented"; return }

type cache struct {
	sync.Mutex
	cap          uint64
	size         uint64
	items        map[string]*cached
	keyList      *list.List
	recordAdd    func(key string) *list.Element
	recordAccess func(key string) *list.Element
}

type CacheOption func(*cache)

type Policy uint8

const (
	LeastRecentlyAdded Policy = iota

	LeastRecentlyUsed
)

func EvictionPolicy(policy Policy) CacheOption { _ = "STUB: not implemented"; return *new(CacheOption) }

func New(capacity uint64, options ...CacheOption) Cache {
	_ = "STUB: not implemented"
	return *new(Cache)
}

func (c *cache) Get(keys ...string) []Item { _ = "STUB: not implemented"; return nil }

func (c *cache) Put(key string, item Item) { _ = "STUB: not implemented"; return }

func (c *cache) Remove(keys ...string) { _ = "STUB: not implemented"; return }

func (c *cache) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *cache) ensureCapacity(toAdd uint64) { _ = "STUB: not implemented"; return }

func (c *cache) remove(key string) { _ = "STUB: not implemented"; return }

func (c *cache) noop(string) *list.Element { _ = "STUB: not implemented"; return nil }

func (c *cache) record(key string) *list.Element { _ = "STUB: not implemented"; return nil }
