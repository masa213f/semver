package semver

import (
	"testing"
)

func equal(t *testing.T, o1, o2 *Version) bool {
	if o1.Major != o2.Major {
		t.Log("Major", o1.Major, o2.Major)
		return false
	}
	if o1.Minor != o2.Minor {
		t.Log("Minor", o1.Minor, o2.Minor)
		return false
	}
	if o1.Patch != o2.Patch {
		t.Log("Patch", o1.Patch, o2.Patch)
		return false
	}
	if len(o1.PreRelease) != len(o2.PreRelease) {
		t.Log("len(pre)", len(o1.PreRelease), len(o2.PreRelease))
		return false
	}
	for i := range o1.PreRelease {
		if o1.PreRelease[i].Str != o2.PreRelease[i].Str {
			t.Log("pre.Str", i, o1.PreRelease[i].Str, o2.PreRelease[i].Str)
			return false
		}
		if o1.PreRelease[i].Num != o2.PreRelease[i].Num {
			t.Log("pre.Num", i, o1.PreRelease[i].Num, o2.PreRelease[i].Num)
			return false
		}
	}
	if len(o1.Build) != len(o2.Build) {
		t.Log("len(build)", len(o1.Build), len(o2.Build))
		return false
	}
	for i := range o1.Build {
		if o1.Build[i] != o2.Build[i] {
			t.Log("build", i, o1.Build[i], o2.Build[i])
			return false
		}
	}
	return true
}

func TestParseOK(t *testing.T) {
	testcase := []struct {
		input    string
		expected Version
	}{
		{
			input: "0.0.0",
			expected: Version{
				Major: 0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{},
				Build:      []BuildID{},
			},
		},
		{
			input: "1111.2222.3333",
			expected: Version{
				Major: 1111, Minor: 2222, Patch: 3333,
				PreRelease: []PreReleaseID{},
				Build:      []BuildID{},
			},
		},
		{
			input: "999999999999.999999999999.999999999999",
			expected: Version{
				Major: 999999999999, Minor: 999999999999, Patch: 999999999999,
				PreRelease: []PreReleaseID{},
				Build:      []BuildID{},
			},
		},
		{
			input: "0.0.0-0",
			expected: Version{
				Major: 0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{Num: 0}},
				Build:      []BuildID{},
			},
		},
		{
			input: "0.0.0-rc.1",
			expected: Version{
				Major: 0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{Str: "rc"}, {Num: 1}},
				Build:      []BuildID{},
			},
		},
		{
			input: "0.0.0-00.01.10.11.aaa.bbb",
			expected: Version{
				Major: 0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{Str: "00"}, {Str: "01"}, {Num: 10}, {Num: 11}, {Str: "aaa"}, {Str: "bbb"}},
				Build:      []BuildID{},
			},
		},
		{
			input: "0.0.0+0",
			expected: Version{
				Major: 0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{},
				Build:      []BuildID{"0"},
			},
		},
		{
			input: "123.456.789-123.456.789.0ab.cde+123.456.789.0ab.cde",
			expected: Version{
				Major: 123, Minor: 456, Patch: 789,
				PreRelease: []PreReleaseID{{Num: 123}, {Num: 456}, {Num: 789}, {Str: "0ab"}, {Str: "cde"}},
				Build:      []BuildID{"123", "456", "789", "0ab", "cde"},
			},
		},
	}
	for _, tc := range testcase {
		actual, err := Parse(tc.input)
		if err != nil {
			t.Errorf("input: '%s', pase error: %s", tc.input, err.Error())
		}
		if !equal(t, actual, &tc.expected) {
			t.Errorf("input: '%s', expected: %v, actual: %v", tc.input, tc.expected, *actual)
		}
	}
}
