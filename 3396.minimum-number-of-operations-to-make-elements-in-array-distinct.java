class Solution {
    public int minimumOperations(int[] nums) {
        Set<Integer> numbers = new HashSet<>();
        int start = 0, end = 0, ans = 0;

        while (end < nums.length) {
            if (numbers.contains(nums[end])) {
                boolean removedDup = false;
                while (!removedDup && start < end) {
                    ans ++;
                    for (int i = 0; i < 3 && start < nums.length; i++) {
                        
                        int numToRm = nums[start];
                        numbers.remove(numToRm);
                        if (numToRm == nums[end]) {
                            removedDup = true;
                        }
                        start++;
                    }

                    end = Math.max(start, end);
                }
            } else {
                numbers.add(nums[end]);
                end++;
            }
        }
        return ans;

    }
}
