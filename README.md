# Randomtext

[![Build Status](https://travis-ci.org/kishaningithub/randomtext.svg?branch=master)](https://travis-ci.org/kishaningithub/randomtext)
[![Go Doc](https://godoc.org/github.com/kishaningithub/randomtext?status.svg)](https://godoc.org/github.com/kishaningithub/randomtext)
[![Go Report Card](https://goreportcard.com/badge/github.com/kishaningithub/randomtext)](https://goreportcard.com/report/github.com/kishaningithub/randomtext)
[![Downloads](https://img.shields.io/github/downloads/kishaningithub/randomtext/latest/total.svg)](https://github.com/kishaningithub/randomtext/releases)
[![Latest release](https://img.shields.io/github/release/kishaningithub/randomtext.svg)](https://github.com/kishaningithub/randomtext/releases)

Command line random text generator

## Usage

```zsh
➜  randomtext git:(master) randomtext -h
Usage of randomtext:
  -size string
        Size of generated random text in KB, MB, GB, TB (default "1MB")
  -type string
        Type of text to be generated - chars, words, zeros (default "chars")
```

### Examples

- Generate 1 KB random charecters
  - `randomtext --size=1KB | pv > output.txt`
- Generate 100MB random words
  - `randomtext --size=100MB | pv > output.txt`
- Generate 1GB of zeros
  - `randomtext --size=1GB --type=zeros | pv > output.txt`

Source of [words](https://github.com/dwyl/english-words). The above commands use [pv](https://www.ivarch.com/programs/pv.shtml) for pipeline visualization only.

## Use case

 Testing the capacity of code editors/IDEs

## Installation

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