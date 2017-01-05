package tv

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"

	"time"

	"github.com/mswift42/goquery"
)

const bbcprefix = "http://www.bbc.co.uk"

type Searcher interface {
	UrlDoc() (*goquery.Document, error)
}

// BeebURL represents an Iplayer URL.
type BeebURL string

type TestHtmlURL string

type IplayerDocument struct {
	idoc *goquery.Document
}

func NewIplayerDocument(doc *goquery.Document) *IplayerDocument {
	return &IplayerDocument{doc}
}

func (b BeebURL) UrlDoc() (*goquery.Document, error) {
	doc, err := goquery.NewDocument(string(b))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (ip *IplayerDocument) tvSelection(selector string) *goquery.Selection {
	return ip.idoc.Find(selector)
}

func (ip *IplayerDocument) NextPages() []BeebURL {
	var bu []BeebURL
	sel := ip.tvSelection(".page > a")
	sel.Each(func(i int, s *goquery.Selection) {
		bu = append(bu, BeebURL(bbcprefix+s.AttrOr("href", "")))
	})
	return bu
}

func (th TestHtmlURL) UrlDoc() (*goquery.Document, error) {
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

func (cat *Category) nextPages() ([]BeebURL, error) {
	return nil, nil
}

// ProgrammeDB stores all queried categories.
type programmeDB struct {
	Categories []*Category `json:"categories"`
	Saved      time.Time   `json:"saved"`
}

func newProgrammeDB(cats []*Category, saved time.Time) *programmeDB {
	return &programmeDB{Categories: cats, Saved: saved}
}

func LoadProgrammeDbFromJSON(jsonfilename string) (*programmeDB, error) {
	file, err := ioutil.ReadFile(jsonfilename)
	if err != nil {
		return nil, err
	}
	var pdb programmeDB
	json.Unmarshal(file, &pdb)
	return &pdb, nil
}

func (pdb *programmeDB) toJSON() ([]byte, error) {
	marshalled, err := json.MarshalIndent(pdb, "", "\t")
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

func (pdb *programmeDB) Save(filename string) error {
	pdb.Saved = time.Now()
	pdb.index()
	json, err := pdb.toJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, json, 0644)
}

func (pdb *programmeDB) index() {
	index := 0
	for _, i := range pdb.Categories {
		for _, j := range i.Programmes {
			j.Index = index
			index++
		}
	}
}

// Programmes iterates over an goquery.Document,
// finding every Programme and finally returning them.
func Programmes(s Searcher) ([]*Programme, error) {
	var programmes []*Programme
	doc, err := s.UrlDoc()
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
		np := newProgramme(title, subtitle, synopsis, pid, thumbnail, url)
		subpage := np.SubPage(s)
		if subpage != bbcprefix {
			subpageprogrammes, err := Programmes(BeebURL(subpage))
			if err != nil {
				log.Println(err)
			}
			programmes = append(programmes, subpageprogrammes...)
		} else {
			if np != nil {
				programmes = append(programmes, np)
			}
		}

	})
	return programmes, nil
}

func (p *Programme) SubPage(s *goquery.Selection) string {
	return bbcprefix + s.Find(".view-more-container").AttrOr("href", "")
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
func findPid(s *goquery.Selection) string {
	return s.Find(".list-item-inner > a").AttrOr("data-episode-id", "")
}
func findSynopsis(s *goquery.Selection) string {
	return s.Find(".synopsis").Text()
}
