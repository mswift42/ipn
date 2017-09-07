package main

import (
	"fmt"
	"os"

	"github.com/mswift42/ipn/categories"
	"github.com/mswift42/ipn/cli"
	"github.com/mswift42/ipn/db"
	"time"
)

const (
	mp = "http://www.bbc.co.uk/iplayer/group/most-popular"
)

func main() {
	app := cli.InitCli()
	app.Run(os.Args)
	cats, err := categories.LoadAllCategories()
	fmt.Println(err)
	for _, i := range cats {
		fmt.Println(i)
	}
	newdb := db.NewProgrammeDB(cats, time.Now() )
	newdb.Save("testdb.json")
	fmt.Println(newdb.ListAvailableCategories())
	fmt.Println(newdb.FindTitle("foster"))
	samedb, _ := db.LoadProgrammeDbFromJSON("testdb.json")
	fmt.Println(samedb.FindTitle("strike"))





}
