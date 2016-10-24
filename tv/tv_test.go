package tv

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/testutils"
	"github.com/stretchr/testify/assert"
)

const mostpopular = "../mostpopular/iplayermostpopular.html"
const films = "../films/films.html"
const crime = "../drama-crime/crime.html"

func TestNewProgramme(t *testing.T) {
	programme := newProgramme("title1", "subtitle1", "synopsys1",
		"a00", "http://thumbnail.url", "http://programme.url")
	assert := assert.New(t)
	assert.Equal(programme.Title, "title1")
	assert.Equal(programme.Index, 0)
	assert.Equal(programme.Url, "http://programme.url")

}

func TestProgrammes(t *testing.T) {
	assert := assert.New(t)
	doc := testutils.LoadTestHtml(mostpopular)
	programmes := Programmes(doc)
	assert.Equal(len(programmes), 40)
}

func TestFindTitle(t *testing.T) {
	doc := testutils.LoadTestHtml(mostpopular)
	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		assert.NotEqual(t, findTitle(s), "")
	})
	programmes := Programmes(doc)
	assert.Equal(t, programmes[0].Title, "Strictly Come Dancing")
	filmdoc := testutils.LoadTestHtml(films)
	filmprogrammes := Programmes(filmdoc)
	assert.Equal(t, filmprogrammes[0].Title, "Adam Curtis")
	crimedoc := testutils.LoadTestHtml(crime)
	crimeprogrammes := Programmes(crimedoc)
	assert.Equal(t, crimeprogrammes[0].Title, "Beck")
}
