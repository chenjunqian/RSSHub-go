#!/bin/sh

if [[ $OSTYPE == 'darwin'* ]]; then
  LIMITCOUNT=$(ulimit -n)
  if [[ $LIMITCOUNT < 524288 ]]; then
    ulimit -n 524288
    echo 'run ulimit -n 524288 to solve too many files problem'
  fi
fi

export env="dev"
gf run main.go
