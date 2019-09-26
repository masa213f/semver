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
	exitStatusSuccess          = 0
	exitStatusParseFailure     = 1
	exitStagusConditionFailure = 2
	exitStatusInvalidOption    = 3
	exitStatusInternalError    = 4
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

func show(o io.Writer, ver *semver.Version) {
	bytes, _ := json.Marshal(ver)
	fmt.Println(string(bytes))
}

func main() {
	opt, err := parseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(exitStatusInvalidOption)
	}

	if opt.showVersion {
		showVersion(os.Stdout)
		os.Exit(exitStatusSuccess)
	}

	if opt.showUsage {
		showUsage(os.Stdout)
		os.Exit(exitStatusSuccess)
	}

	target := strings.TrimSpace(opt.target)
	ver, err := semver.Parse(strings.TrimSpace(opt.target))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(exitStatusParseFailure)
	}

	if opt.isConditionCheck() {
		if opt.isPreRelease && !ver.IsPreRelease() {
			fmt.Fprintf(os.Stderr, "%s is not pre-release\n", target)
			os.Exit(exitStagusConditionFailure)
		}
		if opt.hasBuildMeta && !ver.HasBuildMeta() {
			fmt.Fprintf(os.Stderr, "%s does not have build metadata\n", target)
			os.Exit(exitStagusConditionFailure)
		}
	}
	show(os.Stdout, ver)
	os.Exit(exitStatusSuccess)
}
