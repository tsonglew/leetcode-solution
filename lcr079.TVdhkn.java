class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        if (nums.length == 0) {
            res.add(new ArrayList<>());
            return res;
        }
        List<List<Integer>> subset = this.subsets(Arrays.copyOfRange(nums, 1, nums.length));
        // System.out.printf("subset length: %d\n", subset.size());
        int subsetSize = subset.size();
        for (int i = 0; i < subsetSize; i++) {
            List<Integer> s1 = new ArrayList<>();
            s1.add(nums[0]);
            s1.addAll(subset.get(i));
            subset.add(s1);
        }
        // System.out.printf("res length: %d\n", res.size());
        return subset;
    }
}