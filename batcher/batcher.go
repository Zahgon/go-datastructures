package batcher

import (
	"errors"
	"time"
)

type mutex struct {
	lock chan struct{}
}

func newMutex() *mutex { _ = "STUB: not implemented"; return nil }

func (m *mutex) Lock() { _ = "STUB: not implemented"; return }

func (m *mutex) Unlock() { _ = "STUB: not implemented"; return }

func (m *mutex) TryLock() bool { _ = "STUB: not implemented"; return false }

type Batcher interface {
	Put(interface{}) error

	Get() ([]interface{}, error)

	Flush() error

	Dispose()

	IsDisposed() bool
}

var ErrDisposed = errors.New("batcher: disposed")

type CalculateBytes func(interface{}) uint

type basicBatcher struct {
	maxTime        time.Duration
	maxItems       uint
	maxBytes       uint
	calculateBytes CalculateBytes
	disposed       bool
	items          []interface{}
	batchChan      chan []interface{}
	availableBytes uint
	lock           *mutex
}

func New(maxTime time.Duration, maxItems, maxBytes, queueLen uint, calculate CalculateBytes) (Batcher, error) {
	_ = "STUB: not implemented"
	return *new(Batcher), nil
}

func (b *basicBatcher) Put(item interface{}) error { _ = "STUB: not implemented"; return nil }

func (b *basicBatcher) Get() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *basicBatcher) Flush() error { _ = "STUB: not implemented"; return nil }

func (b *basicBatcher) Dispose() { _ = "STUB: not implemented"; return }

func (b *basicBatcher) IsDisposed() bool { _ = "STUB: not implemented"; return false }

func (b *basicBatcher) flush() { _ = "STUB: not implemented"; return }

func (b *basicBatcher) ready() bool { _ = "STUB: not implemented"; return false }

func (b *basicBatcher) drainBatchChan() { _ = "STUB: not implemented"; return }
