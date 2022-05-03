#!/usr/bin/env python3
"""Counting Valleys.

https://www.hackerrank.com/challenges/counting-valleys/problem
"""

import os
from typing import Any


def countingValleys(steps: Any, path: str) -> int:
    del steps  # Unused
    current_level = 0
    in_valley = False
    valleys = 0
    for step in path:
        if step == 'D':
            current_level -= 1
            if current_level < 0:
                if not in_valley:
                    valleys += 1
                in_valley = True
        elif step == 'U':
            current_level += 1
            if current_level >= 0:
                in_valley = False
    return valleys


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    steps = int(input().strip())

    path = input()

    result = countingValleys(steps, path)

    fptr.write(str(result) + '\n')

    fptr.close()
