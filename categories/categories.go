package categories

import "github.com/mswift42/ipn/tv"

const (
	mostpopular = "http://www.bbc.co.uk/iplayer/group/most-popular"
	films       = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crimedrama  = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	comedy      = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
)

func category(url string, c chan []*tv.Programme) {
	beeburl := tv.Beeburl(url)
	c <- tv.Programmes(beeburl)
}

func AllCategories() ([]*tv.Programme, error) {
	categories := []string{
		mostpopular, films, crimedrama, comedy,
	}
	var beeburl tv.BeebURL
	programmes := make([]*tv.Programme, len(categories))
	for _, i := range categories {
		beeburl = tv.BeebURL(i)
		programme, err := tv.Programmes(beeburl)
		if err != nil {
			return nil, err
		}
		programmes = append(programmes, programme...)

	}
	return programmes, nil

}
