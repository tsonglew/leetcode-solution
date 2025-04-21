class Solution {
    public int numRabbits(int[] answers) {
        Map<Integer, Integer> m = new HashMap<>();
        int ans = 0;
        for (int answer : answers) {
            int maxNum = answer + 1;
            m.compute(maxNum, (k, v) -> v == null ? 1 : v + 1);
            if (m.get(maxNum) == maxNum) {
                m.put(maxNum, 0);
                ans += maxNum;
            }
        }
        for (Map.Entry<Integer, Integer> e : m.entrySet()) {
            if (e.getValue() == 0) {
                continue;
            }
            ans += e.getKey();
        }
        return ans;
    }
}
