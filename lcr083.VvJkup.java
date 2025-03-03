class Solution {

    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList();
        if (nums.length == 0) {
            res.add(new ArrayList<>());
            return res;
        }
        for (int i = 0; i < nums.length; i++) {
            int n = nums[i];
            int[] subArr = new int[nums.length - 1];
            System.arraycopy(nums, 0, subArr, 0, i);
            System.arraycopy(nums, i + 1, subArr, i, nums.length - 1 - i);
            List<List<Integer>> sub = this.permute(subArr);
            for (int j = 0; j < sub.size(); j++) {
                sub.get(j).add(n);
                res.add(sub.get(j));
            }
        }
        return res;
    }

}