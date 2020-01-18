class PeekingIterator:
    def __init__(self, iterator):
        """
        Initialize your data structure here.
        :type iterator: Iterator
        """
        self.it1 = iterator
        self.it2 = []
        

    def peek(self):
        """
        Returns the next element in the iteration without advancing the iterator.
        :rtype: int
        """
        if len(self.it2) == 0:
            top = self.it1.next()
            self.it2.append(top)
        return self.it2[0]


    def next(self):
        """
        :rtype: int
        """
        if len(self.it2) == 0:
            return self.it1.next()
        return self.it2.pop(0)        

    def hasNext(self):
        """
        :rtype: bool
        """
        return self.it1.hasNext() or len(self.it2)>0
