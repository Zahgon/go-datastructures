package btree

type pathBundle struct {
	i    int
	n    *Node
	prev *pathBundle
}

type path struct {
	head *pathBundle
	tail *pathBundle
}

func (p *path) append(pb *pathBundle) { _ = "STUB: not implemented"; return }

func (p *path) pop() *pathBundle { _ = "STUB: not implemented"; return nil }

func (p *path) peek() *pathBundle { _ = "STUB: not implemented"; return nil }
