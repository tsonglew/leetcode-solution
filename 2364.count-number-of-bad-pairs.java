class Solution {
    public long countBadPairs(int[] nums) {
        Map<Long, Long> m = new HashMap<>();
        long ans = 0;
        for (long i = 0; i < nums.length; i++) {
            ans += i - m.getOrDefault(Long.valueOf(nums[(int)i])-i, 0L);
            m.compute(Long.valueOf(nums[(int)i])-i, (k, v) -> v == null ? 1L : v+1L);
        }
        return ans;
    }
}
