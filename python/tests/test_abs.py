import unittest
import src.abs as myAbs


class TestAbs(unittest.TestCase):
    def test_abs(self):
        self.assertEqual(myAbs.abs(9), 9)
