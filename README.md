# Chipper

[![CI Status](https://github.com/mymmrac/chipper/actions/workflows/ci.yml/badge.svg)](https://github.com/mymmrac/chipper/actions/workflows/ci.yml)

Chipper is a small tool for testing CPUs.

## Current tests

- Fibonacci sequence (`1, 1, 2, 3, 5`)
- Factorial (`1, 2, 6, 24, 120`)
- Trigonometry (`atan(tan(atan(... + e)))`)

## Install & Run

Install using `go install`

```shell
go install github.com/mymmrac/chipper@latest
```

Run

```shell
chipper
```

> Note: Make sure to add `$GOPATH/bin` to `$PATH`

## Build & Run

Build binary

```shell
go build -o bin/chipper chipper
```

Run tests

```shell
./bin/chipper
```
