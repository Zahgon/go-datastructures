package avl

type Entries []Entry

type Entry interface {
	Compare(Entry) int
}
