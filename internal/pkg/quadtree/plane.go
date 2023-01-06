package quadtree

import "fmt"

const quadrantCount = 4

type plane[B block] struct {
	xCenter int
	yCenter int
	size    int

	quadrants *[quadrantCount]*plane[B]
	blk       *B
}

func newPlane[B block](xCenter, yCenter, size int) *plane[B] {
	if debug {
		fmt.Printf("quadtree (debug): new plane(%d, %d, %d)\n", xCenter, yCenter, size)
	}

	return &plane[B]{
		xCenter: xCenter,
		yCenter: yCenter,
		size:    size,
	}
}

func (p *plane[B]) put(blk B) {
	if p.size <= blk.Size() {
		p.blk = &blk

		return
	}

	if p.quadrants == nil {
		p.quadrants = new([quadrantCount]*plane[B])
	}

	quadrant := p.calculateQuadrantByCoords(blk.X(), blk.Y())
	if nextPlate := p.quadrants[quadrant]; nextPlate != nil {
		nextPlate.put(blk)

		return
	}

	plateXCenter, plateYCenter := p.calculatePlateCenterFor(quadrant)

	nextPlate := newPlane[B](plateXCenter, plateYCenter, p.size/2)
	p.quadrants[quadrant] = nextPlate

	nextPlate.put(blk)
}

const (
	quadrant1 = 0
	quadrant2 = 1
	quadrant3 = 2
	quadrant4 = 3
)

func (p *plane[B]) calculateQuadrantByCoords(x, y int) int {
	if x >= p.xCenter {
		if y < p.yCenter {
			return quadrant1
		}

		return quadrant4
	} else {
		if y < p.yCenter {
			return quadrant2
		}

		return quadrant3
	}
}

func (p *plane[B]) calculatePlateCenterFor(quadrant int) (plateXCenter, plateYCenter int) {
	step := p.size / quadrantCount

	switch quadrant {
	case quadrant1:
		return p.xCenter + step, p.yCenter - step
	case quadrant2:
		return p.xCenter - step, p.yCenter - step
	case quadrant3:
		return p.xCenter - step, p.yCenter + step
	case quadrant4:
		return p.xCenter + step, p.yCenter + step
	}

	panic(
		fmt.Sprintf("quadtree: failed to calculate plate center for unknown quadrant %d", quadrant),
	)
}
