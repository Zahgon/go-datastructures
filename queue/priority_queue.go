package queue

import "sync"

type Item interface {
	Compare(other Item) int
}

type priorityItems []Item

func (items *priorityItems) swap(i, j int) { _ = "STUB: not implemented"; return }

func (items *priorityItems) pop() Item { _ = "STUB: not implemented"; return *new(Item) }

func (items *priorityItems) get(number int) []Item { _ = "STUB: not implemented"; return nil }

func (items *priorityItems) push(item Item) { _ = "STUB: not implemented"; return }

type PriorityQueue struct {
	waiters         waiters
	items           priorityItems
	itemMap         map[Item]struct{}
	lock            sync.Mutex
	disposeLock     sync.Mutex
	disposed        bool
	allowDuplicates bool
}

func (pq *PriorityQueue) Put(items ...Item) error { _ = "STUB: not implemented"; return nil }

func (pq *PriorityQueue) Get(number int) ([]Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pq *PriorityQueue) Peek() Item { _ = "STUB: not implemented"; return *new(Item) }

func (pq *PriorityQueue) Empty() bool { _ = "STUB: not implemented"; return false }

func (pq *PriorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq *PriorityQueue) Disposed() bool { _ = "STUB: not implemented"; return false }

func (pq *PriorityQueue) Dispose() { _ = "STUB: not implemented"; return }

func NewPriorityQueue(hint int, allowDuplicates bool) *PriorityQueue {
	_ = "STUB: not implemented"
	return nil
}
