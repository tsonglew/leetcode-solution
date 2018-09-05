#include <iostream>
#include <vector>
using namespace std;

class Solution {
private:
    bool checkOverflow(int x, int y, int result) {
        if(x>=0&&y>=0){
            if(result<x||result<y){
                return true;
            }
        } else if((x>=0&&y<=0)||(x<=0&&y>=0)){
            if(result<x&&result<y){
                return true;
            }
        } else if(x<=0&&y<=0){
            if(result>x||result>y){
                return true;
            }
        }
        return false;
    }
public:
    int reverse(int x) {
        if (abs(x)<10) return x;

        vector<int> nums;
        bool negative;
        int absNum, eachBit, multiCount;
        signed int result, toPlus, temp;

        result = 0;
        negative = false;
        if (x<0) {
            negative = true;
        }
        absNum = abs(x);
        while(absNum){
            eachBit = absNum%10;
            nums.push_back(eachBit);
            absNum /= 10;
        }
        multiCount = 1;
        for(vector<int>::iterator it = nums.end()-1; it>=nums.begin(); it--){
            toPlus = (*it)*multiCount;
            if(toPlus%multiCount!=0){
                return 0;
            }
            temp = toPlus + result;
            if (checkOverflow(toPlus, result, temp)){
                return 0;
            }
            multiCount*=10;
            result = temp;
        }
        if(negative){
            return -result;
        }
        return result;
    }
};
