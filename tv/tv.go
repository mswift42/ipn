package tv

import "github.com/PuerkitoBio/goquery"

const bbcprefix = "http://www.bbc.co.uk"

type Searcher interface {
	programmesForURL(url string) ([]*Programme, error)
}

type beebURL string

func (b beebURL) ProgrammesForURL(url string) ([]*Programme, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return programmes(doc)
}

type Programme struct {
	Title     string
	Subtitle  string
	Synopsis  string
	Pid       string
	Thumbnail string
	Url       string
	Index     int
}

func newProgramme(title, subtitle, synopsis, pid,
	thumbnail, url string) *Programme {
	return &Programme{title, subtitle, synopsis, pid,
		thumbnail, url, 0}
}


func (q Queryurl) DocURL() (*goquery.Document, error) {
	doc, err := goquery.NewDocument(string(q))
	if err != nil {
		return nil, err
	}
	return doc, nil
}
func programmes(doc *goquery.Document) []*Programme {
	var programmes []*Programme
	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		title := findTitle(s)
		subtitle := findSubtitle(s)
		synopsis := findSynopsis(s)
		pid := findPid(s)
		thumbnail := findThumbnail(s)
		url := findUrl(s)
		programmes = append(programmes, newProgramme(title, subtitle,
			synopsis, pid, thumbnail, url))
	})
	return programmes
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

func findUrl(s *goquery.Selection) string {
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
