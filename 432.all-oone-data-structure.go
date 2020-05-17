type Elem struct {
    Key string
    Val int 
    Prev *Elem
    Next *Elem
}

type AllOne struct {
    Head *Elem
    Tail *Elem
    M map[string]*Elem
    Len int 
}


/** Initialize your data structure here. */
func Constructor() AllOne {
    head, tail := &Elem{}, &Elem{}
    head.Next = tail
    tail.Prev = head
    return AllOne{
        M: make(map[string]*Elem),
        Head: head,
        Tail: tail,
    }
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
    if n, ok := this.M[key]; ok {
        n.Val++
        for n.Next != this.Tail && n.Val > n.Next.Val {
            nex := n.Next
            this.M[key], this.M[nex.Key] = this.M[nex.Key], this.M[key]
            this.M[key].Val, this.M[nex.Key].Val = n.Val, nex.Val
            this.M[key].Key, this.M[nex.Key].Key = n.Key, nex.Key
            n = n.Next
        }
    } else {
        n := &Elem{Key: key, Val: 1, Prev: this.Head, Next: this.Head.Next}
        this.Head.Next.Prev = n
        this.Head.Next = n
        this.M[key] = n
        this.Len += 1
    }
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
    if n, ok := this.M[key]; ok {
        n.Val --
        if n.Val == 0 {
            n.Prev.Next = n.Next
            n.Next.Prev = n.Prev
            delete(this.M, key)
            this.Len -= 1
            return
        }
        for n.Prev != this.Head && n.Val < n.Prev.Val {
            nex := n.Prev
            this.M[key], this.M[nex.Key] = this.M[nex.Key], this.M[key]
            this.M[key].Val, this.M[nex.Key].Val = n.Val, nex.Val
            this.M[key].Key, this.M[nex.Key].Key = n.Key, nex.Key
            n = n.Prev
        }
    } else {
        return
    }
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
    if this.Len == 0 {
        return ""
    }
    return this.Tail.Prev.Key
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
    if this.Len == 0 {
        return ""
    }
    return this.Head.Next.Key
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
