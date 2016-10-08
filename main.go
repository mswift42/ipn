package main

import (
	"fmt"
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
func main() {
	doc, err := goquery.NewDocument("http://www.bbc.co.uk/iplayer/group/most-popular")
	if err != nil {
		fmt.Println(err)
	}
	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".secondary > .title ").Text()
		fmt.Println(i, title)
	})
}
