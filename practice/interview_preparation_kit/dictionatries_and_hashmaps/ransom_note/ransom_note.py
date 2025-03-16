#!/bin/python3
# SPDX-FileCopyrightText: 2022-2025 Vladimir Rusinov
# SPDX-License-Identifier: Apache-2.0

"""Ransom Note.

https://www.hackerrank.com/challenges/ctci-ransom-note/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
"""  # noqa: E501

import collections


def check_magazine(magazine: str, note: str) -> bool:
    """Check if magazine has enough words to create note.

    Args:
        magazine: the magazine
        note: the note

    Returns:
        whether the note can be created from the magazine.
    """
    magazine_words = collections.Counter()
    for word in magazine.split():
        magazine_words[word] += 1
    for word in note.split():
        if magazine_words[word]:
            magazine_words[word] -= 1
        else:
            return False
    return True


def main():
    """Run the program."""
    first_multiple_input = input().rstrip().split()

    unused_m = int(first_multiple_input[0])  # noqa: F841
    unused_n = int(first_multiple_input[1])  # noqa: F841

    magazine = input().rstrip()
    note = input().rstrip()

    if check_magazine(magazine, note):
        print('Yes')
    else:
        print('No')


if __name__ == '__main__':
    main()
