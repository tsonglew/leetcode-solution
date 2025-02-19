class Solution {
    public int maxDistance(List<List<Integer>> arrays) {
        int minVal = 20000;
        int maxVal = -20000;
        int res = 0;
        for (int i = 0; i < arrays.size(); i++) {
            int arrMin = arrays.get(i).get(0);
            int arrMax = arrays.get(i).get(arrays.get(i).size() - 1);
            res = Math.max(Math.max(arrMax - minVal, maxVal - arrMin), res);
            minVal = Math.min(arrMin, minVal);
            maxVal = Math.max(arrMax, maxVal);
        }
        return res;
    }
}