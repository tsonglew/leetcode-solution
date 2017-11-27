class Solution {
public:
    int divide(int dividend, int divisor) {
        if ((divisor == 0) || (dividend==INT_MIN && divisor==-1)) {
            return INT_MAX;
        }
        long long f_dividend = labs(dividend);
        long long f_divisor = labs(divisor);
        int sign = ((dividend<0) ^ (divisor<0))?-1:1;
        if (f_dividend < f_divisor){
            return 0;
        } else {
            int t = 0;
            int m = 1;
            int count = 0;
            long long tmp = f_divisor;
            while(f_dividend!=0) {
                while((f_dividend-f_divisor) >= 0) {
                    f_dividend -= f_divisor;
                    f_divisor <<= 1;
                    t += m;
                    m += m;
                }
                count += t;
                m = 1;
                t = 0;
                f_divisor = tmp;
                if (f_dividend < f_divisor) {
                    break;
                }
            }
            return sign>0? count:-count;
        }
    }
};
