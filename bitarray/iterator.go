package bitarray

type sparseBitArrayIterator struct {
	index int64
	sba   *sparseBitArray
}

func (iter *sparseBitArrayIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (iter *sparseBitArrayIterator) Value() (uint64, block) {
	_ = "STUB: not implemented"
	return 0, *new(block)
}

func newCompressedBitArrayIterator(sba *sparseBitArray) *sparseBitArrayIterator {
	_ = "STUB: not implemented"
	return nil
}

type bitArrayIterator struct {
	index     int64
	stopIndex uint64
	ba        *bitArray
}

func (iter *bitArrayIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (iter *bitArrayIterator) Value() (uint64, block) {
	_ = "STUB: not implemented"
	return 0, *new(block)
}

func newBitArrayIterator(ba *bitArray) *bitArrayIterator { _ = "STUB: not implemented"; return nil }
