package mostpopular

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

const mostpopular = "iplayermostpopular.html"

func loadTestHtml(filename string) *goquery.Document {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}
	return doc
}

func TestProgrammes(t *testing.T) {
	assert := assert.New(t)
	doc := loadTestHtml(mostpopular)
	prog := Programmes(doc)
	assert.Equal(len(prog), 40)
	assert.Equal(prog[0].Title, "Strictly Come Dancing")
	assert.Equal(prog[0].Synopsis, "The couples must deliver an amazing routine to a classic movie track.")
	assert.Equal(prog[0].Subtitle, "Series 14: Week 3")
	assert.Equal(prog[10].Title, "Our Girl")
	assert.Equal(prog[10].Subtitle, "Series 2: Episode 5")
}
