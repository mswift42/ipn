package categories

import (
	"fmt"

	"github.com/mswift42/goquery"
	"github.com/mswift42/ipn/tv"
)

const (
	mostpopular = "http://www.bbc.co.uk/iplayer/group/most-popular"
	films       = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crimedrama  = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	comedy      = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
	food        = "http://www.bbc.co.uk/iplayer/categories/food/all?sort=atoz"
)

func category(url, name string, c chan *tv.Category) {
	beeburl := tv.BeebURL(url)
	prog, err := tv.Programmes(beeburl)
	if err != nil {
		panic(err)
	}
	cat := tv.NewCategory(name, prog)
	c <- cat
}

func (s *goquery.Selection, c chan bool) hasNextPage() {
	c <- s.Find(".page").AttrOr("href", "") != ""
}

func AllCategories() ([]*tv.Category, error) {
	categories := map[string]tv.BeebURL{
		"most-popular": mostpopular,
		"films":        films,
		"crime/drama":  crimedrama,
		"comedy":       comedy,
		"food":         food,
	}
	cats := make([]*tv.Category, len(categories))
	ch := make(chan *tv.Category)
	for name, url := range categories {
		go func(name string, url tv.BeebURL) {
			fmt.Println("Fetching Cat: ", name)
			beeburl := tv.BeebURL(url)
			prog, err := tv.Programmes(beeburl)
			if err != nil {
				panic(err)
			}
			fmt.Println("Programmes: ", prog[0].Title)
			ch <- tv.NewCategory(name, prog)
		}(name, url)

	}

	for i := 0; i < len(categories); i++ {
		cat := <-ch
		cats = append(cats, cat)
	}
	return cats, nil

}
