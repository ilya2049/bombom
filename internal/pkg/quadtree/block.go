package quadtree

type block interface {
	X() int
	Y() int
	Size() int
}
