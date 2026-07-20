package queue

import (
	"time"
)

func roundUp(v uint64) uint64 { _ = "STUB: not implemented"; return 0 }

type node struct {
	position uint64
	data     interface{}
}

type nodes []node

type RingBuffer struct {
	_padding0      [8]uint64
	queue          uint64
	_padding1      [8]uint64
	dequeue        uint64
	_padding2      [8]uint64
	mask, disposed uint64
	_padding3      [8]uint64
	nodes          nodes
}

func (rb *RingBuffer) init(size uint64) {
	size = roundUp(size)
	rb.nodes = make(nodes, size)
	for i := uint64(0); i < size; i++ {
		rb.nodes[i] = node{position: i}
	}
	rb.mask = size - 1
}

func (rb *RingBuffer) Put(item interface{}) error { _ = "STUB: not implemented"; return nil }

func (rb *RingBuffer) Offer(item interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rb *RingBuffer) put(item interface{}, offer bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rb *RingBuffer) Get() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (rb *RingBuffer) Poll(timeout time.Duration) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rb *RingBuffer) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (rb *RingBuffer) Cap() uint64 { _ = "STUB: not implemented"; return 0 }

func (rb *RingBuffer) Dispose() { _ = "STUB: not implemented"; return }

func (rb *RingBuffer) IsDisposed() bool { _ = "STUB: not implemented"; return false }

func NewRingBuffer(size uint64) *RingBuffer { _ = "STUB: not implemented"; return nil }
