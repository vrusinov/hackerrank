#!/usr/bin/env python3

# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
#
# SPDX-License-Identifier: Apache-2.0

"""Sales by Match.

https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
"""

import collections
import os
from typing import Sequence


def sockMerchant(n, ar: Sequence[int]):
    socks_per_color = collections.Counter()
    for color in ar:
        socks_per_color[color] += 1
    total_paris = 0
    for _color, num in socks_per_color.items():
        total_paris += num // 2
    return total_paris


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())
    ar = list(map(int, input().rstrip().split()))

    result = sockMerchant(n, ar)

    fptr.write(str(result) + '\n')
    fptr.close()
