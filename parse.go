package semver

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	digitRegexp        = "[0-9]{1,16}"
	alnumRegexp        = "[0-9A-Za-z-]{1,16}"
	prefixRegexp       = "[^0-9]{0,16}"
	versionCoreRegexp  = "(" + digitRegexp + ")\\.(" + digitRegexp + ")\\.(" + digitRegexp + ")"
	preReleaseRegexp   = "(" + alnumRegexp + "(\\." + alnumRegexp + "){0,5})"
	buildRegexp        = "(" + alnumRegexp + "(\\." + alnumRegexp + "){0,5})"
	validVersionRegexp = "^(" + prefixRegexp + ")(" + versionCoreRegexp + "(-" + preReleaseRegexp + ")?" + "(\\+" + buildRegexp + ")?)$"
	validNumericRegexp = "^(" + digitRegexp + ")$"

	// <Prefix:Max16><Mmajor:Max16>.<Minor:Max16>.<Patch:Max16>-<PreRelease1:Max16>.(snip).<PreRelease6:Max16>+<Build1:Max16>.(snip).<Build6:Max16>
	maxInputLength = 270
)

var validVersion = regexp.MustCompile(validVersionRegexp)
var validNumeric = regexp.MustCompile(validNumericRegexp)

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

// IsPreRelease determines whether the Version is a pre-release version.
func (v *Version) IsPreRelease() bool {
	return len(v.PreRelease) != 0
}

// HasBuildMeta determines whether the Version has build metadata.
func (v *Version) HasBuildMeta() bool {
	return len(v.Build) != 0
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
	return pr.String == ""
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

func parseUint(str string) uint64 {
	num, _ := strconv.ParseUint(str, 10, 64)
	return num
}

const (
	stringIdentifier = iota
	numberIdentifier
	invalidIdentifier
)

func identifierType(str string) int {
	if !validNumeric.MatchString(str) {
		// string (alphanumeric identifier)
		return stringIdentifier
	} else if str == "0" || str[0] != '0' {
		// number (numeric identifier)
		return numberIdentifier
	}
	// invalid string: Numeric identifiers MUST NOT include leading zeroes.
	return invalidIdentifier
}

func parseCoreVersion(part, str string) (uint64, error) {
	switch identifierType(str) {
	case numberIdentifier:
		return parseUint(str), nil
	case invalidIdentifier:
		return 0, newInvalidNumericError(part, str)
	}
	return 0, errors.New("unexpected error")
}

func parsePreRelease(str string) ([]PreReleaseID, error) {
	if str == "" {
		return nil, nil
	}
	tmp := strings.Split(str, ".")
	ids := make([]PreReleaseID, len(tmp))
	for i, str := range tmp {
		switch identifierType(str) {
		case stringIdentifier:
			ids[i].String = str
		case numberIdentifier:
			ids[i].Number = parseUint(str)
		case invalidIdentifier:
			return nil, newInvalidNumericError("prerelease["+strconv.Itoa(i)+"]", str)
		}
	}
	return ids, nil
}

func parseBuild(str string) []BuildID {
	if str == "" {
		return nil
	}
	tmp := strings.Split(str, ".")
	ids := make([]BuildID, len(tmp))
	for i, id := range tmp {
		ids[i] = BuildID(id)
	}
	return ids
}

// Parse parses "Semantic Versioning 2.0.0".
func Parse(str string) (*Version, error) {
	if len(str) > maxInputLength {
		return nil, newParseError("too long input")
	}
	submatch := validVersion.FindStringSubmatch(str)
	if len(submatch) == 0 {
		return nil, newParseError("format error")
	}

	ver := &Version{}
	ver.Prefix = submatch[1]
	ver.Version = submatch[2]

	major, err := parseCoreVersion("major", submatch[3])
	if err != nil {
		return nil, newParseError(err.Error())
	}
	ver.Major = major

	minor, err := parseCoreVersion("minor", submatch[4])
	if err != nil {
		return nil, newParseError(err.Error())
	}
	ver.Minor = minor

	patch, err := parseCoreVersion("patch", submatch[5])
	if err != nil {
		return nil, newParseError(err.Error())
	}
	ver.Patch = patch

	pre, err := parsePreRelease(submatch[7])
	if err != nil {
		return nil, newParseError(err.Error())
	}
	ver.PreRelease = pre

	ver.Build = parseBuild(submatch[10])
	return ver, nil
}
