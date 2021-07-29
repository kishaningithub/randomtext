# Randomtext

[![Build Status](https://github.com/kishaningithub/kafka-perf/actions/workflows/build.yml/badge.svg)](https://github.com/kishaningithub/randomtext/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/kishaningithub/randomtext.svg)](https://pkg.go.dev/github.com/kishaningithub/randomtext)
[![Go Report Card](https://goreportcard.com/badge/github.com/kishaningithub/randomtext)](https://goreportcard.com/report/github.com/kishaningithub/randomtext)
[![Downloads](https://img.shields.io/github/downloads/kishaningithub/randomtext/latest/total.svg)](https://github.com/kishaningithub/randomtext/releases)
[![Latest release](https://img.shields.io/github/release/kishaningithub/randomtext.svg)](https://github.com/kishaningithub/randomtext/releases)

Command line random text generator

## Usage

```bash
$ randomtext -h
Usage of randomtext:
  -size string
        Size of generated random text in KB, MB, GB, TB (default "1MB")
  -type string
        Type of text to be generated - chars, words, zeros (default "chars")
```

### Examples

- Generate 1 KB random characters
  - `randomtext --size=1KB | pv > output.txt`
- Generate 100MB random words
  - `randomtext --size=100MB --type=words | pv > output.txt`
- Generate 1GB of zeros
  - `randomtext --size=1GB --type=zeros | pv > output.txt`

Source of [words](https://github.com/dwyl/english-words). The above commands use [pv](https://www.ivarch.com/programs/pv.shtml) for pipeline visualization only.

## Use cases

- Testing the capacity of code editors/IDEs
- Generate logs to generate load for log forwarders like [fluent bit](https://github.com/fluent/fluent-bit)

## Installation

## Mac OSX

```shell
$ brew install kishaningithub/tap/randomtext
```

## Other platforms
Prebuilt binaries for Intel 64-bit architecture are available for

- [Linux](https://github.com/kishaningithub/randomtext/releases/download/v1.0.0/randomtext_1.0.0_linux_amd64.tar.gz)
- [Mac OS](https://github.com/kishaningithub/randomtext/releases/download/v1.0.0/randomtext_1.0.0_darwin_amd64.tar.gz)
- [Windows](https://github.com/kishaningithub/randomtext/releases/download/v1.0.0/randomtext_1.0.0_windows_amd64.tar.gz)

Extract and add the binary in your `$PATH`

## Development

To build

```bash
glide install
go install -v ./vendor/github.com/jteeuwen/go-bindata/...
go generate -x .
go install -v ./...
randomtext
```