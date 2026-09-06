package palm

import (
	"sync"

	"github.com/Workiva/go-datastructures/common"
)

type actions []action

type action interface {
	operation() operation
	keys() common.Comparators
	complete()
	addNode(int64, *node)
	nodes() []*node
}

type getAction struct {
	result    common.Comparators
	completer *sync.WaitGroup
}

func (ga *getAction) complete() { _ = "STUB: not implemented"; return }

func (ga *getAction) operation() operation { _ = "STUB: not implemented"; return *new(operation) }

func (ga *getAction) keys() common.Comparators {
	_ = "STUB: not implemented"
	return *new(common.Comparators)
}

func (ga *getAction) addNode(i int64, n *node) { _ = "STUB: not implemented"; return }

func (ga *getAction) nodes() []*node { _ = "STUB: not implemented"; return nil }

func newGetAction(keys common.Comparators) *getAction { _ = "STUB: not implemented"; return nil }

type insertAction struct {
	result    common.Comparators
	completer *sync.WaitGroup
	ns        []*node
}

func (ia *insertAction) complete() { _ = "STUB: not implemented"; return }

func (ia *insertAction) operation() operation { _ = "STUB: not implemented"; return *new(operation) }

func (ia *insertAction) keys() common.Comparators {
	_ = "STUB: not implemented"
	return *new(common.Comparators)
}

func (ia *insertAction) addNode(i int64, n *node) { _ = "STUB: not implemented"; return }

func (ia *insertAction) nodes() []*node { _ = "STUB: not implemented"; return nil }

func newInsertAction(keys common.Comparators) *insertAction { _ = "STUB: not implemented"; return nil }

type removeAction struct {
	*insertAction
}

func (ra *removeAction) operation() operation { _ = "STUB: not implemented"; return *new(operation) }

func newRemoveAction(keys common.Comparators) *removeAction { _ = "STUB: not implemented"; return nil }

type applyAction struct {
	fn          func(common.Comparator) bool
	start, stop common.Comparator
	completer   *sync.WaitGroup
}

func (aa *applyAction) operation() operation { _ = "STUB: not implemented"; return *new(operation) }

func (aa *applyAction) nodes() []*node { _ = "STUB: not implemented"; return nil }

func (aa *applyAction) addNode(i int64, n *node) { _ = "STUB: not implemented"; return }

func (aa *applyAction) keys() common.Comparators {
	_ = "STUB: not implemented"
	return *new(common.Comparators)
}

func (aa *applyAction) complete() { _ = "STUB: not implemented"; return }

func newApplyAction(fn func(common.Comparator) bool, start, stop common.Comparator) *applyAction {
	_ = "STUB: not implemented"
	return nil
}

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
