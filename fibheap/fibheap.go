package fibheap

type FloatingFibonacciHeap struct {
	min  *Entry
	size uint
}

type Entry struct {
	degree                    int
	marked                    bool
	next, prev, child, parent *Entry

	Priority float64
}

type EmptyHeapError string

func (e EmptyHeapError) Error() string { _ = "STUB: not implemented"; return "" }

type NilError string

func (e NilError) Error() string { _ = "STUB: not implemented"; return "" }

func NewFloatFibHeap() FloatingFibonacciHeap {
	_ = "STUB: not implemented"
	return *new(FloatingFibonacciHeap)
}

func (heap *FloatingFibonacciHeap) Enqueue(priority float64) *Entry {
	_ = "STUB: not implemented"
	return nil
}

func (heap *FloatingFibonacciHeap) Min() (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (heap *FloatingFibonacciHeap) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (heap *FloatingFibonacciHeap) Size() uint { _ = "STUB: not implemented"; return 0 }

func (heap *FloatingFibonacciHeap) DequeueMin() (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (heap *FloatingFibonacciHeap) DecreaseKey(node *Entry, newPriority float64) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (heap *FloatingFibonacciHeap) Delete(node *Entry) error { _ = "STUB: not implemented"; return nil }

func (heap *FloatingFibonacciHeap) Merge(other *FloatingFibonacciHeap) (FloatingFibonacciHeap, error) {
	_ = "STUB: not implemented"
	return *new(FloatingFibonacciHeap), nil
}

func newEntry(priority float64) *Entry { _ = "STUB: not implemented"; return nil }

func mergeLists(one, two *Entry) *Entry { _ = "STUB: not implemented"; return nil }

func decreaseKeyUnchecked(heap *FloatingFibonacciHeap, node *Entry, priority float64) {
	_ = "STUB: not implemented"
	return
}

func cutNode(heap *FloatingFibonacciHeap, node *Entry) { _ = "STUB: not implemented"; return }
