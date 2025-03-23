class Solution {
    public int findPeakElement(int[] nums) {
        int i = 0, j = nums.length - 1;
        while (i < j) {
            int mid = (int)Math.floor((i+j)/2);
            if (nums[mid] < nums[mid+1]){
                i = mid+1;
            } else {
                j = mid;
            }
        }
        return j;
    }
}
