package tv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mostpopular = "../categories/mostpopular.html"
const films = "../films/films.html"
const crime = "../drama-crime/crime.html"

func TestNewProgramme(t *testing.T) {
	programme := newProgramme("title1", "subtitle1", "synopsys1",
		"a00", "http://thumbnail.url", "http://programme.url")
	assert := assert.New(t)
	assert.Equal(programme.Title, "title1")
	assert.Equal(programme.URL, "http://programme.url")

}

func TestProgrammes(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL(mostpopular)
	programmes, err := Programmes(th)
	if err != nil {
		panic(err)
	}
	assert.Equal(programmes[0].URL, "www.bbc.co.uk/iplayer/episode/b07zhnf6/strictly-come-dancing-series-14-week-3")
	assert.Equal(programmes[0].Pid, "b07zhnf6")
	assert.Equal(programmes[39].Title, "Cleverman")
	assert.Equal(programmes[39].Synopsis, "Koen must send Kora back to her own dimension and save the Hairypeople from eviction.")
	assert.Equal(programmes[39].Thumbnail, "http://ichef.bbci.co.uk/images/ic/336x189/p049dz62.jpg")
	assert.Equal(programmes[1].Subtitle, "Series 12: 1. Collectables")
	assert.Equal(programmes[1].Title, "The Apprentice")
}

func TestFindTitle(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL(mostpopular)
	programmes, _ := Programmes(th)
	assert.Equal(programmes[0].Title, "Strictly Come Dancing")
	thfilm := TestHtmlURL(films)
	filmprogrammes, _ := Programmes(thfilm)
	assert.Equal(filmprogrammes[0].Title, "Adam Curtis")
	crimeth := TestHtmlURL(crime)
	crimeprogrammes, _ := Programmes(crimeth)
	assert.Equal(crimeprogrammes[0].Title, "Beck")
}

func TestFindSubtitle(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	popprogrammes, _ := Programmes(popth)
	assert.Equal(popprogrammes[0].Subtitle, "Series 14: Week 3")
	filmth := TestHtmlURL(films)
	filmprogrammes, _ := Programmes(filmth)
	assert.Equal(filmprogrammes[0].Subtitle, "HyperNormalisation")
	crimeth := TestHtmlURL(crime)
	crimeprogrammes, _ := Programmes(crimeth)
	assert.Equal(crimeprogrammes[0].Subtitle, "Series 6: 4. The Last Day")
}
