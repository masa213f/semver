# semver

`semver` is a command-line tool and a Go library for parsing "[Semantic Versioning 2.0.0][semver-v2]".

## Installation

```console
$ go get -u github.com/masa213f/semver/cmd/semver
```

## Usage

```console
# Parse a version.
# If the parsing is success, `semver` returns exit status "0".

$ semver v1.1.2-rc.0+build | jq .
{
  "major": 1,
  "minor": 1,
  "patch": 2,
  "prerelease": [
    {
      "string": "rc"
    },
    {
      "number": 0
    }
  ],
  "build": [
    "build"
  ]
}
```

```console
# Check if a version is pre-release.
# If the version is pre-release, `semver` returns exit status "0".

$ semver -p v1.1.2-rc.0+build | jq .
{
  "major": 1,
  "minor": 1,
  "patch": 2,
  "prerelease": [
    {
      "string": "rc"
    },
    {
      "number": 0
    }
  ],
  "build": [
    "build"
  ]
}
```

```console
# Check if a version has build meta data.
# If the version build meta data, `semver` returns exit status "0".

$ semver -p v1.1.2-rc.0+build | jq .
{
  "major": 1,
  "minor": 1,
  "patch": 2,
  "prerelease": [
    {
      "string": "rc"
    },
    {
      "number": 0
    }
  ],
  "build": [
    "build"
  ]
}
```

[semver-v2]: https://semver.org/
