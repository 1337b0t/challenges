"""
You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.

Examples
[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)
"""

def find_outlier(integers):

    result = []
    num1 = integers[0]%2
    num2 = integers[1]%2
    num3 = integers[2]%2

    checkEven = (num1+num2+num3)/3.0

    # even
    if checkEven <= 1.0/3:
        result = list(filter(lambda x: x % 2 !=0, integers))
    # odd 
    else:
        result = list(filter(lambda x: x % 2 ==0, integers))

    result = sorted(result,reverse=True)
    
    return result[0]