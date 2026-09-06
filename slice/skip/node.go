package skip

import "github.com/Workiva/go-datastructures/common"

type widths []uint64

type nodes []*node

type node struct {
	forward nodes

	widths widths

	entry common.Comparator
}

func (n *node) Compare(e common.Comparator) int { _ = "STUB: not implemented"; return 0 }

func newNode(cmp common.Comparator, maxLevels uint8) *node { _ = "STUB: not implemented"; return nil }
