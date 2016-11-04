package categories

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

const (
	mp      = "http://www.bbc.co.uk/iplayer/group/most-popular"
	flms    = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crdrama = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	cmdy    = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
)

func MostPopular(doc *goquery.Document) ([]*tv.Programme, error) {
	doc, err := goquery.NewDocument(mp)
	if err != nil {
		return nil, err
	}
	return tv.Programmes(doc), nil
}
