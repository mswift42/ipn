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

func allCategories() (*[]tv.Programme, error) {

}
