#!/bin/python3
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
#
# SPDX-License-Identifier: Apache-2.0
"""Sherlock and Anagrams

https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
"""

import os
import collections


def sherlock_and_anagrams(s: str) -> int:
    # First count how many times each substring appears.
    # sort substrings since order does not matter for anagrams
    unique_sorted_substrings: collections.Counter[str] = collections.Counter()
    for i in range(1, len(s)):
        for j in range(len(s)-i+1):
            unique_sorted_substrings[''.join(sorted(s[j:j+i]))] += 1
    ret = 0
    # Now for any substring which appears more than ones, count how many pairs
    # is there.
    # E.g. if we have substring which appears twice, we have 2 paris
    # For 3 times we have 3, for 4 => 6.
    #
    # https://en.wikipedia.org/wiki/Combination
    # https://stackoverflow.com/questions/18859430/how-do-i-get-the-total-number-of-unique-pairs-of-a-set-in-the-database
    for count in unique_sorted_substrings.values():
        if count > 1:
            ret += count * (count - 1) // 2
    return ret


if __name__ == '__main__':
    fptr = open(os.environ.get('OUTPUT_PATH', '/dev/stdout'), 'w')

    q = int(input().strip())

    for q_itr in range(q):
        s = input()

        result = sherlock_and_anagrams(s)

        fptr.write(str(result) + '\n')

    fptr.close()
