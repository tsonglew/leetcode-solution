use std::collections::HashMap;

impl Solution {
    pub fn can_partition(nums: Vec<i32>) -> bool {
        let mut mem: HashMap<usize, HashMap<i32, bool>> = HashMap::new();
        let total: i32 = (&nums).iter().sum();
        let cap = total / 2;
        let remainder = total % 2;
        if remainder > 0 {
            return false;
        }
        Solution::f(&nums, nums.len(), cap, &mut mem)
    }
    
    fn f(nums: &Vec<i32>, i: usize, cap: i32, mem: &mut HashMap<usize, HashMap<i32, bool>>) -> bool {
        if let Some(j) = mem.get(&i) {
            if let Some(k) = j.get(&cap) {
                return *k;
            }
        }

        let mut b = false;
        if cap == 0 {
            b = true;
        } else if i >= 1 {
            b = Solution::f(nums, i-1, cap, mem) || Solution::f(nums, i-1, cap-nums[i-1], mem)
        }

        mem.entry(i).or_insert_with(HashMap::new);
        if let Some(mem2) = mem.get_mut(&i) {
            mem2.insert(cap, b);
        }
        return b;
        
    }
}
