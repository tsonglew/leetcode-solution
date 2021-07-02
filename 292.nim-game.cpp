class Solution {
public:
    bool canWinNim(int n) {
        return (n-1)%4 != 3;
    }
};
