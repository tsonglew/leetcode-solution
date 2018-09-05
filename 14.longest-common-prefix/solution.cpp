class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        int size = strs.size();
        if(size<2){
            return size==1?strs[0]:string("");
        }
        string result;
        int i, j = 0;
        char first = strs[0][j];
        while(first!='\0') {
            i = 1;
            while(i<size) {
                if(strs[i][j]!=first){
                    return result;
                }
                i++;
            }
            result.push_back(first);
            j++;
            first = strs[0][j];
        }
        return result;
    }
};
