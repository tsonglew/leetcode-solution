#include <iostream>
#include <vector>
#include <deque>
using namespace std;


class Solution {
public:
    vector<string> letterCombinations(string digits) {
        int inputLen = static_cast<int>(digits.length());
        if(!inputLen) {
            return vector<string>();
        }
        deque<string> result;
        result.push_back("");
        vector<string> map = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};
        for(int i=0; i<inputLen; i++) {
            int inputVal = digits[i] - 48;
            string addStr = map[inputVal];
            int addStrLen = static_cast<int>(addStr.length());
            int resultSize = static_cast<int>(result.size());
            for(int j=0; j<resultSize; j++) {
                string firStr = result[0];
                result.pop_front();
                for(int k=0; k<addStrLen; k++) {
                    result.push_back(firStr+addStr[k]);
                }
            }
        }
        vector<string> strRes;
        int resSize = static_cast<int>(result.size());
        for(int i=0; i<resSize; i++) {
            strRes.push_back(result[i]);
        }
        return strRes;
    }
};
