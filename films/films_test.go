package films

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgrammes(t *testing.T) {
	programmes, err := Programmes()
	assert.Nil(t, err)
	assert.NotEmpty(t, programmes)
}
