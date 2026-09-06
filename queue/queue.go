package queue

import (
	"sync"
	"time"
)

type waiters []*sema

func (w *waiters) get() *sema { _ = "STUB: not implemented"; return nil }

func (w *waiters) put(sema *sema) { _ = "STUB: not implemented"; return }

func (w *waiters) remove(sema *sema) { _ = "STUB: not implemented"; return }

type items []interface{}

func (items *items) get(number int64) []interface{} { _ = "STUB: not implemented"; return nil }

func (items *items) peek() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (items *items) getUntil(checker func(item interface{}) bool) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

type sema struct {
	ready    chan bool
	response *sync.WaitGroup
}

func newSema() *sema { _ = "STUB: not implemented"; return nil }

type Queue struct {
	waiters  waiters
	items    items
	lock     sync.Mutex
	disposed bool
}

func (q *Queue) Put(items ...interface{}) error { _ = "STUB: not implemented"; return nil }

func (q *Queue) Get(number int64) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *Queue) Poll(number int64, timeout time.Duration) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *Queue) Peek() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *Queue) TakeUntil(checker func(item interface{}) bool) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *Queue) Empty() bool { _ = "STUB: not implemented"; return false }

func (q *Queue) Len() int64 { _ = "STUB: not implemented"; return 0 }

func (q *Queue) Disposed() bool { _ = "STUB: not implemented"; return false }

func (q *Queue) Dispose() []interface{} { _ = "STUB: not implemented"; return nil }

func New(hint int64) *Queue { _ = "STUB: not implemented"; return nil }

func ExecuteInParallel(q *Queue, fn func(interface{})) { _ = "STUB: not implemented"; return }
