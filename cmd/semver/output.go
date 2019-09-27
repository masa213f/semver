package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/masa213f/semver"
)

type outputOption struct {
	displayVersionCore bool
	displayPreRelease  bool
	displayBuildMeta   bool
}

func newOutputOption() *outputOption {
	return &outputOption{
		displayVersionCore: false,
		displayPreRelease:  false,
		displayBuildMeta:   false,
	}
}

type outputVersion struct {
	Prefix     *string  `json:"prefix,omitempty"`
	Version    *string  `json:"version,omitempty"`
	Major      *uint64  `json:"major,omitempty"`
	Minor      *uint64  `json:"minor,omitempty"`
	Patch      *uint64  `json:"patch,omitempty"`
	PreRelease []string `json:"prerelease,omitempty"`
	Build      []string `json:"build,omitempty"`
}

func newOutputVersion(ver *semver.Version, opt *outputOption) *outputVersion {
	out := outputVersion{}
	if opt == nil || opt.displayVersionCore {
		if ver.Prefix != "" {
			out.Prefix = &ver.Prefix
		}
		out.Version = &ver.Version
		out.Major = &ver.Major
		out.Minor = &ver.Minor
		out.Patch = &ver.Patch
	}
	if (opt == nil || opt.displayPreRelease) && len(ver.PreRelease) != 0 {
		s := []string{}
		for _, pr := range ver.PreRelease {
			s = append(s, pr.ToString())
		}
		out.PreRelease = s
	}
	if (opt == nil || opt.displayBuildMeta) && len(ver.Build) != 0 {
		s := []string{}
		for _, b := range ver.Build {
			s = append(s, string(b))
		}
		out.Build = s
	}
	return &out
}

func outputJSON(o io.Writer, ver *semver.Version, opt *outputOption) {
	out := newOutputVersion(ver, opt)
	bytes, _ := json.Marshal(out)
	fmt.Fprintln(o, string(bytes))
}

func outputText(o io.Writer, ver *semver.Version, opt *outputOption) {
	out := newOutputVersion(ver, opt)

	if out.Prefix != nil {
		fmt.Fprintf(o, "prefix: %s\n", *out.Prefix)
	}
	if out.Version != nil {
		fmt.Fprintf(o, "version: %s\n", *out.Version)
	}

	if out.Major != nil {
		fmt.Fprintf(o, "major: %d\n", *out.Major)
	}
	if out.Minor != nil {
		fmt.Fprintf(o, "minor: %d\n", *out.Minor)
	}
	if out.Patch != nil {
		fmt.Fprintf(o, "patch: %d\n", *out.Patch)
	}

	if out.PreRelease != nil {
		fmt.Fprintf(o, "prerelease: %s\n", strings.Join(out.PreRelease, "."))
	}
	if out.Build != nil {
		fmt.Fprintf(o, "build: %s\n", strings.Join(out.Build, "."))
	}
}
