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
				PreRelease: []string{"0"},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				PreRelease: []string{"rc"},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				PreRelease: []string{"rc", "0"},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				Build: []string{"foo"},
			},
			isPreRelease: false,
			hasBuildMeta: true,
		},
		{
			input: &Version{
				Build: []string{"foo", "bar"},
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
				PreRelease: []string{"0"},
			},
			isPreRelease: true,
			hasBuildMeta: false,
		},
		{
			input: &Version{
				Version: "0.0.0-0+0",
				Major:   0, Minor: 0, Patch: 0,
				Build: []string{"0"},
			},
			isPreRelease: false,
			hasBuildMeta: true,
		},
		{
			input: &Version{
				Version: "0.0.0-0+0",
				Major:   0, Minor: 0, Patch: 0,
				PreRelease: []string{"0"},
				Build:      []string{"0"},
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
