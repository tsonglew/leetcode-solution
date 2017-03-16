#include <iostream>
using namespace std;

class Solution{
public:
    int myAtoi(string str) {
        int start = static_cast<int>(str.find_first_not_of(' '));
        int result = 0;
        int p = 1;
        if(str[start]=='-' || str[start]=='+'){
            p = str[start++]=='-'?-1:1;
        }
        while(str[start]>='0' && str[start]<='9') {
            if(result>INT_MAX/10 || (result==INT_MAX/10 && str[start]>'7')) {
                return p>0?INT_MAX:INT_MIN;
            }
            result = result*10 + str[start]-'0';
            start++;
        }
        return result*p;
    }
};
