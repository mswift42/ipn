package tv

type Pager interface {
	collectNextPages() []*IplayerDocumentResult
}
