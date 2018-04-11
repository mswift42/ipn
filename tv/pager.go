package tv

type NextPager interface {
	mainDoc() *IplayerDocument
	collectNextPages() []*Searcher
	collectProgramPages() []*Searcher
}
