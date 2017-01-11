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

func category(url tv.BeebURL, name string) *tv.Category {
	prog, err := tv.Programmes(url)
	if err != nil {
		panic(err)
	}
	cat := tv.NewCategory(name, prog)
	return cat
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
	go func() {
		for name, url := range categories {
			doc, err := url.UrlDoc()
			if err != nil {
				panic(err)
			}
			nextpages := doc.NextPages()
			if len(nextpages) > 0 {
				for _, i := range nextpages {
					ch <- category(i, name)
				}
			}
			fmt.Println("Fetching Cat: ", name)
			ch <- category(url, name)

		}
		close(ch)
	}()
	for c := range ch {
		cats = append(cats, c)
	}
	return cats, nil

}
