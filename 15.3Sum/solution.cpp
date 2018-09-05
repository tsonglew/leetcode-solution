#include <iostream>
#include <vector>
#include <algorithm>

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
        sort(nums.begin(), nums.end());
        int length = static_cast<int>(nums.size());
        if(length==0){
            return result;
        }
        for(int i=0; nums[i]<=0&&i<length; i++) {
            int head = i + 1;
            int tail = length - 1;
            int tofind = -nums[i];
            while(head<tail) {
                int sum = nums[head] + nums[tail];
                if(sum == tofind) {
                    vector<int> con;
                    con.push_back(nums[i]);
                    con.push_back(nums[head]);
                    con.push_back(nums[tail]);
                    result.push_back(con);
                    while(nums[head]==nums[head+1]) {
                        head++;
                    }
                    while(nums[tail]==nums[tail-1]) {
                        tail--;
                    }
                    head++;
                    tail--;
                } else if (sum < tofind) {
                    head++;
                } else if (sum > tofind) {
                    tail--;
                }
            }
            while(nums[i]==nums[i+1]){
                i++;
            }
        }
        return result;
    }
};
