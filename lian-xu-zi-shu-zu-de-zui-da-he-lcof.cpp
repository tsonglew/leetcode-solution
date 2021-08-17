class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int n = 0, maxSum = nums[0];
        for (const auto& i : nums) {
            n = max(i, i+n);
            maxSum = max(n, maxSum);
        }
        return maxSum;
    }
};
