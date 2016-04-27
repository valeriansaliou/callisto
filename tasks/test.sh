#!/bin/bash

ABSPATH=$(cd "$(dirname "$0")"; pwd)
BASE_DIR="$ABSPATH/../"

rc=0

pushd "$BASE_DIR" > /dev/null
  lint_output=$(golint *.go)

  if [[ ! -z $lint_output ]]
  then
    echo "$lint_output"
    rc=1
  fi
popd > /dev/null

exit $rc
