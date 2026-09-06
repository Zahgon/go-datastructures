package fastinteger

const ratio = .75

func roundUp(v uint64) uint64 { _ = "STUB: not implemented"; return 0 }

type packet struct {
	key, value uint64
}

type packets []*packet

func (packets packets) find(key uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (packets packets) set(packet *packet) { _ = "STUB: not implemented"; return }

func (packets packets) get(key uint64) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (packets packets) delete(key uint64) bool { _ = "STUB: not implemented"; return false }

func (packets packets) exists(key uint64) bool { _ = "STUB: not implemented"; return false }

type FastIntegerHashMap struct {
	count   uint64
	packets packets
}

func (fi *FastIntegerHashMap) rebuild() { _ = "STUB: not implemented"; return }

func (fi *FastIntegerHashMap) Get(key uint64) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (fi *FastIntegerHashMap) Set(key, value uint64) { _ = "STUB: not implemented"; return }

func (fi *FastIntegerHashMap) Exists(key uint64) bool { _ = "STUB: not implemented"; return false }

func (fi *FastIntegerHashMap) Delete(key uint64) { _ = "STUB: not implemented"; return }

func (fi *FastIntegerHashMap) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (fi *FastIntegerHashMap) Cap() uint64 { _ = "STUB: not implemented"; return 0 }

func New(hint uint64) *FastIntegerHashMap { _ = "STUB: not implemented"; return nil }
