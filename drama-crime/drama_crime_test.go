package dramacrime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgrammes(t *testing.T) {
	programmes, err := Programmes()
	assert.Nil(t, err)
	assert.True(t, len(programmes) > 0)
}
