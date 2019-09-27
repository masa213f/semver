package main

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

const usage = `"semver" is a command-line tool for parsing "Semantic Versioning 2.0.0".

Usage:
  semver [OPTION]... <version>

Condition options:
  -p, --is-prerelease
                    verify the version is a pre-release version and output the pre-release fields
  -b, --has-buildmeta
                    verify the version has build metadata and output the build metadata fields

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
			opt.isPreRelease = true
		case "-b", "--has-buildmeta":
			opt.hasBuildMeta = true
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
