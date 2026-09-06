package ctrie

import (
	"errors"
	"hash"

	"github.com/Workiva/go-datastructures/list"
)

const (
	w = 5

	exp2 = 32
)

type HashFactory func() hash.Hash32

func defaultHashFactory() hash.Hash32 { _ = "STUB: not implemented"; return *new(hash.Hash32) }

type Ctrie struct {
	root        *iNode
	readOnly    bool
	hashFactory HashFactory
}

type generation struct{ _ int }

type iNode struct {
	main *mainNode
	gen  *generation

	rdcss *rdcssDescriptor
}

func (i *iNode) copyToGen(gen *generation, ctrie *Ctrie) *iNode {
	_ = "STUB: not implemented"
	return nil
}

type mainNode struct {
	cNode  *cNode
	tNode  *tNode
	lNode  *lNode
	failed *mainNode

	prev *mainNode
}

type cNode struct {
	bmp   uint32
	array []branch
	gen   *generation
}

func newMainNode(x *sNode, xhc uint32, y *sNode, yhc uint32, lev uint, gen *generation) *mainNode {
	_ = "STUB: not implemented"
	return nil
}

func (c *cNode) inserted(pos, flag uint32, br branch, gen *generation) *cNode {
	_ = "STUB: not implemented"
	return nil
}

func (c *cNode) updated(pos uint32, br branch, gen *generation) *cNode {
	_ = "STUB: not implemented"
	return nil
}

func (c *cNode) removed(pos, flag uint32, gen *generation) *cNode {
	_ = "STUB: not implemented"
	return nil
}

func (c *cNode) renewed(gen *generation, ctrie *Ctrie) *cNode {
	_ = "STUB: not implemented"
	return nil
}

type tNode struct {
	*sNode
}

func (t *tNode) untombed() *sNode { _ = "STUB: not implemented"; return nil }

type lNode struct {
	list.PersistentList
}

func (l *lNode) entry() *sNode { _ = "STUB: not implemented"; return nil }

func (l *lNode) lookup(e *Entry) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (l *lNode) inserted(entry *Entry) *lNode { _ = "STUB: not implemented"; return nil }

func (l *lNode) removed(e *Entry) *lNode { _ = "STUB: not implemented"; return nil }

func (l *lNode) length() uint { _ = "STUB: not implemented"; return 0 }

type branch interface{}

type Entry struct {
	Key   []byte
	Value interface{}
	hash  uint32
}

type sNode struct {
	*Entry
}

func New(hashFactory HashFactory) *Ctrie { _ = "STUB: not implemented"; return nil }

func newCtrie(root *iNode, hashFactory HashFactory, readOnly bool) *Ctrie {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ctrie) Insert(key []byte, value interface{}) { _ = "STUB: not implemented"; return }

func (c *Ctrie) Lookup(key []byte) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Ctrie) Remove(key []byte) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Ctrie) Snapshot() *Ctrie { _ = "STUB: not implemented"; return nil }

func (c *Ctrie) ReadOnlySnapshot() *Ctrie { _ = "STUB: not implemented"; return nil }

func (c *Ctrie) snapshot(readOnly bool) *Ctrie { _ = "STUB: not implemented"; return nil }

func (c *Ctrie) Clear() { _ = "STUB: not implemented"; return }

func (c *Ctrie) Iterator(cancel <-chan struct{}) <-chan *Entry {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ctrie) Size() uint { _ = "STUB: not implemented"; return 0 }

var errCanceled = errors.New("canceled")

func (c *Ctrie) traverse(i *iNode, ch chan<- *Entry, cancel <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ctrie) assertReadWrite() { _ = "STUB: not implemented"; return }

func (c *Ctrie) insert(entry *Entry) { _ = "STUB: not implemented"; return }

func (c *Ctrie) lookup(entry *Entry) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Ctrie) remove(entry *Entry) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Ctrie) hash(k []byte) uint32 { _ = "STUB: not implemented"; return 0 }

func (c *Ctrie) iinsert(i *iNode, entry *Entry, lev uint, parent *iNode, startGen *generation) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Ctrie) ilookup(i *iNode, entry *Entry, lev uint, parent *iNode, startGen *generation) (interface{}, bool, bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func (c *Ctrie) iremove(i *iNode, entry *Entry, lev uint, parent *iNode, startGen *generation) (interface{}, bool, bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func toContracted(cn *cNode, lev uint) *mainNode { _ = "STUB: not implemented"; return nil }

func toCompressed(cn *cNode, lev uint) *mainNode { _ = "STUB: not implemented"; return nil }

func entomb(m *sNode) *mainNode { _ = "STUB: not implemented"; return nil }

func resurrect(iNode *iNode, main *mainNode) branch { _ = "STUB: not implemented"; return *new(branch) }

func clean(i *iNode, lev uint, ctrie *Ctrie) bool { _ = "STUB: not implemented"; return false }

func cleanReadOnly(tn *tNode, lev uint, p *iNode, ctrie *Ctrie, entry *Entry) (val interface{}, exists bool, ok bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func cleanParent(p, i *iNode, hc uint32, lev uint, ctrie *Ctrie, startGen *generation) {
	_ = "STUB: not implemented"
	return
}

func flagPos(hashcode uint32, lev uint, bmp uint32) (uint32, uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func bitCount(x uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func gcas(in *iNode, old, n *mainNode, ct *Ctrie) bool { _ = "STUB: not implemented"; return false }

func gcasRead(in *iNode, ctrie *Ctrie) *mainNode { _ = "STUB: not implemented"; return nil }

func gcasComplete(i *iNode, m *mainNode, ctrie *Ctrie) *mainNode {
	_ = "STUB: not implemented"
	return nil
}

type rdcssDescriptor struct {
	old       *iNode
	expected  *mainNode
	nv        *iNode
	committed int32
}

func (c *Ctrie) readRoot() *iNode { _ = "STUB: not implemented"; return nil }

func (c *Ctrie) rdcssReadRoot(abort bool) *iNode { _ = "STUB: not implemented"; return nil }

func (c *Ctrie) rdcssRoot(old *iNode, expected *mainNode, nv *iNode) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Ctrie) rdcssComplete(abort bool) *iNode { _ = "STUB: not implemented"; return nil }

func (c *Ctrie) casRoot(ov, nv *iNode) bool { _ = "STUB: not implemented"; return false }
