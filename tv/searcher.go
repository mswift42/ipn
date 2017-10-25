package tv

type Searcher interface {
	loadDocument(chan<- *IplayerDocumentResult)
}

type Pager interface {
	collectDocuments() []*IplayerDocumentResult
}
