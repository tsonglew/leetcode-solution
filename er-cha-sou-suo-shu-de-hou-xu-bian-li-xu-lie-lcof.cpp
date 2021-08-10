class Solution {
public:
    bool verifyPostorder(vector<int>& postorder) {
        return verify(postorder, postorder.begin(), postorder.end());
    }

    bool verify(vector<int>& postorder, std::vector<int>::iterator begin, std::vector<int>::iterator end) {
        if (end - begin <= 2) 
            return true;
        auto root = *(end-1);
        auto p = begin;
        while (*p < root) {
            p++;
        }
        auto left_end = p;
        while (*p != root) {
            if (*p < root)
                return false;
            p++;
        }
        return verify(postorder, begin, left_end) && verify(postorder, left_end, p);
    }
};
