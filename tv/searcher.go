package tv


type Searcher interface {
	loadDocument() (*IplayerDocument, error)
}
