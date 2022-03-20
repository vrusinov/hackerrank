#!/bin/python3

import os

def hourglass_sum(arr, i, j):
    """Calculates hourglass for i,j"""
    return (
        arr[i][j]   + arr[i][j+1]   + arr[i][j+2] +
                      arr[i+1][j+1] +
        arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
    )

def max_hourglass_sum(arr):
    # This will be calculated twice.
    # Meh.
    maximum = hourglass_sum(arr, 0, 0)
    for i in range(len(arr)-2):
        for j in range(len(arr[i])-2):
            h = hourglass_sum(arr, i, j)
            maximum = max(maximum, h)
    return maximum

if __name__ == '__main__':
    output_path = os.environ.get('OUTPUT_PATH', '/dev/stdout')
    fptr = open(output_path, 'w')

    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    result = max_hourglass_sum(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
