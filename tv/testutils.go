package tv

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/mswift42/goquery"
)

type TestHtmlURL struct {
	url string
}

type TestMainCategoryDocument struct {
	ip        *IplayerDocument
	NextPages []string
}

// func (t TestMainCategoryDocument) collectNextPages() []*IplayerDocumentResult {
// 	var results []*IplayerDocumentResult
// 	sc := make(chan Searcher)
// 	idrc := make(chan *IplayerDocumentResult)
// 	return results

// }

func (t *TestMainCategoryDocument) collectNextPages() []*IplayerDocumentResult {
	var results []*IplayerDocumentResult
	sc := make(chan Searcher)
	idrc := make(chan *IplayerDocumentResult)
	go collectDocument(sc, idrc)
	for _, i := range t.NextPages {
		go func(url string) {
			th := TestHtmlURL{strings.Replace(url, bbcprefix, "", -1)}
			sc <- th
		}(i)
	}
	for range t.NextPages {
		results = append(results, <-idrc)
	}
	return results
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

}
