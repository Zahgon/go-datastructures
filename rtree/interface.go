package rtree

type Rectangles []Rectangle

type Rectangle interface {
	LowerLeft() (int32, int32)

	UpperRight() (int32, int32)
}

type RTree interface {
	Search(Rectangle) Rectangles

	Len() uint64

	Dispose()

	Delete(...Rectangle)

	Insert(...Rectangle)
}
