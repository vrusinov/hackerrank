#!/usr/bin/env python3
"""Counting Valleys.

https://www.hackerrank.com/challenges/counting-valleys/problem
"""

import math
import os
import random
import re
import sys


def countingValleys(steps, path):
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
