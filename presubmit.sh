#!/bin/bash

set -e

docker run --rm -e "WORKSPACE=${PWD}" -v "$PWD:/app" shiftleft/sast-scan scan --build

# Run gometalinter in all go directories
for d in $(find . -name \*.go | sed -r 's|/[^/]+$||' |sort -u) ; do
    pushd $d > /dev/null
    gometalinter
    popd
done
