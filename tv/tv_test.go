package tv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mostpopular = "mostpopular.html"
const filmspage1 = "filmspage1.html"
const filmspage2 = "filmspage2.html"
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
	assert.Equal(len(programmes), 40)
	assert.Equal(programmes[0].URL, "www.bbc.co.uk/iplayer/episode/b086yqrc/eastenders-24122016")
	assert.Equal(programmes[0].Pid, "b086yqrc")
	assert.Equal(programmes[39].Title, "Mr Stink")
	assert.Equal(programmes[39].Synopsis, "An unhappy, daydreaming schoolgirl befriends a homeless man and his dog in the local park.")
	assert.Equal(programmes[39].Thumbnail, "http://ichef.bbci.co.uk/images/ic/336x189/p049dz62.jpg")
	assert.Equal(programmes[1].Subtitle, "23/12/2016")
	assert.Equal(programmes[1].Title, "EastEnders")
}

func TestFindTitle(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL(mostpopular)
	programmes, err := Programmes(th)
	if err != nil {
		panic(err)
	}
	th2 := TestHtmlURL(filmspage1)
	filmsprog1, err := Programmes(th2)
	if err != nil {
		panic(err)
	}
	assert.Equal(programmes[0].Title, "EastEnders")
	assert.Equal(filmsprog1[1].Title, "Alan Partridge: Alpha Papa")
}

func TestFindSubtitle(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	popprogrammes, _ := Programmes(popth)
	assert.Equal(popprogrammes[0].Subtitle, "Series 14: Week 3")
	filmth := TestHtmlURL(filmspage1)
	filmprogrammes, _ := Programmes(filmth)
	assert.Equal(filmprogrammes[0].Subtitle, "HyperNormalisation")
	crimeth := TestHtmlURL(crime)
	crimeprogrammes, _ := Programmes(crimeth)
	assert.Equal(crimeprogrammes[0].Subtitle, "Series 6: 4. The Last Day")
}

func TestFindThumbnail(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	popprogrammes, err := Programmes(popth)
	if err != nil {
		panic(err)
	}
	assert.Equal(popprogrammes[0].Thumbnail, "https://ichef.bbci.co.uk/images/ic/336x189/p04l711h.jpg")
}
