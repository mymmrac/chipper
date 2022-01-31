<p align="center">
  <img src="docs/chipper.png" alt="Chipper logo" width="512px" style="border-radius: 12px;">
</p>

<p align="center">
    <a href="https://github.com/mymmrac/chipper/actions/workflows/ci.yaml">
        <img src="https://github.com/mymmrac/chipper/actions/workflows/ci.yaml/badge.svg" alt="Chipper CI">
    </a>
    <a href="https://goreportcard.com/report/github.com/mymmrac/chipper">
        <img src="https://goreportcard.com/badge/github.com/mymmrac/chipper" alt="Go Report">
    </a>
</p>

# 👾 Chipper

Chipper is a small tool for testing CPUs.

It runs multiple tests and measures execution time. Based on that time, you can relatively compare CPUs.

> Note:Keep in mind that results may vary from run to run, and results don't show the real picture of which CPU is
> better. Speed of terminal, RAM, or even already running processes may dramatically affect results.

## 📋 Available tests

- Fibonacci sequence (`1, 1, 2, 3, 5`)
- Factorial (`1, 2, 6, 24, 120`)
- Trigonometry (`atan(tan(atan(... + e)))`)

## ⚡️ Install & Run

Install using `go install`

```shell
go install github.com/mymmrac/chipper@latest
```

Run

```shell
chipper
```

Or run in `simple-mode` without pretty TUI (maybe slightly faster)

```shell
chipper -s
```

> Note: Make sure to add `$GOPATH/bin` to `$PATH`

## 💾 Build & Run

Build binary

```shell
go build -o bin/chipper github.com/mymmrac/chipper
```

Run

```shell
./bin/chipper
```
