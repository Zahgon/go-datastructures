package hilbert

import (
	"sync"

	"github.com/Workiva/go-datastructures/rtree"
)

type actions []action

type action interface {
	operation() operation
	keys() hilberts
	rects() []*hilbertBundle
	complete()
	addNode(int64, *node)
	nodes() []*node
}

type getAction struct {
	result    rtree.Rectangles
	completer *sync.WaitGroup
	lookup    *rectangle
}

func (ga *getAction) complete() { _ = "STUB: not implemented"; return }

func (ga *getAction) operation() operation { _ = "STUB: not implemented"; return *new(operation) }

func (ga *getAction) keys() hilberts { _ = "STUB: not implemented"; return *new(hilberts) }

func (ga *getAction) addNode(i int64, n *node) { _ = "STUB: not implemented"; return }

func (ga *getAction) nodes() []*node { _ = "STUB: not implemented"; return nil }

func (ga *getAction) rects() []*hilbertBundle { _ = "STUB: not implemented"; return nil }

func newGetAction(rect rtree.Rectangle) *getAction { _ = "STUB: not implemented"; return nil }

type insertAction struct {
	rs        []*hilbertBundle
	completer *sync.WaitGroup
	ns        []*node
}

func (ia *insertAction) complete() { _ = "STUB: not implemented"; return }

func (ia *insertAction) operation() operation { _ = "STUB: not implemented"; return *new(operation) }

func (ia *insertAction) keys() hilberts { _ = "STUB: not implemented"; return *new(hilberts) }

func (ia *insertAction) addNode(i int64, n *node) { _ = "STUB: not implemented"; return }

func (ia *insertAction) nodes() []*node { _ = "STUB: not implemented"; return nil }

func (ia *insertAction) rects() []*hilbertBundle { _ = "STUB: not implemented"; return nil }

func newInsertAction(rects rtree.Rectangles) *insertAction { _ = "STUB: not implemented"; return nil }

type removeAction struct {
	*insertAction
}

func (ra *removeAction) operation() operation { _ = "STUB: not implemented"; return *new(operation) }

func newRemoveAction(rects rtree.Rectangles) *removeAction { _ = "STUB: not implemented"; return nil }

func minUint64(choices ...uint64) uint64 { _ = "STUB: not implemented"; return 0 }

type interfaces []interface{}

func executeInterfacesInParallel(ifs interfaces, fn func(interface{})) {
	_ = "STUB: not implemented"
	return
}

func executeInterfacesInSerial(ifs interfaces, fn func(interface{})) {
	_ = "STUB: not implemented"
	return
}
