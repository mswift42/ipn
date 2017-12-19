package tv

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
	//prog, err := tv.Programmes(url)
	//if err != nil {
	//	panic(err)
	//}
	doc, err := url.loadDocument()
	if err != nil {
		panic(err)
	}
	nmd := tv.newMainCategoryDocument(doc)
	progs, urls  := nmd.programmes()
	fmt.Println(urls)
	fmt.Println(progs)
	cat := tv.NewCategory(name, progs)
	return cat
}

func LoadAllCategories() ([]*tv.Category, error) {
	categories := map[string]tv.BeebURL{"mostpopular": mostpopular,
		"films":      films,
		"crimedrama": crimedrama,
		"comedy":     comedy}
	cats, err := allCategories(categories)
	if err != nil {
		panic(err)
	}
	return cats, nil
}

func allCategories(categories map[string]tv.BeebURL) ([]*tv.Category, error) {
	cats := make([]*tv.Category, len(categories))
	ch := make(chan *tv.Category)
		for name, url := range categories {
			go func(name string, url tv.BeebURL) {
				fmt.Println(name, url)
				//doc, err := url.UrlDoc()
				//if err != nil {
				//	panic(err)
				//}
				//doc.CollectNextPage()
				//if len(doc.NextPages) > 0 {
				//	for _, i := range doc.NextPages {
				//		ch <- category(tv.BeebURL(i), name)
				//	}
				//}
				fmt.Println("Fetching Cat: ", name)
				ch <- category(url, name)

			}(name, url)
		}

		defer close(ch)

	for c := range ch {
		cats = append(cats, c)
	}
	return cats, nil
}
