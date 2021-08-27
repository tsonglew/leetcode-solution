class Solution:
    def minNumber(self, nums: List[int]) -> str:
        return str(reduce(lambda x, y: str(x)+str(y), sorted(nums, key=cmp_to_key(lambda x, y: int(str(x)+str(y))-int(str(y)+str(x))))))
