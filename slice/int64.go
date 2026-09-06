package slice

type Int64Slice []int64

func (s Int64Slice) Len() int { _ = "STUB: not implemented"; return 0 }

func (s Int64Slice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s Int64Slice) Search(x int64) int { _ = "STUB: not implemented"; return 0 }

func (s Int64Slice) Sort() { _ = "STUB: not implemented"; return }

func (s Int64Slice) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s Int64Slice) Exists(x int64) bool { _ = "STUB: not implemented"; return false }

func (s Int64Slice) Insert(x int64) Int64Slice { _ = "STUB: not implemented"; return *new(Int64Slice) }
