package btree

import (
	"sync"

	"github.com/Workiva/go-datastructures/futures"
)

type cacher struct {
	lock      sync.Mutex
	cache     map[string]*futures.Future
	persister Persister
}

func (c *cacher) asyncLoadNode(t *Tr, key ID, completer chan interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *cacher) clear() { _ = "STUB: not implemented"; return }

func (c *cacher) deleteFromCache(id ID) { _ = "STUB: not implemented"; return }

func (c *cacher) loadNode(t *Tr, key ID) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *cacher) getNode(t *Tr, key ID, useCache bool) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCacher(persister Persister) *cacher { _ = "STUB: not implemented"; return nil }
