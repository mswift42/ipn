package main

import (
	"os"

	"github.com/mswift42/ipn/cli"
)

const (
	mp = "http://www.bbc.co.uk/iplayer/group/most-popular"
)

func main() {
	app := cli.InitCli()
	app.Run(os.Args)
}
