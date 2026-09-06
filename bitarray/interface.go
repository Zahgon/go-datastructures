package bitarray

type BitArray interface {
	SetBit(k uint64) error

	GetBit(k uint64) (bool, error)

	GetSetBits(from uint64, buffer []uint64) []uint64

	ClearBit(k uint64) error

	Reset()

	Blocks() Iterator

	Equals(other BitArray) bool

	Intersects(other BitArray) bool

	Capacity() uint64

	Count() int

	Or(other BitArray) BitArray

	And(other BitArray) BitArray

	Nand(other BitArray) BitArray

	ToNums() []uint64

	IsEmpty() bool
}

type Iterator interface {
	Next() bool

	Value() (uint64, block)
}
