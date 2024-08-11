use std::cmp;

impl Solution {
    pub fn minimum_delete_sum(s1: String, s2: String) -> i32 {
        let mut mem = vec![vec![0;s2.len()+1];s1.len()+1];
        let s1i = if s1.len() == 0 { vec![] } else { s1.chars().map(|a| a as i32).collect() };
        let s2i = if s2.len() == 0 { vec![] } else { s2.chars().map(|a| a as i32).collect() };

        // println!("{:?}", s1i);
        // println!("{:?}", s2i);

        mem[0][0] = 0;
        for i in 0..mem.len() {
            for j in 0..mem[0].len() {
                if i == 0 && j == 0 {
                    continue;
                } else if i == 0 {
                    mem[i][j] = s2i[j-1] + mem[i][j-1];
                } else if j == 0 {
                    mem[i][j] = s1i[i-1] + mem[i-1][j];
                } else {
                    if s1i[i-1] == s2i[j-1] {
                        mem[i][j] = mem[i-1][j-1];
                    } else {
                        mem[i][j] = cmp::min(
                            mem[i-1][j] + s1i[i-1],
                            mem[i][j-1] + s2i[j-1]
                        );
                    }
                }
            }
        }
        // println!("{:?}", mem);
        mem[mem.len()-1][mem[0].len()-1]
    }
