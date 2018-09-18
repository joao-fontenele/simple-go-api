#!/usr/bin/env bash

sigintHandler() {
  kill $PID
  exit
}

trap sigintHandler INT TERM
# trap sigint_handler SIGINT

package="server"
sourceDir="src"

while true; do
  rm $package
  go build $package
  $@ &
  PID=$!
  inotifywait -rq -e modify,create,delete $sourceDir
  kill $PID
done
