# advent-of-code-2023

Golang implementation of [Advent of Code 2023](https://adventofcode.com/2023).

My ultimate goal is to compile these to WebAssembly and run them in the browser; so that anybody can run their input.

## Get input file

```shell
export AOC_SESSION_COOKIE=<your session cookie>
go run ./fetch -day 1 > 1.in
```

## Run it

```shell
go run ./days/1 < 1.in
```

## Compile to WASM

```shell
GOOS=wasip1 GOARCH=wasm go build -o days/1/main.wasm days/1/main.go
```

## Run as WASM

```shell
wasmtime days/1/main.wasm < 1.txt
```
s