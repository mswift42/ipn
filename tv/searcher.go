package tv

type Searcher interface {
	loadDocument(chan<- *IplayerDocumentResult)
}

