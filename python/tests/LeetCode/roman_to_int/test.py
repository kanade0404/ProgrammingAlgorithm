import unittest
from python.LeetCode.roman_to_int.main import romanToInt


class MyTestCase(unittest.TestCase):

    def test_1(self):
        self.assertEqual(romanToInt("III"), 3)

    def test_2(self):
        self.assertEqual(romanToInt("LVIII"), 58)

    def test_3(self):
        self.assertEqual(romanToInt("MCMXCIV"), 1994)


if __name__ == '__main__':
    unittest.main()
