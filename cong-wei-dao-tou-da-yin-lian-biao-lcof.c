/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

void addInt(struct ListNode* head, int* returnSize, int* intHead) {
    if (head == NULL) {
        return;
    }
    addInt(head->next, returnSize, intHead);
    intHead[(*returnSize)++] = head->val;
}

int* reversePrint(struct ListNode* head, int* returnSize){
    *returnSize = 0;
    int* intHead = malloc(sizeof(int)*10000);
    addInt(head, returnSize, intHead);
    return intHead;
}

