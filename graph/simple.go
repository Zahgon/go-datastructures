package graph

import (
	"errors"
	"sync"
)

var (
	ErrVertexNotFound = errors.New("vertex not found")

	ErrSelfLoop = errors.New("self loops not permitted")

	ErrParallelEdge = errors.New("parallel edges are not permitted")
)

type SimpleGraph struct {
	mutex         sync.RWMutex
	adjacencyList map[interface{}]map[interface{}]struct{}
	v, e          int
}

func (g *SimpleGraph) V() int { _ = "STUB: not implemented"; return 0 }

func (g *SimpleGraph) E() int { _ = "STUB: not implemented"; return 0 }

func (g *SimpleGraph) AddEdge(v, w interface{}) error { _ = "STUB: not implemented"; return nil }

func (g *SimpleGraph) Adj(v interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *SimpleGraph) Degree(v interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *SimpleGraph) addVertex(v interface{}) { _ = "STUB: not implemented"; return }

func NewSimpleGraph() *SimpleGraph { _ = "STUB: not implemented"; return nil }
