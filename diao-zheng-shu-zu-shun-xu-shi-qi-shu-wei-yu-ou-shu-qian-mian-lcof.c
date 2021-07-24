/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* exchange(int* nums, int numsSize, int* returnSize){
    int i = 0, j = numsSize-1, s;
    while (i<j) {
        while(nums[i]%2==1 && i < j) 
            i++;
        while(nums[j]%2==0 && j > i)
            j--;
        if (i >= j)
            break;
        s = nums[i];
        nums[i] = nums[j];
        nums[j] = s;
    }
    *returnSize = numsSize;
    return nums;
}
