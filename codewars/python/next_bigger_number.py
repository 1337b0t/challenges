#!/bin/python3
import unittest
"""
You have to create a function that takes a positive integer number and returns the next bigger number formed by the same digits:

12 ==> 21
513 ==> 531
2017 ==> 2071
If no bigger number can be composed using those digits, return -1:

9 ==> -1
111 ==> -1
531 ==> -1
"""

class SampleTests(unittest.TestCase):
    def test(self):
        self.assertEqual(next_bigger(12),21)
        self.assertEqual(next_bigger(513),531)
        self.assertEqual(next_bigger(2017),2071)
        self.assertEqual(next_bigger(414),441)
        self.assertEqual(next_bigger(144),414)

def next_bigger(n):

    sortedList=sorted(str(n))
    nextNumber=n

    if str(n) == ''.join(sortedList[::-1]):
        return -1

    while(True):
        nextNumber+=1
        if sorted(str(nextNumber)) == sortedList:
            return nextNumber
    
if __name__ == "__main__":
    unittest.main()