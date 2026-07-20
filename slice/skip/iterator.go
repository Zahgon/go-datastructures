package skip

import "github.com/Workiva/go-datastructures/common"

const iteratorExhausted = -2

type iterator struct {
	first bool
	n     *node
}

func (iter *iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (iter *iterator) Value() common.Comparator {
	_ = "STUB: not implemented"
	return *new(common.Comparator)
}

func (iter *iterator) exhaust() common.Comparators {
	_ = "STUB: not implemented"
	return *new(common.Comparators)
}

func nilIterator() *iterator { _ = "STUB: not implemented"; return nil }
