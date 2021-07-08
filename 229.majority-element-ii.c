
/**
  * Note: The returned array must be malloced, assume caller calls free().
  */
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *majorityElement(int *nums, int numsSize, int *returnSize)
{
    int *cand = malloc(sizeof(int) * 2);
    int *cnt = malloc(sizeof(int) * 2);
    cand[0] = cand[1] = cnt[0] = cnt[1] = 0;

    for (int i = 0; i < numsSize; ++i)
    {
        if (cnt[0] > 0 && cand[0] == nums[i])
        {
            cnt[0]++;
            continue;
        }
        if (cnt[1] > 0 && cand[1] == nums[i])
        {
            cnt[1]++;
            continue;
        }

        if (cnt[0] == 0)
        {
            cand[0] = nums[i];
            cnt[0] = 1;
            continue;
        }
        if (cnt[1] == 0)
        {
            cand[1] = nums[i];
            cnt[1] = 1;
            continue;
        }
        cnt[0]--;
        cnt[1]--;
    }

    cnt[0] = cnt[1] = 0;
    for (int i = 0; i < numsSize; ++i)
    {
        if (nums[i] == cand[0])
            cnt[0]++;
        else if (nums[i] == cand[1])
            cnt[1]++;
    }
    *returnSize = 0;
    if (cnt[0] > numsSize / 3)
        (*returnSize)++;
    if (cnt[1] > numsSize / 3)
        (*returnSize)++;
    if (*returnSize == 1)
    {
        if (cnt[0] <= numsSize / 3)
        {
            return cand + 1;
        }
    }
    return cand;
}
