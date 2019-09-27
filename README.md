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

`semver` の引数に、パースしたいバージョン番号(文字列)を指定してください。
パースが成功すると、以下のように結果が出力されます。

```console
$ semver v1.1.2-rc.0+build
prefix: v
version: 1.1.2-rc.0+build
major: 1
minor: 1
patch: 2
prerelease: rc.0
build: build
```

また、`-json`オプションを指定すると、結果が JSON で出力されます。
(以下は jq コマンドを使用して出力を整形した例です。)

```console
semver -json v1.1.2-rc.0+build | jq .
{
  "prefix": "v",
  "version": "1.1.2-rc.0+build",
  "major": 1,
  "minor": 1,
  "patch": 2,
  "prerelease": [
    "rc",
    "0"
  ],
  "build": [
    "build"
  ]
}
```

バージョン番号のパースが成功すると、コマンドの終了ステータスとして `0` が返され、パースが失敗すると `1` を返します。

また、プレリリースバージョンの判定には、`-p` オプションを指定します。

```console
$ semver -p v1.1.2-rc.0+build
```

[semver-v2]: https://semver.org/
