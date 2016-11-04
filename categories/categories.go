package categories

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

const (
	mostpop = "http://www.bbc.co.uk/iplayer/group/most-popular"
	films   = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crdrama = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	comedy  = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
)

func MostPopular(doc *goquery.Document) ([]*tv.Programme, error) {
	doc, err := goquery.NewDocument(mostpop)
	if err != nil {
		return nil, err
	}
	return tv.Programmes(doc), nil
}

func Films(doc *goquery.Document) ([]*tv.Programme, error) {
	doc, err := goquery.NewDocument(films)
	if err != nil {
		return nil, err
	}
	return tv.Programmes(doc), nil
}
