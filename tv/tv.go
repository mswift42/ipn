package tv

import (
	"fmt"
	"sync"

	"github.com/mswift42/goquery"
)

const bbcprefix = "http://www.bbc.co.uk"

// BeebURL represents an Iplayer URL.
type BeebURL string

type IplayerDocument struct {
	idoc *goquery.Document
}

type IplayerDocumentResult struct {
	idoc  *goquery.Document
	Error error
}

// iplayerSelection is an iplayer web page
// list-item section of one Programme.
type iplayerSelection struct {
	sel *goquery.Selection
}

// iplayerSelectionResult has either an iplayerSelection for
// an iplayer programme, or, if it has a link to a "more Programmes available"
// site, said link.
type iplayerSelectionResult struct {
	prog     *Programme
	progpage string
}

type MainCategoryDocument struct {
	ip        *IplayerDocument
	NextPages []string
}

type mainCategory struct {
	name string
	docs []*IplayerDocument
}

var seen = make(map[Searcher]bool)
var mutex = &sync.Mutex{}

func seenLink(s Searcher) bool {
	mutex.Lock()
	if seen[s] == false {
		seen[s] = true
		mutex.Unlock()
		return false
	}
	mutex.Unlock()
	return true
}

type page struct {
	result []*iplayerSelectionResult
}

func (pr *page) programPageUrls() []Searcher {
	var urls []Searcher
	for _, i := range pr.result {
		if i.progpage != "" {
			urls = append(urls, BeebURL(bbcprefix+i.progpage))
		}
	}
	return urls
}

func NewIplayerDocument(doc *goquery.Document) *IplayerDocument {
	return &IplayerDocument{doc}
}

// func NewMainCategoryDocument(bu BeebURL) (*MainCategoryDocument, error) {
// 	c := make(chan *IplayerDocumentResult)
// 	doc := <-c
// 	if doc.Error != nil {
// 		return nil, doc.Error
// 	}
// 	idoc := IplayerDocument{doc.idoc}
// 	return &MainCategoryDocument{&idoc, idoc.nextPages()}, nil
// }

// func (mcd *MainCategoryDocument) programmes() []*Programme {
// 	var progs []*Programme
// 	docs := mcd.collectNextPages()
// 	return progs
// }

func newMainCategoryDocument(s Searcher) (*MainCategoryDocument, error) {
	c := make(chan *IplayerDocumentResult)
	go s.loadDocument(c)
	idr := <-c
	if idr.Error != nil {
		return nil, idr.Error
	}
	doc := IplayerDocument{idr.idoc}
	return &MainCategoryDocument{&doc, doc.nextPages()}, nil
}

func collectProgramPages(ires []*iplayerSelectionResult) []*IplayerDocumentResult {
	var docres []*IplayerDocumentResult
	pg := &page{ires}
	morepages := pg.programPageUrls()
	scan := make(chan Searcher, 20)
	idrchan := make(chan *IplayerDocumentResult)
	jobs := 0
	go collectDocument(scan, idrchan)
	for _, i := range morepages {
		if !seenLink(i) {
			jobs++
			scan <- i
		}
	}
	for i := 0; i < jobs; i++ {
		docres = append(docres, <-idrchan)
	}
	close(idrchan)
	return docres
}
func collectDocument(in chan Searcher, out chan *IplayerDocumentResult) {
	c := make(chan *IplayerDocumentResult)
	for u := range in {
		go u.loadDocument(c)
		idr := <-c
		if idr.Error != nil {
			out <- &IplayerDocumentResult{nil, idr.Error}
		} else {
			out <- &IplayerDocumentResult{idr.idoc, nil}
		}
	}
}

func (ipd *IplayerDocument) programmes(c chan<- []*Programme, wg *sync.WaitGroup) {
	var results []*Programme
	ipd.idoc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		results = append(results, findProgramme(s))
	})
	c <- results
}

func (mcd *MainCategoryDocument) collectNextPages() []*IplayerDocumentResult {
	var results []*IplayerDocumentResult
	sc := make(chan Searcher)
	idrc := make(chan *IplayerDocumentResult)
	go collectDocument(sc, idrc)
	for _, i := range mcd.NextPages {
		go func(url string) {
			bu := BeebURL("http://bbc.co.uk" + url)
			sc <- bu
		}(i)
	}
	defer close(sc)
	for range mcd.NextPages {
		results = append(results, <-idrc)
	}
	return results
}

func (mcd *MainCategoryDocument) collectViewMorePages() []*IplayerDocumentResult {
	var results []*IplayerDocumentResult
	iselc := make(chan []*iplayerSelectionResult)
	go mcd.ip.selectionResults(iselc)
	fp := <-iselc
	fpr := collectProgramPages(fp)
	results = append(results, fpr...)
	nextpages := mcd.collectNextPages()
	for _, i := range nextpages {
		idoc := IplayerDocument{i.idoc}
		go idoc.selectionResults(iselc)
	}
	for range nextpages {
		page := <-iselc
		results = append(results, collectProgramPages(page)...)
	}
	return results
}

