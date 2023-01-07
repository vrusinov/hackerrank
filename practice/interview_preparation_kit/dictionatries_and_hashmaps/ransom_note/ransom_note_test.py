# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
#
# SPDX-License-Identifier: Apache-2.0

import unittest

import ransom_note
from parameterized import parameterized


class RansomNoteTest(unittest.TestCase):

    @parameterized.expand([
        ("give me one grand today night", "give one grand today", True),
        ("two times three is not four", "two times two is four", False),
        ("ive got a lovely bunch of coconuts", "ive got some coconuts", False),
    ])
    def test_happy(self, magazine, note, expected):
        self.assertEqual(ransom_note.check_magazine(magazine, note), expected)


if __name__ == '__main__':
    unittest.main()
