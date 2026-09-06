package dtrie

type Dtrie struct {
	root   *node
	hasher func(v interface{}) uint32
}

type entry struct {
	hash  uint32
	key   interface{}
	value interface{}
}

func (e *entry) KeyHash() uint32 { _ = "STUB: not implemented"; return 0 }

func (e *entry) Key() interface{} { _ = "STUB: not implemented"; return nil }

func (e *entry) Value() interface{} { _ = "STUB: not implemented"; return nil }

func New(hasher func(v interface{}) uint32) *Dtrie { _ = "STUB: not implemented"; return nil }

func (d *Dtrie) Size() (size int) { _ = "STUB: not implemented"; return 0 }

func (d *Dtrie) Get(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (d *Dtrie) Insert(key, value interface{}) *Dtrie { _ = "STUB: not implemented"; return nil }

func (d *Dtrie) Remove(key interface{}) *Dtrie { _ = "STUB: not implemented"; return nil }

func (d *Dtrie) Iterator(stop <-chan struct{}) <-chan Entry { _ = "STUB: not implemented"; return nil }
