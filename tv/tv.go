package tv

import (
	"fmt"

	"log"

	"github.com/mswift42/goquery"
	"sync"
)

const bbcprefix = "http://www.bbc.co.uk"

// BeebURL represents an Iplayer URL.
type BeebURL string

type IplayerDocument struct {
	idoc *goquery.Document
}

type iplayerSelection struct {
	sel *goquery.Selection
}

type MainCategoryDocument struct {
	ip        *IplayerDocument
	NextPages []string
}

func NewIplayerDocument(doc *goquery.Document) *IplayerDocument {
	return &IplayerDocument{doc}
}

func NewMainCategoryDocument(ip *IplayerDocument) *MainCategoryDocument {
	nextpages := ip.nextPages()
	return &MainCategoryDocument{ip, nextpages}
}

func (b BeebURL) LoadDocument() (*IplayerDocument, error) {
	doc, err := goquery.NewDocument(string(b))
	if err != nil {
		return nil, err
	}
	return NewIplayerDocument(doc), nil
}

func (ip *IplayerDocument) selection(selector string) *goquery.Selection {
	return ip.idoc.Find(selector)
}

func (ip *IplayerDocument) extraPages() []string {
	return ip.morePages(".view-more-container")
}

func (isel iplayerSelection) hasExtraProgrammes() bool {
	extra := isel.sel.Find(".view-more-container").AttrOr("href", "")
	return extra != ""
}

//// CollectNextPage checks for a pagination div at the bottom of the
//// Programme listing page. If found, it returns a slice of urls
//// for the same category.
//func (ip *IplayerDocument) CollectNextPage() {
//	ip.NextPages = ip.morePages(".page > a")
//}
//
func (ip *IplayerDocument) nextPages() []string {
	return ip.morePages(".page > a")
}

// CollectSubPages collects for every Programme pontentially available
// canonical programme urls.
// (For example, the category comedy site, will only list the most recent
// episode of a Programme, and then link to The Programme's site for more available
// episodes.)
//func (ip *IplayerDocument) CollectSubPages() {
//	ip.SubPages = ip.morePages(".view-more-container")
//}
//

func (ip *IplayerDocument) morePages(selection string) []string {
	var url []string
	sel := ip.selection(selection)
	sel.Each(func(i int, s *goquery.Selection) {
		url = append(url, bbcprefix+s.AttrOr("href", ""))
	})
	return url
}

//
//func (ip *IplayerDocument) pages() []*Pager {
//	ip.CollectNextPage()
//	ip.CollectSubPages()
//	return ip.pages()
//}

var mutex sync.Mutex

func (mp *MainCategoryDocument) Programmes() ([]*Programme, []string) {
	var progs []*Programme
	var extraurls []string
	fmt.Println(mp.NextPages)
		pr, eu := mp.ip.programmes()
		progs = append(progs, pr...)
		extraurls = append(extraurls, eu...)
	for _, i := range mp.NextPages {
		go func(url string) {
			bu := BeebURL(url)
			nd, _ := bu.LoadDocument()
			pr, eu := nd.programmes()
			mutex.Lock()
			progs = append(progs, pr...)
			extraurls = append(extraurls, eu...)
			defer mutex.Unlock()
		}(i)
	}
	fmt.Println(progs)
	fmt.Println(extraurls)
	return progs, extraurls
}

func (ip *IplayerDocument) programmes() ([]*Programme, []string) {
	var progs []*Programme
	var extraurls []string
	ip.idoc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		isel := iplayerSelection{s}
		if isel.hasExtraProgrammes() {
			extraurls = append(extraurls, "http://www.bbc.co.uk"+isel.sel.Find(".view-more-container").AttrOr("href", ""))
		}
		title := findTitle(s)
		subtitle := findSubtitle(s)
		synopsis := findSynopsis(s)
		pid := findPid(s)
		thumbnail := findThumbnail(s)
		url := findURL(s)
		np := newProgramme(title, subtitle, synopsis, pid, thumbnail, url)
		progs = append(progs, np)
	})
	return progs, extraurls
}

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
	progs, urls := doc.programmes()
	var mutex = &sync.Mutex{}
	programmes = append(programmes, progs...)
	for _, i := range urls {
		go func(url string) {
			bu := BeebURL(url)
			doc, err := bu.LoadDocument()
			if err != nil {
				log.Fatal(err)
			}
			pr, _ := doc.programmes()
			mutex.Lock()
			programmes = append(programmes, pr...)
			mutex.Unlock()
		}(i)
	}

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
