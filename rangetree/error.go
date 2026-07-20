package rangetree

type NoEntriesError struct{}

func (nee NoEntriesError) Error() string { _ = "STUB: not implemented"; return "" }

type OutOfDimensionError struct {
	provided, max uint64
}

func (oode OutOfDimensionError) Error() string { _ = "STUB: not implemented"; return "" }
