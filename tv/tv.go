package tv

import (
	"bytes"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

const bbcprefix = "http://www.bbc.co.uk"

type Searcher interface {
	urlDoc() (*goquery.Document, error)
}

type BeebURL string
type TestHtmlURL string

func (b BeebURL) urlDoc() (*goquery.Document, error) {
	doc, err := goquery.NewDocument(string(b))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (th TestHtmlURL) urlDoc() (*goquery.Document, error) {
	file, err := ioutil.ReadFile(string(th))

	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

type Programme struct {
	Title     string
	Subtitle  string
	Synopsis  string
	Pid       string
	Thumbnail string
	URL       string
	Index     int
}

func newProgramme(title, subtitle, synopsis, pid,
	thumbnail, url string) *Programme {
	return &Programme{title, subtitle, synopsis, pid,
		thumbnail, url, 0}
}

// Category struct represents an Iplayer programme category.
// It has a name of the category, like "films" or "comedy" and
// a list of the tv programmes of said category.
type Category struct {
	Name       string
	Programmes []*Programme
}

func NewCategory(name string, programmes []*Programmes) {
	return &Category{name, programmes}
}

func Programmes(s Searcher) ([]*Programme, error) {
	var programmes []*Programme
	doc, err := s.urlDoc()
	if err != nil {
		return nil, err
	}
	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		title := findTitle(s)
		subtitle := findSubtitle(s)
		synopsis := findSynopsis(s)
		pid := findPid(s)
		thumbnail := findThumbnail(s)
		url := findURL(s)
		programmes = append(programmes, newProgramme(title, subtitle,
			synopsis, pid, thumbnail, url))
	})
	return programmes, nil
}
func hasSubPage(doc *goquery.Document) string {
	return doc.Find(".view-more-container").AttrOr("href", "")
}
func SubPage(doc *goquery.Document) *goquery.Document {
	return doc
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
	return s.Find(".r-image").AttrOr("data-ip-src", "")
}
func findPid(s *goquery.Selection) string {
	return s.Find(".list-item-inner > a").AttrOr("data-episode-id", "")
}
func findSynopsis(s *goquery.Selection) string {
	return s.Find(".synopsis").Text()
}
