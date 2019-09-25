package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/masa213f/semver"
)

const (
	exitStatusSuccess        = 0
	exitStatusParseFailure   = 2
	exitStatusInvalidOptions = 3
	exitStatusInternalError  = 5
)

var (
	version string
)

func showVersion(o io.Writer) {
	fmt.Fprintln(o, version)
}

func showUsage(o io.Writer) {
	const usage = `
Usage: semver

T.B.D.

Copyright 2019 xxx.
`
	fmt.Fprintf(o, usage)
}

func show(ver *semver.Version) {
	bytes, _ := json.Marshal(ver)
	fmt.Println(string(bytes))
}

func main() {
	opt, err := parseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(exitStatusInvalidOptions)
	}

	if opt.showVersion {
		showVersion(os.Stdout)
		os.Exit(exitStatusSuccess)
	}

	if opt.showUsage {
		showUsage(os.Stdout)
		os.Exit(exitStatusSuccess)
	}

	ver, err := semver.Parse(strings.TrimSpace(opt.target))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(exitStatusParseFailure)
	}
	show(ver)
	os.Exit(exitStatusSuccess)
}
