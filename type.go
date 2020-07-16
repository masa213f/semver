package semver

import (
	"fmt"
	"strings"
)

// Version represents components of the "Semantic Versioning 2.0.0"<https://semver.org/>.
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
	Prefix     string   `json:"prefix,omitempty"`
	Major      uint64   `json:"major"`
	Minor      uint64   `json:"minor"`
	Patch      uint64   `json:"patch"`
	Prerelease []string `json:"prerelease,omitempty"`
	Metadata   []string `json:"metadata,omitempty"`
}

// IsPrerelease determines whether the Version is a prerelease version.
func (v *Version) IsPrerelease() bool {
	return len(v.Prerelease) != 0
}

// HasMetadata determines whether the Version has metadata.
func (v *Version) HasMetadata() bool {
	return len(v.Metadata) != 0
}

// ToString returns a string representation of the Version.
func (v *Version) ToString() string {
	str := fmt.Sprintf("%s%d.%d.%d", v.Prefix, v.Major, v.Minor, v.Patch)
	if v.IsPrerelease() {
		str += "-" + strings.Join(v.Prerelease, ".")
	}
	if v.HasMetadata() {
		str += "+" + strings.Join(v.Metadata, ".")
	}
	return str
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
