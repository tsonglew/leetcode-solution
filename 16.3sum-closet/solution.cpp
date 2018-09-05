#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        int closet = INT_MAX;
        int msum;
        int length = static_cast<int>(nums.size());
        sort(nums.begin(), nums.end());
        for(int i=0; i<length-2; i++) {
            int head = i+1;
            int tail = length-1;
            while(head<tail) {
                int sum = nums[i] + nums[head] + nums[tail];
                int minus = sum - target;
                int abs_minus = abs(minus);
                if(abs_minus < closet) {
                    closet = abs_minus;
                    msum = sum;
                }
                if(sum < target){
                    head++;
                } else if(sum > target){
                    tail--;
                } else {
                    head++;
                    tail--;
                }
            }
        }
        return msum;
    }
};
