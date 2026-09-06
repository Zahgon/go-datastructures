package list

import "errors"

var (
	Empty PersistentList = &emptyList{}

	ErrEmptyList = errors.New("Empty list")
)

type PersistentList interface {
	Head() (interface{}, bool)

	Tail() (PersistentList, bool)

	IsEmpty() bool

	Length() uint

	Add(head interface{}) PersistentList

	Insert(val interface{}, pos uint) (PersistentList, error)

	Get(pos uint) (interface{}, bool)

	Remove(pos uint) (PersistentList, error)

	Find(func(interface{}) bool) (interface{}, bool)

	FindIndex(func(interface{}) bool) int

	Map(func(interface{}) interface{}) []interface{}
}

type emptyList struct{}

func (e *emptyList) Head() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (e *emptyList) Tail() (PersistentList, bool) {
	_ = "STUB: not implemented"
	return *new(PersistentList), false
}

func (e *emptyList) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (e *emptyList) Length() uint { _ = "STUB: not implemented"; return 0 }

func (e *emptyList) Add(head interface{}) PersistentList {
	_ = "STUB: not implemented"
	return *new(PersistentList)
}

func (e *emptyList) Insert(val interface{}, pos uint) (PersistentList, error) {
	_ = "STUB: not implemented"
	return *new(PersistentList), nil
}

func (e *emptyList) Get(pos uint) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (e *emptyList) Remove(pos uint) (PersistentList, error) {
	_ = "STUB: not implemented"
	return *new(PersistentList), nil
}

func (e *emptyList) Find(func(interface{}) bool) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (e *emptyList) FindIndex(func(interface{}) bool) int { _ = "STUB: not implemented"; return 0 }

func (e *emptyList) Map(func(interface{}) interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

type list struct {
	head interface{}
	tail PersistentList
}

func (l *list) Head() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (l *list) Tail() (PersistentList, bool) {
	_ = "STUB: not implemented"
	return *new(PersistentList), false
}

func (l *list) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (l *list) Length() uint { _ = "STUB: not implemented"; return 0 }

func (l *list) Add(head interface{}) PersistentList {
	_ = "STUB: not implemented"
	return *new(PersistentList)
}

func (l *list) Insert(val interface{}, pos uint) (PersistentList, error) {
	_ = "STUB: not implemented"
	return *new(PersistentList), nil
}

func (l *list) Get(pos uint) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (l *list) Remove(pos uint) (PersistentList, error) {
	_ = "STUB: not implemented"
	return *new(PersistentList), nil
}

func (l *list) Find(pred func(interface{}) bool) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (l *list) FindIndex(pred func(interface{}) bool) int { _ = "STUB: not implemented"; return 0 }

func (l *list) Map(f func(interface{}) interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}
