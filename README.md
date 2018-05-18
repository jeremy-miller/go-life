[![Build Status](https://travis-ci.org/jeremy-miller/life-go.svg?branch=master)](https://travis-ci.org/jeremy-miller/life-go)
[![Coverage Status](https://coveralls.io/repos/github/jeremy-miller/life-go/badge.svg?branch=master)](https://coveralls.io/github/jeremy-miller/life-go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jeremy-miller/life-go)](https://goreportcard.com/report/github.com/jeremy-miller/life-go)
[![GoDoc](https://godoc.org/github.com/jeremy-miller/life-go?status.svg)](https://godoc.org/github.com/jeremy-miller/life-go)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jeremy-miller/life-go/blob/master/LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.10.2-blue.svg)]()

# Life (in Go)
Go implementation of
[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).
The implementation is separated into a library (which implements the game logic)
and an executable.

![Usage](https://github.com/jeremy-miller/life-go/blob/master/usage.gif)

## Table of Contents
- [Prerequisites](#prerequisites)
- [Dependencies](#dependencies)
    - [Update](#update)
    - [Check Status](#check-status)
- [Build](#build)
- [Static Code Analysis](#static-code-analysis)
- [Test](#test)
    - [Execute](#execute)
    - [Coverage](#coverage)
- [Clean](#clean)
- [Simplify](#simplify)
- [Run](#run)
- [License](#license)

## Prerequisites
- [Go](https://golang.org/doc/install) (at least version 1.10.2)

## Dependencies
### Update
Update existing dependencies (or install new dependencies): `make dep`

### Check Status
Check for outdated dependencies: `make dep-versions`

## Build
Build: `make build`

## Static Code Analysis
Lint (using [gometalinter](https://github.com/alecthomas/gometalinter)):
`make lint`

## Test
### Execute
Execute tests (with [data race detection](https://golang.org/doc/articles/race_detector.html)):
`make test`

_NOTE: Data race detection is only available on `linux/amd64`, `freebsd/amd64`,
`darwin/amd64`, and `windows/amd64` platforms._

### Coverage
Output test code coverage for each package: `make cover`

## Clean
Remove all compiled binaries, as well as data created during testing:
`make clean`

## Simplify
Run `gofmt -s` to have the compiler simplify all code: `make simplify`

## Run
By default, the game will run for five iterations, printing the board state to
the terminal between each iteration.

To run the game, execute the following command: `make run`

To run more or less iterations of the game, an optional `iterations` argument
can be passed to the `make run` command: `make run iterations=10`

## License
[MIT](https://github.com/jeremy-miller/life-go/blob/master/LICENSE)
