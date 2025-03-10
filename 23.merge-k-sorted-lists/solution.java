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

    static class Heap {
        ListNode[] mem;
        int idx;

        public Heap(int size) {
            this.mem = new ListNode[size];
            this.idx = 1;
        }

        public boolean isEmpty() {
            return this.idx <= 1;
        }

        public void push(ListNode n) {
            this.mem[this.idx++] = n;
            this.popup();
        }

        public ListNode pop() {
            if (this.idx <= 1) {
                return null;
            }
            ListNode swap = this.mem[1];
            this.mem[1] = this.mem[this.idx - 1];
            this.mem[this.idx - 1] = swap;
            this.mem[--this.idx] = null;
            this.popdown();
            return swap;
        }

        private void popup() {
            int i = this.idx - 1;
            int pi = (int) Math.floor(i / 2);
            while (i > 0 && pi > 0 && this.mem[pi].val > this.mem[i].val) {
                ListNode swap = this.mem[pi];
                this.mem[pi] = this.mem[i];
                this.mem[i] = swap;
                i = pi;
                pi = (int) Math.floor(i / 2);
            }
        }

        private void popdown() {
            int i = 1;
            int cil = i * 2;
            int cir = i * 2 + 1;
            while (i < this.idx && (cil < this.idx || cir < this.idx)) {
                boolean useLeftChild = false;
                boolean useRightChild = false;
                if (cil < this.idx && this.mem[i].val > this.mem[cil].val) {
                    useLeftChild = true;
                }
                if (cir < this.idx && this.mem[i].val > this.mem[cir].val) {
                    useRightChild = true;
                }
                if (useLeftChild && useRightChild) {
                    if (this.mem[cil].val <= this.mem[cir].val) {
                        useRightChild = false;
                    } else {
                        useLeftChild = false;
                    }
                }
                if (!useLeftChild && !useRightChild) {
                    return;
                }
                if (useLeftChild) {
                    this.swap(i, cil);
                    i = cil;
                } else {
                    this.swap(i, cir);
                    i = cir;
                }
                cil = i * 2;
                cir = i * 2 + 1;
            }
        }

        private void swap(int i, int j) {
            ListNode s = this.mem[i];
            this.mem[i] = this.mem[j];
            this.mem[j] = s;
        }

        public void print() {
            System.out.printf("Heap: ");
            for (int i = 1; i < this.idx; i++) {
                System.out.printf("%d ", this.mem[i].val);
            }
            System.out.println();
        }

    }

    public ListNode mergeKLists(ListNode[] lists) {
        Heap h = new Heap(lists.length + 1);
        for (ListNode n : lists) {
            if (n == null) {
                continue;
            }
            h.push(n);
        }

        ListNode head = new ListNode();
        ListNode tail = head;
        while (!h.isEmpty()) {
            ListNode n = h.pop();
            tail.next = n;
            tail = tail.next;
            if (n.next != null) {
                h.push(n.next);
            }
        }
        return head.next;
    }
}