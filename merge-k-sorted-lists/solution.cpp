#include <iostream>
#include <vector>
using namespace std;


struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};


class Solution {
private:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *r, *newHead = new ListNode(0);
        newHead->next = NULL;
        r = newHead;
        while(l1&&l2) {
            if(l1->val<l2->val) {
                newHead->next = l1;
                l1 = l1->next;
            } else {
                newHead->next = l2;
                l2 = l2->next;
            }
            newHead = newHead->next;
        }
        if(l1==NULL) {
            newHead->next = l2;
        } else if(l2==NULL) {
            newHead->next = l1;
        }
        return r->next;
    }
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if(lists.size()==0) {
            return NULL;
        }
        while(lists.size() > 1) {
            lists.push_back(mergeTwoLists(lists[0], lists[1]));
            lists.erase(lists.begin(), lists.begin()+2);
        }
        return lists[0];
    }
};
