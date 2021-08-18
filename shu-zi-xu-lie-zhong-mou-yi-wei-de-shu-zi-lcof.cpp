class Solution {
public:
    int findNthDigit(int n) {
        int numSection = 1;
        long numSectionLowerBound = 0;
        while (1) {
            long nexLowerBound = 0;
            if (numSection == 1) {
                nexLowerBound = 10;
            } else if (numSection == 2){
                nexLowerBound = 9 * 10 * 2;
            } else {
                nexLowerBound = numSectionLowerBound * 10 * numSection / (numSection-1);
            }
            if (n > nexLowerBound) {
                n -= nexLowerBound;
                numSection ++;
                numSectionLowerBound = nexLowerBound;
            } else {
                break;
            }
        }
        int numSectionBase = numSection>1 ? pow(10, numSection-1) : 0; 
        int cnt = n / numSection;
        int rem = n % numSection;
        return to_string(cnt+numSectionBase)[rem]-48;
    }
};
