package tv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mostpopular = "mostpopular.html"
const filmspage1 = "filmspage1.html"
const filmspage2 = "filmspage2.html"
const crm = "../drama-crime/crime.html"
const crime = "drama_and_crime.html"
const comedy = "comedy.html"

func TestBeebURLUrlDoc(t *testing.T) {
	assert := assert.New(t)
	b := BeebURL("http://www.example.com/")
	c := make(chan *IplayerDocumentResult)
	go b.loadDocument(c)
	idr := <-c
	assert.Nil(idr.Error)
	assert.NotNil(idr.idoc)
	b1 := BeebURL("")
	go b1.loadDocument(c)
	idr2 := <-c
	assert.NotNil(idr2.Error)
	assert.Nil(idr2.idoc)
}

func TestTestHtmLLoadDocument(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{mostpopular}
	c := make(chan *IplayerDocumentResult)
	go th.loadDocument(c)
	idr := <-c
	assert.Nil(idr.Error)
	assert.NotNil(idr.idoc)
	th2 := TestHtmlURL{""}
	go th2.loadDocument(c)
	idr2 := <-c
	assert.Nil(idr2.idoc)
	assert.NotNil(idr2.Error)

}

func TestCollectDocuments(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{crime}
	mcd, err := newMainCategoryDocument(th)
	assert.Nil(err)
	tmcd := TestMainCategoryDocument{mcd.ip, mcd.NextPages}
	results := tmcd.collectDocuments()
	assert.NotNil(results)
	assert.Equal(len(tmcd.NextPages), 2)
	for _, i := range results {
		assert.Nil(i.Error)
		assert.NotNil(i.idoc)
	}
	assert.Equal(len(results), 2)
}

func TestNewMainCategoryDocument(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{mostpopular}
	mcd, err := newMainCategoryDocument(th)
	assert.Nil(err)
	assert.Equal(len(mcd.NextPages), 0)
	assert.NotNil(mcd)
}

func TestSelectionResults(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{mostpopular}
	ic := make(chan *IplayerDocumentResult)
	go th.loadDocument(ic)
	doc := <-ic
	assert.Nil(doc.Error)
	assert.NotNil(doc.idoc)
	reschan := make(chan []*iplayerSelectionResult)
	idoc := &IplayerDocument{doc.idoc}
	go idoc.selectionResults(reschan)
	popres := <-reschan
	assert.Equal(len(popres), 40)
	for _, i := range popres {
		if i.progpage == "" {
			assert.NotNil(i.prog)
		}
	}
	assert.Equal(popres[0].prog.Title, "Strictly Come Dancing")
	assert.Equal(popres[0].progpage, "")
	th2 := TestHtmlURL{crime}
	go th2.loadDocument(ic)
	crimedoc := <-ic
	assert.Nil(doc.Error)
	assert.NotNil(doc.idoc)
	crimeidoc := &IplayerDocument{crimedoc.idoc}
	go crimeidoc.selectionResults(reschan)
	crimeres := <-reschan
	assert.Equal(crimeres[0].progpage, "")
	assert.Equal(crimeres[0].prog.Title, "A Nightingale Sang in Berkeley Square")
	for _, i := range crimeres {
		if i.progpage == "" {
			assert.NotNil(i.prog)
		}
		if i.prog == Nil {
			assert.NotEqual(i.progpage, "")
		}
	}

}

func TestNewProgramme(t *testing.T) {
	programme := newProgramme("title1", "subtitle1", "synopsys1",
		"a00", "http://thumbnail.url", "http://programme.url")
	assert := assert.New(t)
	assert.Equal(programme.Title, "title1")
	assert.Equal(programme.URL, "http://programme.url")

}

// func TestProgrammes(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{mostpopular}
// 	programmes, err := Programmes(th)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(len(programmes), 40)
// 	assert.Equal(programmes[0].URL, "www.bbc.co.uk/iplayer/episode/b0957wrf/strictly-come-dancing-series-15-1-launch")
// 	assert.Equal(programmes[0].Pid, "b0957wrf")
// 	assert.Equal(programmes[39].Title, "Hey Duggee")
// 	assert.Equal(programmes[39].Synopsis, "The Squirrels dress up as brave knights in cardboard costumes.")
// 	assert.Equal(programmes[39].Thumbnail, "https://ichef.bbci.co.uk/images/ic/336x189/p02typd6.jpg")
// 	assert.Equal(programmes[1].Subtitle, "08/09/2017")
// 	assert.Equal(programmes[1].Title, "EastEnders")
// }

// func TestFindTitle(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{mostpopular}
// 	programmes, err := Programmes(th)
// 	if err != nil {
// 		panic(err)
// 	}
// 	th2 := TestHtmlURL{filmspage1}
// 	filmsprog1, err := Programmes(th2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(programmes[0].Title, "Strictly Come Dancing")
// 	assert.Equal(filmsprog1[1].Title, "Broken")
// }

// func TestFindSubtitle(t *testing.T) {
// 	assert := assert.New(t)
// 	popth := TestHtmlURL{mostpopular}
// 	popprogrammes, err := Programmes(popth)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(popprogrammes[0].Subtitle, "Series 15: 1. Launch")
// 	film1th := TestHtmlURL{filmspage1}
// 	film1prog, _ := Programmes(film1th)
// 	assert.Equal(film1prog[0].Subtitle, "HyperNormalisation")
// 	assert.Equal(film1prog[1].Subtitle, "")
// }

