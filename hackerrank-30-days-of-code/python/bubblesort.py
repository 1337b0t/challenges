#!/bin/python3

import sys

n = int(input().strip())
a = list(map(int, input().strip().split(' ')))

numSwaps = 0


for i in range(n):
    for x in range(n-1):
        if a[x] > a[x+1]:
            temp = a[x]
            a[x] = a[x+1]
            a[x+1] = temp
            numSwaps+=1
    
    if numSwaps == 0:
        break

print("Array is sorted in " + str(numSwaps) + " swaps.")
print("First Element: " + str(a[0]))
print("Last Element: " + str(a[n-1]))



