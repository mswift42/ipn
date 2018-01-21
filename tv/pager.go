package tv

type Pager interface {
	collectNextPages() []*IplayerDocumentResult
	collectViewMorePages() []*IplayerDocumentResult
}
