package link

func (keys Keys) search(key Key) int { _ = "STUB: not implemented"; return 0 }

func (keys *Keys) insert(key Key) Key { _ = "STUB: not implemented"; return *new(Key) }

func (keys *Keys) insertAt(key Key, i int) Key { _ = "STUB: not implemented"; return *new(Key) }

func (keys *Keys) split() (Key, Keys, Keys) {
	_ = "STUB: not implemented"
	return *new(Key), *new(Keys), *new(Keys)
}

func (keys *Keys) splitAt(i int) (Keys, Keys) {
	_ = "STUB: not implemented"
	return *new(Keys), *new(Keys)
}

func (keys Keys) last() Key { _ = "STUB: not implemented"; return *new(Key) }

func (keys Keys) first() Key { _ = "STUB: not implemented"; return *new(Key) }

func (keys Keys) needsSplit() bool { _ = "STUB: not implemented"; return false }

func (keys Keys) reverse() Keys { _ = "STUB: not implemented"; return *new(Keys) }

func chunkKeys(keys Keys, numParts int64) []Keys { _ = "STUB: not implemented"; return nil }
