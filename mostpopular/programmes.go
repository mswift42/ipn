package mostpopular

import "github.com/mswift42/ipn/tv"

func Programmes() ([]*tv.Programme, error) {
	popurl := "http://www.bbc.co.uk/iplayer/group/most-popular"
	programmes, err := tv.Programmes(popurl)
	if err != nil {
		return nil, err
	}
	return programmes, err
}
