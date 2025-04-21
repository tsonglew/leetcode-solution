class Solution {
    public int numberOfArrays(int[] differences, int lower, int upper) {
        long releventMax = 0, releventMin =0;
        long relevent = 0;
        for (int diff : differences) {
            relevent += (long) diff;
            releventMax = Math.max(relevent, releventMax);
            releventMin = Math.min(relevent, releventMin);
        }
        // System.out.printf("max: %d, min: %d\n", releventMax, releventMin);
        return (int)Math.max(0, (upper - lower) - (releventMax - releventMin) + 1);
    }
}
