package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/masa213f/semver"
)

const (
	exitStatusSuccess          = 0
	exitStatusParseFailure     = 1
	exitStatusConditionFailure = 2
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
	fmt.Fprintf(o, usage)
}

func main() {
	cmdOpt, err := parseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(exitStatusInvalidOption)
	}

	if cmdOpt.showVersion {
		showVersion(os.Stdout)
		os.Exit(exitStatusSuccess)
	}

	if cmdOpt.showUsage {
		showUsage(os.Stdout)
		os.Exit(exitStatusSuccess)
	}

	ver, err := semver.Parse(strings.TrimSpace(cmdOpt.target))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(exitStatusParseFailure)
	}

	if cmdOpt.isPreRelease {
		if !ver.IsPrerelease() {
			fmt.Fprintln(os.Stderr, "not prerelease version")
			os.Exit(exitStatusConditionFailure)
		}
	}
	if cmdOpt.hasBuildMeta {
		if !ver.HasMetadata() {
			fmt.Fprintln(os.Stderr, "no metadata")
			os.Exit(exitStatusConditionFailure)
		}
	}

	outOpt := newOutputOption()
	if cmdOpt.jsonOutput {
		outOpt.format = "json"
	}
	output(os.Stdout, ver, outOpt)

	os.Exit(exitStatusSuccess)
}
