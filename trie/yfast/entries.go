package yfast

type entriesWrapper struct {
	key     uint64
	entries Entries
}

func (ew *entriesWrapper) Key() uint64 { _ = "STUB: not implemented"; return 0 }

type Entries []Entry

func (entries Entries) search(key uint64) int { _ = "STUB: not implemented"; return 0 }

func (entries *Entries) insert(entry Entry) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (entries *Entries) delete(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (entries Entries) max() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (entries Entries) get(key uint64) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (entries Entries) successor(key uint64) (Entry, int) {
	_ = "STUB: not implemented"
	return *new(Entry), 0
}

func (entries Entries) predecessor(key uint64) (Entry, int) {
	_ = "STUB: not implemented"
	return *new(Entry), 0
}
