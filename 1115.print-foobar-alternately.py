from threading import Semaphore, Thread

class FooBar:
    def __init__(self, n):
        self.n = n
        self.sem_foo = Semaphore(1)
        self.sem_bar = Semaphore(0)


    def foo(self, printFoo: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.sem_foo.acquire()
            printFoo()
            self.sem_bar.release()


    def bar(self, printBar: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.sem_bar.acquire()
            printBar()
            self.sem_foo.release()
