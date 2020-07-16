package semver

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

// Version represents components of the "Semantic Versioning 2.0.0" <https://semver.org/>.
// memo: "Prefix" is not defined in the specification. It is defined to improve convenience of this program.
//
// If the original version number is "v1.2.3-rc.0+build.20190928", the following values are stored.
//   - Prefix: "v"
//   - Major: 1
//   - Minor: 2
//   - Patch: 3
//   - Prerelease: ["rc", "0"]
//   - Metadata: ["build", "20190928"]
type Version struct {
	Prefix     string     `json:"prefix,omitempty"`
	Major      uint64     `json:"major"`
	Minor      uint64     `json:"minor"`
	Patch      uint64     `json:"patch"`
	Prerelease Prerelease `json:"prerelease,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
}

// Prerelease represents prerelease parts of a version.
type Prerelease []string

func (p Prerelease) String() string {
	return strings.Join(p, ".")
}

// Metadata represents metadata parts of a version.
type Metadata []string

func (m Metadata) String() string {
	return strings.Join(m, ".")
}

// IsPrerelease determines whether the Version is a prerelease version.
func (v *Version) IsPrerelease() bool {
	return len(v.Prerelease) != 0
}

// HasMetadata determines whether the Version has metadata.
func (v *Version) HasMetadata() bool {
	return len(v.Metadata) != 0
}

// String returns a string representation of the Version.
func (v *Version) String() string {
	str := fmt.Sprintf("%s%d.%d.%d", v.Prefix, v.Major, v.Minor, v.Patch)
	if v.IsPrerelease() {
		str += "-" + v.Prerelease.String()
	}
	if v.HasMetadata() {
		str += "+" + v.Metadata.String()
	}
	return str
}

// Format returns a string representation of the Version.
func (v *Version) Format(tmpl string) string {
	out := new(bytes.Buffer)
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	if err = t.Execute(out, v); err != nil {
		log.Fatal(err)
	}
	return out.String()
}

// Copy returns a copy of the Version.
func (v *Version) Copy() *Version {
	ret := Version{
		Prefix:     v.Prefix,
		Major:      v.Major,
		Minor:      v.Minor,
		Patch:      v.Patch,
		Prerelease: v.Prerelease[:],
		Metadata:   v.Metadata[:],
	}
	return &ret
}
