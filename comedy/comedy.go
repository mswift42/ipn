package comedy

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/tv"
)

func Programmes(doc *goquery.Document) []*tv.Programme {
	programmes := tv.Programmes(doc)
	return programmes
}
