impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut s: Vec<i32> = Vec::new();
        for t in tokens.iter() {
            match t.as_str() {
                "+"|"-"|"*"|"/" => {
                    let n1 = s.pop().unwrap();
                    let n2 = s.pop().unwrap();
                    s.push(match t.as_str() {
                        "+" => n1 + n2,
                        "-" => n2 - n1,
                        "*" => n2 * n1,
                        "/" => n2 / n1,
                        _ => 0
                    });
                },
                _ => {
                    s.push(t.parse::<i32>().unwrap());
                }
            }
        }
        s.pop().unwrap()
    }
}
