package link

type Keys []Key

type Key interface {
	Compare(Key) int
}
