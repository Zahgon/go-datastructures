package rangetree

import "sync"

var entriesPool = sync.Pool{
	New: func() interface{} {
		return make(Entries, 0, 10)
	},
}

type Entries []Entry

func (entries *Entries) Dispose() { _ = "STUB: not implemented"; return }

func NewEntries() Entries { _ = "STUB: not implemented"; return *new(Entries) }
