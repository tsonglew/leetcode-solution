/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        return this.mergeSection(lists, 0, lists.length);
    }

    private ListNode mergeSection(ListNode[] lists, int from, int to) {
        if (from == to)
            return null;
        if (from + 1 == to)
            return lists[from];
        ListNode left = this.mergeSection(lists, from, (int) Math.floor((from + to) / 2));
        ListNode right = this.mergeSection(lists, (int) Math.floor((from + to) / 2), to);
        ListNode head = new ListNode();
        ListNode headC = head;
        ListNode leftC = left;
        ListNode rightC = right;
        while (!(leftC == null && rightC == null)) {
            ListNode picked = null;
            if (leftC == null) {
                picked = rightC;
                rightC = rightC.next;
            } else if (rightC == null) {
                picked = leftC;
                leftC = leftC.next;
            } else {
                if (leftC.val <= rightC.val) {
                    picked = leftC;
                    leftC = leftC.next;
                } else {
                    picked = rightC;
                    rightC = rightC.next;
                }
            }
            picked.next = null;
            headC.next = picked;
            headC = headC.next;
        }

        return head.next;
    }

    private void iterNode(ListNode head) {
        ListNode n = head;
        while (n != null) {
            System.out.printf("%d ", n.val);
            n = n.next;
        }
        System.out.println();
    }
}