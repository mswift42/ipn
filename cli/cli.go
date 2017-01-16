package cli

import (
	"fmt"

	"github.com/mswift42/cli"
	"github.com/mswift42/ipn/db"
)

func InitCli() *cli.App {
	tdb, _ := db.LoadProgrammeDbFromJSON("db/testjson.json")
	app := cli.NewApp()
	app.Setup()
	app.Name = "ipn"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all available categories.",
			Action: func(c *cli.Context) error {
				fmt.Println(tdb.ListAvailableCategories())
				return nil
			},
		},
		{
			Name:    "category",
			Aliases: []string{"c"},
			Usage:   "List all programmes for a category.",
			Action: func(c *cli.Context) error {
				fmt.Println(tdb.ListCategory(c.Args().Get(0)))
				return nil
			},
		},
		{
			Name:    "find",
			Aliases: []string{"f"},
			Usage:   "Find a programme.",
			Action: func(c *cli.Context) error {
				fmt.Println(tdb.FindTitle(c.Args().Get(0)))
				return nil
			},
		},
	}
	return app
}
