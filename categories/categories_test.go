package categories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	assert := assert.New(t)
	nc1 := newCategory("mostpopular", "mostpop")
	assert.Equal(nc1.name, "mostpopular")
	assert.Equal(nc1.url, "http://www.bbc.co.uk/iplayer/group/most-popular")
}
