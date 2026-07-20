package bitarray

import (
	"unsafe"
)

type block uint64

const s = uint64(unsafe.Sizeof(block(0)) * 8)

const maximumBlock = block(0) | ^block(0)

func (b block) toNums(offset uint64, nums *[]uint64) { _ = "STUB: not implemented"; return }

func (b block) findLeftPosition() uint64 { _ = "STUB: not implemented"; return 0 }

func (b block) findRightPosition() uint64 { _ = "STUB: not implemented"; return 0 }

func (b block) insert(position uint64) block { _ = "STUB: not implemented"; return *new(block) }

func (b block) remove(position uint64) block { _ = "STUB: not implemented"; return *new(block) }

func (b block) or(other block) block { _ = "STUB: not implemented"; return *new(block) }

func (b block) and(other block) block { _ = "STUB: not implemented"; return *new(block) }

func (b block) nand(other block) block { _ = "STUB: not implemented"; return *new(block) }

func (b block) get(position uint64) bool { _ = "STUB: not implemented"; return false }

func (b block) equals(other block) bool { _ = "STUB: not implemented"; return false }

func (b block) intersects(other block) bool { _ = "STUB: not implemented"; return false }

func (b block) String() string { _ = "STUB: not implemented"; return "" }
