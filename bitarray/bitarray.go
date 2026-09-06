package bitarray

type bitArray struct {
	blocks  []block
	lowest  uint64
	highest uint64
	anyset  bool
}

func getIndexAndRemainder(k uint64) (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

func (ba *bitArray) setLowest() { _ = "STUB: not implemented"; return }

func (ba *bitArray) setHighest() { _ = "STUB: not implemented"; return }

func (ba *bitArray) Capacity() uint64 { _ = "STUB: not implemented"; return 0 }

func (ba *bitArray) ToNums() []uint64 { _ = "STUB: not implemented"; return nil }

func (ba *bitArray) SetBit(k uint64) error { _ = "STUB: not implemented"; return nil }

func (ba *bitArray) GetBit(k uint64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (ba *bitArray) GetSetBits(from uint64, buffer []uint64) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

func getSetBitsInBlocks(
	fromBlockIndex, fromOffset uint64,
	blocks []block,
	indices []uint64,
	buffer []uint64,
) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

func (ba *bitArray) ClearBit(k uint64) error { _ = "STUB: not implemented"; return nil }

func (ba *bitArray) Count() int { _ = "STUB: not implemented"; return 0 }

func (ba *bitArray) Or(other BitArray) BitArray { _ = "STUB: not implemented"; return *new(BitArray) }

func (ba *bitArray) And(other BitArray) BitArray { _ = "STUB: not implemented"; return *new(BitArray) }

func (ba *bitArray) Nand(other BitArray) BitArray { _ = "STUB: not implemented"; return *new(BitArray) }

func (ba *bitArray) Reset() { _ = "STUB: not implemented"; return }

func (ba *bitArray) Equals(other BitArray) bool { _ = "STUB: not implemented"; return false }

func (ba *bitArray) Intersects(other BitArray) bool { _ = "STUB: not implemented"; return false }

func (ba *bitArray) Blocks() Iterator { _ = "STUB: not implemented"; return *new(Iterator) }

func (ba *bitArray) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (ba *bitArray) complement() { _ = "STUB: not implemented"; return }

func (ba *bitArray) intersectsSparseBitArray(other *sparseBitArray) bool {
	_ = "STUB: not implemented"
	return false
}

func (ba *bitArray) intersectsDenseBitArray(other *bitArray) bool {
	_ = "STUB: not implemented"
	return false
}

func (ba *bitArray) copy() BitArray { _ = "STUB: not implemented"; return *new(BitArray) }

func newBitArray(size uint64, args ...bool) *bitArray { _ = "STUB: not implemented"; return nil }

func NewBitArray(size uint64, args ...bool) BitArray {
	_ = "STUB: not implemented"
	return *new(BitArray)
}
