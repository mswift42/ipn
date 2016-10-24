package films

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

func Programmes() ([]*tv.Programme, error) {
	filmurl := "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	doc, err := goquery.NewDocument(filmurl)
	if err != nil {
		return nil, err
	}
	programmes := tv.Programmes(doc)
	return programmes, nil
}
