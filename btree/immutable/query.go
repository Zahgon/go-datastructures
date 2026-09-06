package btree

func (t *Tr) Apply(fn func(item *Item), keys ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tr) filter(start, stop interface{}, n *Node, fn func(key *Key) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Tr) iter(start, stop interface{}, fn func(*Key) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tr) iterativeFind(value interface{}, id ID) (*path, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tr) iterativeFindWithoutPath(value interface{}, id ID) (*Node, interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
