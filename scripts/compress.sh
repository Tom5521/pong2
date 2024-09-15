#!/bin/bash

bin="$1"

which upx >/dev/null 2>&1
if [[ $? != 0 ]]; then
  exit 0
fi

upx -9 $bin
