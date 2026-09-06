package plus

type Keys []Key

type Key interface {
	Compare(Key) int
}

type Iterator interface {
	Next() bool

	Value() Key

	exhaust() keys
}
