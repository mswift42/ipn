package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Programme struct {
	index     int
	name      string
	pid       string
	episode   int
	series    int
	thumbnail string
	url       string
}

func main() {
	doc, err := goquery.NewDocument("http://www.bbc.co.uk/iplayer/group/most-popular")
	if err != nil {
		fmt.Println(err)
	}
	doc.Find(".list-item episode numbered").Each(func(i int, s *goquery.Selection) {
		title := s.Find("").Text()
		fmt.Println(title)
	})
}
