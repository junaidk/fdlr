Download file in Golang
==================

[![Build Status](https://github.com/Imputes/fdlr/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Imputes/fdlr/actions)

## Featues:
* Support HTTP and HTTPS
* you can set the number of parallel to download
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

```
fdlr download -h
downloads a file from URL or file name

Usage:
  ./bin/fdlr download [flags]

Examples:
fdlr download [-c=goroutines_count] URL

Flags:
  -c, --goroutines count int   default is your CPU threads count (default 4)
  -h, --help                   help for download
```

```
fdlr resume -h
resume downloading task

Usage:
  ./bin/fdlr resume [flags]

Examples:
fdlr resume URL or file name

Flags:
  -h, --help   help for resume
```

## Inspired
- [https://github.com/ytdl-org/youtube-dl](https://github.com/ytdl-org/youtube-dl)
- [https://github.com/cavaliercoder/grab](https://github.com/cavaliercoder/grab)
