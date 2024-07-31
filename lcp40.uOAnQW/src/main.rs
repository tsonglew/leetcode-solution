impl Solution {
    pub fn maxmium_score(cards: Vec<i32>, cnt: i32) -> i32 {
        let mut cardscopy = cards.clone();
        cardscopy.sort_by(|a, b| b.cmp(a));
        let mut r = 0;
        let mut evn = -1;
        let mut odd = -1;
        for i in 0..cnt as usize {
            let now = cardscopy[i];
            if now % 2 == 0 {
                evn = now;
            } else {
                odd = now;
            }
            r += &cardscopy[i];
        }
        if r % 2 == 0 {
            return r;
        }

        for i in cnt as usize..cardscopy.len() as usize {
            if odd > 0 && cardscopy[i] % 2 == 0 {
                r += cardscopy[i] - odd;
                break;
            }
            if evn > 0 && cardscopy[i] % 2 != 0 {
                r += cardscopy[i] - evn;
                break;
            }
        }
        if r % 2 == 0 {
            r
        } else {
            0
        }
    }
}
