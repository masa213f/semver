package semver

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	digitRegexp        = "0|[1-9][0-9]{0,15}"
	alnumRegexp        = "[0-9A-Za-z-]{1,16}"
	versionCoreRegexp  = "(" + digitRegexp + ")\\.(" + digitRegexp + ")\\.(" + digitRegexp + ")"
	preReleaseRegexp   = "(" + alnumRegexp + "(\\." + alnumRegexp + "){0,5})"
	buildRegexp        = "(" + alnumRegexp + "(\\." + alnumRegexp + "){0,5})"
	validVersionRegexp = "^" + versionCoreRegexp + "(-" + preReleaseRegexp + ")?" + "(\\+" + buildRegexp + ")?$"
	validNumericRegexp = "^(" + digitRegexp + ")$"

	// <Mmajor:Max16>.<Minor:Max16>.<Patch:Max16>-<PreRelease1:Max16>.(snip).<PreRelease6:Max16>+<Build1:Max16>.(snip).<Build6:Max16>
	maxInputLength = 254
)

var validVersion = regexp.MustCompile(validVersionRegexp)
var validNumeric = regexp.MustCompile(validNumericRegexp)

// Version is xxx.
type Version struct {
	Major      uint64         `json:"major"`
	Minor      uint64         `json:"minor"`
	Patch      uint64         `json:"patch"`
	PreRelease []PreReleaseID `json:"prerelease,omitempty"`
	Build      []BuildID      `json:"build,omitempty"`
}

// IsPreRelease is xxx.
func (v *Version) IsPreRelease() bool {
	if len(v.PreRelease) == 0 {
		return false
	}
	return true
}

// PreReleaseID is xxx.
type PreReleaseID struct {
	String string
	Number uint64
}

type preReleaseIDString struct {
	String string `json:"string"`
}

type preReleaseIDNumber struct {
	Number uint64 `json:"number"`
}

// IsNumber is xxx.
func (pr *PreReleaseID) IsNumber() bool {
	if pr.String == "" {
		return true
	}
	return false
}

// MarshalJSON is xxx.
func (pr PreReleaseID) MarshalJSON() ([]byte, error) {
	if pr.IsNumber() {
		tmp := preReleaseIDNumber{Number: pr.Number}
		return json.Marshal(tmp)
	}
	tmp := preReleaseIDString{String: pr.String}
	return json.Marshal(tmp)
}

// BuildID is xxx.
type BuildID string

func parseUint(str string) uint64 {
	num, _ := strconv.ParseUint(str, 10, 64)
	return num
}

func isValidNumber(str string) bool {
	return validNumeric.MatchString(str)
}

func parsePreRelease(str string) []PreReleaseID {
	if str == "" {
		return make([]PreReleaseID, 0)
	}
	tmp := strings.Split(str, ".")
	ids := make([]PreReleaseID, len(tmp))
	for i, str := range tmp {
		if isValidNumber(str) {
			ids[i].Number = parseUint(str)
		} else {
			ids[i].String = str
		}
	}
	return ids
}

func parseBuild(str string) []BuildID {
	if str == "" {
		return make([]BuildID, 0)
	}
	tmp := strings.Split(str, ".")
	ids := make([]BuildID, len(tmp))
	for i, id := range tmp {
		ids[i] = BuildID(id)
	}
	return ids
}

// Parse is xxx.
func Parse(str string) (*Version, error) {
	if len(str) > maxInputLength {
		return nil, fmt.Errorf("max length")
	}
	submatch := validVersion.FindStringSubmatch(str)
	if len(submatch) == 0 {
		return nil, fmt.Errorf("parse error")
	}
	// for i, v := range submatch {
	// 	fmt.Println(">", i, v)
	// }
	ver := &Version{}
	ver.Major = parseUint(submatch[1])
	ver.Minor = parseUint(submatch[2])
	ver.Patch = parseUint(submatch[3])
	ver.PreRelease = parsePreRelease(submatch[5])
	ver.Build = parseBuild(submatch[8])
	return ver, nil
}
