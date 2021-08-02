/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */


bool equal(struct TreeNode* t1, struct TreeNode* t2) {
    if (t1 == NULL && t2 == NULL)
        return true;
    if (t1 == NULL || t2 == NULL || t1->val != t2->val)
        return false;
    return equal(t1->left, t2->right) && equal(t1->right, t2->left);
}
bool isSymmetric(struct TreeNode* root){
    if (root == NULL)
        return true;
    return equal(root->left, root->right);
}
