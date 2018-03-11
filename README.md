[![Build Status](https://travis-ci.org/jeremy-miller/life-go.svg?branch=master)](https://travis-ci.org/jeremy-miller/life-go)
[![Coverage Status](https://coveralls.io/repos/github/jeremy-miller/life-go/badge.svg?branch=master)](https://coveralls.io/github/jeremy-miller/life-go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jeremy-miller/life-go)](https://goreportcard.com/report/github.com/jeremy-miller/life-go)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jeremy-miller/life-go/blob/master/LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.10-blue.svg)]()

Library - [![GoDoc](https://godoc.org/github.com/jeremy-miller/life-go/internal/life?status.svg)](https://godoc.org/github.com/jeremy-miller/life-go/internal/life)

Executable - [![GoDoc](https://godoc.org/github.com/jeremy-miller/life-go/cmd/life?status.svg)](https://godoc.org/github.com/jeremy-miller/life-go/cmd/life)

# Life (in Go)
Go implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).
The implementation is separated into a library (which implements the game logic) and an executable.

![Usage](https://github.com/jeremy-miller/life-go/blob/master/usage.gif)

## Table of Contents
- [Prerequisites](#prerequisites)
- [Documentation](#documentation)
- [Build](#build)
- [Static Code Analysis](#static-code-analysis)
- [Test](#test)
- [Run](#run)
- [License](#license)

## Prerequisites
- [Go](https://golang.org/doc/install) (at least version 1.10)

## Documentation
Documentation for the library can be found [here](https://godoc.org/github.com/jeremy-miller/life-go/internal/life).
Documentation for the executable can be found [here](https://godoc.org/github.com/jeremy-miller/life-go/cmd/life).

## Build
To build the game, execute the following command: ```make build```

## Static Code Analysis
To run the [gometalinter](https://github.com/alecthomas/gometalinter), execute the following command:
```make lint```

## Test
_NOTE: The source files will be linted before the tests run._

### Tests
To run the tests, execute the following command: ```make test```

### Data Race Detection
To run the tests with [data race detection](https://golang.org/doc/articles/race_detector.html),
execute the following command: ```make race```

_NOTE: The data race detection is only available on `linux/amd64`, `freebsd/amd64`, `darwin/amd64`, and `windows/amd64`._

## Run
By default, the game will run for five iterations, printing the board state to the terminal between each iteration.

To run the game, execute the following command: ```make run```

To run more or less iterations of the game, an optional argument can be passed to the `make run` command: ```make run iterations=10```

## License
[MIT](https://github.com/jeremy-miller/life-go/blob/master/LICENSE)
