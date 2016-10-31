package films

import (
	"testing"

	"github.com/mswift42/ipn/testutils"
	"github.com/stretchr/testify/assert"
)

const films = "films.html"

func TestProgrammes(t *testing.T) {
	assert := assert.New(t)
	filmdoc := testutils.LoadTestHtml(films)
	programmes := Programmes(filmdoc)
	assert.NotEmpty(programmes)
	assert.Equal(programmes[0].Title, "Adam Curtis")
	assert.Equal(programmes[0].Subtitle, "HyperNormalisation")
	assert.Equal(programmes[2].Title, "Begin Again")
	assert.Equal(programmes[2].Synopsis, "A former music company producer and a talented singer-songwriter start to collaborate.")
}
