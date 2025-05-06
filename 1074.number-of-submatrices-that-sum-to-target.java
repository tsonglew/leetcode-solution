class Solution {
    public int numSubmatrixSumTarget(int[][] matrix, int target) {
        int ans = 0;
        
        for (int t = 0; t < matrix.length; t++) {
            // t for top
            int[] row = new int[matrix[0].length];
            
            Map<Integer, Integer> m = new HashMap<>();

            for (int b = t; b < matrix.length; b++) {
                // b for bottom
                m.clear();
                int rowSum = 0;

                for (int c = 0; c < matrix[0].length; c++) {
                    // c for column

                    rowSum += matrix[b][c];

                    row[c] += rowSum;
                    
                    if (row[c] == target) {
                        ans += 1;
                    }

                    // row[c] - x = target => x = row[c] - target
                    ans += m.getOrDefault(row[c] - target, 0);
                    // System.out.printf("[t:%d]b: %d, c: %d: row[c]: %d, ans: %d\n", t, b, c, row[c], ans);
                    // this.iterRow(row);
                    // System.out.printf("matrix[b][c]: %d, row[c-1]: %d\n", matrix[b][c], (c == 0 ? 0 : row[c-1]));
                    m.compute(row[c], (k, v) -> v == null ? 1 : v + 1);
                }

            }
        }

        return ans;
    }

    private void iterRow(int[] row) {
        System.out.printf("row: ");
        for (int i : row) {
            System.out.printf("%d ", i);
        }
        System.out.println();
    }
}
