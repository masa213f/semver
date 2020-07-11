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
		toString     string
	}{
		{
			input:        &Version{},
			isPreRelease: false,
			hasBuildMeta: false,
			toString:     "0.0.0",
		},
		{
			input: &Version{
				Prefix: "version",
			},
			isPreRelease: false,
			hasBuildMeta: false,
			toString:     "version0.0.0",
		},
		{
			input: &Version{
				PreRelease: []string{"rc"},
			},
			isPreRelease: true,
			hasBuildMeta: false,
			toString:     "0.0.0-rc",
		},
		{
			input: &Version{
				PreRelease: []string{"rc", "0"},
			},
			isPreRelease: true,
			hasBuildMeta: false,
			toString:     "0.0.0-rc.0",
		},
		{
			input: &Version{
				Build: []string{"foo"},
			},
			isPreRelease: false,
			hasBuildMeta: true,
			toString:     "0.0.0+foo",
		},
		{
			input: &Version{
				Build: []string{"foo", "bar"},
			},
			isPreRelease: false,
			hasBuildMeta: true,
			toString:     "0.0.0+foo.bar",
		},
		{
			input: &Version{
				Prefix: "v",
				Major:  1, Minor: 2, Patch: 3,
				PreRelease: []string{"rc.0"},
				Build:      []string{"foo", "bar"},
			},
			isPreRelease: true,
			hasBuildMeta: true,
			toString:     "v1.2.3-rc.0+foo.bar",
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
