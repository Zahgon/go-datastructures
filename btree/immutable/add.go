package btree

func (t *Tr) AddItems(its ...*Item) ([]*Item, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tr) add(keys Keys) (Keys, error) { _ = "STUB: not implemented"; return *new(Keys), nil }

func (t *Tr) determinePaths(keys Keys) (map[string][]*nodeBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func insertByMerge(comparator Comparator, n *Node, bundles []*nodeBundle) (Keys, error) {
	_ = "STUB: not implemented"
	return *new(Keys), nil
}

func insertLastDimension(t *Tr, n *Node, bundles []*nodeBundle) (Keys, error) {
	_ = "STUB: not implemented"
	return *new(Keys), nil
}

func (t *Tr) iterativeSplit(n *Node) Keys { _ = "STUB: not implemented"; return *new(Keys) }

func (t *Tr) walkupInsert(nodes map[string]*path) error { _ = "STUB: not implemented"; return nil }
