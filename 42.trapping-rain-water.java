import java.util.Stack;
// import java.io.*;

class C {
    int width;
    int height;

    public C(int width, int height) {
        this.width = width;
        this.height = height;
    }

    public String toString() {
        return String.format("C[width:%d,height:%d]", this.width, this.height);
    }
}

class Solution {
    public int trap(int[] height) {
        Stack<C> s = new Stack();
        int res = 0;
        for (int i = 0; i < height.length; i++) {
            // System.out.printf("i: %d, height: %d\n", i, height[i]);
            int heightI = height[i];
            int accWidth = 0;
            while (!s.isEmpty() && s.peek().height <= heightI) {
                C top = s.pop();
                if (s.isEmpty()) {
                    break;
                }
                C preTop = s.peek();
                res += (Math.min(preTop.height, heightI) - top.height) * (accWidth + top.width);
                accWidth += top.width;
                // System.out.printf("update res: %d\n", res);
            }
            if (s.isEmpty() || s.peek().height > heightI) {
                s.push(new C(1 + accWidth, heightI));
            }
            // this.printStack(s);
        }
        return res;
    }

    public void printStack(Stack<C> s) {
        Iterator<C> i = s.iterator();
        while (i.hasNext()) {
            System.out.printf("%s,", i.next());
        }
        System.out.println("");
    }
}