package tv

import (
	"bytes"
	"io/ioutil"

	"github.com/mswift42/goquery"
)

type TestHtmlURL struct {
	url string
}

func (th TestHtmlURL) loadDocument() (*IplayerDocument,  error) {
	file, err := ioutil.ReadFile(th.url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}
	newidoc := NewIplayerDocument(doc)
	return newidoc,  nil
}


