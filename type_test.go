package semver

import (
	"strconv"
	"testing"
)

func TestVersion(t *testing.T) {
	testcase := []struct {
		input        *Version
		isPrerelease bool
		hasMetadata  bool
		toString     string
	}{
		{
			input:        &Version{},
			isPrerelease: false,
			hasMetadata:  false,
			toString:     "0.0.0",
		},
		{
			input: &Version{
				Prefix: "version",
			},
			isPrerelease: false,
			hasMetadata:  false,
			toString:     "version0.0.0",
		},
		{
			input: &Version{
				Prerelease: []string{"rc"},
			},
			isPrerelease: true,
			hasMetadata:  false,
			toString:     "0.0.0-rc",
		},
		{
			input: &Version{
				Prerelease: []string{"rc", "0"},
			},
			isPrerelease: true,
			hasMetadata:  false,
			toString:     "0.0.0-rc.0",
		},
		{
			input: &Version{
				Metadata: []string{"foo"},
			},
			isPrerelease: false,
			hasMetadata:  true,
			toString:     "0.0.0+foo",
		},
		{
			input: &Version{
				Metadata: []string{"foo", "bar"},
			},
			isPrerelease: false,
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
			isPrerelease: true,
			hasMetadata:  true,
			toString:     "v1.2.3-rc.0+foo.bar",
		},
	}
	for no, tc := range testcase {
		t.Run(strconv.Itoa(no), func(t *testing.T) {
			t.Run("IsPrerelease", func(t *testing.T) {
				actual := tc.input.IsPrerelease()
				if actual != tc.isPrerelease {
					t.Errorf("expected=%t, actual=%t", tc.isPrerelease, actual)
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
