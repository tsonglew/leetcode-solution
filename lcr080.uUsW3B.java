class Solution {
    public List<List<Integer>> combine(int n, int k) {
        int[] nums = new int[n];
        // System.out.println(n);
        for (int i = 0; i < nums.length; i++) {
            nums[i] = i + 1;
        }
        return this.sub(nums, k);
    }

    public List<List<Integer>> sub(int[] nums, int k) {
        // for (int i = 0; i < nums.length; i++) {
        // System.out.printf("%d ", nums[i]);
        // }
        // System.out.println();
        List<List<Integer>> res = new ArrayList<>();
        if (k == 0) {
            res.add(new ArrayList<>());
            return res;
        }
        if (k == nums.length) {
            res.add(new ArrayList<>(Arrays.stream(nums).boxed().collect(Collectors.toList())));
            return res;
        }
        List<List<Integer>> subset1 = this.sub(Arrays.copyOfRange(nums, 1, nums.length), k);
        List<List<Integer>> subset2 = this.sub(Arrays.copyOfRange(nums, 1, nums.length), k - 1);
        int subsize = subset2.size();
        for (int i = 0; i < subsize; i++) {
            subset2.get(i).add(nums[0]);
            subset1.add(subset2.get(i));
        }
        return subset1;
    }
}