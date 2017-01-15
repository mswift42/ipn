package db

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/mswift42/ipn/tv"
	"github.com/stretchr/testify/assert"
)

const comedy = "../tv/comedy.html"
const mostpopular = "../tv/mostpopular.html"
const filmspage1 = "../tv/filmspage1.html"

func TestNewProgrammeDB(t *testing.T) {
	assert := assert.New(t)
	popth := tv.TestHtmlURL(mostpopular)
	popprogrammes, _ := tv.Programmes(popth)
	cat := tv.NewCategory("mostpopular", popprogrammes)
	filmdoc := tv.TestHtmlURL(filmspage1)
	filmprog, _ := tv.Programmes(filmdoc)
	cat2 := tv.NewCategory("films", filmprog)
	cats := []*tv.Category{cat, cat2}
	now := time.Now()
	pdb := newProgrammeDB(cats, now)
	assert.Equal(len(pdb.Categories), 2)
	assert.Equal(pdb.Saved, now)
	assert.Equal(now.Day(), time.Now().Day())
}

func TestProgrammeDB_Save(t *testing.T) {
	assert := assert.New(t)
	popth := tv.TestHtmlURL(mostpopular)
	film1 := tv.TestHtmlURL(filmspage1)
	popprog, _ := tv.Programmes(popth)
	filmprog, _ := tv.Programmes(film1)
	cat1 := tv.NewCategory("mostpopular", popprog)
	cat2 := tv.NewCategory("films", filmprog)
	pdb := newProgrammeDB([]*tv.Category{cat1, cat2}, time.Now())
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
	popth := tv.TestHtmlURL(mostpopular)
	film1 := tv.TestHtmlURL(filmspage1)
	popprog, _ := tv.Programmes(popth)
	filmprog, _ := tv.Programmes(film1)
	cat1 := tv.NewCategory("mostpopular", popprog)
	cat2 := tv.NewCategory("films", filmprog)
	pdb := newProgrammeDB([]*tv.Category{cat1, cat2}, time.Now())
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

func TestListCategory(t *testing.T) {
	assert := assert.New(t)
	db, err := LoadProgrammeDbFromJSON("testjson.json")
	assert.Nil(err)
	assert.NotNil(db)
	assert.Equal(len(db.Categories), 2)
	catstring := db.ListCategory("mostpopular")
	assert.NotNil(catstring)
	assert.Equal(len(strings.Split(catstring, "\n")), 41)
}

func TestFindTitle(t *testing.T) {
	assert := assert.New(t)
	db, _ := LoadProgrammeDbFromJSON("testjson.json")
	assert.Equal(len(db.Categories), 2)
	findString := db.FindTitle("Adam")
	assert.Equal(findString, "40:  Adam Curtis  HyperNormalisation\n41:  Adam Curtis  Bitter Lake\n")
	findString = db.FindTitle("EastEnders")
	assert.Equal(findString, "0:  EastEnders  24/12/2016\n1:  EastEnders  23/12/2016\n7:  EastEnders  22/12/2016\n22:  EastEnders  20/12/2016\n")
	findString = db.FindTitle("eastenders")
	assert.Equal(findString, "0:  EastEnders  24/12/2016\n1:  EastEnders  23/12/2016\n7:  EastEnders  22/12/2016\n22:  EastEnders  20/12/2016\n")
	findString = db.FindTitle("BROWN")
	assert.Equal(findString, "12:  Mrs Brown's Boys  Christmas Specials 2012: 2. The Virgin Mammy\n23:  Father Brown  Series 5: 1. The Star of Jacob\n")
}
