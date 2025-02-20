#!/bin/python3

# SPDX-FileCopyrightText: 2022,2025 Vladimir Rusinov
# SPDX-License-Identifier: Apache-2.0

"""A Very Big Sum.

https://www.hackerrank.com/challenges/a-very-big-sum/problem
"""

import os
from typing import Iterable


def a_very_big_sum(arr: Iterable[int]) -> int:
    """Calculate a very big sum.

    Args:
        arr: an arary to sum.

    Returns:
        a sum of all elements of array.
    """
    return sum(arr)


if __name__ == '__main__':
    with open(os.environ['OUTPUT_PATH'], 'w', encoding='utf-8') as fptr:
        ar_count = int(input())
        ar = list(map(int, input().rstrip().split()))
        result = a_very_big_sum(ar)
        fptr.write(str(result) + '\n')
