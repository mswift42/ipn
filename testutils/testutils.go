package testutils

import (
	"bytes"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

func (th tv.TestHtmLURL) urlDoc() (*goquery.Document, error) {
	file, err := ioutil.ReadFile(string(th))
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

type TestIplayerDocument struct {
	tdoc *goquery.Document
}

func (td *TestIplayerDocument) NextPages() []string {

}

func (td *TestIplayerDocument) morePages(selection string) []string {
	var results []string
	sel := td.tdoc.Find(selection)
}

func LoadTestHtml(filename string) *goquery.Document {
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
