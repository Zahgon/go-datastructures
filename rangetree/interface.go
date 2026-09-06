package rangetree

type Entry interface {
	ValueAtDimension(dimension uint64) int64
}

type Interval interface {
	LowAtDimension(dimension uint64) int64

	HighAtDimension(dimension uint64) int64
}

type RangeTree interface {
	Add(entries ...Entry) Entries

	Len() uint64

	Delete(entries ...Entry) Entries

	Query(interval Interval) Entries

	Apply(interval Interval, fn func(Entry) bool)

	Get(entries ...Entry) Entries

	InsertAtDimension(dimension uint64, index, number int64) (Entries, Entries)
}
