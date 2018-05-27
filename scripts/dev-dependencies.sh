#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

function main() {
    passGoDevDependencies
}

function passGoDevDependencies() {
  echo "Installing developer tools."

  echo "errcheck https://github.com/kisielk/errcheck"
  go get -u github.com/kisielk/errcheck

  echo "golint https://github.com/golang/lint/golint"
  go get -u github.com/golang/lint/golint

  echo "megacheck https://github.com/dominikh/go-tools/tree/master/cmd/megacheck"
  go get -u honnef.co/go/tools/cmd/megacheck
}

main