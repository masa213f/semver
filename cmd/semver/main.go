package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/masa213f/semver"
)

var (
	version         string
	showVersionFlag = flag.Bool("v", false, "version")
	showHelpFlag    = flag.Bool("h", false, "help")
)

func showVersion() {
	fmt.Println(version)
}

func showHelp() {
	fmt.Println("T.B.D.")
}

func show(ver *semver.Version) {
	fmt.Println(*ver)
}

func main() {
	flag.Parse()

	if *showVersionFlag {
		showVersion()
		os.Exit(0)
	}

	if *showHelpFlag {
		showHelp()
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) != 1 {
		showHelp()
		os.Exit(1)
	}

	ver, err := semver.Parse(args[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	show(ver)
}
