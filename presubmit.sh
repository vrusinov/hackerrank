#!/bin/bash

set -e

if which docker ;
then
    docker run --rm -e "WORKSPACE=${PWD}" -v "$PWD:/app" shiftleft/sast-scan scan --build
else
    echo "WARN: docker not installed"
fi

# Run gometalinter in all go directories
for d in $(find . -name \*.go | sed -r 's|/[^/]+$||' |sort -u) ; do
    golint $d
    which gometalinter && gometalinter $d
    pushd $d
    golangci-lint run
    popd
done

pytype .
