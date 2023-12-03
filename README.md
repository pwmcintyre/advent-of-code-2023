# advent-of-code-2023

Golang implementation of [Advent of Code 2023](https://adventofcode.com/2023).

My ultimate goal is to compile these to WebAssembly and run them in the browser; so that anybody can run their input.

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
go run ./fetch -day 1 > ./days/1/input.txt
```

OR; Install and then use it:

```shell
go install ./fetch
fetch -day 1 > ./days/1/input.txt
```

## Run it

```shell
go run ./days/1 < ./days/1/input.txt
```

## Compile to WASM

```shell
GOOS=wasip1 GOARCH=wasm go build -o days/1/main.wasm days/1/main.go
```

## Run as WASM

```shell
wasmtime days/1/main.wasm < ./days/1/input.txt
```
s