<div align="right">

![CI](https://github.com/sheepla/fzwiki/actions/workflows/ci.yml/badge.svg)
![Relase](https://github.com/sheepla/fzwiki/actions/workflows/release.yml/badge.svg)

<a href="https://github.com/sheepla/qiitaz/releases/latest">

![Latest Release](https://img.shields.io/github/v/release/sheepla/qiitaz?style=flat-square)

</a>

</div>

<div align="center">

# qiitaz

</div>

<div align="center">

A command line [Qiita](https://qiita.com) searcher with fuzzyfinder UI

</div>

## Usage

```
Usage:
  qiitaz [OPTIONS] QUERY...

Application Options:
  -V, --version  Show version
  -s, --sort=    Sort key to search e.g. "created", "like", "stock", "rel",
                 (default: "rel")
  -o, --open     Open URL in your web browser

Help Options:
  -h, --help     Show this help message
```

## Installation

### Build from Source

Clone this repository then run `go install`.
Requires Go, testing on `v1.17.8 linux/amd64`

### Download Executable Binary

> [Latest Release](https://github.com/sheepla/qiitaz/releases/latest)

## LICENSE

[MIT](./LICENSE)

## Related Projects

- [sheepla/fzwiki](https://github.com/sheepla/fzwiki)
- [sheepla/fzenn](https://github.com/sheepla/fzenn)

