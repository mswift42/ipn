package main

import (
	"os"
	"time"

	"github.com/mswift42/ipn/categories"
	"github.com/mswift42/ipn/cli"
	"github.com/mswift42/ipn/db"
)

const (
	mp    = "http://www.bbc.co.uk/iplayer/group/most-popular"
	film1 = "http://www.bbc.co.uk/iplayer/categories/films/all?sort=atoz"
)

func main() {

	cats, err := categories.LoadAllCategories()
	newdb := db.NewProgrammeDB(cats, time.Now())
	newdb.Save("db/testjson.json")
	app := cli.InitCli()
	app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
