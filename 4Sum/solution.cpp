#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;


class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        vector<vector<int>> result;
        int numsSize = static_cast<int>(nums.size());
        if(numsSize == 0) {
            return result;
        }
        sort(nums.begin(), nums.end());
        for(int i=0; i<numsSize-3; i++) {
            int t3 = target - nums[i];
            for(int j=i+1; j<numsSize-2; j++) {
                int t2 = t3 - nums[j];
                int front = j + 1;
                int rear = numsSize - 1;
                while(front < rear) {
                    int s2 = nums[front] + nums[rear];
                    if(s2 == t2) {
                        vector<int> ans;
                        ans.push_back(nums[i]);
                        ans.push_back(nums[j]);
                        ans.push_back(nums[front]);
                        ans.push_back(nums[rear]);
                        result.push_back(ans);
                        while(nums[front]==nums[front+1]) {
                        front++;
                        }
                        while(nums[rear]==nums[rear-1]) {
                            rear--;
                        }
                        front++;
                        rear--;
                    } else if(s2 < t2) {
                        front++;
                    } else if(s2 > t2) {
                        rear--;
                    }
                }
                while(nums[j] == nums[j+1]){
                    j++;
                }
            }
            while(nums[i] == nums[i+1]){
                i++;
            }
        }
        return result;
    }
};
