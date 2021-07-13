/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */


struct TreeNode* buildTree(int* preorder, int preorderSize, int* inorder, int inorderSize){
    if (preorderSize==0) {
        return NULL;
    }
    struct TreeNode* head = malloc(sizeof(struct TreeNode));
    head->val = preorder[0];
    int i;
    for (i=0; i<inorderSize && inorder[i]!=preorder[0]; ++i) {}
    head->left = buildTree(preorder+1, i,inorder, i);
    head->right = buildTree(preorder+1+i, preorderSize-i-1, inorder+i+1, preorderSize-i-1);
    return head;
}
