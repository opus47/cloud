package main

import (
	"github.com/mkideal/cli"
	"log"
	"os"

	discd "github.com/opus47/cloud/discd/lib"
)

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * - Command Implementations
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

func doInfo() {
	info, err := discd.ReadDiscInfo()
	if err != nil {
		log.Printf("%s", err)
	} else {
		log.Printf("%s", info)
	}
}

func doRip(i int) {
	log.Printf("rip ~ %d", i)
	discd.RipTrack(int(i))
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * - Command Setup
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

type rootT struct {
	cli.Helper
}

type infoT struct {
	cli.Helper
}

type ripT struct {
	cli.Helper
	Track int `cli:"*t" usage:"integer"`
}

var help = cli.HelpCommand("display help information")

var root = &cli.Command{
	Desc: "discd, the compact disc daemon",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		//command without subcommand does nothing
		log.Print(ctx.Command().Usage(ctx))
		return nil
	},
}

var info = &cli.Command{
	Desc: "get disc information",
	Name: "info",
	Argv: func() interface{} { return new(infoT) },
	Fn: func(ctx *cli.Context) error {
		doInfo()
		return nil
	},
}

var rip = &cli.Command{
	Desc: "rip a track",
	Name: "rip",
	Argv: func() interface{} { return new(ripT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*ripT)
		doRip(argv.Track)
		return nil
	},
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * - Entrypoint
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

func main() {
	log.SetFlags(0)

	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(info),
		cli.Tree(rip),
	).Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
