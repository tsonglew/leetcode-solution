/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* oddEvenList(struct ListNode* head) {
	struct ListNode* odd_head = malloc(sizeof(struct ListNode));
	struct ListNode* eve_head = malloc(sizeof(struct ListNode));
	odd_head->val = 0;
	odd_head->next = NULL;
	eve_head->val = 0;
	eve_head->next = NULL;
	struct ListNode* odd_p = odd_head;
	struct ListNode* eve_p = eve_head;
	int is_node_odd = 1;
	struct ListNode* p = head;
	while (p != NULL) {
		if (is_node_odd) {
			odd_p->next = p;
			p = p->next;
			odd_p = odd_p->next;
			odd_p->next = NULL;

			is_node_odd = 0;
			}
		else {
			eve_p->next = p;
			p = p->next;
			eve_p = eve_p->next;
			eve_p->next = NULL;
			is_node_odd = 1;
		}
	}
	odd_p->next = eve_head->next;
	return odd_head->next;
}
