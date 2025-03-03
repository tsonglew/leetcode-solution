class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        if (nums.length == 0) {
            res.add(new ArrayList<>());
            return res;
        }
        List<List<Integer>> subset = this.subsets(Arrays.copyOfRange(nums, 1, nums.length));
        // System.out.printf("subset length: %d\n", subset.size());

        for (int i = 0; i < subset.size(); i++) {
            List<Integer> s1 = new ArrayList<>();
            List<Integer> s2 = new ArrayList<>();
            s1.add(nums[0]);
            s1.addAll(subset.get(i));
            s2.addAll(subset.get(i));
            res.add(s1);
            res.add(s2);
        }
        // System.out.printf("res length: %d\n", res.size());
        return res;
    }
}