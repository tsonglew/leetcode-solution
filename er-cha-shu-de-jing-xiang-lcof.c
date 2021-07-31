/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */


struct TreeNode* mirrorTree(struct TreeNode* root){
    if (root == NULL)
        return root;
    struct TreeNode* c = root->left;
    root->left = root->right;
    root->right = c;
    mirrorTree(root->left);
    mirrorTree(root->right);
    return root;
}
