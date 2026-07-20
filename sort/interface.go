package merge

type Comparators []Comparator

func (c Comparators) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (c Comparators) Len() int { _ = "STUB: not implemented"; return 0 }

func (c Comparators) Swap(i, j int) { _ = "STUB: not implemented"; return }

type Comparator interface {
	Compare(Comparator) int
}
