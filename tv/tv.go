package tv

import "github.com/PuerkitoBio/goquery"

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
func Programmes(url string) ([]*Programme, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
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
	return programmes, nil
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
