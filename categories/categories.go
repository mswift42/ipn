package categories

import (
	"fmt"

	"github.com/mswift42/ipn/tv"
)

const (
	mostpopular = "http://www.bbc.co.uk/iplayer/group/most-popular"
	films       = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crimedrama  = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	comedy      = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
	food        = "http://www.bbc.co.uk/iplayer/categories/food/all?sort=atoz"
)

type Category struct {
	name       string
	programmes []*tv.Programme
}

func category(url string, c chan []*tv.Programme) {
	beeburl := tv.BeebURL(url)
	prog, err := tv.Programmes(beeburl)
	if err != nil {
		panic(err)
	}
	c <- prog
}

func AllCategories() ([]*tv.Programme, error) {
	categories := []string{
		mostpopular, films, crimedrama, comedy, food,
	}
	programmes := make([]*tv.Programme, len(categories))
	ch := make(chan []*tv.Programme)
	for _, i := range categories {
		go func(i string) {
			fmt.Println("Fetching Cat: ", i)
			beeburl := tv.BeebURL(i)
			prog, err := tv.Programmes(beeburl)
			if err != nil {
				panic(err)
			}
			fmt.Println("Programmes: ", prog[0].Title)
			ch <- prog
		}(i)

	}

	for i := 0; i < len(categories); i++ {
		prog := <-ch
		programmes = append(programmes, prog...)
	}
	return programmes, nil

}
