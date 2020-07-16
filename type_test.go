package semver

import (
	"strconv"
	"testing"
)

func TestVersion(t *testing.T) {
	testcase := []struct {
		input        *Version
		isPreRelease bool
		hasMetadata  bool
		toString     string
	}{
		{
			input:        &Version{},
			isPreRelease: false,
			hasMetadata:  false,
			toString:     "0.0.0",
		},
		{
			input: &Version{
				Prefix: "version",
			},
			isPreRelease: false,
			hasMetadata:  false,
			toString:     "version0.0.0",
		},
		{
			input: &Version{
				Prerelease: []string{"rc"},
			},
			isPreRelease: true,
			hasMetadata:  false,
			toString:     "0.0.0-rc",
		},
		{
			input: &Version{
				Prerelease: []string{"rc", "0"},
			},
			isPreRelease: true,
			hasMetadata:  false,
			toString:     "0.0.0-rc.0",
		},
		{
			input: &Version{
				Metadata: []string{"foo"},
			},
			isPreRelease: false,
			hasMetadata:  true,
			toString:     "0.0.0+foo",
		},
		{
			input: &Version{
				Metadata: []string{"foo", "bar"},
			},
			isPreRelease: false,
			hasMetadata:  true,
			toString:     "0.0.0+foo.bar",
		},
		{
			input: &Version{
				Prefix: "v",
				Major:  1, Minor: 2, Patch: 3,
				Prerelease: []string{"rc.0"},
				Metadata:   []string{"foo", "bar"},
			},
			isPreRelease: true,
			hasMetadata:  true,
			toString:     "v1.2.3-rc.0+foo.bar",
		},
	}
	for no, tc := range testcase {
		t.Run(strconv.Itoa(no), func(t *testing.T) {
			t.Run("IsPreRelease", func(t *testing.T) {
				actual := tc.input.IsPrerelease()
				if actual != tc.isPreRelease {
					t.Errorf("expected=%t, actual=%t", tc.isPreRelease, actual)
					return
				}
			})
			t.Run("HasMetadata", func(t *testing.T) {
				actual := tc.input.HasMetadata()
				if actual != tc.hasMetadata {
					t.Errorf("expected=%t, actual=%t", tc.hasMetadata, actual)
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
