[![Go Report Card](https://goreportcard.com/badge/github.com/joao-fontenele/simple-go-api)](https://goreportcard.com/report/github.com/joao-fontenele/simple-go-api)

# Simple go http rest api

## Requirements

- golang v1.11

## How to run

- add project directory to GOPATH: `export GOPATH=$GOPATH:$PWD`
- `go build server # builds server on .`
- `go install server # builds server binaries on ./bin folder`
- `./server # runs server`

## Use the server

- `curl -X GET localhost:8080/hello`
