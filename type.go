package semver

import (
	"encoding/json"
	"strconv"
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
	Prefix     string         `json:"prefix"`
	Version    string         `json:"version"`
	Major      uint64         `json:"major"`
	Minor      uint64         `json:"minor"`
	Patch      uint64         `json:"patch"`
	PreRelease []PreReleaseID `json:"prerelease,omitempty"`
	Build      []BuildID      `json:"build,omitempty"`
}

// ToString returns a string representation of the Version.
func (v *Version) String() string {
	return v.Version
}

// IsPreRelease determines whether the Version is a pre-release version.
func (v *Version) IsPreRelease() bool {
	if len(v.PreRelease) == 0 {
		return false
	}
	return true
}

// HasBuildMeta determines whether the Version has build metadata.
func (v *Version) HasBuildMeta() bool {
	if len(v.Build) == 0 {
		return false
	}
	return true
}

// PreReleaseID represents "dot-separated pre-release identifier".
// "dot-separated pre-release identifier" can be either a string type or a numeric type.
// Since the comparison method differs depending on the type, it makes it possible to distinguish the type.
type PreReleaseID struct {
	String string
	Number uint64
}

// IsNumber determines the PreReleaseID type.
func (pr *PreReleaseID) IsNumber() bool {
	if pr.String == "" {
		return true
	}
	return false
}

// ToString returns a string representation of the PreReleaseID.
func (pr *PreReleaseID) ToString() string {
	if pr.IsNumber() {
		return strconv.FormatUint(pr.Number, 10)
	}
	return pr.String
}

// preReleaseIDString is just used for MarshalJSON().
type preReleaseIDString struct {
	String string `json:"string"`
}

// preReleaseIDNumber is just used for MarshalJSON().
type preReleaseIDNumber struct {
	Number uint64 `json:"number"`
}

// MarshalJSON is custom machaler of PreReleaseID.
// It outputs either a string type value or a numeric type value.
func (pr PreReleaseID) MarshalJSON() ([]byte, error) {
	if pr.IsNumber() {
		tmp := preReleaseIDNumber{Number: pr.Number}
		return json.Marshal(tmp)
	}
	tmp := preReleaseIDString{String: pr.String}
	return json.Marshal(tmp)
}

// BuildID represents "dot-separated build identifiers".
type BuildID string
