package tv

import (
	"bytes"
	"io/ioutil"

	"github.com/mswift42/goquery"
)

type TestHtmlURL struct {
	url string
}

func (th TestHtmlURL) loadDocument(c chan<- *IplayerDocumentResult) {
	file, err := ioutil.ReadFile(th.url)
	if err != nil {
		c <- &IplayerDocumentResult{nil, err}
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		c <- &IplayerDocumentResult{nil, err}
	}
	c <- &IplayerDocumentResult{doc, nil}
	close(c)
}
