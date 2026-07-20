package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/Workiva/go-datastructures/batcher"
)

var _ batcher.Batcher = new(Batcher)

type Batcher struct {
	mock.Mock
	PutChan chan bool
}

func (m *Batcher) Put(items interface{}) error { _ = "STUB: not implemented"; return nil }

func (m *Batcher) Get() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Batcher) Flush() error { _ = "STUB: not implemented"; return nil }

func (m *Batcher) Dispose() { _ = "STUB: not implemented"; return }

func (m *Batcher) IsDisposed() bool { _ = "STUB: not implemented"; return false }
