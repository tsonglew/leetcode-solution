/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */


bool compare(struct TreeNode* A, struct TreeNode* B) {
    if (B == NULL)
        return true;
    if (A != NULL && B != NULL && A->val == B->val)
        return compare(A->left, B->left) && compare(A->right, B->right);
    return false;

}
bool isSubStructure(struct TreeNode* A, struct TreeNode* B){
    if (A == NULL || B == NULL)
        return false;
    return compare(A, B) || isSubStructure(A->left, B) || isSubStructure(A->right, B);
}
