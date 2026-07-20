package skip

import "github.com/Workiva/go-datastructures/common"

type Iterator interface {
	Next() bool

	Value() common.Comparator

	exhaust() common.Comparators
}
