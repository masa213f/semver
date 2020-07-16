package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/masa213f/semver"
)

type outputOption struct {
	stripPrefix     bool
	stripPrerelease bool
	stripMetadata   bool
	format          string
}

func newOutputOption() *outputOption {
	return &outputOption{
		stripPrefix:     false,
		stripPrerelease: false,
		stripMetadata:   false,
		format:          "text",
	}
}

func output(o io.Writer, ver *semver.Version, opt *outputOption) {
	out := ver.Copy()
	if opt.stripPrefix {
		out.Prefix = ""
	}
	if opt.stripPrerelease {
		out.Prerelease = nil
	}
	if opt.stripMetadata {
		out.Metadata = nil
	}

	switch opt.format {
	case "text":
		outputText(o, out, opt)
	case "json":
		outputJSON(o, out, opt)
	}
}

func outputText(o io.Writer, ver *semver.Version, opt *outputOption) {
	if ver.Prefix != "" {
		fmt.Fprintf(o, "prefix, %s\n", ver.Prefix)
	}
	fmt.Fprintf(o, "major, %d\n", ver.Major)
	fmt.Fprintf(o, "minor, %d\n", ver.Minor)
	fmt.Fprintf(o, "patch, %d\n", ver.Patch)
	if ver.IsPrerelease() {
		fmt.Fprintf(o, "prerelease, %s\n", strings.Join(ver.Prerelease, ", "))
	}
	if ver.HasMetadata() {
		fmt.Fprintf(o, "metadata, %s\n", strings.Join(ver.Metadata, ", "))
	}
}

func outputJSON(o io.Writer, ver *semver.Version, opt *outputOption) {
	bytes, _ := json.MarshalIndent(ver, "", "  ")
	fmt.Fprintln(o, string(bytes))
}
