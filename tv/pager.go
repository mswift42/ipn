package tv

type NextPager interface {
	collectNextPages() []*IplayerDocumentResult
	collectViewMorePages() []*IplayerDocumentResult
}
