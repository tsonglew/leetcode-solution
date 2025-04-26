class Solution {
    public int countCompleteSubarrays(int[] nums) {
        Set<Integer> uniqueNums = new HashSet<>();
        for (int num : nums) {
            uniqueNums.add(num);
        }
        int total = uniqueNums.size();

        int ans = 0;
        int startP = 0, endP = 0;
        Map<Integer, Integer> numsCount = new HashMap<>();
        // System.out.printf("total: %d\n", total);
        while (endP <= nums.length) {
            if (numsCount.size() == total) {
                ans += (nums.length - endP + 1);
                // System.out.printf("%d %d\n", startP, endP);
                if (numsCount.get(nums[startP]) == 1) {
                    numsCount.remove(nums[startP]);
                } else {
                    numsCount.put(nums[startP], numsCount.get(nums[startP]) - 1);
                }
                startP ++;
            } else {
                if (endP == nums.length) {
                    break;
                }
                numsCount.compute(nums[endP], (k, v) -> v == null ? 1 : v + 1);
                endP ++; 
            }
        }
        return ans;
    }
}
