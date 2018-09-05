#include <iostream>
using namespace std;

/**
 * Definition for singly-linked list.
**/
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode *front = head;
        ListNode *front2 = head;
        while(n) {
            front2 = front2->next;
            n--;
        }
        if(front2==NULL) {
            return head->next;
        }
        while(front2->next!=NULL) {
            front = front->next;
            front2 = front2->next;
        }
        front->next = front->next->next;
        return head;
    }
};
