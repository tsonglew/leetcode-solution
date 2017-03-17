#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    string convert(string s, int numRows) {
        string result;
        int sec_size = 2*numRows - 2;
        if(sec_size==0) {
            return s;
        }
        int str_length = static_cast<int>(s.length());
        int sec_nums = str_length / sec_size + 1;
        int sec_count = 0;
        int m = 0;
        while(m < numRows) {
            sec_count = 0;
            while(sec_count<sec_nums){
                int first = sec_count*sec_size + m;
                int last = (sec_count+1)*sec_size - m;
                if (first < str_length) {
                    result.push_back(s[first]);
                }
                if (last<(sec_count+1)*sec_size && last < str_length && first!=last) {
                    result.push_back(s[last]);
                }
                sec_count++;
            }
            m++;
        }
        return result;
    }
};
