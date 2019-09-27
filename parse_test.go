package semver

import (
	"reflect"
	"testing"
)

func cmpValue(o1, o2 *Version) bool {
	if o1 == nil && o2 == nil {
		return true
	}

	if o1.Prefix != o2.Prefix {
		return false
	}
	if o1.Version != o2.Version {
		return false
	}
	if o1.Major != o2.Major {
		return false
	}
	if o1.Minor != o2.Minor {
		return false
	}
	if o1.Patch != o2.Patch {
		return false
	}
	if len(o1.PreRelease) != len(o2.PreRelease) {
		return false
	}
	for i := range o1.PreRelease {
		if o1.PreRelease[i].String != o2.PreRelease[i].String {
			return false
		}
		if o1.PreRelease[i].Number != o2.PreRelease[i].Number {
			return false
		}
	}
	if len(o1.Build) != len(o2.Build) {
		return false
	}
	for i := range o1.Build {
		if o1.Build[i] != o2.Build[i] {
			return false
		}
	}
	return true
}

func cmpError(e1, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	t1 := reflect.TypeOf(e1)
	t2 := reflect.TypeOf(e2)

	if t1 == t2 {
		return true
	}
	return false
}

func TestParse(t *testing.T) {
	testcase := []struct {
		input         string
		expectedValue *Version
		expectedError error
	}{
		{
			input: "0.0.0",
			expectedValue: &Version{
				Version: "0.0.0",
				Major:   0, Minor: 0, Patch: 0,
			},
		},
		{
			input: "0.0.0-0",
			expectedValue: &Version{
				Version: "0.0.0-0",
				Major:   0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{Number: 0}},
			},
		},
		{
			input: "0.0.0--",
			expectedValue: &Version{
				Version: "0.0.0--",
				Major:   0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{String: "-"}},
			},
		},
		{
			input: "0.0.0+0",
			expectedValue: &Version{
				Version: "0.0.0+0",
				Major:   0, Minor: 0, Patch: 0,
				Build: []BuildID{"0"},
			},
		},
		{
			input: "0.0.0+0-0",
			expectedValue: &Version{
				Version: "0.0.0+0-0",
				Major:   0, Minor: 0, Patch: 0,
				Build: []BuildID{"0-0"},
			},
		},
		{
			input: "0.0.0-0+0",
			expectedValue: &Version{
				Version: "0.0.0-0+0",
				Major:   0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{Number: 0}},
				Build:      []BuildID{"0"},
			},
		},
		{
			input: "ver123.456.789-1234.5678.90ab.ceef+1234.5678.90ab.ceef",
			expectedValue: &Version{
				Prefix: "ver", Version: "123.456.789-1234.5678.90ab.ceef+1234.5678.90ab.ceef",
				Major: 123, Minor: 456, Patch: 789,
				PreRelease: []PreReleaseID{{Number: 1234}, {Number: 5678}, {String: "90ab"}, {String: "ceef"}},
				Build:      []BuildID{"1234", "5678", "90ab", "ceef"},
			},
		},
		{
			input: "v1.2.15-rc.1+build20190907",
			expectedValue: &Version{
				Prefix: "v", Version: "1.2.15-rc.1+build20190907",
				Major: 1, Minor: 2, Patch: 15,
				PreRelease: []PreReleaseID{{String: "rc"}, {Number: 1}},
				Build:      []BuildID{"build20190907"},
			},
		},
		{
			input: "pppppppppppppppp9999999999999999.9999999999999999.9999999999999999-1111111111111111.2222222222222222.3333333333333333.aaaaaaaaaaaaaaaa.bbbbbbbbbbbbbbbb.cccccccccccccccc+1111111111111111.2222222222222222.3333333333333333.aaaaaaaaaaaaaaaa.bbbbbbbbbbbbbbbb.cccccccccccccccc",
			expectedValue: &Version{
				Prefix:  "pppppppppppppppp",
				Version: "9999999999999999.9999999999999999.9999999999999999-1111111111111111.2222222222222222.3333333333333333.aaaaaaaaaaaaaaaa.bbbbbbbbbbbbbbbb.cccccccccccccccc+1111111111111111.2222222222222222.3333333333333333.aaaaaaaaaaaaaaaa.bbbbbbbbbbbbbbbb.cccccccccccccccc",
				Major:   9999999999999999, Minor: 9999999999999999, Patch: 9999999999999999,
				PreRelease: []PreReleaseID{
					{Number: 1111111111111111}, {Number: 2222222222222222}, {Number: 3333333333333333}, {String: "aaaaaaaaaaaaaaaa"}, {String: "bbbbbbbbbbbbbbbb"}, {String: "cccccccccccccccc"},
				},
				Build: []BuildID{
					"1111111111111111", "2222222222222222", "3333333333333333", "aaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbb", "cccccccccccccccc",
				},
			},
		},

		// invalid case
		{
			input:         "",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.0+",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.0-",
			expectedError: newParseError(""),
		},
		{
			input:         "12345678901234567.0.0",
			expectedError: newParseError(""),
		},
		{
			input:         "0.12345678901234567.0",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.12345678901234567",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.0-12345678901234567",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.0-1.2.3.4.5.6.7",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.0+12345678901234567",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.0+1.2.3.4.5.6.7",
			expectedError: newParseError(""),
		},
		{
			input:         "01.0.0",
			expectedError: newParseError(""),
		},
		{
			input:         "0.01.0",
			expectedError: newParseError(""),
		},
		{
			input:         "0.0.01",
			expectedError: newParseError(""),
		},
		{
			input:         "12345.67890.12345-rc.00",
			expectedError: newParseError(""),
		},
		{
			input:         "12345.67890.12345-rc.12345.00+build.0000",
			expectedError: newParseError(""),
		},
		{
			input:         "xpppppppppppppppp9999999999999999.9999999999999999.9999999999999999-1111111111111111.2222222222222222.3333333333333333.aaaaaaaaaaaaaaaa.bbbbbbbbbbbbbbbb.cccccccccccccccc+1111111111111111.2222222222222222.3333333333333333.aaaaaaaaaaaaaaaa.bbbbbbbbbbbbbbbb.cccccccccccccccc",
			expectedError: newParseError(""),
		},
	}
	for _, tc := range testcase {
		t.Run(tc.input, func(t *testing.T) {
			actualValue, actualError := Parse(tc.input)
			if !cmpValue(actualValue, tc.expectedValue) {
				t.Errorf("expected=%v, actual=%v", *tc.expectedValue, *actualValue)
				return
			}
			if !cmpError(actualError, tc.expectedError) {
				t.Errorf("error='%v'", actualError)
				return
			}
		})
	}
}
