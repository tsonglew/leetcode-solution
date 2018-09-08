public class Solution {
    public double myPow(double x, int n) {
        return pow(x, n);
    }
    public double pow(double x, int n) {
        if(n == 0)
            return 1;
        if(n<0){
            if (n == -2147483648) {
                n = 2147483647;
                x = 1/(x*x);
            } else {
                n = -n;
                x = 1/x;
            }
        }
        return (n%2 == 0) ? pow(x*x, n/2) : x*pow(x*x, n/2);
    }
}