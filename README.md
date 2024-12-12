# Advent of Code

Golang implementation of [Advent of Code](https://adventofcode.com).

My ultimate goal is to compile these to WebAssembly and run them in the browser; so that anybody can run their input. (WIP)

## Get input file

NOTE: this assumes a cookie value set via:
```shell
export AOC_SESSION_COOKIE=<your session cookie>
```

OR

```shell
source .env 
```

... and then you can fetch:

```shell
go run ./fetch -year 2023 -day 1 > ./2023/1/input.txt
```

OR; Install and then use it:

```shell
go install ./fetch
fetch -year 2023 -day 1 > ./2023/1/input.txt
```

## Run it

```shell
go run ./2023/1 < ./2023/1/input.txt
```

## Compile to WASM

```shell
GOOS=wasip1 GOARCH=wasm go build -o 2023/1/main.wasm 2023/1/main.go
```

## Run as WASM

```shell
wasmtime 2023/1/main.wasm < ./2023/1/input.txt
```
