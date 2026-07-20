package bitarray

type uintSlice []uint64

func (u uintSlice) Len() int64 { _ = "STUB: not implemented"; return 0 }

func (u uintSlice) Swap(i, j int64) { _ = "STUB: not implemented"; return }

func (u uintSlice) Less(i, j int64) bool { _ = "STUB: not implemented"; return false }

func (u uintSlice) search(x uint64) int64 { _ = "STUB: not implemented"; return 0 }

func (u *uintSlice) insert(x uint64) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func (u *uintSlice) deleteAtIndex(i int64) { _ = "STUB: not implemented"; return }

func (u uintSlice) get(x uint64) int64 { _ = "STUB: not implemented"; return 0 }

type blocks []block

func (b *blocks) insert(index int64) { _ = "STUB: not implemented"; return }

func (b *blocks) deleteAtIndex(i int64) { _ = "STUB: not implemented"; return }

type sparseBitArray struct {
	blocks  blocks
	indices uintSlice
}

func (sba *sparseBitArray) SetBit(k uint64) error { _ = "STUB: not implemented"; return nil }

func (sba *sparseBitArray) GetBit(k uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (sba *sparseBitArray) GetSetBits(from uint64, buffer []uint64) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

func (sba *sparseBitArray) ToNums() []uint64 { _ = "STUB: not implemented"; return nil }

func (sba *sparseBitArray) ClearBit(k uint64) error { _ = "STUB: not implemented"; return nil }

func (sba *sparseBitArray) Reset() { _ = "STUB: not implemented"; return }

func (sba *sparseBitArray) Blocks() Iterator { _ = "STUB: not implemented"; return *new(Iterator) }

func (sba *sparseBitArray) Capacity() uint64 { _ = "STUB: not implemented"; return 0 }

func (sba *sparseBitArray) Equals(other BitArray) bool { _ = "STUB: not implemented"; return false }

func (sba *sparseBitArray) Count() int { _ = "STUB: not implemented"; return 0 }

func (sba *sparseBitArray) Or(other BitArray) BitArray {
	_ = "STUB: not implemented"
	return *new(BitArray)
}

func (sba *sparseBitArray) And(other BitArray) BitArray {
	_ = "STUB: not implemented"
	return *new(BitArray)
}

func (sba *sparseBitArray) Nand(other BitArray) BitArray {
	_ = "STUB: not implemented"
	return *new(BitArray)
}

func (sba *sparseBitArray) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (sba *sparseBitArray) copy() *sparseBitArray { _ = "STUB: not implemented"; return nil }

func (sba *sparseBitArray) Intersects(other BitArray) bool { _ = "STUB: not implemented"; return false }

func (sba *sparseBitArray) IntersectsBetween(other BitArray, start, stop uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func newSparseBitArray() *sparseBitArray { _ = "STUB: not implemented"; return nil }

func NewSparseBitArray() BitArray { _ = "STUB: not implemented"; return *new(BitArray) }
