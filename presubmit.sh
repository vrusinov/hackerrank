#!/bin/bash

# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
# SPDX-License-Identifier: Apache-2.0

set -e

if which docker ;
then
    docker run --rm -e "WORKSPACE=${PWD}" -v "$PWD:/app" shiftleft/sast-scan scan --build
else
    echo "WARN: docker not installed"
fi

pytype .

reuse lint
