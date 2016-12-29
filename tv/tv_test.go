package tv

import (
	"strings"
	"testing"

	"io/ioutil"
	"time"

	"github.com/stretchr/testify/assert"
)

const mostpopular = "mostpopular.html"
const filmspage1 = "filmspage1.html"
const filmspage2 = "filmspage2.html"
const crime = "../drama-crime/crime.html"

func TestBeebURLUrlDoc(t *testing.T) {
	assert := assert.New(t)
	b := BeebURL("http://www.example.com/")
	ex, err := b.urlDoc()
	assert.Nil(err)
	assert.NotNil(ex)
	b1 := BeebURL("")
	ex1, err := b1.urlDoc()
	assert.NotNil(err)
	assert.Nil(ex1)
}

func TestTestHtmlURLDoc(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL(mostpopular)
	succ, err := th.urlDoc()
	assert.Nil(err)
	assert.NotNil(succ)
	th2 := TestHtmlURL("")
	fail, err := th2.urlDoc()
	assert.Nil(fail)
	assert.NotNil(err)

}

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
	assert.Equal(programmes[39].Thumbnail, "https://ichef.bbci.co.uk/images/ic/336x189/p01j0d94.jpg")
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
	assert.Equal(filmsprog1[1].Title, "Adam Curtis")
}

func TestFindSubtitle(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	popprogrammes, err := Programmes(popth)
	if err != nil {
		panic(err)
	}
	assert.Equal(popprogrammes[0].Subtitle, "24/12/2016")
	film1th := TestHtmlURL(filmspage1)
	film1prog, _ := Programmes(film1th)
	assert.Equal(film1prog[0].Subtitle, "HyperNormalisation")
	assert.Equal(film1prog[1].Subtitle, "Bitter Lake")
	assert.Equal(film1prog[2].Subtitle, "")
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

func TestNewCategory(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	popprogrammes, err := Programmes(popth)
	if err != nil {
		panic(err)
	}
	cat := NewCategory("mostpopular", popprogrammes)
	assert.Equal(cat.Name, "mostpopular")
	assert.Equal(cat.Programmes[0].Title, "EastEnders")
}

func TestNewProgrammeDB(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	popprogrammes, _ := Programmes(popth)
	cat := NewCategory("mostpopular", popprogrammes)
	filmdoc := TestHtmlURL(filmspage1)
	filmprog, _ := Programmes(filmdoc)
	cat2 := NewCategory("films", filmprog)
	cats := []*Category{cat, cat2}
	now := time.Now()
	pdb := newProgrammeDB(cats, now)
	assert.Equal(len(pdb.Categories), 2)
	assert.Equal(pdb.Saved, now)
	assert.Equal(now.Day(), time.Now().Day())
}

func TestProgrammeDB_Save(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	film1 := TestHtmlURL(filmspage1)
	popprog, _ := Programmes(popth)
	filmprog, _ := Programmes(film1)
	cat1 := NewCategory("mostpopular", popprog)
	cat2 := NewCategory("films", filmprog)
	pdb := newProgrammeDB([]*Category{cat1, cat2}, time.Now())
	json, err := pdb.toJSON()
	assert.Nil(err)
	assert.NotNil(json)
	pdb.Save("testjson.json")
	file, err := ioutil.ReadFile("testjson.json")
	assert.NotNil(file)
	assert.Nil(err)
	assert.True(strings.Contains(string(file), "categories"))
	assert.True(strings.Contains(string(file), "saved"))
	assert.True(strings.Contains(string(file), "synopsis"))

}

func TestProgrammeDB_Index(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL(mostpopular)
	film1 := TestHtmlURL(filmspage1)
	popprog, _ := Programmes(popth)
	filmprog, _ := Programmes(film1)
	cat1 := NewCategory("mostpopular", popprog)
	cat2 := NewCategory("films", filmprog)
	pdb := newProgrammeDB([]*Category{cat1, cat2}, time.Now())
	pdb.Save("testjson.json")
	file, _ := ioutil.ReadFile("testjson.json")
	assert.True(strings.Contains(string(file), "39"))
}

func TestNewProgrammeDBFromJSON(t *testing.T) {
	assert := assert.New(t)
	db, err := LoadProgrammeDbFromJSON("testjson.json")
	assert.Nil(err)
	assert.Equal(len(db.Categories), 2)
	assert.Equal(db.Categories[0].Name, "mostpopular")
	assert.Equal(db.Categories[1].Name, "films")
}

func TestProgrammeString(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL(mostpopular)
	programmes, err := Programmes(th)
	if err != nil {
		panic(err)
	}
	p0 := programmes[0]
	assert.Equal(p0.String(), "0:  EastEnders  24/12/2016")
}
