package tv

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

const mostpopular = "mostpopular.html"
const filmspage1 = "filmspage1.html"
const filmspage2 = "filmspage2.html"
const crime = "../drama-crime/crime.html"
const comedy = "comedy.html"

func TestBeebURLUrlDoc(t *testing.T) {
	assert := assert.New(t)
	b := BeebURL("http://www.example.com/")
	ex, err := b.loadDocument()
	assert.Nil(err)
	assert.NotNil(ex)
	b1 := BeebURL("")
	ex1, err := b1.loadDocument()
	assert.NotNil(err)
	assert.Nil(ex1)
}

func TestTestHtmLLoadDocument(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{mostpopular}
	succ, err := th.loadDocument()
	assert.Nil(err)
	assert.NotNil(succ)
	th2 := TestHtmlURL{""}
	fail, err := th2.loadDocument()
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
	th := TestHtmlURL{mostpopular}
	programmes, err := Programmes(th)
	if err != nil {
		panic(err)
	}
	assert.Equal(len(programmes), 40)
	assert.Equal(programmes[0].URL, "www.bbc.co.uk/iplayer/episode/b0957wrf/strictly-come-dancing-series-15-1-launch")
	assert.Equal(programmes[0].Pid, "b0957wrf")
	assert.Equal(programmes[39].Title, "Hey Duggee")
	assert.Equal(programmes[39].Synopsis, "The Squirrels dress up as brave knights in cardboard costumes.")
	assert.Equal(programmes[39].Thumbnail, "https://ichef.bbci.co.uk/images/ic/336x189/p02typd6.jpg")
	assert.Equal(programmes[1].Subtitle, "08/09/2017")
	assert.Equal(programmes[1].Title, "EastEnders")
}

func TestFindTitle(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{mostpopular}
	programmes, err := Programmes(th)
	if err != nil {
		panic(err)
	}
	th2 := TestHtmlURL{filmspage1}
	filmsprog1, err := Programmes(th2)
	fmt.Println(filmsprog1)
	if err != nil {
		panic(err)
	}
	assert.Equal(programmes[0].Title, "Strictly Come Dancing")
	assert.Equal(filmsprog1[1].Title, "Adam Curtis")
}

func TestFindSubtitle(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL{mostpopular}
	popprogrammes, err := Programmes(popth)
	if err != nil {
		panic(err)
	}
	assert.Equal(popprogrammes[0].Subtitle, "Series 15: 1. Launch")
	film1th := TestHtmlURL{filmspage1}
	film1prog, _ := Programmes(film1th)
	assert.Equal(film1prog[0].Subtitle, "HyperNormalisation")
	assert.Equal(film1prog[1].Subtitle, "Bitter Lake")
}

func TestFindThumbnail(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL{mostpopular}
	popprogrammes, err := Programmes(popth)
	if err != nil {
		panic(err)
	}
	assert.Equal(popprogrammes[0].Thumbnail, "https://ichef.bbci.co.uk/images/ic/336x189/p05fb1zb.jpg")
}

func TestNewCategory(t *testing.T) {
	assert := assert.New(t)
	popth := TestHtmlURL{mostpopular}
	popprogrammes, err := Programmes(popth)
	if err != nil {
		panic(err)
	}
	cat := NewCategory("mostpopular", popprogrammes)
	assert.Equal(cat.Name, "mostpopular")
	assert.Equal(cat.Programmes[0].Title, "Strictly Come Dancing")
}

func TestProgrammeString(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{mostpopular}
	programmes, err := Programmes(th)
	if err != nil {
		panic(err)
	}
	p0 := programmes[0]
	assert.Equal(p0.String(), "0:  Strictly Come Dancing  Series 15: 1. Launch")
}

func TestTVSelection(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{filmspage1}
	doc, err := th.loadDocument()
	assert.Nil(err)
	assert.NotNil(doc)
	tvsel := doc.selection(".page > a")
	assert.Equal(tvsel.AttrOr("href", ""),
		"/iplayer/categories/films/all?sort=atoz&page=2")
}

func TestNextPages(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{filmspage1}
	doc, _ := th.loadDocument()
	np := doc.nextPages()
	assert.Equal(len(np), 1)
	assert.Equal(np[0], "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz&page=2")
	 th = TestHtmlURL{comedy}
	doc, err := th.loadDocument()
	assert.Nil(err)
	np = doc.nextPages()
	assert.Equal(len(np), 4)
	assert.Equal(np[0],
		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=2")
	assert.Equal(np[1],
		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=3")
	assert.Equal(np[2],
		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=4")
	assert.Equal(np[3],
		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=5")
}

//func TestSubPages(t *testing.T) {
//	assert := assert.New(t)
//	th := TestHtmlURL(filmspage1)
//	doc, err := th.UrlDoc()
//	assert.Nil(err)
//	doc.CollectSubPages()
//	sp := doc.SubPages
//	assert.Equal(len(sp), 1)
//	assert.Equal(string(sp[0]),
//		bbcprefix+"/iplayer/episodes/p04bkttz")
//	th = TestHtmlURL(comedy)
//	doc, _ = th.UrlDoc()
//	doc.CollectSubPages()
//	sp = doc.SubPages
//	assert.Equal(len(doc.SubPages), 10)
//	assert.Equal(string(sp[0]),
//		bbcprefix+"/iplayer/episodes/b07zyh6k")
//	assert.Equal(string(sp[1]),
//		bbcprefix+"/iplayer/episodes/p01djw5m")
//	assert.Equal(string(sp[2]),
//		bbcprefix+"/iplayer/episodes/b00hqlc4")
//	assert.Equal(string(sp[3]),
//		bbcprefix+"/iplayer/episodes/b006p76t")
//}
