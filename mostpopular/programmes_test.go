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
	doc := loadTestHtml(mostpopular)
	prog := Programmes(doc)
	assert.Equal(t, len(prog), 40)
}
