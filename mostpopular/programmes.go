package mostpopular

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

func Programmes() ([]*tv.Programme, error) {
	popurl := "http://www.bbc.co.uk/iplayer/group/most-popular"
	doc, err := goquery.NewDocument(popurl)
	if err != nil {
		return nil, err
	}
	programmes := tv.Programmes(doc)
	return programmes, err
}
