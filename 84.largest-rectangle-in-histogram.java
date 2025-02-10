import java.util.Stack;
import java.io.*;


class Column {
    public int index;
    public int height;

    public Column(int index, int height) {
        this.index = index;
        this.height = height;
    }

    public String toString(){ 
        return String.format("[index: %d, height: %d]", this.index, this.height);
    }
}

class Solution {
    public int largestRectangleArea(int[] heights) {
        int largest = 0;
        Stack<Column> s = new Stack();
        for (int i = 0; i < heights.length; i++) {
            // System.out.printf("%d\n", i);
            if (s.isEmpty()) {
                if (heights[i] > 0) {
                    s.push(new Column(i, heights[i]));
                }
                continue;
            }
            Column top = s.peek();
            if (heights[i] >= top.height) {
                s.push(new Column(i, heights[i]));
                continue;
            }
            int n = 0;
            while (!s.isEmpty()) {
                Column top2 = s.peek();
                if (top2.height <= heights[i]) {
                    break;
                }
                int newVal = top2.height * (i-top2.index);
                largest = largest > newVal ? largest : newVal;
                // System.out.printf("largest: %d, Column: %s\n", largest, top2);
                s.pop();
                n = top2.index;
            }
            if (heights[i] > 0) {
                s.push(new Column(n, heights[i]));
            }

        }
        // System.out.println("================");
        while (!s.isEmpty()) {            
            Column top = s.pop();
            int newVal = top.height * (heights.length - top.index);
            largest = largest > newVal ? largest : newVal;
            // System.out.printf("largest: %d, Column: %s\n", largest, top);
        }
        return largest;
    }
}