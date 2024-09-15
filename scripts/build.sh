#!/bin/bash

os="$1"
arch="$2"

bin_name=./builds/pong-$os-$arch
compiler=gcc
cgo_enabled=1
ldflags="-s -w"
tags="-tags 'rgfw'"

if [[ $1 == "windows" ]]; then
  bin_name="$bin_name.exe"
  compiler=x86_64-w64-mingw32-gcc
  ldflags="$ldflags -H=windowsgui"
  tags=""
fi
CGO_ENABLED=$cgo_enabled CC=$compiler GOOS=$os GOARCH=$arch \
  go build -ldflags "$ldflags" -v $tags -o "$bin_name"

if [[ $SKIP_COMPRESS == 1 ]]; then
  exit 0
fi

just compress "$bin_name"
