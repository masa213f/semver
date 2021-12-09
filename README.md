# semver

`semver` is a command-line tool to parse "[Semantic Versioning 2.0.0][semver-v2]".

## Install

Go is required to install `semver`. Please run `go get` as follows.

```console
$ go install github.com/masa213f/semver/cmd/semver@v0.1.0
```

## Usage

Specify the version number (string) as the argument of `semver`.
If the version number is parsed successfully, this command outputs the result as follows.
At this time, `0` is returned as the exit status.

```console
$ semver v1.2.3-rc.0+build.20190925
prefix: v
version: 1.2.3-rc.0+build.20190925
major: 1
minor: 2
patch: 3
prerelease: rc.0
build: build.20190925
# => exit status: 0
```

If you specify the `--json` option, the result will be output in JSON format.

```console
$ semver v1.2.3-rc.0+build.20190925 --json
{
  "prefix": "v",
  "version": "1.2.3-rc.0+build.20190925",
  "major": 1,
  "minor": 2,
  "patch": 3,
  "prerelease": [
    "rc",
    "0"
  ],
  "build": [
    "build",
    "20190925"
  ]
}
# => exit status: 0
```

If the argument does not follow "Semantic Versioning 2.0.0", `1` will be returned as the exit status.

```console
$ semver v1.12
parse error: format error
# => exit status: 1

$ semver v1.01.0
parse error: invalid numeric identifier (leading zeros): minor = 01
# => exit status: 1
```

You can also specify the `-p`(`--is-prerelease`) option to determine the pre-release version.

```console
$ semver -p v1.1.2-rc.0
prerelease: rc.0
# => exit status: 0

$ semver -p v1.1.2
official version: version = 1.1.2
# => exit status: 2
```

[semver-v2]: https://semver.org/
