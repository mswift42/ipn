package tv

import (

	"github.com/mswift42/goquery"
	"fmt"
)

const bbcprefix = "http://www.bbc.co.uk"

//type Searcher interface {
//	UrlDoc() (*IplayerDocument, error)
//}

// BeebURL represents an Iplayer URL.
type BeebURL string


type IplayerDocument struct {
	idoc      *goquery.Document
}

func NewIplayerDocument(doc *goquery.Document) *IplayerDocument {
	return &IplayerDocument{doc}
}


type ProgrammeSelection struct {
	sel *goquery.Selection
}

func (b BeebURL) loadDocument() (*IplayerDocument, error) {
	doc, err := goquery.NewDocument(string(b))
	if err != nil {
		return nil, err
	}
	return NewIplayerDocument(doc), nil
}

func (ip *IplayerDocument) selection(selector string) *goquery.Selection {
	return ip.idoc.Find(selector)
}

func (ip *IplayerDocument) extraPages() []BeebURL {
	return ip.morePages(".view-more-container")
}

func (ip *IplayerDocument) newProgrammeSelection() *ProgrammeSelection {
	return &ProgrammeSelection(ip.idoc.Find(".list-item.programme"))
}

func (ps ProgrammeSelection) hasExtraProgrammes() bool {
	sel := ps.sel.Find(".view-more-container").AttrOr("href", "")
	return sel != ""
}

//// CollectNextPage checks for a pagination div at the bottom of the
//// Programme listing page. If found, it returns a slice of urls
//// for the same category.
//func (ip *IplayerDocument) CollectNextPage() {
//	ip.NextPages = ip.morePages(".page > a")
//}
//
//func (ip *IplayerDocument) nextPage() string {
//	return ip.selection(".page > a").AttrOr("href", "")
//}

// CollectSubPages collects for every Programme pontentially available
// canonical programme urls.
// (For example, the category comedy site, will only list the most recent
// episode of a Programme, and then link to The Programme's site for more available
// episodes.)
//func (ip *IplayerDocument) CollectSubPages() {
//	ip.SubPages = ip.morePages(".view-more-container")
//}
//

func (ip *IplayerDocument) morePages(selection string) []BeebURL {
	var bu []BeebURL
	sel := ip.selection(selection)
	sel.Each(func(i int, s *goquery.Selection) {
		bu = append(bu, BeebURL(bbcprefix+s.AttrOr("href", "")))
	})
	return bu
}

//
//func (ip *IplayerDocument) pages() []*Pager {
//	ip.CollectNextPage()
//	ip.CollectSubPages()
//	return ip.pages()
//}


func (ip *IplayerDocument) programmes(c chan<- []*Programme) {
	var programmes []*Programme
		ip.idoc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
			title := findTitle(s)
			subtitle := findSubtitle(s)
			synopsis := findSynopsis(s)
			pid := findPid(s)
			thumbnail := findThumbnail(s)
			url := findURL(s)
			np := newProgramme(title, subtitle, synopsis, pid, thumbnail, url)
		programmes = append(programmes, np)
		})
	c <- programmes
}

//func (th TestHtmlURL) UrlDoc() (*IplayerDocument, error) {
//	file, err := ioutil.ReadFile(string(th))
//
//	if err != nil {
//		return nil, err
//	}
//	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
//	if err != nil {
//		return nil, err
//	}
//	return NewIplayerDocument(doc), nil
//}

// Programme represents an Iplayer TV programme. It consists of
// the programme's title, subtitle, a short programme description,
// The Iplayer Programme ID, the url to its thumbnail, and the url
// to the programme's website.
type Programme struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Synopsis  string `json:"synopsis"`
	Pid       string `json:"pid"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
	Index     int    `json:"index"`
}

func newProgramme(title, subtitle, synopsis, pid,
	thumbnail, url string) *Programme {
	return &Programme{
		Title:     title,
		Subtitle:  subtitle,
		Synopsis:  synopsis,
		Pid:       pid,
		Thumbnail: thumbnail,
		URL:       url,
	}
}

// String returns a string for type Programme,
// starting with Index, followed by Programme Title and Programme Subtitle.
func (p *Programme) String() string {
	return fmt.Sprintf("%d:  %s  %s", p.Index, p.Title, p.Subtitle)
}

// Category struct represents an Iplayer programme category.
// It has the name of the category, like "films" or "comedy" and
// a list of the tv programmes of said category.
type Category struct {
	Name       string       `json:"category"`
	Programmes []*Programme `json:"programmes"`
}

// NewCategory returns a new Category struct for a given
// category name and list of programmes.
func NewCategory(name string, programmes []*Programme) *Category {
	return &Category{name, programmes}
}

// Programmes iterates over an goquery.Document,
// finding every Programme and finally returning them.
func Programmes(s Searcher) ([]*Programme, error) {
	var programmes []*Programme
	doc, err := s.loadDocument()
	if err != nil {
		panic(err)
	}
	progs := make(chan []*Programme)
	go doc.programmes(progs)
	programmes = append(programmes, <-progs...)
	return programmes, nil
}

////func nextPages(pager Pager) []string {
////	var results []string
////	results = append(results, pager.NextPages()...)
////	return results
////}
//
//func subPages(pager Pager) []string {
//	var results []string
//	results = append(results, pager.SubPages()...)
//	return results
//}

func findTitle(s *goquery.Selection) string {
	return s.Find(".secondary > .title").Text()
}

func findSubtitle(s *goquery.Selection) string {
	return s.Find(".secondary > .subtitle").Text()
}

func findURL(s *goquery.Selection) string {
	return "www.bbc.co.uk" + s.Find("a").AttrOr("href", "")
}

func findThumbnail(s *goquery.Selection) string {
	return s.Find(".rs-image > picture > source").AttrOr("srcset", "")
}
// TODO make sure findPID works with films a-z.
func findPid(s *goquery.Selection) string {
	pid := s.AttrOr("data-ip-id", "")
	if pid != "" {
		return pid
	}
	return s.Find(".list-item-inner > a").AttrOr("data-episode-id", "")
}

func findSynopsis(s *goquery.Selection) string {
	return s.Find(".synopsis").Text()
}
