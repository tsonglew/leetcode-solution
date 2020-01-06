type Stack struct {
	Elements []int
}

func (s *Stack) Push(x int) {
	s.Elements = append(s.Elements, x)
}

func (s *Stack) Pop() int {
	result := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return result
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

type MyQueue struct {
	Front int
	S1 Stack
	S2 Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		0,
		Stack{make([]int, 0)},
		Stack{make([]int, 0)},
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	if this.Empty() {
		this.Front = x
	}
	this.S1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.S2.IsEmpty() {
		for !this.S1.IsEmpty() {
			this.S2.Push(this.S1.Pop())
		}
	} 
	poped := this.S2.Pop()
	if this.S2.IsEmpty() {
		for !this.S1.IsEmpty() {
			this.S2.Push(this.S1.Pop())
		}
	}
	if !this.Empty() {
		this.Front = this.S2.Pop()
		this.S2.Push(this.Front)
	}
	return poped
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Front
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.S1.IsEmpty() && this.S2.IsEmpty()
}
