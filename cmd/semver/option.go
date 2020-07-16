package main

type cmdlineOption struct {
	showVersion  bool
	showUsage    bool
	isPrerelease bool
	hasMetadata  bool
	jsonOutput   bool
	target       string
}

const usage = `"semver" is a command-line tool for parsing "Semantic Versioning 2.0.0".

Usage:
  semver [OPTION]... <version>

Condition options:
  -p, --is-prerelease
                    Check the version is a prerelease version
  -m, --has-metadata
                    Check the version has metadata

Output options:
  --json            output in JSON format (2-space indentation)

Other options:
  -?, -h, --help    display this help and exit
  -v, --version     output program version and exit

GitHub repository URL: https://github.com/masa213f/semver
`

func parseOptions(args []string) (*cmdlineOption, error) {
	opt := cmdlineOption{}
	for i := 0; i < len(args); i++ {
		o := args[i]
		switch o {
		// Condition
		case "-p", "--is-prerelease":
			opt.isPrerelease = true
		case "-m", "--has-metadata":
			opt.hasMetadata = true
		// output options
		case "--json":
			opt.jsonOutput = true
		// other options
		case "-?", "-h", "--help":
			opt.showUsage = true
		case "-v", "--version":
			opt.showVersion = true
		// target
		default:
			opt.target = o
		}
	}

	return &opt, nil
}
