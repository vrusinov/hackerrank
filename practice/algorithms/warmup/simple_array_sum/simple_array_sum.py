#!/bin/python3
"""Simple Array Sum

https://www.hackerrank.com/challenges/simple-array-sum/problem
"""

import os
import sys

from typing import Sequence

def simpleArraySum(ar: Sequence):
    return sum(ar)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    ar_count = int(input())

    ar = list(map(int, input().rstrip().split()))

    result = simpleArraySum(ar)

    fptr.write(str(result) + '\n')

    fptr.close()
