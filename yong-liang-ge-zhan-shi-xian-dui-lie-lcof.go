type Stack struct {
    Values []int
}

func (this *Stack) Push(value int) {
    this.Values = append(this.Values, value)
}

func (this *Stack) Pop() (int, bool) {
    if len(this.Values) == 0 {
        return -1, false
    }
    value := this.Values[len(this.Values)-1]
    this.Values = this.Values[:len(this.Values)-1]
    return value, true
}

type CQueue struct {
    SIn *Stack
    SOut *Stack
}


func Constructor() CQueue {
    return CQueue{
        SIn: &Stack{[]int{}},
        SOut: &Stack{[]int{}},
    }
}


func (this *CQueue) AppendTail(value int)  {
    this.SIn.Push(value)
}


func (this *CQueue) DeleteHead() int {
    value, ok := this.SOut.Pop()
    if ok {
        return value
    }
    for value, ok := this.SIn.Pop(); ok; value, ok = this.SIn.Pop(){
        this.SOut.Push(value)
    }
    value, _ = this.SOut.Pop()
    return value
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
