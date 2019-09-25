package main

import (
	"errors"
	"flag"
	"io/ioutil"
)

type options struct {
	showVersion  bool
	showUsage    bool
	isPreRelease bool
	hasBuildTag  bool
	target       string
}

func parseOptions(args []string) (*options, error) {
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	// Suppress default error output
	flagSet.SetOutput(ioutil.Discard)
	flagSet.Usage = func() {}

	opt := options{}
	flagSet.BoolVar(&opt.showVersion, "v", false, "version")
	flagSet.BoolVar(&opt.showUsage, "h", false, "help")
	flagSet.BoolVar(&opt.isPreRelease, "-p", false, "pre-release")
	flagSet.BoolVar(&opt.hasBuildTag, "-b", false, "build")

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
