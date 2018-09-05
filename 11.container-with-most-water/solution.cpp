#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int maxArea(vector<int>& height) {
        int right = static_cast<int>(height.size())-1;
        int maxA = 0, left = 0;
        while(left<right){
            maxA = max(maxA, min(height[left], height[right])*(right-left));
            height[left]<=height[right]?left++:right--;
        }
        return maxA;
    }
};
