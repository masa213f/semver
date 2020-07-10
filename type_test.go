package semver

import (
	"strconv"
	"testing"
)

func TestVersion(t *testing.T) {
	testcase := []struct {
		input        *Version
		isPreRelease bool
		hasBuildMeta bool
	}{
		{
			input: &Version{
				PreRelease: []PreReleaseID{{Number: 0}},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				PreRelease: []PreReleaseID{{String: "rc"}},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				PreRelease: []PreReleaseID{{String: "rc"}, {Number: 0}},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				Build: []BuildID{"foo"},
			},
			isPreRelease: false,
			hasBuildMeta: true,
		},
		{
			input: &Version{
				Build: []BuildID{"foo", "bar"},
			},
			isPreRelease: false,
			hasBuildMeta: true,
		},
		{
			input:        &Version{},
			isPreRelease: false,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				Version: "0.0.0-0+0",
				Major:   0, Minor: 0, Patch: 0,
			},
			isPreRelease: false,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				Version: "0.0.0-0+0",
				Major:   0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{Number: 0}},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				Version: "0.0.0-0+0",
				Major:   0, Minor: 0, Patch: 0,
				Build: []BuildID{"0"},
			},
			isPreRelease: false,
			hasBuildMeta: true,
		},
		{
			input: &Version{
				Version: "0.0.0-0+0",
				Major:   0, Minor: 0, Patch: 0,
				PreRelease: []PreReleaseID{{Number: 0}},
				Build:      []BuildID{"0"},
			},
			isPreRelease: true,
			hasBuildMeta: true,
		},
	}
	for no, tc := range testcase {
		t.Run(strconv.Itoa(no), func(t *testing.T) {
			t.Run("IsPreRelease", func(t *testing.T) {
				actual := tc.input.IsPreRelease()
				if actual != tc.isPreRelease {
					t.Errorf("expected=%t, actual=%t", tc.isPreRelease, actual)
					return
				}
			})
			t.Run("HasBuildMeta", func(t *testing.T) {
				actual := tc.input.HasBuildMeta()
				if actual != tc.hasBuildMeta {
					t.Errorf("expected=%t, actual=%t", tc.hasBuildMeta, actual)
					return
				}
			})
		})

	}
}

func TestPreReleaseID(t *testing.T) {
	testcase := []struct {
		input    *PreReleaseID
		isNumber bool
		toString string
	}{
		{
			input:    &PreReleaseID{},
			isNumber: true,
			toString: "0",
		},
		{
			input: &PreReleaseID{
				Number: 1,
			},
			isNumber: true,
			toString: "1",
		},
		{
			input: &PreReleaseID{
				String: "1",
			},
			isNumber: false,
			toString: "1",
		},
		{
			// invalid case
			input: &PreReleaseID{
				Number: 1,
				String: "1",
			},
			isNumber: false,
			toString: "1",
		},
	}
	for no, tc := range testcase {
		t.Run(strconv.Itoa(no), func(t *testing.T) {
			t.Run("IsNumber", func(t *testing.T) {
				actual := tc.input.IsNumber()
				if actual != tc.isNumber {
					t.Errorf("expected=%t, actual=%t", tc.isNumber, actual)
					return
				}
			})
			t.Run("ToString", func(t *testing.T) {
				actual := tc.input.ToString()
				if actual != tc.toString {
					t.Errorf("expected=%s, actual=%s", tc.toString, actual)
					return
				}
			})
		})
	}
}
