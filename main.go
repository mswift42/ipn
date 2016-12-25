package main

import (
	"fmt"

	"github.com/mswift42/ipn/categories"
)

const (
	mp = "http://www.bbc.co.uk/iplayer/group/most-popular"
)

func main() {
	prog, err := categories.AllCategories()
	if err != nil {
		fmt.Println("Oops, error in fetching all categories: ", err)
	}
	for _, i := range prog {
		if i != nil {
			fmt.Println("\n\n", i.Name, "\n\n")
			for _, j := range i.Programmes {
				if j != nil {
					fmt.Println(j.Title)
				}
			}
		}
	}
}
