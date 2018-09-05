#include <iostream>
#include <vector>
using namespace std;


class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int count = 0;
        for(int i=0; i<nums.size(); i++) {
            count++;
            while(nums[i+1]==nums[i]) {
                i++;
            }
            nums[count] = nums[i+1];
        }
        return count;
    }
};
