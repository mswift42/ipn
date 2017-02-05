package main

import (
	"fmt"
	"os"

	"github.com/mswift42/ipn/categories"
	"github.com/mswift42/ipn/cli"
)

const (
	mp = "http://www.bbc.co.uk/iplayer/group/most-popular"
)

func main() {
	app := cli.InitCli()
	app.Run(os.Args)
	cats, _ := categories.AllCategories()
	for _, i := range cats {
		fmt.Println(i.Name)
		for _, j := range i.Programmes {
			fmt.Println(j)
		}
	}
}
