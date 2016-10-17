package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/mswift42/ipn/tv"

	"github.com/PuerkitoBio/goquery"
)

func Programmes() ([]*tv.Programme, error) {
	popurl := "http://www.bbc.co.uk/iplayer/group/most-popular"
	doc, err := goquery.NewDocument(popurl)
	if err != nil {
		return nil, err
	}
	var programmes []*tv.Programme
	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		title := findTitle(s)
		subtitle := findSubtitle(s)
		synopsis := findSynopsis(s)
		pid := findPid(s)
		thumbnail := findThumbnail(s)
		url := findUrl(s)
		programmes = append(programmes, tv.NewProgramme(title, subtitle, synopsis, pid,
			thumbnail, url))

	})
	return programmes, nil
}

func loadTestHtml(filename string) *goquery.Document {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}
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
func main() {
	html := loadTestHtml("iplayermostpopular.html")
	html.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		pid := findPid(s)
		fmt.Println(pid)
		title := findTitle(s)
		subtitle := findSubtitle(s)
		fmt.Println(title)
		fmt.Println(subtitle)
	})
}
