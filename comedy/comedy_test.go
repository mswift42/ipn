package comedy

import (
	"testing"

	"github.com/mswift42/ipn/testutils"
	"github.com/stretchr/testify/assert"
)

func TestProgrammes(t *testing.T) {
	assert := assert.New(t)
	cmddoc := testutils.LoadTestHtml("comedy.html")
	programmes := Programmes(cmddoc)
	assert.Equal(programmes[0].Title, "Angry Boys")
	assert.Equal(programmes[0].Subtitle, "Episode 12")
	assert.Equal(programmes[3].Title, "The Blame Game")
	assert.Equal(programmes[3].Synopsis, "With Tim McGarry, Colin Murphy, Jake O'Kane, Neil Delamere and Kevin Bridges.")
}
