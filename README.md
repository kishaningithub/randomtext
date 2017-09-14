# Randomtext 

Command line random text generator (Development in progress)

## Usage

```bash
  randomtext --size=1KB
  randomtext --size=1MB
  randomtext --size=1GB
  randomtext --size=1TB
 ```

To generate content with valid random words. 

```bash

  randomtext --type=words

```

To generate zeros

```bash

 randomtext --type=zeros

 ```

_Source of [words](https://github.com/dwyl/english-words)_

## Use case

 Testing the capacity of code editors/IDEs

## Installation

With the [Go compiler](https://golang.org/dl/) installed, run:

`go get -u github.com/kishaningithub/randomtext/...`

The `randomtext` command will now be compiled and located in `$GOPATH/bin`

## Development

To build

```bash

go generate -x .
go install -v ./...

```