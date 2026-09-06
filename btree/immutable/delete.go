package btree

func (t *Tr) DeleteItems(values ...interface{}) ([]*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tr) delete(keys Keys) error { _ = "STUB: not implemented"; return nil }

func (t *Tr) walkupDelete(key *Key, node *Node, path *path, mapping map[string]*Node) error {
	_ = "STUB: not implemented"
	return nil
}
