package tv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mostpopular = "../mostpopular/iplayermostpopular.html"
const films = "../films/films.go"

func TestNewProgramme(t *testing.T) {
	programme := newProgramme("title1", "subtitle1", "synopsys1",
		"a00", "http://thumbnail.url", "http://programme.url")
	assert := assert.New(t)
	assert.Equal(programme.Title, "title1")
	assert.Equal(programme.Index, 0)
	assert.Equal(programme.Url, "http://programme.url")

}
