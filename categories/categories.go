package categories

import "github.com/mswift42/ipn/tv"

const (
	mostpopular = "http://www.bbc.co.uk/iplayer/group/most-popular"
	films       = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crimedrama  = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	comedy      = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
)

type Category struct {
	name string
	url  tv.BeebURL
}

func newCategory(name, url tv.BeebURL) *Category {
	return &Category{name, url}
}

func AllCategories() (*[]tv.Programme, error) {
	categories := []*Category{
		{"mostpoular", mostpopular},
		{"films", films},
		{"crimedrama", crimedrama},
		{"comedy", comedy},
	}
	var beeburl tv.Beeburl
	programmes := make([]*Programme, len(categories))
	for _, i := range categories {
		beeburl = tv.BeebURL(i.url)
		programme, err := tv.Programmes(beeburl)
		if err != nil {
			return nil, err
		}
		programmes = append(programmes, programme)

	}
	return programmes, nil

}
