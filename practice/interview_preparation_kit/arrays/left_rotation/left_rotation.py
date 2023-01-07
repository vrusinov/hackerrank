#!/bin/python3

# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
#
# SPDX-License-Identifier: Apache-2.0

"""
Hackerrank Left Rotation problem

https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem
"""

import os


def rotLeft(a, d):
    """
    Performs left rotation of array a by d rotations.

    Args:
        a: input array.
        d: number of rotations to do.
    Returns:
        list of ints, rotated array.
    """
    if d == len(a):
        return a
    # No point in rotating over len of array.
    if d > len(a):
        d = d % len(a)
    # Very easy to do in python - split array by 2 in d'th
    # emelent. Join two of them together.
    return a[d:] + a[:d]


if __name__ == '__main__':
    output_path = os.environ.get('OUTPUT_PATH', '/dev/stdout')
    fptr = open(output_path, 'w')

    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    d = int(first_multiple_input[1])
    a = list(map(int, input().rstrip().split()))

    result = rotLeft(a, d)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
