class Difference:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = 0

	# Add your code here
    
    def computeDifference(self):
      
        _ = len(self.__elements)
        items = (self.__elements)
        
        for i in range(_):
            for x in range(i):
                result = abs(items[i]-items[x])
                if result >= self.maximumDifference:
                    self.maximumDifference = result


# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()