package tv

import "github.com/mswift42/goquery"

type Searcher interface {
	loadDocument() (*IplayerDocument, error)
}
