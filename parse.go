package semver

import (
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
	Major      uint64
	Minor      uint64
	Patch      uint64
	PreRelease []PreReleaseID
	Build      []BuildID
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
	Str string
	Num uint64
}

// BuildID is xxx.
type BuildID string

// IsNum is xxx.
func (pr *PreReleaseID) IsNum() bool {
	if pr.Str != "" {
		return false
	}
	return true
}

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
			ids[i].Num = parseUint(str)
		} else {
			ids[i].Str = str
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
