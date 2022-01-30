# Chipper

[![CI Status](https://github.com/mymmrac/chipper/actions/workflows/ci.yml/badge.svg)](https://github.com/mymmrac/chipper/actions/workflows/ci.yml)

Chipper is a small tool for testing CPUs.

## Current tests

- Fibonacci sequence (`1, 1, 2, 3, 5`)
- Factorial (`1, 2, 6, 24, 120`)
- Trigonometry (`atan(tan(atan(... + e)))`)

## Run

Build binary:

```shell
go build -o bin/chipper chipper
```

Run test:

```shell
./bin/chipper
```

Results of tests will be located in `results` folder.