package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/mswift42/ipn/films"
	"github.com/mswift42/ipn/mostpopular"
)

const (
	mp = "http://www.bbc.co.uk/iplayer/group/most-popular"
)

func main() {
	pop, _ := goquery.NewDocument(mp)
	programmes := mostpopular.Programmes(pop)

	for _, i := range programmes {
		fmt.Println(i.Title)
		fmt.Println(i.Synopsis)
	}
	films, _ := films.Programmes()
	for _, i := range films {
		fmt.Println(i.Title)
		fmt.Println(i.Synopsis)
	}
}
