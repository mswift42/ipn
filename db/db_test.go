package db

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProgrammeDB(t *testing.T) {
	assert := assert.New(t)
	popth := tv.TestHtmlURL(mostpopular)
	popprogrammes, _ := Programmes(popth)
	cat := NewCategory("mostpopular", popprogrammes)
	filmdoc := TestHtmlURL(filmspage1)
	filmprog, _ := Programmes(filmdoc)
	cat2 := NewCategory("films", filmprog)
	cats := []*Category{cat, cat2}
	now := time.Now()
	pdb := db.newProgrammeDB(cats, now)
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
	pdb := tv.newProgrammeDB([]*Category{cat1, cat2}, time.Now())
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
