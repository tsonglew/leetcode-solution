int *singleNumber(int *nums, int numsSize, int *returnSize)
{
    int xor = 0;
    for (int i = 0; i < numsSize; ++i)
        xor ^= nums[i];
    int n = 0;
    while ((xor&1) == 0)
    {
        n++;
        xor >>= 1;
    }
    int a = 0, b = 0;
    for (int i = 0; i < numsSize; ++i)
    {
        if ((nums[i] >> n & 1) == 0)
            a ^= nums[i];
        else
            b ^= nums[i];
    }
    int *res = malloc(sizeof(int) * 2);
    *returnSize = 2;
    res[0] = a;
    res[1] = b;
    return res;
}
