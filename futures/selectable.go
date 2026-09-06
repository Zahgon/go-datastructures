package futures

import (
	"errors"
	"sync"
)

var ErrFutureCanceled = errors.New("future canceled")

type Selectable struct {
	m      sync.Mutex
	val    interface{}
	err    error
	wait   chan struct{}
	filled uint32
}

func NewSelectable() *Selectable { _ = "STUB: not implemented"; return nil }

func (f *Selectable) wchan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (f *Selectable) WaitChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (f *Selectable) GetResult() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Selectable) Fill(v interface{}, e error) error { _ = "STUB: not implemented"; return nil }

func (f *Selectable) SetValue(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (f *Selectable) SetError(e error) { _ = "STUB: not implemented"; return }

func (f *Selectable) Cancel() { _ = "STUB: not implemented"; return }

var closed = make(chan struct{})

func init() {
	close(closed)
}
