package quadtree_test

import (
	"bombom/internal/pkg/quadtree"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTreePut(t *testing.T) {
	tree := quadtree.New[*fakeBlock](64)

	tree.Put(newFakeBlock(8, 48, 8))

	block, ok := tree.GetBlock(9, 49)

	require.True(t, ok)
	assert.Equal(t, 8, block.X())
	assert.Equal(t, 48, block.Y())
	assert.Equal(t, 8, block.Size())
}
