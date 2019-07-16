from threading import Semaphore

class ZeroEvenOdd:
    def __init__(self, n):
        self.n = n
        self.sem_zero = Semaphore(1)
        self.sem_even = Semaphore(0)
        self.sem_odd = Semaphore(0)
        
    def generate_nums(self, n):
        i = n
        result = []
        while i <= self.n:
            yield i
            i += 2
    
	# printNumber(x) outputs "x", where x is an integer.
    def zero(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(self.n):
            self.sem_zero.acquire()
            printNumber(0)
            if i % 2 == 0:
                self.sem_odd.release()
            else:
                self.sem_even.release()        
        
    def even(self, printNumber: 'Callable[[int], None]') -> None:
        for i in self.generate_nums(2):
            self.sem_even.acquire()
            printNumber(i)
            self.sem_zero.release()        
        
    def odd(self, printNumber: 'Callable[[int], None]') -> None:
        for i in self.generate_nums(1):
            self.sem_odd.acquire()
            printNumber(i)
            self.sem_zero.release()