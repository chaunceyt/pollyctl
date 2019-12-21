#!/bin/bash -e

# Build for supported platforms.

for GOOS in darwin linux; do
  for GOARCH in amd64; do
    env GOOS=$GOOS GOARCH=$GOARCH go build -v -o releases/pollyctl-$GOOS-$GOARCH
  done
done
