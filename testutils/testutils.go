package testutils

import (
	"bytes"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)
type TestHtmlURL struct {
	url string
}

func (th TestHtmlURL) loadDocument() (*tv.IplayerDocument, error) {
	file, err := ioutil.ReadFile(th.url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (th TestHtmlURL) UrlDoc() (*goquery.Document, error) {
	file, err := ioutil.ReadFile(th.url)
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
