package testutils

import (
	"bytes"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

// func (th TestHtmLURL) urlDoc() (*goquery.Document, error) {
// 	file, err := ioutil.ReadFile(string(th))
// 	if err != nil {
// 		return nil, err
// 	}
// 	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return doc, nil
// }

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
