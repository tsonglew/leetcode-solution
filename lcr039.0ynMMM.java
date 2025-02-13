class Solution {
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