#!/bin/sh

OS="darwin windows linux"
ARCH="amd64 arm64"

mkdir -p output

for os in $OS; do
  for arch in $ARCH; do
    echo "Building $os/$arch"
    GOOS=$os GOARCH=$arch go build -o output/ttoys
    tar -czf output/ttoys-$os-$arch.tar.gz output/ttoys
    rm output/ttoys
  done
done
