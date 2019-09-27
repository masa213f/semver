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

	target := strings.TrimSpace(cmdOpt.target)
	ver, err := semver.Parse(strings.TrimSpace(cmdOpt.target))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(exitStatusParseFailure)
	}

	if !cmdOpt.isConditionCheck() {
		// output all fields
		if cmdOpt.jsonOutput {
			outputJSON(os.Stdout, ver, nil)
		} else {
			outputText(os.Stdout, ver, nil)
		}
		os.Exit(exitStatusSuccess)
	}

	outOpt := newOutputOption()

	if cmdOpt.isPreRelease {
		if !ver.IsPreRelease() {
			fmt.Fprintf(os.Stderr, "%s is not pre-release version\n", target)
			os.Exit(exitStatusConditionFailure)
		}
		outOpt.displayPreRelease = true
	}
	if cmdOpt.hasBuildMeta {
		if !ver.HasBuildMeta() {
			fmt.Fprintf(os.Stderr, "%s does not have build metadata\n", target)
			os.Exit(exitStatusConditionFailure)
		}
		outOpt.displayBuildMeta = true
	}

	// output selected fields
	if cmdOpt.jsonOutput {
		outputJSON(os.Stdout, ver, outOpt)
	} else {
		outputText(os.Stdout, ver, outOpt)
	}
	os.Exit(exitStatusSuccess)
}
