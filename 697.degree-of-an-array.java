class Solution {
    static class Degree {
        int val;
        int firstIdx;
        int lastIdx;
        int maxCnt;

        public Degree(int val, int firstIdx, int lastIdx) {
            this.val = val;
            this.firstIdx = firstIdx;
            this.lastIdx = lastIdx;
            this.maxCnt = 0;
        }
    }

    public int findShortestSubArray(int[] nums) {
        int maxCnt = 0;
        Set<Integer> maxCntVal = new HashSet<>();
        HashMap<Integer, Degree> m = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (!m.containsKey(nums[i])) {
                m.put(nums[i], new Degree(nums[i], i, i));
            }
            Degree d = m.get(nums[i]);
            d.maxCnt += 1;
            d.lastIdx = i;
            if (d.maxCnt > maxCnt) {
                maxCnt = d.maxCnt;
                maxCntVal.clear();
            }
            if (d.maxCnt >= maxCnt) {
                maxCntVal.add(nums[i]);
            }
        }
        // m.forEach((k, v) -> {
        // System.out.printf("k: %d, v: %d\n", k, v.maxCnt);
        // });
        int res = 50000;
        Iterator<Integer> it = maxCntVal.iterator();
        while (it.hasNext()) {
            Degree d = m.get(it.next());
            res = Math.min(res, d.lastIdx - d.firstIdx + 1);
        }

        return res;
    }
}