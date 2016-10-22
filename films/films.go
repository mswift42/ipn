package films

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

func Programmes() ([]*tv.Programme, error) {
	popurl := "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	doc, err := goquery.NewDocument(popurl)
	if err != nil {
		return nil, err
	}
	var programmes []*tv.Programme

}
