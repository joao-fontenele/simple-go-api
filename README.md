[![Go Report Card](https://goreportcard.com/badge/github.com/joao-fontenele/simple-go-api)](https://goreportcard.com/report/github.com/joao-fontenele/simple-go-api)

# Simple go http rest api

## Requirements

- golang v1.11
- notify-tools: for inotifywait

## How to run

- add project directory to GOPATH: `export GOPATH=$GOPATH:$PWD`
- `./reload.sh ./server # nodemon like server`
- `go install server # builds server binaries on ./bin folder`

## Use the server

- `curl -X GET localhost:8080/hello`
