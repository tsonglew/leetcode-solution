#[derive(Debug)]
struct MinStack {
    el_stack: Vec<i32>,
    min_stack: Vec<i32>,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {

    fn new() -> Self {
        MinStack{
            el_stack: Vec::new(),
            min_stack: Vec::new(),
        }
    }
    
    fn push(&mut self, val: i32) {
        self.el_stack.push(val);
        if let Some(last) = self.min_stack.last() {
            if val <= *last {
                self.min_stack.push(val);
            }
        } else {
            self.min_stack.push(val);
        }
        // println!("{:?}", self);
    }
    
    fn pop(&mut self) {
        if let Some(last) = self.el_stack.pop() {
            if let Some(min_last) = self.min_stack.last() {
                if *min_last == last {
                    self.min_stack.pop();
                }
            }
        }
        // println!("{:?}", self);

    }
    
    fn top(&self) -> i32 {
        *self.el_stack.last().unwrap()
    }
    
    fn get_min(&self) -> i32 {
        *self.min_stack.last().unwrap()
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * let obj = MinStack::new();
 * obj.push(val);
 * obj.pop();
 * let ret_3: i32 = obj.top();
 * let ret_4: i32 = obj.get_min();
 */
