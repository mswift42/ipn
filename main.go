package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Programme struct {
	index     int
	name      string
	subtitle  string
	pid       string
	episode   int
	series    int
	thumbnail string
	url       string
}

func NewProgramme(index, series, episode int, name, subtitle,
	pid, thumbnail, url string) *Programme {
	return &Programme{index, name, subtitle, pid, episode, series, thumbnail, url}
}
func mostPopular() []*Programme {
	popurl := "http://www.bbc.co.uk/iplayer/group/most-popular"
	doc, err := goquery.NewDocument(popurl)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".list-item")
	return nil
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
func main() {
	html := loadTestHtml("iplayermostpopular.html")
	html.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		thumbnail := findThumbnail(s)
		fmt.Println(thumbnail)
	})
}
