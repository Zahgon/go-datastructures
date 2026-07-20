package bitarray

type Bitmap32 uint32

func (b Bitmap32) SetBit(pos uint) Bitmap32 { _ = "STUB: not implemented"; return *new(Bitmap32) }

func (b Bitmap32) ClearBit(pos uint) Bitmap32 { _ = "STUB: not implemented"; return *new(Bitmap32) }

func (b Bitmap32) GetBit(pos uint) bool { _ = "STUB: not implemented"; return false }

func (b Bitmap32) PopCount() int { _ = "STUB: not implemented"; return 0 }

type Bitmap64 uint64

func (b Bitmap64) SetBit(pos uint) Bitmap64 { _ = "STUB: not implemented"; return *new(Bitmap64) }

func (b Bitmap64) ClearBit(pos uint) Bitmap64 { _ = "STUB: not implemented"; return *new(Bitmap64) }

func (b Bitmap64) GetBit(pos uint) bool { _ = "STUB: not implemented"; return false }

func (b Bitmap64) PopCount() int { _ = "STUB: not implemented"; return 0 }
