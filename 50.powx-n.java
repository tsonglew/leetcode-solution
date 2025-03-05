class Solution {
    public double myPow(double x, int n) {
        if (n == 0)
            return 1;
        if (n == -1 << 31)
            return 1 / x * this.myPow(x, n + 1);
        if (n < 0)
            return 1 / this.myPow(x, -n);
        double subVal = this.myPow(x, (int) Math.floor(n / 2));
        return n % 2 == 1 ? x * subVal * subVal : subVal * subVal;
    }
}