typedef struct NumArrayNode {
    int nums[512];
    int length;
    struct NumArrayNode *next;
} NumArrayNode;

typedef struct {
    NumArrayNode *head;
} NumArray;


NumArray* numArrayCreate(int* nums, int numsSize) {
    int i;
    int sum;

    sum = 0;
    NumArray *na = (NumArray*)malloc(sizeof(NumArray));
    na->head = (NumArrayNode*)malloc(sizeof(NumArrayNode));
    NumArrayNode *node = na->head;
    node->length = 0;
    
    for (i=0; i<numsSize; i++) {
        sum += nums[i];
        if (node->length < 512) {
            node->nums[i%512] = sum;
            node->length ++;
        } else{
            node->next = (NumArrayNode*)malloc(sizeof(NumArrayNode));
            node = node->next;
            node->length = 0;
        }
    }
    return na;
}

int numArraySumRange(NumArray* obj, int left, int right) {
    int left_group = left == 0? -1:(left-1) / 512;
    int left_offset = (left-1) % 512;
    int right_group = right / 512;
    int right_offset = right % 512;
    int i;
    NumArrayNode *node = obj->head;
    NumArrayNode *right_node = obj->head;
    for (i = 0; i < right_group; ++i) {
        if (i < left_group) {
            node = node->next;
        }
        right_node = right_node->next;
    }
    int left_value = left_group == -1? 0: node->nums[left_offset];
    int right_value = right_node->nums[right_offset];
    return right_value - left_value;
}

void numArrayFree(NumArray* obj) {

}

/**
 * Your NumArray struct will be instantiated and called as such:
 * NumArray* obj = numArrayCreate(nums, numsSize);
 * int param_1 = numArraySumRange(obj, left, right);
 
 * numArrayFree(obj);
*/
