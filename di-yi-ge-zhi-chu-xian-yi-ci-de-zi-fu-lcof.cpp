class Solution {
public:
    char firstUniqChar(string s) {
        std::unordered_map<char, int> chars;
        for (auto c : s) {
            chars[c]++;
        }
        for (auto c : s) {
            if (chars[c] == 1)
                return c;
        }
        return ' ';
    }
};
