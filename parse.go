package semver

import (
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
	prereleaseRegexp   = "(" + alnumRegexp + "(\\." + alnumRegexp + "){0,5})"
	metadataRegexp     = "(" + alnumRegexp + "(\\." + alnumRegexp + "){0,5})"
	validVersionRegexp = "^(" + prefixRegexp + ")(" + versionCoreRegexp + "(-" + prereleaseRegexp + ")?" + "(\\+" + metadataRegexp + ")?)$"
	validNumericRegexp = "^(" + digitRegexp + ")$"

	// <Prefix:Max16><Mmajor:Max16>.<Minor:Max16>.<Patch:Max16>-<Prerelease[0]:Max16>.(snip).<Prerelease[5]:Max16>+<Metadata[0]:Max16>.(snip).<Metadata[5]:Max16>
	maxInputLength = 270
)

var validVersion = regexp.MustCompile(validVersionRegexp)
var validNumeric = regexp.MustCompile(validNumericRegexp)

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

func parsePrerelease(str string) ([]string, error) {
	if str == "" {
		return nil, nil
	}
	ret := strings.Split(str, ".")
	for i, str := range ret {
		if identifierType(str) == invalidIdentifier {
			return nil, newInvalidNumericError("prerelease["+strconv.Itoa(i)+"]", str)
		}
	}
	return ret, nil
}

func parseMetadata(str string) []string {
	if str == "" {
		return nil
	}
	return strings.Split(str, ".")
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

	pre, err := parsePrerelease(submatch[7])
	if err != nil {
		return nil, newParseError(err.Error())
	}
	ver.Prerelease = pre

	ver.Metadata = parseMetadata(submatch[10])
	return ver, nil
}
