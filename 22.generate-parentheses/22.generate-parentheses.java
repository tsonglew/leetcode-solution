class Solution {
    HashMap<Integer, List<String>> cache;

    public Solution() {
        this.cache = new HashMap<>();
        this.cache.put(0, List.of(""));
        this.cache.put(-1, List.of(""));
    }
    public List<String> generateParenthesis(int n) {
        if (this.cache.containsKey(n)) {
            return this.cache.get(n);
        }
        ArrayList<String> ans = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            // (A)B
            List<String> A = this.generateParenthesis(i);
            List<String> B = this.generateParenthesis(n-i-1);
            // System.out.println(i);
            for (int j = 0; j < A.size(); j++) {
                for (int k = 0; k < B.size(); k++) {
                    // System.out.printf("%d, %d\n", j, k);
                    ans.add("(" + A.get(j) + ")" + B.get(k));
                }
            }
        }
        this.cache.put(n, ans);
        return ans;
    }
}
