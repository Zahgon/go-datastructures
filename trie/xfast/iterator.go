package xfast

type Entries []Entry

type Iterator struct {
	n     *node
	first bool
}

func (iter *Iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (iter *Iterator) Value() Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (iter *Iterator) exhaust() Entries { _ = "STUB: not implemented"; return *new(Entries) }
