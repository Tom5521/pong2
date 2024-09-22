#!/bin/bash

os="$1"
arch="$2"

bin_name=./builds/pong-$os-$arch
cc="gcc"
cgo_enabled=1
ldflags="-s -w"
tags="-tags sdl"

if [[ "$1" == "windows" ]]; then
  bin_name="$bin_name.exe"
  cc="x86_64-w64-mingw32-gcc"
  ldflags="$ldflags -H=windowsgui"
  tags=""
fi
CGO_ENABLED=$cgo_enabled CC=$cc GOOS=$os GOARCH=$arch \
  go build -ldflags "$ldflags" -v $tags -o "$bin_name"
