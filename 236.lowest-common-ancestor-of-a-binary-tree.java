/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode(int x) { val = x; }
 * }
 */
class Solution {

    static class Triple {
        TreeNode root;
        boolean p;
        boolean q;

        public Triple() {
            this.root = null;
            this.p = false;
            this.q = false;
        }

        public Triple(TreeNode root, boolean p, boolean q) {
            this.root = root;
            this.p = p;
            this.q = q;
        }

        public boolean found() {
            return p && q;
        }
    }

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        Triple ans = this.f(root, p, q);
        return ans.root;
    }

    private Triple f(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null) {
            return new Triple();
        }
        boolean isP = false;
        boolean isQ = false;
        if (root.equals(p)) {
            isP = true;
        }
        if (root.equals(q)) {
            isQ = true;
        }
        Triple left = this.f(root.left, p, q);
        if (left.found()) {
            return left;
        }
        Triple right = this.f(root.right, p, q);
        if (right.found()) {
            return right;
        }
        return new Triple(root, isP || left.p || right.p, isQ || left.q || right.q);
    }
}