package tv

import (
	"fmt"

)

const (
	mostpopular = "http://www.bbc.co.uk/iplayer/group/most-popular"
	films       = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
	crimedrama  = "http://www.bbc.co.uk/iplayer/categories/drama-crime/all?sort=atoz"
	comedy      = "http://www.bbc.co.uk/iplayer/categories/comedy/all?sort=atoz"
	food        = "http://www.bbc.co.uk/iplayer/categories/food/all?sort=atoz"
)

func category(url BeebURL, name string) *tv.Category {
	doc, err := url.loadDocument()
	if err != nil {
		panic(err)
	}
	nmd := newMainCategoryDocument(doc)
	progs, urls  := nmd.programmes()
	fmt.Println(urls)
	fmt.Println(progs)
	cat := NewCategory(name, progs)
	return cat
}

func LoadAllCategories() ([]*Category, error) {
	categories := map[string]BeebURL{"mostpopular": mostpopular,
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
	cats := make([]*Category, len(categories))
	ch := make(chan *Category)
		for name, url := range categories {
			go func(name string, url BeebURL) {
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
