class Solution {
    public int maximalRectangle(char[][] matrix) {
        int res = 0;
        int[][] columns = new int[matrix.length][matrix[0].length];
        for (int i = columns.length - 1; i >= 0; i--) {
            for (int j = columns[0].length - 1; j >= 0; j--) {
                if (matrix[i][j] == '0') {
                    columns[i][j] = 0;
                } else {
                    if (i == columns.length - 1) {
                        columns[i][j] = matrix[i][j] == '1' ? 1 : 0;
                    } else {
                        columns[i][j] = matrix[i][j] == '0' ? 0 : columns[i + 1][j] + 1;
                    }
                }
            }
        }
        for (int i = 0; i < columns.length; i++) {
            res = Math.max(res, this.largestRectangleArea(columns[i]));
        }
        return res;
    }

    public int largestRectangleArea(int[] heights) {
        Stack<Rect> s = new Stack();
        int res = 0;
        for (int i = 0; i < heights.length; i++) {
            int idx = i;
            while (!s.isEmpty() && heights[i] <= s.peek().height) {
                Rect ti = s.pop();
                idx = ti.index;
                res = Math.max(res, (i - ti.index) * ti.height);
            }
            if (heights[i] > 0) {
                s.push(new Rect(idx, heights[i]));
            }

        }
        while (!s.isEmpty()) {
            Rect ti = s.pop();
            res = Math.max(res, (heights.length - ti.index) * ti.height);
        }
        return res;
    }

    class Rect {
        int index;
        int height;

        public Rect(int index, int height) {
            this.index = index;
            this.height = height;
        }
    }
}