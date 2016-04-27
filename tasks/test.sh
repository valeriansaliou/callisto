#!/bin/bash

ABSPATH=$(cd "$(dirname "$0")"; pwd)
BASE_DIR="$ABSPATH/../"

rc=0

pushd "$BASE_DIR" > /dev/null
  if [[ ! -z $(golint *.go) ]]
  then
    rc=1
  fi
popd > /dev/null

exit $rc
