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
//   - Version: "1.2.3-rc.0+build.20190928"
//   - Major: 1
//   - Minor: 2
//   - Patch: 3
//   - PreRelease: ["rc", 0]
//   - Build: ["build", "20190928"]
type Version struct {
	Prefix     string   `json:"prefix,omitempty"`
	Major      uint64   `json:"major"`
	Minor      uint64   `json:"minor"`
	Patch      uint64   `json:"patch"`
	PreRelease []string `json:"prerelease,omitempty"`
	Build      []string `json:"build,omitempty"`
}

// IsPreRelease determines whether the Version is a pre-release version.
func (v *Version) IsPreRelease() bool {
	return len(v.PreRelease) != 0
}

// HasBuildMeta determines whether the Version has build metadata.
func (v *Version) HasBuildMeta() bool {
	return len(v.Build) != 0
}

// ToString returns a string representation of the Version.
func (v *Version) ToString() string {
	str := fmt.Sprintf("%s%d.%d.%d", v.Prefix, v.Major, v.Minor, v.Patch)
	if v.IsPreRelease() {
		str += "-" + strings.Join(v.PreRelease, ".")
	}
	if v.HasBuildMeta() {
		str += "+" + strings.Join(v.Build, ".")
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
		PreRelease: v.PreRelease[:],
		Build:      v.Build[:],
	}
	return &ret
}
