package categories

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

const (
	mostpop = "http://www.bbc.co.uk/iplayer/group/most-popular"
	flms    = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crdrama = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	cmdy    = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
)

type Category struct {
	name, url string
}

func newCategory(name, url string) *Category {
	return &Category{name, url}
}

func mostPopular(doc *goquery.Document) []*tv.Programme {
	return tv.Programmes(doc)
}

func films(doc *goquery.Document) []*tv.Programme {
	return tv.Programmes(doc)
}

func cdrdama(doc *goquery.Document) []*tv.Programme {
	return tv.Programmes(doc)
}

func comedy(doc *goquery.Document) []*tv.Programme {
	return tv.Programmes(doc)
}
