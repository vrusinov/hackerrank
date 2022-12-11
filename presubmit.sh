#!/bin/bash

# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
# SPDX-License-Identifier: Apache-2.0

set -e

# Install hook:
ln -sf ../../presubmit.sh .git/hooks/pre-commit

if which docker ;
then
    docker run --rm -e "WORKSPACE=${PWD}" -v "$PWD:/app" shiftleft/sast-scan scan --build
else
    echo "WARN: docker not installed"
fi

# Run gometalinter in all go directories
for d in $(find . -name \*.go | sed -r 's|/[^/]+$||' |sort -u) ; do
    pushd $d
    golangci-lint run
    popd
done

pytype .

reuse lint

# Install presubmit
ln -sf ../../presubmit.sh .git/hooks/pre-commit
