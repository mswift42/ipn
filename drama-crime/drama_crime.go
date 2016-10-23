package dramacrime

import "github.com/mswift42/ipn/tv"

func Programmes() ([]*tv.Programme, error) {
	crimeurl := "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	programmes, err := tv.Programmes(crimeurl)
	if err != nil {
		return nil, err
	}
	return programmes, err
}
