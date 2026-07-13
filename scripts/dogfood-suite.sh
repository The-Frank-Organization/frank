#!/bin/sh
set -eu
export FRANK_DOGFOOD_NESTED=1
exec go test -p=1 ./... -count=1