// func TestFindThumbnail(t *testing.T) {
// 	assert := assert.New(t)
// 	popth := TestHtmlURL{mostpopular}
// 	popprogrammes, err := Programmes(popth)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(popprogrammes[0].Thumbnail, "https://ichef.bbci.co.uk/images/ic/336x189/p05fb1zb.jpg")
// }

// func TestNewCategory(t *testing.T) {
// 	assert := assert.New(t)
// 	popth := TestHtmlURL{mostpopular}
// 	popprogrammes, err := Programmes(popth)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cat := NewCategory("mostpopular", popprogrammes)
// 	assert.Equal(cat.Name, "mostpopular")
// 	assert.Equal(cat.Programmes[0].Title, "Strictly Come Dancing")
// }

// func TestProgrammeString(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{mostpopular}
// 	programmes, err := Programmes(th)
// 	if err != nil {
// 		panic(err)
// 	}
// 	p0 := programmes[0]
// 	assert.Equal(p0.String(), "0:  Strictly Come Dancing  Series 15: 1. Launch")
// }

// func TestTVSelection(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{filmspage1}
// 	doc, err := th.loadDocument()
// 	assert.Nil(err)
// 	assert.NotNil(doc)
// 	tvsel := doc.selection(".page > a")
// 	assert.Equal(tvsel.AttrOr("href", ""),
// 		"/iplayer/categories/films/all?sort=atoz&page=2")
// }

// func TestNextPages(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{filmspage1}
// 	doc, _ := th.loadDocument()
// 	np := doc.nextPages()
// 	assert.Equal(len(np), 1)
// 	assert.Equal(np[0], "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz&page=2")
// 	th = TestHtmlURL{comedy}
// 	doc, err := th.loadDocument()
// 	assert.Nil(err)
// 	np = doc.nextPages()
// 	assert.Equal(len(np), 4)
// 	assert.Equal(np[0],
// 		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=2")
// 	assert.Equal(np[1],
// 		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=3")
// 	assert.Equal(np[2],
// 		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=4")
// 	assert.Equal(np[3],
// 		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=5")
// }

// func TestMainCategoryDocument(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{comedy}
// 	doc, err := th.loadDocument()
// 	assert.Nil(err)
// 	mcd := NewMainCategoryDocument(doc)
// 	assert.Equal(len(mcd.NextPages), 4)
// 	assert.Equal(mcd.NextPages[0],
// 		bbcprefix+"/iplayer/categories/comedy/all?sort=atoz&page=2")
// }

// func TestHasExtraProgrammes(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{filmspage1}
// 	doc, err := th.loadDocument()
// 	assert.Nil(err)
// 	programmeselection := iplayerSelection{doc.idoc.Find(".list-item.programme")}
// 	assert.Equal(len(programmeselection.sel.Nodes), 20)

// 	//npl := doc.newProgrammesListItem()
// 	//assert.Equal(len(npl.sel.Nodes), 20)
// 	//first := programmesListItem{npl.sel.First()}
// 	//assert.Equal(first.hasExtraProgrammes(), true)
// 	//last := programmesListItem{npl.sel.Last()}
// 	//assert.Equal(last.hasExtraProgrammes(), false)
// 	//second := programmesListItem{npl.sel.Eq(1)}
// 	//assert.Equal(second.hasExtraProgrammes(), false)

// }

// func TestMainCategoryDocument_Programmes(t *testing.T) {
// 	assert := assert.New(t)
// 	th := TestHtmlURL{filmspage1}
// 	doc, err := th.loadDocument()
// 	assert.Nil(err)
// 	nmcd := NewMainCategoryDocument(doc)
// 	assert.NotNil(nmcd)
// 	progs, urls := nmcd.Programmes()
// 	assert.NotNil(progs)
// 	assert.NotNil(urls)
// 	//progs, urls := nmcd.Programmes()
// 	//assert.Equal(progs[0].Title, "")
// 	//if len(urls) > 0 {
// 	//	assert.Equal(urls[0], "")
// 	//}
// }

// //func TestSubPages(t *testing.T) {
// //	assert := assert.New(t)
// //	th := TestHtmlURL(filmspage1)
// //	doc, err := th.UrlDoc()
// //	assert.Nil(err)
// //	doc.CollectSubPages()
// //	sp := doc.SubPages
// //	assert.Equal(len(sp), 1)
// //	assert.Equal(string(sp[0]),
// //		bbcprefix+"/iplayer/episodes/p04bkttz")
// //	th = TestHtmlURL(comedy)
// //	doc, _ = th.UrlDoc()
// //	doc.CollectSubPages()
// //	sp = doc.SubPages
// //	assert.Equal(len(doc.SubPages), 10)
// //	assert.Equal(string(sp[0]),
// //		bbcprefix+"/iplayer/episodes/b07zyh6k")
// //	assert.Equal(string(sp[1]),
// //		bbcprefix+"/iplayer/episodes/p01djw5m")
// //	assert.Equal(string(sp[2]),
// //		bbcprefix+"/iplayer/episodes/b00hqlc4")
// //	assert.Equal(string(sp[3]),
// //		bbcprefix+"/iplayer/episodes/b006p76t")
// //}
