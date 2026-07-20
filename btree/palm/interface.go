package palm

import "github.com/Workiva/go-datastructures/common"

type BTree interface {
	Insert(...common.Comparator)

	Delete(...common.Comparator)

	Get(...common.Comparator) common.Comparators

	Len() uint64

	Query(start, stop common.Comparator) common.Comparators

	Dispose()
}
