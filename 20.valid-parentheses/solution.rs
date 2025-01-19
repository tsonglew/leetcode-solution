impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut st: Vec<char> = vec![];
        for c in s.chars() {
            // println!("{}", c);
            match c {
                '(' | '[' | '{' => st.push(c),
                ')' | ']' | '}' => {
                    if let Some(last_el) = st.pop() {
                        if last_el != Solution::get_pair(c) {
                            return false;
                        }
                    } else {
                        return false;
                    }
                }
                _ => {}
            };
            // println!("{:?}", st);
        }

        st.len() == 0
    }

    pub fn get_pair(c: char) -> char {
        match c {
            ')' => '(',
            '}' => '{',
            ']' => '[',
            _ => '1'
        }
    }
}
