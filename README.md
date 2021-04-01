Download file in Golang
==================

[![Build Status](https://github.com/Imputes/fdlr/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Imputes/fdlr/actions)

## Featues:
* skip TLS checking
* download batches of files concurrently
* resume incomplete downloads

## Requires
Go v1.6+

## USAGE
```
fdlr
file downloader written in Go

Usage:
fdlr [command]

Available Commands:
  download    downloads a file from URL or file name
  help        Help about any command
  resume      resume downloading task
  task        show current downloading task
  version     prints meta info

Flags:
  -h, --help   help for ./bin/fdlr

Use "./bin/fdlr [command] --help" for more information about a command.
```

## Inspired
- [https://github.com/ytdl-org/youtube-dl](https://github.com/ytdl-org/youtube-dl)
- [https://github.com/cavaliercoder/grab](https://github.com/cavaliercoder/grab)
