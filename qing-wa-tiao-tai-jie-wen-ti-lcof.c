int numWays(int n){
    int a = 1, b = 1, c;
    for (int i=2; i<=n; ++i) {
        c = b;
        b = (a+b)%1000000007;
        a = c;
    }
    return b;
}
