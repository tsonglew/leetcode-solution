import java.util.*;

class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        Deque<Integer> q = new LinkedList();
        int[] res = new int[nums.length - k + 1];
        for (int i = 0; i < nums.length; i++) {
            while (!q.isEmpty() && nums[q.peekLast()] < nums[i]) {
                q.removeLast();
            }
            q.addLast(i);
            if (q.peekFirst() <= i - k) {
                q.removeFirst();
            }
            // System.out.printf("i: %d\n", i);
            // Iterator<Integer> it = q.iterator();
            // while (it.hasNext()) {
            // System.out.printf("%d, ", it.next());
            // }
            // System.out.println();
            if (i - k < -1) {
                continue;
            }
            res[i - k + 1] = nums[q.peekFirst()];
        }
        return res;
    }
}