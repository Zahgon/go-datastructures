package btree

type Config struct {
	NodeWidth int

	Persister Persister

	Comparator Comparator `msg:"-"`
}

func DefaultConfig(persister Persister, comparator Comparator) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}
