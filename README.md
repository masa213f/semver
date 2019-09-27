# semver

`semver` は "[Semantic Versioning 2.0.0][semver-v2]" をパースするコマンドラインツールです。

"Semantic Versioning 2.0.0" をコマンドラインで処理するツールがなかったので作成しました。
CircleCI を使った GitHub プロジェクトの自動リリース等で、 バージョン番号のバリデーションに使用できます。

## インストール

`semver` のインストールには Go が必要です。以下のように `go get` を実行してください。

```console
$ go get -u github.com/masa213f/semver/cmd/semver
```

## 使用方法

`semver` の引数に、バージョン番号(文字列)を指定してください。
バージョン番号のパースが成功すると、以下のように結果が出力されます。
この時、コマンドの終了ステータスとして `0` が返されます。

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

`--json`オプションを指定すると、結果がJSON形式で出力されます。

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

引数に指定された文字列が"Semantic Versioning 2.0.0"に従っていない場合は、終了ステータスとして`1`が返されます。

```console
$ semver v1.12
parse error: format error
# => exit status: 1

$ semver v1.01.0
parse error: invalid numeric identifier (leading zeros): minor = 01
# => exit status: 1
```

また、`-p`(`--is-prerelease`) オプションを指定すると、プレリリースバージョンの判定ができます。
プレリリースバージョンの場合は終了ステータス`0`、そうでない場合は終了ステータス`2`が返されます。

```console
$ semver -p v1.1.2-rc.0
prerelease: rc.0
# => exit status: 0

$ semver -p v1.1.2
official version: version = 1.1.2
# => exit status: 2
```

[semver-v2]: https://semver.org/
