package err

import "sync"

type Error struct {
	lock sync.RWMutex
	err  error
}

func (e *Error) Set(err error) { _ = "STUB: not implemented"; return }

func (e *Error) Get() error { _ = "STUB: not implemented"; return nil }

func New() *Error { _ = "STUB: not implemented"; return nil }
