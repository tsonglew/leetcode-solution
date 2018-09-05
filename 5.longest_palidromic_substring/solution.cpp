#include <iostream>
using namespace std;

class Solution {
private:
    int max = 0, start = 0;
    void extendEven(string s, int i, int string_len){
        int k = i+1;
        int length_count = 0;
        while(i>=0&&k<=string_len-1&&s[i] == s[k]) {
            length_count += 2;
            i--;
            k++;
        }
        if(length_count>max) {
            max = length_count;
            start = i+1;
        }

    }
    void extendOdd(string s, int i, int string_len) {
        int length_count = 1;
        int carry = 0;
        while((i-carry)>=0&&(i+carry)<=string_len-1&&s[i+carry] == s[i-carry]) {
            length_count = carry*2 + 1;
            carry++;
        }
        if(length_count>max){
            max = length_count;
            start = i-carry+1;
        }
    }
public:
    string longestPalindrome(string s) {
        int i = 0;
        int len = static_cast<int>(s.length());
        if(len<2) return s;
        
        while(i<len-1){
            extendEven(s, i, len);
            extendOdd(s, i, len);
            i++;
        }
        return s.substr(start, max);
    }
};
