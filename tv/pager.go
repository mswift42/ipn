package tv

type Pager interface {
	collectDocuments() []*IplayerDocumentResult
}
