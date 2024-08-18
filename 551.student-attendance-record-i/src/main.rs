impl Solution {
    pub fn check_record(s: String) -> bool {
        let mut continue_l = 0;
        let mut a = 0;
        let c: Vec<char> = s.chars().collect();
        for &i in c.iter() {
            if i == 'A' {
                a += 1
            }
            if i == 'L' {
                continue_l += 1;
            } else {
                continue_l = 0;
            }
            if continue_l >= 3 || a > 1 {
                return false;
            }
        }
        return true;
    }
}
