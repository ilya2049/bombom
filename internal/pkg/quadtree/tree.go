package quadtree

import "fmt"

type Tree[B block] struct {
	root *plane[B]
}

func New[B block](size int) *Tree[B] {
	if debug {
		fmt.Printf("quadtree (debug): new tree(%d)\n", size)
	}

	center := size / 2

	return &Tree[B]{
		root: newPlane[B](center, center, size),
	}
}

func (t *Tree[B]) Put(blk B) {
	if debug {
		fmt.Printf("quadtree (debug): put block(%d, %d, %d)\n", blk.X(), blk.Y(), blk.Size())
	}

	t.root.put(blk)
}

func (t *Tree[B]) GetBlock(x, y int) (blk B, ok bool) {
	plane := t.root

	for plane.quadrants != nil {
		nextQuadrant := plane.calculateQuadrantByCoords(x, y)

		if debug {
			fmt.Printf("quadtree (debug): get block: plane(%d, %d, %d)\n",
				plane.xCenter, plane.yCenter, plane.size,
			)
			fmt.Printf("quadtree (debug): get block: next quadrant: %d\n", nextQuadrant)
		}

		if nextPlane := plane.quadrants[nextQuadrant]; nextPlane != nil {
			plane = nextPlane
		} else {
			return blk, false
		}
	}

	return *plane.blk, true
}
