package bitarray

func Marshal(ba BitArray) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func Unmarshal(input []byte) (BitArray, error) {
	_ = "STUB: not implemented"
	return *new(BitArray), nil
}

func (ba *sparseBitArray) Serialize() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func Uint64FromBytes(b []byte) (uint64, int) { _ = "STUB: not implemented"; return 0, 0 }

func (ret *sparseBitArray) Deserialize(incoming []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (ba *bitArray) Serialize() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ret *bitArray) Deserialize(incoming []byte) error { _ = "STUB: not implemented"; return nil }
