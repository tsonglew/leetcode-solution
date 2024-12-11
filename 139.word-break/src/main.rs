use std::collections::HashSet;

impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let mut hs: HashSet<String> = HashSet::new();
        for i in word_dict {
            hs.insert(i);
        }

        let mut dp = vec![0; s.len()];
        for i in (0..s.len()).rev() {
            if hs.contains(&s[i..s.len()]){
                dp[i] = 1;
                continue;
            }
            for j in i..s.len() {
                // println!("i: {}, j:{}", i, j);
                if dp[j] == 0 {
                    continue;
                }
                if hs.contains(&s[i..j]) {
                    dp[i] = 1;
                    break;
                }
            }
        }
        // println!("dp: {:?}", dp);
        dp[0] == 1
    }
}
