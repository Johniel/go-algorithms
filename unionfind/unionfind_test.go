package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	uf := New(10)

	assert.False(t, uf.IsSameSet(1, 2))
	uf.Union(1, 2)
	assert.True(t, uf.IsSameSet(1, 2))
}

func TestFind(t *testing.T) {
	uf := New(10)

	assert.Equal(t, uf.Find(1), 1)
	assert.Equal(t, uf.Find(2), 2)
}
