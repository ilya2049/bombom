package quadtree_test

type fakeBlock struct {
	x, y, size int
}

func newFakeBlock(x, y, size int) *fakeBlock {
	return &fakeBlock{
		x:    x,
		y:    y,
		size: size,
	}
}

func (blk *fakeBlock) X() int    { return blk.x }
func (blk *fakeBlock) Y() int    { return blk.y }
func (blk *fakeBlock) Size() int { return blk.size }
