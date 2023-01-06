# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
#
# SPDX-License-Identifier: Apache-2.0

import unittest
import sherlock_and_anagrams

from parameterized import parameterized


class SherlockTest(unittest.TestCase):

    @parameterized.expand([
        ("abba", 4),
        ("abcd", 0),
    ])
    def test_happy(self, s, expected):
        self.assertEqual(sherlock_and_anagrams.sherlock_and_anagrams(s), expected)


if __name__ == '__main__':
    unittest.main()
