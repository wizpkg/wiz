#!/usr/bin/env sh
for i in $(go list ./...); do
  curl -L godoc.org/$i
done
