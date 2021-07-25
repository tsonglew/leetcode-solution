/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* getKthFromEnd(struct ListNode* head, int k){
    int cnt = 0;
    struct ListNode* p = head;
    while (p != NULL){
        p = p->next;
        cnt ++;
    }

    p = head;
    int num = cnt-k;
    while (num--)
        p = p->next;
    return p;
}
