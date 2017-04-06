#include <iostream>
using namespace std;


class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        ListNode *newhead = head;
        ListNode *formertail = NULL;
        bool gethead = false;

        while(true) {

            if(!head) {
                break;
            }
            ListNode *tail = head;
            int count = k-1;
            int count2 = k;
            while(count && tail->next) {
                tail = tail->next;
                count--;
            }
            if(formertail) {
                formertail->next = tail;
                if(count>0 || !tail) {
                    formertail->next = head;
                }
            }
            formertail = head;
            if(count>0 || !tail) {
                break;
            }
            ListNode *former = tail->next;
            while(count2) {
                if(!gethead) {
                    newhead = tail;
                    gethead = true;
                }
                ListNode *hp = head->next;

                head->next = former;
                former = head;
                head = hp;
                count2--;
            }
        }
        return newhead;
    }
};