func (mcd *MainCategoryDocument) programmes() []*Programme {
	var progs []*Programme
	var selres []*iplayerSelectionResult
	selreschan := make(chan []*iplayerSelectionResult)
	progchan := make(chan []*Programme)
	wg := &sync.WaitGroup{}
	go mcd.ip.selectionResults(selreschan)
	selres = append(selres, <-selreschan...)
	docs := mcd.collectNextPages()
	for _, i := range docs {
		idoc := &IplayerDocument{i.idoc}
		go idoc.selectionResults(selreschan)
	}
	for range docs {
		selres = append(selres, <-selreschan...)
	}
	idr := collectProgramPages(selres)
	for _, i := range idr {
		doc := &IplayerDocument{i.idoc}
		go doc.programmes(progchan, wg)
	}
	for _, i := range selres {
		if i.prog != nil {
			progs = append(progs, i.prog)
		}
	}
	return progs
}

// TODO add method that iterates over selecetionResults and
// collects all progpage documents.

func (bu BeebURL) loadDocument(c chan<- *IplayerDocumentResult) {
	doc, err := goquery.NewDocument(string(bu))
	if err != nil {
		c <- &IplayerDocumentResult{nil, err}
	}
	c <- &IplayerDocumentResult{doc, nil}
}

func (ipd *IplayerDocument) selectionResults(c chan<- []*iplayerSelectionResult) {
	var res []*iplayerSelectionResult
	ipd.idoc.Find(".list-item").Each(func(i int, selection *goquery.Selection) {
		expage := findProgrammeSite(selection)
		if expage != "" {
			res = append(res, &iplayerSelectionResult{nil, expage})
		} else {
			res = append(res,
				&iplayerSelectionResult{findProgramme(selection), ""})
		}
	})
	c <- res
}

//func (b BeebURL) LoadDocument() (*IplayerDocument, error) {
//	doc, err := goquery.NewDocument(string(b))
//	if err != nil {
//		return nil, err
//	}
//	return NewIplayerDocument(doc), nil
//}

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
//// Programme listing page. If found, git@github.com:mswift42/rtc.gitit returns a slice of urls
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

// func (mp *MainCategoryDocument) Programmes() ([]*Programme, []string) {
// 	var progs []*Programme
// 	var extraurls []string
// 	progch := make(chan []*Programme)
// 	urlch := make(chan []string)
// 	fmt.Println(mp.NextPages)
// 	go mp.ip.programmes(progch, urlch)
// 	//for _, i := range mp.NextPages {
// 	//	go func(url string) {
// 	//		bu := BeebURL(url)
// 	//		nd, _ := bu.LoadDocument()
// 	//		go nd.programmes(progch, urlch)
// 	//	}(i)
// 	//}
// 	for p := range progch {
// 		progs = append(progs, p...)
// 	}
// 	for u := range urlch {
// 		extraurls = append(extraurls, u...)
// 	}
// 	fmt.Println(extraurls)
// 	return progs, extraurls
// }

// func (ip *IplayerDocument) programmes(progch chan []*Programme, urlch chan<- []string) {
// 	var progs []*Programme
// 	var extraurls []string
// 	ip.idoc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
// 		prog, eu := findProgramme(i, s)
// 		progs = append(progs, prog)
// 		fmt.Println(progs)
// 		if eu != "" {
// 			extraurls = append(extraurls, string(eu))
// 			fmt.Println(extraurls)
// 		}
// 	})
// 	progch <- progs
// 	urlch <- extraurls
// }

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
// func Programmes(s Searcher) ([]*Programme, error) {
// 	var programmes []*Programme
// 	c := make(chan *IplayerDocumentResult)
// 	go s.loadDocument(c)
// 	doc := <-c
// 	var urls []string
// 	progch := make(chan []*Programme)
// 	extraurls := make(chan []string)
// 	doc.programmes(progch, extraurls)
// 	for p := range progch {
// 		programmes = append(programmes, p...)
// 	}
// 	fmt.Println(programmes)
// 	for u := range extraurls {
// 		urls = append(urls, u...)
// 	}
// 	fmt.Println(extraurls)
// 	//for _, i := range urls {
// 	//	go func(url string) {
// 	//		bu := BeebURL(url)
// 	//		doc, err := bu.LoadDocument()
// 	//		if err != nil {
// 	//			log.Fatal(err)
// 	//		}
// 	//		doc.programmes(progch, extraurls)
// 	//		for p := range progch {
// 	//			programmes = append(programmes, p...)
// 	//		}
// 	//	}(i)
// 	//}
// 	return programmes, nil
// }

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

func findProgramme(s *goquery.Selection) *Programme {
	title := findTitle(s)
	subtitle := findSubtitle(s)
	synopsis := findSynopsis(s)
	pid := findPid(s)
	thumbnail := findThumbnail(s)
	url := findURL(s)
	np := newProgramme(title, subtitle, synopsis, pid, thumbnail, url)
	return np
}

// findProgrammeSite returns the link the iplayer programme's page
// if a class view-more-container is present.
func findProgrammeSite(s *goquery.Selection) string {
	return s.Find(".view-more-container").AttrOr("href", "")
}

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
