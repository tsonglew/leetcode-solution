use std::cmp;

impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let mut state = vec![0; cost.len()];
        state[0] = cost[0];
        state[1] = cost[1];
        for i in 2..state.len() {
            state[i] = cmp::min(state[i - 1], state[i - 2]) + cost[i];
        }
        return cmp::min(state[cost.len() - 1], state[cost.len() - 2]);
    }
}
