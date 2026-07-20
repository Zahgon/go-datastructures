package augmentedtree

import "sync"

var intervalsPool = sync.Pool{
	New: func() interface{} {
		return make(Intervals, 0, 10)
	},
}

type Intervals []Interval

func (ivs *Intervals) Dispose() { _ = "STUB: not implemented"; return }
