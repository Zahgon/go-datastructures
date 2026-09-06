package rangetree

type nodes []*node

type node struct {
	value        int64
	entry        Entry
	orderedNodes orderedNodes
}

func newNode(value int64, entry Entry, needNextDimension bool) *node {
	_ = "STUB: not implemented"
	return nil
}
