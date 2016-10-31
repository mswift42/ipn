package dramacrime

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

func Programmes(doc *goquery.Document) ([]*tv.Programme, error) {
	crimeurl := "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	doc, err := goquery.NewDocument(crimeurl)
	if err != nil {
		return nil, err
	}
	programmes := tv.Programmes(doc)
	return programmes, err
}
