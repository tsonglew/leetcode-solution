type Temp struct {
    Val int
    Idx int
}

func dailyTemperatures(T []int) []int {
    if len(T) == 0 { return []int{} }
    stack := []Temp{}
    stack = append(stack, Temp{T[len(T)-1], len(T)-1})
    T[len(T)-1] = 0
    for i := len(T)-2; i>=0; i-- {
        n := len(stack)-1
        curVal := T[i]
        for n >= 0 && stack[n].Val <= T[i] {
            n --
        }
        stack = stack[0:n+1]
        if len(stack) == 0 {
            T[i] = 0
        } else {
            T[i] = stack[len(stack)-1].Idx-i
        }
        stack = append(stack, Temp{curVal, i})
    }
    return T
}
