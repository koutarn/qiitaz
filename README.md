<div align="right">

![Relase](https://github.com/sheepla/qiitaz/actions/workflows/release.yml/badge.svg)
![Relase](https://github.com/sheepla/qiitaz/actions/workflows/test.yml/badge.svg)
![Relase](https://github.com/sheepla/qiitaz/actions/workflows/golangci-lint.yml/badge.svg)

<a href="https://github.com/sheepla/qiitaz/releases/latest">

![Latest Release](https://img.shields.io/github/v/release/sheepla/qiitaz?style=flat-square)

</a>

</div>

<div align="center">

# 📝 qiitaz

</div>

<div align="center">

[Qiita](https://qiita.com)の記事を素早く検索し、ターミナル上で閲覧できるコマンドラインツール

</div>

## 使い方

```
Usage:
  qiitaz [OPTIONS] QUERY...

Application Options:
  -V, --version  Show version
  -s, --sort=    Sort key to search e.g. "created", "like", "stock", "rel",
                 (default: "rel")
  -o, --open     Open URL in your web browser
  -p, --preview  Preview page on your terminal
  -n, --pageno=  Max page number of search page (default: 1)
  -j, --json     Output result in JSON format

Help Options:
  -h, --help     Show this help message
```

1. 引数に検索したいキーワードを指定してコマンドを実行します。
1. 検索結果をfuzzyfinderで絞り込みます。`Ctrl-N`, `Ctrl-P` または `Ctrl-J`, `Ctrl-K` でフォーカスを移動します。 `Tab`キーで選択し `Enter` キーで確定します。
1. 選択した記事のURLが出力されます。また、`--open`, `--json`, `--preview` などのオプションを指定することで、選択した記事をブラウザで開いたりターミナル上で閲覧したりすることができます。

### ブラウザで記事のページを開く

`-o`, `--open`オプションを付けるとデフォルトのブラウザが起動し、選択した記事のページが開きます。

### 記事を閲覧する

`-p`, `--preview` オプションを付けると、ターミナルに記事を色付きでレンダリングします。

```bash
qiitaz -p QUERY...
```

lessページャのような感覚でスクロールして、記事を閲覧することができます。

|キー       |説明                                   |
|-----------|---------------------------------------|
|`j` / `k`  |上下スクロール                         |
|`d` / `u`  |半ページスクロール                     |
|`f` / `b`  |1ページスクロール                      |
|`g` / `G`  |ページの先頭へ移動 / ページの末尾へ移動|
|`q` / `Esc`|終了                                   |

### 高度な検索

クエリ引数に次のオプションや演算子を指定することで、条件を詳細に指定して検索することができます。

```
qiitaz title:Go created:\>2022-03-01
```

|オプション               |説明                              |
|-------------------------|----------------------------------|
|`title:{{タイトル}}`     |タイトルにそのキーワードが含まれる|
|`body:{{キーワード}}`          |本文にそのキーワードが含まれる    |
|`code:{{コードの一部}}`  |コードにそのキーワードが含まれる  |
|`tag:{{タグ}}`           |記事に付けられているタグ          |
|`-tag:{{タグ}}`          |除外するタグ                      |
|`user:{{ユーザー名}}`    |ユーザー名                        |
|`stocks:>{{数値}}`       |ストック数                        |
|`created:>{{YYYY-MM-DD}}`|作成日がその日以降                |
|`updated:>{{YYYY-MM-DD}}`|更新日がその日以降                |

**注**: bash, fish, zsh等のシェルでは `>` がファイル上書きのリダイレクトの記号と認識されてしまいます。そのため`\>` のようにエスケープするか、シングルクォートないしはダブルクォートで引数を囲む必要があります。

どちらかの条件にマッチするものを検索したい場合は `OR` 演算子を使います。

|演算子                |説明  |
|----------------------|------|
|`{{条件}} OR {{条件}}`|OR条件|

### ソート条件の変更

`-s`, `--sort` オプションを指定することで、ソート条件を変更することができます。

**例**: 

```
qiitaz -s like Go
```

|値       |説明              |
|---------|------------------|
|`rel`    |関連度順          |
|`like`   |LGTM数の多い順    |
|`stock`  |ストック数の多い順|
|`created`|作成日順          |

### JSON形式で出力

`-j`, `--json` オプションを指定すると、検索結果をJSON形式で出力することができます。

```
qiitaz -j QUERY...
```

## インストール

リリースページから実行可能なバイナリをダウンロードしてください。

> [Latest Release](https://github.com/sheepla/qiitaz/releases/latest)

ソースからビルドする場合は、このリポジトリをクローンして `go install` を実行してください。
`v1.18.1 linux/amd64`にて開発しています。

## ライセンス

[MIT](LICENSE)

## 関連

- [sheepla/fzwiki](https://github.com/sheepla/fzwiki)
- [sheepla/fzenn](https://github.com/sheepla/fzenn)

