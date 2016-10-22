package films

import "github.com/mswift42/ipn/tv"

func Programmes() ([]*tv.Programme, error) {
	filmurl := "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	programmes, err := tv.Programmes(filmurl)
	if err != nil {
		return nil, err
	}
	return programmes, nil
}
