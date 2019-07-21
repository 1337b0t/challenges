from math import sqrt

class Prime:
   
    def isPrime(self,n):

        for i in range(2,int(sqrt(n))+1):

            if n % i is 0:
                return False

        return True   

if __name__ == "__main__":

    n = int(input())
    for __ in range(n):
        num = int(input())
        p = Prime()

        if p.isPrime(num) and num >=2:
            print("Prime")
        else:
            print("Not prime")
