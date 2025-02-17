class Solution {
    public int findSpecialInteger(int[] arr) {
        int continues = 0;
        int v = (int) Math.floor(arr.length * 0.25 + 1);
        for (int i = 0; i < arr.length; i++) {
            if (i == 0 || arr[i] == arr[i-1]) {
                continues ++;
            } else {
                continues = 1;
            }
            if (continues >= v) {
                return arr[i];
            }
        }
        return 0;
    }
}
