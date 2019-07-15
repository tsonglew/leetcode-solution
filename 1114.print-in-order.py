from threading import Semaphore

class Foo:
    def __init__(self):
        self.second_sem = Semaphore(0)
        self.third_sem = Semaphore(0)


    def first(self, printFirst: 'Callable[[], None]') -> None:
        
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.second_sem.release()


    def second(self, printSecond: 'Callable[[], None]') -> None:
        with self.second_sem:
        # printSecond() outputs "second". Do not change or remove this line.
            printSecond()
            self.third_sem.release()


    def third(self, printThird: 'Callable[[], None]') -> None:
        with self.third_sem:
        # printThird() outputs "third". Do not change or remove this line.
            printThird()