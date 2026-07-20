package plus

const iteratorExhausted = -2

type iterator struct {
	node  *lnode
	index int
}

func (iter *iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (iter *iterator) Value() Key { _ = "STUB: not implemented"; return *new(Key) }

func (iter *iterator) exhaust() keys { _ = "STUB: not implemented"; return *new(keys) }

func nilIterator() *iterator { _ = "STUB: not implemented"; return nil }
