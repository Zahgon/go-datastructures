package avl

type nodes []*node

func (ns nodes) reset() { _ = "STUB: not implemented"; return }

type node struct {
	balance  int8
	children [2]*node
	entry    Entry
}

func (n *node) copy() *node { _ = "STUB: not implemented"; return nil }

func newNode(entry Entry) *node { _ = "STUB: not implemented"; return nil }
