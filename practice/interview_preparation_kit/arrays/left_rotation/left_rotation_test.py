# SPDX-FileCopyrightText: 2022 Vladimir Rusinov
#
# SPDX-License-Identifier: Apache-2.0

import unittest
import left_rotation

from parameterized import parameterized


class LeftRotationTest(unittest.TestCase):

    @parameterized.expand([
        ([1, 2, 3, 4, 5], 4, [5, 1, 2, 3, 4]),
        ([1, 2, 3, 4, 5], 5, [1, 2, 3, 4, 5]),
        ([1, 2, 3, 4, 5], 9, [5, 1, 2, 3, 4]),
        ([1, 2, 3, 4, 5], 0, [1, 2, 3, 4, 5]),
    ])
    def test_happy(self, a, d, expected):
        self.assertListEqual(left_rotation.rotLeft(a, d), expected)

    def test_boolean(self):
        a = True
        b = True
        self.assertEqual(a, b)


if __name__ == '__main__':
    unittest.main()
