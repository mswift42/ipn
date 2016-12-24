package tv

import (
	"bytes"
	"io/ioutil"

	"encoding/json"

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
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Synopsis  string `json:"synopsis"`
	Pid       string `json:"pid"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
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

// ProgrammeDB stores all queried categories.
type programmeDB struct {
	Catogories []*Category `json:"categories"`
}

func newProgrammeDB(cats []*Category) *programmeDB {
	return &programmeDB{cats}
}

func (pdb *programmeDB) toJson() ([]byte, error) {
	marshalled, err := json.MarshalIndent(pdb)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

func (pdb *programmeDB) save(jsonfile []byte, filename string) error {
	return ioutil.WriteFile(filename, jsonfile, 0644)
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
