package btree

type Tree interface {
	Apply(fn func(item *Item), keys ...interface{}) error

	ID() ID

	Len() int
}

type ReadableTree interface {
	Tree

	AsMutable() MutableTree
}

type MutableTree interface {
	Tree

	Commit() (ReadableTree, error)

	AddItems(items ...*Item) ([]*Item, error)

	DeleteItems(keys ...interface{}) ([]*Item, error)
}

type Comparator func(item1, item2 interface{}) int

type Payload struct {
	Key     []byte
	Payload []byte
}

type Persister interface {
	Save(items ...*Payload) error
	Load(keys ...[]byte) ([]*Payload, error)
}
