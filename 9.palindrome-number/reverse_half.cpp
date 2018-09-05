#include <iostream>
using namespace std;

class Solution {
public:
    bool isPalindrome(int x) {
        if(x < 0 || (!(x%10) && x!=0)) {
            return false;
        }
        int half = 0;
        while(half < x) {
            half = (half*10 + x%10);
            x /= 10;
        }
        if(x == half || x==half/10) {
            return true;
        }
        return false;
    }
};
