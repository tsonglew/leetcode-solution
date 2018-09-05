#include <iostream>
using namespace std;


struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};


class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        ListNode *headptr = new ListNode(0);
        headptr->next = head;
        ListNode *hp = headptr;
        bool getHead = false;
        while(hp->next && hp->next->next) {
            ListNode *tmp = hp->next->next;
            hp->next->next = hp->next->next->next;
            tmp->next = hp->next;
            if(!getHead) {
                headptr->next = tmp;
                getHead = true;
            } else {
                hp->next = tmp;
            }
            hp = tmp->next;
        }
        return headptr->next;
    }
};
