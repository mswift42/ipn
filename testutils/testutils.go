package testutils

import (
	"bytes"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

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
