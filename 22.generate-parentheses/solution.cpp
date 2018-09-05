#include <iostream>
#include <vector>
using namespace std;


class Solution {
private:
    void addBr(vector<string> *result, int lLeft, int rLeft, string currentStr) {
        if(rLeft==0 && lLeft==0) {
            (*result).push_back(currentStr);
            return;
        }
        if(lLeft==0) {
            addBr(result, 0, rLeft-1, currentStr+")");
            return;
        }
        if(lLeft < rLeft) {
            addBr(result, lLeft-1, rLeft, currentStr+"(");
            addBr(result, lLeft, rLeft-1, currentStr+")");
        } else if (lLeft == rLeft) {
            addBr(result, lLeft-1, rLeft, currentStr+"(");
        }
    }

public:
    vector<string> generateParenthesis(int n) {
        vector<string> *result = new vector<string>;
        int lLeft = n;
        int rLeft = n;
        string currentStr("(");
        addBr(result, lLeft - 1, rLeft, currentStr);
        return *result;
    }
};
