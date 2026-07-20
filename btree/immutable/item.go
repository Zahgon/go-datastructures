package btree

type Item struct {
	Value   interface{}
	Payload []byte
}

type items []*Item

func (its items) split(numParts int) []items { _ = "STUB: not implemented"; return nil }
