class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        return list(map(lambda x: x >= max(candies)-extraCandies, candies))
