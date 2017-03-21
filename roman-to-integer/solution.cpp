#include <iostream>
#include <unordered_map>
using namespace std;

class Solution {
public:
    int romanToInt(string s) {
        unordered_map<char, int> T = {{'I', 1}, {'V', 5}, {'X', 10}, {'L', 50}, {'C', 100}, {'D', 500}, {'M', 1000}};
        int sum = 0;
        sum += T[s[static_cast<int>(s.length())-1]];
        for(string::iterator it=s.end()-2; it>=s.begin(); it--){
            sum += T[*it]<T[*(it+1)]?(-T[*it]):T[*it];
        }
        return sum;
    }
};
