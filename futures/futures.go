package futures

import (
	"sync"
	"time"
)

type Completer <-chan interface{}

type Future struct {
	triggered bool
	item      interface{}
	err       error
	lock      sync.Mutex
	wg        sync.WaitGroup
}

func (f *Future) GetResult() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Future) HasResult() bool { _ = "STUB: not implemented"; return false }

func (f *Future) setItem(item interface{}, err error) { _ = "STUB: not implemented"; return }

func listenForResult(f *Future, ch Completer, timeout time.Duration, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func New(completer Completer, timeout time.Duration) *Future { _ = "STUB: not implemented"; return nil }
