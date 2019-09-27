package main

import (
	"errors"
	"flag"
	"io/ioutil"
)

type cmdlineOption struct {
	showVersion  bool
	showUsage    bool
	isPreRelease bool
	hasBuildMeta bool
	jsonOutput   bool
	target       string
}

func (opt *cmdlineOption) isConditionCheck() bool {
	if opt.isPreRelease || opt.hasBuildMeta {
		return true
	}
	return false
}

func parseOptions(args []string) (*cmdlineOption, error) {
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	// Suppress default error output
	flagSet.SetOutput(ioutil.Discard)
	flagSet.Usage = func() {}

	opt := cmdlineOption{}
	flagSet.BoolVar(&opt.showVersion, "v", false, "version")
	flagSet.BoolVar(&opt.showUsage, "h", false, "help")
	flagSet.BoolVar(&opt.isPreRelease, "p", false, "pre-release")
	flagSet.BoolVar(&opt.hasBuildMeta, "b", false, "build")
	flagSet.BoolVar(&opt.jsonOutput, "json", false, "json output")

	err := flagSet.Parse(args)
	if err != nil {
		return nil, err
	}

	if opt.showVersion || opt.showUsage {
		return &opt, nil
	}

	tmp := flagSet.Args()
	if len(tmp) != 1 {
		return nil, errors.New("requires 1 argument")
	}
	opt.target = tmp[0]
	return &opt, nil
}
