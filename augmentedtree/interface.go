package augmentedtree

type Interval interface {
	LowAtDimension(uint64) int64

	HighAtDimension(uint64) int64

	OverlapsAtDimension(Interval, uint64) bool

	ID() uint64
}

type Tree interface {
	Add(intervals ...Interval)

	Len() uint64

	Delete(intervals ...Interval)

	Query(interval Interval) Intervals

	Traverse(func(Interval))
}
